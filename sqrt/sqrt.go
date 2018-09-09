package main

import (
	"fmt"
	"math"
)

// This is the example code out of the go-tour
// that you can find in flowcontrol/8
func Sqrt(x float64, d float64) (z float64) {
	z = 1.0
	y := z * 2
	for math.Abs(z-y) > d {
		y = z
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2, 0.001))
	fmt.Println(math.Sqrt(2))
}
