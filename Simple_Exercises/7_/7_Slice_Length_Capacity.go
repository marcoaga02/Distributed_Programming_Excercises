package main

import(
	"fmt"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Printf("The lenght of the slice 's' is %d and the capacity is %d\n", len(s), cap(s))
	slice := make([]int, 5, 5)
	fmt.Printf("The lenght of the slice 'slice' is %d and the capacity is %d\n", len(slice), cap(slice))
	for i:=0; i < len(slice); i++ {
		slice[i] = i*i;
	}
	slice = append(slice, slice[len(slice) - 1] + 1);
	fmt.Printf("The lenght of the slice 'slice' after append is %d and the capacity is %d\n", len(slice), cap(slice))
	for i:=0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	var pow[3] int
	pow[0] = 1
	pow[1] = 2
	pow[2] = 4
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

}