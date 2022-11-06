package noise

import (
	"math"
	"testing"
)

func TestNoise3D(t *testing.T) {
	const (
		span           = 100
		start          = 0.0
		stepsPerDim    = 20.0
		totalSteps     = stepsPerDim * stepsPerDim * stepsPerDim
		step           = span / stepsPerDim
		permissibleMax = 3.6
	)
	for x := start; x < start+span; x += step {
		for y := start; y < start+span; y += step {
			for z := start; z < start+span; z += step {
				n := Simplex3D(x, y, z)
				ok := math.Abs(n) < permissibleMax
				if !ok {
					t.Error(x, y, z, n)
				}
			}
		}
	}
}

func TestSimplex2D(t *testing.T) {
	const (
		span           = 100.0
		start          = 0.0
		stepsPerDim    = 1000
		totalSteps     = stepsPerDim * stepsPerDim
		step           = float64(span) / stepsPerDim
		permissibleMax = 1
	)
	for x := start; x < start+span; x += step {
		for y := start; y < start+span; y += step {
			n := Simplex2D(x, y)
			ok := math.Abs(n) < permissibleMax
			if !ok {
				t.Error(x, y, n)
			}
		}
	}
}

func TestSimplex1D(t *testing.T) {
	const (
		span           = 1000.0
		start          = -500.0
		stepsPerDim    = 1_000_000
		totalSteps     = stepsPerDim * stepsPerDim
		step           = float64(span) / stepsPerDim
		permissibleMax = 1
	)
	for x := start; x < start+span; x += step {
		n := Simplex1D(x)
		ok := math.Abs(n) < permissibleMax
		if !ok {
			t.Error(x, n)
		}
	}
}
