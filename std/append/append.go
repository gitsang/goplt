package main

import "fmt"

func main()  {
	var a []int
	printSlice("a", a)

	a = append(a, 0)
	printSlice("a", a)

	a = append(a, 1)
	printSlice("a", a)

	a = append(a, 2, 3, 4)
	printSlice("a", a)
}

func printSlice(name string, slice []int)  {
	fmt.Printf("%s: len:%d, cap:%d, value:%v\n",
		name, len(slice), cap(slice), slice)
}