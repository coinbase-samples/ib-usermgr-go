package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func LogInit(app AppConfig) *log.Entry {
	logger := log.New()
	logLevel, _ := log.ParseLevel(app.LogLevel)
	logger.SetLevel(logLevel)
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetReportCaller(true)
	logger.SetOutput(os.Stdout)
	return log.NewEntry(logger)
}
