package main

import "testing"

func TestNilValue(t *testing.T) {
	s := make([]int64, 0)
	m := make(map[int64]interface{})
	m[123] = nil
	m[124] = nil

	for k, _ := range m {
		s = append(s, k)
	}

	for _, v := range s {
		println(v)
	}
}
