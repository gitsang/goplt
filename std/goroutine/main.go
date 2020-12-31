package main

import (
	"log"
	"strconv"
	"sync"
	"time"
)

func loop() {
	for range make([]int, 10) {
		log.Println("looping")
	}
}

func Parallelize(functions ...func()) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(functions))
	defer waitGroup.Wait()
	for _, function := range functions {
		go func(copy func()) {
			defer waitGroup.Done()
			copy()
		}(function)
	}
}

func Parallel() (s string) {
	var (
		loop = 5
	 	waitGroup sync.WaitGroup
	)

	log.Println("start")

	waitGroup.Add(loop)
	defer waitGroup.Wait()

	for i := 0; i < loop; i++ {
		i := i
		go func() {
			defer waitGroup.Done()
			time.Sleep(3 * time.Second)
			s = "loop" + strconv.Itoa(i)
			//_s := "loop" + strconv.Itoa(i)
			//s = _s
		}()
		time.Sleep(time.Second)
	}

	return s
}

func main() {
	log.Println("return: ", Parallel())
}
