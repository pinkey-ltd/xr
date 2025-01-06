package mat4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMat4_String(t *testing.T) {
	m := Mat[float64]{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	expected := "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16"
	assert.Equal(t, expected, m.String())
}

func TestMat4_RowsColsSize(t *testing.T) {
	m := Mat[float64]{}
	assert.Equal(t, 4, m.Rows())
	assert.Equal(t, 4, m.Cols())
	assert.Equal(t, 16, m.Size())
}

func TestMat4_IsZero(t *testing.T) {
	var zero Mat[float64]
	nonZero := Mat[float64]{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}
	assert.True(t, zero.IsZero())
	assert.False(t, nonZero.IsZero())
}

func TestMat4_Scale(t *testing.T) {
	m := &Mat[float64]{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	scaleFactor := 2.0
	m.Scale(scaleFactor)
	expected := Mat[float64]{{2, 4, 6, 8}, {10, 12, 14, 16}, {18, 20, 22, 24}, {26, 28, 30, 32}}
	assert.Equal(t, expected, *m)
}

func TestMat4_Scaled(t *testing.T) {
	m := Mat[float64]{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	scaleFactor := 2.0
	expected := Mat[float64]{{2, 4, 6, 8}, {10, 12, 14, 16}, {18, 20, 22, 24}, {26, 28, 30, 32}}
	result := m.Scaled(scaleFactor)
	assert.Equal(t, expected, result)
}

func TestMat4_Mul(t *testing.T) {
	m := &Mat[float64]{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	factor := 2.0
	m.Mul(factor)
	expected := Mat[float64]{{2, 4, 6, 8}, {10, 12, 14, 16}, {18, 20, 22, 24}, {26, 28, 30, 32}}
	assert.Equal(t, expected, *m)
}

func TestMat4_Muled(t *testing.T) {
	m := Mat[float64]{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	factor := 2.0
	expected := Mat[float64]{{2, 4, 6, 8}, {10, 12, 14, 16}, {18, 20, 22, 24}, {26, 28, 30, 32}}
	result := m.Muled(factor)
	assert.Equal(t, expected, result)
}

func TestMat4_Trace(t *testing.T) {
	m := Mat[float64]{{1, 0, 0, 0}, {0, 2, 0, 0}, {0, 0, 3, 0}, {0, 0, 0, 4}}
	trace := m.Trace()
	assert.Equal(t, 10.0, trace)
}

func TestMat4_Trace3(t *testing.T) {
	m := Mat[float64]{{1, 0, 0, 0}, {0, 2, 0, 0}, {0, 0, 3, 0}, {0, 0, 0, 4}}
	trace3 := m.Trace3()
	assert.Equal(t, 6.0, trace3)
}

// Additional tests for remaining methods can be added following similar patterns.
