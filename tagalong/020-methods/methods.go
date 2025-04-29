package main

import "fmt"

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

func (r *rectangle) scale(scale float64) {
	r.height *= scale
	r.width *= scale
}

func main() {
	r := rectangle{width: 12.7, height: 10}
	fmt.Println("area [mm²]:", r.area())
	fmt.Println("perimeter [mm]:", r.perim())

	// Convert to inches.
	r.scale(1. / 25.4)
	fmt.Println("area[inches²]:", r.area())
	fmt.Println("perim[inches]:", r.perim())
}
