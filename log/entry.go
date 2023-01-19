package log

import "github.com/sirupsen/logrus"

func (e *Entry) GetUnderneath() *logrus.Entry {
	return e.l
}

func (e *Entry) Tracef(format string, args ...interface{}) {
	e.l.Tracef(format, args...)
}

func (e *Entry) Trace(args ...interface{}) {
	e.l.Trace(args...)
}

func (e *Entry) Debugf(format string, args ...interface{}) {
	e.l.Debugf(format, args...)
}

func (e *Entry) Debug(args ...interface{}) {
	e.l.Debug(args...)
}

func (e *Entry) Infof(format string, args ...interface{}) {
	e.l.Infof(format, args...)
}

func (e *Entry) Info(args ...interface{}) {
	e.l.Info(args...)
}

func (e *Entry) Warnf(format string, args ...interface{}) {
	e.l.Warnf(format, args...)
}

func (e *Entry) Warn(args ...interface{}) {
	e.l.Warn(args...)
}

func (e *Entry) Errorf(format string, args ...interface{}) {
	e.l.Errorf(format, args...)
}

func (e *Entry) Error(args ...interface{}) {
	e.l.Error(args...)
}

func (e *Entry) Fatalf(format string, args ...interface{}) {
	e.l.Fatalf(format, args...)
}

func (e *Entry) Fatal(args ...interface{}) {
	e.l.Fatal(args...)
}

func (e *Entry) Panicf(format string, args ...interface{}) {
	e.l.Panicf(format, args...)
}

func (e *Entry) Panic(args ...interface{}) {
	e.l.Panic(args...)
}

func (e *Entry) WithField(key string, value interface{}) *Entry {
	e.l = e.l.WithField(key, value)
	return e
}

func (e *Entry) WithFields(fields map[string]interface{}) *Entry {
	e.l = e.l.WithFields(fields)
	return e
}
