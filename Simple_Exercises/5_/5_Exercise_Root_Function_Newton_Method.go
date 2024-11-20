package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	var z_old float64 // initialize with 0 that is its zero value
	i := 1
	
	for ; math.Abs(z - z_old) > 1e-8; i++{
		z_old = z
		z -= (z * z -x) / (2 * z)
		fmt.Printf("it. n: %d -> actual sqrt = %f\n", i, z)
	}
	
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}