//go:build !netgo

package mat4

import "unsafe"

// Array returns the elements of the matrix as array pointer.
// The data may be a copy depending on the platform implementation.
func (mat *Mat[T]) Array() *[16]T {
	return (*[16]T)(unsafe.Pointer(mat)) //#nosec G103 -- unsafe OK
}
