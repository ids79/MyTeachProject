package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	if e < 0 {
		return fmt.Sprintf("cannot Sqrt negative number: %.1f", e)
	}
	return fmt.Sprint(float64(e))
}

func Sqrt(x float64) ErrNegativeSqrt {
	if x < 0 {
		return ErrNegativeSqrt(x)
	}
	var z = 1.0
	for range 10 {
		if math.Abs((z*z-x)/(2*z)) < 0.000000000001 {
			break
		}
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return ErrNegativeSqrt(z)
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

func fibonacci() func() int {
	var f []int
	return func() int {
		switch len(f) {
		case 0:
			f = append(f, 0)
			return 0
		case 1:
			f = append(f, 1)
			return 1
		default:
			i := f[len(f)-2] + f[len(f)-1]
			f = append(f, i)
			return i
		}
	}
}

func append[T any](slice []T, data ...T) []T {
	initialLength := len(slice)
	finalLength := initialLength + len(data)
	if finalLength > cap(slice) {
		newSlice := make([]T, initialLength*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:finalLength]
	copy(slice[initialLength:finalLength], data)
	return slice
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	l, err := r.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i, s := range b {
		b[i] = rot13(s)
	}
	return l, err
}

func rot13(r byte) byte {
	if r >= 'a' && r <= 'z' { // Строчная буква
		if r >= 'm' {
			return r - 13
		} else {
			return r + 13
		}
	} else if r >= 'A' && r <= 'Z' { // Заглавная буква
		if r >= 'M' {
			return r - 13
		} else {
			return r + 13
		}
	}
	return r
}

func main() {
	//fmt.Println(Sqrt(2))
	// fmt.Println(math.Sqrt(2))
	// fmt.Println(5 ^ 1)
	// fmt.Println(WordCount("111 222 111 55 222 111"))
	// f := fibonacci()
	// for range 10 {
	// 	fmt.Println(f())
	// }
	// s := []byte("1234")
	// ptr := unsafe.SliceData(s)
	// str := unsafe.String(ptr, len(s))
	// p := unsafe.StringData("Str")
	// sl := unsafe.Slice(p, len("Str"))
	// i := [...]int{1, 2, 3}
	// i2 := i
	// p1 := &(i)
	// p2 := &(i)
	// fmt.Printf("%p - %p - %p - %p\n", &i, &i2, p1, p2)
	// s := make([]int, 0, 6)
	// fmt.Println(s[:6])
	// nameSlice := []string{"D", "a", "n", "i", "i", "l"}
	// nameSlice = append(nameSlice, "1", "3")
	// fmt.Println(nameSlice)
	// fmt.Println(len(nameSlice), cap(nameSlice))
	// nameSlice = slices.Grow(nameSlice, 18)
	// fmt.Println(len(nameSlice), cap(nameSlice))

	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

}
