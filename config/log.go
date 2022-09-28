package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func LogInit(app AppConfig) {
	logLevel, _ := log.ParseLevel(app.LogLevel)
	log.SetLevel(logLevel)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
}
