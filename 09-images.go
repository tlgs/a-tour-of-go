package main

import (
	"image"
	"image/color"
	"math"

	"golang.org/x/tour/pic"
)

type Image struct {
	w int
	h int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	dx := float64(x - i.w/2)
	dy := float64(y - i.h/2)

	v := uint8(math.Sqrt(dx*dx+dy*dy)) * 4
	return color.RGBA{v, v, v, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
