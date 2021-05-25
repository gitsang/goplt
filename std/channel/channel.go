package main

import "fmt"

func main() {
	ch := make(chan struct{})
	_, ok := <- ch
	if ok {
		fmt.Println("ok")
	} else {
		fmt.Println("!ok")
	}
	close(ch)
	if ok {
		fmt.Println("ok")
	} else {
		fmt.Println("!ok")
	}
}
