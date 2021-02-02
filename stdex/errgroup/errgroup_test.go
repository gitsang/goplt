package errgroup

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

func errFunc0() error {
	fmt.Println("func0 in")
	defer fmt.Println("func0 out")
	time.Sleep(1 * time.Second)
	return nil
}

func errFunc1() error {
	fmt.Println("func1 in")
	defer fmt.Println("func1 out")
	time.Sleep(3 * time.Second)
	return errors.New("error1")
}

func errFunc2() error {
	fmt.Println("func2 in")
	defer fmt.Println("func2 out")
	time.Sleep(5 * time.Second)
	return errors.New("error2")
}

func errFunc3() error {
	fmt.Println("func3 in")
	defer fmt.Println("func3 out")
	time.Sleep(10 * time.Second)
	return errors.New("error3")
}

func errFunc4() error {
	fmt.Println("func4 in")
	defer fmt.Println("func4 out")
	return errors.New("error4")
}

func TestErrGroup(t *testing.T) {
	for {
		select {
		case <-time.Tick(time.Second):
		}

		fmt.Println("(re)start")
		eg := new(errgroup.Group)

		eg.Go(errFunc0)
		eg.Go(errFunc2)
		eg.Go(errFunc3)
		eg.Go(errFunc1)
		eg.Go(errFunc4)

		err := eg.Wait()
		if err != nil {
			fmt.Println("errgroup recv error", err.Error())
			continue
		} else {
			fmt.Println("success")
			break
		}
	}
}
