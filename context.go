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
