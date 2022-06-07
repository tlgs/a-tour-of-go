package main

import (
	"image"
	"image/color"
	"math/rand"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (Image) At(x, y int) color.Color {
	v := uint8(rand.Intn(256))
	return color.RGBA{255, v, v, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
