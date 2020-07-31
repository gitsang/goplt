package main

import (
	"log"
	"time"
)

func usedTime() func() {
	start := time.Now()
	log.Printf("start")
	return func() {
		log.Printf("used time: %s", time.Since(start))
	}
}

func usetime_test() {
	defer usedTime()()
	// ...
	time.Sleep(5 * time.Second)
}

func main() {
	usetime_test()
}
