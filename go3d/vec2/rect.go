package vec2

import (
	"fmt"
	"math"
)

// Rect is a coordinate system aligned rectangle defined by a Min and Max vector.
type Rect[T float64 | float32] struct {
	Min Vec[T]
	Max Vec[T]
}

// NewRect creates a Rect from two points.
func NewRect[T float64 | float32](a, b *Vec[T]) (rect Rect[T]) {
	rect.Min = Min(a, b)
	rect.Max = Max(a, b)
	return rect
}

// ParseRect parses a Rect from a string. See also String()
func ParseRect[T float64 | float32](s string) (r Rect[T], err error) {
	_, err = fmt.Sscan(s, &r.Min[0], &r.Min[1], &r.Max[0], &r.Max[1])
	return r, err
}

func (rect *Rect[T]) Width() T {
	return T(math.Abs(float64(rect.Max[0] - rect.Min[0])))
}

func (rect *Rect[T]) Height() T {
	return T(math.Abs(float64(rect.Max[1] - rect.Min[1])))
}

func (rect *Rect[T]) Size() T {
	width := rect.Width()
	height := rect.Height()
	return T(math.Max(float64(width), float64(height)))
}

// Slice returns the elements of the vector as slice.
func (rect *Rect[T]) Slice() []T {
	return rect.Array()[:]
}

func (rect *Rect[T]) Array() *[4]T {
	return &[...]T{
		rect.Min[0], rect.Min[1],
		rect.Max[0], rect.Max[1],
	}
}

// String formats Rect as string. See also ParseRect().
func (rect *Rect[T]) String() string {
	return rect.Min.String() + " " + rect.Max.String()
}

// ContainsPoint returns if a point is contained within the rectangle.
func (rect *Rect[T]) ContainsPoint(p *Vec[T]) bool {
	return p[0] >= rect.Min[0] && p[0] <= rect.Max[0] &&
		p[1] >= rect.Min[1] && p[1] <= rect.Max[1]
}

// Contains returns if other Rect is contained within the rectangle.
func (rect *Rect[T]) Contains(other *Rect[T]) bool {
	return rect.Min[0] <= other.Min[0] &&
		rect.Min[1] <= other.Min[1] &&
		rect.Max[0] >= other.Max[0] &&
		rect.Max[1] >= other.Max[1]
}

// Area calculates the area of the rectangle.
func (rect *Rect[T]) Area() T {
	return (rect.Max[0] - rect.Min[0]) * (rect.Max[1] - rect.Min[1])
}

func (rect *Rect[T]) Intersects(other *Rect[T]) bool {
	return other.Max[0] >= rect.Min[0] &&
		other.Min[0] <= rect.Max[0] &&
		other.Max[1] >= rect.Min[1] &&
		other.Min[1] <= rect.Max[1]
}

// Join enlarges this rectangle to contain also the given rectangle.
func (rect *Rect[T]) Join(other *Rect[T]) {
	rect.Min = Min(&rect.Min, &other.Min)
	rect.Max = Max(&rect.Max, &other.Max)
}

func (rect *Rect[T]) Extend(p *Vec[T]) {
	rect.Min = Min(&rect.Min, p)
	rect.Max = Max(&rect.Max, p)
}

// Joined returns the minimal rectangle containing both a and b.
func Joined[T float64 | float32](a, b *Rect[T]) (rect Rect[T]) {
	rect.Min = Min(&a.Min, &b.Min)
	rect.Max = Max(&a.Max, &b.Max)
	return rect
}
