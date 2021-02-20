package context

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

func TestCancel(t *testing.T) {
	origCtx := context.Background()
	ctx, cancel := context.WithCancel(origCtx)

	errg, ectx := errgroup.WithContext(ctx)

	errg.Go(func() error {
		for {
			select {
			case <-ectx.Done():
				fmt.Println("fucn1 ectx cancel")
				return nil
			default:
			}
			time.Sleep(time.Second)
			fmt.Println("func1 print")
		}
	})
	errg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("func2 ctx cancel")
				return nil
			default:
			}
			time.Sleep(time.Second)
			fmt.Println("func2 print")
		}
	})
	errg.Go(func() error {
		time.Sleep(500 * time.Second)
		fmt.Println("func3 error")
		return errors.New("err")
	})
	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	err := errg.Wait()
	if err != nil {
		fmt.Println("error", err.Error())
	}

	select {
	}
}
