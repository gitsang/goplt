package cron

import (
	"fmt"
	"github.com/robfig/cron"
	"testing"
	"time"
)

// | Seconds    | 0-59	          | * / , -
// | Minutes    | 0-59	          | * / , -
// | Hours      | 0-23	          | * / , -
// | DayOfMonth | 1-31	          | * / , – ?
// | Month      | 1-12 or JAN-DEC | * / , -
// | DayOfWeek  | 0-6 or SUM-SAT  | * / , – ?

func TestCronSec(t *testing.T) {
	c := cron.New()
	_ = c.AddFunc("* * * * *",
		func() { fmt.Println("every sec", time.Now().Unix()) })
	c.Start()
	defer c.Stop()

	select {}
}

func TestCron(t *testing.T) {
	c := cron.New()
	_ = c.AddFunc("0 * * * * *",
		func() { fmt.Println("every min", time.Now().Unix()) })
	_ = c.AddFunc("0 0 * * * *",
		func() { fmt.Println("every hour", time.Now().Unix()) })

	_ = c.AddFunc("CRON_TZ=UTC 0 0 3-6 * * *",
		func() { fmt.Println("every hour(range 3-6)", time.Now().Unix()) })
	_ = c.AddFunc("CRON_TZ=UTC 0 30 18 * * *",
		func() { fmt.Println("every day 18:30UTC") })

	_ = c.AddFunc("@hourly",
		func() { fmt.Println("Every hour, starting an hour from now") })
	_ = c.AddFunc("@every 1h30m",
		func() { fmt.Println("Every hour thirty, starting an hour thirty from now") })
	c.Start()
	defer c.Stop()

	select {}
}

func TestCronStop(t *testing.T) {
	c := cron.New()
	for {
		c.Stop()
		c = cron.New()
		_ = c.AddFunc("* * * * * *", func() {
			fmt.Println(time.Now().Unix())
		})
		t.Log("start begin")
		c.Start()
		t.Log("start end")
		time.Sleep(10 * time.Second)
	}
}

func TestCronTrash(t *testing.T) {
	t.Log(t.Name(), "test start")
	c := cron.New()
	for {
		c.Stop()
		c = cron.New()
		for i := 0; i < 100000; i++ {
			_ = c.AddFunc("0 * * * * *", func() {
				fmt.Println(time.Now().Unix())
			})
		}
		c.Start()
	}
}
