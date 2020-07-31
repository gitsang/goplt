package main

import (
	"fmt"
	"os"
)

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

func errorf_test() {
	linenum := 12
	name := "name"
	num := 1
	errorf(linenum, "underfined: %s|%d", name, num)
}