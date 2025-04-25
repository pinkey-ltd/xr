// Package mat2 contains a 2x2 float matrix type Mat and functions.
package mat2

import (
	"errors"
	"fmt"
	"pinkey.ltd/xr/go3d"
	"pinkey.ltd/xr/go3d/vec2"
)

//var (
//	// Zero holds a zero matrix.
//	Zero = T{}
//
//	// Ident holds an ident matrix.
//	Ident = T{
//		vec2.Vec{1, 0},
//		vec2.Vec{0, 1},
//	}
//)

// Mat represents a 2x2 matrix.
type Mat[T float64 | float32] [2]vec2.Vec[T]

// From copies a 2x2 matrix from a Generic implementation.
func From[T float64 | float32](other go3d.T[T]) Mat[T] {
	r := Mat[T]{
		vec2.Vec[T]{1, 0},
		vec2.Vec[T]{0, 1},
	}

	cols := other.Cols()
	rows := other.Rows()
	if (cols == 3 && rows == 3) || (cols == 4 && rows == 4) {
		cols = 2
		rows = 2
	} else if !(cols == 2 && rows == 2) {
		panic("Unsupported type")
	}
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			r[col][row] = other.Get(col, row)
		}
	}
	return r
}

// Parse parses Mat from a string. See also String()
func Parse[T float64 | float32](s string) (r Mat[T], err error) {
	_, err = fmt.Sscan(s,
		&r[0][0], &r[0][1],
		&r[1][0], &r[1][1],
	)
	return r, err
}

// String formats T as string. See also Parse().
func (mat *Mat[T]) String() string {
	return fmt.Sprintf("%s %s", mat[0].String(), mat[1].String())
}

// Rows returns the number of rows of the matrix.
func (mat *Mat[T]) Rows() int {
	return 2
}

// Cols returns the number of columns of the matrix.
func (mat *Mat[T]) Cols() int {
	return 2
}

// Size returns the number elements of the matrix.
func (mat *Mat[T]) Size() int {
	return 4
}

// Slice returns the elements of the matrix as slice.
func (mat *Mat[T]) Slice() []T {
	return mat.Array()[:]
}

// Get returns one element of the matrix.
// Matrices are defined by (two) column vectors.
//
// Note that this function use the opposite reference order of rows and columns to the mathematical matrix indexing.
//
// A value in this matrix is referenced by <col><row> where both row and column is in the range [0..1].
// This notation and range reflects the underlying representation.
//
// A value in a matrix A is mathematically referenced by A<row><col>
// where both row and column is in the range [1..2].
// (It is really the lower case 'a' followed by <row><col> but this documentation syntax is somewhat limited.)
//
// matrixA.Get(0, 1) == matrixA[0][1] ( == A21 in mathematical indexing)
func (mat *Mat[T]) Get(col, row int) T {
	return mat[col][row]
}

// Set assigns the given value to the specified element at the given column and row indices within the matrix.
func (mat *Mat[T]) Set(col, row int, value T) {
	mat[col][row] = value
}

// IsZero checks if all elements of the matrix are zero.
func (mat *Mat[T]) IsZero() bool {
	zero := Mat[T]{
		vec2.Vec[T]{0, 0},
		vec2.Vec[T]{0, 0},
	}
	return *mat == zero
}

// Scale multiplies the diagonal scale elements by f returns mat.
func (mat *Mat[T]) Scale(f T) *Mat[T] {
	mat[0][0] *= f
	mat[1][1] *= f
	return mat
}

// Scaled returns a copy of the matrix with the diagonal scale elements multiplied by f.
func (mat *Mat[T]) Scaled(f T) Mat[T] {
	r := *mat
	return *r.Scale(f)
}

// Scaling returns the scaling diagonal of the matrix.
func (mat *Mat[T]) Scaling() vec2.Vec[T] {
	return vec2.Vec[T]{mat[0][0], mat[1][1]}
}

// SetScaling sets the scaling diagonal of the matrix.
func (mat *Mat[T]) SetScaling(s *vec2.Vec[T]) *Mat[T] {
	mat[0][0] = s[0]
	mat[1][1] = s[1]
	return mat
}

// Trace returns the trace value for the matrix.
func (mat *Mat[T]) Trace() T {
	return mat[0][0] + mat[1][1]
}

// AssignMul multiplies a and b and assigns the result to mat.
func (mat *Mat[T]) AssignMul(a, b *Mat[T]) *Mat[T] {
	mat[0] = a.MulVec2(&b[0])
	mat[1] = a.MulVec2(&b[1])
	return mat
}

// MulVec2 multiplies vec with mat.
func (mat *Mat[T]) MulVec2(vec *vec2.Vec[T]) vec2.Vec[T] {
	return vec2.Vec[T]{
		mat[0][0]*vec[0] + mat[1][0]*vec[1],
		mat[0][1]*vec[1] + mat[1][1]*vec[1],
	}
}

// TransformVec2 multiplies v with mat and saves the result in v.
func (mat *Mat[T]) TransformVec2(v *vec2.Vec[T]) {
	// Use intermediate variables to not alter further computations.
	x := mat[0][0]*v[0] + mat[1][0]*v[1]
	v[1] = mat[0][1]*v[0] + mat[1][1]*v[1]
	v[0] = x
}

func (mat *Mat[T]) Determinant() T {
	return mat[0][0]*mat[1][1] - mat[1][0]*mat[0][1]
}

// PracticallyEquals compares two matrices if they are equal with each other within a delta tolerance.
func (mat *Mat[T]) PracticallyEquals(matrix *Mat[T], allowedDelta T) bool {
	return mat[0].PracticallyEquals(&matrix[0], allowedDelta) &&
		mat[1].PracticallyEquals(&matrix[1], allowedDelta)
}

// Transpose transposes the matrix.
func (mat *Mat[T]) Transpose() *Mat[T] {
	temp := mat[0][1]
	mat[0][1] = mat[1][0]
	mat[1][0] = temp
	return mat
}

// Transposed returns a transposed copy the matrix.
func (mat *Mat[T]) Transposed() Mat[T] {
	result := *mat
	result.Transpose()
	return result
}

// Invert inverts the given matrix. Destructive operation.
// Does not check if matrix is singular and may lead to strange results!
func (mat *Mat[T]) Invert() (*Mat[T], error) {
	determinant := mat.Determinant()
	zero := Mat[T]{
		vec2.Vec[T]{0, 0},
		vec2.Vec[T]{0, 0},
	}

	if determinant == 0 {
		return &zero, errors.New("can not create inverted matrix as determinant is 0")
	}

	invDet := 1.0 / determinant

	mat[0][0], mat[1][1] = invDet*mat[1][1], invDet*mat[0][0]
	mat[0][1] = -invDet * mat[0][1]
	mat[1][0] = -invDet * mat[1][0]

	return mat, nil
}

// Inverted inverts a copy of the given matrix.
// Does not check if matrix is singular and may lead to strange results!
func (mat *Mat[T]) Inverted() (Mat[T], error) {
	result := *mat
	_, err := result.Invert()
	return result, err
}
