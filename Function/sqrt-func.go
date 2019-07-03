package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for n := 0; n < 10; n++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(3))
}
