package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"time"
)

const (
	imageSize  = 200
	stepLength = 1.0
	iterations = 10
)

func main() {
	start := time.Now()
	rect := image.Rect(0, 0, imageSize, imageSize)
	base := image.NewGray(rect)
	for i := range base.Pix {
		base.Pix[i] = 0xff // set to white.
	}
	s := Spirograph{}
	for it := 0; it < iterations; it++ {
		s.SetTime(float64(it) * stepLength)
		for ix := 0; ix < imageSize; ix++ {
			x := float64(2*ix-imageSize) / imageSize
			for iy := 0; iy < imageSize; iy++ {
				y := float64(2*iy-imageSize) / imageSize
				actual := s.at(x, y)
				previousGray := base.GrayAt(ix, iy)
				previous := float64(previousGray.Y) / 255
				newColor := mix(math.Min(previous, actual), 1, 0.01)
				base.SetGray(ix, iy, color.Gray{Y: uint8(newColor * 255)})
			}
		}
	}
	fp, _ := os.Create("spiro.png")
	err := png.Encode(fp, base)
	if err != nil {
		panic(err)
	}
	fmt.Println("elapsed: ", time.Since(start).String())
}

// https://www.shadertoy.com/view/XlfGzX
type Spirograph struct {
	Time  float64
	Scale float64
}

type vec2 struct{ x, y float64 }

func (s Spirograph) F(t float64) vec2 {
	var q vec2
	r := 1.0
	for j := 0; j < 9; j++ {
		sin, cos := math.Sincos(t)
		q = add2(q, scale2(r, vec2{sin, cos}))
		t *= s.Scale
		r /= math.Abs(s.Scale)
	}
	return q
}

// DF is a gradient calculation. Is used to slow down the stepping.
func (s Spirograph) DF(p vec2, t float64) (df, dt float64) {
	pf := sub2(p, s.F(t))
	d1 := math.Hypot(pf.x, pf.y)
	dt = 0.1 * d1
	pfdt := sub2(p, s.F(t+dt))
	d2 := math.Hypot(pfdt.x, pfdt.y)
	dt /= math.Max(dt, d1-d2)
	return math.Min(d1, d2), 0.4 * math.Log(d1*dt+1)
}

func (s Spirograph) at(x, y float64) float64 {
	p := vec2{x, y}
	p = scale2(1.75, p)
	t := s.Time
	d := 100.0
	for i := 0; i < 40; i++ {
		df, dt := s.DF(p, t)
		d = math.Min(d, df)
		t += dt
	}
	d = smoothstep(0, 0.01, d)
	return d * d * d
}

func (s *Spirograph) SetTime(Time float64) {
	s.Time = Time
	sgn := math.Copysign(1.0, 27-math.Mod(Time, 54))
	Time = math.Floor(math.Mod(Time, 27))
	var newScale float64
	switch {
	case Time < 0:
		newScale = (2 + Time/4) * sgn
	case Time < 20:
		newScale = (2 + Time/3) * sgn
	case Time < 21:
		newScale = 3.82845 * sgn
	case Time < 22:
		newScale = 3.64575 * sgn
	case Time < 23:
		newScale = 3.44955 * sgn
	case Time < 24:
		newScale = 2.7913 * sgn
	case Time < 25:
		newScale = 2.5616 * sgn
	case Time < 26:
		newScale = 2.4495 * sgn
	default:
		newScale = 2.30275 * sgn
	}
	s.Scale = newScale
}

// mix performs a linear interpolation between x and y using a to weight between them.
func mix(x, y, a float64) float64   { return x*(1-a) + y*a }
func add2(a, b vec2) vec2           { return vec2{a.x + b.x, a.y + b.y} }
func sub2(a, b vec2) vec2           { return vec2{a.x - b.x, a.y - b.y} }
func uclamp(x float64) float64      { return math.Max(0, math.Min(1, x)) }
func scale2(f float64, v vec2) vec2 { return vec2{f * v.x, f * v.y} }
func smoothstep(edge0, edge1, x float64) float64 {
	t := uclamp((x - edge0) / (edge1 - edge0))
	return t * t * (3 - 2*t)
}
