package main

import "fmt"

func collatz(a int) (down int, up int) {
	down = a / 2
	up = a*3 + 1
	return down, up
}

func main() {
	const v = 60
	down, up := collatz(v)
	fmt.Printf("empezando en %d hay que saber subir %d, y bajar %d", v, up, down)
}
