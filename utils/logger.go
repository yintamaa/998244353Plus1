package utils

import (
	"sync"

	"go.uber.org/zap"
)

var logger *zap.Logger

var loggerOnce sync.Once

// 获取日志实例
func GetLoggerInstance() *zap.Logger {
	loggerOnce.Do(func() {
		var err error
		logger, err = zap.NewProduction()
		if err != nil {
			panic("logger init error")
		}
	})
	return logger
}
