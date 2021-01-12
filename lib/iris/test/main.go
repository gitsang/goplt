package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const addr = "http://localhost:8085"

type Info struct {
	ID   int64  `json:"id"`
	City string `json:"city"`
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Info Info
}

func getJson()  {
	fmt.Println("getJson")

	resp, err := http.Get(addr + "/json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(s))
}

func postJson() {
	fmt.Println("postJson")

	body := &User{
		Name: "Friend A",
		Age:  15,
		Info: Info{
			ID:   56,
			City: "D",
		},
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)
	resp, _ := http.Post(addr + "/json", "UTF-8", buf)
	s, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("status: %s, body: %s", resp.Status, string(s))
}

func main() {
	//getJson()
	postJson()
}
