package main
import "fmt"
import "math"

func trial(x, n, lim float64) (val float64) {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
		return -v
	}
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	/* for {
		fmt.Println("sono nel looooooop...")
	} */
	
	val := trial(2, 4, 10); if val != 0 {
		fmt.Printf("The value v is %g\n", val)
	}

	
}