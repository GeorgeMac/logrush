LogRush
=======

[![GoDoc](https://godoc.org/github.com/GeorgeMac/logrush?status.svg)](https://godoc.org/github.com/GeorgeMac/logrush)

# What is this?

It’s `logrus` with some common contextual stuff I write a lot

# Why?

1. Because I love `logrus` and I keep configuring it a similar way between my applications
2. Because I want some common context between my log lines

# So show me what I can do with it!

```go
import (
    log "github.com/GeorgeMac/logrush"
    "github.com/Sirupsen/logrus"
)

func main() {
    env := os.GetEnv("GO_ENV")
    if env == "" {
        env = "development"
    }

    fields := logrus.Fields{"environment": env}

    // global logger
    log.Set(log.App("my-service"),
        log.Common(fields),
        log.Formatter(&logrus.JSONFormatter{}),
        log.Level(logrus.DebugLevel))

    log.Printf(...)
    log.Debugf(...)

    // app="my-service" environment="development" context="Request for 192.168.0.1" start="1st Jan, 2015" msg="Started Request from 192.168.0.1"
    log.For("Request from 192.168.0.1").Info("Something about the request")

    // app="my-service" environment="development" context="Request for 192.168.0.2" start="1st Jan, 2015" msg="Started Request from 192.168.0.2"
    logger := log.StartFor("Request from 192.168.0.2")
    // app="my-service" environment="development" context="Request for 192.168.0.2" start="1st Jan, 2015" msg="Processing Request"
    logger.Infof("Processing request")
    // app="my-service" environment="development" context="Request for 192.168.0.2" end="2nd Jan, 2015" duration="1d" msg="Ended Request from 192.168.0.2"
    defer logger.End()

    // separate logger
    logger := log.New("other-service",
        log.Common(fields),
        log.Level(logrus.InfoLevel))
}
```
