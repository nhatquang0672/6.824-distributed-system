package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z, d := 1., 1.
	for {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-d) < 1e-15 {
			break
		}
		d = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
