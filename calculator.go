// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply returns first times second
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide returns division result from a to b
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division by zero")
	}
	return a / b, nil
}

// IsPrimo return boolean for prime numbers
func IsPrimo(number int) (bool, error) {
	var isprime bool = true
	for numLoop := 2; numLoop < number; numLoop++ {
		got, err := Divide(float64(number), float64(numLoop))
		receivedErr := err != nil
		if receivedErr == true {
			return false, errors.New("Unexpected error")
		}
		if got == math.Trunc(got) {
			isprime = false
			break
		}
	}
	return isprime, nil
}

// Remainder returns the remainder of division result from a to b
func Remainder(a, b int) int {
	return a % b
}
