package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func f(vs ...int) { }
func g(vs []int) { }

func valsparam_test() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(3, 4, 1 ,2))
	values := []int{1, 2, 4, 8}
	fmt.Println(sum(values...))
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)
}

