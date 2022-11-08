package noise

import "math"

type vec2 struct{ x, y float64 }
type vec3 struct{ x, y, z float64 }
type vec4 struct{ x, y, z, w float64 }

func Basic2D(x, y float64) float64 {
	i, f := modf2(vec2{x, y})
	u := mul2(f, mul2(f, sub2(elem2(3), scale2(2, f))))
	n2 := hash2(add2(i, vec2{1, 1}))
	n1 := mix(hash2(i), hash2(add2(i, vec2{1, 0})), u.x)
	n2 = mix(hash2(add2(i, vec2{0, 1})), n2, u.x)
	return -1 + 2*mix(n1, n2, u.y)
}

// Simplex noise implementation. Has a bug.
// [reference implementation]: https://github.com/ashima/webgl-noise
func snoise2(v vec2) float64 {
	const (
		Cx = 0.211324865405187117745425 // (3.0-sqrt(3.0))/6.0
		Cy = 0.366025403784438646763723 // 0.5*(sqrt(3.0)-1.0)
		Cz = -1.0 + 2*Cx                //  -1.0 + 2.0 * C.x
		Cw = 1.0 / 41.                  //  1.0 / 41.0
	)
	// First corner.
	vcy := dot2(v, elem2(Cy))
	i := floor2(addScalar2(vcy, v))
	icx := dot2(i, elem2(Cx))
	x0 := addScalar2(icx, sub2(v, i))

	// Other corners.
	var i1 vec2
	if x0.x > x0.y {
		i1.x = 1
	} else {
		i1.y = 1
	}

	x12 := vec4{
		Cx + x0.x - i1.x,
		Cx + x0.y - i1.y,
		Cz + x0.x,
		Cz + x0.y,
	}

	// Permutations
	i = mod289_2(i) // Avoid truncation effects in permutation
	p := permute3(vec3{i.y, i.y + i1.y, i.y + 1.0})
	p = addScalar3(i.x, add3(p, vec3{0, i1.x, 1}))
	p = permute3(p)
	x12xy := vec2{x12.x, x12.y}
	x12zw := vec2{x12.z, x12.w}
	m := vec3{
		math.Max(0, 0.5-dot2(x0, x0)),
		math.Max(0, 0.5-dot2(x12xy, x12xy)),
		math.Max(0, 0.5-dot2(x12zw, x12zw)),
	}
	m = mul3(m, m)
	m = mul3(m, m)

	// Gradients: 41 points uniformly over a line, mapped onto a diamond.
	// The ring size 17*17 = 289 is close to a multiple of 41 (41*7 = 287)
	x := vec3{
		2*(1-math.Floor(p.x*Cw)) - 1,
		2*(1-math.Floor(p.y*Cw)) - 1,
		2*(1-math.Floor(p.z*Cw)) - 1,
	}
	h := abs3(x)
	ox := floor3(addScalar3(0.5, x))
	a0 := floor3(sub3(x, ox))

	// Normalise gradients implicitly by scaling m
	// Approximation of: m *= inversesqrt( a0*a0 + h*h );
	mx := add3(mul3(a0, a0), mul3(h, h))
	mx = addScalar3(1.79284291400159, scale3(-0.85373472095314, mx))
	m = mul3(m, mx)
	g := vec3{
		a0.x*x0.x + h.x*x0.y,
		a0.y*x12.x + h.y*x12.y,
		a0.z*x12.z + h.z*x12.w,
	}
	result := 130 * dot3(m, g)
	return result
}

// Simplex noise implementation. See [reference implementation].
//
// [reference implementation]: https://github.com/ashima/webgl-noise
func snoise3(v vec3) float64 {
	// https://www.youtube.com/watch?v=lctXaT9pxA0&ab_channel=SebastianLague
	const (
		Cx, Cy = 1.0 / 6.0, 1.0 / 3.0
	)
	i := floor3(addScalar3(dot3(v, elem3(Cy)), v))
	x0 := addScalar3(dot3(i, elem3(Cx)), sub3(v, i))

	//Other corners
	g := step3(vec3{x0.y, x0.z, x0.x}, x0)
	l := sub3(elem3(1), g)
	i1 := min3(g, vec3{l.z, l.x, l.y})
	i2 := max3(g, vec3{l.z, l.x, l.y})

	x1 := addScalar3(Cx, sub3(x0, i1))
	x2 := addScalar3(Cy, sub3(x0, i2))
	x3 := addScalar3(-0.5, x0)

	// Permutations
	i = mod289_3(i)
	p := permute4(addScalar4(i.z, vec4{y: i1.z, z: i2.z, w: 1.0}))
	p = add4(p, permute4(addScalar4(i.y, vec4{y: i1.y, z: i2.y, w: 1.0})))
	p = add4(p, permute4(addScalar4(i.x, vec4{y: i1.x, z: i2.x, w: 1.0})))

	// Gradients: 7x7 points over square, mapped onto octahedron. The ring size 17x17 = 289 is close to multiple of 49 (49*6 = 294) ????
	const (
		Dx, Dy, Dz, Dw = 0.0, 0.5, 1.0, 2.0
		d7             = 1.0 / 7.0
		// Why does the source do this? It is a mystery to me.
		nsx, nsy, nsz = d7*Dw - Dx, d7*Dy - Dz, d7*Dz - Dx
	)
	j := sub4(p, scale4(49, floor4(scale4(nsz*nsz, p)))) // mod(p,7*7)

	x_ := floor4(scale4(nsz, j))
	y_ := floor4(sub4(j, scale4(7, x_))) // mod(j, n)

	x := addScalar4(nsy, scale4(nsx, x_))
	y := addScalar4(nsy, scale4(nsx, y_))
	h := sub4(scale4(-1, abs4(x)), abs4(y))
	h = addScalar4(1, h)
	// x := addScalar4(0.5, scale4(2, x_))
	// x = addScalar4(-1, scale4(d7, x))
	// y := addScalar4(0.5, scale4(2, y_))
	// y = addScalar4(-1, scale4(d7, y))
	// h := addScalar4(1, sub4(scale4(-1, abs4(x)), abs4(y)))

	b0 := vec4{x.x, x.y, y.x, y.y}
	b1 := vec4{x.z, x.w, y.z, y.w}

	s0 := addScalar4(1, scale4(2, floor4(b0)))
	s1 := addScalar4(1, scale4(2, floor4(b1)))
	sh := scale4(-1, step4(h, elem4(0)))

	a0 := add4(vec4{b0.x, b0.z, b0.y, b0.w}, mul4(vec4{s0.x, s0.z, s0.y, s0.w}, vec4{sh.x, sh.x, sh.y, sh.y}))
	a1 := add4(vec4{b1.x, b1.z, b1.y, b1.w}, mul4(vec4{s1.x, s1.z, s1.y, s1.w}, vec4{sh.z, sh.z, sh.w, sh.w}))

	g0 := vec3{a0.x, a0.y, h.x}
	g1 := vec3{a0.z, a0.w, h.y}
	g2 := vec3{a1.x, a1.y, h.z}
	g3 := vec3{a1.z, a1.w, h.w}

	// Normalize gradients
	norm := taylorInvSqrt(vec4{dot3(g0, g0), dot3(g1, g1), dot3(g2, g2), dot3(g3, g3)})
	g0 = scale3(norm.x, g0)
	g1 = scale3(norm.y, g1)
	g2 = scale3(norm.z, g2)
	g3 = scale3(norm.w, g3)

	// Mix final noise value.
	m := max4(elem4(0), vec4{0.5 - dot3(x0, x0), 0.5 - dot3(x1, x1), 0.5 - dot3(x2, x2), 0.5 - dot3(x3, x3)})
	m = mul4(m, m)
	m = mul4(m, m)
	px := vec4{dot3(x0, g0), dot3(x1, g1), dot3(x2, g2), dot3(x3, g3)}
	return 105.0 * dot4(m, px)
}

func mod289_2(v vec2) vec2 {
	v.x -= math.Floor(v.x*(1.0/289.0)) * 289.0
	v.y -= math.Floor(v.y*(1.0/289.0)) * 289.0
	return v
}
func mod289_3(v vec3) vec3 {
	v.x -= math.Floor(v.x*(1.0/289.0)) * 289.0
	v.y -= math.Floor(v.y*(1.0/289.0)) * 289.0
	v.z -= math.Floor(v.z*(1.0/289.0)) * 289.0
	return v
}
func mod289_4(v vec4) vec4 {
	v.x -= math.Floor(v.x*(1.0/289.0)) * 289.0
	v.y -= math.Floor(v.y*(1.0/289.0)) * 289.0
	v.z -= math.Floor(v.z*(1.0/289.0)) * 289.0
	v.w -= math.Floor(v.w*(1.0/289.0)) * 289.0
	return v
}

func taylorInvSqrt(r vec4) vec4 {
	const a, b = 1.79284291400159, 0.85373472095314
	return vec4{
		a - r.x*b, a - r.y*b, a - r.z*b, a - r.w*b,
	}
}

func mod4(a, b vec4) vec4 {
	return vec4{math.Mod(a.x, b.x), math.Mod(a.y, b.y), math.Mod(a.z, b.z), math.Mod(a.w, b.w)}
}
func abs4(v vec4) vec4 {
	return vec4{math.Abs(v.x), math.Abs(v.y), math.Abs(v.z), math.Abs(v.w)}
}
func abs3(a vec3) vec3    { return vec3{math.Abs(a.x), math.Abs(a.y), math.Abs(a.z)} }
func max3(a, b vec3) vec3 { return vec3{math.Max(a.x, b.x), math.Max(a.y, b.y), math.Max(a.z, b.z)} }
func min3(a, b vec3) vec3 { return vec3{math.Min(a.x, b.x), math.Min(a.y, b.y), math.Min(a.z, b.z)} }
func floor3(v vec3) vec3  { return vec3{math.Floor(v.x), math.Floor(v.y), math.Floor(v.z)} }
func floor4(v vec4) vec4 {
	return vec4{math.Floor(v.x), math.Floor(v.y), math.Floor(v.z), math.Floor(v.w)}
}
func max4(a, b vec4) vec4 {
	return vec4{math.Max(a.x, b.x), math.Max(a.y, b.y), math.Max(a.z, b.z), math.Max(a.w, b.w)}
}

// hadamard product.
func mul4(a, b vec4) vec4 { return vec4{a.x * b.x, a.y * b.y, a.z * b.z, a.w * b.w} }

// hadamard product.
func mul3(a, b vec3) vec3 { return vec3{a.x * b.x, a.y * b.y, a.z * b.z} }

func add3(a, b vec3) vec3     { return vec3{a.x + b.x, a.y + b.y, a.z + b.z} }
func dot3(a, b vec3) float64  { return a.x*b.x + a.y*b.y + a.z*b.z }
func dot4(a, b vec4) float64  { return a.x*b.x + a.y*b.y + a.z*b.z + a.w*b.w }
func sub3(a, b vec3) vec3     { return vec3{a.x - b.x, a.y - b.y, a.z - b.z} }
func step3(edge, p vec3) vec3 { return vec3{step(edge.x, p.x), step(edge.y, p.y), step(edge.z, p.z)} }
func step4(edge, p vec4) vec4 {
	return vec4{step(edge.x, p.x), step(edge.y, p.y), step(edge.z, p.z), step(edge.w, p.w)}
}
func step(edge, point float64) float64 {
	if point >= edge {
		return 1
	}
	return 0
}
func permute3(x vec3) vec3 {
	x.x, x.y, x.z = (x.x*34+10)*x.x, (x.y*34+10)*x.y, (x.z*34+10)*x.z
	return mod289_3(x)
}
func permute4(x vec4) vec4 {
	x.x, x.y, x.z, x.w = (x.x*34+10)*x.x, (x.y*34+10)*x.y, (x.z*34+10)*x.z, (x.w*34+10)*x.w
	return mod289_4(x)
}
func add4(a, b vec4) vec4 { return vec4{a.x + b.x, a.y + b.y, a.z + b.z, a.w + b.w} }
func sub4(a, b vec4) vec4 { return vec4{a.x - b.x, a.y - b.y, a.z - b.z, a.w - b.w} }
func cross3(a, b vec3) vec3 {
	return vec3{
		a.y*b.z - b.y*a.z,
		a.z*b.x - b.z*a.x,
		a.x*b.y - b.x*a.y,
	}
}
func scale3(f float64, v vec3) vec3     { return vec3{v.x * f, v.y * f, v.z * f} }
func scale4(f float64, v vec4) vec4     { return vec4{v.x * f, v.y * f, v.z * f, v.w * f} }
func elem3(x float64) vec3              { return vec3{x, x, x} }
func elem4(x float64) vec4              { return vec4{x, x, x, x} }
func addScalar3(f float64, v vec3) vec3 { return vec3{v.x + f, v.y + f, v.z + f} }
func addScalar4(f float64, v vec4) vec4 { return vec4{v.x + f, v.y + f, v.z + f, v.w + f} }

func floor2(v vec2) vec2                { return vec2{math.Floor(v.x), math.Floor(v.y)} }
func add2(a, b vec2) vec2               { return vec2{a.x + b.x, a.y + b.y} }
func sub2(a, b vec2) vec2               { return vec2{a.x - b.x, a.y - b.y} }
func mul2(a, b vec2) vec2               { return vec2{a.x * b.x, a.y * b.y} }
func scale2(f float64, v vec2) vec2     { return vec2{v.x * f, v.y * f} }
func addScalar2(f float64, v vec2) vec2 { return vec2{v.x + f, v.y + f} }
func dot2(a, b vec2) float64            { return a.x*b.x + a.y*b.y }
func elem2(x float64) vec2              { return vec2{x, x} }
func mix(x, y, a float64) float64       { return x*(1-a) + y*a }

func hash2(p vec2) float64 {
	h := dot2(p, vec2{127.1, 311.7})
	_, frac := math.Modf(math.Sin(h) * 43758.5453123)
	return frac
}
func modf2(p vec2) (int, frac vec2) {
	intx, fracx := math.Modf(p.x)
	inty, fracy := math.Modf(p.y)
	return vec2{intx, inty}, vec2{fracx, fracy}
}
