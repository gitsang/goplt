package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for k, v := range pow {
		fmt.Printf("2**%d = %d\n", k, v)
	}
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}