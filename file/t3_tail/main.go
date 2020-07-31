package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"os"
)

func main() {
	seek := &tail.SeekInfo{}
	seek.Offset = 0
	seek.Whence = os.SEEK_END

	config := tail.Config{}
	config.Follow = true
	config.ReOpen = true
	config.MustExist = false
	config.Location = seek

	t, _ := tail.TailFile("log.txt", config)
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
