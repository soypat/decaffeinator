package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"net/http"
)

type ImageServer struct{}

const (
	imgSize = 1000
	squares = 8 // This is an approximate parameter.
	address = ":8080"
)

func main() {
	fmt.Printf("start server at http://localhost%s\n", address)
	is := ImageServer{}
	http.Handle("/", is)
	http.ListenAndServe(address, nil)
}

func (is ImageServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, is)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, html, imgSize, imgSize, base64.StdEncoding.EncodeToString(buf.Bytes()))
}

func (is ImageServer) At(i, j int) color.Color {
	x := math.J0((float64(i)) / (imgSize / (3 * squares)))
	y := math.J0((float64(j)) / (imgSize / (3 * squares)))
	z := math.Min(1, x*y+.05)
	return color.RGBA{
		R: uint8(x * 255),
		G: uint8(y * 255),
		B: uint8(z * 255),
		A: 255,
	}
}

func (is ImageServer) Bounds() image.Rectangle {
	return image.Rect(0, 0, imgSize, imgSize)
}

func (is ImageServer) ColorModel() color.Model {
	return color.RGBAModel
}

const html = `<!DOCTYPE html>
<html>
	<head>
		<title>My Server</title>
	</head>
	<body>
		<img style='display:block; width:%dpx;height:%dpx;' src="data:image/png;base64,%s"/>
	</body>
</html>`
