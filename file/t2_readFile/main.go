package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	mp := make(map[string]int)
	filenames := os.Args[1:]

	if len(filenames) != 0 {
		/* read file */
		for _, filename := range filenames {
			file, err := os.Open(filename)
			if err == nil {
				input := bufio.NewScanner(file)
				for input.Scan() {
					mp[input.Text()] = mp[input.Text()] + 1
				}
			}
		}

		/* print */
		for k, v := range mp {
			fmt.Printf("k:%s, v:%d", k, v)
		}
	}
}