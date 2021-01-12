package main

import (
	"fmt"
	"path"
	"runtime"
)

func FUNCTION() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func FUNCTION_NAME() string {
	pc, _, _, _ := runtime.Caller(1)
	return path.Base(runtime.FuncForPC(pc).Name())
}

func CALLER_FUNCTION() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

func CALLER_FUNCTION_NAME() string {
	pc, _, _, _ := runtime.Caller(2)
	return path.Base(runtime.FuncForPC(pc).Name())
}

func lv2()  {
	fmt.Println("FUNCTION:", FUNCTION())
	fmt.Println("FUNCTION_NAME:", FUNCTION_NAME())
	fmt.Println("CALLER:", CALLER_FUNCTION())
	fmt.Println("CALLER_NAME", CALLER_FUNCTION_NAME())
}

func lv1() {
	lv2()
}

func main() {
	lv1()
}

