package main

import (
	"fmt"
	"time"
)

func main() {
	// it creates a stack of call and then before the return it is printed at the start the interation 10-th, then the 9-th and so on
	for i:=0; i<=10; i++ {
		defer fmt.Printf("The time at iteration %d is %s\n", i, time.Now())
	}
}