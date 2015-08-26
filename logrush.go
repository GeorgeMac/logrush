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
	formatter   logrus.Formatter
	level       logrus.Level
	logger      *logrus.Logger
}

func New(app string, opts ...Option) *LogRush {
	l := &LogRush{
		appKey: "app",
		app:    app,
		level:  logrus.InfoLevel,
	}

	// apply options to LogRush instance
	l.Set(opts...)

	return l
}

func (l *LogRush) Set(opts ...Option) {
	for _, opt := range opts {
		opt(l)
	}
	l.refresh()
}

func (l *LogRush) Context(context string) *Context {
	c := &Context{Entry: l.log()}
	c.context = context
	return c
}

func (l *LogRush) BeginTimed(context string) *Context {
	c := l.Context(context)
	c.Start()
	return c
}

func (l *LogRush) refresh() {
	l.logger = logrus.New()
	l.logger.Level = l.level

	if l.formatter != nil {
		l.logger.Formatter = l.formatter
	}
}

func (l *LogRush) log() *logrus.Entry {
	e := l.logger.WithField(l.appKey, l.app)
	if l.common != nil {
		e = e.WithFields(l.common)
	}
	return e
}
