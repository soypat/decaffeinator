package noise

import (
	"math"
	"math/rand"
)

type PerlinGrid struct {
	xSize, ySize int
	grid         [][2]float64
	interp       func(a, b, x float64) float64
}

func NewPerlinGrid(x, y int, interp func(a, b, x float64) float64) PerlinGrid {
	p := PerlinGrid{grid: make([][2]float64, (x+1)*(y+1)), xSize: x, ySize: y}
	for i := range p.grid {
		p.grid[i][0] = rand.Float64()
		p.grid[i][1] = rand.Float64()
	}
	p.interp = interp
	if interp == nil {
		p.interp = smoothInterp
	}
	return p
}

// x and y contained in interval [0,1)
func (p PerlinGrid) At(x, y float64) float64 {
	// Scale [0,1) interval to perlin grid domain.
	x *= float64(p.xSize)
	y *= float64(p.ySize)

	// Determine surrounding grid cell coordinates.
	ix0, iy0 := int(x), int(y)
	ix1, iy1 := ix0+1, iy0+1

	// Determine interpolation weights.
	sx := x - math.Floor(x)
	sy := y - math.Floor(y)

	// Interpolate between grid point gradients.
	n0 := p.dotGradient(ix0, iy0, x, y)
	n1 := p.dotGradient(ix1, iy0, x, y)
	interp0 := p.interp(n0, n1, sx)

	n0 = p.dotGradient(ix0, iy1, x, y)
	n1 = p.dotGradient(ix1, iy1, x, y)
	interp1 := p.interp(n0, n1, sx)
	return p.interp(interp0, interp1, sy)*0.5 + 0.5
}

func (p PerlinGrid) dotGradient(ix, iy int, x, y float64) float64 {
	gradient := p.grid[iy*p.xSize+ix] // Row major storage, though this is not important for the purposes of the algorithm.
	dx := x - float64(ix)
	dy := y - float64(iy)
	return dx*gradient[0] + dy*gradient[1]
}

func linearInterp(a0, a1, x float64) float64 { return (a1-a0)*x + a0 }
func cubicInterp(a0, a1, x float64) float64  { return (a1-a0)*(3.0-x*2.0)*x*x + a0 }
func smoothInterp(a0, a1, x float64) float64 { return (a1-a0)*((x*(x*6.0-15.0)+10.0)*x*x*x) + a0 }
