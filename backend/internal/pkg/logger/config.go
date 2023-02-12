package logger

import (
	"go.uber.org/zap"
)

var config zap.Config

func init() {
	// TODO config.yamlから設定できるようにする
	config = zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
}
