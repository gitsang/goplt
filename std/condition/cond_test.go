package condition

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var global = false

func TestCond(t *testing.T) {
	var wg sync.WaitGroup
	//var mtx sync.Mutex
	cond := sync.Cond{L: &sync.Mutex{}}

	wg.Add(1)
	go func() {
		cond.L.Lock()
		defer func() {
			cond.L.Unlock()
			wg.Done()
		}()

		for global == false {
			fmt.Println("goroutine 1 wait")
			cond.Wait()
		}
		fmt.Println("goroutine 1 done")
	}()

	wg.Add(1)
	go func() {
		cond.L.Lock()
		defer func() {
			cond.L.Unlock()
			wg.Done()
		}()

		for global == false {
			fmt.Println("goroutine 2 wait")
			cond.Wait()
		}
		fmt.Println("goroutine 2 done")
	}()

	time.Sleep(2 * time.Second)
	cond.L.Lock()
	fmt.Println("main ready")
	global = true
	cond.Broadcast()
	fmt.Println("broadcast end")
	time.Sleep(2 * time.Second)
	fmt.Println("main unlock")
	cond.L.Unlock()
	wg.Wait()
}
