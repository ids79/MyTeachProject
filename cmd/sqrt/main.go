package main

import (
	"fmt"
	"math"
	"strings"
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

func Pic(dx, dy int) [][]uint8 {
	var img = make([][]uint8, dx)

	for i := range dx {
		img[i] = make([]uint8, dy)
		for j := range dy {
			img[i][j] = uint8((i + j) / 2)
		}
	}
	return img
}

func WordCount(s string) map[string]int {
	var c = make(map[string]int)
	strs := strings.Fields(s)
	for _, w := range strs {
		c[w]++
	}
	return c
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
	fmt.Println(5 ^ 1)
	fmt.Println(WordCount("111 222 111 55 222 111"))
}
