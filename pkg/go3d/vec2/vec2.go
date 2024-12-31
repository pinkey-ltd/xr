// Package vec2 contains a 2D float vector type Vec and functions.
package vec2

import (
	"fmt"
	"math"

	"pinkey.ltd/xr/pkg/go3d"
)

// Vec represents a 2D vector.
type Vec[T float64 | float32] [2]T

// NewVec2 creates and returns a pointer to a new Vec[T] instance with the given x and y components.
// The generic type parameter T restricts the type to either float64 or float32.
func NewVec2[T float64 | float32](x, y T) *Vec[T] {
	return &Vec[T]{x, y}
}

// From copies a Vec[T] from a go3d.T[Vec[T]] implementation.
func From[T float64 | float32](other go3d.T[T]) Vec[T] {
	return Vec[T]{other.Get(0, 0), other.Get(0, 1)}
}

// Parse parses Vec from a string. See also String()
func Parse[T float64 | float32](s string) (r Vec[T], err error) {
	_, err = fmt.Sscan(s, &r[0], &r[1])
	return r, err
}

// String formats Vec as string. See also Parse().
func (vec *Vec[T]) String() string {
	return fmt.Sprint(vec[0], vec[1])
}

// Clone creates a new copy of the Vec[T].
func (vec *Vec[T]) Clone() *Vec[T] {
	return &Vec[T]{vec[0], vec[1]}
}

// Rows returns the number of rows of the vector.
func (vec *Vec[T]) Rows() int {
	return 2
}

// Cols returns the number of columns of the vector.
func (vec *Vec[T]) Cols() int {
	return 1
}

// Size returns the number elements of the vector.
func (vec *Vec[T]) Size() int {
	return 2
}

// Slice returns the elements of the vector as slice.
func (vec *Vec[T]) Slice() []T {
	return vec[:]
}

// Get returns one element of the vector.
func (vec *Vec[T]) Get(col, row int) T {
	return vec[row]
}

// IsZero checks if all elements of the vector are zero.
func (vec *Vec[T]) IsZero() bool {
	return vec[0] == 0 && vec[1] == 0
}

// Length returns the length of the vector.
// See also LengthSqr and Normalize.
func (vec *Vec[T]) Length() T {
	return T(math.Hypot(float64(vec[0]), float64(vec[1])))
}

// LengthSqr returns the squared length of the vector.
// See also Length and Normalize.
func (vec *Vec[T]) LengthSqr() T {
	return vec[0]*vec[0] + vec[1]*vec[1]
}

// Scale multiplies all element of the vector by f and returns vec.
func (vec *Vec[T]) Scale(f T) *Vec[T] {
	vec[0] *= f
	vec[1] *= f
	return vec
}

// Scaled returns a copy of vec with all elements multiplies by f.
func (vec *Vec[T]) Scaled(f T) Vec[T] {
	return Vec[T]{vec[0] * f, vec[1] * f}
}

// PracticallyEquals compares two vectors if they are equal with each other within a delta tolerance.
func (vec *Vec[T]) PracticallyEquals(compareVector *Vec[T], allowedDelta T) bool {
	return (math.Abs(float64(vec[0]-compareVector[0])) <= float64(allowedDelta)) &&
		(math.Abs(float64(vec[1]-compareVector[1])) <= float64(allowedDelta))
}

// PracticallyEquals compares two values if they are equal with each other within a delta tolerance.
func PracticallyEquals[T float64 | float32](v1, v2, allowedDelta T) bool {
	return math.Abs(float64(v1-v2)) <= float64(allowedDelta)
}

// Invert inverts the vector.
func (vec *Vec[T]) Invert() *Vec[T] {
	vec[0] = -vec[0]
	vec[1] = -vec[1]
	return vec
}

// Inverted returns an inverted copy of the vector.
func (vec *Vec[T]) Inverted() Vec[T] {
	return Vec[T]{-vec[0], -vec[1]}
}

// Abs sets every component of the vector to its absolute value.
func (vec *Vec[T]) Abs() *Vec[T] {
	vec[0] = T(math.Abs(float64(vec[0])))
	vec[1] = T(math.Abs(float64(vec[1])))
	return vec
}

// Absed returns a copy of the vector containing the absolute values.
func (vec *Vec[T]) Absed() Vec[T] {
	return Vec[T]{T(math.Abs(float64(vec[0]))), T(math.Abs(float64(vec[1])))}
}

// Normalize normalizes the vector to unit length.
func (vec *Vec[T]) Normalize() *Vec[T] {
	sl := vec.LengthSqr()
	if sl == 0 || sl == 1 {
		return vec
	}
	return vec.Scale(1 / T(math.Sqrt(float64(sl))))
}

// Normalized returns a unit length normalized copy of the vector.
func (vec *Vec[T]) Normalized() Vec[T] {
	v := *vec
	v.Normalize()
	return v
}

// Normal returns a new normalized orthogonal vector.
// The normal is orthogonal clockwise to the vector.
// See also function Rotate90DegRight.
func (vec *Vec[T]) Normal() Vec[T] {
	n := *vec
	n[0], n[1] = n[1], -n[0]
	return *n.Normalize()
}

// NormalCCW returns a new normalized orthogonal vector.
// The normal is orthogonal counterclockwise to the vector.
// See also function Rotate90DegLeft.
func (vec *Vec[T]) NormalCCW() Vec[T] {
	n := *vec
	n[0], n[1] = -n[1], n[0]
	return *n.Normalize()
}

// Add adds another vector to vec.
func (vec *Vec[T]) Add(v *Vec[T]) *Vec[T] {
	vec[0] += v[0]
	vec[1] += v[1]
	return vec
}

// Added adds another vector to vec and returns a copy of the result
func (vec *Vec[T]) Added(v *Vec[T]) Vec[T] {
	return Vec[T]{vec[0] + v[0], vec[1] + v[1]}
}

// Sub subtracts another vector from vec.
func (vec *Vec[T]) Sub(v *Vec[T]) *Vec[T] {
	vec[0] -= v[0]
	vec[1] -= v[1]
	return vec
}

// Subed subtracts another vector from vec and returns a copy of the result
func (vec *Vec[T]) Subed(v *Vec[T]) Vec[T] {
	return Vec[T]{vec[0] - v[0], vec[1] - v[1]}
}

// Mul multiplies the components of the vector with the respective components of v.
func (vec *Vec[T]) Mul(v *Vec[T]) *Vec[T] {
	vec[0] *= v[0]
	vec[1] *= v[1]
	return vec
}

// Muled multiplies the components of the vector with the respective components of v and returns a copy of the result
func (vec *Vec[T]) Muled(v *Vec[T]) Vec[T] {
	return Vec[T]{vec[0] * v[0], vec[1] * v[1]}
}

// Rotate rotates the vector counter-clockwise by angle.
func (vec *Vec[T]) Rotate(angle T) *Vec[T] {
	*vec = vec.Rotated(angle)
	return vec
}

// Rotated returns a counter-clockwise rotated copy of the vector.
func (vec *Vec[T]) Rotated(angle T) Vec[T] {
	sinA := T(math.Sin(float64(angle)))
	cosA := T(math.Cos(float64(angle)))
	return Vec[T]{
		vec[0]*cosA - vec[1]*sinA,
		vec[0]*sinA + vec[1]*cosA,
	}
}

// RotateAroundPoint rotates the vector counter-clockwise around a point.
func (vec *Vec[T]) RotateAroundPoint(point *Vec[T], angle T) *Vec[T] {
	return vec.Sub(point).Rotate(angle).Add(point)
}

// Rotate90DegLeft rotates the vector 90 degrees left (counter-clockwise).
func (vec *Vec[T]) Rotate90DegLeft() *Vec[T] {
	temp := vec[0]
	vec[0] = -vec[1]
	vec[1] = temp
	return vec
}

// Rotate90DegRight rotates the vector 90 degrees right (clockwise).
func (vec *Vec[T]) Rotate90DegRight() *Vec[T] {
	temp := vec[0]
	vec[0] = vec[1]
	vec[1] = -temp
	return vec
}

// Sinus returns the sinus value of the (shortest/smallest) angle between the two vectors a and b.
// The returned sine value is in the range -1.0 ≤ value ≤ 1.0.
// The angle is always considered to be in the range 0 to Pi radians and thus the sine value returned is always positive.
func Sinus[T float64 | float32](a, b *Vec[T]) T {
	v := float64(cross(a, b)) / math.Sqrt(float64(a.LengthSqr()*b.LengthSqr()))

	if v > 1.0 {
		return 1.0
	} else if v < -1.0 {
		return -1.0
	}
	return T(v)
}

// Cosine returns the cosine value of the angle between the two vectors.
// The returned cosine value is in the range -1.0 ≤ value ≤ 1.0.
func Cosine[T float64 | float32](a, b *Vec[T]) T {
	v := float64(dot(a, b)) / math.Sqrt(float64(a.LengthSqr()*b.LengthSqr()))

	if v > 1.0 {
		return 1.0
	} else if v < -1.0 {
		return -1.0
	}
	return T(v)
}

// Angle returns the counter-clockwise angle of the vector from the x axis.
func (vec *Vec[T]) Angle() T {
	return T(math.Atan2(float64(vec[1]), float64(vec[0])))
}

// Add returns the sum of two vectors.
func Add[T float64 | float32](a, b *Vec[T]) Vec[T] {
	return Vec[T]{a[0] + b[0], a[1] + b[1]}
}

// Sub returns the difference of two vectors.
func Sub[T float64 | float32](a, b *Vec[T]) Vec[T] {
	return Vec[T]{a[0] - b[0], a[1] - b[1]}
}

// Mul returns the component wise product of two vectors.
func Mul[T float64 | float32](a, b *Vec[T]) Vec[T] {
	return Vec[T]{a[0] * b[0], a[1] * b[1]}
}

// Dot returns the dot product of two vectors.
func Dot[T float64 | float32](a, b *Vec[T]) T {
	return a[0]*b[0] + a[1]*b[1]
}

// dot returns the dot product of two vectors.
func dot[T float64 | float32](a, b *Vec[T]) T {
	return a[0]*b[0] + a[1]*b[1]
}

// Cross returns the cross product of two vectors.
func Cross[T float64 | float32](a, b *Vec[T]) Vec[T] {
	return Vec[T]{
		a[1]*b[0] - a[0]*b[1],
		a[0]*b[1] - a[1]*b[0],
	}
}

// cross returns the "cross product" of two vectors.
// In 2D space it is a scalar value.
// It is the same as the determinant value of the 2D matrix constructed by the two vectors.
// Cross product in 2D is not well-defined but this is the implementation stated at https://mathworld.wolfram.com/CrossProduct.html .
func cross[T float64 | float32](a, b *Vec[T]) T {
	return a[0]*b[1] - a[1]*b[0]
}

// Angle returns the angle between two vectors.
func Angle[T float64 | float32](a, b *Vec[T]) T {
	v := Dot(a, b) / (a.Length() * b.Length())
	// prevent NaN
	if v > 1. {
		v = v - 2
	} else if v < -1. {
		v = v + 2
	}
	return T(math.Acos(float64(v)))
}

// angle returns the angle value of the (shortest/smallest) angle between the two vectors a and b.
// The returned value is in the range 0 ≤ angle ≤ Pi radians.
func angle[T float64 | float32](a, b *Vec[T]) T {
	return T(math.Acos(float64(Cosine(a, b))))
}

// IsLeftWinding returns if the angle from a to b is left winding.
func IsLeftWinding[T float64 | float32](a, b *Vec[T]) bool {
	ab := b.Rotated(-a.Angle())
	return ab.Angle() > 0
}

// isLeftWinding returns if the angle from a to b is left winding.
// Two parallell or anti parallell vectors will give a false result.
func isLeftWinding[T float64 | float32](a, b *Vec[T]) bool {
	return cross(a, b) > 0 // It's really the sign changing part of the Sinus(a, b) function
}

// IsRightWinding returns if the angle from a to b is right winding.
func IsRightWinding[T float64 | float32](a, b *Vec[T]) bool {
	ab := b.Rotated(-a.Angle())
	return ab.Angle() < 0
}

// isRightWinding returns if the angle from a to b is right winding.
// Two parallell or anti parallell vectors will give a false result.
func isRightWinding[T float64 | float32](a, b *Vec[T]) bool {
	return cross(a, b) < 0 // It's really the sign changing part of the Sinus(a, b) function
}

// Min returns the component wise minimum of two vectors.
func Min[T float64 | float32](a, b *Vec[T]) Vec[T] {
	minVal := *a
	if b[0] < minVal[0] {
		minVal[0] = b[0]
	}
	if b[1] < minVal[1] {
		minVal[1] = b[1]
	}
	return minVal
}

// Max returns the component wise maximum of two vectors.
func Max[T float64 | float32](a, b *Vec[T]) Vec[T] {
	maxVal := *a
	if b[0] > maxVal[0] {
		maxVal[0] = b[0]
	}
	if b[1] > maxVal[1] {
		maxVal[1] = b[1]
	}
	return maxVal
}

// Interpolate interpolates between a and b at t (0,1).
func Interpolate[T float64 | float32](a, b *Vec[T], t T) Vec[T] {
	t1 := 1 - t
	return Vec[T]{
		a[0]*t1 + b[0]*t,
		a[1]*t1 + b[1]*t,
	}
}

// Clamp clamps the vector's components to be in the range of min to max.
func (vec *Vec[T]) Clamp(min, max *Vec[T]) *Vec[T] {
	for i := range vec {
		if vec[i] < min[i] {
			vec[i] = min[i]
		} else if vec[i] > max[i] {
			vec[i] = max[i]
		}
	}
	return vec
}

// Clamped returns a copy of the vector with the components clamped to be in the range of min to max.
func (vec *Vec[T]) Clamped(min, max *Vec[T]) Vec[T] {
	res := *vec
	res.Clamp(min, max)
	return res
}

// Clamp01 clamps the vector's components to be in the range of 0 to 1.
func (vec *Vec[T]) Clamp01() *Vec[T] {
	zero := &Vec[T]{}
	unitXY := &Vec[T]{1, 1}
	return vec.Clamp(zero, unitXY)
}

// Clamped01 returns a copy of the vector with the components clamped to be in the range of 0 to 1.
func (vec *Vec[T]) Clamped01() Vec[T] {
	result := *vec
	result.Clamp01()
	return result
}

func (vec *Vec[T]) SetMin(c Vec[T]) {
	if c[0] < vec[0] {
		vec[0] = c[0]
	}
	if c[1] < vec[1] {
		vec[1] = c[1]
	}
}

func (vec *Vec[T]) SetMax(c Vec[T]) {
	if c[0] > vec[0] {
		vec[0] = c[0]
	}
	if c[1] > vec[1] {
		vec[1] = c[1]
	}
}

func PointSegmentDistance[T float64 | float32](p1 *Vec[T], x1 *Vec[T], x2 *Vec[T]) T {
	v := Sub(x2, x1)
	w := Sub(p1, x1)
	if w.IsZero() {
		return 0
	}

	vn := v.Normalized()
	wn := w.Normalized()
	cn := Dot(&vn, &wn)
	if math.Abs(math.Abs(float64(cn))-1) < 1e-8 {
		return 0
	}
	return T(w.Length()) * T(math.Sqrt(float64(1-cn*cn)))
}

func PointSegmentVerticalPoint[T float64 | float32](p1 *Vec[T], x1 *Vec[T], x2 *Vec[T]) *Vec[T] {
	v := Sub(x2, x1)
	w := Sub(p1, x1)
	if w.IsZero() {
		return &Vec[T]{p1[0], p1[1]}
	}

	vn := v.Normalized()
	wn := w.Normalized()
	cn := Dot(&vn, &wn)
	if math.Abs(math.Abs(float64(cn))-1) < 1e-8 {
		return &Vec[T]{p1[0], p1[1]}
	}
	c1 := Dot(&w, &v)
	c2 := Dot(&v, &v)
	b := c1 / c2
	v1 := v.Scaled(b)
	res := Add(x1, &v1)
	return &res
}
