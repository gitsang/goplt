package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "just a flag")
var ac = flag.String("ac", " ", "count of value")

func main()  {
	flag.Parse()
	for _, arg := range flag.Args() {
		fmt.Println(strings.Count(arg, *ac))
	}

	if !*n {
		fmt.Println("n: false")
	} else {
		fmt.Println("n: true")
	}
}