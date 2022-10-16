package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main() {
	fp, _ := os.Open("redshirt.png")
	img, _ := png.Decode(fp)
	fp.Close()
	newshirt, _ := os.Create("newshirt.png")
	defer newshirt.Close()

	// First layer is the color switched shirt.
	h := newHue(.9)
	layer1 := overlay{
		img,
		func(x, y int, c color.Color) color.Color {
			return huedColor{c, &h}
		},
	}

	// Second layer is the logo.
	const scaleDiv = 8 // fraction of image height
	bounds := layer1.Bounds()
	width := bounds.Max.X - bounds.Min.X
	midx := bounds.Min.X + width/2
	height := bounds.Max.Y - bounds.Min.Y
	midy := bounds.Min.Y + height/2

	layer2 := overlay{
		layer1,
		func(x, y int, c color.Color) color.Color {
			if abs(x-midx) > width/scaleDiv || abs(y-midy) > height/scaleDiv {
				return c
			}
			// Insert logo here. For now it's a black box.
			return color.Black
		},
	}

	// Write new image to disk.
	err := png.Encode(newshirt, layer2)
	if err != nil {
		panic(err)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type overlay struct {
	image.Image
	ol func(x, y int, c color.Color) color.Color
}

func (l overlay) At(x, y int) color.Color {
	c := l.Image.At(x, y)
	return l.ol(x, y, c)
}

func newHue(turns float64) hue {
	// http: //www.graficaobscura.com/matrix/index.html
	const sqrt13 = 0.5773502691896257645091488 // sqrt(1/3)
	h := hue{}
	radians := 2 * math.Pi * turns
	sinA, cosA := math.Sincos(radians)
	sqrt13sinA := sqrt13 * sinA
	h[0][0] = float32(cosA + (1.0-cosA)/3.0)
	h[0][1] = float32(1./3.*(1.0-cosA) - sqrt13sinA)
	h[0][2] = float32(1./3.*(1.0-cosA) + sqrt13sinA)
	h[1][0] = float32(1./3.*(1.0-cosA) + sqrt13sinA)
	h[1][1] = float32(cosA + 1./3.*(1.0-cosA))
	h[1][2] = float32(1./3.*(1.0-cosA) - sqrt13sinA)
	h[2][0] = float32(1./3.*(1.0-cosA) - sqrt13sinA)
	h[2][1] = float32(1./3.*(1.0-cosA) + sqrt13sinA)
	h[2][2] = float32(cosA + 1./3.*(1.0-cosA))
	return h
}

type huedColor struct {
	Color color.Color
	h     *hue
}

func (hc huedColor) RGBA() (r, g, b, a uint32) {
	r, g, b, a = hc.Color.RGBA()
	if a != 0 {
		r, g, b = hc.h.apply(r, g, b)
	}
	return r, g, b, a
}

type hue [3][3]float32

func (h *hue) apply(r, g, b uint32) (uint32, uint32, uint32) {
	rf, gf, bf := float32(r), float32(g), float32(b)
	rx := rf*h[0][0] + gf*h[0][1] + bf*h[0][2]
	gx := rf*h[1][0] + gf*h[1][1] + bf*h[1][2]
	bx := rf*h[2][0] + gf*h[2][1] + bf*h[2][2]
	return clamp(rx), clamp(gx), clamp(bx)
}

func clamp(v float32) uint32 {
	if v < 0 {
		return 0
	} else if v > math.MaxUint16 {
		return math.MaxUint16
	}
	return uint32(v + 0.5)
}
