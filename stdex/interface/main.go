package _interface

type TestInterface interface {
	func1()
	func2()
}

// use this to ensure interface implement is completed
//var _ TestInterface = test{}

type test struct {
}

func (t *test) func1() {
}

