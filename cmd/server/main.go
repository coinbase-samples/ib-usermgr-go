package main

import (
	"fmt"
	"time"

	"github.com/coinbase-samples/ib-usermgr-go/config"
	"github.com/coinbase-samples/ib-usermgr-go/dba"
)

var (
	wait time.Duration
)

func main() {
	var app config.AppConfig

	config.Setup(&app)
	fmt.Println("starting app with config", app)
	logrusLogger := config.LogInit(app)

	//setup cognito client
	cip := InitAuth(&app)
	aw := authMiddleware{cip} //setup dynamodb connection

	//setup database
	repo := dba.NewRepo(&app)
	dba.NewDBA(repo)

	// Start gRPC Server
	gRPCListen(app, aw, logrusLogger)
}
