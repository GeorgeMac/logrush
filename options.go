package logrush

import "github.com/Sirupsen/logrus"

// Option is a function which takes a pointer to a LogRush value.
// Its purpose is to be used as a functional Option for configuring
// LogRush
type Option func(*LogRush)

// AppKey returns an Option to set the application key for logging against.
// For example, passing LogRush.AppKey("service") will result in log lines
// with a field named "service".
func AppKey(appKey string) Option {
	return func(l *LogRush) {
		l.appKey = appKey
	}
}

// App returns an Option to set the application value for logging against.
// For example, passing LogRush.App("my-app") will result in log lines
// with a field `appKey` (which will be "app" by default) with a value "my-app".
func App(app string) Option {
	return func(l *LogRush) {
		l.app = app
	}
}

// Common returns an Option to set the default fields to be set on all logrush lines
// For example, passing LogRush.Common(logrus.Fields{"env":"production"}) will
// result in all log lines having a field named "env" with a value "production".
func Common(common logrus.Fields) Option {
	return func(l *LogRush) {
		l.common = common
	}
}

// Formatter returns an Option to set the logrus.Formatter type on the logrush instance.
// For example, logrus supports JSON formatted logs for logging in to things like logstash.
// To use this formatter just use logrush.Formatter(&logrus.JSONFormatter{}) to get the Option.
func Formatter(formatter logrus.Formatter) Option {
	return func(l *LogRush) {
		l.formatter = formatter
	}
}

// Level returns an Option for setting the logrus.Level on the logrush instance.
// For example, use logrush.Level(logrus.WarningLevel) to produce an Option whichs
// sets the logging level to Warning.
func Level(level logrus.Level) Option {
	return func(l *LogRush) {
		l.level = level
	}
}
