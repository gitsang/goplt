package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func defer_test() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	bytesRead, err := file.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("size:%d, string:%s", bytesRead, string(buffer))
}

func main() {
	defer_test()
}
