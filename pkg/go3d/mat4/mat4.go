// Package mat4 contains a 4x4 float64 matrix type T and functions.
package mat4

import (
	"fmt"
	"math"

	"pinkey.ltd/xr/pkg/go3d"
	"pinkey.ltd/xr/pkg/go3d/mat2"
	"pinkey.ltd/xr/pkg/go3d/mat3"
	"pinkey.ltd/xr/pkg/go3d/quaternion"
	"pinkey.ltd/xr/pkg/go3d/vec3"
	"pinkey.ltd/xr/pkg/go3d/vec4"
)

//var (
//	// Zero holds a zero matrix.
//	Zero = T{}
//
//	// Ident holds an ident matrix.
//	Ident = T{
//		vec4.T{1, 0, 0, 0},
//		vec4.T{0, 1, 0, 0},
//		vec4.T{0, 0, 1, 0},
//		vec4.T{0, 0, 0, 1},
//	}
//)

// Mat represents a 4x4 matrix.
type Mat[T float64 | float32] [4]vec4.Vec[T]

// From copies a T from a generic.T implementation.
func From[T float64 | float32](other go3d.T[T]) Mat[T] {
	r := Mat[T]{
		vec4.Vec[T]{1, 0, 0, 0},
		vec4.Vec[T]{0, 1, 0, 0},
		vec4.Vec[T]{0, 0, 1, 0},
		vec4.Vec[T]{0, 0, 0, 1},
	}
	cols := other.Cols()
	rows := other.Rows()
	if !((cols == 2 && rows == 2) || (cols == 3 && rows == 3) || (cols == 4 && rows == 4)) {
		panic("Unsupported type")
	}
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			r[col][row] = other.Get(col, row)
		}
	}
	return r
}

func FromArray[T float64 | float32](mtw [16]T) Mat[T] {
	mt := Mat[T]{
		vec4.Vec[T]{mtw[0], mtw[1], mtw[2], mtw[3]},
		vec4.Vec[T]{mtw[4], mtw[5], mtw[6], mtw[7]},
		vec4.Vec[T]{mtw[8], mtw[9], mtw[10], mtw[11]},
		vec4.Vec[T]{mtw[12], mtw[13], mtw[14], mtw[15]},
	}
	return mt
}

// Parse parses T from a string. See also String()
func Parse[T float64 | float32](s string) (r Mat[T], err error) {
	_, err = fmt.Sscan(s,
		&r[0][0], &r[0][1], &r[0][2], &r[0][3],
		&r[1][0], &r[1][1], &r[1][2], &r[1][3],
		&r[2][0], &r[2][1], &r[2][2], &r[2][3],
		&r[3][0], &r[3][1], &r[3][2], &r[3][3],
	)
	return r, err
}

// String formats T as string. See also Parse().
func (mat *Mat[T]) String() string {
	return fmt.Sprintf("%s %s %s %s", mat[0].String(), mat[1].String(), mat[2].String(), mat[3].String())
}

// Rows returns the number of rows of the matrix.
func (mat *Mat[T]) Rows() int {
	return 4
}

// Cols returns the number of columns of the matrix.
func (mat *Mat[T]) Cols() int {
	return 4
}

// Size returns the number elements of the matrix.
func (mat *Mat[T]) Size() int {
	return 16
}

// Slice returns the elements of the matrix as slice.
func (mat *Mat[T]) Slice() []T {
	return mat.Array()[:]
}

// Get returns one element of the matrix.
func (mat *Mat[T]) Get(col, row int) T {
	return mat[col][row]
}

// IsZero checks if all elements of the matrix are zero.
func (mat *Mat[T]) IsZero() bool {
	zero := Mat[T]{}
	return *mat == zero
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

// Mul multiplies every element by f and returns mat.
func (mat *Mat[T]) Mul(f T) *Mat[T] {
	for i, col := range mat {
		for j := range col {
			mat[i][j] *= f
		}
	}
	return mat
}

// Muled returns a copy of the matrix with every element multiplied by f.
func (mat *Mat[T]) Muled(f T) Mat[T] {
	result := *mat
	result.Mul(f)
	return result
}

// Trace returns the trace value for the matrix.
func (mat *Mat[T]) Trace() T {
	return mat[0][0] + mat[1][1] + mat[2][2] + mat[3][3]
}

// Trace3 returns the trace value for the 3x3 sub-matrix.
func (mat *Mat[T]) Trace3() T {
	return mat[0][0] + mat[1][1] + mat[2][2]
}

// AssignMat2x2 assigns a 2x2 sub-matrix and sets the rest of the matrix to the ident value.
func (mat *Mat[T]) AssignMat2x2(m *mat2.Mat[T]) *Mat[T] {
	*mat = Mat[T]{
		vec4.Vec[T]{m[0][0], m[1][0], 0, 0},
		vec4.Vec[T]{m[0][1], m[1][1], 0, 0},
		vec4.Vec[T]{0, 0, 1, 0},
		vec4.Vec[T]{0, 0, 0, 1},
	}
	return mat
}

// AssignMat3x3 assigns a 3x3 sub-matrix and sets the rest of the matrix to the ident value.
func (mat *Mat[T]) AssignMat3x3(m *mat3.Mat[T]) *Mat[T] {
	*mat = Mat[T]{
		vec4.Vec[T]{m[0][0], m[1][0], m[2][0], 0},
		vec4.Vec[T]{m[0][1], m[1][1], m[2][1], 0},
		vec4.Vec[T]{m[0][2], m[1][2], m[2][2], 0},
		vec4.Vec[T]{0, 0, 0, 1},
	}
	return mat
}

// AssignMul multiplies a and b and assigns the result to T.
func (mat *Mat[T]) AssignMul(a, b *Mat[T]) *Mat[T] {
	mat[0] = a.MulVec4(&b[0])
	mat[1] = a.MulVec4(&b[1])
	mat[2] = a.MulVec4(&b[2])
	mat[3] = a.MulVec4(&b[3])
	return mat
}

// MulVec4 multiplies v with mat and returns a new vector v' = M * v.
func (mat *Mat[T]) MulVec4(v *vec4.Vec[T]) vec4.Vec[T] {
	return vec4.Vec[T]{
		mat[0][0]*v[0] + mat[1][0]*v[1] + mat[2][0]*v[2] + mat[3][0]*v[3],
		mat[0][1]*v[0] + mat[1][1]*v[1] + mat[2][1]*v[2] + mat[3][1]*v[3],
		mat[0][2]*v[0] + mat[1][2]*v[1] + mat[2][2]*v[2] + mat[3][2]*v[3],
		mat[0][3]*v[0] + mat[1][3]*v[1] + mat[2][3]*v[2] + mat[3][3]*v[3],
	}
}

// TransformVec4 multiplies v with mat and saves the result in v.
func (mat *Mat[T]) TransformVec4(v *vec4.Vec[T]) {
	// Use intermediate variables to not alter further computations.
	x := mat[0][0]*v[0] + mat[1][0]*v[1] + mat[2][0]*v[2] + mat[3][0]*v[3]
	y := mat[0][1]*v[0] + mat[1][1]*v[1] + mat[2][1]*v[2] + mat[3][1]*v[3]
	z := mat[0][2]*v[0] + mat[1][2]*v[1] + mat[2][2]*v[2] + mat[3][2]*v[3]
	v[3] = mat[0][3]*v[0] + mat[1][3]*v[1] + mat[2][3]*v[2] + mat[3][3]*v[3]
	v[0] = x
	v[1] = y
	v[2] = z
}

// MulVec3 multiplies v (converted to a vec4 as (v_1, v_2, v_3, 1))
// with mat and divides the result by w. Returns a new vec3.
func (mat *Mat[T]) MulVec3(v *vec3.Vec[T]) vec3.Vec[T] {
	v4 := vec4.FromVec3(v)
	v4 = mat.MulVec4(&v4)
	return v4.Vec3DividedByW()
}

// TransformVec3 multiplies v (converted to a vec4 as (v_1, v_2, v_3, 1))
// with mat, divides the result by w and saves the result in v.
func (mat *Mat[T]) TransformVec3(v *vec3.Vec[T]) {
	x := mat[0][0]*v[0] + mat[1][0]*v[1] + mat[2][0]*v[2] + mat[3][0]
	y := mat[0][1]*v[0] + mat[1][1]*v[1] + mat[2][1]*v[2] + mat[3][1]
	z := mat[0][2]*v[0] + mat[1][2]*v[1] + mat[2][2]*v[2] + mat[3][2]
	w := mat[0][3]*v[0] + mat[1][3]*v[1] + mat[2][3]*v[2] + mat[3][3]
	oow := 1 / w
	v[0] = x * oow
	v[1] = y * oow
	v[2] = z * oow
}

// MulVec3W multiplies v with mat with w as fourth component of the vector.
// Useful to differentiate between vectors (w = 0) and points (w = 1)
// without transforming them to vec4.
func (mat *Mat[T]) MulVec3W(v *vec3.Vec[T], w T) vec3.Vec[T] {
	result := *v
	mat.TransformVec3W(&result, w)
	return result
}

// TransformVec3W multiplies v with mat with w as fourth component of the vector and
// saves the result in v.
// Useful to differentiate between vectors (w = 0) and points (w = 1)
// without transforming them to vec4.
func (mat *Mat[T]) TransformVec3W(v *vec3.Vec[T], w T) {
	// use intermediate variables to not alter further computations
	x := mat[0][0]*v[0] + mat[1][0]*v[1] + mat[2][0]*v[2] + mat[3][0]*w
	y := mat[0][1]*v[0] + mat[1][1]*v[1] + mat[2][1]*v[2] + mat[3][1]*w
	v[2] = mat[0][2]*v[0] + mat[1][2]*v[1] + mat[2][2]*v[2] + mat[3][2]*w
	v[0] = x
	v[1] = y
}

// SetTranslation sets the translation elements of the matrix.
func (mat *Mat[T]) SetTranslation(v *vec3.Vec[T]) *Mat[T] {
	mat[3][0] = v[0]
	mat[3][1] = v[1]
	mat[3][2] = v[2]
	return mat
}

// Translate adds v to the translation part of the matrix.
func (mat *Mat[T]) Translate(v *vec3.Vec[T]) *Mat[T] {
	mat[3][0] += v[0]
	mat[3][1] += v[1]
	mat[3][2] += v[2]
	return mat
}

// TranslateX adds dx to the X-translation element of the matrix.
func (mat *Mat[T]) TranslateX(dx T) *Mat[T] {
	mat[3][0] += dx
	return mat
}

// TranslateY adds dy to the Y-translation element of the matrix.
func (mat *Mat[T]) TranslateY(dy T) *Mat[T] {
	mat[3][1] += dy
	return mat
}

// TranslateZ adds dz to the Z-translation element of the matrix.
func (mat *Mat[T]) TranslateZ(dz T) *Mat[T] {
	mat[3][2] += dz
	return mat
}

// Scaling returns the scaling diagonal of the matrix.
func (mat *Mat[T]) Scaling() vec4.Vec[T] {
	return vec4.Vec[T]{mat[0][0], mat[1][1], mat[2][2], mat[3][3]}
}

// SetScaling sets the scaling diagonal of the matrix.
func (mat *Mat[T]) SetScaling(s *vec4.Vec[T]) *Mat[T] {
	mat[0][0] = s[0]
	mat[1][1] = s[1]
	mat[2][2] = s[2]
	mat[3][3] = s[3]
	return mat
}

// ScaleVec3 multiplies the scaling diagonal of the matrix by s.
func (mat *Mat[T]) ScaleVec3(s *vec3.Vec[T]) *Mat[T] {
	mat[0][0] *= s[0]
	mat[1][1] *= s[1]
	mat[2][2] *= s[2]
	return mat
}

func (mat *Mat[T]) Quaternion() quaternion.H[T] {
	// http://www.euclideanspace.com/maths/geometry/rotations/conversions/matrixToQuaternion/index.htm
	// assumes the upper 3x3 of m is a pure rotation matrix (i.e, unscaled)
	m11, m12, m13 := mat[0][0], mat[1][0], mat[2][0]
	m21, m22, m23 := mat[0][1], mat[1][1], mat[2][1]
	m31, m32, m33 := mat[0][2], mat[1][2], mat[2][2]

	trace := m11 + m22 + m33
	var s, _w, _x, _y, _z T

	if trace > 0 {
		s = T(0.5 / math.Sqrt(float64(trace+1.0)))

		_w = 0.25 / s
		_x = (m32 - m23) * s
		_y = (m13 - m31) * s
		_z = (m21 - m12) * s
	} else if m11 > m22 && m11 > m33 {
		s = T(2.0 * math.Sqrt(float64(1.0+m11-m22-m33)))

		_w = (m32 - m23) / s
		_x = 0.25 * s
		_y = (m12 + m21) / s
		_z = (m13 + m31) / s
	} else if m22 > m33 {
		s = T(2.0 * math.Sqrt(float64(1.0+m22-m11-m33)))

		_w = (m13 - m31) / s
		_x = (m12 + m21) / s
		_y = 0.25 * s
		_z = (m23 + m32) / s
	} else {
		s = T(2.0 * math.Sqrt(float64(1.0+m33-m11-m22)))

		_w = (m21 - m12) / s
		_x = (m13 + m31) / s
		_y = (m23 + m32) / s
		_z = 0.25 * s
	}
	return quaternion.H[T]{
		_x,
		_y,
		_z,
		_w,
	}
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
	mat[3][0] = 0

	mat[0][1] = xy + wz
	mat[1][1] = 1 - (xx + zz)
	mat[2][1] = yz - wx
	mat[3][1] = 0

	mat[0][2] = xz - wy
	mat[1][2] = yz + wx
	mat[2][2] = 1 - (xx + yy)
	mat[3][2] = 0

	mat[0][3] = 0
	mat[1][3] = 0
	mat[2][3] = 0
	mat[3][3] = 1

	return mat
}

// AssignXRotation assigns a rotation around the x axis to the rotation part of the matrix and sets the remaining elements to their ident value.
func (mat *Mat[T]) AssignXRotation(angle float64) *Mat[T] {
	cosine := T(math.Cos(angle))
	sine := T(math.Sin(angle))

	mat[0][0] = 1
	mat[1][0] = 0
	mat[2][0] = 0
	mat[3][0] = 0

	mat[0][1] = 0
	mat[1][1] = cosine
	mat[2][1] = -sine
	mat[3][1] = 0

	mat[0][2] = 0
	mat[1][2] = sine
	mat[2][2] = cosine
	mat[3][2] = 0

	mat[0][3] = 0
	mat[1][3] = 0
	mat[2][3] = 0
	mat[3][3] = 1

	return mat
}

// AssignYRotation assigns a rotation around the y axis to the rotation part of the matrix and sets the remaining elements to their ident value.
func (mat *Mat[T]) AssignYRotation(angle float64) *Mat[T] {
	cosine := T(math.Cos(angle))
	sine := T(math.Sin(angle))

	mat[0][0] = cosine
	mat[1][0] = 0
	mat[2][0] = sine
	mat[3][0] = 0

	mat[0][1] = 0
	mat[1][1] = 1
	mat[2][1] = 0
	mat[3][1] = 0

	mat[0][2] = -sine
	mat[1][2] = 0
	mat[2][2] = cosine
	mat[3][2] = 0

	mat[0][3] = 0
	mat[1][3] = 0
	mat[2][3] = 0
	mat[3][3] = 1

	return mat
}

// AssignZRotation assigns a rotation around the z axis to the rotation part of the matrix and sets the remaining elements to their ident value.
func (mat *Mat[T]) AssignZRotation(angle float64) *Mat[T] {
	cosine := T(math.Cos(angle))
	sine := T(math.Sin(angle))

	mat[0][0] = cosine
	mat[1][0] = -sine
	mat[2][0] = 0
	mat[3][0] = 0

	mat[0][1] = sine
	mat[1][1] = cosine
	mat[2][1] = 0
	mat[3][1] = 0

	mat[0][2] = 0
	mat[1][2] = 0
	mat[2][2] = 1
	mat[3][2] = 0

	mat[0][3] = 0
	mat[1][3] = 0
	mat[2][3] = 0
	mat[3][3] = 1

	return mat
}

// AssignCoordinateSystem assigns the rotation of a orthogonal coordinates system to the rotation part of the matrix and sets the remaining elements to their ident value.
func (mat *Mat[T]) AssignCoordinateSystem(x, y, z *vec3.Vec[T]) *Mat[T] {
	mat[0][0] = x[0]
	mat[1][0] = x[1]
	mat[2][0] = x[2]
	mat[3][0] = 0

	mat[0][1] = y[0]
	mat[1][1] = y[1]
	mat[2][1] = y[2]
	mat[3][1] = 0

	mat[0][2] = z[0]
	mat[1][2] = z[1]
	mat[2][2] = z[2]
	mat[3][2] = 0

	mat[0][3] = 0
	mat[1][3] = 0
	mat[2][3] = 0
	mat[3][3] = 1

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
	mat[3][0] = 0

	mat[0][1] = sinR*cosH + cosR*sinP*sinH
	mat[1][1] = cosR * cosP
	mat[2][1] = sinR*sinH - cosR*sinP*cosH
	mat[3][1] = 0

	mat[0][2] = -cosP * sinH
	mat[1][2] = sinP
	mat[2][2] = cosP * cosH
	mat[3][2] = 0

	mat[0][3] = 0
	mat[1][3] = 0
	mat[2][3] = 0
	mat[3][3] = 1

	return mat
}

// ExtractEulerAngles extracts the rotation part of the matrix as Euler angle rotation values.
func (mat *Mat[T]) ExtractEulerAngles() (yHead, xPitch, zRoll float64) {
	xPitch = math.Asin(float64(mat[1][2]))
	f12 := math.Abs(float64(mat[1][2]))
	if f12 > (1.0-0.0001) && f12 < (1.0+0.0001) { // f12 == 1.0
		yHead = 0.0
		zRoll = math.Atan2(float64(mat[0][1]), float64(mat[0][0]))
	} else {
		yHead = math.Atan2(float64(-mat[0][2]), float64(mat[2][2]))
		zRoll = math.Atan2(float64(-mat[1][0]), float64(mat[1][1]))
	}
	return yHead, xPitch, zRoll
}

// AssignPerspectiveProjection assigns a perspective projection transformation.
func (mat *Mat[T]) AssignPerspectiveProjection(left, right, bottom, top, znear, zfar T) *Mat[T] {
	near2 := znear + znear
	ooFarNear := 1 / (zfar - znear)

	mat[0][0] = near2 / (right - left)
	mat[1][0] = 0
	mat[2][0] = (right + left) / (right - left)
	mat[3][0] = 0

	mat[0][1] = 0
	mat[1][1] = near2 / (top - bottom)
	mat[2][1] = (top + bottom) / (top - bottom)
	mat[3][1] = 0

	mat[0][2] = 0
	mat[1][2] = 0
	mat[2][2] = -(zfar + znear) * ooFarNear
	mat[3][2] = -2 * zfar * znear * ooFarNear

	mat[0][3] = 0
	mat[1][3] = 0
	mat[2][3] = -1
	mat[3][3] = 0

	return mat
}

// AssignOrthogonalProjection assigns an orthogonal projection transformation.
func (mat *Mat[T]) AssignOrthogonalProjection(left, right, bottom, top, znear, zfar T) *Mat[T] {
	ooRightLeft := 1 / (right - left)
	ooTopBottom := 1 / (top - bottom)
	ooFarNear := 1 / (zfar - znear)

	mat[0][0] = 2 * ooRightLeft
	mat[1][0] = 0
	mat[2][0] = 0
	mat[3][0] = -(right + left) * ooRightLeft

	mat[0][1] = 0
	mat[1][1] = 2 * ooTopBottom
	mat[2][1] = 0
	mat[3][1] = -(top + bottom) * ooTopBottom

	mat[0][2] = 0
	mat[1][2] = 0
	mat[2][2] = -2 * ooFarNear
	mat[3][2] = -(zfar + znear) * ooFarNear

	mat[0][3] = 0
	mat[1][3] = 0
	mat[2][3] = 0
	mat[3][3] = 1

	return mat
}

// Determinant3x3 returns the determinant of the 3x3 sub-matrix.
func (mat *Mat[T]) Determinant3x3() T {
	return mat[0][0]*mat[1][1]*mat[2][2] +
		mat[1][0]*mat[2][1]*mat[0][2] +
		mat[2][0]*mat[0][1]*mat[1][2] -
		mat[2][0]*mat[1][1]*mat[0][2] -
		mat[1][0]*mat[0][1]*mat[2][2] -
		mat[0][0]*mat[2][1]*mat[1][2]
}

func (mat *Mat[T]) Determinant() T {
	s1 := mat[0][0]
	det1 := mat[1][1]*mat[2][2]*mat[3][3] +
		mat[2][1]*mat[3][2]*mat[1][3] +
		mat[3][1]*mat[1][2]*mat[2][3] -
		mat[3][1]*mat[2][2]*mat[1][3] -
		mat[2][1]*mat[1][2]*mat[3][3] -
		mat[1][1]*mat[3][2]*mat[2][3]

	s2 := mat[0][1]
	det2 := mat[1][0]*mat[2][2]*mat[3][3] +
		mat[2][0]*mat[3][2]*mat[1][3] +
		mat[3][0]*mat[1][2]*mat[2][3] -
		mat[3][0]*mat[2][2]*mat[1][3] -
		mat[2][0]*mat[1][2]*mat[3][3] -
		mat[1][0]*mat[3][2]*mat[2][3]
	s3 := mat[0][2]
	det3 := mat[1][0]*mat[2][1]*mat[3][3] +
		mat[2][0]*mat[3][1]*mat[1][3] +
		mat[3][0]*mat[1][1]*mat[2][3] -
		mat[3][0]*mat[2][1]*mat[1][3] -
		mat[2][0]*mat[1][1]*mat[3][3] -
		mat[1][0]*mat[3][1]*mat[2][3]
	s4 := mat[0][3]
	det4 := mat[1][0]*mat[2][1]*mat[3][2] +
		mat[2][0]*mat[3][1]*mat[1][2] +
		mat[3][0]*mat[1][1]*mat[2][2] -
		mat[3][0]*mat[2][1]*mat[1][2] -
		mat[2][0]*mat[1][1]*mat[3][2] -
		mat[1][0]*mat[3][1]*mat[2][2]
	return s1*det1 - s2*det2 + s3*det3 - s4*det4
}

// IsReflective returns true if the matrix can be reflected by a plane.
func (mat *Mat[T]) IsReflective() bool {
	return mat.Determinant3x3() < 0
}

// Transpose transposes the matrix.
func (mat *Mat[T]) Transpose() *Mat[T] {
	//swap(&mat[3][0], &mat[0][3])
	mat[3][0], mat[0][3] = mat[0][3], mat[3][0]
	//swap(&mat[3][1], &mat[1][3])
	mat[3][1], mat[1][3] = mat[1][3], mat[3][1]
	//swap(&mat[3][2], &mat[2][3])
	mat[3][2], mat[2][3] = mat[2][3], mat[3][2]
	return mat.Transpose3x3()
}

// Transpose3x3 transposes the 3x3 sub-matrix.
func (mat *Mat[T]) Transpose3x3() *Mat[T] {
	//swap(&mat[1][0], &mat[0][1])
	mat[1][0], mat[0][1] = mat[0][1], mat[1][0]
	//swap(&mat[2][0], &mat[0][2])
	mat[2][0], mat[0][2] = mat[0][2], mat[2][0]
	//swap(&mat[2][1], &mat[1][2])
	mat[2][1], mat[1][2] = mat[1][2], mat[2][1]
	return mat
}

// Adjugate computes the adjugate of this matrix and returns mat
func (mat *Mat[T]) Adjugate() *Mat[T] {
	matOrig := *mat
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			// - 1 for odd i+j, 1 for even i+j
			sign := T(((i+j)%2)*-2 + 1)
			mat[i][j] = matOrig.maskedBlock(i, j).Determinant() * sign
		}
	}
	return mat.Transpose()
}

// Adjugated returns an adjugated copy of the matrix.
func (mat *Mat[T]) Adjugated() Mat[T] {
	result := *mat
	result.Adjugate()
	return result
}

// returns a 3x3 matrix without the i-th column and j-th row
func (mat *Mat[T]) maskedBlock(blockI, blockJ int) *mat3.Mat[T] {
	var m mat3.Mat[T]
	m_i := 0
	for i := 0; i < 4; i++ {
		if i == blockI {
			continue
		}
		m_j := 0
		for j := 0; j < 4; j++ {
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

// Inverts the given matrix.
// Does not check if matrix is singular and may lead to strange results!
func (mat *Mat[T]) Invert() *Mat[T] {
	initialDet := mat.Determinant()
	mat.Adjugate()
	mat.Mul(1 / initialDet)
	return mat
}

// Inverted returns an inverted copy of the matrix.
// Does not check if matrix is singular and may lead to strange results!
func (mat *Mat[T]) Inverted() Mat[T] {
	result := *mat
	result.Invert()
	return result
}

func v3Combine[T float64 | float32](a *vec3.Vec[T], b *vec3.Vec[T], result *vec3.Vec[T], ascl T, bscl T) {

	result[0] = (ascl * a[0]) + (bscl * b[0])
	result[1] = (ascl * a[1]) + (bscl * b[1])
	result[2] = (ascl * a[2]) + (bscl * b[2])
}

func Decompose[T float64 | float32](mat *Mat[T]) (*vec3.Vec[T], *quaternion.H[T], *vec3.Vec[T]) {
	sx := (&vec3.Vec[T]{mat[0][0], mat[0][1], mat[0][2]}).Length()
	sy := (&vec3.Vec[T]{mat[1][0], mat[1][1], mat[1][2]}).Length()
	sz := (&vec3.Vec[T]{mat[2][0], mat[2][1], mat[2][2]}).Length()

	// if determine is negative, we need to invert one scale
	det := mat.Determinant()
	if det < 0 {
		sx = -sx
	}

	position := vec3.Vec[T]{mat[3][0], mat[3][1], mat[3][2]}

	// scale the rotation part
	invSX, invSY, invSZ := 1.0/sx, 1.0/sy, 1.0/sz
	matrix := Mat[T]{
		vec4.Vec[T]{mat[0][0] * invSX, mat[0][1] * invSX, mat[0][2] * invSX, 0},
		vec4.Vec[T]{mat[1][0] * invSY, mat[1][1] * invSY, mat[1][2] * invSY, 0},
		vec4.Vec[T]{mat[2][0] * invSZ, mat[2][1] * invSZ, mat[2][2] * invSZ, 0},
		vec4.Vec[T]{0, 0, 0, 1},
	}
	quat := matrix.Quaternion()

	scale := vec3.Vec[T]{sx, sy, sz}

	return &position, &quat, &scale
}

func Compose[T float64 | float32](pos *vec3.Vec[T], quat *quaternion.H[T], scale *vec3.Vec[T]) *Mat[T] {
	posMat := Mat[T]{
		vec4.Vec[T]{1, 0, 0, 0},
		vec4.Vec[T]{0, 1, 0, 0},
		vec4.Vec[T]{0, 0, 1, 0},
		vec4.Vec[T]{0, 0, 0, 1},
	}
	posMat.SetTranslation(pos)
	quatMat := Mat[T]{
		vec4.Vec[T]{1, 0, 0, 0},
		vec4.Vec[T]{0, 1, 0, 0},
		vec4.Vec[T]{0, 0, 1, 0},
		vec4.Vec[T]{0, 0, 0, 1},
	}
	quatMat.AssignQuaternion(quat)
	scaleMat := Mat[T]{
		vec4.Vec[T]{1, 0, 0, 0},
		vec4.Vec[T]{0, 1, 0, 0},
		vec4.Vec[T]{0, 0, 1, 0},
		vec4.Vec[T]{0, 0, 0, 1},
	}
	scaleMat.ScaleVec3(scale)

	qsMat := Mat[T]{
		vec4.Vec[T]{1, 0, 0, 0},
		vec4.Vec[T]{0, 1, 0, 0},
		vec4.Vec[T]{0, 0, 1, 0},
		vec4.Vec[T]{0, 0, 0, 1},
	}
	qsMat.AssignMul(&quatMat, &scaleMat)

	result := Mat[T]{
		vec4.Vec[T]{1, 0, 0, 0},
		vec4.Vec[T]{0, 1, 0, 0},
		vec4.Vec[T]{0, 0, 1, 0},
		vec4.Vec[T]{0, 0, 0, 1},
	}
	result.AssignMul(&posMat, &qsMat)

	return &result
}

func (mat *Mat[T]) LookAt(eye, target, up vec3.Vec[T]) *Mat[T] {
	_z := vec3.Sub(&eye, &target)

	if _z.LengthSqr() == 0 {
		// eye and target are in the same position
		_z[2] = 1
	}

	_z.Normalize()
	_x := vec3.Cross(&up, &_z)

	if _x.LengthSqr() == 0 {
		// up and z are parallel
		if math.Abs(float64(up[2])) == 1 {
			_z[0] += 0.0001
		} else {
			_z[2] += 0.0001
		}

		_z.Normalize()
		_x = vec3.Cross(&up, &_z)
	}

	_x.Normalize()
	_y := vec3.Cross(&_z, &_x)

	mat[0][0] = _x[0]
	mat[0][1] = _y[0]
	mat[0][2] = _z[0]
	mat[1][0] = _x[1]
	mat[1][1] = _y[1]
	mat[1][2] = _z[1]
	mat[2][0] = _x[2]
	mat[2][1] = _y[2]
	mat[2][2] = _z[2]
	mat.Transpose()
	return mat
}

func AssignMul[T float64 | float32](a, b *Mat[T]) *Mat[T] {
	mat := Mat[T]{
		vec4.Vec[T]{1, 0, 0, 0},
		vec4.Vec[T]{0, 1, 0, 0},
		vec4.Vec[T]{0, 0, 1, 0},
		vec4.Vec[T]{0, 0, 0, 1},
	}
	mat[0] = a.MulVec4(&b[0])
	mat[1] = a.MulVec4(&b[1])
	mat[2] = a.MulVec4(&b[2])
	mat[3] = a.MulVec4(&b[3])
	return &mat
}
