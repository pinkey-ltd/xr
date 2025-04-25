// Package quaternion contains a float unit-quaternion type H and functions.
package quaternion

import (
	"fmt"
	"math"

	"pinkey.ltd/xr/go3d/vec3"
	"pinkey.ltd/xr/go3d/vec4"
)

// H represents an orientatin/rotation as a unit quaternion.
// See http://en.wikipedia.org/wiki/Quaternions_and_spatial_rotation
type H[T float64 | float32] [4]T

// FromAxisAngle returns a quaternion representing a rotation around and axis.
func FromAxisAngle[T float64 | float32](axis *vec3.Vec[T], angle float64) H[T] {
	angle *= 0.5
	sin := T(math.Sin(angle))
	q := H[T]{axis[0] * sin, axis[1] * sin, axis[2] * sin, T(math.Cos(angle))}
	return q.Normalized()
}

// FromXAxisAngle returns a quaternion representing a rotation around the x axis.
func FromXAxisAngle[T float64 | float32](angle T) H[T] {
	angle *= 0.5
	return H[T]{T(math.Sin(float64(angle))), 0, 0, T(math.Cos(float64(angle)))}
}

// FromYAxisAngle returns a quaternion representing a rotation around the y axis.
func FromYAxisAngle[T float64 | float32](angle T) H[T] {
	angle *= 0.5
	return H[T]{0, T(math.Sin(float64(angle))), 0, T(math.Cos(float64(angle)))}
}

// FromZAxisAngle returns a quaternion representing a rotation around the z axis.
func FromZAxisAngle[T float64 | float32](angle T) H[T] {
	angle *= 0.5
	return H[T]{0, 0, T(math.Sin(float64(angle))), T(math.Cos(float64(angle)))}
}

// FromEulerAngles returns a quaternion representing Euler angle(in radian) rotations.
func FromEulerAngles[T float64 | float32](yHead, xPitch, zRoll T) H[T] {
	qy := FromYAxisAngle[T](yHead)
	qx := FromXAxisAngle[T](xPitch)
	qz := FromZAxisAngle[T](zRoll)
	return Mul3(&qy, &qx, &qz)
}

// ToEulerAngles returns the euler angle(in radian) rotations of the quaternion.
func (quat *H[T]) ToEulerAngles() (yHead, xPitch, zRoll T) {
	z := quat.RotatedVec3(&vec3.Vec[T]{0, 0, 1})
	yHead = T(math.Atan2(float64(z[0]), float64(z[2])))
	xPitch = T(-math.Atan2(float64(z[1]), math.Sqrt(float64(z[0]*z[0]+z[2]*z[2]))))

	quatNew := FromEulerAngles[T](yHead, xPitch, 0)

	x2 := quatNew.RotatedVec3(&vec3.Vec[T]{1, 0, 0})
	x := quat.RotatedVec3(&vec3.Vec[T]{1, 0, 0})
	r := vec3.Cross(&x, &x2)
	sin := float64(vec3.Dot(&r, &z))
	cos := float64(vec3.Dot(&x, &x2))
	zRoll = T(-math.Atan2(sin, cos))
	return
}

// FromVec4 converts a vec4.H into a quaternion.
func FromVec4[T float64 | float32](v *vec4.Vec[T]) H[T] {
	return H[T](*v)
}

// Vec4 converts the quaternion into a vec4.H.
func (quat *H[T]) Vec4() vec4.Vec[T] {
	return vec4.Vec[T](*quat)
}

// Parse parses H from a string. See also String()
func Parse[T float64 | float32](s string) (r H[T], err error) {
	_, err = fmt.Sscan(s, &r[0], &r[1], &r[2], &r[3])
	return r, err
}

// String formats H as string. See also Parse().
func (quat *H[T]) String() string {
	return fmt.Sprint(quat[0], quat[1], quat[2], quat[3])
}

// AxisAngle extracts the rotation from a normalized quaternion in the form of an axis and a rotation angle.
func (quat *H[T]) AxisAngle() (axis vec3.Vec[T], angle T) {
	cos := float64(quat[3])
	sin := T(math.Sqrt(1 - cos*cos))
	angle = T(math.Acos(cos) * 2)

	var ooSin T
	if math.Abs(float64(sin)) < 0.0005 {
		ooSin = 1
	} else {
		ooSin = 1 / sin
	}
	axis[0] = quat[0] * ooSin
	axis[1] = quat[1] * ooSin
	axis[2] = quat[2] * ooSin

	return axis, angle * 2
}

// Norm returns the norm value of the quaternion.
func (quat *H[T]) Norm() T {
	return quat[0]*quat[0] + quat[1]*quat[1] + quat[2]*quat[2] + quat[3]*quat[3]
}

// Normalize normalizes to a unit quaternation.
func (quat *H[T]) Normalize() *H[T] {
	norm := quat.Norm()
	if norm != 1 && norm != 0 {
		ool := T(1 / math.Sqrt(float64(norm)))
		quat[0] *= ool
		quat[1] *= ool
		quat[2] *= ool
		quat[3] *= ool
	}
	return quat
}

// Normalized returns a copy normalized to a unit quaternation.
func (quat *H[T]) Normalized() H[T] {
	norm := quat.Norm()
	if norm != 1 && norm != 0 {
		ool := T(1 / math.Sqrt(float64(norm)))
		return H[T]{
			quat[0] * ool,
			quat[1] * ool,
			quat[2] * ool,
			quat[3] * ool,
		}
	} else {
		return *quat
	}
}

// Negate negates the quaternion.
func (quat *H[T]) Negate() *H[T] {
	quat[0] = -quat[0]
	quat[1] = -quat[1]
	quat[2] = -quat[2]
	quat[3] = -quat[3]
	return quat
}

// Negated returns a negated copy of the quaternion.
func (quat *H[T]) Negated() H[T] {
	return H[T]{-quat[0], -quat[1], -quat[2], -quat[3]}
}

// Invert inverts the quaterion.
func (quat *H[T]) Invert() *H[T] {
	quat[0] = -quat[0]
	quat[1] = -quat[1]
	quat[2] = -quat[2]
	return quat
}

// IsIdent checks if the quaternion represents the identity rotation (no rotation).
// It returns true if the quaternion is (0, 0, 0, 1), indicating no rotation.
func (quat *H[T]) IsIdent() bool {
	return quat[0] == 0 && quat[1] == 0 && quat[2] == 0 && quat[3] == 1
}

// Inverted returns an inverted copy of the quaternion.
func (quat *H[T]) Inverted() H[T] {
	return H[T]{-quat[0], -quat[1], -quat[2], quat[3]}
}

// SetShortestRotation negates the quaternion if it does not represent the shortest rotation from quat to the orientation of other.
// (there are two directions to rotate from the orientation of quat to the orientation of other)
// See IsShortestRotation()
func (quat *H[T]) SetShortestRotation(other *H[T]) *H[T] {
	if !IsShortestRotation(quat, other) {
		quat.Negate()
	}
	return quat
}

// IsShortestRotation returns if the rotation from a to b is the shortest possible rotation.
// (there are two directions to rotate from the orientation of quat to the orientation of other)
// See H.SetShortestRotation
func IsShortestRotation[T float64 | float32](a, b *H[T]) bool {
	return Dot(a, b) >= 0
}

// IsUnitQuat returns if the quaternion is within tolerance of the unit quaternion.
func (quat *H[T]) IsUnitQuat(tolerance T) bool {
	norm := quat.Norm()
	return norm >= (1.0-tolerance) && norm <= (1.0+tolerance)
}

// RotateVec3 rotates v by the rotation represented by the quaternion.
// using the algorithm mentioned here https://gamedev.stackexchange.com/questions/28395/rotating-vector3-by-a-quaternion
func (quat *H[T]) RotateVec3(v *vec3.Vec[T]) {
	u := vec3.Vec[T]{quat[0], quat[1], quat[2]}
	s := quat[3]
	vt1 := u.Scaled(2 * vec3.Dot(&u, v))
	vt2 := v.Scaled(s*s - vec3.Dot(&u, &u))
	vt3 := vec3.Cross(&u, v)
	vt3 = vt3.Scaled(2 * s)
	v[0] = vt1[0] + vt2[0] + vt3[0]
	v[1] = vt1[1] + vt2[1] + vt3[1]
	v[2] = vt1[2] + vt2[2] + vt3[2]
}

// RotatedVec3 returns a rotated copy of v.
// using the algorithm mentioned here https://gamedev.stackexchange.com/questions/28395/rotating-vector3-by-a-quaternion
func (quat *H[T]) RotatedVec3(v *vec3.Vec[T]) vec3.Vec[T] {
	u := vec3.Vec[T]{quat[0], quat[1], quat[2]}
	s := quat[3]
	vt1 := u.Scaled(2 * vec3.Dot(&u, v))
	vt2 := v.Scaled(s*s - vec3.Dot(&u, &u))
	vt3 := vec3.Cross(&u, v)
	vt3 = vt3.Scaled(2 * s)
	return vec3.Vec[T]{vt1[0] + vt2[0] + vt3[0], vt1[1] + vt2[1] + vt3[1], vt1[2] + vt2[2] + vt3[2]}
}

// Dot returns the dot product of two quaternions.
func Dot[T float64 | float32](a, b *H[T]) T {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2] + a[3]*b[3]
}

// Mul multiplies two quaternions.
func Mul[T float64 | float32](a, b *H[T]) H[T] {
	q := H[T]{
		a[3]*b[0] + a[0]*b[3] + a[1]*b[2] - a[2]*b[1],
		a[3]*b[1] + a[1]*b[3] + a[2]*b[0] - a[0]*b[2],
		a[3]*b[2] + a[2]*b[3] + a[0]*b[1] - a[1]*b[0],
		a[3]*b[3] - a[0]*b[0] - a[1]*b[1] - a[2]*b[2],
	}
	return q.Normalized()
}

// Mul3 multiplies three quaternions.
func Mul3[T float64 | float32](a, b, c *H[T]) H[T] {
	q := Mul(a, b)
	return Mul(&q, c)
}

// Mul4 multiplies four quaternions.
func Mul4[T float64 | float32](a, b, c, d *H[T]) H[T] {
	q := Mul(a, b)
	q = Mul(&q, c)
	return Mul(&q, d)
}

// Slerp returns the spherical linear interpolation quaternion between a and b at t (0,1).
// See http://en.wikipedia.org/wiki/Slerp
func Slerp[T float64 | float32](a, b *H[T], t float64) H[T] {
	d := math.Acos(float64(a[0]*b[0] + a[1]*b[1] + a[2]*b[2] + a[3]*b[3]))
	ooSinD := 1 / math.Sin(d)

	t1 := T(math.Sin(d*(1-t)) * ooSinD)
	t2 := T(math.Sin(d*t) * ooSinD)

	q := H[T]{
		a[0]*t1 + b[0]*t2,
		a[1]*t1 + b[1]*t2,
		a[2]*t1 + b[2]*t2,
		a[3]*t1 + b[3]*t2,
	}

	return q.Normalized()
}

// Vec3Diff returns the rotation quaternion between two vectors.
func Vec3Diff[T float64 | float32](a, b *vec3.Vec[T]) H[T] {
	cr := vec3.Cross(a, b)
	sr := T(math.Sqrt(2 * (1 + float64(vec3.Dot(a, b)))))
	oosr := 1 / sr

	q := H[T]{cr[0] * oosr, cr[1] * oosr, cr[2] * oosr, sr * 0.5}
	return q.Normalized()
}
