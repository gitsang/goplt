package main

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"
)

type tss struct {
	A string `json:"a"`
	B string `json:"b"`
}

type tssList struct {
	list []*tss
}

func (l *tssList) clone() *tssList {
	ret := *l
	ret.list = make([]*tss, 0)
	for _, v := range l.list {
		ret.list = append(ret.list, v)
	}
	return &ret
}

func TestGc(t *testing.T) {
	mgr := tssList{
		list: []*tss{
			{
				A: "a1",
				B: "b1",
			},
			{
				A: "a2",
				B: "b2",
			},
		},
	}

	// append
	go func() {
		for {
			nn := mgr.clone()
			nn.list = append(nn.list, &tss{
				A: "a3",
				B: "b3",
			})
		}
	}()

	// delete
	go func() {
		for {
			nn := mgr.clone()
			nn.list = nn.list[1:]
		}
	}()

	// marshal
	go func() {
		for {
			nn := mgr.clone()
			j, _ := json.Marshal(nn.list)
			if bytes.Contains(j, []byte("null")) {
				t.Log(string(j))
			}
		}
	}()

	for {
		time.Sleep(time.Hour)
	}
}