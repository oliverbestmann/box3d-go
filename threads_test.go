package b3

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Ported from box3d test/test_thread.c

func TestSemaphoreInitialCount(t *testing.T) {
	s := __builtin_b3CreateSemaphore(nil, 3)
	defer __builtin_b3DestroySemaphore(nil, s)

	// these do not block, the count starts at 3
	for range 3 {
		__builtin_b3WaitSemaphore(nil, s)
	}

	__builtin_b3SignalSemaphore(nil, s)
	__builtin_b3WaitSemaphore(nil, s)
}

func TestThreadSignalAndJoin(t *testing.T) {
	sem := __builtin_b3CreateSemaphore(nil, 0)
	defer __builtin_b3DestroySemaphore(nil, sem)

	var value int
	worker := threadFunction(func(_ *_Stack, _ uintptr) {
		value = 99
		__builtin_b3SignalSemaphore(nil, sem)
	})

	// __ccgo_fp hides the closure from the compiler and GC: move it to the heap and keep it alive
	escapes(worker)
	defer runtime.KeepAlive(worker)

	thread := __builtin_b3CreateThread(nil, __ccgo_fp(worker), 0, 0)
	__builtin_b3WaitSemaphore(nil, sem)
	assert.Equal(t, 99, value)

	__builtin_b3JoinThread(nil, thread)
}

func TestThreadMutex(t *testing.T) {
	mutex := __builtin_b3CreateMutex(nil)
	defer __builtin_b3DestroyMutex(nil, mutex)

	var sum int
	worker := threadFunction(func(_ *_Stack, _ uintptr) {
		for range 1000 {
			__builtin_b3LockMutex(nil, mutex)
			sum++
			__builtin_b3UnlockMutex(nil, mutex)
		}
	})

	// __ccgo_fp hides the closure from the compiler and GC: move it to the heap and keep it alive
	escapes(worker)
	defer runtime.KeepAlive(worker)

	var threads []uintptr
	for range 4 {
		threads = append(threads, __builtin_b3CreateThread(nil, __ccgo_fp(worker), 0, 0))
	}

	for _, thread := range threads {
		__builtin_b3JoinThread(nil, thread)
	}

	assert.Equal(t, 4000, sum)
}
