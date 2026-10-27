//go:build !tinygo

package b3

import "unsafe"

// funcPtr returns a pointer to a func value for a function pointer from __ccgo_fp. With gc, a
// func value is a single pointer and the function pointer is that value.
func funcPtr(fp uintptr) unsafe.Pointer {
	return unsafe.Pointer(&fp)
}
