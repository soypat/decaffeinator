package main

import "fmt"

func main() {
	a := 1
	notSoRand := func() int {
		a = a * 7
		return a
	}
	call1 := notSoRand()
	call2 := notSoRand()
	fmt.Println("calling function yields different results:", call1, call2)

	superrand1 := SuperRandom(notSoRand)
	superrand2 := SuperRandom(staticRandom)
	fmt.Println("a function can take another function as argument:", superrand1, superrand2)
}

// SuperRandom returns a super random number using a
// not so random source function "normalRandom".
// (this is not actually more random!)
func SuperRandom(normalRandom func() int) int {
	superrand := 12345678
	for i := 0; i < 3; i++ {
		rand := normalRandom()
		superrand = superrand*7 + rand*31
	}
	return superrand
}

func staticRandom() int {
	return 4
}
