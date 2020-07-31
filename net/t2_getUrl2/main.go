package main

import (
	"io"
	"net/http"
	"os"
)

func main()  {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		os.Exit(1)
	}
	io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
}
