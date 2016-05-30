package main

import "fmt"

func fact(n int) int {
	var rst = 1
	if n > 1 {
		rst = n * fact(n-1)
	}
	return rst
}

func main() {
	fmt.Println(fact(20))
}
