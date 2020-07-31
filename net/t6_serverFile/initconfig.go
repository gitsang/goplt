package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var configMap = make(map[string]string)

func initConfig(configFileName string) {
	configFile, err := os.Open(configFileName)
	if err != nil {
		log.Fatal(err)
	}
	input := bufio.NewScanner(configFile)
	for i := 0; input.Scan(); i++ {
		line := strings.TrimSpace(input.Text())
		index := strings.Index(line, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(line[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(line[index+1:])
		if len(value) == 0 {
			continue
		}
		configMap[key] = value
	}
}

func loadConfig(key string) string {
	if _, exist := configMap[key]; exist {
		return configMap[key]
	}
	return ""
}

