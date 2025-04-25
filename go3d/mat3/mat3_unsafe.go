//go:build !netgo
// +build !netgo

package mat3

import "unsafe"

// Array returns the elements of the matrix as array pointer.
// The data may be a copy depending on the platform implementation.
func (mat *Mat[T]) Array() *[9]T {
	return (*[9]T)(unsafe.Pointer(mat)) //#nosec G103 -- unsafe OK
}
