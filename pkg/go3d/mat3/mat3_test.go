package mat3

import (
	"math"
	"testing"

	"github.com/flywave/go3d/mat2"
	"github.com/flywave/go3d/vec2"
	"github.com/stretchr/testify/assert"
	"pinkey.ltd/xr/pkg/go3d/vec3"
)

func TestMat3From(t *testing.T) {
	// 创建一个 3x3 矩阵
	mat := Mat[float64]{
		vec3.Vec[float64]{1, 2, 3},
		vec3.Vec[float64]{4, 5, 6},
		vec3.Vec[float64]{7, 8, 9},
	}

	// 使用 From 函数复制矩阵
	copiedMat := From(&mat)

	// 检查复制的矩阵是否与原始矩阵相同
	assert.Equal(t, mat, copiedMat, "From should correctly copy the matrix")
}

func TestMat3Scale(t *testing.T) {
	mat := Mat[float64]{
		vec3.Vec[float64]{1, 0, 0},
		vec3.Vec[float64]{0, 1, 0},
		vec3.Vec[float64]{0, 0, 1},
	}

	// 缩放矩阵
	mat.Scale(2.0)

	// 检查缩放后的矩阵
	expected := Mat[float64]{
		vec3.Vec[float64]{2, 0, 0},
		vec3.Vec[float64]{0, 2, 0},
		vec3.Vec[float64]{0, 0, 2},
	}

	assert.Equal(t, expected, mat, "Scale should correctly scale the matrix")
}

func TestMat3Determinant(t *testing.T) {
	mat := Mat[float64]{
		vec3.Vec[float64]{1, 2, 3},
		vec3.Vec[float64]{4, 5, 6},
		vec3.Vec[float64]{7, 8, 9},
	}

	det := mat.Determinant()

	// 计算预期的行列式值
	expected := 1.0*(5.0*9.0-6.0*8.0) - 2.0*(4.0*9.0-6.0*7.0) + 3.0*(4.0*8.0-5.0*7.0)

	assert.Equal(t, expected, det, "Determinant should correctly calculate the determinant")
}

func TestMat3Transpose(t *testing.T) {
	mat := Mat[float64]{
		vec3.Vec[float64]{1, 2, 3},
		vec3.Vec[float64]{4, 5, 6},
		vec3.Vec[float64]{7, 8, 9},
	}

	mat.Transpose()

	expected := Mat[float64]{
		vec3.Vec[float64]{1, 4, 7},
		vec3.Vec[float64]{2, 5, 8},
		vec3.Vec[float64]{3, 6, 9},
	}

	assert.Equal(t, expected, mat, "Transpose should correctly transpose the matrix")
}

func TestMat3Quaternion(t *testing.T) {
	mat := Mat[float64]{
		vec3.Vec[float64]{1, 0, 0},
		vec3.Vec[float64]{0, 1, 0},
		vec3.Vec[float64]{0, 0, 1},
	}

	q := mat.Quaternion()

	// 检查四元数的值
	expected := vec3.Vec[float64]{0, 0, 0}

	assert.Equal(t, expected, q, "Quaternion should correctly extract the quaternion from the matrix")
}

func TestMat3AssignXRotation(t *testing.T) {
	mat := Mat[float64]{}
	angle := math.Pi / 2 // 90 degrees

	mat.AssignXRotation(angle)

	expected := Mat[float64]{
		vec3.Vec[float64]{1, 0, 0},
		vec3.Vec[float64]{0, 0, -1},
		vec3.Vec[float64]{0, 1, 0},
	}

	assert.Equal(t, expected, mat, "AssignXRotation should correctly assign the X rotation to the matrix")
}

func TestMat3AssignYRotation(t *testing.T) {
	mat := Mat[float64]{}
	angle := math.Pi / 2 // 90 degrees

	mat.AssignYRotation(angle)

	expected := Mat[float64]{
		vec3.Vec[float64]{0, 0, 1},
		vec3.Vec[float64]{0, 1, 0},
		vec3.Vec[float64]{-1, 0, 0},
	}

	assert.Equal(t, expected, mat, "AssignYRotation should correctly assign the Y rotation to the matrix")
}

func TestMat3AssignZRotation(t *testing.T) {
	mat := Mat[float64]{}
	angle := math.Pi / 2 // 90 degrees

	mat.AssignZRotation(angle)

	expected := Mat[float64]{
		vec3.Vec[float64]{0, -1, 0},
		vec3.Vec[float64]{1, 0, 0},
		vec3.Vec[float64]{0, 0, 1},
	}

	assert.Equal(t, expected, mat, "AssignZRotation should correctly assign the Z rotation to the matrix")
}

func TestMat3AssignCoordinateSystem(t *testing.T) {
	mat := Mat[float64]{}
	x := &vec3.Vec[float64]{1, 0, 0}
	y := &vec3.Vec[float64]{0, 1, 0}
	z := &vec3.Vec[float64]{0, 0, 1}

	mat.AssignCoordinateSystem(x, y, z)

	expected := Mat[float64]{
		vec3.Vec[float64]{1, 0, 0},
		vec3.Vec[float64]{0, 1, 0},
		vec3.Vec[float64]{0, 0, 1},
	}

	assert.Equal(t, expected, mat, "AssignCoordinateSystem should correctly assign the coordinate system to the matrix")
}

func TestMat3AssignEulerRotation(t *testing.T) {
	mat := Mat[float64]{}
	yHead := math.Pi / 2  // 90 degrees
	xPitch := math.Pi / 4 // 45 degrees
	zRoll := math.Pi / 6  // 30 degrees

	mat.AssignEulerRotation(yHead, xPitch, zRoll)

	// 预期的旋转矩阵
	expected := Mat[float64]{
		vec3.Vec[float64]{0.43301270189221935, -0.75, 0.5},
		vec3.Vec[float64]{0.8660254037844387, 0.43301270189221935, -0.25},
		vec3.Vec[float64]{-0.25, 0.5, 0.8660254037844387},
	}

	assert.Equal(t, expected, mat, "AssignEulerRotation should correctly assign the Euler rotation to the matrix")
}

func TestMat3ExtractEulerAngles(t *testing.T) {
	mat := Mat[float64]{
		vec3.Vec[float64]{0.43301270189221935, -0.75, 0.5},
		vec3.Vec[float64]{0.8660254037844387, 0.43301270189221935, -0.25},
		vec3.Vec[float64]{-0.25, 0.5, 0.8660254037844387},
	}

	yHead, xPitch, zRoll := mat.ExtractEulerAngles()

	// 预期的欧拉角
	expectedYHead := math.Pi / 2
	expectedXPitch := math.Pi / 4
	expectedZRoll := math.Pi / 6

	assert.InDelta(t, expectedYHead, float64(yHead), 1e-6, "ExtractEulerAngles should correctly extract the yHead")
	assert.InDelta(t, expectedXPitch, float64(xPitch), 1e-6, "ExtractEulerAngles should correctly extract the xPitch")
	assert.InDelta(t, expectedZRoll, float64(zRoll), 1e-6, "ExtractEulerAngles should correctly extract the zRoll")
}

func TestMat3IsReflective(t *testing.T) {
	mat := Mat[float64]{
		vec3.Vec[float64]{1, 0, 0},
		vec3.Vec[float64]{0, 1, 0},
		vec3.Vec[float64]{0, 0, -1},
	}

	assert.True(t, mat.IsReflective(), "IsReflective should return true for a reflective matrix")
}

func Test_ColsAndRows(t *testing.T) {
	A := Mat[float64]{
		vec3.Vec[float64]{23, -4, -0.5},
		vec3.Vec[float64]{-12, 20.5, -5},
		vec3.Vec[float64]{7, -17, 13},
	}

	a11 := A.Get(0, 0)
	a21 := A.Get(0, 1)
	a31 := A.Get(0, 2)

	a12 := A.Get(1, 0)
	a22 := A.Get(1, 1)
	a32 := A.Get(1, 2)

	a13 := A.Get(2, 0)
	a23 := A.Get(2, 1)
	a33 := A.Get(2, 2)

	assert.True(t, a11 == 23 && a21 == -4 && a31 == -0.5 &&
		a12 == -12 && a22 == 20.5 && a32 == -5 &&
		a13 == 7 && a23 == -17 && a33 == 13, "matrix ill referenced")
}

func TestT_Transposed(t *testing.T) {
	matrix := Mat[float64]{
		vec3.Vec[float64]{1, 2, 3},
		vec3.Vec[float64]{4, 5, 6},
		vec3.Vec[float64]{7, 8, 9},
	}
	expectedMatrix := Mat[float64]{
		vec3.Vec[float64]{1, 4, 7},
		vec3.Vec[float64]{2, 5, 8},
		vec3.Vec[float64]{3, 6, 9},
	}

	transposedMatrix := matrix.Transposed()

	if transposedMatrix != expectedMatrix {
		t.Errorf("matrix transposed wrong: %v --> %v", matrix, transposedMatrix)
	}
}

func TestT_Transpose(t *testing.T) {
	matrix := Mat[float64]{
		vec3.Vec[float64]{10, 20, 30},
		vec3.Vec[float64]{40, 50, 60},
		vec3.Vec[float64]{70, 80, 90},
	}

	expectedMatrix := Mat[float64]{
		vec3.Vec[float64]{10, 40, 70},
		vec3.Vec[float64]{20, 50, 80},
		vec3.Vec[float64]{30, 60, 90},
	}

	transposedMatrix := matrix
	transposedMatrix.Transpose()

	if transposedMatrix != expectedMatrix {
		t.Errorf("matrix transposed wrong: %v --> %v", matrix, transposedMatrix)
	}
}

func TestDeterminant_2(t *testing.T) {
	detTwo := Ident
	detTwo[0][0] = 2
	if det := detTwo.Determinant(); det != 2 {
		t.Errorf("Wrong determinant: %f", det)
	}
}

func TestDeterminant_3(t *testing.T) {
	scale2 := Ident.Scaled(2)
	if det := scale2.Determinant(); det != 2*2*2*1 {
		t.Errorf("Wrong determinant: %f", det)
	}
}

func TestDeterminant_4(t *testing.T) {
	row1changed, _ := Parse("3 0 0   2 2 0   1 0 2")
	if det := row1changed.Determinant(); det != 12 {
		t.Errorf("Wrong determinant: %f", det)
	}
}

func TestDeterminant_5(t *testing.T) {
	row12changed, _ := Parse("3 1 0   2 5 0   1 6 2")
	if det := row12changed.Determinant(); det != 26 {
		t.Errorf("Wrong determinant: %f", det)
	}
}

func TestDeterminant_7(t *testing.T) {
	randomMatrix, err := Parse("0.43685 0.81673 0.63721   0.16600 0.40608 0.53479   0.37328 0.36436 0.56356")
	randomMatrix.Transpose()
	if err != nil {
		t.Errorf("Could not parse random matrix: %v", err)
	}
	if det := randomMatrix.Determinant(); PracticallyEquals(det, 0.043437) {
		t.Errorf("Wrong determinant for random sub 3x3 matrix: %f", det)
	}
}

func PracticallyEquals(value1 float64, value2 float64) bool {
	return math.Abs(value1-value2) > EPSILON
}

func TestDeterminant_6(t *testing.T) {
	row123changed := row123Changed
	if det := row123changed.Determinant(); det != 60.500 {
		t.Errorf("Wrong determinant for 3x3 matrix: %f", det)
	}
}

func TestDeterminant_1(t *testing.T) {
	detId := Ident.Determinant()
	if detId != 1 {
		t.Errorf("Wrong determinant for identity matrix: %f", detId)
	}
}

func TestMaskedBlock(t *testing.T) {
	m := row123Changed
	blockedExpected := mat2.T{vec2.T{5, 2}, vec2.T{6, 7}}
	if blocked := m.maskedBlock(0, 0); *blocked != blockedExpected {
		t.Errorf("Did not block 0,0 correctly: %#v", blocked)
	}
}

func TestAdjugate(t *testing.T) {
	adj := row123Changed

	// Computed in octave:
	adjExpected := T{
		vec3.T{23, -4, -0.5},
		vec3.T{-12, 20.5, -5},
		vec3.T{7, -17, 13},
	}

	adj.Adjugate()

	if adj != adjExpected {
		t.Errorf("Adjugate not computed correctly: %#v", adj)
	}
}

func TestAdjugated(t *testing.T) {
	sqrt2 := math.Sqrt(2)
	A := T{
		vec3.T{1, 0, -1},
		vec3.T{0, sqrt2, 0},
		vec3.T{1, 0, 1},
	}

	expectedAdjugated := T{
		vec3.T{1.4142135623730951, -0, 1.4142135623730951},
		vec3.T{-0, 2, -0},
		vec3.T{-1.4142135623730951, -0, 1.4142135623730951},
	}

	adjA := A.Adjugated()

	if adjA != expectedAdjugated {
		t.Errorf("Adjugate not computed correctly: %v", adjA)
	}
}

func TestInvert_ok(t *testing.T) {
	inv := invertableMatrix1
	_, err := inv.Invert()

	if err != nil {
		t.Error("Inverse not computed correctly", err)
	}

	invExpected := invertedMatrix1
	if inv != invExpected {
		t.Errorf("Inverse not computed correctly: %#v", inv)
	}
}

func TestInvert_ok2(t *testing.T) {
	sqrt2 := math.Sqrt(2)
	A := T{
		vec3.T{1, 0, -1},
		vec3.T{0, sqrt2, 0},
		vec3.T{1, 0, 1},
	}

	expectedInverted := T{
		vec3.T{0.5, 0, 0.5},
		vec3.T{0, 0.7071067811865475, 0},
		vec3.T{-0.5, 0, 0.5},
	}

	invA, err := A.Inverted()
	if err != nil {
		t.Error("Inverse not computed correctly", err)
	}

	if invA != expectedInverted {
		t.Errorf("Inverse not computed correctly: %v", A)
	}
}

func TestInvert_nok_1(t *testing.T) {
	inv := nonInvertableMatrix1
	_, err := inv.Inverted()
	if err == nil {
		t.Error("Inverse should not be possible", err)
	}
}

func TestInvert_nok_2(t *testing.T) {
	inv := nonInvertableMatrix2
	_, err := inv.Inverted()
	if err == nil {
		t.Error("Inverse should not be possible", err)
	}
}

func BenchmarkAssignMul(b *testing.B) {
	m1 := testMatrix1
	m2 := testMatrix2
	var mMult T
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mMult.AssignMul(&m1, &m2)
	}
}
