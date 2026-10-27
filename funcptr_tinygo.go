//go:build tinygo

package b3

import "unsafe"

// funcPtr returns a pointer to a func value for a function pointer from __ccgo_fp. TinyGo
// stores a func value in two words, so __ccgo_fp already returns a pointer to the boxed value.
func funcPtr(fp uintptr) unsafe.Pointer {
	return unsafe.Pointer(fp)
}
