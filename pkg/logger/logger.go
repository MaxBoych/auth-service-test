package logger

import (
	"fmt"
	"go.uber.org/zap"
)

var log *zap.Logger

func Init() (err error) {
	log, err = zap.NewDevelopment()
	return
}

func Infof(template string, args ...interface{}) {
	log.Info(fmt.Sprintf(template, args...))
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

func Errorf(template string, args ...interface{}) {
	log.Error(fmt.Sprintf(template, args...))
}
