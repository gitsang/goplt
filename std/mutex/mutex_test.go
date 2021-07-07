package mutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRWMutex(t *testing.T) {
	var l sync.RWMutex
	m := make(map[string]int64)

	// write
	go func() {
		for {
			l.Lock()
			m["NOW"] = time.Now().Unix()
			l.Unlock()
		}
	}()

	// read-1
	go func() {
		for {
			l.RLock()
			fmt.Println("NOW-1", m["NOW"])
			l.RUnlock()
		}
	}()

	// read-2
	go func() {
		for {
			l.RLock()
			fmt.Println("NOW-2", m["NOW"])
			l.RUnlock()
		}
	}()

	select {}
}

func TestUnlock(t *testing.T) {
	var mtx sync.Mutex
	mtx.Lock()
	t.Log("lock")
	mtx.Unlock()
	t.Log("unlock once")
	mtx.Unlock()
	t.Log("unlock twice")
}
