package vec3

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
