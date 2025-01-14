package mat2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"pinkey.ltd/xr/pkg/go3d/vec2"
)

var IdentTest = Mat[float64]{
	vec2.Vec[float64]{1, 0},
	vec2.Vec[float64]{0, 1},
}

func TestParse(t *testing.T) {
	// 测试用例
	testCases := []struct {
		input    string
		expected Mat[float64]
		err      error
	}{
		{"1 2 3 4", Mat[float64]{vec2.Vec[float64]{1, 2}, vec2.Vec[float64]{3, 4}}, nil},
		{"1.5 2.5 3.5 4.5", Mat[float64]{vec2.Vec[float64]{1.5, 2.5}, vec2.Vec[float64]{3.5, 4.5}}, nil},
		{"1 2 3", Mat[float64]{}, fmt.Errorf("unexpected EOF")}, // 输入不完整
	}

	for _, tc := range testCases {
		result, err := Parse[float64](tc.input)
		if tc.err != nil {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		}
	}
}

func TestScale(t *testing.T) {
	m := Mat[float64]{vec2.Vec[float64]{1, 2}, vec2.Vec[float64]{3, 4}}
	m.Scale(2.0)

	expected := Mat[float64]{vec2.Vec[float64]{2, 2}, vec2.Vec[float64]{3, 8}}
	assert.Equal(t, expected, m)
}

func TestScaled(t *testing.T) {
	m := Mat[float64]{vec2.Vec[float64]{1, 2}, vec2.Vec[float64]{3, 4}}
	scaled := m.Scaled(2.0)

	expected := Mat[float64]{vec2.Vec[float64]{2, 2}, vec2.Vec[float64]{3, 8}}
	assert.Equal(t, expected, scaled)
}

func TestScaling(t *testing.T) {
	m := Mat[float64]{vec2.Vec[float64]{1, 2}, vec2.Vec[float64]{3, 4}}
	scaling := m.Scaling()

	expected := vec2.Vec[float64]{1, 4}
	assert.Equal(t, expected, scaling)
}

func TestSetScaling(t *testing.T) {
	m := Mat[float64]{vec2.Vec[float64]{1, 2}, vec2.Vec[float64]{3, 4}}
	newScaling := &vec2.Vec[float64]{5, 6}
	m.SetScaling(newScaling)

	expected := Mat[float64]{vec2.Vec[float64]{5, 2}, vec2.Vec[float64]{3, 6}}
	assert.Equal(t, expected, m)
}

func TestTrace(t *testing.T) {
	m := Mat[float64]{vec2.Vec[float64]{1, 2}, vec2.Vec[float64]{3, 4}}
	trace := m.Trace()

	expected := 5.0
	assert.Equal(t, expected, trace)
}

func TestAssignMul(t *testing.T) {
	a := Mat[float64]{vec2.Vec[float64]{1, 2}, vec2.Vec[float64]{3, 4}}
	b := Mat[float64]{vec2.Vec[float64]{5, 6}, vec2.Vec[float64]{7, 8}}
	result := Mat[float64]{}
	result.AssignMul(&a, &b)

	expected := Mat[float64]{
		vec2.Vec[float64]{19, 22},
		vec2.Vec[float64]{43, 50},
	}
	assert.Equal(t, expected, result)
}

func TestMulVec2(t *testing.T) {
	m := Mat[float64]{vec2.Vec[float64]{1, 2}, vec2.Vec[float64]{3, 4}}
	v := &vec2.Vec[float64]{5, 6}
	result := m.MulVec2(v)

	expected := vec2.Vec[float64]{17, 39}
	assert.Equal(t, expected, result)
}

func TestTransformVec2(t *testing.T) {
	m := Mat[float64]{vec2.Vec[float64]{1, 2}, vec2.Vec[float64]{3, 4}}
	v := &vec2.Vec[float64]{5, 6}
	m.TransformVec2(v)

	expected := vec2.Vec[float64]{17, 39}
	assert.Equal(t, expected, *v)
}

func TestPracticallyEquals(t *testing.T) {
	m1 := Mat[float64]{vec2.Vec[float64]{1.000001, 2.000001}, vec2.Vec[float64]{3.000001, 4.000001}}
	m2 := Mat[float64]{vec2.Vec[float64]{1.0, 2.0}, vec2.Vec[float64]{3.0, 4.0}}

	assert.True(t, m1.PracticallyEquals(&m2, 0.00001))
	assert.False(t, m1.PracticallyEquals(&m2, 0.0000001))
}

func TestInvert(t *testing.T) {
	m := Mat[float64]{vec2.Vec[float64]{4, 7}, vec2.Vec[float64]{2, 6}}
	_, err := m.Invert()

	assert.NoError(t, err)
	expected := Mat[float64]{vec2.Vec[float64]{0.6, -0.7}, vec2.Vec[float64]{-0.2, 0.4}}
	assert.InDelta(t, expected[0][0], m[0][0], 1e-7)
	assert.InDelta(t, expected[0][1], m[0][1], 1e-7)
	assert.InDelta(t, expected[1][0], m[1][0], 1e-7)
	assert.InDelta(t, expected[1][1], m[1][1], 1e-7)
}

func TestInverted(t *testing.T) {
	m := Mat[float64]{vec2.Vec[float64]{4, 7}, vec2.Vec[float64]{2, 6}}
	inverted, err := m.Inverted()

	assert.NoError(t, err)
	expected := Mat[float64]{vec2.Vec[float64]{0.6, -0.7}, vec2.Vec[float64]{-0.2, 0.4}}
	assert.InDelta(t, expected[0][0], inverted[0][0], 1e-7)
	assert.InDelta(t, expected[0][1], inverted[0][1], 1e-7)
	assert.InDelta(t, expected[1][0], inverted[1][0], 1e-7)
	assert.InDelta(t, expected[1][1], inverted[1][1], 1e-7)
}

func TestDeterminant(t *testing.T) {
	m := Mat[float64]{vec2.Vec[float64]{1, 2}, vec2.Vec[float64]{3, 4}}
	det := m.Determinant()

	expected := -2.0
	assert.Equal(t, expected, det)
}

func TestMat2String(t *testing.T) {
	m1 := Mat[float64]{vec2.Vec[float64]{1, 0}, vec2.Vec[float64]{0, 1}}
	m2 := Mat[float64]{vec2.Vec[float64]{2, 3}, vec2.Vec[float64]{4, 5}}

	assert.Equal(t, "1 0 0 1", m1.String())
	assert.Equal(t, "2 3 4 5", m2.String())
}

func TestMat2RowsColsSize(t *testing.T) {
	m := Mat[float64]{}

	assert.Equal(t, 2, m.Rows())
	assert.Equal(t, 2, m.Cols())
	assert.Equal(t, 4, m.Size())
}

func TestMat2Slice(t *testing.T) {
	m := Mat[float64]{vec2.Vec[float64]{1, 2}, vec2.Vec[float64]{3, 4}}

	assert.Equal(t, []float64{1, 2, 3, 4}, m.Slice())
}

func TestMat2Get(t *testing.T) {
	m := Mat[float64]{vec2.Vec[float64]{1, 2}, vec2.Vec[float64]{3, 4}}

	assert.Equal(t, 1.0, m.Get(0, 0))
	assert.Equal(t, 2.0, m.Get(0, 1))
	assert.Equal(t, 3.0, m.Get(1, 0))
	assert.Equal(t, 4.0, m.Get(1, 1))

	for row := 0; row < m.Rows(); row++ {
		for col := 0; col < m.Cols(); col++ {
			expected := float64(row*2 + col + 1)
			if got := m.Get(row, col); got != expected {
				t.Errorf("Get(%d, %d) = %v, want %v", row, col, got, expected)
			}
		}
	}
}

func TestMat2Set(t *testing.T) {
	// 初始化一个 2x2 矩阵
	mat := Mat[float64]{
		vec2.Vec[float64]{1, 0},
		vec2.Vec[float64]{0, 1},
	}

	// 设置矩阵中的一个值
	col, row, newValue := 0, 1, 5.0
	mat.Set(col, row, newValue)

	// 验证设置后的值是否正确
	expectedMat := Mat[float64]{
		vec2.Vec[float64]{1, 5},
		vec2.Vec[float64]{0, 1},
	}
	assert.Equal(t, mat, expectedMat)
}

func TestMat2IsZero(t *testing.T) {
	zero := Mat[float64]{vec2.Vec[float64]{0, 0}, vec2.Vec[float64]{0, 0}}
	nonZero := Mat[float64]{vec2.Vec[float64]{1, 0}, vec2.Vec[float64]{0, 1}}

	assert.True(t, zero.IsZero())
	assert.False(t, nonZero.IsZero())
}

func TestMat2Scale(t *testing.T) {
	m := Mat[float64]{vec2.Vec[float64]{1, 2}, vec2.Vec[float64]{3, 4}}
	m.Scale(2.0)

	assert.Equal(t, Mat[float64]{vec2.Vec[float64]{2, 2}, vec2.Vec[float64]{3, 8}}, m)
}

func TestT_Transposed(t *testing.T) {
	matrix := Mat[float64]{
		vec2.Vec[float64]{1, 2},
		vec2.Vec[float64]{3, 4},
	}

	expectedMatrix := Mat[float64]{
		vec2.Vec[float64]{1, 3},
		vec2.Vec[float64]{2, 4},
	}

	transposedMatrix := matrix.Transposed()

	assert.Equal(t, expectedMatrix, transposedMatrix)
}

func TestT_Transpose(t *testing.T) {
	matrix := Mat[float64]{
		vec2.Vec[float64]{10, 20},
		vec2.Vec[float64]{30, 40},
	}

	expectedMatrix := Mat[float64]{
		vec2.Vec[float64]{10, 30},
		vec2.Vec[float64]{20, 40},
	}

	transposedMatrix := matrix
	transposedMatrix.Transpose()

	if transposedMatrix != expectedMatrix {
		t.Errorf("matrix transposed wrong: %v --> %v", matrix, transposedMatrix)
	}
}

func TestDeterminant_1(t *testing.T) {
	detId := IdentTest.Determinant()
	if detId != 1 {
		t.Errorf("Wrong determinant for identity matrix: %f", detId)
	}
}

func TestDeterminant_2(t *testing.T) {
	detTwo := IdentTest
	detTwo[0][0] = 2
	if det := detTwo.Determinant(); det != (2*1 - 0*0) {
		t.Errorf("Wrong determinant: %f", det)
	}
}

func TestDeterminant_3(t *testing.T) {
	scale2 := IdentTest.Scaled(2)
	if det := scale2.Determinant(); det != (2*2 - 0*0) {
		t.Errorf("Wrong determinant: %f", det)
	}
}

func TestDeterminant_4(t *testing.T) {
	row1changed, _ := Parse[float64]("3 0   2 2")
	if det := row1changed.Determinant(); det != (3*2 - 0*2) {
		t.Errorf("Wrong determinant: %f", det)
	}
}

func TestDeterminant_5(t *testing.T) {
	row12changed, _ := Parse[float64]("3 1 2 5")
	if det := row12changed.Determinant(); det != (3*5 - 1*2) {
		t.Errorf("Wrong determinant: %f", det)
	}
}

func TestDeterminant_6(t *testing.T) {
	row123changed, _ := Parse[float64]("3 1 2 5")
	if det := row123changed.Determinant(); det != (3*5 - 2*1) {
		t.Errorf("Wrong determinant for 3x3 matrix: %f", det)
	}
}

func TestDeterminant_7(t *testing.T) {
	randomMatrix, err := Parse[float64]("0.43685 0.81673 0.16600 0.40608")
	randomMatrix.Transpose()
	// if err != nil {
	// 	t.Errorf("Could not parse random matrix: %v", err)
	// }
	assert.NoError(t, err, "Could not parse random matrix")
	// if det := randomMatrix.Determinant(); PracticallyEquals(det, 0.0418189) {
	// 	t.Errorf("Wrong determinant for random sub 3x3 matrix: %f", det)
	// }
	assert.InDelta(t, 0.0418189, randomMatrix.Determinant(), 1e-7)
}

func TestInvert_ok(t *testing.T) {
	inv := Mat[float64]{vec2.Vec[float64]{4, -2}, vec2.Vec[float64]{8, -3}}
	_, err := inv.Invert()

	if err != nil {
		t.Error("Inverse not computed correctly", err)
	}

	invExpected := Mat[float64]{vec2.Vec[float64]{-3.0 / 4.0, 1.0 / 2.0}, vec2.Vec[float64]{-2, 1}}
	if inv != invExpected {
		t.Errorf("Inverse not computed correctly: %#v", inv)
	}
}

func TestInvert_nok_1(t *testing.T) {
	inv := Mat[float64]{vec2.Vec[float64]{1, 1}, vec2.Vec[float64]{1, 1}}
	_, err := inv.Inverted()
	if err == nil {
		t.Error("Inverse should not be possible", err)
	}
}

func TestInvert_nok_2(t *testing.T) {
	inv := Mat[float64]{vec2.Vec[float64]{2, 0}, vec2.Vec[float64]{1, 0}}
	_, err := inv.Inverted()
	if err == nil {
		t.Error("Inverse should not be possible", err)
	}
}
