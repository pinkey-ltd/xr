package cbezier2

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pinkey.ltd/xr/pkg/go3d/vec2"
)

func TestParse(t *testing.T) {
	s := "1.0 2.0 3.0 4.0 5.0 6.0 7.0 8.0"
	bez, err := Parse[float64](s)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	expected := Bez[float64]{
		P0: vec2.Vec[float64]{1.0, 2.0},
		P1: vec2.Vec[float64]{3.0, 4.0},
		P2: vec2.Vec[float64]{5.0, 6.0},
		P3: vec2.Vec[float64]{7.0, 8.0},
	}

	if bez.P0 != expected.P0 || bez.P1 != expected.P1 || bez.P2 != expected.P2 || bez.P3 != expected.P3 {
		t.Errorf("Parse result incorrect, got: %v, want: %v", bez, expected)
	}
}

func TestString(t *testing.T) {
	bez := Bez[float64]{
		P0: vec2.Vec[float64]{1.0, 2.0},
		P1: vec2.Vec[float64]{3.0, 4.0},
		P2: vec2.Vec[float64]{5.0, 6.0},
		P3: vec2.Vec[float64]{7.0, 8.5},
	}
	expected := "1 2 3 4 5 6 7 8.5"

	assert.Equal(t, expected, bez.String())
}

func TestPoint(t *testing.T) {
	bez := Bez[float64]{
		P0: vec2.Vec[float64]{0.0, 0.0},
		P1: vec2.Vec[float64]{1.0, 1.0},
		P2: vec2.Vec[float64]{2.0, 1.0},
		P3: vec2.Vec[float64]{3.0, 0.0},
	}

	assert.Equal(t, vec2.Vec[float64]{0, 0}, bez.Point(0))
	assert.Equal(t, vec2.Vec[float64]{3, 0}, bez.Point(1))
	assert.Equal(t, vec2.Vec[float64]{1.5, 0.75}, bez.Point(0.5))
	assert.Equal(t, vec2.Vec[float64]{0.75, 0.5625}, bez.Point(0.25))
	assert.Equal(t, vec2.Vec[float64]{2.25, 0.5625}, bez.Point(0.75))
}

func TestTangent(t *testing.T) {
	bez := Bez[float64]{
		P0: vec2.Vec[float64]{0.0, 0.0},
		P1: vec2.Vec[float64]{1.0, 1.0},
		P2: vec2.Vec[float64]{2.0, 1.0},
		P3: vec2.Vec[float64]{3.0, 0.0},
	}

	assert.Equal(t, vec2.Vec[float64]{3, 3}, bez.Tangent(0))
	assert.Equal(t, vec2.Vec[float64]{3, -3}, bez.Tangent(1))
	assert.Equal(t, vec2.Vec[float64]{3, 0}, bez.Tangent(0.5))
	assert.Equal(t, vec2.Vec[float64]{3, 1.5}, bez.Tangent(0.25))
	assert.Equal(t, vec2.Vec[float64]{3, -1.5}, bez.Tangent(0.75))
}

func TestLength(t *testing.T) {
	bez := Bez[float64]{
		P0: vec2.Vec[float64]{0.0, 0.0},
		P1: vec2.Vec[float64]{1.0, 1.0},
		P2: vec2.Vec[float64]{2.0, 2.0},
		P3: vec2.Vec[float64]{3.0, 3.0},
	}

	t.Run("t=0", func(t *testing.T) {
		length := bez.Length(0.0)
		expected := 0.0
		if length != expected {
			t.Errorf("Length at t=0 incorrect, got: %v, want: %v", length, expected)
		}
	})

	t.Run("t=1", func(t *testing.T) {
		length := bez.Length(1.0)
		expected := 4.242640687119285 // Approximate length of the line from (0,0) to (3,3)
		if length != expected {
			t.Errorf("Length at t=1 incorrect, got: %v, want: %v", length, expected)
		}
	})

	t.Run("t=0.5", func(t *testing.T) {
		length := bez.Length(0.5)
		expected := 2.1213203435596424 // Approximate length of the line from (0,0) to (1.5,1.5)
		if length != expected {
			t.Errorf("Length at t=0.5 incorrect, got: %v, want: %v", length, expected)
		}
	})
}
