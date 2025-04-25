package vec4

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"pinkey.ltd/xr/go3d/vec2"
	"pinkey.ltd/xr/go3d/vec3"
)

func TestVec3(t *testing.T) {
	vec := Vec[float64]{1, 2, 3, 4}
	v3 := vec.Vec3()

	expected := vec3.Vec[float64]{1, 2, 3}
	assert.Equal(t, expected, v3)
}

func TestAssignVec3(t *testing.T) {
	vec := Vec[float64]{0, 0, 0, 0}
	v3 := &vec3.Vec[float64]{1, 2, 3}
	vec.AssignVec3(v3)

	expected := Vec[float64]{1, 2, 3, 1}
	assert.Equal(t, expected, vec)
}

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

func TestVecSub(t *testing.T) {
	vecA := Vec[float64]{1, 2, 3, 1}
	vecB := Vec[float64]{4, 5, 6, 1}
	vecA.Sub(&vecB)

	expected := Vec[float64]{-3, -3, -3, 1}
	assert.Equal(t, expected, vecA)
}

func TestVecSubFunction(t *testing.T) {
	vecA := Vec[float64]{1, 2, 3, 1}
	vecB := Vec[float64]{4, 5, 6, 1}
	result := Sub(&vecA, &vecB)

	expected := Vec[float64]{-3, -3, -3, 1}
	assert.Equal(t, expected, result)
}

func TestCross(t *testing.T) {
	vecA := Vec[float64]{1, 0, 0, 1}
	vecB := Vec[float64]{0, 1, 0, 1}
	cross := Cross(&vecA, &vecB)

	expected := Vec[float64]{0, 0, 1, 1}
	assert.Equal(t, expected, cross)
}

func TestVecClamp(t *testing.T) {
	vec := Vec[float64]{-1, 2, 3, 1}
	minVal := Vec[float64]{0, 0, 0, 0}
	maxVal := Vec[float64]{5, 5, 5, 5}

	assert.Equal(t, Vec[float64]{0, 2, 3, 1}, *vec.Clamp(&minVal, &maxVal))
}

func TestAngle(t *testing.T) {
	vecA := Vec[float64]{1, 0, 0, 1}
	vecB := Vec[float64]{0, 1, 0, 1}
	angle := Angle(&vecA, &vecB)

	expected := math.Pi / 2
	assert.InDelta(t, expected, float64(angle), 1e-9)
}

func TestFrom(t *testing.T) {
	vec2D := &vec2.Vec[float64]{1, 2}
	vec3D := &vec3.Vec[float64]{1, 2, 3}
	vec4D := &Vec[float64]{1, 2, 3, 4}

	assert.Equal(t, Vec[float64]{1, 2, 0, 1}, From[float64](vec2D))
	assert.Equal(t, Vec[float64]{1, 2, 3, 1}, From[float64](vec3D))
	assert.Equal(t, Vec[float64]{1, 2, 3, 4}, From[float64](vec4D))
}

func TestVecScale(t *testing.T) {
	vec := Vec[float64]{1, 2, 3, 4}
	vec.Scale(2)

	expected := Vec[float64]{2, 4, 6, 4}
	assert.Equal(t, expected, vec)
}

func TestVecScaled(t *testing.T) {
	vec := Vec[float64]{1, 2, 3, 4}
	scaled := vec.Scaled(2)

	expected := Vec[float64]{2, 4, 6, 4}
	assert.Equal(t, expected, scaled)
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

func TestDot(t *testing.T) {
	vecA := Vec[float64]{1, 2, 3, 1}
	vecB := Vec[float64]{4, 5, 6, 1}
	dot := Dot(&vecA, &vecB)

	expected := float64(1*4 + 2*5 + 3*6)
	assert.Equal(t, expected, dot)
}

func TestVecDot4(t *testing.T) {
	vecA := Vec[float64]{1, 2, 3, 4}
	vecB := Vec[float64]{4, 5, 6, 7}
	dot := Dot4(&vecA, &vecB)

	expected := float64(1*4 + 2*5 + 3*6 + 4*7)
	assert.Equal(t, expected, dot)
}

func TestInvert(t *testing.T) {
	vec := Vec[float64]{1, 2, 3, 4}
	vec.Invert()

	expected := Vec[float64]{-1, -2, -3, 4}
	assert.Equal(t, expected, vec)
}

func TestInverted(t *testing.T) {
	vec := Vec[float64]{1, 2, 3, 4}
	inverted := vec.Inverted()

	expected := Vec[float64]{-1, -2, -3, 4}
	assert.Equal(t, expected, inverted)
}

func TestNormalize(t *testing.T) {
	vec := Vec[float64]{1, 2, 3, 4}
	vec.Normalize()

	length := vec.Length()
	assert.InDelta(t, 1.0, length, 1e-9)
}

func TestNormalized(t *testing.T) {
	vec := Vec[float64]{1, 2, 3, 4}
	normalized := vec.Normalized()

	length := normalized.Length()
	assert.InDelta(t, 1.0, length, 1e-9)
}

func TestDivideByW(t *testing.T) {
	vec := Vec[float64]{2, 4, 6, 2}
	vec.DivideByW()

	expected := Vec[float64]{1, 2, 3, 1}
	assert.Equal(t, expected, vec)
}

func TestDividedByW(t *testing.T) {
	vec := Vec[float64]{2, 4, 6, 2}
	divided := vec.DividedByW()

	expected := Vec[float64]{1, 2, 3, 1}
	assert.Equal(t, expected, divided)
}

func TestVec3DividedByW(t *testing.T) {
	vec := Vec[float64]{2, 4, 6, 2}
	v3 := vec.Vec3DividedByW()

	expected := vec3.Vec[float64]{1, 2, 3}
	assert.Equal(t, expected, v3)
}

func TestInterpolate(t *testing.T) {
	vecA := Vec[float64]{1, 2, 3, 1}
	vecB := Vec[float64]{4, 5, 6, 1}
	interpolated := Interpolate(&vecA, &vecB, 0.5)

	expected := Vec[float64]{2.5, 3.5, 4.5, 1}
	assert.Equal(t, expected, interpolated)
}

func TestClamp01(t *testing.T) {
	vec := Vec[float64]{-1, 2, 0.5, 1}
	vec.Clamp01()

	expected := Vec[float64]{0, 1, 0.5, 1}
	assert.Equal(t, expected, vec)
}

func TestClamped01(t *testing.T) {
	vec := Vec[float64]{-1, 2, 0.5, 1}
	clamped := vec.Clamped01()

	expected := Vec[float64]{0, 1, 0.5, 1}
	assert.Equal(t, expected, clamped)
}
