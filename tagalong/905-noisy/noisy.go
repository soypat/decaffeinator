package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"

	noise "github.com/soypat/decaffeinator/tagalong/pkg-noise"
)

const (
	imageSize = 1000
	perlinDim = 10
)

func main() {
	seed := time.Now().Unix() % 1000
	rand.Seed(seed)
	p := Noisy{offset: rand.Float64() * 20}
	fp, _ := os.Create("noisy.png")
	fmt.Println("creating noisy.png with seed", seed)
	png.Encode(fp, p)
}

type Noisy struct {
	offset float64
}

func (p Noisy) At(i, j int) color.Color {
	const maxNoise = 1
	const span = 100
	x, y := float64(i)/imageSize, float64(j)/imageSize
	x += p.offset
	y += p.offset
	n := math.Max(0, (noise.Simplex3D(x*span, y*span, 5)+maxNoise)/(maxNoise*2))
	return color.RGBA{R: uint8(n * 255), A: 255}
}

func (p Noisy) Bounds() image.Rectangle { return image.Rect(0, 0, imageSize, imageSize) }

func (p Noisy) ColorModel() color.Model { return color.RGBAModel }
