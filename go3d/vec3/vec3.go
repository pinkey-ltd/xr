// Package vec3 contains a 3D float64 vector type Vec and functions.
package vec3

import (
	"fmt"
	"math"

	"pinkey.ltd/xr/go3d"
)

var (
	//	// Zero holds a zero vector.
	//	Zero = Vec[float64]{}
	//
	//	// UnitX holds a vector with X set to one.
	//	UnitX = Vec{1, 0, 0}
	//	// UnitY holds a vector with Y set to one.
	//	UnitY = Vec{0, 1, 0}
	//	// UnitZ holds a vector with Z set to one.
	//	UnitZ = Vec{0, 0, 1}
	//	// UnitXYZ holds a vector with X, Y, Z set to one.
	//	UnitXYZ = Vec{1, 1, 1}
	//
	//	// Red holds the color red.
	//	Red = Vec{1, 0, 0}
	//	// Green holds the color green.
	//	Green = Vec{0, 1, 0}
	//	// Blue holds the color black.
	//	Blue = Vec{0, 0, 1}
	//	// Black holds the color black.
	//	Black = Vec{0, 0, 0}
	//	// White holds the color white.
	//	White = Vec{1, 1, 1}
	//
	// MinVal holds a vector with the smallest possible component values.
	MinVal = Vec[float64]{-math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64}
	// MaxVal holds a vector with the highest possible component values.
	MaxVal = Vec[float64]{+math.MaxFloat64, +math.MaxFloat64, +math.MaxFloat64}
)

// Vec represents a 3D vector.
type Vec[T float64 | float32] [3]T

// From copies a Vec from a generic.Vec implementation.
func From[T float64 | float32](other go3d.T[T]) Vec[T] {
	switch other.Size() {
	case 2:
		return Vec[T]{other.Get(0, 0), other.Get(0, 1), 0}
	case 3, 4:
		return Vec[T]{other.Get(0, 0), other.Get(0, 1), other.Get(0, 2)}
	default:
		panic("Unsupported type")
	}
}

// Parse parses Vec from a string. See also String()
func Parse[T float64 | float32](s string) (r Vec[T], err error) {
	_, err = fmt.Sscan(s, &r[0], &r[1], &r[2])
	return r, err
}

// String formats Vec as string. See also Parse().
func (vec *Vec[T]) String() string {
	return fmt.Sprint(vec[0], vec[1], vec[2])
}

// Rows returns the number of rows of the vector.
func (vec *Vec[T]) Rows() int {
	return 3
}

// Cols returns the number of columns of the vector.
func (vec *Vec[T]) Cols() int {
	return 1
}

// Size returns the number elements of the vector.
func (vec *Vec[T]) Size() int {
	return 3
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
	return vec[0] == 0 && vec[1] == 0 && vec[2] == 0
}

// Length returns the length of the vector.
// See also LengthSqr and Normalize.
func (vec *Vec[T]) Length() T {
	return T(math.Sqrt(float64(vec.LengthSqr())))
}

// LengthSqr returns the squared length of the vector.
// See also Length and Normalize.
func (vec *Vec[T]) LengthSqr() T {
	return vec[0]*vec[0] + vec[1]*vec[1] + vec[2]*vec[2]
}

// Scale multiplies all element of the vector by f and returns vec.
func (vec *Vec[T]) Scale(f T) *Vec[T] {
	vec[0] *= f
	vec[1] *= f
	vec[2] *= f
	return vec
}

// Scaled returns a copy of vec with all elements multiplies by f.
func (vec *Vec[T]) Scaled(f T) Vec[T] {
	return Vec[T]{vec[0] * f, vec[1] * f, vec[2] * f}
}

// Invert inverts the vector.
func (vec *Vec[T]) Invert() *Vec[T] {
	vec[0] = -vec[0]
	vec[1] = -vec[1]
	vec[2] = -vec[2]
	return vec
}

// Inverted returns an inverted copy of the vector.
func (vec *Vec[T]) Inverted() Vec[T] {
	return Vec[T]{-vec[0], -vec[1], -vec[2]}
}

// Abs sets every component of the vector to its absolute value.
func (vec *Vec[T]) Abs() *Vec[T] {
	vec[0] = T(math.Abs(float64(vec[0])))
	vec[1] = T(math.Abs(float64(vec[1])))
	vec[2] = T(math.Abs(float64(vec[2])))
	return vec
}

// Absed returns a copy of the vector containing the absolute values.
func (vec *Vec[T]) Absed() Vec[T] {
	return Vec[T]{T(math.Abs(float64(vec[0]))), T(math.Abs(float64(vec[1]))), T(math.Abs(float64(vec[2])))}
}

// Normalize normalizes the vector to unit length.
func (vec *Vec[T]) Normalize() *Vec[T] {
	sl := vec.LengthSqr()
	if sl == 0 || sl == 1 {
		return vec
	}
	vec.Scale(1 / T(math.Sqrt(float64(sl))))
	return vec
}

// Normalized returns a unit length normalized copy of the vector.
func (vec *Vec[T]) Normalized() Vec[T] {
	v := *vec
	v.Normalize()
	return v
}

// Normal returns an orthogonal vector.
func (vec *Vec[T]) Normal() Vec[T] {
	unitZ := Vec[T]{0, 0, 1}
	unitX := Vec[T]{1, 0, 0}
	n := Cross(vec, &unitZ)
	if n.IsZero() {
		return unitX
	}
	return n.Normalized()
}

// Add adds another vector to vec.
func (vec *Vec[T]) Add(v *Vec[T]) *Vec[T] {
	vec[0] += v[0]
	vec[1] += v[1]
	vec[2] += v[2]
	return vec
}

// Added adds another vector to vec and returns a copy of the result
func (vec *Vec[T]) Added(v *Vec[T]) Vec[T] {
	return Vec[T]{vec[0] + v[0], vec[1] + v[1], vec[2] + v[2]}
}

// Sub subtracts another vector from vec.
func (vec *Vec[T]) Sub(v *Vec[T]) *Vec[T] {
	vec[0] -= v[0]
	vec[1] -= v[1]
	vec[2] -= v[2]
	return vec
}

// Subed subtracts another vector from vec and returns a copy of the result
func (vec *Vec[T]) Subed(v *Vec[T]) Vec[T] {
	return Vec[T]{vec[0] - v[0], vec[1] - v[1], vec[2] - v[2]}
}

// Mul multiplies the components of the vector with the respective components of v.
func (vec *Vec[T]) Mul(v *Vec[T]) *Vec[T] {
	vec[0] *= v[0]
	vec[1] *= v[1]
	vec[2] *= v[2]
	return vec
}

// Muled multiplies the components of the vector with the respective components of v and returns a copy of the result
func (vec *Vec[T]) Muled(v *Vec[T]) Vec[T] {
	return Vec[T]{vec[0] * v[0], vec[1] * v[1], vec[2] * v[2]}
}

func (vec *Vec[T]) ProjectOnVector(vector *Vec[T]) *Vec[T] {
	scalar := Dot(vector, vec) / vector.LengthSqr()
	copy(vec[:], vector[:])
	return vec.Scale(scalar)
}

func (vec *Vec[T]) ProjectOnPlane(planeNormal *Vec[T]) *Vec[T] {
	v1 := Vec[T]{}
	copy(v1[:], vec[:])
	v1.ProjectOnVector(planeNormal)
	result := Sub(vec, &v1)
	copy(vec[:], result[:])
	return vec
}

// Add returns the sum of two vectors.
func Add[T float64 | float32](a, b *Vec[T]) Vec[T] {
	return Vec[T]{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

// SquareDistance the distance between two vectors squared (= distance*distance)
func SquareDistance[T float64 | float32](a, b *Vec[T]) T {
	d := Sub(a, b)
	return d.LengthSqr()
}

// Distance between two vectors
func Distance[T float64 | float32](a, b *Vec[T]) T {
	d := Sub(a, b)
	return d.Length()
}

// Sub returns the difference of two vectors.
func Sub[T float64 | float32](a, b *Vec[T]) Vec[T] {
	return Vec[T]{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

// Mul returns the component wise product of two vectors.
func Mul[T float64 | float32](a, b *Vec[T]) Vec[T] {
	return Vec[T]{a[0] * b[0], a[1] * b[1], a[2] * b[2]}
}

// Dot returns the dot product of two vectors.
func Dot[T float64 | float32](a, b *Vec[T]) T {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

// Cross returns the cross product of two vectors.
func Cross[T float64 | float32](a, b *Vec[T]) Vec[T] {
	return Vec[T]{
		a[1]*b[2] - a[2]*b[1],
		a[2]*b[0] - a[0]*b[2],
		a[0]*b[1] - a[1]*b[0],
	}
}

// Sinus returns the sinus value of the (shortest/smallest) angle between the two vectors a and b.
// The returned sine value is in the range 0.0 ≤ value ≤ 1.0.
// The angle is always considered to be in the range 0 to Pi radians and thus the sine value returned is always positive.
func Sinus[T float64 | float32](a, b *Vec[T]) T {
	cross := Cross(a, b)
	v := cross.Length() / T(math.Sqrt(float64(a.LengthSqr()*b.LengthSqr())))

	if v > 1.0 {
		return 1.0
	} else if v < 0.0 {
		return 0.0
	}
	return v
}

// Cosine returns the cosine value of the angle between the two vectors.
// The returned cosine value is in the range -1.0 ≤ value ≤ 1.0.
func Cosine[T float64 | float32](a, b *Vec[T]) T {
	v := Dot(a, b) / T(math.Sqrt(float64(a.LengthSqr()*b.LengthSqr())))

	if v > 1.0 {
		return 1.0
	} else if v < -1.0 {
		return -1.0
	}
	return v
}

// Angle returns the angle between two vectors.
func Angle[T float64 | float32](a, b *Vec[T]) T {
	v := Dot(a, b) / (a.Length() * b.Length())
	// prevent NaN
	if v > 1. {
		return 0
	} else if v < -1. {
		return math.Pi
	}
	return T(math.Acos(float64(v)))
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
	if b[2] < minVal[2] {
		minVal[2] = b[2]
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
	if b[2] > maxVal[2] {
		maxVal[2] = b[2]
	}
	return maxVal
}

// Interpolate interpolates between a and b at t (0,1).
func Interpolate[T float64 | float32](a, b *Vec[T], t T) Vec[T] {
	t1 := 1 - t
	return Vec[T]{
		a[0]*t1 + b[0]*t,
		a[1]*t1 + b[1]*t,
		a[2]*t1 + b[2]*t,
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
	result := *vec
	result.Clamp(min, max)
	return result
}

// Clamp01 clamps the vector's components to be in the range of 0 to 1.
func (vec *Vec[T]) Clamp01() *Vec[T] {
	zero := Vec[T]{}
	unitXYZ := Vec[T]{1, 1, 1}
	return vec.Clamp(&zero, &unitXYZ)
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
	if c[2] < vec[2] {
		vec[2] = c[2]
	}
}

func (vec *Vec[T]) SetMax(c Vec[T]) {
	if c[0] > vec[0] {
		vec[0] = c[0]
	}
	if c[1] > vec[1] {
		vec[1] = c[1]
	}
	if c[2] > vec[2] {
		vec[2] = c[2]
	}
}

func Rotated[T float64 | float32](vec *Vec[T], axis *Vec[T], rad T) *Vec[T] {
	sinVal := T(math.Sin(float64(rad)))
	cosVal := T(math.Cos(float64(rad)))

	crs := Cross(vec, axis)
	crs.Scale(sinVal)

	scl := vec.Scaled(cosVal)
	dt := Dot(vec, axis) * (1 - cosVal)
	dv := axis.Scaled(dt)

	crs.Add(&scl)
	crs.Add(&dv)
	return &crs
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
	return w.Length() * T(math.Sqrt(float64(1-cn*cn)))
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
