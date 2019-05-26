package log

import (
	"github.com/sirupsen/logrus"
)

func Panic(f interface{}, v ...interface{}) {
	logger.Panic(formatLog(f, v...))
}

func Fatal(f interface{}, v ...interface{}) {
	//logger.Fatal(formatLog(f, v...))
	logger.Error(formatLog(f, v...))
}

func Error(f interface{}, v ...interface{}) {
	logger.Error(formatLog(f, v...))
}

func Warn(f interface{}, v ...interface{}) {
	logger.Warn(formatLog(f, v...))
}

func Warning(f interface{}, v ...interface{}) {
	logger.Warn(formatLog(f, v...))
}

func Info(f interface{}, v ...interface{}) {
	logger.Info(formatLog(f, v...))
}

func Notice(f interface{}, v ...interface{}) {
	logger.Info(formatLog(f, v...))
}

func Debug(f interface{}, v ...interface{}) {
	logger.Debug(formatLog(f, v...))
}

func Close() {
	Info("log: closing...")
	logrus.Exit(0)
}
