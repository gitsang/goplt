package runtime

import "testing"

func printFunction() {
	println(FUNCTION())
	println(CALLER_FUNCTION())
}

func TestCallerFunction(t *testing.T) {
	printFunction()
}
