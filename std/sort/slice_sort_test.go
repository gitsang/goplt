package sort

import (
	"fmt"
	"sort"
	"testing"
)

type people struct {
	id   int
	name string
}

func TestSilceSort(t *testing.T) {
	s := []people{
		{id: 9, name: "9"},
		{id: 4, name: "4"},
		{id: 3, name: "3"},
		{id: 5, name: "5"},
		{id: 8, name: "8"},
	}
	sort.SliceStable(s, func(i, j int) bool {
		return s[i].id < s[j].id
	})

	fmt.Println(s)
}
