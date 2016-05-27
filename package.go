package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(34750)
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("My favorite number is", rand.Intn(10))
}
