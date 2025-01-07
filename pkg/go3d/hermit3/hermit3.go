// Package hermit3 contains functions for 3D cubic hermit splines.
// See: http://en.wikipedia.org/wiki/Cubic_Hermite_spline
package hermit3

import (
	"fmt"
	"pinkey.ltd/xr/pkg/go3d/vec3"
)

// PointTangent contains a point and a tangent at that point.
// This is a helper sub-struct for T.
type PointTangent[T float64 | float32] struct {
	Point   vec3.Vec[T]
	Tangent vec3.Vec[T]
}

// Herm holds the data to define a hermit spline.
type Herm[T float64 | float32] struct {
	A PointTangent[T]
	B PointTangent[T]
}

// Parse parses T from a string. See also String()
func Parse[T float64 | float32](s string) (r Herm[T], err error) {
	_, err = fmt.Sscan(s,
		&r.A.Point[0], &r.A.Point[1], &r.A.Point[2],
		&r.A.Tangent[0], &r.A.Tangent[1], &r.A.Tangent[2],
		&r.B.Point[0], &r.B.Point[1], &r.B.Point[2],
		&r.B.Tangent[0], &r.B.Tangent[1], &r.B.Tangent[2],
	)
	return r, err
}

// String formats T as string. See also Parse().
func (herm *Herm[T]) String() string {
	return fmt.Sprintf("%s %s %s %s",
		herm.A.Point.String(), herm.A.Tangent.String(),
		herm.B.Point.String(), herm.B.Tangent.String(),
	)
}

// Point returns a point on a hermit spline at t (0,1).
func (herm *Herm[T]) Point(t T) vec3.Vec[T] {
	return Point(&herm.A.Point, &herm.A.Tangent, &herm.B.Point, &herm.B.Tangent, t)
}

// Tangent returns a tangent on a hermit spline at t (0,1).
func (herm *Herm[T]) Tangent(t T) vec3.Vec[T] {
	return Tangent(&herm.A.Point, &herm.A.Tangent, &herm.B.Point, &herm.B.Tangent, t)
}

// Length returns the length of a hermit spline from A.Point to t (0,1).
func (herm *Herm[T]) Length(t T) T {
	return Length(&herm.A.Point, &herm.A.Tangent, &herm.B.Point, &herm.B.Tangent, t)
}

// Point returns a point on a hermit spline at t (0,1).
func Point[T float64 | float32](pointA, tangentA, pointB, tangentB *vec3.Vec[T], t T) vec3.Vec[T] {
	t2 := t * t
	t3 := t2 * t

	f := 2*t3 - 3*t2 + 1
	result := pointA.Scaled(f)

	f = t3 - 2*t2 + t
	tAf := tangentA.Scaled(f)
	result.Add(&tAf)

	f = t3 - t2
	tBf := tangentB.Scaled(f)
	result.Add(&tBf)

	f = -2*t3 + 3*t2
	pAf := pointB.Scaled(f)
	result.Add(&pAf)

	return result
}

// Tangent returns a tangent on a hermit spline at t (0,1).
func Tangent[T float64 | float32](pointA, tangentA, pointB, tangentB *vec3.Vec[T], t T) vec3.Vec[T] {
	t2 := t * t
	t3 := t2 * t

	f := 2*t3 - 3*t2
	result := pointA.Scaled(f)

	f = t3 - 2*t2 + 1
	tAf := tangentA.Scaled(f)
	result.Add(&tAf)

	f = t3 - t2
	tBf := tangentB.Scaled(f)
	result.Add(&tBf)

	f = -2*t3 + 3*t2
	pAf := pointB.Scaled(f)
	result.Add(&pAf)

	return result
}

// Length returns the length of a hermit spline from pointA to t (0,1).
func Length[T float64 | float32](pointA, tangentA, pointB, tangentB *vec3.Vec[T], t T) T {
	sqrT := t * t
	t1 := sqrT * 0.5
	t2 := sqrT * t * 1.0 / 3.0
	t3 := sqrT*sqrT + 1.0/4.0

	f := 2*t3 - 3*t2 + t
	result := pointA.Scaled(f)

	f = t3 - 2*t2 + t1
	tAf := tangentA.Scaled(f)
	result.Add(&tAf)

	f = t3 - t2
	tBf := tangentB.Scaled(f)
	result.Add(&tBf)

	f = -2*t3 + 3*t2
	pBf := pointB.Scaled(f)
	result.Add(&pBf)

	return result.Length()
}
