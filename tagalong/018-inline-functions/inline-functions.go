package main

import "fmt"

func main() {
	a := 287117
	notSoRand := func() int {
		a = a * 7
		return a
	}
	superrand1 := SuperRandom(notSoRand)
	superrand2 := SuperRandom(notSoRand)
	superrand3 := SuperRandom(notSoRand)

	fmt.Println(superrand1, superrand2, superrand3)
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
