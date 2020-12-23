package slice

import "testing"

type tss struct {
	A string `json:"a"`
	B string `json:"b"`
}

func TestSlice(t *testing.T) {
	list := make([]*tss, 0)
	for i := 0; i < 100; i++ {
		list = append(list, nil)
	}
	t.Log("listSize", len(list))
	list = list[50:]
	t.Log("listSize", len(list))
	t.Log(list[0].A)
}