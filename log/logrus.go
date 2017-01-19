package log

import (
	"encoding/json"
	"io"
	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	glog "github.com/labstack/gommon/log"
)

type logrusLog struct {
	sync.Mutex
	logger *logrus.Logger
	prefix string
	off    bool
}

func (l *logrusLog) Output() io.Writer {
	return l.logger.Out
}

func (l *logrusLog) SetOutput(out io.Writer) {
	l.Lock()
	defer l.Unlock()
	l.logger.Out = out
}

func (l *logrusLog) Level() glog.Lvl {
	switch l.logger.Level {
	case logrus.InfoLevel:
		return glog.INFO
	case logrus.WarnLevel:
		return glog.WARN
	case logrus.ErrorLevel:
		return glog.ERROR
	}
	if l.off {
		return glog.OFF
	}
	return glog.DEBUG
}

func (l *logrusLog) SetLevel(lvl glog.Lvl) {
	l.Lock()
	defer l.Unlock()
	logLevel := logrus.DebugLevel
	l.off = false
	switch lvl {
	case glog.INFO:
		logLevel = logrus.InfoLevel
	case glog.WARN:
		logLevel = logrus.WarnLevel
	case glog.ERROR:
		logLevel = logrus.ErrorLevel
	case glog.OFF:
		l.off = true
	}
	l.logger.Level = logLevel
}
func (l *logrusLog) Prefix() string {
	return l.prefix
}

func (l *logrusLog) SetPrefix(p string) {
	// TODO implement logging with prefix feature
	l.prefix = p
}

func (l *logrusLog) Print(i ...interface{}) {
	if !l.off {
		l.logger.Print(i...)
	}
}

func (l *logrusLog) Printf(format string, i ...interface{}) {
	if !l.off {
		l.logger.Printf(format, i...)
	}
}

func (l *logrusLog) Printj(j glog.JSON) {
	if !l.off {
		b, _ := json.Marshal(j)
		l.logger.Print(string(b))
	}
}

func (l *logrusLog) Debug(i ...interface{}) {
	if !l.off {
		l.logger.Debug(i...)
	}
}

func (l *logrusLog) Debugf(format string, i ...interface{}) {
	if !l.off {
		l.logger.Debugf(format, i...)
	}
}

func (l *logrusLog) Debugj(j glog.JSON) {
	if !l.off {
		b, _ := json.Marshal(j)
		l.logger.Debug(string(b))
	}
}

func (l *logrusLog) Info(i ...interface{}) {
	if !l.off {
		l.logger.Info(i...)
	}
}

func (l *logrusLog) Infof(format string, i ...interface{}) {
	if !l.off {
		l.logger.Infof(format, i...)
	}
}

func (l *logrusLog) Infoj(j glog.JSON) {
	if !l.off {
		b, _ := json.Marshal(j)
		l.logger.Info(string(b))
	}
}

func (l *logrusLog) Warn(i ...interface{}) {
	if !l.off {
		l.logger.Warn(i...)
	}
}

func (l *logrusLog) Warnf(format string, i ...interface{}) {
	if !l.off {
		l.logger.Warnf(format, i...)
	}
}

func (l *logrusLog) Warnj(j glog.JSON) {
	if !l.off {
		b, _ := json.Marshal(j)
		l.logger.Warn(string(b))
	}
}

func (l *logrusLog) Error(i ...interface{}) {
	if !l.off {
		l.logger.Error(i...)
	}
}

func (l *logrusLog) Errorf(format string, i ...interface{}) {
	if !l.off {
		l.logger.Errorf(format, i...)
	}
}

func (l *logrusLog) Errorj(j glog.JSON) {
	if !l.off {
		b, _ := json.Marshal(j)
		l.logger.Error(string(b))
	}
}

func (l *logrusLog) Fatal(i ...interface{}) {
	if !l.off {
		l.logger.Fatal(i...)
	}
}

func (l *logrusLog) Fatalf(format string, i ...interface{}) {
	if !l.off {
		l.logger.Fatalf(format, i...)
	}
}

func (l *logrusLog) Fatalj(j glog.JSON) {
	if !l.off {
		b, _ := json.Marshal(j)
		l.logger.Fatal(string(b))
	}
}

func (l *logrusLog) Panic(i ...interface{}) {
	if !l.off {
		l.logger.Panic(i...)
	}
}

func (l *logrusLog) Panicf(format string, i ...interface{}) {
	if !l.off {
		l.logger.Panicf(format, i...)
	}
}

func (l *logrusLog) Panicj(j glog.JSON) {
	if !l.off {
		b, _ := json.Marshal(j)
		l.logger.Panic(string(b))
	}
}

// Logrus returns logger with default logrus instance
func Logrus() echo.Logger {
	return &logrusLog{
		logger: logrus.StandardLogger(),
		off:    false,
	}
}

// LogrusFromLogger returns logger with custom logrus instance
func LogrusFromLogger(logger *logrus.Logger) echo.Logger {
	return &logrusLog{
		logger: logger,
		off:    false,
	}
}

var _ echo.Logger = &logrusLog{}
