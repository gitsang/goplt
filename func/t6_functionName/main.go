package main

import (
	"fmt"
	"runtime"
	"strings"
)

func FUNCTION() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func CALLER_FUNCTION() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

func CALLER_FUNCTION_NAME() string {
	pc, _, _, _ := runtime.Caller(2)
	split := strings.Split(runtime.FuncForPC(pc).Name(), "/")
	return split[len(split)-1]
}

func lv2()  {
	fmt.Println(FUNCTION())
	fmt.Println(CALLER_FUNCTION())
	fmt.Printf(CALLER_FUNCTION_NAME())
}

func lv1() {
	lv2()
}

func main() {
	lv1()
}

