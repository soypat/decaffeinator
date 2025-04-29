package main

import "fmt"

func append3(s []int) {
	s = append(s, 3)
}

func main() {
	l := []int{0, 1}
	append3(l)
	fmt.Println(l)

	_ = append(l, 3) // result discarded.
	fmt.Println(l)
}
