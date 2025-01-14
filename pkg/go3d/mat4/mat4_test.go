package mat4

import (
	"math"
	"pinkey.ltd/xr/pkg/go3d/mat3"
	"pinkey.ltd/xr/pkg/go3d/vec3"
	"pinkey.ltd/xr/pkg/go3d/vec4"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMat4String(t *testing.T) {
	m := Mat[float64]{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	expected := "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16"
	assert.Equal(t, expected, m.String())
}

func TestMat4RowsColsSize(t *testing.T) {
	m := Mat[float64]{}
	assert.Equal(t, 4, m.Rows())
	assert.Equal(t, 4, m.Cols())
	assert.Equal(t, 16, m.Size())
}

func TestMat4IsZero(t *testing.T) {
	var zero Mat[float64]
	nonZero := Mat[float64]{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}
	assert.True(t, zero.IsZero())
	assert.False(t, nonZero.IsZero())
}

func TestMat4Set(t *testing.T) {
	// 初始化一个4x4矩阵
	m := Mat[float64]{
		vec4.Vec[float64]{1, 0, 0, 0},
		vec4.Vec[float64]{0, 1, 0, 0},
		vec4.Vec[float64]{0, 0, 1, 0},
		vec4.Vec[float64]{0, 0, 0, 1},
	}

	// 定义测试用例
	testCases := []struct {
		col      int
		row      int
		value    float64
		expected Mat[float64]
	}{
		{0, 0, 5, Mat[float64]{{5, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}},
		{1, 1, -1, Mat[float64]{{5, 0, 0, 0}, {0, -1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}},
		{2, 2, 7.5, Mat[float64]{{5, 0, 0, 0}, {0, -1, 0, 0}, {0, 0, 7.5, 0}, {0, 0, 0, 1}}},
		{3, 3, 100, Mat[float64]{{5, 0, 0, 0}, {0, -1, 0, 0}, {0, 0, 7.5, 0}, {0, 0, 0, 100}}},
	}

	// 遍历并执行测试用例
	for _, tc := range testCases {
		// 调用 Set 方法
		m.Set(tc.col, tc.row, tc.value)

		// 检查设置后的矩阵是否与预期相符
		//if m != tc.expected {
		//	t.Errorf("Set(%d, %d, %v): expected %v, got %v", tc.col, tc.row, tc.value, tc.expected, m)
		//}
		assert.Equal(t, tc.expected, m)
	}
}

func TestMat4Scale(t *testing.T) {
	m := &Mat[float64]{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	scaleFactor := 2.0
	m.Scale(scaleFactor)
	expected := Mat[float64]{{2, 4, 6, 8}, {10, 12, 14, 16}, {18, 20, 22, 24}, {26, 28, 30, 32}}
	assert.Equal(t, expected, *m)
}

func TestMat4Scaled(t *testing.T) {
	m := Mat[float64]{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	scaleFactor := 2.0
	expected := Mat[float64]{{2, 4, 6, 8}, {10, 12, 14, 16}, {18, 20, 22, 24}, {26, 28, 30, 32}}
	result := m.Scaled(scaleFactor)
	assert.Equal(t, expected, result)
}

func TestMat4Mul(t *testing.T) {
	m := &Mat[float64]{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	factor := 2.0
	m.Mul(factor)
	expected := Mat[float64]{{2, 4, 6, 8}, {10, 12, 14, 16}, {18, 20, 22, 24}, {26, 28, 30, 32}}
	assert.Equal(t, expected, *m)
}

func TestMat4Muled(t *testing.T) {
	m := Mat[float64]{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	factor := 2.0
	expected := Mat[float64]{{2, 4, 6, 8}, {10, 12, 14, 16}, {18, 20, 22, 24}, {26, 28, 30, 32}}
	result := m.Muled(factor)
	assert.Equal(t, expected, result)
}

func TestMat4Trace(t *testing.T) {
	m := Mat[float64]{{1, 0, 0, 0}, {0, 2, 0, 0}, {0, 0, 3, 0}, {0, 0, 0, 4}}
	trace := m.Trace()
	assert.Equal(t, 10.0, trace)
}

func TestMat4Trace3(t *testing.T) {
	m := Mat[float64]{{1, 0, 0, 0}, {0, 2, 0, 0}, {0, 0, 3, 0}, {0, 0, 0, 4}}
	trace3 := m.Trace3()
	assert.Equal(t, 6.0, trace3)
}

func TestDeterminant(t *testing.T) {
	ident := Mat[float64]{
		vec4.Vec[float64]{1, 0, 0, 0},
		vec4.Vec[float64]{0, 1, 0, 0},
		vec4.Vec[float64]{0, 0, 1, 0},
		vec4.Vec[float64]{0, 0, 0, 1},
	}
	detId := ident.Determinant()
	if detId != 1 {
		t.Errorf("Wrong determinant for identity matrix: %f", detId)
	}

	detTwo := ident
	detTwo[0][0] = 2
	if det := detTwo.Determinant(); det != 2 {
		t.Errorf("Wrong determinant: %f", det)
	}

	scale2 := ident.Scale(2)
	if det := scale2.Determinant(); det != 2*2*2*1 {
		t.Errorf("Wrong determinant: %f", det)
	}

	row1changed, _ := Parse[float64]("3 0 0 0 2 2 0 0 1 0 2 0 2 0 0 1")
	if det := row1changed.Determinant(); det != 12 {
		t.Errorf("Wrong determinant: %f", det)
	}

	row12changed, _ := Parse[float64]("3 1 0 0 2 5 0 0 1 6 2 0 2 100 0 1")
	if det := row12changed.Determinant(); det != 26 {
		t.Errorf("Wrong determinant: %f", det)
	}

	row123changed, _ := Parse[float64]("3 1 0.5 0 2 5 2 0 1 6 7 0 2 100 1 1")
	if det := row123changed.Determinant3x3(); det != 60.500 {
		t.Errorf("Wrong determinant for 3x3 matrix: %f", det)
	}
	if det := row123changed.Determinant(); det != 60.500 {
		t.Errorf("Wrong determinant: %f", det)
	}
	randomMatrix, err := Parse[float64]("0.43685 0.81673 0.63721 0.23421 0.16600 0.40608 0.53479 0.43210 0.37328 0.36436 0.56356 0.66830 0.32475 0.14294 0.42137 0.98046")
	randomMatrix.Transpose() //transpose for easy comparability with octave output
	//if err != nil {
	//	t.Errorf("Could not parse random matrix: %v", err)
	//}
	assert.Nil(t, err, "Could not parse random matrix: ", err)

	if det := randomMatrix.Determinant3x3(); math.Abs(det-0.043437) > 1e-5 {
		t.Errorf("Wrong determinant for random sub 3x3 matrix: %f", det)
	}

	if det := randomMatrix.Determinant(); math.Abs(det-0.012208) > 1e-5 {
		t.Errorf("Wrong determinant for random matrix: %f", det)
	}
}

func TestMaskedBlock(t *testing.T) {
	m, _ := Parse[float64]("3 1 0.5 0 2 5 2 0 1 6 7 0 2 100 1 1")
	blocked_expected := mat3.Mat[float64]{vec3.Vec[float64]{5, 2, 0}, vec3.Vec[float64]{6, 7, 0}, vec3.Vec[float64]{100, 1, 1}}
	if blocked := m.maskedBlock(0, 0); *blocked != blocked_expected {
		t.Errorf("Did not block 0,0 correctly: %#v", blocked)
	}
}

func TestAdjugate(t *testing.T) {
	adj, _ := Parse[float64]("3 1 0.5 0 2 5 2 0 1 6 7 0 2 100 1 1")
	adj.Adjugate()
	// Computed in octave:
	adj_expected := Mat[float64]{vec4.Vec[float64]{23, -4, -0.5, -0}, vec4.Vec[float64]{-12, 20.5, -5, 0}, vec4.Vec[float64]{7, -17, 13, -0}, vec4.Vec[float64]{1147, -2025, 488, 60.5}}
	if adj != adj_expected {
		t.Errorf("Adjugate not computed correctly: %#v", adj)
	}
}

func TestInvert(t *testing.T) {
	inv, _ := Parse[float64]("3 1 0.5 0 2 5 2 0 1 6 7 0 2 100 1 1")
	inv.Invert()
	// Computed in octave:
	inv_expected := Mat[float64]{vec4.Vec[float64]{0.38016528, -0.0661157, -0.008264462, -0}, vec4.Vec[float64]{-0.19834709, 0.33884296, -0.08264463, 0}, vec4.Vec[float64]{0.11570247, -0.28099173, 0.21487603, -0}, vec4.Vec[float64]{18.958677, -33.471073, 8.066115, 0.99999994}}
	if inv != inv_expected {
		t.Errorf("Inverse not computed correctly: %#v", inv)
	}
}

func TestMultSimpleMatrices(t *testing.T) {
	m1 := Mat[float64]{vec4.Vec[float64]{1, 0, 0, 2},
		vec4.Vec[float64]{0, 1, 2, 0},
		vec4.Vec[float64]{0, 2, 1, 0},
		vec4.Vec[float64]{2, 0, 0, 1}}
	m2 := m1
	var mMult Mat[float64]
	mMult.AssignMul(&m1, &m2)
	t.Log(&m1)
	t.Log(&m2)
	m1.MultMatrix(&m2)
	if m1 != mMult {
		t.Errorf("Multiplication of matrices above failed, expected: \n%v \ngotten: \n%v", &mMult, &m1)
	}
}

func TestMultMatrixVsAssignMul(t *testing.T) {
	m1 := Mat[float64]{
		vec4.Vec[float64]{0.38016528, -0.0661157, -0.008264462, -0},
		vec4.Vec[float64]{-0.19834709, 0.33884296, -0.08264463, 0},
		vec4.Vec[float64]{0.11570247, -0.28099173, 0.21487603, -0},
		vec4.Vec[float64]{18.958677, -33.471073, 8.066115, 0.99999994},
	}
	m2 := Mat[float64]{
		vec4.Vec[float64]{23, -4, -0.5, -0},
		vec4.Vec[float64]{-12, 20.5, -5, 0},
		vec4.Vec[float64]{7, -17, 13, -0},
		vec4.Vec[float64]{1147, -2025, 488, 60.5},
	}
	var mMult Mat[float64]
	mMult.AssignMul(&m1, &m2)
	t.Log(&m1)
	t.Log(&m2)
	m1.MultMatrix(&m2)
	if m1 != mMult {
		t.Errorf("Multiplication of matrices above failed, expected: \n%v \ngotten: \n%v", &mMult, &m1)
	}
}

func BenchmarkAssignMul(b *testing.B) {
	m1 := Mat[float64]{
		vec4.Vec[float64]{0.38016528, -0.0661157, -0.008264462, -0},
		vec4.Vec[float64]{-0.19834709, 0.33884296, -0.08264463, 0},
		vec4.Vec[float64]{0.11570247, -0.28099173, 0.21487603, -0},
		vec4.Vec[float64]{18.958677, -33.471073, 8.066115, 0.99999994},
	}
	m2 := Mat[float64]{
		vec4.Vec[float64]{23, -4, -0.5, -0},
		vec4.Vec[float64]{-12, 20.5, -5, 0},
		vec4.Vec[float64]{7, -17, 13, -0},
		vec4.Vec[float64]{1147, -2025, 488, 60.5},
	}
	var mMult Mat[float64]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mMult.AssignMul(&m1, &m2)
	}
}

func BenchmarkMultMatrix(b *testing.B) {
	m1 := Mat[float64]{
		vec4.Vec[float64]{0.38016528, -0.0661157, -0.008264462, -0},
		vec4.Vec[float64]{-0.19834709, 0.33884296, -0.08264463, 0},
		vec4.Vec[float64]{0.11570247, -0.28099173, 0.21487603, -0},
		vec4.Vec[float64]{18.958677, -33.471073, 8.066115, 0.99999994},
	}
	m2 := Mat[float64]{
		vec4.Vec[float64]{23, -4, -0.5, -0},
		vec4.Vec[float64]{-12, 20.5, -5, 0},
		vec4.Vec[float64]{7, -17, 13, -0},
		vec4.Vec[float64]{1147, -2025, 488, 60.5},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m1.MultMatrix(&m2)
	}
}

func TestMulVec4vsTransformVec4(t *testing.T) {
	m1 := Mat[float64]{
		vec4.Vec[float64]{0.38016528, -0.0661157, -0.008264462, -0},
		vec4.Vec[float64]{-0.19834709, 0.33884296, -0.08264463, 0},
		vec4.Vec[float64]{0.11570247, -0.28099173, 0.21487603, -0},
		vec4.Vec[float64]{18.958677, -33.471073, 8.066115, 0.99999994},
	}

	v := vec4.Vec[float64]{1, 1.5, 2, 2.5}
	v_1 := m1.MulVec4(&v)
	v_2 := m1.MulVec4(&v_1)

	m1.TransformVec4(&v)
	m1.TransformVec4(&v)

	if v_2 != v {
		t.Error(v_2, v)
	}

}

func BenchmarkMulVec4(b *testing.B) {
	m1 := Mat[float64]{
		vec4.Vec[float64]{0.38016528, -0.0661157, -0.008264462, -0},
		vec4.Vec[float64]{-0.19834709, 0.33884296, -0.08264463, 0},
		vec4.Vec[float64]{0.11570247, -0.28099173, 0.21487603, -0},
		vec4.Vec[float64]{18.958677, -33.471073, 8.066115, 0.99999994},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := vec4.Vec[float64]{1, 1.5, 2, 2.5}
		v_1 := m1.MulVec4(&v)
		m1.MulVec4(&v_1)
	}
}

func BenchmarkTransformVec4(b *testing.B) {
	m1 := Mat[float64]{
		vec4.Vec[float64]{0.38016528, -0.0661157, -0.008264462, -0},
		vec4.Vec[float64]{-0.19834709, 0.33884296, -0.08264463, 0},
		vec4.Vec[float64]{0.11570247, -0.28099173, 0.21487603, -0},
		vec4.Vec[float64]{18.958677, -33.471073, 8.066115, 0.99999994},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := vec4.Vec[float64]{1, 1.5, 2, 2.5}
		m1.TransformVec4(&v)
		m1.TransformVec4(&v)
	}
}

func (mat *Mat[T]) TransformVec4PassByValue(v vec4.Vec[T]) (r vec4.Vec[T]) {
	// Use intermediate variables to not alter further computations.
	x := mat[0][0]*v[0] + mat[1][0]*v[1] + mat[2][0]*v[2] + mat[3][0]*v[3]
	y := mat[0][1]*v[0] + mat[1][1]*v[1] + mat[2][1]*v[2] + mat[3][1]*v[3]
	z := mat[0][2]*v[0] + mat[1][2]*v[1] + mat[2][2]*v[2] + mat[3][2]*v[3]
	r[3] = mat[0][3]*v[0] + mat[1][3]*v[1] + mat[2][3]*v[2] + mat[3][3]*v[3]
	r[0] = x
	r[1] = y
	r[2] = z
	return r
}

func Vec4AddPassByValue[T float64 | float32](a, b vec4.Vec[T]) vec4.Vec[T] {
	if a[3] == b[3] {
		return vec4.Vec[T]{a[0] + b[0], a[1] + b[1], a[2] + b[2], 1}
	} else {
		a3 := a.Vec3DividedByW()
		b3 := b.Vec3DividedByW()
		return vec4.Vec[T]{a3[0] + b3[0], a3[1] + b3[1], a3[2] + b3[2], 1}
	}
}

func BenchmarkMulAddVec4_PassByPointer(b *testing.B) {
	m1 := Mat[float64]{
		vec4.Vec[float64]{0.38016528, -0.0661157, -0.008264462, -0},
		vec4.Vec[float64]{-0.19834709, 0.33884296, -0.08264463, 0},
		vec4.Vec[float64]{0.11570247, -0.28099173, 0.21487603, -0},
		vec4.Vec[float64]{18.958677, -33.471073, 8.066115, 0.99999994},
	}
	m2 := Mat[float64]{
		vec4.Vec[float64]{23, -4, -0.5, -0},
		vec4.Vec[float64]{-12, 20.5, -5, 0},
		vec4.Vec[float64]{7, -17, 13, -0},
		vec4.Vec[float64]{1147, -2025, 488, 60.5},
	}

	var v1 vec4.Vec[float64]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := vec4.Vec[float64]{1, 1.5, 2, 2.5}
		m1.TransformVec4(&v)
		m2.TransformVec4(&v)
		v = vec4.Add(&v, &v1)
		v = vec4.Add(&v, &v1)
	}
}

// Demonstrate that
func BenchmarkMulAddVec4_PassByValue(b *testing.B) {
	m1 := Mat[float64]{
		vec4.Vec[float64]{0.38016528, -0.0661157, -0.008264462, -0},
		vec4.Vec[float64]{-0.19834709, 0.33884296, -0.08264463, 0},
		vec4.Vec[float64]{0.11570247, -0.28099173, 0.21487603, -0},
		vec4.Vec[float64]{18.958677, -33.471073, 8.066115, 0.99999994},
	}
	m2 := Mat[float64]{
		vec4.Vec[float64]{23, -4, -0.5, -0},
		vec4.Vec[float64]{-12, 20.5, -5, 0},
		vec4.Vec[float64]{7, -17, 13, -0},
		vec4.Vec[float64]{1147, -2025, 488, 60.5},
	}
	var v1 vec4.Vec[float64]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := vec4.Vec[float64]{1, 1.5, 2, 2.5}
		m1.TransformVec4PassByValue(v)
		m2.TransformVec4PassByValue(v)
		v = Vec4AddPassByValue(v, v1)
		v = Vec4AddPassByValue(v, v1)
	}
}
