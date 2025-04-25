package vec3

import (
	"fmt"
)

// Box is a coordinate system aligned 3D box defined by a Min and Max vector.
type Box[T float64 | float32] struct {
	Min Vec[T]
	Max Vec[T]
}

var (
	// MaxBox holds a box that contains the entire R3 space that can be represented as vec3
	MaxBox = Box[float64]{MinVal, MaxVal}
	MinBox = Box[float64]{MaxVal, MinVal}
)

func FromSlice[T float64 | float32](s []T) *Box[T] {
	return &Box[T]{Min: Vec[T]{s[0], s[1], s[2]}, Max: Vec[T]{s[3], s[4], s[5]}}
}

// ParseBox parses a Box from a string. See also String()
func ParseBox[T float64 | float32](s string) (r Box[T], err error) {
	_, err = fmt.Sscan(s, &r.Min[0], &r.Min[1], &r.Min[2], &r.Max[0], &r.Max[1], &r.Max[2])
	return r, err
}

// String formats Box as string. See also ParseBox().
func (box *Box[T]) String() string {
	return box.Min.String() + " " + box.Max.String()
}

// Slice returns the elements of the vector as slice.
func (box *Box[T]) Slice() []T {
	return box.Array()[:]
}

// Array returns a pointer to an array containing the Min and Max coordinates of the Box, providing a fixed-size sequential representation of the box's boundaries.
func (box *Box[T]) Array() *[6]T {
	return &[...]T{
		box.Min[0], box.Min[1], box.Min[2],
		box.Max[0], box.Max[1], box.Max[2],
	}
}

// ContainsPoint returns if a point is contained within the box.
func (box *Box[T]) ContainsPoint(p *Vec[T]) bool {
	return p[0] >= box.Min[0] && p[0] <= box.Max[0] &&
		p[1] >= box.Min[1] && p[1] <= box.Max[1] &&
		p[2] >= box.Min[2] && p[2] <= box.Max[2]
}

// Contains checks whether the box `t` is completely inside the current box. It compares the minimum and maximum coordinates of both boxes and returns true if `t` is contained within.
func (box *Box[T]) Contains(t *Box[T]) bool {
	return t.Min[0] >= box.Min[0] && t.Max[0] <= box.Max[0] &&
		t.Min[1] >= box.Min[1] && t.Max[1] <= box.Max[1] &&
		t.Min[2] >= box.Min[2] && t.Max[2] <= box.Max[2]
}

// Center returns the center point of the Box by averaging the Min and Max coordinates and scaling by 0.5.
func (box *Box[T]) Center() Vec[T] {
	c := Add(&box.Min, &box.Max)
	c.Scale(0.5)
	return c
}

// Diagonal calculates and returns the diagonal vector of the box from Min to Max.
func (box *Box[T]) Diagonal() Vec[T] {
	return Sub(&box.Max, &box.Min)
}

// Intersects returns true if this and the given box intersect.
// For an explanation of the algorithm, see
// http://rbrundritt.wordpress.com/2009/10/03/determining-if-two-bounding-boxes-overlap/
func (box *Box[T]) Intersects(other *Box[T]) bool {
	d1 := box.Diagonal()
	d2 := other.Diagonal()
	sizes := Add(&d1, &d2)
	c1 := box.Center()
	c2 := other.Center()
	distCenters2 := Sub(&c1, &c2)
	distCenters2.Scale(2)
	distCenters2.Abs()
	return distCenters2[0] <= sizes[0] && distCenters2[1] <= sizes[1] && distCenters2[2] <= sizes[2]
}

// Join enlarges this box to contain also the given box.
func (box *Box[T]) Join(other *Box[T]) {
	box.Min = Min(&box.Min, &other.Min)
	box.Max = Max(&box.Max, &other.Max)
}

// Extend grows the Box to include the given point specified by the Vec other.
func (box *Box[T]) Extend(other *Vec[T]) {
	box.Min = Min(&box.Min, other)
	box.Max = Max(&box.Max, other)
}

// Joined returns the minimal box containing both a and b.
func Joined[T float64 | float32](a, b *Box[T]) Box[T] {
	var joined Box[T]
	joined.Min = Min(&a.Min, &b.Min)
	joined.Max = Max(&a.Max, &b.Max)
	return joined
}
