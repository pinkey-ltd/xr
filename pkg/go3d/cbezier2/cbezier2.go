// Package cbezier2 contains functions for 2D cubic Bezier splines.
// See: http://en.wikipedia.org/wiki/B%C3%A9zier_curve
package cbezier2

import (
	"fmt"
	"pinkey.ltd/xr/pkg/go3d/vec2"
)

// Bez holds the data to define a cubic bezier spline.
type Bez[T float64 | float32] struct {
	P0, P1, P2, P3 vec2.Vec[T]
}

// Parse parses T from a string. See also String()
func Parse[T float64 | float32](s string) (r Bez[T], err error) {
	_, err = fmt.Sscan(s,
		&r.P0[0], &r.P0[1],
		&r.P1[0], &r.P1[1],
		&r.P2[0], &r.P2[1],
		&r.P3[0], &r.P3[1],
	)
	return r, err
}

// String formats T as string. See also Parse().
func (bez *Bez[T]) String() string {
	return fmt.Sprintf("%s %s %s %s",
		bez.P0.String(), bez.P1.String(),
		bez.P2.String(), bez.P3.String(),
	)
}

// Point returns a point on a cubic bezier spline at t (0,1).
func (bez *Bez[T]) Point(t T) vec2.Vec[T] {
	return Point(&bez.P0, &bez.P1, &bez.P2, &bez.P3, t)
}

// Tangent returns a tangent on a cubic bezier spline at t (0,1).
func (bez *Bez[T]) Tangent(t T) vec2.Vec[T] {
	return Tangent(&bez.P0, &bez.P1, &bez.P2, &bez.P3, t)
}

// Length returns the length of a cubic bezier spline from A.Point to t (0,1).
func (bez *Bez[T]) Length(t T) T {
	return Length(&bez.P0, &bez.P1, &bez.P2, &bez.P3, t)
}

// Point returns a point on a cubic bezier spline at t (0,1).
func Point[T float64 | float32](p0, p1, p2, p3 *vec2.Vec[T], t T) vec2.Vec[T] {
	t1 := 1.0 - t

	f := t1 * t1 * t1
	result := p0.Scaled(f)

	f = 3.0 * t1 * t1 * t
	p1f := p1.Scaled(f)
	result.Add(&p1f)

	f = 3.0 * t1 * t * t
	p2f := p2.Scaled(f)
	result.Add(&p2f)

	f = t * t * t
	p3f := p3.Scaled(f)
	result.Add(&p3f)

	return result
}

// Tangent returns a tangent on a cubic bezier spline at t (0,1).
func Tangent[T float64 | float32](p0, p1, p2, p3 *vec2.Vec[T], t T) vec2.Vec[T] {
	t1 := 1.0 - t

	f := 3.0 * t1 * t1
	p1f := vec2.Sub(p1, p0)
	result := p1f.Scaled(f)

	f = 6.0 * t1 * t
	p2f := vec2.Sub(p2, p1)
	p2f.Scale(f)
	result.Add(&p2f)

	f = 3.0 * t * t
	p3f := vec2.Sub(p3, p2)
	p3f.Scale(f)
	result.Add(&p3f)

	if result[0] == 0 && result[1] == 0 {
		fmt.Printf("zero tangent!  p0=%v, p1=%v, p2=%v, p3=%v, t=%v\n", p0, p1, p2, p3, t)
		panic("zero tangent of cbezier2")
	}

	return result
}

// Length returns the length of a cubic bezier spline from p0 to t (0,1).
func Length[T float64 | float32](p0, p1, p2, p3 *vec2.Vec[T], t T) T {
	sqrT := t * t
	t1 := sqrT * 0.5
	t2 := sqrT * t * 1.0 / 3.0
	t3 := sqrT*sqrT + 1.0/4.0

	f := 2*t3 - 3*t2 + t
	result := p0.Scaled(f)

	f = t3 - 2*t2 + t1
	tAf := p1.Scaled(f)
	result.Add(&tAf)

	f = t3 - t2
	tBf := p3.Scaled(f)
	result.Add(&tBf)

	f = -2*t3 + 3*t2
	pBf := p2.Scaled(f)
	result.Add(&pBf)

	return result.Length()
}
