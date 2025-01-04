package vec2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase[T float64 | float32] struct {
	name     string
	rect     Rect[T]
	expected T
}

func TestNewRect(t *testing.T) {
	vecA := Vec[float64]{1, 2}
	vecB := Vec[float64]{3, 4}
	rect := NewRect(&vecA, &vecB)

	expected := Rect[float64]{
		Min: Vec[float64]{1, 2},
		Max: Vec[float64]{3, 4},
	}
	assert.Equal(t, expected, rect)
}

func TestParseRect(t *testing.T) {
	s := "1 2 3 4"
	rect, err := ParseRect[float64](s)

	assert.NoError(t, err)
	expected := Rect[float64]{
		Min: Vec[float64]{1, 2},
		Max: Vec[float64]{3, 4},
	}
	assert.Equal(t, expected, rect)
}

func TestRectSize(t *testing.T) {
	tests := []testCase[float64]{
		{"Square", Rect[float64]{Vec[float64]{0, 0}, Vec[float64]{5, 5}}, 5},
		{"Rectangle", Rect[float64]{Vec[float64]{0, 0}, Vec[float64]{10, 5}}, 10},
		{"ZeroSize", Rect[float64]{Vec[float64]{0, 0}, Vec[float64]{0, 0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.rect.Size()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRectSlice(t *testing.T) {
	rect := Rect[float64]{
		Min: Vec[float64]{1, 2},
		Max: Vec[float64]{3, 4},
	}
	expectedSlice := []float64{1, 2, 3, 4}
	assert.Equal(t, expectedSlice, rect.Slice())
}

func TestRectArray(t *testing.T) {
	rect := Rect[float64]{
		Min: Vec[float64]{1, 2},
		Max: Vec[float64]{3, 4},
	}
	expectedArray := [4]float64{1, 2, 3, 4}
	assert.Equal(t, &expectedArray, rect.Array())
}

func TestRectContainsPoint(t *testing.T) {
	rect := Rect[float64]{
		Min: Vec[float64]{0, 0},
		Max: Vec[float64]{2, 2},
	}
	pointInside := Vec[float64]{1, 1}
	pointOutside := Vec[float64]{3, 3}

	assert.True(t, rect.ContainsPoint(&pointInside), "Point inside should be contained")
	assert.False(t, rect.ContainsPoint(&pointOutside), "Point outside should not be contained")
}

func TestRectContains(t *testing.T) {
	rectA := Rect[float64]{
		Min: Vec[float64]{0, 0},
		Max: Vec[float64]{2, 2},
	}
	rectB := Rect[float64]{
		Min: Vec[float64]{1, 1},
		Max: Vec[float64]{1, 1},
	}
	rectC := Rect[float64]{
		Min: Vec[float64]{3, 3},
		Max: Vec[float64]{4, 4},
	}

	assert.True(t, rectA.Contains(&rectB), "Rect A should contain Rect B")
	assert.False(t, rectA.Contains(&rectC), "Rect A should not contain Rect C")
}

func TestRectArea(t *testing.T) {
	rect := Rect[float64]{
		Min: Vec[float64]{0, 0},
		Max: Vec[float64]{2, 3},
	}
	expectedArea := float64(6)
	assert.Equal(t, expectedArea, rect.Area())
}

func TestRectIntersects(t *testing.T) {
	rectA := Rect[float64]{
		Min: Vec[float64]{0, 0},
		Max: Vec[float64]{2, 2},
	}
	rectB := Rect[float64]{
		Min: Vec[float64]{1, 1},
		Max: Vec[float64]{3, 3},
	}
	rectC := Rect[float64]{
		Min: Vec[float64]{3, 3},
		Max: Vec[float64]{4, 4},
	}

	assert.True(t, rectA.Intersects(&rectB), "Rect A and B should intersect")
	assert.False(t, rectA.Intersects(&rectC), "Rect A and C should not intersect")
}

func TestRectJoin(t *testing.T) {
	rectA := Rect[float64]{
		Min: Vec[float64]{0, 0},
		Max: Vec[float64]{2, 2},
	}
	rectB := Rect[float64]{
		Min: Vec[float64]{1, 1},
		Max: Vec[float64]{3, 3},
	}
	rectA.Join(&rectB)

	expected := Rect[float64]{
		Min: Vec[float64]{0, 0},
		Max: Vec[float64]{3, 3},
	}
	assert.Equal(t, expected, rectA)
}

func TestRectJoined(t *testing.T) {
	rectA := Rect[float64]{
		Min: Vec[float64]{0, 0},
		Max: Vec[float64]{2, 2},
	}
	rectB := Rect[float64]{
		Min: Vec[float64]{1, 1},
		Max: Vec[float64]{3, 3},
	}
	joined := Joined(&rectA, &rectB)

	expected := Rect[float64]{
		Min: Vec[float64]{0, 0},
		Max: Vec[float64]{3, 3},
	}
	assert.Equal(t, expected, joined)
}

func TestRectExtend(t *testing.T) {
	rect := Rect[float64]{
		Min: Vec[float64]{0, 0},
		Max: Vec[float64]{2, 2},
	}
	point := Vec[float64]{3, 3}
	rect.Extend(&point)

	expected := Rect[float64]{
		Min: Vec[float64]{0, 0},
		Max: Vec[float64]{3, 3},
	}
	assert.Equal(t, expected, rect)
}

func TestRectWidth(t *testing.T) {
	tests := []testCase[float64]{
		{"PositiveWidth", Rect[float64]{Vec[float64]{0, 0}, Vec[float64]{5, 10}}, 5},
		{"NegativeWidth", Rect[float64]{Vec[float64]{5, 0}, Vec[float64]{0, 10}}, 5},
		{"ZeroWidth", Rect[float64]{Vec[float64]{0, 0}, Vec[float64]{0, 10}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.rect.Width()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRectHeight(t *testing.T) {
	tests := []testCase[float64]{
		{"PositiveHeight", Rect[float64]{Vec[float64]{0, 0}, Vec[float64]{10, 5}}, 5},
		{"NegativeHeight", Rect[float64]{Vec[float64]{0, 5}, Vec[float64]{10, 0}}, 5},
		{"ZeroHeight", Rect[float64]{Vec[float64]{0, 0}, Vec[float64]{10, 0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.rect.Height()
			assert.Equal(t, tt.expected, result)
		})
	}
}
