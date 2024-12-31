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
