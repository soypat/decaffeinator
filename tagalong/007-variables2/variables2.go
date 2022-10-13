package main

import (
	"fmt"
	"strings"
)

var c, python, java bool

var (
	s           int64
	s2          string  = "This is long text"
	start, stop float64 = 1, 20
)

func main() {
	afloat := 6.02
	casted := int(afloat)
	words := strings.Split(s2, " ")
	fmt.Println(c, python, java, s, words, start, stop, afloat, casted)
}
