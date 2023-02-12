package logger

import (
	"log"

	"go.uber.org/zap"
)

type Logger interface {
	Zap() *zap.Logger
}

type logger struct {
	zap   *zap.Logger
	sugar *zap.SugaredLogger
}

func (l logger) Zap() *zap.Logger {
	return l.zap
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
