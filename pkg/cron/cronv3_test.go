package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"testing"
	"time"
)
func TestCronV3(t *testing.T) {
	c := cron.New()
	_, _ = c.AddFunc("* * * * *",
		func() { fmt.Println("every min", time.Now().Unix()) })
	c.Start()
	defer c.Stop()

	select {}
}

func TestCronRemove(t *testing.T) {
	t.Log(t.Name(), "test start")
	c := cron.New()
	for {
		c.Stop()
		c = cron.New()
		for i := 0; i < 100000; i++ {
			_, _ = c.AddFunc("* * * * * ", func() {
				fmt.Println(time.Now().Unix())
			})
		}
		c.Run()
	}
}

func TestNil(t *testing.T) {
	var c *cron.Cron
	c.Stop()
}
