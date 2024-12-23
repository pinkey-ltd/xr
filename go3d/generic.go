// Package go3d Package generic contains an interface T that
// that all float vector and matrix types implement.
package go3d

// M is an interface that all float64 vector and matrix types implement.
type M[T float64 | float32] interface {

	// Cols returns the number of columns of the vector or matrix.
	Cols() int

	// Rows returns the number of rows of the vector or matrix.
	Rows() int

	// Size returns the number elements of the vector or matrix.
	Size() int

	// Slice returns the elements of the vector or matrix as slice.
	Slice() []T

	// Get returns one element of the vector or matrix.
	Get(row, col int) T

	// IsZero checks if all elements of the vector or matrix are zero.
	IsZero() bool
}
