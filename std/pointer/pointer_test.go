package pointer

import (
	log "github.com/gitsang/golog"
	"go.uber.org/zap"
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
