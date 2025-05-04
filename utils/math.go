// Package utils provides utility functions for the xr project.
package utils

import "math"

// IsEven checks if a number is even.
func IsEven(n int) bool {
	return n%2 == 0
}

// IsPrime checks if a number is prime.
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if IsEven(n) || n%3 == 0 {
		return false
	}

	// Check for divisibility by numbers of the form 6k±1 up to sqrt(n)
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

// Factorial calculates the factorial of a number.
func Factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// Round rounds a float64 to the nearest integer.
func Round(x float64) int {
	return int(math.Round(x))
}
