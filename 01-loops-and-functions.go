package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for e := 1.0; math.Abs(e) > 1e-9; z -= e {
		e = (z*z - x) / (2 * z)
	}

	return z
}

func main() {
	fmt.Println("Sqrt(2)     :", Sqrt(2))
	fmt.Println("math.Sqrt(2):", math.Sqrt(2))
}
