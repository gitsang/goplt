package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCounts := make(map[string]int)
	for i := range words {
		wordCounts[words[i]]++
	}
	return wordCounts
}

func main()  {
	fmt.Println(WordCount("What a pitty go what pitty go go Go a."))
}