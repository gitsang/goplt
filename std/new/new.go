package main

import "fmt"

func newInt() *int {
	var v int
	return &v
}

func main()  {
	p1 := new(int)
	fmt.Println(*p1)
	*p1 = 2
	fmt.Println(*p1)

	p2 := newInt()
	fmt.Println(*p2)
	*p2 = 2
	fmt.Println(*p2)
}