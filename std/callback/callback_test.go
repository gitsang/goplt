package callback

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type Config struct {
	time string
	time_callback func()
}

var config = Config{
	time:          "",
	time_callback: nil,
}

func printTimeCallback() {
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(config)
		}
	}()
}

type Callback func(ctx context.Context)

func configCallback(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("done")
				return
			default:
				time.Sleep(time.Second)
			}
			fmt.Println(config)
		}
	}()
}

func configChanged(new string, callback Callback) {
	//config = new
	ctx, _ := context.WithCancel(context.Background())
	callback(ctx)
}

func watchingEvent() {
	for {
		// pretend to revive config
		configChanged(time.Now().String(), configCallback)

		select {
		case <-time.Tick(10 * time.Second):
		}
	}
}

func TestEventCallback(t *testing.T) {
	go watchingEvent()

	for {
		time.Sleep(time.Hour)
	}
}
