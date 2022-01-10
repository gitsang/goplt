package main

import "fmt"

type A struct {
	Name string
	Age  int
}

func DoubleAge(a *A) int {
	return a.Age * 2
}

func CallDoubleAge(a *A) int {
	return DoubleAge(a)
}

func Print(a *A) {
	fmt.Println(a)
}

func Quote() {
	q := make([]*int, 1)
	qn := 2
	q[0] = &qn
}

func main() {
	a := &A{
		Name: "a",
		Age:  10,
	}
	DoubleAge(a)
	CallDoubleAge(a)
	Print(a)
	Quote()
}
