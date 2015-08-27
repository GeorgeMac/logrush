package logrush

import (
	"time"

	"github.com/Sirupsen/logrus"
)

type Context struct {
	*logrus.Entry
	context string
	start   time.Time
}

func For(context string) *Context {
	return logger.For(context)
}

func StartFor(context string) *Context {
	return logger.StartFor(context)
}

func (l *LogRush) For(context string) *Context {
	c := &Context{Entry: l.log()}
	c.context = context
	return c
}

func (l *LogRush) StartFor(context string) *Context {
	c := l.For(context)
	c.Start()
	return c
}

func (c *Context) Start() {
	c.start = now()
	c.WithField("start", c.start).
		WithField("context", c.context).
		Infof("Started %v", c.context)
}

func (c *Context) End() {
	end := now()
	c.WithField("end", end).
		WithField("context", c.context).
		WithField("duration", end.Sub(c.start)).
		Infof("Ended %v", c.context)
}
