package main

import (
	"fmt"
	"math"
)

// Define the custom error type as instructed
type ErrNegativeSqrt float64

// Implement the Error() method for ErrNegativeSqrt
func (e ErrNegativeSqrt) Error() string {
	// Convert e to float64 to avoid infinite recursion in fmt.Sprintf
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt function modified to return an ErrNegativeSqrt when x is negative
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x) // Return ErrNegativeSqrt type
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
