package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if z < 0.1 {
			return z
		}
	}
	return z
}

func main() {
	if Sqrt(2) == math.Sqrt(2) {
		fmt.Println(Sqrt(2))
		fmt.Println(math.Sqrt(2))
	}
}
