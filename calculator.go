// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64, c ...float64) float64 {
	total := a + b
	for _, num := range c {
		total += num
	}
	return total
}

// Subtract takes two numbers and returns the result of subtracting the second from the first.
func Subtract(a, b float64, c ...float64) float64 {
	total := a - b
	for _, num := range c {
		total -= num
	}
	return total
}

// Multiply returns first times second
func Multiply(a, b float64, c ...float64) float64 {
	total := a * b
	for _, num := range c {
		total *= num
	}
	return total
}

// Divide returns division result from a to b
func Divide(a, b float64, c ...float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	total := a / b
	for _, num := range c {
		total /= num
	}
	return total, nil
}

// IsPrimo return boolean result for prime numbers
func IsPrimo(number int) (bool, error) {
	isprime := true
	for numLoop := 2; numLoop < number; numLoop++ {
		got, err := Divide(float64(number), float64(numLoop))
		receivedErr := err != nil
		if receivedErr == true {
			return false, errors.New("division error")
		}
		if got == math.Trunc(got) {
			isprime = false
			break
		}
	}
	return isprime, nil
}

// Remainder returns the remainder of division result from a to b
func Remainder(a, b int64) int64 {
	return a % b
}

//Sqrt returns square root of number
func Sqrt(number float64) (float64, error) {
	if number < 0 {
		return 0, errors.New("cannot accept negative numbers on square root")
	}
	sqrtresult := math.Sqrt(float64(number))
	return sqrtresult, nil
}
