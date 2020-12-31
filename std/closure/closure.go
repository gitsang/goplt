package main

import (
	"fmt"
	"math"
)

func main()  {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))

	funca := func(i int) int {
		return i * 2
	}
	b := funca(10)
	fmt.Println(b)
}
