package main

import (
	"fmt"
	"os"
)

//func main() {
//	//valsparam_test()
//	errorf_test()
//}

func formater(key int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "key %d; ", key)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

func main() {
	key := 10
	name := "name"
	num := 1
	formater(key, "formater: %s|%d", name, num)
}