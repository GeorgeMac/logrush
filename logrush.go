package logrush

import (
	"time"

	"github.com/Sirupsen/logrus"
)

var now = time.Now

var logger *LogRush

func init() {
	logger = New("default")
}

// Set calls LogRush.Set on the global logrush logger.
func Set(opts ...Option) {
	logger.Set(opts...)
}

// LogRush wraps logrus and adapts it with some "useful" idioms and configurations.
type LogRush struct {
	appKey, app string
	common      logrus.Fields
	formatter   logrus.Formatter
	level       logrus.Level
	logger      *logrus.Logger
}

// New returns a pointer to a new LogRush value.
// It takes an app name and a variadic number of logrush.Option types.
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

// Set can be called to update the configuration of the LogRush instance.
//
// note: It will create a new logrus logger underneath the bonnet!
func (l *LogRush) Set(opts ...Option) {
	for _, opt := range opts {
		opt(l)
	}
	l.refresh()
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
