package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEven(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Zero", 0, true},
		{"Positive Even", 2, true},
		{"Positive Odd", 3, false},
		{"Negative Even", -4, true},
		{"Negative Odd", -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEven(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Negative", -1, false},
		{"Zero", 0, false},
		{"One", 1, false},
		{"Two", 2, true},
		{"Three", 3, true},
		{"Four", 4, false},
		{"Five", 5, true},
		{"Six", 6, false},
		{"Seven", 7, true},
		{"Nine", 9, false},
		{"Eleven", 11, true},
		{"Twenty Five", 25, false},
		{"Large Prime", 97, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPrime(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"Negative", -1, 1},
		{"Zero", 0, 1},
		{"One", 1, 1},
		{"Two", 2, 2},
		{"Three", 3, 6},
		{"Four", 4, 24},
		{"Five", 5, 120},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Factorial(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected int
	}{
		{"Integer", 5.0, 5},
		{"Round Up", 5.6, 6},
		{"Round Down", 5.4, 5},
		{"Round Half", 5.5, 6},
		{"Negative Round Up", -5.6, -6},
		{"Negative Round Down", -5.4, -5},
		{"Negative Round Half", -5.5, -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Round(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
