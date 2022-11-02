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
		imageH: 2000,
		imageW: 3000,
		xmin:   -2,
		xmax:   1,
		ymin:   -1,
		ymax:   1,
	}

	png.Encode(fp, mandel)
}

var _ image.Image = Mandelbrot{}

type Mandelbrot struct {
	imageH, imageW int
	xmin, xmax     float64
	ymin, ymax     float64
}

func (m Mandelbrot) ColorModel() color.Model {
	return color.GrayModel
}

func (m Mandelbrot) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.imageW, m.imageH)
}

const (
	contrast      = 7
	maxIterations = 255 / 7
)

func (m Mandelbrot) At(i, j int) color.Color {
	// Map pixel position to the complex plane.
	x := float64(i)/float64(m.imageW)*(m.xmax-m.xmin) + m.xmin
	y := float64(j)/float64(m.imageH)*(m.ymax-m.ymin) + m.ymin
	c := complex(x, y)
	iterations := mandelbrotIterations(c)
	if iterations == 255 {
		return color.Black
	}
	// You can go crazy with the colors here.
	return color.Gray{255 - contrast*iterations}
}

// mandelbrotIterations calculates how many iterations of
// the mandelbrot equation z undergoes before diverging.
//
//	v_next = v*v + z
func mandelbrotIterations(z complex128) uint8 {
	var v complex128
	for n := uint8(0); n < maxIterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return n
		}
	}
	return 255
}
