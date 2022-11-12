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
	seed := time.Now().Unix() % 1000 // Seed 471 is particularily nice.
	rand.Seed(seed)
	is := NewImageServer("myimg.png")
	http.Handle("/", &is)
	fmt.Printf("started evolution server at http://localhost%s with seed %d. Spam F5 to run evolution.\n", address, seed)
	http.ListenAndServe(address, nil)
}

func NewImageServer(imagePath string) (is Montecarlo) {
	is.lock = new(sync.Mutex)
	return is
}

func randColor() color.Color {
	return color.RGBA{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 255}
}

func (is *Montecarlo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	is.lock.Lock()
	defer is.lock.Unlock()
	// We add `increment` points to the list of points.
	for i := 0; i < increment; i++ {
		is.points = append(is.points, Point{rand.Float64()*2 - 1, rand.Float64()*2 - 1})
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
	fmt.Fprintf(w, html, imgSize, imgSize, base64.StdEncoding.EncodeToString(buf.Bytes()), pi, math.Abs(math.Pi-pi))
}

func (is Montecarlo) At(i, j int) color.Color {
	x, y := 2*float64(i)/imgSize-1, 2*float64(j)/imgSize-1
	for _, point := range is.points {
		if math.Abs(x-point.X) < pointSize && math.Abs(y-point.Y) < pointSize {
			if math.Hypot(point.X, point.Y) < 1 {
				return color.RGBA{A: 0xff, R: 0xff}
			}
			return color.Black
		}
	}
	return color.White
}

func (is Montecarlo) Bounds() image.Rectangle {
	return image.Rect(0, 0, imgSize, imgSize)
}

func (is Montecarlo) ColorModel() color.Model {
	return color.RGBAModel
}

const html = `<!DOCTYPE html><html><head><title>Evolve</title></head><body>
	<img style='display:block; width:%dpx;height:%dpx;' src="data:image/png;base64,%s"/>
	<p>Approximation: %f,  Error:%.5g</p>
</body></html>`

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
