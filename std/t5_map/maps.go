package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main()  {
	// define
	var mp = map[string]Vertex {
		"Astring": Vertex{
			13.2344, -234.22,
		},
		"Bstring": {
			Lat:  342.123,
			Long: 12.5324,
		},
		"Cstring" : {123.243, 234.123},
	}
	fmt.Println(mp)

	// add
	mp["Dstring"] = Vertex{40.23421, -124.234}
	fmt.Println(mp)

	// change
	mp["Astring"] = Vertex{123.123, 123.456}
	fmt.Println(mp)
	mp["Astring"] = Vertex{111.222, 542.126}
	fmt.Println(mp)

	// delete
	delete(mp, "Astring")
	fmt.Println(mp)

	// check
	v, ok := mp["Astring"]
	fmt.Println("Value:", v, " Present?", ok)
}
