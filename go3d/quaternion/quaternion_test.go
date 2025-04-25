package quaternion

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"pinkey.ltd/xr/go3d/vec3"
	"pinkey.ltd/xr/go3d/vec4"
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

func TestFromAxisAngle(t *testing.T) {
	axis := &vec3.Vec[float64]{1, 0, 0}
	angle := math.Pi / 2
	q := FromAxisAngle(axis, angle)

	expected := H[float64]{math.Sqrt(2) / 2, 0, 0, math.Sqrt(2) / 2}
	assert.InDelta(t, expected[0], q[0], 1e-9)
	assert.InDelta(t, expected[1], q[1], 1e-9)
	assert.InDelta(t, expected[2], q[2], 1e-9)
	assert.InDelta(t, expected[3], q[3], 1e-9)
}

func TestFromXAxisAngle(t *testing.T) {
	angle := math.Pi / 2
	q := FromXAxisAngle[float64](angle)

	expected := H[float64]{math.Sqrt(2) / 2, 0, 0, math.Sqrt(2) / 2}
	assert.InDelta(t, expected[0], q[0], 1e-9)
	assert.InDelta(t, expected[1], q[1], 1e-9)
	assert.InDelta(t, expected[2], q[2], 1e-9)
	assert.InDelta(t, expected[3], q[3], 1e-9)
}

func TestFromYAxisAngle(t *testing.T) {
	angle := math.Pi / 2
	q := FromYAxisAngle[float64](angle)

	expected := H[float64]{0, math.Sqrt(2) / 2, 0, math.Sqrt(2) / 2}
	assert.InDelta(t, expected[0], q[0], 1e-9)
	assert.InDelta(t, expected[1], q[1], 1e-9)
	assert.InDelta(t, expected[2], q[2], 1e-9)
	assert.InDelta(t, expected[3], q[3], 1e-9)
}

func TestFromZAxisAngle(t *testing.T) {
	angle := math.Pi / 2
	q := FromZAxisAngle[float64](angle)

	expected := H[float64]{0, 0, math.Sqrt(2) / 2, math.Sqrt(2) / 2}
	assert.InDelta(t, expected[0], q[0], 1e-9)
	assert.InDelta(t, expected[1], q[1], 1e-9)
	assert.InDelta(t, expected[2], q[2], 1e-9)
	assert.InDelta(t, expected[3], q[3], 1e-9)
}

func TestFromVec4(t *testing.T) {
	v := &vec4.Vec[float64]{1, 2, 3, 4}
	q := FromVec4(v)

	expected := H[float64]{1, 2, 3, 4}
	assert.Equal(t, expected, q)
}

func TestMul(t *testing.T) {
	q1 := H[float64]{1, 0, 0, 0}
	q2 := H[float64]{0, 1, 0, 0}
	q := Mul(&q1, &q2)

	expected := H[float64]{0, 0, 1, 0}
	assert.InDelta(t, expected[0], q[0], 1e-9)
	assert.InDelta(t, expected[1], q[1], 1e-9)
	assert.InDelta(t, expected[2], q[2], 1e-9)
	assert.InDelta(t, expected[3], q[3], 1e-9)
}

func TestSlerp(t *testing.T) {
	q1 := H[float64]{1, 0, 0, 0}
	q2 := H[float64]{0, 1, 0, 0}
	q := Slerp(&q1, &q2, 0.5)

	expected := H[float64]{math.Sqrt(2) / 2, math.Sqrt(2) / 2, 0, 0}
	assert.InDelta(t, expected[0], q[0], 1e-9)
	assert.InDelta(t, expected[1], q[1], 1e-9)
	assert.InDelta(t, expected[2], q[2], 1e-9)
	assert.InDelta(t, expected[3], q[3], 1e-9)
}

func TestVec3Diff(t *testing.T) {
	v1 := &vec3.Vec[float64]{1, 0, 0}
	v2 := &vec3.Vec[float64]{0, 1, 0}
	q := Vec3Diff(v1, v2)

	expected := H[float64]{0, 0, math.Sqrt(2) / 2, math.Sqrt(2) / 2}
	assert.InDelta(t, expected[0], q[0], 1e-9)
	assert.InDelta(t, expected[1], q[1], 1e-9)
	assert.InDelta(t, expected[2], q[2], 1e-9)
	assert.InDelta(t, expected[3], q[3], 1e-9)
}

func TestIsUnitQuat(t *testing.T) {
	q := H[float64]{math.Sqrt(2) / 2, math.Sqrt(2) / 2, 0, 0}
	assert.True(t, q.IsUnitQuat(1e-9))

	q2 := H[float64]{1, 2, 3, 4}
	assert.False(t, q2.IsUnitQuat(1e-9))
}

func TestSetShortestRotation(t *testing.T) {
	q1 := H[float64]{1, 0, 0, 0}
	q2 := H[float64]{-1, 0, 0, 0}
	q1.SetShortestRotation(&q2)

	expected := H[float64]{1, 0, 0, 0}
	assert.Equal(t, expected, q1)
}

func TestIsShortestRotation(t *testing.T) {
	q1 := H[float64]{1, 0, 0, 0}
	q2 := H[float64]{-1, 0, 0, 0}
	assert.False(t, IsShortestRotation(&q1, &q2))

	q3 := H[float64]{1, 0, 0, 0}
	q4 := H[float64]{1, 0, 0, 0}
	assert.True(t, IsShortestRotation(&q3, &q4))
}

func TestParse(t *testing.T) {
	s := "1 2 3 4"
	q, err := Parse[float64](s)
	assert.NoError(t, err)
	assert.Equal(t, H[float64]{1, 2, 3, 4}, q)
}

func TestRotatedVec3(t *testing.T) {
	q := H[float64]{math.Sqrt(2) / 2, 0, 0, math.Sqrt(2) / 2}
	v := &vec3.Vec[float64]{1, 0, 0}
	rotated := q.RotatedVec3(v)

	expected := vec3.Vec[float64]{1, 0, 0}
	assert.InDelta(t, expected[0], rotated[0], 1e-9)
	assert.InDelta(t, expected[1], rotated[1], 1e-9)
	assert.InDelta(t, expected[2], rotated[2], 1e-9)
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

func TestMul4(t *testing.T) {
	// 测试用例1：恒等四元数相乘
	identity := H[float64]{0, 0, 0, 1}
	q1 := H[float64]{0.5, 0.5, 0.5, 0.5}
	result := Mul4(&identity, &identity, &identity, &q1)
	if !isEqual(result, q1) {
		t.Errorf("Mul4 with identity quaternions failed: got %v, want %v", result, q1)
	}

	// 测试用例2：旋转组合
	// 用不同的轴角创建四个四元数
	qa := FromXAxisAngle[float64](math.Pi / 4) // 绕X轴旋转45度
	qb := FromYAxisAngle[float64](math.Pi / 3) // 绕Y轴旋转60度
	qc := FromZAxisAngle[float64](math.Pi / 6) // 绕Z轴旋转30度
	qd := FromXAxisAngle[float64](math.Pi / 2) // 绕X轴旋转90度

	// 直接计算结果
	result = Mul4(&qa, &qb, &qc, &qd)

	// 分步计算结果进行比较
	expected := Mul(&qa, &qb)
	expected = Mul(&expected, &qc)
	expected = Mul(&expected, &qd)

	if !isEqual(result, expected) {
		t.Errorf("Mul4 rotation combination failed: got %v, want %v", result, expected)
	}

	// 测试用例3：float32类型
	qa32 := FromXAxisAngle[float32](math.Pi / 4)
	qb32 := FromYAxisAngle[float32](math.Pi / 3)
	qc32 := FromZAxisAngle[float32](math.Pi / 6)
	qd32 := FromXAxisAngle[float32](math.Pi / 2)

	result32 := Mul4(&qa32, &qb32, &qc32, &qd32)

	expected32 := Mul(&qa32, &qb32)
	expected32 = Mul(&expected32, &qc32)
	expected32 = Mul(&expected32, &qd32)

	if !isEqual(result32, expected32) {
		t.Errorf("Mul4 with float32 failed: got %v, want %v", result32, expected32)
	}
}

// 帮助函数：比较两个四元数是否近似相等
func isEqual[T float64 | float32](a, b H[T]) bool {
	const epsilon = 1e-6 // 浮点比较的误差范围
	for i := 0; i < 4; i++ {
		if math.Abs(float64(a[i]-b[i])) > epsilon {
			return false
		}
	}
	return true
}
