package main

import "fmt"

//Vertex is exported type
type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v)

}
