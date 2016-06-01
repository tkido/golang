package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(e))
}

/*Sqrt is exproted function */
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	n := x
	for i := 0; i < 5; i++ {
		n = n - ((n*n - x) / (2 * n))
	}
	return n, nil
}

func main() {
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
