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

package dba

import (
	"context"
	"testing"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/coinbase-samples/ib-usermgr-go/config"
)

func TestSetup(t *testing.T) {
	app := config.AppConfig{
		BaseConfig: config.BaseConfig{
			Env: "local",
		},
		DatabaseEndpoint: "http://localhost:4566",
	}
	cfg, _ := awsConfig.LoadDefaultConfig(context.Background())

	dr := NewRepo(&app, cfg)

	if dr.App.IsLocalEnv() {
		t.Fatal("did not correctly set config")
	}

	if dr.Svc == nil {
		t.Fatal("did not generate service")
	}
}

func TestStageSetup(t *testing.T) {
	app := config.AppConfig{
		BaseConfig: config.BaseConfig{
			Env: "stage",
		},
		DatabaseEndpoint: "http://localhost:4566",
	}
	cfg, _ := awsConfig.LoadDefaultConfig(context.Background())

	dr := NewRepo(&app, cfg)

	if dr.App.Env != "stage" {
		t.Fatal("did not correctly set config")
	}

	if dr.Svc == nil {
		t.Fatal("did not generate service")
	}
}

func TestNewDBA(t *testing.T) {
	app := config.AppConfig{
		BaseConfig: config.BaseConfig{
			Env: "local",
		},
		DatabaseEndpoint: "http://localhost:4566",
	}
	repo := &DynamoRepository{
		App: &app,
		Svc: nil,
	}

	NewDBA(repo)

	if !Repo.App.IsLocalEnv() {
		t.Fatal("config not set on repo")
	}

	if Repo.Svc != nil {
		t.Fatal("service expected nil")
	}
}
