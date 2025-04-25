//go:build !netgo

package mat2

import "unsafe"

// Array returns the elements of the matrix as array pointer.
// The data may be a copy depending on the platform implementation.
func (mat *Mat[T]) Array() *[4]T {
	return (*[4]T)(unsafe.Pointer(mat)) //#nosec G103 -- unsafe OK
}
