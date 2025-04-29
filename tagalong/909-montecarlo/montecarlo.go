package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Montecarlo struct {
	lock   *sync.Mutex
	rng    *rand.Rand
	points []Point
}

type Point struct {
	X, Y float64
}

const (
	pointSize = 0.01
	imgSize   = 500
	increment = 20
	address   = ":8080"
)

func main() {
	seed := time.Now().Unix() % 1000
	is := NewImageServer("myimg.png", seed)
	http.Handle("/", &is)
	fmt.Printf("started evolution server at http://localhost%s with seed %d. Spam F5 to run evolution.\n", address, seed)
	http.ListenAndServe(address, nil)
}

func NewImageServer(imagePath string, seed int64) (is Montecarlo) {
	is.lock = new(sync.Mutex)
	is.rng = rand.New(rand.NewSource(seed))
	return is
}

func (is *Montecarlo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	is.lock.Lock()
	defer is.lock.Unlock()
	start := time.Now()
	// We add `increment` points to the list of points.
	for i := 0; i < increment; i++ {
		is.points = append(is.points, is.randPoint())
	}

	// Calculate pi by counting points inside circle.
	insideCircle := 0
	for _, point := range is.points {
		if math.Hypot(point.X, point.Y) <= 1 {
			insideCircle++
		}
	}

	pi := 4 * float64(insideCircle) / float64(len(is.points))
	buf := new(bytes.Buffer)
	png.Encode(buf, is)
	elapsed := time.Since(start)
	fmt.Fprintf(w, `<!DOCTYPE html><html><head><title>Evolve</title></head><body>
	<img style='display:block; width:%dpx;height:%dpx;' src="data:image/png;base64,%s"/>
	<p>Approximation: %f,  Error:%.5g,  Elapsed:%s</p>
</body></html>`, imgSize, imgSize, base64.StdEncoding.EncodeToString(buf.Bytes()), pi, math.Abs(math.Pi-pi), elapsed.Round(time.Millisecond))
}

func (is Montecarlo) At(i, j int) color.Color {
	ij := Point{X: 2*float64(i)/imgSize - 1, Y: 2*float64(j)/imgSize - 1}
	for _, point := range is.points {
		if squareDrawer(ij, point, pointSize) {
			if math.Hypot(point.X, point.Y) < 1 {
				return color.RGBA{A: 0xff, R: 0xff}
			}
			return color.Black
		}
	}
	return color.White
}

func (is Montecarlo) randPoint() Point {
	return Point{is.rng.Float64()*2 - 1, is.rng.Float64()*2 - 1}
}

func (is Montecarlo) Bounds() image.Rectangle {
	return image.Rect(0, 0, imgSize, imgSize)
}

func (is Montecarlo) ColorModel() color.Model {
	return color.RGBAModel
}

func squareDrawer(pixPos, pointPos Point, size float64) bool {
	return math.Abs(pixPos.X-pointPos.X) < pointSize && math.Abs(pixPos.Y-pointPos.Y) < pointSize
}

func circleDrawer(pixPos, pointPos Point, size float64) bool {
	return math.Hypot(pixPos.X-pointPos.X, pixPos.Y-pointPos.Y) < pointSize
}
