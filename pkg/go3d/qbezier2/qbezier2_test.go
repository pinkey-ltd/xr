package qbezier2

import (
	"math"
	"pinkey.ltd/xr/pkg/go3d/vec2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBezString(t *testing.T) {
	bez := Bez[float64]{
		P0: vec2.Vec[float64]{0, 0},
		P1: vec2.Vec[float64]{1, 1},
		P2: vec2.Vec[float64]{2, 2},
	}
	expected := "(0 0) (1 1) (2 2)"
	assert.Equal(t, expected, bez.String())
}

func TestParse(t *testing.T) {
	for _, tc := range []struct {
		input    string
		expected Bez[float64]
		err      bool
	}{
		{"0 0 1 1 2 2", Bez[float64]{P0: vec2.Vec[float64]{0, 0}, P1: vec2.Vec[float64]{1, 1}, P2: vec2.Vec[float64]{2, 2}}, false},
		{"0 0 1 1", Bez[float64]{}, true},
	} {
		result, err := Parse[float64](tc.input)
		if tc.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		}
	}
}

func TestPoint(t *testing.T) {
	p0 := vec2.Vec[float64]{0, 0}
	p1 := vec2.Vec[float64]{1, 1}
	p2 := vec2.Vec[float64]{2, 2}
	testCases := []struct {
		t        float64
		expected vec2.Vec[float64]
	}{
		{0, vec2.Vec[float64]{0, 0}},
		{1, vec2.Vec[float64]{2, 2}},
		{0.5, vec2.Vec[float64]{1, 1}},
	}

	for _, tc := range testCases {
		result := Point[float64](&p0, &p1, &p2, tc.t)
		assert.InDeltaSlice(t, tc.expected.Slice(), result.Slice(), 1e-9)
	}
}

func TestBezier2Point(t *testing.T) {
	bez := Bez[float64]{
		vec2.Vec[float64]{0, 0},
		vec2.Vec[float64]{1, 1},
		vec2.Vec[float64]{2, 0},
	}

	assert.Equal(t, vec2.Vec[float64]{0, 0}, bez.Point(0))
	assert.Equal(t, vec2.Vec[float64]{2, 0}, bez.Point(1))
	assert.Equal(t, vec2.Vec[float64]{1, 0.5}, bez.Point(0.5))
	assert.Equal(t, vec2.Vec[float64]{0.5, 0.375}, bez.Point(0.25))
	assert.Equal(t, vec2.Vec[float64]{1.5, 0.375}, bez.Point(0.75))
}

func TestBezier2Tangent(t *testing.T) {
	bez := Bez[float64]{
		vec2.Vec[float64]{0, 0},
		vec2.Vec[float64]{1, 1},
		vec2.Vec[float64]{2, 0},
	}

	assert.Equal(t, vec2.Vec[float64]{2, 2}, bez.Tangent(0))
	assert.Equal(t, vec2.Vec[float64]{2, -2}, bez.Tangent(1))
	assert.Equal(t, vec2.Vec[float64]{2, 0}, bez.Tangent(0.5))
	assert.Equal(t, vec2.Vec[float64]{2, 1}, bez.Tangent(0.25))
	assert.Equal(t, vec2.Vec[float64]{2, -1}, bez.Tangent(0.75))
}

func TestTangent(t *testing.T) {
	p0 := vec2.Vec[float64]{0, 0}
	p1 := vec2.Vec[float64]{1, 1}
	p2 := vec2.Vec[float64]{2, 2}
	testCases := []struct {
		t        float64
		expected vec2.Vec[float64]
	}{
		{0, vec2.Vec[float64]{1, 1}},
		{1, vec2.Vec[float64]{1, 1}},
		{0.5, vec2.Vec[float64]{1, 1}},
	}

	for _, tc := range testCases {
		result := Tangent[float64](&p0, &p1, &p2, tc.t)
		assert.InDeltaSlice(t, tc.expected.Slice(), result.Slice(), 1e-9)
	}
}

func TestLength(t *testing.T) {
	p0 := vec2.Vec[float64]{0, 0}
	p1 := vec2.Vec[float64]{1, 1}
	p2 := vec2.Vec[float64]{2, 2}
	testCases := []struct {
		t        float64
		expected float64
	}{
		{0, 0},
		{1, math.Sqrt(8)},
		{0.5, math.Sqrt(2) / 2},
	}

	for _, tc := range testCases {
		result := Length[float64](&p0, &p1, &p2, tc.t)
		assert.InDelta(t, tc.expected, result, 1e-9)
	}
}
