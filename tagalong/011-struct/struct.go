package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v0 := Vertex{}
	fmt.Println(v0)

	v1 := Vertex{X: 1, Y: 2}
	fmt.Println("X:", v1.X)

	v1.X = 1e9
	fmt.Println("new v1:", v1)
}
