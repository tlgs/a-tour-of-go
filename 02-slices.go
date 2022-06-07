package main

import (
	"math/rand"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		s[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			s[y][x] = uint8(rand.Intn(256))
		}
	}

	return s
}

func main() {
	pic.Show(Pic)
}
