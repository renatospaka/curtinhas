package main

import (
	"errors"
	"fmt"
)

// Common errors
var (
	ErrNegSqrt    = errors.New("sqrt of negative number")
	ErrNoSolution = errors.New("no solution found")
)

func Abs(val float64) float64 {
	if val < 0 {
		return -val
	}
	return val
}

func Sqrt(val float64) (float64, error) {
	if val < 0.0 {
		return 0.0, ErrNegSqrt
	}

	if val == 0.0 {
		return 0.0, nil
	}

	guess, epsilon := 1.0, 0.00001
	for i := 0; i < 10000; i++ {
		if Abs(guess * guess - val) <= epsilon {
			return guess, nil
		}
		guess = (val / guess + guess) / 2.0
	}
	return 0.0, ErrNoSolution
}

func main() {
	for i := 2; i <= 15; i++ {
		n := i * i
		sq, _ := Sqrt(float64(n))
		fmt.Printf("the sqrt of %d is %f\n", n, sq)
	}
}
