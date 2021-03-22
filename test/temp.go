package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 0
	b := 0
	x := a/b
	fmt.Println(x)
	c := float64(a)/float64(b)
	fmt.Println(c)
	d := int(c*100)
	fmt.Println(d)
	s := strconv.Itoa(d)
	fmt.Println(s)
}