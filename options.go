package logrush

import "github.com/Sirupsen/logrus"

type Option func(*LogRush)

func AppKey(appKey string) Option {
	return func(l *LogRush) {
		l.appKey = appKey
	}
}

func App(app string) Option {
	return func(l *LogRush) {
		l.app = app
	}
}

func Common(common logrus.Fields) Option {
	return func(l *LogRush) {
		l.common = common
	}
}

func Formatter(formatter logrus.Formatter) Option {
	return func(l *LogRush) {
		l.formatter = formatter
	}
}

func Level(level logrus.Level) Option {
	return func(l *LogRush) {
		l.level = level
	}
}
