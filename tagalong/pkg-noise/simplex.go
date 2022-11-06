package noise

// Simplex1D is a basic implementation of Ken Perlin's
// Simplex noise algorithm, which is an improvement on his original's
// classical noise algorithm known as classic or Perlin noise.
func Simplex1D(x float64) float64 {
	// https://github.com/devdad/SimplexNoise/blob/master/Source/SimplexNoise/Private/SimplexNoiseBPLibrary.cpp
	i0 := fastfloor(x)
	i1 := i0 + 1
	x0 := x - float64(i0)
	x1 := x0 - 1

	t0 := 1 - x0*x0
	t0 *= t0
	n0 := t0 * t0 * grad1(perm[i0&0xff], x0)

	t1 := 1 - x1*x1
	t1 *= t1
	n1 := t1 * t1 * grad1(perm[i1&0xff], x1)
	// From reference:
	// The maximum value of this noise is 8*(3/4)^4 = 2.53125
	// A factor of 0.395 would scale to fit exactly within [-1,1], but
	// we want to match PRMan's 1D noise, so we scale it down some more.
	return 0.395 * (n0 + n1)
}

// Simplex2D returns simplex noise on a 2D field.
func Simplex2D(x, y float64) float64 {
	return simplext2D(x, y)
}

// Returns a simplex noise sample at x, y, z.
// Value returned tends to be in interval (-1,1) though a bug
// is causing peaks of magnitude (-4.5,4.5) to appear randomly.
func Simplex3D(x, y, z float64) float64 {
	return snoise3(vec3{x, y, z})
}

func simplext2D(x, y float64) float64 {
	// https://github.com/devdad/SimplexNoise/blob/master/Source/SimplexNoise/Private/SimplexNoiseBPLibrary.cpp
	const (
		F2 = 0.3660254037844386  // 0.5*(sqrt(3.0)-1.0)
		G2 = 0.21132486540518713 // (3.0-Math.sqrt(3.0))/6.0
		S2 = 40.0 / 0.884343445  // final scaler.
	)
	s := (x + y) * F2 // Hairy factor for 2D.
	xs := x + s
	ys := y + s
	i := fastfloor(xs)
	j := fastfloor(ys)

	t := float64(i+j) * G2
	// Unskew the cell origin back to (x,y) space
	X0 := float64(i) - t
	Y0 := float64(j) - t
	// The x,y distances from the cell origin
	x0 := x - X0
	y0 := y - Y0

	// For the 2D case, the simplex shape is an equilateral triangle.
	// Determine which simplex we are in.
	var i1, j1 uint8
	if x0 > y0 {
		// lower triangle, XY order: (0,0)->(1,0)->(1,1)
		i1, j1 = 1, 0
	} else {
		// upper triangle, YX order: (0,0)->(0,1)->(1,1)
		i1, j1 = 0, 1
	}
	// Offsets for middle corner in (x,y) unskewed coords
	x1 := x0 - float64(i1) + G2
	y1 := y0 - float64(j1) + G2
	// Offsets for last corner in (x,y) unskewed coords
	x2 := x0 - 1 + 2*G2
	y2 := y0 - 1 + 2*G2

	// Wrap the integer indices at 256, to avoid indexing perm[] out of bounds
	ii := uint8(i & 0xff)
	jj := uint8(j & 0xff)
	// Calculate noise contributions from the three corners n0, n1, n2.
	var n0, n1, n2 float64
	t0 := 0.5 - x0*x0 - y0*y0
	if t0 >= 0 {
		t0 *= t0
		n0 = t0 * t0 * grad2(perm[ii+perm[jj]], x0, y0)
	}
	t1 := 0.5 - x1*x1 - y1*y1
	if t1 >= 0 {
		t1 *= t1
		n1 = t1 * t1 * grad2(perm[ii+i1+perm[jj+j1]], x1, y1)
	}
	t2 := 0.5 - x2*x2 - y2*y2
	if t2 >= 0 {
		t2 *= t2
		n2 = t2 * t2 * grad2(perm[ii+1+perm[jj+1]], x2, y2)
	}

	// Add contributions from each corner to get the final noise value.
	// The result is scaled to return values in the interval [-1,1]
	return S2 * (n0 + n1 + n2)
}

func grad2(hash uint8, x, y float64) float64 {
	// Convert low 3 bits of hash code
	h := hash & 0b111
	if h > 4 {
		x, y = y, x
	}
	if h&1 != 0 {
		x = -x
	}
	if h&2 != 0 {
		y = -y
	}
	return x + 2*y
}

func fastfloor(x float64) int {
	// x = math.Floor(x)
	if x > 0 {
		return int(x)
	}
	return int(x) - 1
}

var perm = [512]uint8{151, 160, 137, 91, 90, 15,
	131, 13, 201, 95, 96, 53, 194, 233, 7, 225, 140, 36, 103, 30, 69, 142, 8, 99, 37, 240, 21, 10, 23,
	190, 6, 148, 247, 120, 234, 75, 0, 26, 197, 62, 94, 252, 219, 203, 117, 35, 11, 32, 57, 177, 33,
	88, 237, 149, 56, 87, 174, 20, 125, 136, 171, 168, 68, 175, 74, 165, 71, 134, 139, 48, 27, 166,
	77, 146, 158, 231, 83, 111, 229, 122, 60, 211, 133, 230, 220, 105, 92, 41, 55, 46, 245, 40, 244,
	102, 143, 54, 65, 25, 63, 161, 1, 216, 80, 73, 209, 76, 132, 187, 208, 89, 18, 169, 200, 196,
	135, 130, 116, 188, 159, 86, 164, 100, 109, 198, 173, 186, 3, 64, 52, 217, 226, 250, 124, 123,
	5, 202, 38, 147, 118, 126, 255, 82, 85, 212, 207, 206, 59, 227, 47, 16, 58, 17, 182, 189, 28, 42,
	223, 183, 170, 213, 119, 248, 152, 2, 44, 154, 163, 70, 221, 153, 101, 155, 167, 43, 172, 9,
	129, 22, 39, 253, 19, 98, 108, 110, 79, 113, 224, 232, 178, 185, 112, 104, 218, 246, 97, 228,
	251, 34, 242, 193, 238, 210, 144, 12, 191, 179, 162, 241, 81, 51, 145, 235, 249, 14, 239, 107,
	49, 192, 214, 31, 181, 199, 106, 157, 184, 84, 204, 176, 115, 121, 50, 45, 127, 4, 150, 254,
	138, 236, 205, 93, 222, 114, 67, 29, 24, 72, 243, 141, 128, 195, 78, 66, 215, 61, 156, 180,
	151, 160, 137, 91, 90, 15,
	131, 13, 201, 95, 96, 53, 194, 233, 7, 225, 140, 36, 103, 30, 69, 142, 8, 99, 37, 240, 21, 10, 23,
	190, 6, 148, 247, 120, 234, 75, 0, 26, 197, 62, 94, 252, 219, 203, 117, 35, 11, 32, 57, 177, 33,
	88, 237, 149, 56, 87, 174, 20, 125, 136, 171, 168, 68, 175, 74, 165, 71, 134, 139, 48, 27, 166,
	77, 146, 158, 231, 83, 111, 229, 122, 60, 211, 133, 230, 220, 105, 92, 41, 55, 46, 245, 40, 244,
	102, 143, 54, 65, 25, 63, 161, 1, 216, 80, 73, 209, 76, 132, 187, 208, 89, 18, 169, 200, 196,
	135, 130, 116, 188, 159, 86, 164, 100, 109, 198, 173, 186, 3, 64, 52, 217, 226, 250, 124, 123,
	5, 202, 38, 147, 118, 126, 255, 82, 85, 212, 207, 206, 59, 227, 47, 16, 58, 17, 182, 189, 28, 42,
	223, 183, 170, 213, 119, 248, 152, 2, 44, 154, 163, 70, 221, 153, 101, 155, 167, 43, 172, 9,
	129, 22, 39, 253, 19, 98, 108, 110, 79, 113, 224, 232, 178, 185, 112, 104, 218, 246, 97, 228,
	251, 34, 242, 193, 238, 210, 144, 12, 191, 179, 162, 241, 81, 51, 145, 235, 249, 14, 239, 107,
	49, 192, 214, 31, 181, 199, 106, 157, 184, 84, 204, 176, 115, 121, 50, 45, 127, 4, 150, 254,
	138, 236, 205, 93, 222, 114, 67, 29, 24, 72, 243, 141, 128, 195, 78, 66, 215, 61, 156, 180,
}

func grad1(hash uint8, x float64) float64 {
	hash = hash & 15
	// Gradient value 1.0, 2.0, ..., 8.0
	grad := 1.0 + float64(hash&7)
	// Set a random sign for the gradient
	if (hash & 8) != 0 {
		grad = -grad
	}
	return grad * x // Multiply the gradient with the distance
}
