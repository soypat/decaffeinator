// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	fp, _ := os.Create("mandelbrot.png")

	// Keep dimension aspect ratio for best results.
	mandel := Mandelbrot{
		imageH:  2000,
		imageW:  3000,
		xmin:    -2,
		xmax:    1,
		ymin:    -1,
		ymax:    1,
		maxIter: 255 / 7,
	}

	png.Encode(fp, mandel)
}

var _ image.Image = Mandelbrot{}

type Mandelbrot struct {
	imageH, imageW int
	xmin, xmax     float64
	ymin, ymax     float64
	maxIter        int
}

func (m Mandelbrot) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Mandelbrot) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.imageW, m.imageH)
}

func (m Mandelbrot) At(i, j int) color.Color {
	// Map pixel position to the complex plane.
	x := float64(i)/float64(m.imageW)*(m.xmax-m.xmin) + m.xmin
	y := float64(j)/float64(m.imageH)*(m.ymax-m.ymin) + m.ymin
	c := complex(x, y)
	iterations := mandelbrotIterations(c, m.maxIter)
	if iterations == m.maxIter {
		return color.Black
	}
	return ultraColoring(iterations, m.maxIter)
}

// mandelbrotIterations calculates how many iterations of
// the mandelbrot equation z undergoes before diverging.
//
//	v_next = v*v + z
func mandelbrotIterations(z complex128, maxIter int) int {
	var v complex128
	for n := 0; n < maxIter; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return n
		}
	}
	return maxIter
}

// matrixColoring replicates 1999 film's color palette.
func matrixColoring(iter, maxIter int) color.RGBA {
	// Roughly corresponds to distance to edge of mandelbrot set.
	dist := 256 * iter / maxIter
	if dist < 128 {
		// Far away color quickly fades away.
		return color.RGBA{R: uint8(dist), G: uint8(dist * 2), B: uint8(dist), A: 255}
	} else if dist > 255 {
		dist = 255
	}
	// Close to mandelbrot set color changes from green to white.
	return color.RGBA{R: uint8(dist), G: 255, B: uint8(dist), A: 255}
}

func grayscaleColoring(iter, maxIter int) color.Gray {
	return color.Gray{Y: uint8(256 * iter / maxIter)}
}

func ultraColoring(iter, maxiter int) color.RGBA {
	if iter < maxiter {
		return ultraColor[iter%len(ultraColor)]
	}
	return color.RGBA{} // Black.
}

// https://stackoverflow.com/questions/16500656/which-color-gradient-is-used-to-color-mandelbrot-in-wikipedia
var ultraColor = [...]color.RGBA{
	{R: 66, G: 30, B: 15, A: 255},    // brown 3
	{R: 25, G: 7, B: 26, A: 255},     // dark violett
	{R: 9, G: 1, B: 47, A: 255},      // darkest blue
	{R: 4, G: 4, B: 73, A: 255},      // blue 5
	{R: 0, G: 7, B: 100, A: 255},     // blue 4
	{R: 12, G: 44, B: 138, A: 255},   // blue 3
	{R: 24, G: 82, B: 177, A: 255},   // blue 2
	{R: 57, G: 125, B: 209, A: 255},  // blue 1
	{R: 134, G: 181, B: 229, A: 255}, // blue 0
	{R: 211, G: 236, B: 248, A: 255}, // lightest blue
	{R: 241, G: 233, B: 191, A: 255}, // lightest yellow
	{R: 248, G: 201, B: 95, A: 255},  // light yellow
	{R: 255, G: 170, B: 0, A: 255},   // dirty yellow
	{R: 204, G: 128, B: 0, A: 255},   // brown 0
	{R: 153, G: 87, B: 0, A: 255},    // brown 1
	{R: 106, G: 52, B: 3, A: 255},    // brown 2
}
