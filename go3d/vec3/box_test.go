package vec3

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoxString(t *testing.T) {
	box := Box[float64]{
		Min: Vec[float64]{1, 2, 3},
		Max: Vec[float64]{4, 5, 6},
	}
	expected := "1 2 3 4 5 6"
	if box.String() != expected {
		t.Errorf("Box.String() = %v, want %v", box.String(), expected)
	}
}

func TestBoxSlice(t *testing.T) {
	box := Box[float64]{
		Min: Vec[float64]{1, 2, 3},
		Max: Vec[float64]{4, 5, 6},
	}
	expected := []float64{1, 2, 3, 4, 5, 6}
	if !equalSlices(box.Slice(), expected) {
		t.Errorf("Box.Slice() = %v, want %v", box.Slice(), expected)
	}
}

func TestBoxArray(t *testing.T) {
	box := Box[float64]{
		Min: Vec[float64]{1, 2, 3},
		Max: Vec[float64]{4, 5, 6},
	}
	expected := [6]float64{1, 2, 3, 4, 5, 6}
	if !equalArrays(box.Array(), &expected) {
		t.Errorf("Box.Array() = %v, want %v", box.Array(), expected)
	}
}

func TestBoxContainsPoint(t *testing.T) {
	box := Box[float64]{
		Min: Vec[float64]{0, 0, 0},
		Max: Vec[float64]{1, 1, 1},
	}
	pointInside := Vec[float64]{0.5, 0.5, 0.5}
	pointOutside := Vec[float64]{2, 2, 2}

	if !box.ContainsPoint(&pointInside) {
		t.Error("Point inside should be contained")
	}
	if box.ContainsPoint(&pointOutside) {
		t.Error("Point outside should not be contained")
	}
}

func TestBoxContains(t *testing.T) {
	box := Box[float64]{
		Min: Vec[float64]{0, 0, 0},
		Max: Vec[float64]{2, 2, 2},
	}
	smallerBox := Box[float64]{
		Min: Vec[float64]{1, 1, 1},
		Max: Vec[float64]{1, 1, 1},
	}
	biggerBox := Box[float64]{
		Min: Vec[float64]{-1, -1, -1},
		Max: Vec[float64]{3, 3, 3},
	}
	notContainedBox := Box[float64]{
		Min: Vec[float64]{3, 3, 3},
		Max: Vec[float64]{4, 4, 4},
	}

	if !box.Contains(&smallerBox) {
		t.Error("Smaller box should be contained")
	}
	if !box.Contains(&biggerBox) {
		t.Error("Bigger box should contain smaller box")
	}
	if box.Contains(&notContainedBox) {
		t.Error("Not contained box should not be contained")
	}
}

func TestBoxCenter(t *testing.T) {
	box := Box[float64]{
		Min: Vec[float64]{0, 0, 0},
		Max: Vec[float64]{2, 2, 2},
	}

	assert.Equal(t, Vec[float64]{1, 1, 1}, box.Center())
}

func TestBoxDiagonal(t *testing.T) {
	box := Box[float64]{
		Min: Vec[float64]{0, 0, 0},
		Max: Vec[float64]{2, 2, 2},
	}

	assert.Equal(t, Vec[float64]{2, 2, 2}, box.Diagonal())
}

func TestBoxIntersects(t *testing.T) {
	boxA := Box[float64]{
		Min: Vec[float64]{0, 0, 0},
		Max: Vec[float64]{1, 1, 1},
	}
	boxB := Box[float64]{
		Min: Vec[float64]{1, 1, 1},
		Max: Vec[float64]{2, 2, 2},
	}
	boxC := Box[float64]{
		Min: Vec[float64]{1, 2, 1},
		Max: Vec[float64]{2, 3, 2},
	}

	assert.True(t, boxA.Intersects(&boxB), "Boxes A and B should intersect")
	assert.True(t, !boxA.Intersects(&boxC), "Boxes A and C should not intersect")
}

func TestFromSlice(t *testing.T) {
	slice := []float64{1, 2, 3, 4, 5, 6}
	box := FromSlice(slice)

	expected := Box[float64]{
		Min: Vec[float64]{1, 2, 3},
		Max: Vec[float64]{4, 5, 6},
	}
	assert.Equal(t, expected, *box)
}

func TestParseBox(t *testing.T) {
	s := "1 2 3 4 5 6"
	box, err := ParseBox[float64](s)

	assert.NoError(t, err)
	expected := Box[float64]{
		Min: Vec[float64]{1, 2, 3},
		Max: Vec[float64]{4, 5, 6},
	}
	assert.Equal(t, expected, box)
}

func TestJoin(t *testing.T) {
	boxA := Box[float64]{
		Min: Vec[float64]{0, 0, 0},
		Max: Vec[float64]{1, 1, 1},
	}
	boxB := Box[float64]{
		Min: Vec[float64]{2, 2, 2},
		Max: Vec[float64]{3, 3, 3},
	}
	boxA.Join(&boxB)

	expected := Box[float64]{
		Min: Vec[float64]{0, 0, 0},
		Max: Vec[float64]{3, 3, 3},
	}
	assert.Equal(t, expected, boxA)
}

func TestJoined(t *testing.T) {
	boxA := Box[float64]{
		Min: Vec[float64]{0, 0, 0},
		Max: Vec[float64]{1, 1, 1},
	}
	boxB := Box[float64]{
		Min: Vec[float64]{2, 2, 2},
		Max: Vec[float64]{3, 3, 3},
	}
	joined := Joined(&boxA, &boxB)

	expected := Box[float64]{
		Min: Vec[float64]{0, 0, 0},
		Max: Vec[float64]{3, 3, 3},
	}
	assert.Equal(t, expected, joined)
}

func TestExtend(t *testing.T) {
	box := Box[float64]{
		Min: Vec[float64]{0, 0, 0},
		Max: Vec[float64]{1, 1, 1},
	}
	point := Vec[float64]{2, 2, 2}
	box.Extend(&point)

	expected := Box[float64]{
		Min: Vec[float64]{0, 0, 0},
		Max: Vec[float64]{2, 2, 2},
	}
	assert.Equal(t, expected, box)
}

func TestIntersectsBoundary(t *testing.T) {
	boxA := Box[float64]{
		Min: Vec[float64]{0, 0, 0},
		Max: Vec[float64]{1, 1, 1},
	}
	boxB := Box[float64]{
		Min: Vec[float64]{1, 1, 1},
		Max: Vec[float64]{2, 2, 2},
	}
	boxC := Box[float64]{
		Min: Vec[float64]{1.1, 1.1, 1.1},
		Max: Vec[float64]{2, 2, 2},
	}

	assert.True(t, boxA.Intersects(&boxB), "Boxes A and B should intersect at boundary")
	assert.False(t, boxA.Intersects(&boxC), "Boxes A and C should not intersect")
}

func TestDiagonal(t *testing.T) {
	box := Box[float64]{
		Min: Vec[float64]{0, 0, 0},
		Max: Vec[float64]{2, 2, 2},
	}

	expected := Vec[float64]{2, 2, 2}
	assert.Equal(t, expected, box.Diagonal())
}

func TestCenter(t *testing.T) {
	box := Box[float64]{
		Min: Vec[float64]{0, 0, 0},
		Max: Vec[float64]{2, 2, 2},
	}

	expected := Vec[float64]{1, 1, 1}
	assert.Equal(t, expected, box.Center())
}

func TestContainsBoundary(t *testing.T) {
	boxA := Box[float64]{
		Min: Vec[float64]{0, 0, 0},
		Max: Vec[float64]{2, 2, 2},
	}
	boxB := Box[float64]{
		Min: Vec[float64]{1, 1, 1},
		Max: Vec[float64]{2, 2, 2},
	}
	boxC := Box[float64]{
		Min: Vec[float64]{2, 2, 2},
		Max: Vec[float64]{3, 3, 3},
	}

	assert.True(t, boxA.Contains(&boxB), "Box A should contain Box B")
	assert.False(t, boxA.Contains(&boxC), "Box A should not contain Box C")
}

func TestContainsPointBoundary(t *testing.T) {
	box := Box[float64]{
		Min: Vec[float64]{0, 0, 0},
		Max: Vec[float64]{2, 2, 2},
	}
	pointInside := Vec[float64]{1, 1, 1}
	pointBoundary := Vec[float64]{2, 2, 2}
	pointOutside := Vec[float64]{3, 3, 3}

	assert.True(t, box.ContainsPoint(&pointInside), "Point inside should be contained")
	assert.True(t, box.ContainsPoint(&pointBoundary), "Point on boundary should be contained")
	assert.False(t, box.ContainsPoint(&pointOutside), "Point outside should not be contained")
}

func TestMaxBox(t *testing.T) {
	point := Vec[float64]{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64}
	assert.True(t, MaxBox.ContainsPoint(&point), "MaxBox should contain the maximum point")
}

func TestMinBox(t *testing.T) {
	point := Vec[float64]{-math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64}
	assert.True(t, MinBox.ContainsPoint(&point), "MinBox should contain the minimum point")
}

// Helper functions for testing

func equalSlices[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalArrays[T comparable](a, b *[6]T) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
