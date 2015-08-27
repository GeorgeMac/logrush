package logrush

// Panic shhh golint
func Panic(v ...interface{}) { logger.Panic(v...) }

// Panicf shhh golint
func Panicf(f string, v ...interface{}) { logger.Panicf(f, v...) }

// Panicln shhh golint
func Panicln(v ...interface{}) { logger.Panicln(v...) }

// Fatal shhh golint
func Fatal(v ...interface{}) { logger.Fatal(v...) }

// Fatalf shhh golint
func Fatalf(f string, v ...interface{}) { logger.Fatalf(f, v...) }

// Fatalln shhh golint
func Fatalln(v ...interface{}) { logger.Fatalln(v...) }

// Error shhh golint
func Error(v ...interface{}) { logger.Error(v...) }

// Errorf shhh golint
func Errorf(f string, v ...interface{}) { logger.Errorf(f, v...) }

// Errorln shhh golint
func Errorln(v ...interface{}) { logger.Errorln(v...) }

// Warning shhh golint
func Warning(v ...interface{}) { logger.Warning(v...) }

// Warningf shhh golint
func Warningf(f string, v ...interface{}) { logger.Warningf(f, v...) }

// Warningln shhh golint
func Warningln(v ...interface{}) { logger.Warningln(v...) }

// Info shhh golint
func Info(v ...interface{}) { logger.Info(v...) }

// Infof shhh golint
func Infof(f string, v ...interface{}) { logger.Infof(f, v...) }

// Infoln shhh golint
func Infoln(v ...interface{}) { logger.Infoln(v...) }

// Debug shhh golint
func Debug(v ...interface{}) { logger.Debug(v...) }

// Debugf shhh golint
func Debugf(f string, v ...interface{}) { logger.Debugf(f, v...) }

// Debugln shhh golint
func Debugln(v ...interface{}) { logger.Debugln(v...) }

// LogRush type specific

// Panic shhh golint
func (l *LogRush) Panic(v ...interface{}) { l.log().Panic(v...) }

// Panicf shhh golint
func (l *LogRush) Panicf(f string, v ...interface{}) { l.log().Panicf(f, v...) }

// Panicln shhh golint
func (l *LogRush) Panicln(v ...interface{}) { l.log().Panicln(v...) }

// Fatal shhh golint
func (l *LogRush) Fatal(v ...interface{}) { l.log().Fatal(v...) }

// Fatalf shhh golint
func (l *LogRush) Fatalf(f string, v ...interface{}) { l.log().Fatalf(f, v...) }

// Fatalln shhh golint
func (l *LogRush) Fatalln(v ...interface{}) { l.log().Fatalln(v...) }

// Error shhh golint
func (l *LogRush) Error(v ...interface{}) { l.log().Error(v...) }

// Errorf shhh golint
func (l *LogRush) Errorf(f string, v ...interface{}) { l.log().Errorf(f, v...) }

// Errorln shhh golint
func (l *LogRush) Errorln(v ...interface{}) { l.log().Errorln(v...) }

// Warning shhh golint
func (l *LogRush) Warning(v ...interface{}) { l.log().Warning(v...) }

// Warningf shhh golint
func (l *LogRush) Warningf(f string, v ...interface{}) { l.log().Warningf(f, v...) }

// Warningln shhh golint
func (l *LogRush) Warningln(v ...interface{}) { l.log().Warningln(v...) }

// Info shhh golint
func (l *LogRush) Info(v ...interface{}) { l.log().Info(v...) }

// Infof shhh golint
func (l *LogRush) Infof(f string, v ...interface{}) { l.log().Infof(f, v...) }

// Infoln shhh golint
func (l *LogRush) Infoln(v ...interface{}) { l.log().Infoln(v...) }

// Debug shhh golint
func (l *LogRush) Debug(v ...interface{}) { l.log().Debug(v...) }

// Debugf shhh golint
func (l *LogRush) Debugf(f string, v ...interface{}) { l.log().Debugf(f, v...) }

// Debugln shhh golint
func (l *LogRush) Debugln(v ...interface{}) { l.log().Debugln(v...) }
