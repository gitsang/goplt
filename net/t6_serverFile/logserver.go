package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func printLog(filename string, out io.Writer, filter string) {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	// read file
	input := bufio.NewScanner(file)
	for i := 0; input.Scan(); i++ {
		if line := input.Text(); strings.Contains(line, filter) {
			fmt.Fprintf(out, "%s\n", line)
		}
	}
}

func main() {
	initConfig("logserver.conf")
	logPath := loadConfig("log.path")
	serverIp := loadConfig("server.ip")
	serverPort := loadConfig("server.port")
	serverAddr := serverIp + ":" + serverPort

	// start filter
	for i := 1; ; i++ {
		fkey := loadConfig("filter." + strconv.Itoa(i) + ".key")
		fvalue:= loadConfig("filter." + strconv.Itoa(i) + ".value")
		if fkey == "" {
			break
		}
		http.HandleFunc("/" + fkey,
			func(writer http.ResponseWriter, request *http.Request) {
				printLog(logPath, writer, fvalue)
			})
	}

	// serve
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
