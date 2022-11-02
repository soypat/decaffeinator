package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

func main() {
	art := NewRandomArt(500, 500, 160)
	fp, _ := os.Create("agentart.png")
	err := png.Encode(fp, art)
	if err != nil {
		panic(err)
	}
	fp.Close()
}

type Art struct {
	width, height int
	last          image.Image
	agents        []Agent
}

type Agent struct {
	size       int
	posx, posy int
	c          color.Color
}

func NewRandomArt(width, height, agents int) Art {
	// Start on an empty white canvas.
	art := Art{
		width:  width,
		height: height,
		agents: make([]Agent, agents),
		last:   image.NewUniform(color.White),
	}
	for i := 0; i < agents; i++ {
		art.agents[i] = Agent{
			posx: rand.Intn(width),
			posy: rand.Intn(height),
			size: rand.Intn(width/20) + 1,
			c:    color.RGBA{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 255},
		}
	}
	return art
}

func (a Art) Bounds() image.Rectangle {
	return image.Rect(0, 0, a.width, a.height)
}

func (a Art) ColorModel() color.Model {
	return color.RGBAModel
}

func (a Art) At(i, j int) color.Color {
	for _, agent := range a.agents {
		if abs(i-agent.posx) < agent.size && abs(j-agent.posy) < agent.size {
			return agent.c
		}
	}
	return a.last.At(i, j)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
