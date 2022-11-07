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
)

const (
	imageSize = 1000
)

func main() {
	s := SDFDrawing{}
	seed := time.Now().Unix() % 1000
	rand.Seed(seed)
	for i := 0; i < 8; i++ {
		sdf := TriangleSDF(randvec2(imageSize), randvec2(imageSize), randvec2(imageSize))
		s.sdfs = append(s.sdfs, sdf)
	}
	fmt.Println("creating sdf.png with seed", seed)
	fp, _ := os.Create("sdf.png")
	png.Encode(fp, s)
}

type SDFDrawing struct {
	sdfs []func(vec2) float64
}

func (s SDFDrawing) At(i, j int) color.Color {
	v := vec2{x: float64(i), y: float64(j)}
	for i := range s.sdfs {
		if value := s.sdfs[i](v); value < 0 {
			g := uint8(i) % 3
			b := uint8(i) % 4
			return color.RGBA{R: uint8(0xff + int(value)), G: g * 64, B: b * 64, A: 0xff}
		}
	}
	return color.Black
}

func (s SDFDrawing) Bounds() image.Rectangle { return image.Rect(0, 0, imageSize, imageSize) }

func (s SDFDrawing) ColorModel() color.Model { return color.RGBAModel }

type vec2 struct{ x, y float64 }

// TriangleSDF returns the signed distance function for a 2D triangle with vertices at the given points.
// MIT license. Copyright © 2014 Inigo Quilez, https://www.shadertoy.com/view/XsXSz4
func TriangleSDF(p0, p1, p2 vec2) func(vec2) float64 {
	return func(p vec2) float64 {
		e0 := sub2(p1, p0)
		e1 := sub2(p2, p1)
		e2 := sub2(p0, p2)

		v0 := sub2(p, p0)
		v1 := sub2(p, p1)
		v2 := sub2(p, p2)

		s0 := uclamp(dot2(v0, e0) / dot2(e0, e0))
		pq0 := sub2(v0, scale2(s0, e0))
		s1 := uclamp(dot2(v1, e1) / dot2(e1, e1))
		pq1 := sub2(v1, scale2(s1, e1))
		s2 := uclamp(dot2(v2, e2) / dot2(e2, e2))
		pq2 := sub2(v2, scale2(s2, e2))

		s := e0.x*e2.y - e0.y*e2.x
		distance := math.Min(dot2(pq0, pq0), math.Min(dot2(pq1, pq1), dot2(pq2, pq2)))
		sign := math.Min(s*(v0.x*e0.y-v0.y*e0.x), math.Min(s*(v1.x*e1.y-v1.y*e1.x), s*(v2.x*e2.y-v2.y*e2.x)))
		return math.Copysign(math.Sqrt(distance), -sign)
	}
}

func sub2(a, b vec2) vec2           { return vec2{a.x - b.x, a.y - b.y} }
func dot2(a, b vec2) float64        { return a.x*b.x + a.y*b.y }
func uclamp(x float64) float64      { return math.Max(0, math.Min(1, x)) }
func scale2(f float64, v vec2) vec2 { return vec2{f * v.x, f * v.y} }
func randvec2(f float64) vec2       { return vec2{f * rand.Float64(), f * rand.Float64()} }
