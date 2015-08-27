package logrush

func Panic(v ...interface{})            { logger.Panic(v...) }
func Panicf(f string, v ...interface{}) { logger.Panicf(f, v...) }
func Panicln(v ...interface{})          { logger.Panicln(v...) }

func Fatal(v ...interface{})            { logger.Fatal(v...) }
func Fatalf(f string, v ...interface{}) { logger.Fatalf(f, v...) }
func Fatalln(v ...interface{})          { logger.Fatalln(v...) }

func Error(v ...interface{})            { logger.Error(v...) }
func Errorf(f string, v ...interface{}) { logger.Errorf(f, v...) }
func Errorln(v ...interface{})          { logger.Errorln(v...) }

func Warning(v ...interface{})            { logger.Warning(v...) }
func Warningf(f string, v ...interface{}) { logger.Warningf(f, v...) }
func Warningln(v ...interface{})          { logger.Warningln(v...) }

func Info(v ...interface{})            { logger.Info(v...) }
func Infof(f string, v ...interface{}) { logger.Infof(f, v...) }
func Infoln(v ...interface{})          { logger.Infoln(v...) }

func Debug(v ...interface{})            { logger.Debug(v...) }
func Debugf(f string, v ...interface{}) { logger.Debugf(f, v...) }
func Debugln(v ...interface{})          { logger.Debugln(v...) }

// LogRush type specific

func (l *LogRush) Panic(v ...interface{})            { l.log().Panic(v...) }
func (l *LogRush) Panicf(f string, v ...interface{}) { l.log().Panicf(f, v...) }
func (l *LogRush) Panicln(v ...interface{})          { l.log().Panicln(v...) }

func (l *LogRush) Fatal(v ...interface{})            { l.log().Fatal(v...) }
func (l *LogRush) Fatalf(f string, v ...interface{}) { l.log().Fatalf(f, v...) }
func (l *LogRush) Fatalln(v ...interface{})          { l.log().Fatalln(v...) }

func (l *LogRush) Error(v ...interface{})            { l.log().Error(v...) }
func (l *LogRush) Errorf(f string, v ...interface{}) { l.log().Errorf(f, v...) }
func (l *LogRush) Errorln(v ...interface{})          { l.log().Errorln(v...) }

func (l *LogRush) Warning(v ...interface{})            { l.log().Warning(v...) }
func (l *LogRush) Warningf(f string, v ...interface{}) { l.log().Warningf(f, v...) }
func (l *LogRush) Warningln(v ...interface{})          { l.log().Warningln(v...) }

func (l *LogRush) Info(v ...interface{})            { l.log().Info(v...) }
func (l *LogRush) Infof(f string, v ...interface{}) { l.log().Infof(f, v...) }
func (l *LogRush) Infoln(v ...interface{})          { l.log().Infoln(v...) }

func (l *LogRush) Debug(v ...interface{})            { l.log().Debug(v...) }
func (l *LogRush) Debugf(f string, v ...interface{}) { l.log().Debugf(f, v...) }
func (l *LogRush) Debugln(v ...interface{})          { l.log().Debugln(v...) }
