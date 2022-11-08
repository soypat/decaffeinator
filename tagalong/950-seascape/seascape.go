// Directly inspired by Seascape by TDM https://www.shadertoy.com/view/Ms2SD1
package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
)

const (
	imageSize        = 300
	Numsteps         = 8
	eps              = 1e-3
	epsNorm          = 0.1 / imageSize
	iterGeom         = 3
	iterFragDetailed = 5

	seaHeight = 0.6
	seaChoppy = 4.0
	seaSpeed  = 0.8
	seaFreq   = 0.16
)

var (
	seaBase          = vec3{0.0, 0.09, 0.18}
	seaWaterColor    = scale3(0.6, vec3{0.8, 0.9, 0.6})
	seaWaterColorP12 = scale3(0.12, seaWaterColor)
	light            = unit3(vec3{0.0, 1.0, 0.8})
	octave           = mat2{1.6, -1.2, 1.2, 1.6} // Row major differnece?
)

func main() {
	const colorMul = 255
	time := rand.Float64()
	_ = time
	img := image.NewRGBA(image.Rect(0, 0, imageSize, imageSize))
	for x := 0.0; x < imageSize; x++ {
		for y := 0.0; y < imageSize; y++ {
			uv := vec2{x, y}
			col := getPixel(uv, time)
			// col := vec3{seaOctave(uv, 4), seaOctave(uv, 4), seaOctave(uv, 4)}
			img.SetRGBA(int(x), int(y), color.RGBA{
				R: uint8(col.x * colorMul),
				G: uint8(col.y * colorMul),
				B: uint8(col.z * colorMul),
				A: 0xff,
			})
		}
	}
	fp, _ := os.Create("seascape.png")
	png.Encode(fp, img)
}

// lighting
func diffuse(n, l vec3, p float64) float64 { return math.Pow(dot3(n, l)*0.4+0.6, p) }
func specular(n, l, e vec3, s float64) float64 {
	nrm := (s + 8.0) / (math.Pi * 8)
	dot := dot3(reflect3(e, n), l)
	result := math.Pow(math.Max(0, dot), s) * nrm
	return result
}
func skyColor(e vec3) vec3 {
	e.y = (math.Max(e.y, 0)*0.8 + 0.2) * 0.8
	oneMinusEy := 1 - e.y
	return vec3{oneMinusEy * oneMinusEy, oneMinusEy, 0.6 + oneMinusEy*0.4}
}

func seaOctave(uv vec2, choppy float64) float64 {
	uv = addScalar2(noise2d(uv), uv)
	suv, cuv := sincos2(uv)
	wv := sub2(elem2(1), abs2(suv))
	swv := abs2(cuv)
	wv = mix2(wv, swv, wv)
	return math.Pow(1-math.Pow(wv.x*wv.y, 0.65), choppy)
}

func seaMap(detail int, p vec3, t float64) float64 {
	freq := seaFreq
	amp := seaHeight
	choppy := seaChoppy
	uv := vec2{p.x, p.z}
	uv.x *= 0.75
	var d, h float64
	for i := 0; i < detail; i++ {
		d = seaOctave(scale2(freq, addScalar2(t, uv)), choppy)
		d += seaOctave(scale2(freq, addScalar2(-t, uv)), choppy)
		h += d * amp
		uv = octave.mulvec(uv)
		freq *= 1.9
		amp *= 0.22
		choppy = mix(choppy, 1.0, 0.2)
	}
	return p.y - h
}

func seaColor(p, n, l, eye, dist vec3) vec3 {
	fresnel := uclamp(1 - dot3(n, scale3(-1, eye)))
	fresnel = fresnel * fresnel * fresnel * 0.5

	reflected := skyColor(reflect3(eye, n))
	diff := diffuse(n, l, 80.0)
	refracted := add3(seaBase, scale3(diff, seaWaterColorP12))

	color := mix3(refracted, reflected, elem3(fresnel))
	atten := math.Max(0.0, 1-0.001*dot3(dist, dist))

	color = add3(color, scale3(atten*0.18*(p.y-seaHeight), seaWaterColor))
	color = add3(color, elem3(specular(n, l, eye, 60.0)))
	return color
}

func heightMapTracing(ori, dir vec3, t float64) (p vec3, tmid float64) {
	tm := 0.0
	tx := 1000.0

	hx := seaMap(iterGeom, add3(ori, scale3(tx, dir)), t)
	if hx > 0 {
		p = add3(ori, scale3(tx, dir))
		return p, tx
	}
	hm := seaMap(iterGeom, add3(ori, scale3(tm, dir)), t)
	for i := 0; i < Numsteps; i++ {
		tmid = mix(tm, tx, hm/(hm-hx))
		p = add3(ori, scale3(tmid, dir))
		hmid := seaMap(iterGeom, p, t)
		if hmid < 0 {
			tx = tmid
			hx = hmid
		} else {
			tm = tmid
			hm = hmid
		}
	}
	return p, tmid
}

func getNormal(p vec3, eps, t float64) (n vec3) {
	n.y = seaMap(iterFragDetailed, p, t)
	n.x = seaMap(iterFragDetailed, vec3{p.x + eps, p.y, p.z}, t) - n.y
	n.z = seaMap(iterFragDetailed, vec3{p.x, p.y, p.z + eps}, t) - n.y
	n.y = eps
	return unit3(n)
}

func getPixel(coord vec2, t float64) vec3 {
	uv := scale2(2.0/imageSize, coord)
	uv = addScalar2(-1, uv)
	// ray
	// ang := vec3{math.Sin(3*t) * 0.1, math.Sin(t)*0.2 + 0.3, t}
	ori := vec3{0, 3.5, t * 5}
	dir := unit3(vec3{uv.x, uv.y, -2 + 0.14*math.Hypot(uv.x, uv.y)})
	dir = unit3(dir)
	// dir := dirGeneral //vec3{0.7, 0, 0}
	// Height map tracing.
	p, _ := heightMapTracing(ori, dir, t)
	dist := sub3(p, ori)
	n := scale3(epsNorm, getNormal(p, dot3(dist, dist), t))
	return mix3(
		skyColor(dir),
		seaColor(p, n, light, dir, dist),
		elem3(math.Pow(smoothstep(0, -0.02, dir.y), 0.2)),
	)
}

// //////////
// Vectors //
// //////////
type vec2 struct{ x, y float64 }
type vec3 struct{ x, y, z float64 }
type vec4 struct{ x, y, z, w float64 }

func mod4(a, b vec4) vec4 {
	return vec4{math.Mod(a.x, b.x), math.Mod(a.y, b.y), math.Mod(a.z, b.z), math.Mod(a.w, b.w)}
}
func abs2(a vec2) vec2 { return vec2{math.Abs(a.x), math.Abs(a.y)} }
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

func add4(a, b vec4) vec4 { return vec4{a.x + b.x, a.y + b.y, a.z + b.z, a.w + b.w} }
func sub4(a, b vec4) vec4 { return vec4{a.x - b.x, a.y - b.y, a.z - b.z, a.w - b.w} }
func cross3(a, b vec3) vec3 {
	return vec3{
		a.y*b.z - b.y*a.z,
		a.z*b.x - b.z*a.x,
		a.x*b.y - b.x*a.y,
	}
}
func unit3(v vec3) vec3 { // Equivalent to GLSL's normalize.
	norm := math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
	return scale3(1./norm, v)
}
func scale3(f float64, v vec3) vec3     { return vec3{v.x * f, v.y * f, v.z * f} }
func scale4(f float64, v vec4) vec4     { return vec4{v.x * f, v.y * f, v.z * f, v.w * f} }
func elem3(x float64) vec3              { return vec3{x, x, x} }
func elem4(x float64) vec4              { return vec4{x, x, x, x} }
func addScalar3(f float64, v vec3) vec3 { return vec3{v.x + f, v.y + f, v.z + f} }
func addScalar4(f float64, v vec4) vec4 { return vec4{v.x + f, v.y + f, v.z + f, v.w + f} }

func pow2(v vec2, exp float64) vec2 { return vec2{math.Pow(v.x, exp), math.Pow(v.y, exp)} }
func pow3(v vec3, exp float64) vec3 {
	return vec3{math.Pow(v.x, exp), math.Pow(v.y, exp), math.Pow(v.z, exp)}
}

// func normalize2(p vec2) vec2{}
func floor2(v vec2) vec2                { return vec2{math.Floor(v.x), math.Floor(v.y)} }
func add2(a, b vec2) vec2               { return vec2{a.x + b.x, a.y + b.y} }
func sub2(a, b vec2) vec2               { return vec2{a.x - b.x, a.y - b.y} }
func mul2(a, b vec2) vec2               { return vec2{a.x * b.x, a.y * b.y} }
func scale2(f float64, v vec2) vec2     { return vec2{v.x * f, v.y * f} }
func addScalar2(f float64, v vec2) vec2 { return vec2{v.x + f, v.y + f} }
func dot2(a, b vec2) float64            { return a.x*b.x + a.y*b.y }
func elem2(x float64) vec2              { return vec2{x, x} }
func sincos2(v vec2) (sin, cos vec2) {
	sx, cx := math.Sincos(v.x)
	sy, cy := math.Sincos(v.y)
	return vec2{sx, sy}, vec2{cx, cy}
}
func frac2(p vec2) vec2 {
	_, fracx := math.Modf(p.x)
	_, fracy := math.Modf(p.y)
	return vec2{fracx, fracy}
}
func reflect3(I, N vec3) vec3 { return sub3(I, scale3(2*dot3(N, I), N)) }
func modf2(p vec2) (int, frac vec2) {
	intx, fracx := math.Modf(p.x)
	inty, fracy := math.Modf(p.y)
	return vec2{intx, inty}, vec2{fracx, fracy}
}

func mix(x, y, a float64) float64 { return x*(1-a) + y*a }
func mix2(x, y, a vec2) vec2      { return vec2{mix(x.x, y.x, a.x), mix(x.y, y.y, a.y)} }
func mix3(x, y, a vec3) vec3      { return vec3{mix(x.x, y.x, a.x), mix(x.y, y.y, a.y), mix(x.z, y.z, a.z)} }
func uclamp(x float64) float64    { return math.Max(0, math.Min(1, x)) }
func smoothstep(edge0, edge1, x float64) float64 {
	t := uclamp((x - edge0) / (edge1 - edge0))
	return t * t * (3 - 2*t)
}

// Row major matrices

type mat2 [2 * 2]float64
type mat3 [3 * 3]float64

func (m *mat3) rowset(row int, v vec3) { m[row*3], m[row*3+1], m[row*3+2] = v.x, v.y, v.z }
func (m *mat3) colset(col int, v vec3) { m[col], m[3+col], m[6+col] = v.x, v.y, v.z }
func (m *mat3) mulvec(v vec3) vec3 {
	v.x = m[0]*v.x + m[1]*v.y + m[2]*v.z
	v.y = m[3]*v.x + m[4]*v.y + m[5]*v.z
	v.z = m[6]*v.x + m[7]*v.y + m[8]*v.z
	return v
}

func (m *mat2) mulvec(v vec2) vec2 {
	v.x = m[0]*v.x + m[1]*v.y
	v.y = m[2]*v.x + m[3]*v.y
	return v
}
func hash(p vec2) float64 {
	h := dot2(p, vec2{127.1, 311.7})
	_, frac := math.Modf(math.Sin(h) * 43758.5453123)
	return frac
}

func noise2d(p vec2) float64 {
	i, f := modf2(p)
	u := mul2(f, mul2(f, sub2(elem2(3), scale2(2, f))))
	n2 := hash(add2(i, vec2{1, 1}))
	n1 := mix(hash(i), hash(add2(i, vec2{1, 0})), u.x)
	n2 = mix(hash(add2(i, vec2{0, 1})), n2, u.x)
	return -1 + 2*mix(n1, n2, u.y)
}
