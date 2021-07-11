package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for y := range p {
		q := make([]uint8, dx)
		for x := range q {
			q[x] = uint8(x ^ y)
		}
		p[y] = q
	}

	return p
}

func main() {
	pic.Show(Pic)
}
