package vec3

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const radian45Degree = math.Pi / 4.0

type testData[T float64 | float32] struct {
	vec         Vec[T]
	expected    Vec[T]
	description string
}

func TestAdd(t *testing.T) {
	v1 := Vec[float64]{1, 2, 3}
	v2 := Vec[float64]{9, 8, 7}
	v3 := v1.Add(&v2)

	assert.Equal(t, Vec[float64]{10, 10, 10}, v1)
	assert.Equal(t, Vec[float64]{9, 8, 7}, v2)
	assert.Equal(t, Vec[float64]{10, 10, 10}, *v3)
}

func TestAdded(t *testing.T) {
	v1 := Vec[float64]{1, 2, 3}
	v2 := Vec[float64]{9, 8, 7}
	v3 := v1.Added(&v2)

	assert.Equal(t, Vec[float64]{1, 2, 3}, v1)
	assert.Equal(t, Vec[float64]{9, 8, 7}, v2)
	assert.Equal(t, Vec[float64]{10, 10, 10}, v3)
}

func TestSub(t *testing.T) {
	v1 := Vec[float64]{1, 2, 3}
	v2 := Vec[float64]{9, 8, 7}
	v3 := v1.Sub(&v2)

	assert.Equal(t, Vec[float64]{-8, -6, -4}, v1)
	assert.Equal(t, Vec[float64]{9, 8, 7}, v2)
	assert.Equal(t, Vec[float64]{-8, -6, -4}, *v3)
}

func TestSubed(t *testing.T) {
	v1 := Vec[float64]{1, 2, 3}
	v2 := Vec[float64]{9, 8, 7}
	v3 := v1.Subed(&v2)

	assert.Equal(t, Vec[float64]{1, 2, 3}, v1)
	assert.Equal(t, Vec[float64]{9, 8, 7}, v2)
	assert.Equal(t, Vec[float64]{-8, -6, -4}, v3)
}

func TestMul(t *testing.T) {
	v1 := Vec[float64]{1, 2, 3}
	v2 := Vec[float64]{9, 8, 7}
	v3 := v1.Mul(&v2)

	assert.Equal(t, Vec[float64]{9, 16, 21}, v1)
	assert.Equal(t, Vec[float64]{9, 8, 7}, v2)
	assert.Equal(t, Vec[float64]{9, 16, 21}, *v3)
}

func TestMuled(t *testing.T) {
	v1 := Vec[float64]{1, 2, 3}
	v2 := Vec[float64]{9, 8, 7}
	v3 := v1.Muled(&v2)

	assert.Equal(t, Vec[float64]{1, 2, 3}, v1)
	assert.Equal(t, Vec[float64]{9, 8, 7}, v2)
	assert.Equal(t, Vec[float64]{9, 16, 21}, v3)
}

func TestScale(t *testing.T) {
	vec := Vec[float64]{1, 2, 3}
	vec.Scale(2)

	expected := Vec[float64]{2, 4, 6}
	assert.Equal(t, expected, vec)
}

func TestScaled(t *testing.T) {
	vec := Vec[float64]{1, 2, 3}
	scaled := vec.Scaled(2)

	expected := Vec[float64]{2, 4, 6}
	assert.Equal(t, expected, scaled)
}

func TestString(t *testing.T) {
	vec := Vec[float64]{1, 2, 3}
	if vec.String() != "1 2 3" {
		t.Errorf("String representation incorrect, got: %s", vec.String())
	}
}

func TestRowsColsSize(t *testing.T) {
	vec := Vec[float64]{}
	if vec.Rows() != 3 {
		t.Error("Rows method returned incorrect value")
	}
	if vec.Cols() != 1 {
		t.Error("Cols method returned incorrect value")
	}
	if vec.Size() != 3 {
		t.Error("Size method returned incorrect value")
	}
}

func TestGet(t *testing.T) {
	vec := Vec[float64]{1, 2, 3}
	if vec.Get(0, 0) != 1 {
		t.Error("Get method returned incorrect value for row 0")
	}
	if vec.Get(0, 1) != 2 {
		t.Error("Get method returned incorrect value for row 1")
	}
	if vec.Get(0, 2) != 3 {
		t.Error("Get method returned incorrect value for row 2")
	}
}

func TestInvert(t *testing.T) {
	vec := Vec[float64]{1, 2, 3}
	vec.Invert()

	expected := Vec[float64]{-1, -2, -3}
	assert.Equal(t, expected, vec)
}

func TestInverted(t *testing.T) {
	vec := Vec[float64]{1, 2, 3}
	inverted := vec.Inverted()

	expected := Vec[float64]{-1, -2, -3}
	assert.Equal(t, expected, inverted)
}

func TestAbs(t *testing.T) {
	vec := Vec[float64]{-1, -2, -3}
	vec.Abs()

	expected := Vec[float64]{1, 2, 3}
	assert.Equal(t, expected, vec)
}

func TestAbsed(t *testing.T) {
	vec := Vec[float64]{-1, -2, -3}
	absed := vec.Absed()

	expected := Vec[float64]{1, 2, 3}
	assert.Equal(t, expected, absed)
}

func TestNormal(t *testing.T) {
	vec := Vec[float64]{1, 0, 0}
	normal := vec.Normal()

	expected := Vec[float64]{0, 0, 1}
	assert.Equal(t, expected, normal)
}

func TestIsZero(t *testing.T) {
	var tests = []struct {
		vec         Vec[float64]
		expected    bool
		description string
	}{
		{Vec[float64]{}, true, "Zero vector"},
		{Vec[float64]{1, 0, 0}, false, "Non-zero X component"},
		{Vec[float64]{0, 1, 0}, false, "Non-zero Y component"},
		{Vec[float64]{0, 0, 1}, false, "Non-zero Z component"},
	}

	for _, test := range tests {
		if test.vec.IsZero() != test.expected {
			t.Errorf("%s: IsZero returned incorrect value, expected %v, got %v", test.description, test.expected, !test.expected)
		}
	}
}

func TestAngle(t *testing.T) {
	assert.InDelta(t, 0*radian45Degree, Angle(&Vec[float64]{1, 0, 0}, &Vec[float64]{1, 0, 0}), 1e-8, "0/360 degree angle, equal/parallell vectors")
	assert.InDelta(t, 1*radian45Degree, Angle(&Vec[float64]{1, 0, 0}, &Vec[float64]{1, 1, 0}), 1e-8, "45 degree angle")
	assert.InDelta(t, 2*radian45Degree, Angle(&Vec[float64]{1, 0, 0}, &Vec[float64]{0, 1, 0}), 1e-8, "90 degree angle, orthogonal vectors")
	assert.InDelta(t, 3*radian45Degree, Angle(&Vec[float64]{1, 0, 0}, &Vec[float64]{-1, 1, 0}), 1e-8, "135 degree angle")
	assert.InDelta(t, 4*radian45Degree, Angle(&Vec[float64]{1, 0, 0}, &Vec[float64]{-1, 0, 0}), 1e-8, "180 degree angle, inverted/anti parallell vectors")
	assert.InDelta(t, (8-5)*radian45Degree, Angle(&Vec[float64]{1, 0, 0}, &Vec[float64]{-1, -1, 0}), 1e-8, "225 degree angle")
	assert.InDelta(t, (8-6)*radian45Degree, Angle(&Vec[float64]{1, 0, 0}, &Vec[float64]{0, -1, 0}), 1e-8, "270 degree angle, orthogonal vectors")
	assert.InDelta(t, (8-7)*radian45Degree, Angle(&Vec[float64]{1, 0, 0}, &Vec[float64]{1, -1, 0}), 1e-8, "315 degree angle")
}

func TestCosine(t *testing.T) {
	assert.InDelta(t, math.Cos(0*radian45Degree), Cosine(&Vec[float64]{1, 0, 0}, &Vec[float64]{1, 0, 0}), 1e-8, "0/360 degree angle, equal/parallell vectors")
	assert.InDelta(t, math.Cos(1*radian45Degree), Cosine(&Vec[float64]{1, 0, 0}, &Vec[float64]{1, 1, 0}), 1e-8, "45 degree angle")
	assert.InDelta(t, math.Cos(2*radian45Degree), Cosine(&Vec[float64]{1, 0, 0}, &Vec[float64]{0, 1, 0}), 1e-8, "90 degree angle, orthogonal vectors")
	assert.InDelta(t, math.Cos(3*radian45Degree), Cosine(&Vec[float64]{1, 0, 0}, &Vec[float64]{-1, 1, 0}), 1e-8, "135 degree angle")
	assert.InDelta(t, math.Cos(4*radian45Degree), Cosine(&Vec[float64]{1, 0, 0}, &Vec[float64]{-1, 0, 0}), 1e-8, "180 degree angle, inverted/anti parallell vectors")
	assert.InDelta(t, math.Cos(5*radian45Degree), Cosine(&Vec[float64]{1, 0, 0}, &Vec[float64]{-1, -1, 0}), 1e-8, "225 degree angle")
	assert.InDelta(t, math.Cos(6*radian45Degree), Cosine(&Vec[float64]{1, 0, 0}, &Vec[float64]{0, -1, 0}), 1e-8, "270 degree angle, orthogonal vectors")
	assert.InDelta(t, math.Cos(7*radian45Degree), Cosine(&Vec[float64]{1, 0, 0}, &Vec[float64]{1, -1, 0}), 1e-8, "315 degree angle")
}

func TestSinus(t *testing.T) {
	assert.InDelta(t, math.Sin(0*radian45Degree), Sinus(&Vec[float64]{1, 0, 0}, &Vec[float64]{1, 0, 0}), 1e-8, "0/360 degree angle, equal/parallell vectors")
	assert.InDelta(t, math.Sin(1*radian45Degree), Sinus(&Vec[float64]{1, 0, 0}, &Vec[float64]{1, 1, 0}), 1e-8, "45 degree angle")
	assert.InDelta(t, math.Sin(2*radian45Degree), Sinus(&Vec[float64]{1, 0, 0}, &Vec[float64]{0, 1, 0}), 1e-8, "90 degree angle, orthogonal vectors")
	assert.InDelta(t, math.Sin(3*radian45Degree), Sinus(&Vec[float64]{1, 0, 0}, &Vec[float64]{-1, 1, 0}), 1e-8, "135 degree angle")
	assert.InDelta(t, math.Sin(4*radian45Degree), Sinus(&Vec[float64]{1, 0, 0}, &Vec[float64]{-1, 0, 0}), 1e-8, "180 degree angle, inverted/anti parallell vectors")
	assert.InDelta(t, math.Abs(math.Sin(5*radian45Degree)), Sinus(&Vec[float64]{1, 0, 0}, &Vec[float64]{-1, -1, 0}), 1e-8, "225 degree angle")
	assert.InDelta(t, math.Abs(math.Sin(6*radian45Degree)), Sinus(&Vec[float64]{1, 0, 0}, &Vec[float64]{0, -1, 0}), 1e-8, "270 degree angle, orthogonal vectors")
	assert.InDelta(t, math.Abs(math.Sin(7*radian45Degree)), Sinus(&Vec[float64]{1, 0, 0}, &Vec[float64]{1, -1, 0}), 1e-8, "315 degree angle")
}

func TestProjectOnVector(t *testing.T) {
	vec := Vec[float64]{1, 2, 3}
	vector := Vec[float64]{1, 0, 0}
	vec.ProjectOnVector(&vector)

	expected := Vec[float64]{1, 0, 0}
	assert.Equal(t, expected, vec)
}

func TestProjectOnPlane(t *testing.T) {
	vec := Vec[float64]{1, 2, 3}
	planeNormal := Vec[float64]{1, 0, 0}
	vec.ProjectOnPlane(&planeNormal)

	expected := Vec[float64]{0, 2, 3}
	assert.Equal(t, expected, vec)
}

func TestNormalize(t *testing.T) {
	var tests = []testData[float64]{
		{Vec[float64]{1, 0, 0}, Vec[float64]{1, 0, 0}, "Unit vector along X"},
		{Vec[float64]{1, 2, 2}, Vec[float64]{0.37139068, 0.74278136, 0.74278136}, "Non-unit vector"},
		// ... Additional test cases ...
	}

	for _, td := range tests {
		td.vec.Normalize()
		if !td.vec.ApproxEqual(td.expected) {
			t.Errorf("%s: Normalize failed, expected %v, got %v", td.description, td.expected, td.vec)
		}
	}
}

func TestSquareDistance(t *testing.T) {
	vecA := Vec[float64]{1, 2, 3}
	vecB := Vec[float64]{4, 5, 6}
	squareDistance := SquareDistance(&vecA, &vecB)

	expected := float64(27)
	assert.Equal(t, expected, squareDistance)
}

func TestDistance(t *testing.T) {
	vecA := Vec[float64]{1, 2, 3}
	vecB := Vec[float64]{4, 5, 6}
	distance := Distance(&vecA, &vecB)

	expected := math.Sqrt(27)
	assert.Equal(t, expected, distance)
}

func TestCross(t *testing.T) {
	vecA := Vec[float64]{1, 0, 0}
	vecB := Vec[float64]{0, 1, 0}
	cross := Cross(&vecA, &vecB)

	expected := Vec[float64]{0, 0, 1}
	assert.Equal(t, expected, cross)
}

func TestMin(t *testing.T) {
	vecA := Vec[float64]{1, 2, 3}
	vecB := Vec[float64]{0, 3, 2}
	min := Min(&vecA, &vecB)

	expected := Vec[float64]{0, 2, 2}
	assert.Equal(t, expected, min)
}

func TestMax(t *testing.T) {
	vecA := Vec[float64]{1, 2, 3}
	vecB := Vec[float64]{0, 3, 2}
	max := Max(&vecA, &vecB)

	expected := Vec[float64]{1, 3, 3}
	assert.Equal(t, expected, max)
}

func TestSetMin(t *testing.T) {
	vec := Vec[float64]{1, 2, 3}
	vec.SetMin(Vec[float64]{0, 3, 2})

	expected := Vec[float64]{0, 2, 2}
	assert.Equal(t, expected, vec)
}

func TestSetMax(t *testing.T) {
	vec := Vec[float64]{1, 2, 3}
	vec.SetMax(Vec[float64]{0, 3, 2})

	expected := Vec[float64]{1, 3, 3}
	assert.Equal(t, expected, vec)
}

func TestRotated(t *testing.T) {
	vec := Vec[float64]{1, 0, 0}
	axis := Vec[float64]{0, 0, 1}
	rotated := Rotated(&vec, &axis, math.Pi/2)

	expected := Vec[float64]{0, 1, 0}
	assert.InDelta(t, expected[0], rotated[0], 1e-9)
	assert.InDelta(t, expected[1], rotated[1], 1e-9)
	assert.InDelta(t, expected[2], rotated[2], 1e-9)
}

func TestPointSegmentDistance(t *testing.T) {
	p1 := Vec[float64]{1, 1, 0}
	x1 := Vec[float64]{0, 0, 0}
	x2 := Vec[float64]{2, 0, 0}
	distance := PointSegmentDistance(&p1, &x1, &x2)

	expected := float64(1)
	assert.Equal(t, expected, distance)
}

func TestPointSegmentVerticalPoint(t *testing.T) {
	p1 := Vec[float64]{1, 1, 0}
	x1 := Vec[float64]{0, 0, 0}
	x2 := Vec[float64]{2, 0, 0}
	verticalPoint := PointSegmentVerticalPoint(&p1, &x1, &x2)

	expected := Vec[float64]{1, 0, 0}
	assert.Equal(t, expected, *verticalPoint)
}

func TestClamp(t *testing.T) {
	// Test setup and assertions would follow the pattern above
}

func TestInterpolate(t *testing.T) {
	// Test setup and assertions would follow the pattern above
}

// ... Complete all test cases ...

func (vec *Vec[T]) ApproxEqual(other Vec[T]) bool {
	const epsilon = 1e-9
	return math.Abs(float64(vec[0]-other[0])) <= epsilon &&
		math.Abs(float64(vec[1]-other[1])) <= epsilon &&
		math.Abs(float64(vec[2]-other[2])) <= epsilon
}
