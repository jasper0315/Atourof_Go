package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	delta := 0.00001

	for math.Abs(z*z-x) > delta {
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func main() {
	x := 16.0
	fmt.Printf("Sqrt(%g) = %g\n", x, Sqrt(x))
}
