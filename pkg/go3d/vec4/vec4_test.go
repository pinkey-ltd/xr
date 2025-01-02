package vec4

import (
	"github.com/stretchr/testify/assert"
	"math"
	"pinkey.ltd/xr/pkg/go3d/vec2"
	"pinkey.ltd/xr/pkg/go3d/vec3"
	"testing"
)

func TestVecString(t *testing.T) {
	for _, tc := range []struct {
		vec    Vec[float64]
		expect string
	}{
		{Vec[float64]{1, 2, 3, 4}, "1 2 3 4"},
		{Vec[float64]{0, -1, 5.5, 0}, "0 -1 5.5 0"},
	} {
		if got := tc.vec.String(); got != tc.expect {
			t.Errorf("Vec.String() = %v, want %v", got, tc.expect)
		}
	}
}

func TestVecRowsColsSize(t *testing.T) {
	vec := Vec[float64]{}
	if got, want := vec.Rows(), 4; got != want {
		t.Errorf("Rows() = %v, want %v", got, want)
	}
	if got, want := vec.Cols(), 1; got != want {
		t.Errorf("Cols() = %v, want %v", got, want)
	}
	if got, want := vec.Size(), 4; got != want {
		t.Errorf("Size() = %v, want %v", got, want)
	}
}

func TestVecIsZero(t *testing.T) {
	v1 := Vec[float64]{}
	v2 := Vec[float64]{1, 0, 0, 0}
	v3 := Vec[float64]{0, 0, 0, 0}

	assert.True(t, v1.IsZero())
	assert.False(t, v2.IsZero())
	assert.True(t, v3.IsZero())
}

func TestVecLength(t *testing.T) {
	vec := Vec[float64]{3, 4, 8, 2}
	expectedLength := math.Sqrt((3./2)*(3./2) + 4/2*4/2 + 8/2*8/2)

	assert.Equal(t, expectedLength, vec.Length())
}

func TestVecAdd(t *testing.T) {
	vecA := Vec[float64]{1, 2, 3, 1}
	vecB := Vec[float64]{4, 5, 6, 1}

	assert.Equal(t, Vec[float64]{5, 7, 9, 1}, *vecA.Add(&vecB))
}

func TestVecClamp(t *testing.T) {
	vec := Vec[float64]{-1, 2, 3, 1}
	minVal := Vec[float64]{0, 0, 0, 0}
	maxVal := Vec[float64]{5, 5, 5, 5}

	assert.Equal(t, Vec[float64]{0, 2, 3, 1}, *vec.Clamp(&minVal, &maxVal))
}

func TestFrom(t *testing.T) {
	vec2D := &vec2.Vec[float64]{1, 2}
	vec3D := &vec3.Vec[float64]{1, 2, 3}
	vec4D := &Vec[float64]{1, 2, 3, 4}

	assert.Equal(t, Vec[float64]{1, 2, 0, 1}, From[float64](vec2D))
	assert.Equal(t, Vec[float64]{1, 2, 3, 1}, From[float64](vec3D))
	assert.Equal(t, Vec[float64]{1, 2, 3, 4}, From[float64](vec4D))
}

func TestShuffle(t *testing.T) {
	vec := Vec[float64]{1, 2, 3, 4}
	vec.Shuffle(WZYX)

	assert.Equal(t, Vec[float64]{4, 3, 2, 1}, vec)
}

func TestShuffled(t *testing.T) {
	vec := Vec[float64]{1, 2, 3, 4}

	assert.Equal(t, Vec[float64]{4, 3, 2, 1}, vec.Shuffled(WZYX))
}
