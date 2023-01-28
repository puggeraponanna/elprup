package logger

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/sirupsen/logrus"
)

type CallerHook struct {
	once sync.Once
}

func (ch *CallerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (ch *CallerHook) Fire(entry *logrus.Entry) error {
	ch.once.Do(func() {
		_, file, line, ok := runtime.Caller(12)
		if ok {
			entry.Data["file"] = fmt.Sprintf("%s:%d", file, line)
		}
	})
	return nil
}

func NewCallerHook() *CallerHook {
	return &CallerHook{}
}
