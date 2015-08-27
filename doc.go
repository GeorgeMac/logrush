// LogRush logrus wrapper
//
// Examples:
//
//	logger := logrush.New("my-app")
//
// This will returns a new LogRush with the field {"app": "my-app"}
//
//	logger := logrush.New("my-service", logrush.AppKey("service"), logrush.Level(logrus.InfoLevel))
//
// This will return a new LogRush with field {"service": "my-service"} with the logging
// level set to INFO.
//
//	logger := logrush.New("my-app",
//		logrush.Formatter(&logrus.JSONFomatter{}),
//		logrush.Common(logrus.Fields{"env":"development"}))
//
// This will return a new LogRush with fields {"app": "my-app", "env": "development"} and output
// using logrus JSON formatting for logging via logstash.
//
// If you just want to configure the global logrush logger. This can be done via the logrush.Set function
//
//	logrush.Set(logrush.App("my-app"), logrush.Level(logrus.DebugLevel))
//	logrush.Debugf("Some %s", "Information")
//
// Will result is a log line with the fields {"app": "my-app", level:"DEBUG", msg:"Some Information"}
package logrush
