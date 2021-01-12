package runtime

import (
	"path"
	"runtime"
)

func FUNCTION() string {
	pc, _, _, _ := runtime.Caller(1)
	return path.Base(runtime.FuncForPC(pc).Name())
}

func CALLER_FUNCTION() string {
	pc, _, _, _ := runtime.Caller(2)
	return path.Base(runtime.FuncForPC(pc).Name())
}
