package cron

import (
	"testing"
	"time"

	log "github.com/gitsang/golog"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

// V3 format: min hour dayOfMouth mouth dayOfWeek
func TestCronV3(t *testing.T) {
	c := cron.New()

	// normal
	_, _ = c.AddFunc("* * * * *", func() {
		log.Info("every min", zap.String("time", time.Now().String()))
	})
	_, _ = c.AddFunc("0 * * * *", func() {
		log.Info("every hour", zap.String("time", time.Now().String()))
	})
	_, _ = c.AddFunc("0 3 * * *", func() {
		log.Info("at 3:00", zap.String("time", time.Now().String()))
	})

	// timezone: default is machine timezone
	_, _ = c.AddFunc("CRON_TZ=UTC 0 3-6 * * *", func() {
		log.Info("At minute 0 past every hour from 3 through 6 (UTC)", zap.String("time", time.Now().String()))
	})
	_, _ = c.AddFunc("CRON_TZ=UTC 30 18 * * *", func() {
		log.Info("At 18:30 (UTC)", zap.String("time", time.Now().String()))
	})

	// predefined
	_, _ = c.AddFunc("@hourly", func() {
		log.Info("Every hour", zap.String("time", time.Now().String()))
	})
	_, _ = c.AddFunc("@every 1h30m", func() {
		log.Info("Every 1.5 hour", zap.String("time", time.Now().String()))
	})

	c.Start()
	defer c.Stop()

	select {}
}
