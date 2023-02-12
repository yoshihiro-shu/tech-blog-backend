package logger

import (
	"go.uber.org/zap"
)

var config zap.Config

func init() {
	// TODO config.yamlから設定できるようにする
	config = zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	encoderConfig := zap.NewProductionEncoderConfig()
	// encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig
}
