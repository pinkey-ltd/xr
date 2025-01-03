package quaternion

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"pinkey.ltd/xr/pkg/go3d/vec3"
	"pinkey.ltd/xr/pkg/go3d/vec4"
	"testing"
)

func TestVec4(t *testing.T) {
	q := H[float64]{1, 2, 3, 4}
	expected := vec4.Vec[float64]{1, 2, 3, 4}
	q2 := H[float32]{1, 2, 3, 4}
	expected2 := vec4.Vec[float32]{1, 2, 3, 4}

	assert.Equal(t, expected, q.Vec4())
	assert.Equal(t, expected2, q2.Vec4())
}

func TestString(t *testing.T) {
	q := H[float64]{1, 2, 3, 4}
	expected := "1 2 3 4"
	assert.Equal(t, expected, q.String())
}

// TODO: fix
func TestAxisAngle(t *testing.T) {
	testCases := []struct {
		quat  H[float64]
		axis  vec3.Vec[float64]
		angle float64
	}{
		{H[float64]{1, 0, 0, 0}, vec3.Vec[float64]{1, 0, 0}, math.Pi},
		{H[float64]{0, 1, 0, 0}, vec3.Vec[float64]{0, 1, 0}, math.Pi},
		{H[float64]{0, 0, 1, 0}, vec3.Vec[float64]{0, 0, 1}, math.Pi},
		{H[float64]{0, 0, 0, 1}, vec3.Vec[float64]{0, 0, 0}, 0},
	}

	for _, tc := range testCases {
		axis, angle := tc.quat.AxisAngle()
		//if !axis.Equals(tc.axis) || math.Abs(angle-tc.angle) > 1e-9 {
		//	t.Errorf("AxisAngle() = (%v, %v), want (%v, %v)", axis, angle, tc.axis, tc.angle)
		//}
		assert.Equal(t, tc.axis, axis)
		assert.InDelta(t, tc.angle, angle, 1e-9)
	}
}

func TestNorm(t *testing.T) {
	q := H[float64]{1, 2, 3, 4}
	if got := q.Norm(); got != 30 {
		t.Errorf("Norm() = %v, want %v", got, 30)
	}
}

func TestNormalize(t *testing.T) {
	q := H[float64]{1, 2, 3, 4}
	q.Normalize()

	assert.InDelta(t, 1, q.Norm(), 1e-8)
}

func TestNormalized(t *testing.T) {
	q := H[float64]{1, 2, 3, 4}
	normalized := q.Normalized()

	assert.InDelta(t, 1, normalized.Norm(), 1e-8)
}

func TestNegate(t *testing.T) {
	q := H[float64]{1, 2, 3, 4}
	q.Negate()

	assert.Equal(t, H[float64]{-1, -2, -3, -4}, q)
}

func TestNegated(t *testing.T) {
	q := H[float64]{1, 2, 3, 4}

	assert.Equal(t, H[float64]{-1, -2, -3, -4}, q.Negated())
}

func TestInvert(t *testing.T) {
	q := H[float64]{1, 2, 3, 4}
	q.Invert()

	assert.Equal(t, H[float64]{-1, -2, -3, 4}, q)
}

func TestIsIdent(t *testing.T) {
	q0 := H[float64]{0, 0, 0, 1}
	nq0 := H[float64]{1, 0, 0, 1}

	assert.True(t, q0.IsIdent(), "the quaternion represents the identity")
	assert.False(t, nq0.IsIdent())
}

func TestInverted(t *testing.T) {
	q := H[float64]{1, 2, 3, 4}

	assert.Equal(t, H[float64]{-1, -2, -3, 4}, q.Inverted())
}

func TestQuaternionRotateVec3(t *testing.T) {
	eularAngles := []vec3.Vec[float64]{
		{90, 20, 21},
		{-90, 0, 0},
		{28, 1043, -38},
	}
	vecs := []vec3.Vec[float64]{
		{2, 3, 4},
		{1, 3, -2},
		{-6, 2, 9},
	}
	for _, vec := range vecs {
		for _, eularAngle := range eularAngles {
			func() {
				q := FromEulerAngles[float64](eularAngle[1]*math.Pi/180.0, eularAngle[0]*math.Pi/180.0, eularAngle[2]*math.Pi/180.0)
				vecR1 := vec
				vecR2 := vec
				magSqr := vecR1.LengthSqr()
				rotateAndNormalizeVec3(&q, &vecR2)
				q.RotateVec3(&vecR1)
				//vecd := q.RotatedVec3(&vec)
				magSqr2 := vecR1.LengthSqr()

				//if !vecd.PracticallyEquals(&vecR1, 1e-15) {
				//	t.Logf("test case %v rotates %v failed - vector rotation: %+v, %+v\n", eularAngle, vec, vecd, vecR1)
				//	t.Fail()
				//}

				angle := vec3.Angle(&vecR1, &vecR2)
				length := math.Abs(magSqr - magSqr2)

				if angle > 1e-7 {
					t.Logf("test case %v rotates %v failed - angle difference to large\n", eularAngle, vec)
					t.Logf("vectors: %+v, %+v\n", vecR1, vecR2)
					t.Logf("angle: %v\n", angle)
					t.Fail()
				}

				if length > 1e-12 {
					t.Logf("test case %v rotates %v failed - squared length difference to large\n", eularAngle, vec)
					t.Logf("vectors: %+v %+v\n", vecR1, vecR2)
					t.Logf("squared lengths: %v, %v\n", magSqr, magSqr2)
					t.Fail()
				}
			}()
		}
	}
}

func TestToEulerAngles(t *testing.T) {
	specialValues := []float64{-5, -math.Pi, -2, -math.Pi / 2, 0, math.Pi / 2, 2.4, math.Pi, 3.9}
	for _, x := range specialValues {
		for _, y := range specialValues {
			for _, z := range specialValues {
				quat1 := FromEulerAngles[float64](y, x, z)
				ry, rx, rz := quat1.ToEulerAngles()
				quat2 := FromEulerAngles[float64](ry, rx, rz)
				// quat must be equivalent
				const e64 = 1e-14
				cond1 := math.Abs(quat1[0]-quat2[0]) < e64 && math.Abs(quat1[1]-quat2[1]) < e64 && math.Abs(quat1[2]-quat2[2]) < e64 && math.Abs(quat1[3]-quat2[3]) < e64
				cond2 := math.Abs(quat1[0]+quat2[0]) < e64 && math.Abs(quat1[1]+quat2[1]) < e64 && math.Abs(quat1[2]+quat2[2]) < e64 && math.Abs(quat1[3]+quat2[3]) < e64
				if !cond1 && !cond2 {
					fmt.Printf("test case %v, %v, %v failed\n", x, y, z)
					fmt.Printf("result is %v, %v, %v\n", rx, ry, rz)
					fmt.Printf("quat1 is %v\n", quat1)
					fmt.Printf("quat2 is %v\n", quat2)
					t.Fail()
				}
			}
		}
	}
}

// RotateVec3 rotates v by the rotation represented by the quaternion.
func rotateAndNormalizeVec3[T float64 | float32](quat *H[T], v *vec3.Vec[T]) {
	qv := H[T]{v[0], v[1], v[2], 0}
	inv := quat.Inverted()
	q := Mul3(quat, &qv, &inv)
	v[0] = q[0]
	v[1] = q[1]
	v[2] = q[2]
}
