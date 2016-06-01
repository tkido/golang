package main

import (
	"fmt"
	"math"
)

/*Sqrt is exproted function */
func Sqrt(x float64) float64 {
	n := x
	for i := 0; i < 5; i++ {
		n = n - ((n*n - x) / (2 * n))
	}
	return n
}

func main() {
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(2))
}
