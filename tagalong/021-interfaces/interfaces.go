package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func shapeEfficiency(g geometry) float64 {
	// Area enclosed per unit perimeter used.
	return g.area() / g.perim()
}

func detectCircle(g geometry) (radius float64, isCircle bool) {
	c, isCircle := g.(circle)
	if isCircle {
		return c.radius, true
	}
	return 0, false
}

func main() {
	rect := rectangle{width: 1, height: 4}
	square := rectangle{width: 4, height: 4}
	c := circle{radius: 4}

	fmt.Println("1x4 rectangle efficiency:", shapeEfficiency(rect))
	fmt.Println("square efficiency:", shapeEfficiency(square))
	fmt.Println("circle efficiency:", shapeEfficiency(c))

	fmt.Println(detectCircle(rect))
	fmt.Println(detectCircle(c))
}
