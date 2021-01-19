package pointer

import (
	log "github.com/gitsang/golog"
	"go.uber.org/zap"
	"strings"
	"testing"
)

func f(s *[]string) {
	*s = append(*s, "a")
}

func TestPointer(t *testing.T) {
	var slice []string
	f(&slice)
	log.Info(zap.Any("slice", slice))
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