package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	mp := make(map[string]string)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		mp[input.Text()] = input.Text()
		fmt.Printf("scan finished.\n")
	}
	for k, v := range mp {
		fmt.Printf("k:%s, v:%s", k, v)
	}
}

