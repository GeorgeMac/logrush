package logrush

import (
	"time"

	"github.com/Sirupsen/logrus"
)

var now func() time.Time = time.Now

var logger *LogRush

func init() {
	logger = New("default")
}

func Set(opts ...Option) {
	logger.Set(opts...)
}

type LogRush struct {
	appKey, app string
	common      logrus.Fields
}

func New(app string, opts ...Option) *LogRush {
	l := &LogRush{
		appKey: "app",
		app:    app,
	}

	// apply options to LogRush instance
	l.Set(opts...)

	return l
}

func (l *LogRush) Set(opts ...Option) {
	for _, opt := range opts {
		opt(l)
	}
}

func (l *LogRush) logger() *logrus.Entry {
	logger := logrus.WithFields(logrus.Fields{
		l.appKey: l.app,
	})
	if l.common != nil {
		logger = logger.WithFields(l.common)
	}
	return logger
}
