package main

import "fmt"

func Sqrt(x float64) float64 {
	z := 1.0
	for epsilon := 1.0; epsilon > 1e-10; {
		diff := (z*z - x) / (2 * z)
		z -= diff

		if diff > 0 {
			epsilon = diff
		} else {
			epsilon = -diff
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
