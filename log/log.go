package log

// TODO move me to shared util package

import (
	"fmt"
	"os"

	"github.com/coinbase-samples/ib-usermgr-go/config"
	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.New()
)

func Init(app config.AppConfig) {
	logLevel, _ := logrus.ParseLevel(app.LogLevel)
	logger.SetLevel(logLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: caller(),
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyFile: "caller",
		},
	})
	if app.LogToFile == "true" {
		// open a file
		f, err := os.OpenFile("testing.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			fmt.Printf("error opening file: %v", err)
		}
		logger.SetOutput(f)
	} else {
		logger.SetOutput(os.Stdout)
	}
}

func NewEntry() *Entry {
	return &Entry{l: logrus.NewEntry(logger)}
}

func Tracef(format string, args ...interface{}) {
	logger.Tracef(format, args...)
}

func Trace(args ...interface{}) {
	logger.Trace(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}
