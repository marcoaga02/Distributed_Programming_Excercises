package main

import "fmt"

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func test2(a,b float64) float64{
	return a/b
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	fmt.Println("# FIRST PART")
	val := compute(test2)
	fmt.Println(val)
	fmt.Println("# SECOND PART")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}