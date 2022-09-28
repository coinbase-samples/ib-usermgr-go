package main

import (
	"fmt"
	"time"

	"github.com/coinbase-samples/ib-usermgr-go/config"
	"github.com/coinbase-samples/ib-usermgr-go/dba"
	log "github.com/sirupsen/logrus"
)

var (
	//setup logrus for interceptor
	logrusLogger = log.New()
	wait         time.Duration
)

func main() {
	var app config.AppConfig

	config.Setup(&app)
	fmt.Println("starting app with config", app)
	config.LogInit(app)

	logrusLogger.SetFormatter(&log.JSONFormatter{})
	//setup cognito client
	cip := InitAuth(&app)
	aw := authMiddleware{cip} //setup dynamodb connection

	//setup database
	repo := dba.NewRepo(&app)
	dba.NewDBA(repo)

	// Start gRPC Server
	gRPCListen(app, aw)
}
