package main

import "fmt"

func main() {
	value := 1
	doNothing(value)
	fmt.Println(value)
	ptr := &value
	doOp(ptr)
	fmt.Println(ptr, *ptr)
}

func doNothing(v int) {
	v = 23
}

func doOp(v *int) {
	*v = 23
}
