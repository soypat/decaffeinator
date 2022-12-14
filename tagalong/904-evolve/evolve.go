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
	"os"
	"sync"
	"time"
)

type ImageServer struct {
	lock       *sync.Mutex
	agents     []Agent
	background image.Image
	imagePath  string
}

type Agent struct {
	velx, vely float64
	posx, posy int
	c          color.Color
}

const (
	imgSize   = 500
	numAgents = 3
	agentSize = imgSize/256 + 1
	address   = ":8080"
)

func main() {
	seed := time.Now().Unix() % 1000 // Seed 471 is particularily nice.
	rand.Seed(seed)
	is := NewImageServer("myimg.png")
	http.Handle("/", is)
	fmt.Printf("started evolution server at http://localhost%s with seed %d. Spam F5 to run evolution.\n", address, seed)
	http.ListenAndServe(address, nil)
}

func NewImageServer(imagePath string) (is ImageServer) {
	is.lock = new(sync.Mutex)
	is.imagePath = imagePath
	is.background = image.NewUniform(color.White)
	fp, err := os.Create(imagePath)
	if err != nil {
		panic(err)
	}
	png.Encode(fp, is) // Creates white background image.
	fp.Close()
	for i := 0; i < numAgents; i++ {
		is.agents = append(is.agents, Agent{posx: rand.Intn(imgSize), posy: rand.Intn(imgSize), c: randColor()})
	}
	return is
}

func randColor() color.Color {
	return color.RGBA{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 255}
}

func (is ImageServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	is.lock.Lock()
	defer is.lock.Unlock()
	const softening = 1.0 / imgSize // Softening parameter to avoid 0 distance singularities.
	const G = imgSize / 500.0       // "Gravitational constant".
	// First update agents.
	for i, agent1 := range is.agents {
		for j, agent2 := range is.agents[i+1:] {
			// Calculate force and thus acceleration on agent1.
			distancex, distancey := float64(agent2.posx-agent1.posx), float64(agent2.posy-agent1.posy)
			euclidean := math.Max(math.Sqrt(distancex*distancex+distancey*distancey), softening)
			accelx, accely := G*distancex/euclidean, G*distancey/euclidean
			is.agents[i].velx += accelx
			is.agents[i].vely += accely
			is.agents[i+j+1].velx -= accelx
			is.agents[i+j+1].vely -= accely
		}
	}
	for i, agent := range is.agents {
		// Update agent positions with updated velocities.
		is.agents[i].posx += int(agent.velx)
		is.agents[i].posy += int(agent.vely)
	}
	// Next load previous image.
	fp, err := os.Open(is.imagePath)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	is.background, err = png.Decode(fp)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fp.Close()
	// Now generate new image. The At method composes current agents with loaded image.
	buf := new(bytes.Buffer)
	err = png.Encode(buf, is)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	newImage := buf.Bytes()
	fmt.Fprintf(w, html, imgSize, imgSize, base64.StdEncoding.EncodeToString(newImage))
	// Don't forget to update image on disk for next read.
	fp, err = os.Create(is.imagePath)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = png.Encode(fp, is)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fp.Close()
}

func (is ImageServer) At(i, j int) color.Color {
	for _, agent := range is.agents {
		if abs(i-agent.posx) < agentSize && abs(j-agent.posy) < agentSize {
			return agent.c
		}
	}
	return is.background.At(i, j)
}

func (is ImageServer) Bounds() image.Rectangle {
	return image.Rect(0, 0, imgSize, imgSize)
}

func (is ImageServer) ColorModel() color.Model {
	return color.RGBAModel
}

const html = `<!DOCTYPE html><html><head><title>Evolve</title></head><body>
	<img style='display:block; width:%dpx;height:%dpx;' src="data:image/png;base64,%s"/>
</body></html>`

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
