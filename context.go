package logrush

import (
	"time"

	"github.com/Sirupsen/logrus"
)

// Context adapts a *logrus.Entry in order to simplify the logging of
// operations and their time to execute over a context.
type Context struct {
	*logrus.Entry
	context string
	start   time.Time
}

// For returns a new Context instance from the global logrush instance.
func For(context string) *Context {
	return logger.For(context)
}

// StartFor returns a new Context instance from the global logrush instance
// and calls Start on the instance.
func StartFor(context string) *Context {
	return logger.StartFor(context)
}

// For returns a new Context instance on the LogRush instance.
// Note all subsequence chained calls on the Context will result in a log
// line with the field "context" with the value of the context string.
func (l *LogRush) For(context string) *Context {
	c := &Context{Entry: l.log()}
	c.context = context
	return c
}

// StartFor returns a new Context instance on the LogRush instance and it calls Start.
// Note all subsequence chained calls on the Context will result in a log
// line with the field "context" with the value of the context string.
func (l *LogRush) StartFor(context string) *Context {
	c := l.For(context)
	c.Start()
	return c
}

// Start products a log line through the provided Context with the follows fields
// start: time.Now() of call
// context: context string stored on the Context object
// msg: Started {{ context }}
func (c *Context) Start() {
	c.start = now()
	c.WithField("start", c.start).
		WithField("context", c.context).
		Infof("Started %v", c.context)
}

// End products a log line through the provided Context with the follows fields
// end: time.Now() of call
// duration: duration since call to Start() i.e. time.Now().Sub(start)
// context: context string stored on the Context object
// msg: Ended {{ context }}
func (c *Context) End() {
	end := now()
	c.WithField("end", end).
		WithField("context", c.context).
		WithField("duration", end.Sub(c.start)).
		Infof("Ended %v", c.context)
}
