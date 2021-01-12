package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func t_anonymous() {
	f1 := squares()
	f2 := squares()
	fmt.Println(f1) // 0x49e3b0
	fmt.Println(f1()) // 1
	fmt.Println(f1()) // 4
	fmt.Println(f1()) // 9
	fmt.Println(f2) // 0x49e3b0
	fmt.Println(f2()) // 1
	fmt.Println(f2()) // 4
	fmt.Println(f2()) // 9
	fmt.Println(squares()) // 0x49e3b0
	fmt.Println(squares()()) // 1
	fmt.Println(squares()()) // 1
	fmt.Println(squares()()) // 1
}

func main() {
	t_anonymous()
}
