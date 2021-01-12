package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mtx sync.Mutex
var cnt int

func handler(respW http.ResponseWriter, req *http.Request) {
	mtx.Lock()
	cnt ++
	mtx.Unlock()
	fmt.Fprintf(respW, "URL.Path = %q\n", req.URL.Path)

	fmt.Fprintf(respW, "%s %s %s\n", req.Method, req.URL, req.Proto)
	for k, v := range req.Header {
		fmt.Fprintf(respW, "Header[%q]: %q\n", k, v)
	}
	fmt.Fprintf(respW, "Host: %q\n", req.Host)
	fmt.Fprintf(respW, "RemoteAddr: %q\n", req.RemoteAddr)
	if err := req.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range req.Form {
		fmt.Fprintf(respW, "Form[%q]: %q\n", k, v)
	}
}

func counter(respW http.ResponseWriter, req *http.Request)  {
	mtx.Lock()
	fmt.Fprintf(respW, "cnt:%d\n", cnt)
	mtx.Unlock()
}

func main()  {
	http.HandleFunc("/", handler)
	http.HandleFunc("/cnt", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}


