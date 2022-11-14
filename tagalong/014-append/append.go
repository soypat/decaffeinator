package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s)

	// append works on nil slices.
	s = append(s, 0)
	fmt.Println(s)

	// The slice grows as needed.
	s = append(s, 1)
	fmt.Println(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	fmt.Println(s)

	// We can also append a list to a list.
	g := []int{5, 6, 7}
	s = append(s, g...)
	fmt.Println(s)
}
