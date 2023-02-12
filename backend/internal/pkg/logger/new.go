package logger

import (
	"log"

	"go.uber.org/zap"
)

type Logger interface {
	Debug(msg string, opts ...zap.Field)
	Info(msg string, opts ...zap.Field)
	Warn(msg string, opts ...zap.Field)
	Error(msg string, opts ...zap.Field)
}

type logger struct {
	zap   *zap.Logger
	sugar *zap.SugaredLogger
}

func (l logger) Debug(msg string, opts ...zap.Field) {
	l.zap.Debug(msg, opts...)
}

func (l logger) Info(msg string, opts ...zap.Field) {
	l.zap.Info(msg, opts...)
}

func (l logger) Warn(msg string, opts ...zap.Field) {
	l.zap.Warn(msg, opts...)
}

func (l logger) Error(msg string, opts ...zap.Field) {
	l.zap.Error(msg, opts...)
}

func New() Logger {
	zap, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("ERROR AT Init Logger %s\n", err.Error())
		panic(err)
	}
	return logger{
		zap:   zap,
		sugar: zap.Sugar(),
	}
}
