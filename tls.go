package b3

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

var theStack *_Stack = newStack(1024 * 256)

type _Stack struct {
	Stack  []byte
	allocs []int
	pos    int

	// yields counts the calls to sched_yield, each goroutine of box3d has its own stack
	yields uint32
}

//go:noinline
func newStack(stackSize int) *_Stack {
	stack := &_Stack{
		Stack: make([]byte, stackSize),
	}

	escapes(stack)

	return stack
}

func (s *_Stack) CheckEmpty() {
	if s.pos != 0 || len(s.allocs) > 0 {
		panic("stack is not empty")
	}
}

func (s *_Stack) Alloc(n int) uintptr {
	if n%8 > 0 {
		// round up to a multiple of 8
		n += 8 - n%8
	}

	if s.pos+n > len(s.Stack) {
		panic("out of stack space")
	}

	buf := uintptr(unsafe.Pointer(unsafe.SliceData(s.Stack[s.pos:])))
	s.pos += n
	s.allocs = append(s.allocs, n)

	return buf
}

func (s *_Stack) Free(n int) uintptr {
	if n%8 > 0 {
		// round up to a multiple of 8
		n += 8 - n%8
	}

	lastAlloc := s.allocs[len(s.allocs)-1]
	if lastAlloc != n {
		panic("alloc/free is not paired")
	}

	s.allocs = s.allocs[:len(s.allocs)-1]
	s.pos -= n

	return 0
}

func (s *_Stack) toCString(src string) alloc {
	n := len(src) + 1
	mem := s.Alloc(n)

	buf := sliceOf(mem, size_t(n))
	copy(buf, src)
	buf[n-1] = 0

	return alloc{Addr: mem, tls: s, size: n}
}

type alloc struct {
	Addr uintptr
	tls  *_Stack
	size int
}

func (a alloc) Free() {
	a.tls.Free(a.size)
}

// copyToStack copies value to the box3d stack. The stack is a byte slice, the GC and the stack
// mover do not see pointers in it, so T must not contain any, see copyToStackAs.
func copyToStack[T any](tls *_Stack, value T) alloc {
	assertPointerFree[T]()

	n := int(unsafe.Sizeof(value))
	stack := tls.Alloc(n)
	*(*T)(unsafe.Pointer(stack)) = value

	return alloc{Addr: stack, tls: tls, size: n}
}

// copyToStackAs copies value to the box3d stack as C, a pointer-free mirror of T with the same
// layout. box3d sees the pointers of value as uintptr, the caller must keep value alive.
func copyToStackAs[C, T any](tls *_Stack, value T) alloc {
	if unsafe.Sizeof(*new(C)) != unsafe.Sizeof(value) {
		panic(fmt.Sprintf("copyToStackAs: %T and %T differ in size", *new(C), value))
	}

	// the pointees must not live on a goroutine stack, the stack mover does not update the copy
	escapes(value)

	return copyToStack(tls, *(*C)(unsafe.Pointer(&value)))
}

func sliceToStack[T any](tls *_Stack, values []T) alloc {
	assertPointerFree[T]()

	n := len(values) * int(unsafe.Sizeof(*new(T)))
	stack := tls.Alloc(n)
	copy(unsafe.Slice((*T)(unsafe.Pointer(stack)), len(values)), values)

	return alloc{Addr: stack, tls: tls, size: n}
}

// pointerFree caches assertPointerFree, reflect.Type -> bool
var pointerFree sync.Map

func assertPointerFree[T any]() {
	t := reflect.TypeFor[T]()
	ok, cached := pointerFree.Load(t)
	if !cached {
		ok, _ = pointerFree.LoadOrStore(t, !hasPointers(t))
	}

	if !ok.(bool) {
		panic(fmt.Sprintf("box3d stack: %s contains pointers", t))
	}
}

// hasPointers reports if values of t contain anything the GC or the stack mover must see.
func hasPointers(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Pointer, reflect.UnsafePointer, reflect.Map, reflect.Chan,
		reflect.Func, reflect.Interface, reflect.Slice, reflect.String:
		return true
	case reflect.Array:
		return t.Len() > 0 && hasPointers(t.Elem())
	case reflect.Struct:
		for i := range t.NumField() {
			if hasPointers(t.Field(i).Type) {
				return true
			}
		}
	}

	return false
}
