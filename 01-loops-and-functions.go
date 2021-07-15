package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for epsilon := 1.0; math.Abs(epsilon) > 1e-10; {
		epsilon = (z*z - x) / (2 * z)
		z -= epsilon
	}

	return z
}

func main() {
	fmt.Println("Sqrt(2)     :", Sqrt(2))
	fmt.Println("math.Sqrt(2):", math.Sqrt(2))
}
