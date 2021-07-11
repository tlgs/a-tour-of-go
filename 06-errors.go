package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0.0, ErrNegativeSqrt(x)
	}

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

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
