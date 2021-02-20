package main

import (
	"configor/config"
	CfgAgent "configor/config/cfgagent"
	"gitcode.yealink.com/server/server_framework/go-utils/ylog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func main() {
	err := ylog.InitYlog(ylog.Config{
		Level:  zapcore.DebugLevel,
		Stdout: ylog.LogStdoutConfig{
			Enable: true,
			Color:  true,
		},
	})

	err = config.LoadConfig("main_test.yml")
	if err != nil {
		ylog.Error("load config failed", zap.Error(err))
	}

	CfgAgent.RegisterCallback("id_one", func(v string) {
		ylog.Info("my call back", zap.Any("v", v))
	})

	go func() {
		for {
			config.PrintConfig()
			select {
			case <-time.Tick(10 * time.Second):
			}
		}
	}()

	select {
	}
}
