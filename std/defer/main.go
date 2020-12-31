package main

import "fmt"

func deferTest1() string {
	s := "init"
	defer func() {
		s = "defer"
	}()
	return s
}

func deferTest2() (s string) {
	s = "init"
	defer func() {
		s = "defer"
	}()
	return s
}

func main() {
	fmt.Println(deferTest1())
	fmt.Println(deferTest2())
}