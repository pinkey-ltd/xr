// Package mat3 contains a 3x3 float64 matrix type T and functions.
package mat3

import (
	"errors"
	"fmt"
	"math"

	"pinkey.ltd/xr/go3d"
	"pinkey.ltd/xr/go3d/mat2"
	"pinkey.ltd/xr/go3d/quaternion"
	"pinkey.ltd/xr/go3d/vec2"
	"pinkey.ltd/xr/go3d/vec3"
)

// var (
// 	// Zero holds a zero matrix.
// 	Zero = T{}
// 	// Ident holds an ident matrix.
// 	Ident = T{
// 		vec3.T{1, 0, 0},
// 		vec3.T{0, 1, 0},
// 		vec3.T{0, 0, 1},
// 	}
// )

// Mat represents a 3x3 matrix.
type Mat[T float64 | float32] [3]vec3.Vec[T]

// From copies a Mat3 from a go3d.T implementation.
func From[T float64 | float32](other go3d.T[T]) Mat[T] {
	r := Mat[T]{
		vec3.Vec[T]{1, 0, 0},
		vec3.Vec[T]{0, 1, 0},
		vec3.Vec[T]{0, 0, 1},
	}

	cols := other.Cols()
	rows := other.Rows()
	if cols == 4 && rows == 4 {
		cols = 3
		rows = 3
	} else if !((cols == 2 && rows == 2) || (cols == 3 && rows == 3)) {
		panic("Unsupported type")
	}
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			r[col][row] = other.Get(col, row)
		}
	}
	return r
}

// Parse parses T from a string. See also String()
func Parse[T float64 | float32](s string) (r Mat[T], err error) {
	_, err = fmt.Sscan(s,
		&r[0][0], &r[0][1], &r[0][2],
		&r[1][0], &r[1][1], &r[1][2],
		&r[2][0], &r[2][1], &r[2][2],
	)
	return r, err
}

// String formats T as string. See also Parse().
func (mat *Mat[T]) String() string {
	return fmt.Sprintf("%s %s %s", mat[0].String(), mat[1].String(), mat[2].String())
}

// Rows returns the number of rows of the matrix.
func (mat *Mat[T]) Rows() int {
	return 3
}

// Cols returns the number of columns of the matrix.
func (mat *Mat[T]) Cols() int {
	return 3
}

// Size returns the number elements of the matrix.
func (mat *Mat[T]) Size() int {
	return 9
}

// Slice returns the elements of the matrix as slice.
func (mat *Mat[T]) Slice() []T {
	return mat.Array()[:]
}

// Get returns one element of the matrix.
func (mat *Mat[T]) Get(col, row int) T {
	return mat[col][row]
}

// Set assigns the given value to the element at the specified column and row in the matrix.
func (mat *Mat[T]) Set(col, row int, value T) { mat[col][row] = value }

// IsZero checks if all elements of the matrix are zero.
func (mat *Mat[T]) IsZero() bool {
	return *mat == Mat[T]{}
}

// Scale multiplies the diagonal scale elements by f returns mat.
func (mat *Mat[T]) Scale(f T) *Mat[T] {
	mat[0][0] *= f
	mat[1][1] *= f
	mat[2][2] *= f
	return mat
}

// Scaled returns a copy of the matrix with the diagonal scale elements multiplied by f.
func (mat *Mat[T]) Scaled(f T) Mat[T] {
	r := *mat
	return *r.Scale(f)
}

// Scaling returns the scaling diagonal of the matrix.
func (mat *Mat[T]) Scaling() vec3.Vec[T] {
	return vec3.Vec[T]{mat[0][0], mat[1][1], mat[2][2]}
}

// SetScaling sets the scaling diagonal of the matrix.
func (mat *Mat[T]) SetScaling(s *vec3.Vec[T]) *Mat[T] {
	mat[0][0] = s[0]
	mat[1][1] = s[1]
	mat[2][2] = s[2]
	return mat
}

// ScaleVec2 multiplies the 2D scaling diagonal of the matrix by s.
func (mat *Mat[T]) ScaleVec2(s *vec2.Vec[T]) *Mat[T] {
	mat[0][0] *= s[0]
	mat[1][1] *= s[1]
	return mat
}

// SetTranslation sets the 2D translation elements of the matrix.
func (mat *Mat[T]) SetTranslation(v *vec2.Vec[T]) *Mat[T] {
	mat[2][0] = v[0]
	mat[2][1] = v[1]
	return mat
}

// Translate adds v to the 2D translation part of the matrix.
func (mat *Mat[T]) Translate(v *vec2.Vec[T]) *Mat[T] {
	mat[2][0] += v[0]
	mat[2][1] += v[1]
	return mat
}

// TranslateX adds dx to the 2D X-translation element of the matrix.
func (mat *Mat[T]) TranslateX(dx T) *Mat[T] {
	mat[2][0] += dx
	return mat
}

// TranslateY adds dy to the 2D Y-translation element of the matrix.
func (mat *Mat[T]) TranslateY(dy T) *Mat[T] {
	mat[2][1] += dy
	return mat
}

// Trace returns the trace value for the matrix.
func (mat *Mat[T]) Trace() T {
	return mat[0][0] + mat[1][1] + mat[2][2]
}

// AssignMul multiplies a and b and assigns the result to mat.
func (mat *Mat[T]) AssignMul(a, b *Mat[T]) *Mat[T] {
	mat[0] = a.MulVec3(&b[0])
	mat[1] = a.MulVec3(&b[1])
	mat[2] = a.MulVec3(&b[2])
	return mat
}

// AssignMat2x2 assigns a 2x2 sub-matrix and sets the rest of the matrix to the ident value.
func (mat *Mat[T]) AssignMat2x2(m *mat2.Mat[T]) *Mat[T] {
	*mat = Mat[T]{
		vec3.Vec[T]{m[0][0], m[1][0], 0},
		vec3.Vec[T]{m[0][1], m[1][1], 0},
		vec3.Vec[T]{0, 0, 1},
	}
	return mat
}

// Mul multiplies every element by f and returns mat.
func (mat *Mat[T]) Mul(f T) *Mat[T] {
	mat[0][0] *= f
	mat[0][1] *= f
	mat[0][2] *= f

	mat[1][0] *= f
	mat[1][1] *= f
	mat[1][2] *= f

	mat[2][0] *= f
	mat[2][1] *= f
	mat[2][2] *= f

	return mat
}

// Muled returns a copy of the matrix with every element multiplied by f.
func (mat *Mat[T]) Muled(f T) Mat[T] {
	result := *mat
	result.Mul(f)
	return result
}

// MulVec3 multiplies v with T.
func (mat *Mat[T]) MulVec3(v *vec3.Vec[T]) vec3.Vec[T] {
	return vec3.Vec[T]{
		mat[0][0]*v[0] + mat[1][0]*v[1] + mat[2][0]*v[2],
		mat[0][1]*v[1] + mat[1][1]*v[1] + mat[2][1]*v[2],
		mat[0][2]*v[2] + mat[1][2]*v[1] + mat[2][2]*v[2],
	}
}

// TransformVec3 multiplies v with mat and saves the result in v.
func (mat *Mat[T]) TransformVec3(v *vec3.Vec[T]) {
	// Use intermediate variables to not alter further computations.
	x := mat[0][0]*v[0] + mat[1][0]*v[1] + mat[2][0]*v[2]
	y := mat[0][1]*v[0] + mat[1][1]*v[1] + mat[2][1]*v[2]
	v[2] = mat[0][2]*v[0] + mat[1][2]*v[1] + mat[2][2]*v[2]
	v[0] = x
	v[1] = y
}

// Quaternion extracts a quaternion from the rotation part of the matrix.
func (mat *Mat[T]) Quaternion() quaternion.H[T] {
	tr := mat.Trace()

	s := T(math.Sqrt(float64(tr + 1)))
	w := s * 0.5
	s = 0.5 / s

	q := quaternion.H[T]{
		(mat[1][2] - mat[2][1]) * s,
		(mat[2][0] - mat[0][2]) * s,
		(mat[0][1] - mat[1][0]) * s,
		w,
	}
	return q.Normalized()
}

// AssignQuaternion assigns a quaternion to the rotations part of the matrix and sets the other elements to their ident value.
func (mat *Mat[T]) AssignQuaternion(q *quaternion.H[T]) *Mat[T] {
	xx := q[0] * q[0] * 2
	yy := q[1] * q[1] * 2
	zz := q[2] * q[2] * 2
	xy := q[0] * q[1] * 2
	xz := q[0] * q[2] * 2
	yz := q[1] * q[2] * 2
	wx := q[3] * q[0] * 2
	wy := q[3] * q[1] * 2
	wz := q[3] * q[2] * 2

	mat[0][0] = 1 - (yy + zz)
	mat[1][0] = xy - wz
	mat[2][0] = xz + wy

	mat[0][1] = xy + wz
	mat[1][1] = 1 - (xx + zz)
	mat[2][1] = yz - wx

	mat[0][2] = xz - wy
	mat[1][2] = yz + wx
	mat[2][2] = 1 - (xx + yy)

	return mat
}

// AssignXRotation assigns a rotation around the x axis to the rotation part of the matrix and sets the remaining elements to their ident value.
func (mat *Mat[T]) AssignXRotation(angle float64) *Mat[T] {
	cosine := T(math.Cos(angle))
	sine := T(math.Sin(angle))

	mat[0][0] = 1
	mat[1][0] = 0
	mat[2][0] = 0

	mat[0][1] = 0
	mat[1][1] = cosine
	mat[2][1] = -sine

	mat[0][2] = 0
	mat[1][2] = sine
	mat[2][2] = cosine

	return mat
}

// AssignYRotation assigns a rotation around the y axis to the rotation part of the matrix and sets the remaining elements to their ident value.
func (mat *Mat[T]) AssignYRotation(angle float64) *Mat[T] {
	cosine := T(math.Cos(angle))
	sine := T(math.Sin(angle))

	mat[0][0] = cosine
	mat[1][0] = 0
	mat[2][0] = sine

	mat[0][1] = 0
	mat[1][1] = 1
	mat[2][1] = 0

	mat[0][2] = -sine
	mat[1][2] = 0
	mat[2][2] = cosine

	return mat
}

// AssignZRotation assigns a rotation around the z axis to the rotation part of the matrix and sets the remaining elements to their ident value.
func (mat *Mat[T]) AssignZRotation(angle float64) *Mat[T] {
	cosine := T(math.Cos(angle))
	sine := T(math.Sin(angle))

	mat[0][0] = cosine
	mat[1][0] = -sine
	mat[2][0] = 0

	mat[0][1] = sine
	mat[1][1] = cosine
	mat[2][1] = 0

	mat[0][2] = 0
	mat[1][2] = 0
	mat[2][2] = 1

	return mat
}

// AssignCoordinateSystem assigns the rotation of a orthogonal coordinates system to the rotation part of the matrix and sets the remaining elements to their ident value.
func (mat *Mat[T]) AssignCoordinateSystem(x, y, z *vec3.Vec[T]) *Mat[T] {
	mat[0][0] = x[0]
	mat[1][0] = x[1]
	mat[2][0] = x[2]

	mat[0][1] = y[0]
	mat[1][1] = y[1]
	mat[2][1] = y[2]

	mat[0][2] = z[0]
	mat[1][2] = z[1]
	mat[2][2] = z[2]

	return mat
}

// AssignEulerRotation assigns Euler angle rotations to the rotation part of the matrix and sets the remaining elements to their ident value.
func (mat *Mat[T]) AssignEulerRotation(yHead, xPitch, zRoll float64) *Mat[T] {
	sinH := T(math.Sin(yHead))
	cosH := T(math.Cos(yHead))
	sinP := T(math.Sin(xPitch))
	cosP := T(math.Cos(xPitch))
	sinR := T(math.Sin(zRoll))
	cosR := T(math.Cos(zRoll))

	mat[0][0] = cosR*cosH - sinR*sinP*sinH
	mat[1][0] = -sinR * cosP
	mat[2][0] = cosR*sinH + sinR*sinP*cosH

	mat[0][1] = sinR*cosH + cosR*sinP*sinH
	mat[1][1] = cosR * cosP
	mat[2][1] = sinR*sinH - cosR*sinP*cosH

	mat[0][2] = -cosP * sinH
	mat[1][2] = sinP
	mat[2][2] = cosP * cosH

	return mat
}

// ExtractEulerAngles extracts the rotation part of the matrix as Euler angle rotation values.
func (mat *Mat[T]) ExtractEulerAngles() (yHead, xPitch, zRoll T) {
	xPitch = T(math.Asin(float64(mat[1][2])))
	f12 := math.Abs(float64(mat[1][2]))
	if f12 > (1.0-0.0001) && f12 < (1.0+0.0001) { // f12 == 1.0
		yHead = 0.0
		zRoll = T(math.Atan2(float64(mat[0][1]), float64(mat[0][0])))
	} else {
		yHead = T(math.Atan2(float64(-mat[0][2]), float64(mat[2][2])))
		zRoll = T(math.Atan2(float64(-mat[1][0]), float64(mat[1][1])))
	}
	return yHead, xPitch, zRoll
}

// Determinant returns the determinant of the matrix.
func (mat *Mat[T]) Determinant() T {
	return mat[0][0]*mat[1][1]*mat[2][2] +
		mat[1][0]*mat[2][1]*mat[0][2] +
		mat[2][0]*mat[0][1]*mat[1][2] -
		mat[2][0]*mat[1][1]*mat[0][2] -
		mat[1][0]*mat[0][1]*mat[2][2] -
		mat[0][0]*mat[2][1]*mat[1][2]
}

// IsReflective returns true if the matrix can be reflected by a plane.
func (mat *Mat[T]) IsReflective() bool {
	return mat.Determinant() < 0
}

// Transpose transposes the matrix.
func (mat *Mat[T]) Transpose() *Mat[T] {
	mat[1][0], mat[0][1] = mat[0][1], mat[1][0]
	mat[2][0], mat[0][2] = mat[0][2], mat[2][0]
	mat[2][1], mat[1][2] = mat[1][2], mat[2][1]
	return mat
}

// Transposed returns a transposed copy the matrix.
func (mat *Mat[T]) Transposed() Mat[T] {
	result := *mat
	result.Transpose()
	return result
}

// Adjugate computes the adjugate of this matrix and returns mat
func (mat *Mat[T]) Adjugate() *Mat[T] {
	m := *mat

	mat[0][0] = +(m[1][1]*m[2][2] - m[1][2]*m[2][1])
	mat[0][1] = -(m[0][1]*m[2][2] - m[0][2]*m[2][1])
	mat[0][2] = +(m[0][1]*m[1][2] - m[0][2]*m[1][1])

	mat[1][0] = -(m[1][0]*m[2][2] - m[1][2]*m[2][0])
	mat[1][1] = +(m[0][0]*m[2][2] - m[0][2]*m[2][0])
	mat[1][2] = -(m[0][0]*m[1][2] - m[0][2]*m[1][0])

	mat[2][0] = +(m[1][0]*m[2][1] - m[1][1]*m[2][0])
	mat[2][1] = -(m[0][0]*m[2][1] - m[0][1]*m[2][0])
	mat[2][2] = +(m[0][0]*m[1][1] - m[0][1]*m[1][0])

	return mat
}

// Adjugated returns an adjugated copy of the matrix.
func (mat *Mat[T]) Adjugated() Mat[T] {
	result := *mat
	result.Adjugate()
	return result
}

// returns a 3x3 matrix without the i-th column and j-th row
func (mat *Mat[T]) maskedBlock(blockI, blockJ int) *mat2.Mat[T] {
	var m mat2.Mat[T]
	m_i := 0
	for i := 0; i < 3; i++ {
		if i == blockI {
			continue
		}
		m_j := 0
		for j := 0; j < 3; j++ {
			if j == blockJ {
				continue
			}
			m[m_i][m_j] = mat[i][j]
			m_j++
		}
		m_i++
	}
	return &m
}

// Invert inverts the given matrix. Destructive operation.
// Does not check if matrix is singular and may lead to strange results!
func (mat *Mat[T]) Invert() (*Mat[T], error) {
	initialDet := mat.Determinant()
	zero := Mat[T]{}
	if initialDet == 0 {
		return &zero, errors.New("can not create inverted matrix as determinant is 0")
	}

	mat.Adjugate()
	mat.Mul(1.0 / initialDet)
	return mat, nil
}

// Inverted inverts a copy of the given matrix.
// Does not check if matrix is singular and may lead to strange results!
func (mat *Mat[T]) Inverted() (Mat[T], error) {
	result := *mat
	_, err := result.Invert()
	return result, err
}
