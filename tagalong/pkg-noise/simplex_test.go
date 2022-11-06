package noise

import "testing"

func TestNoise2D(t *testing.T) {
	for x := 0.0; x < 1; x += 0.01 {
		n := Simplex2D(x, 1)
		t.Error(n)
	}
}
