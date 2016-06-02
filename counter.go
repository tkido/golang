package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func fibonacci() func() int {
	x, y := 1, 0
	return func() int {
		x, y = y, x+y
		return x
	}
}

func main() {
	c := counter()
	for i := 0; i < 10; i++ {
		fmt.Println(c())
	}

	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
