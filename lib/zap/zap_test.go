package zap

import (
	log "github.com/gitsang/golog"
	"go.uber.org/zap"
	"testing"
)

func TestField(t *testing.T) {
	Field := make([]zap.Field, 0)
	field := make([]zap.Field, 0)

	field = append(field, zap.String("a", "a"), zap.String("b", "b"))
	Field = append(Field, zap.Any("field", field))
	log.Info("info", Field...)
}

type queues []int

type subInfo struct {
	Topic     string
	Group     string
	QueueMaps map[string]queues
}

func TestField2(t *testing.T) {
	sub := subInfo{
		Topic: "topic2",
		Group: "group2",
		QueueMaps: make(map[string]queues),
	}
	log.Info("queue list update", zap.Reflect("subInfo", sub))

	//sub.QueueMaps["broker-a"] = make([]int, 0)
	sub.QueueMaps["broker-a"] = append(sub.QueueMaps["broker-a"], 1)
	sub.QueueMaps["broker-a"] = append(sub.QueueMaps["broker-a"], 2)
	sub.QueueMaps["broker-a"] = append(sub.QueueMaps["broker-a"], 3)
	//sub.QueueMaps["broker-b"] = make([]int, 0)
	sub.QueueMaps["broker-b"] = append(sub.QueueMaps["broker-b"], 0)
	sub.QueueMaps["broker-b"] = append(sub.QueueMaps["broker-b"], 1)
	sub.QueueMaps["broker-b"] = append(sub.QueueMaps["broker-b"], 2)

	log.Info("queue list update", zap.Any("subInfo", sub))
}