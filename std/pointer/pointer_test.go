package pointer

import (
	"fmt"
	"strings"
	"testing"
)

func f(s *[]string) {
	*s = append(*s, "a")
}

func TestPointer(t *testing.T) {
	var slice []string
	f(&slice)
	t.Log("slice", slice)
}

func splitHostPorts(orig *string, hps *[]string) {
	*hps = make([]string, 0)
	ss := strings.Split(*orig, ",")
	for _, s := range ss {
		tmp := strings.TrimSpace(s)
		if tmp != "" {
			*hps = append(*hps, tmp)
		}
	}
}

type conf struct {
	orig string
	hps []string
}

func TestSplit(t *testing.T) {
	c := conf{
		orig: "abc, def, ghi, jkl",
	}
	splitHostPorts(&c.orig, &c.hps)
	t.Log(c.hps)
}

func TestMapNotExist(t *testing.T) {
	s := "INFO"
	t.Log("s.addr:", &s)

	m := make(map[string]*string)
	itemA, e := m["A"]
	if !e {
		itemA = &s
	}
	t.Log("m[\"A\"]:", m["A"], ", itemB:", itemA)

	itemB, e := m["B"]
	if !e {
		m["B"] = &s
	}
	t.Log("m[\"B\"]:", m["B"], ", itemB:", itemB)

	itemC, e := m["C"]
	if !e {
		m["C"] = &s
		itemC, _ = m["C"]
	}
	t.Log("m[\"C\"]:", m["C"], ", itemC:", itemC)
}

func newDefer() *conf {
	c := new(conf)
	defer func() {
		fmt.Println("c.orig", c.orig)
	}()

	return nil
}

func TestNewDefer(t *testing.T) {
	_ = newDefer()
}
