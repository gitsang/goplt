package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

func InitLogger(level zapcore.Level, path string) {
	sugarLogger = NewLogger(level, path).Sugar()
}

func Sync() {
	err := sugarLogger.Sync()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Debug(args ...interface{}) {
	sugarLogger.Debug(args)
}

func Info(args ...interface{}) {
	sugarLogger.Info(args)
}

func Warn(args ...interface{}) {
	sugarLogger.Warn(args)
}

func Error(args ...interface{}) {
	sugarLogger.Error(args)
}
