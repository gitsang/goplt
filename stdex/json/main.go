package main

import (
	"encoding/json"
	"fmt"
)

type RegionInfo struct {
	Region     string   `json:"region"`
	DumpServer []string `json:"dump-server"`
	Disable    bool     `json:"disable"`
}

func getRegionList() []*RegionInfo {
	list := make([]*RegionInfo, 0)
	list = append(list, &RegionInfo{
		Region:     "region-a",
		DumpServer: nil,
		Disable:    false,
	})
	list = append(list, &RegionInfo{
		Region:     "region-b",
		DumpServer: nil,
		Disable:    true,
	})
	return list
}

type Response struct {
	Data interface{} `json:"data"`
}

func ResponseOK(data interface{}) *Response {
	return &Response{Data: data}
}

func main() {
	regions := getRegionList()
	jsons, _ := json.Marshal(regions)
	fmt.Printf("json: %v", string(jsons))
	fmt.Printf("response: %v", ResponseOK(regions))
}
