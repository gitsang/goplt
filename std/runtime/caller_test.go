package runtime

import "testing"

func printFunction() {
	println(FUNCTION())
	println(CALLER_FUNCTION())
	println(CALLER_LINE())
}

func TestCallerFunction(t *testing.T) {
	printFunction()
}
