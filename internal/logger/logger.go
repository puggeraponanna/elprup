package logger

import (
	"context"
	"elprup/constant"

	"github.com/sirupsen/logrus"
)

func Init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.AddHook(NewCallerHook())
}

func WithContext(c context.Context) *logrus.Entry {
	r := c.Value(constant.REQUEST_ID)
	return logrus.WithField(constant.REQUEST_ID, r)
}

func Fatal(v ...any) {
	logrus.Fatal(v...)
}

func Panic(v ...any) {
	logrus.Panic(v...)
}

func Info(v ...any) {
	logrus.Info(v...)
}

func Warn(v ...any) {
	logrus.Warn(v...)
}

func Debug(v ...any) {
	logrus.Debug(v...)
}

func Error(v ...any) {
	logrus.Error(v...)
}
