package logic

import (
	"bytes"
	"encoding/json"
	"github.com/pingcap/log"
	"go.uber.org/zap"
	"sync"
	"testing"
	"time"
)

// etcd -----------------------------------------------------------------------------

var stLock sync.Mutex
var storage []byte

// struct  -----------------------------------------------------------------------------

type Info interface {
	Marshal() ([]byte, error)
	Gc()
}

type History interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
	Clone() []Info
	DeleteOne()
	UpdateOne(data []byte)
}

type HistoryReviseItem struct {
	A string `json:"a"`
	B string `json:"b"`
}

type HistoryReviseInfo struct {
	list []*HistoryReviseItem
}

type historyReviseList struct {
	list map[string]*HistoryReviseInfo
	sync.RWMutex
}

// func ---------------------------------------------------------------------- global

func Persist(info Info) {
	body, err := info.Marshal()
	if err != nil {
		panic(err)
	}
	if bytes.Contains(body, []byte("null")) {
		panic(body)
	}
	//log.Info("", zap.ByteString("body", body))

	stLock.Lock()
	storage = body
	stLock.Unlock()
}

// func ---------------------------------------------------------------------- info

func (src *HistoryReviseInfo) clone() *HistoryReviseInfo {
	ret := *src
	ret.list = make([]*HistoryReviseItem, 0)
	for _, v := range src.list {
		ret.list = append(ret.list, v)
	}
	return &ret
}

func (src *HistoryReviseInfo) Marshal() ([]byte, error) {
	return json.Marshal(src.list)
}

func (src *HistoryReviseInfo) Gc() {
	for len(src.list) > 0 {
		if len(src.list) > 1 {
			src.list = src.list[1:]
		} else {
			src.list = make([]*HistoryReviseItem, 0)
		}
	}
	Persist(src)
}

// func ---------------------------------------------------------------------- list

var hrl *historyReviseList

func Init() {
	hrl = new(historyReviseList)
	hrl.list = make(map[string]*HistoryReviseInfo)
}

func (src *historyReviseList) Clone() []Info {
	list := make([]Info, 0)
	for _, v := range src.list {
		list = append(list, v.clone())
	}
	return list
}

func (src *historyReviseList) UpdateOne(data []byte) {
	var list []*HistoryReviseItem
	err := json.Unmarshal(data, &list)
	if err != nil {
		log.Warn("unmarshal failed", zap.Error(err), zap.ByteString("data", data))
		list = make([]*HistoryReviseItem, 0)
	}
	log.Info("unmarshal success", zap.ByteString("data", data))
	src.list["A"] = &HistoryReviseInfo{list: list}
}

func (src *historyReviseList) DeleteOne() {
	delete(src.list, "A")
}

// func ---------------------------------------------------------------------- test

func TestClone(t *testing.T) {
	Init()

	// add
	go func() {
		for {
			log.Info("add")
			time.Sleep(10 * time.Millisecond)
			hrl.RLock()
			info := hrl.list["A"]
			if info != nil {
				info = info.clone()
			} else {
				info = &HistoryReviseInfo{list: make([]*HistoryReviseItem, 0)}
			}
			hrl.RUnlock()

			info.list = append(info.list, &HistoryReviseItem{
				A: "a",
				B: "b",
			})
			Persist(info)
		}
	}()

	// watch
	go func(history History) {
		for {
			log.Info("watch")
			time.Sleep(10 * time.Millisecond)
			history.Lock()
			history.UpdateOne(storage)
			history.Unlock()
		}
	}(hrl)

	// gc
	go func(history History) {
		for {
			log.Info("gc")
			time.Sleep(1000 * time.Millisecond)
			history.RLock()
			list := history.Clone()
			history.RUnlock()

			for _, info := range list {
				info.Gc()
			}
		}
	}(hrl)

	for {
		time.Sleep(time.Hour)
	}
}
