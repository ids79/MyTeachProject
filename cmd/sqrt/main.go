package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z = 1.0
	for range 10 {
		if math.Abs((z*z-x)/(2*z)) < 0.000000000001 {
			break
		}
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
