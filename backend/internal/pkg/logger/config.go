package logger

import (
	"go.uber.org/zap"
)

var config zap.Config

func init() {
	// doesn't work!
	// jst, err := time.LoadLocation("Asia/Tokyo")
	// if err != nil {
	// 	panic(err)
	// }
	// time.Local = jst

	// TODO config.yamlから設定できるようにする
	config = zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	encoderConfig := zap.NewProductionEncoderConfig()
	// TOOD Setting TimeZone
	// encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig
}
