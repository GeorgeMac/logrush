LogRush
=======

# What is this?

It’s `logrus` with some common contextual stuff I write a lot

# Why?

1. Because I love `logrus` and I keep configuring it a similar way between my applications
2. Because I want some common context between my log lines

# So show me what I can do with it!

```go
import (
    log “github.com/GeorgeMac/logrush”
    “github.com/Sirupsen/logrus”
)

func main() {
    env := os.GetEnv(“GO_ENV”)
    if env == “” {
        env = “development”
    }

    fields := logrus.Fields{“environment”: env}

    // global logger
    log.Set(log.App(“my-service”),
        log.Common(fields),
        log.Formatter(&logrus.JSONFormatter{}),
        log.Level(logrus.DebugLevel))

    log.Printf(...)
    log.Debugf(...)

    // separate logger
    logger := log.New(“other-service”, 
        log.Common(fields),
        log.Level(logrus.InfoLevel))
}
```
