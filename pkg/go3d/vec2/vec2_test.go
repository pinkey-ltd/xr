package vec2

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const radian45Degree = math.Pi / 4.0

func TestAbs(t *testing.T) {
	vec := Vec[float64]{1.4, -3.5}
	vecRef := vec.Abs()
	assert.Equal(t, Vec[float64]{1.4, 3.5}, vec)
	assert.Equal(t, Vec[float64]{1.4, 3.5}, *vecRef)
}

func TestAbsed(t *testing.T) {
	vec := Vec[float64]{1.4, -3.5}
	vecRef := vec.Absed()
	assert.Equal(t, Vec[float64]{1.4, -3.5}, vec)
	assert.Equal(t, Vec[float64]{1.4, 3.5}, vecRef)
}

func TestNormal(t *testing.T) {
	vec1 := Vec[float64]{1.0, 1.0}
	vecLength1 := math.Sqrt(1*1 + 1*1)

	vec2 := Vec[float64]{4, 6}
	vecLength2 := math.Sqrt(4*4 + 6*6)

	assert.Equal(t, Vec[float64]{1 / vecLength1, -1 / vecLength1}, vec1.Normal())
	assert.Equal(t, Vec[float64]{6.0 / vecLength2, -4 / vecLength2}, vec2.Normal())
}

func TestNormalCCW(t *testing.T) {
	vec1 := Vec[float64]{1.0, 1.0}
	vecLength1 := math.Sqrt(1*1 + 1*1)

	vec2 := Vec[float64]{4, 6}
	vecLength2 := math.Sqrt(4*4 + 6*6)

	assert.Equal(t, Vec[float64]{-1 / vecLength1, 1 / vecLength1}, vec1.NormalCCW())
	assert.Equal(t, Vec[float64]{-6.0 / vecLength2, 4 / vecLength2}, vec2.NormalCCW())
}

func TestString(t *testing.T) {
	vec := Vec[float64]{1, 2}
	assert.Equal(t, "1 2", vec.String())
}

func TestRowsColsSize(t *testing.T) {
	vec := Vec[float64]{}
	assert.Equal(t, 2, vec.Rows())
	assert.Equal(t, 1, vec.Cols())
	assert.Equal(t, 2, vec.Size())
}

func TestGet(t *testing.T) {
	gl := Vec[float64]{3.4, 4.5} // openGL format
	// TODO: Transpose()
	//dx := gl.Transpose()
	assert.Equal(t, 3.4, gl.Get(0, 0))
	assert.Equal(t, 4.5, gl.Get(0, 1))
}

func TestIsZero(t *testing.T) {
	zeroVec := Vec[float64]{}
	nonZeroVec := Vec[float64]{1, 0}
	assert.True(t, zeroVec.IsZero())
	assert.False(t, nonZeroVec.IsZero())
}

func TestLengthLengthSqr(t *testing.T) {
	vec := Vec[float64]{3, 4}
	assert.InDelta(t, 5., vec.Length(), 1e-7)
	assert.Equal(t, 25., vec.LengthSqr())
}

func TestScale(t *testing.T) {
	vec := &Vec[float64]{3, 4}
	scaledVec := vec.Scale(2)
	assert.Equal(t, Vec[float64]{6, 8}, *scaledVec)
}

func TestScaled(t *testing.T) {
	vec := Vec[float64]{3, 4}
	assert.Equal(t, Vec[float64]{6, 8}, vec.Scaled(2))
}

func TestInvertInverted(t *testing.T) {
	vec := Vec[float64]{3, -4}
	invertedVec := vec.Invert()
	assert.Equal(t, Vec[float64]{-3, 4}, *invertedVec)
	assert.InDelta(t, vec.Inverted()[0], 3, 1e-8)
	assert.InDelta(t, vec.Inverted()[1], -4, 1e-8)
}

func TestNormalizeNormalized(t *testing.T) {
	vec := &Vec[float64]{4, 3}
	normalizedVec := vec.Normalize()
	assert.InDelta(t, 0.8, normalizedVec[0], 1e-7)
	assert.InDelta(t, 0.6, normalizedVec[1], 1e-7)
}

func TestAddSub(t *testing.T) {
	vec1 := Vec[float64]{1, 2}
	vec2 := Vec[float64]{9, 8}

	vec3 := vec1.Add(&vec2)
	assert.Equal(t, Vec[float64]{10, 10}, vec1)
	assert.Equal(t, Vec[float64]{9, 8}, vec2)
	assert.Equal(t, Vec[float64]{10, 10}, *vec3)
	vec4 := vec1.Sub(&vec2)
	assert.Equal(t, Vec[float64]{1, 2}, vec1)
	assert.Equal(t, Vec[float64]{9, 8}, vec2)
	assert.Equal(t, Vec[float64]{1, 2}, *vec4)
}

func TestAddedSubed(t *testing.T) {
	vec1 := Vec[float64]{1, 2}
	vec2 := Vec[float64]{9, 8}

	vec3 := vec1.Added(&vec2)
	assert.Equal(t, Vec[float64]{1, 2}, vec1)
	assert.Equal(t, Vec[float64]{9, 8}, vec2)
	assert.Equal(t, Vec[float64]{10, 10}, vec3)
	vec4 := vec2.Subed(&vec1)
	assert.Equal(t, Vec[float64]{1, 2}, vec1)
	assert.Equal(t, Vec[float64]{9, 8}, vec2)
	assert.Equal(t, Vec[float64]{8, 6}, vec4)
}

func TestMul(t *testing.T) {
	vec1 := &Vec[float64]{2, 3}
	vec2 := Vec[float64]{3, 4}
	vec3 := vec1.Mul(&vec2)

	assert.Equal(t, Vec[float64]{6, 12}, *vec1)
	assert.Equal(t, Vec[float64]{3, 4}, vec2)
	assert.Equal(t, Vec[float64]{6, 12}, *vec3)
}

func TestMuled(t *testing.T) {
	vec1 := &Vec[float64]{2, 3}
	vec2 := Vec[float64]{3, 4}
	vec3 := vec1.Muled(&vec2)

	assert.Equal(t, Vec[float64]{2, 3}, *vec1)
	assert.Equal(t, Vec[float64]{3, 4}, vec2)
	assert.Equal(t, Vec[float64]{6, 12}, vec3)
}

func TestRotateRotated(t *testing.T) {
	vec := Vec[float64]{1, 0}
	angle := math.Pi / 2
	rotatedVec := vec.Rotated(angle)
	assert.InDelta(t, 0., rotatedVec[0], 1e-7)
	assert.InDelta(t, 1., rotatedVec[1], 1e-7)
	vec.Rotate(angle)
	assert.Equal(t, rotatedVec, vec)
}

func TestRotateAroundPoint(t *testing.T) {
	point := Vec[float64]{0, 0}
	vec := Vec[float64]{1, 0}
	angle := math.Pi / 2
	rotatedVec := vec.RotateAroundPoint(&point, angle)
	assert.InDelta(t, 0., rotatedVec[0], 1e-8)
	assert.InDelta(t, 1., rotatedVec[1], 1e-8)
}

func TestRotate90Deg(t *testing.T) {
	vec := Vec[float64]{1, 0}
	vec.Rotate90DegLeft()
	assert.Equal(t, Vec[float64]{0, 1}, vec)
	vec = Vec[float64]{1, 0}
	vec.Rotate90DegRight()
	assert.Equal(t, Vec[float64]{0, -1}, vec)
}

func TestAngle(t *testing.T) {
	vec := Vec[float64]{1, 0}
	assert.InDelta(t, 0., vec.Angle(), 1e-8)
	vec = Vec[float64]{0, 1}
	assert.InDelta(t, math.Pi/2, vec.Angle(), 1e-8)
}

func TestAngle2(t *testing.T) {
	assert.InDelta(t, 0*radian45Degree, Angle(&Vec[float64]{1, 0}, &Vec[float64]{1, 0}), 1e-8, "0/360 degree angle, equal/parallell vectors")
	assert.InDelta(t, 1*radian45Degree, Angle(&Vec[float64]{1, 0}, &Vec[float64]{1, 1}), 1e-8, "45 degree angle")
	assert.InDelta(t, 2*radian45Degree, Angle(&Vec[float64]{1, 0}, &Vec[float64]{0, 1}), 1e-8, "90 degree angle, orthogonal vectors")
	assert.InDelta(t, 3*radian45Degree, Angle(&Vec[float64]{1, 0}, &Vec[float64]{-1, 1}), 1e-8, "135 degree angle")
	assert.InDelta(t, 4*radian45Degree, Angle(&Vec[float64]{1, 0}, &Vec[float64]{-1, 0}), 1e-8, "180 degree angle, inverted/anti parallell vectors")
	assert.InDelta(t, (8-5)*radian45Degree, Angle(&Vec[float64]{1, 0}, &Vec[float64]{-1, -1}), 1e-8, "225 degree angle")
	assert.InDelta(t, (8-6)*radian45Degree, Angle(&Vec[float64]{1, 0}, &Vec[float64]{0, -1}), 1e-8, "270 degree angle, orthogonal vectors")
	assert.InDelta(t, (8-7)*radian45Degree, Angle(&Vec[float64]{1, 0}, &Vec[float64]{1, -1}), 1e-8, "315 degree angle")
}

func TestClampClamped(t *testing.T) {
	vec := Vec[float64]{-1, 3}
	min := Vec[float64]{0, 0}
	max := Vec[float64]{2, 4}
	clampedVec := vec.Clone().Clamp(&min, &max)
	assert.Equal(t, Vec[float64]{0, 3}, *clampedVec)
	assert.Equal(t, Vec[float64]{0, 3}, vec.Clamped(&min, &max))
}

func TestClamp01(t *testing.T) {
	vec := Vec[float64]{-1, 2}
	clampedVec := vec.Clone().Clamp01()
	assert.Equal(t, Vec[float64]{0, 1}, *clampedVec)
	assert.Equal(t, Vec[float64]{0, 1}, vec.Clamped01())
}

func TestCosine(t *testing.T) {
	assert.InDelta(t, math.Cos(0*radian45Degree), Cosine(&Vec[float64]{1, 0}, &Vec[float64]{1, 0}), 1e-8, "0/360 degree angle, equal/parallell vectors")
	assert.InDelta(t, math.Cos(1*radian45Degree), Cosine(&Vec[float64]{1, 0}, &Vec[float64]{1, 1}), 1e-8, "45 degree angle")
	assert.InDelta(t, math.Cos(2*radian45Degree), Cosine(&Vec[float64]{1, 0}, &Vec[float64]{0, 1}), 1e-8, "90 degree angle, orthogonal vectors")
	assert.InDelta(t, math.Cos(3*radian45Degree), Cosine(&Vec[float64]{1, 0}, &Vec[float64]{-1, 1}), 1e-8, "135 degree angle")
	assert.InDelta(t, math.Cos(4*radian45Degree), Cosine(&Vec[float64]{1, 0}, &Vec[float64]{-1, 0}), 1e-8, "180 degree angle, inverted/anti parallell vectors")
	assert.InDelta(t, math.Cos(5*radian45Degree), Cosine(&Vec[float64]{1, 0}, &Vec[float64]{-1, -1}), 1e-8, "225 degree angle")
	assert.InDelta(t, math.Cos(6*radian45Degree), Cosine(&Vec[float64]{1, 0}, &Vec[float64]{0, -1}), 1e-8, "270 degree angle, orthogonal vectors")
	assert.InDelta(t, math.Cos(7*radian45Degree), Cosine(&Vec[float64]{1, 0}, &Vec[float64]{1, -1}), 1e-8, "315 degree angle")
}

func TestSinus(t *testing.T) {
	assert.InDelta(t, math.Sin(0*radian45Degree), Sinus(&Vec[float64]{1, 0}, &Vec[float64]{1, 0}), 1e-8, "0/360 degree angle, equal/parallell vectors")
	assert.InDelta(t, math.Sin(1*radian45Degree), Sinus(&Vec[float64]{1, 0}, &Vec[float64]{1, 1}), 1e-8, "45 degree angle")
	assert.InDelta(t, math.Sin(2*radian45Degree), Sinus(&Vec[float64]{1, 0}, &Vec[float64]{0, 1}), 1e-8, "90 degree angle, orthogonal vectors")
	assert.InDelta(t, math.Sin(3*radian45Degree), Sinus(&Vec[float64]{1, 0}, &Vec[float64]{-1, 1}), 1e-8, "135 degree angle")
	assert.InDelta(t, math.Sin(4*radian45Degree), Sinus(&Vec[float64]{1, 0}, &Vec[float64]{-1, 0}), 1e-8, "180 degree angle, inverted/anti parallell vectors")
	assert.InDelta(t, math.Sin(5*radian45Degree), Sinus(&Vec[float64]{1, 0}, &Vec[float64]{-1, -1}), 1e-8, "225 degree angle")
	assert.InDelta(t, math.Sin(6*radian45Degree), Sinus(&Vec[float64]{1, 0}, &Vec[float64]{0, -1}), 1e-8, "270 degree angle, orthogonal vectors")
	assert.InDelta(t, math.Sin(7*radian45Degree), Sinus(&Vec[float64]{1, 0}, &Vec[float64]{1, -1}), 1e-8, "315 degree angle")
}

func TestSetMinMax(t *testing.T) {
	vecMin := Vec[float64]{1, 5}
	vecMax := Vec[float64]{1, 5}
	minVal := Vec[float64]{2, 2}
	maxVal := Vec[float64]{3, 4}
	vecMin.SetMin(minVal)
	assert.Equal(t, Vec[float64]{1, 2}, vecMin)
	vecMax.SetMax(maxVal)
	assert.Equal(t, Vec[float64]{3, 5}, vecMax)
}

func TestLeftRightWinding(t *testing.T) {
	a := Vec[float64]{1.0, 0.0}
	for angle := 0; angle <= 360; angle += 15 {
		rad := (math.Pi / 180.0) * float64(angle)

		bx := clampDecimals(math.Cos(rad), 4)
		by := clampDecimals(math.Sin(rad), 4)
		b := Vec[float64]{bx, by}

		t.Run("left winding angle "+strconv.Itoa(angle), func(t *testing.T) {
			lw := isLeftWinding(&a, &b)
			rw := isRightWinding(&a, &b)

			if angle%180 == 0 {
				// No winding at 0, 180 and 360 degrees
				if lw || rw {
					t.Errorf("Neither left or right winding should be true on angle %d. Left winding=%t, right winding=%t", angle, lw, rw)
				}
			} else if angle < 180 {
				// Left winding at 0 < angle < 180
				if !lw || rw {
					t.Errorf("Left winding should be true (not right winding) on angle %d. Left winding=%t, right winding=%t", angle, lw, rw)
				}
			} else if angle > 180 {
				// Right winding at 180 < angle < 360
				if lw || !rw {
					t.Errorf("Right winding should be true (not left winding) on angle %d. Left winding=%t, right winding=%t", angle, lw, rw)
				}
			}
		})
	}
}

func TestDataConversion(t *testing.T) {
	other := NewVec2[float64](1, 2)
	vec := From[float64](other)
	assert.Equal(t, Vec[float64]{1, 2}, vec)
}

func TestParse(t *testing.T) {
	str := "3.14 2.71"
	vec, err := Parse[float64](str)
	assert.NoError(t, err)
	assert.Equal(t, Vec[float64]{3.14, 2.71}, vec)
}

func TestPracticallyEquals(t *testing.T) {
	vec1 := Vec[float64]{1.0, 2.0}
	vec2 := Vec[float64]{1.0000001, 2.0000001}
	vec3 := Vec[float64]{1.1, 2.1}

	assert.True(t, vec1.PracticallyEquals(&vec2, 1e-6), "Vectors should be practically equal")
	assert.False(t, vec1.PracticallyEquals(&vec3, 1e-6), "Vectors should not be practically equal")
}

func TestPracticallyEqualsStatic(t *testing.T) {
	assert.True(t, PracticallyEquals(1.0, 1.0000001, 1e-6), "Values should be practically equal")
	assert.False(t, PracticallyEquals(1.0, 1.1, 1e-6), "Values should not be practically equal")
}

func TestRotate90DegLeft(t *testing.T) {
	vec := Vec[float64]{1, 0}
	vec.Rotate90DegLeft()

	expected := Vec[float64]{0, 1}
	assert.Equal(t, expected, vec)
}

func TestRotate90DegRight(t *testing.T) {
	vec := Vec[float64]{1, 0}
	vec.Rotate90DegRight()

	expected := Vec[float64]{0, -1}
	assert.Equal(t, expected, vec)
}

func TestCross(t *testing.T) {
	vecA := Vec[float64]{1, 0}
	vecB := Vec[float64]{0, 1}
	cross := Cross(&vecA, &vecB)

	expected := Vec[float64]{0, 0}
	assert.Equal(t, expected, cross)
}

func TestCrossScalar(t *testing.T) {
	vecA := Vec[float64]{1, 0}
	vecB := Vec[float64]{0, 1}
	cross := cross(&vecA, &vecB)

	expected := float64(1)
	assert.Equal(t, expected, cross)
}

func TestAngleFunction(t *testing.T) {
	vecA := Vec[float64]{1, 0}
	vecB := Vec[float64]{0, 1}
	angle := angle(&vecA, &vecB)

	expected := math.Pi / 2
	assert.InDelta(t, expected, float64(angle), 1e-8)
}

func TestIsLeftWinding(t *testing.T) {
	vecA := Vec[float64]{1, 0}
	vecB := Vec[float64]{0, 1}
	assert.True(t, IsLeftWinding(&vecA, &vecB), "Angle from A to B should be left winding")

	vecC := Vec[float64]{0, 1}
	vecD := Vec[float64]{1, 0}
	assert.False(t, IsLeftWinding(&vecC, &vecD), "Angle from C to D should not be left winding")
}

func TestIsRightWinding(t *testing.T) {
	vecA := Vec[float64]{1, 0}
	vecB := Vec[float64]{0, 1}
	assert.False(t, IsRightWinding(&vecA, &vecB), "Angle from A to B should not be right winding")

	vecC := Vec[float64]{0, 1}
	vecD := Vec[float64]{1, 0}
	assert.True(t, IsRightWinding(&vecC, &vecD), "Angle from C to D should be right winding")
}

func TestPointSegmentDistance(t *testing.T) {
	p1 := Vec[float64]{1, 1}
	x1 := Vec[float64]{0, 0}
	x2 := Vec[float64]{2, 0}
	distance := PointSegmentDistance(&p1, &x1, &x2)

	expected := float64(1)
	assert.InDelta(t, expected, distance, 1e-8)
}

func TestPointSegmentVerticalPoint(t *testing.T) {
	p1 := Vec[float64]{1, 1}
	x1 := Vec[float64]{0, 0}
	x2 := Vec[float64]{2, 0}
	verticalPoint := PointSegmentVerticalPoint(&p1, &x1, &x2)

	expected := Vec[float64]{1, 0}
	assert.Equal(t, expected, *verticalPoint)
}

func TestInterpolate(t *testing.T) {
	vecA := Vec[float64]{1, 2}
	vecB := Vec[float64]{3, 4}
	interpolated := Interpolate(&vecA, &vecB, 0.5)

	expected := Vec[float64]{2, 3}
	assert.Equal(t, expected, interpolated)
}

func TestClamp(t *testing.T) {
	vec := Vec[float64]{-1, 3}
	min := Vec[float64]{0, 0}
	max := Vec[float64]{2, 4}
	clampedVec := vec.Clone().Clamp(&min, &max)

	expected := Vec[float64]{0, 3}
	assert.Equal(t, expected, *clampedVec)
}

func TestClamped(t *testing.T) {
	vec := Vec[float64]{-1, 3}
	min := Vec[float64]{0, 0}
	max := Vec[float64]{2, 4}
	clampedVec := vec.Clamped(&min, &max)

	expected := Vec[float64]{0, 3}
	assert.Equal(t, expected, clampedVec)
}

func TestClamped01(t *testing.T) {
	vec := Vec[float64]{-1, 2}
	clampedVec := vec.Clamped01()

	expected := Vec[float64]{0, 1}
	assert.Equal(t, expected, clampedVec)
}

// clampDecimals clamps a floating-point number to a specified number of decimal places by rounding.
// It takes a value and the number of decimals to retain, then rounds the input to that decimal precision.
func clampDecimals[T float64 | float32](decimalValue T, amountDecimals T) T {
	factor := math.Pow(10, float64(amountDecimals))
	return T(math.Round(float64(decimalValue)*factor) / factor)
}
