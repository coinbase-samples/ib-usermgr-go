/**
 * Copyright 2022 Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"time"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/coinbase-samples/ib-usermgr-go/auth"
	"github.com/coinbase-samples/ib-usermgr-go/config"
	"github.com/coinbase-samples/ib-usermgr-go/dba"
	"github.com/coinbase-samples/ib-usermgr-go/log"
)

var (
	wait time.Duration
)

func main() {
	var app config.AppConfig

	config.Setup(&app)
	log.Init(app)
	log.Debugf("starting app with config - %v", app)

	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := awsConfig.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	// Setup cognito client
	cip := auth.InitAuth(&app, cfg)
	aw := auth.Middleware{Cip: cip}

	// Setup database
	repo := dba.NewRepo(&app, cfg)
	dba.NewDBA(repo)

	// Start gRPC Server
	gRPCListen(app, aw)
}
