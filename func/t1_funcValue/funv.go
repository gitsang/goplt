package main

import (
	"fmt"
	"math"
)

func main() {
	f1 := sqrt
	fmt.Println(f1(25))

	f2 := square
	fmt.Println(f2(5))
}

func square(n int) int {
	return n * n
}

func sqrt(n int) float64 {
	return math.Sqrt(float64(n))
}
