/**
 * Copyright 2022-present Coinbase Global, Inc.
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

package auth

import (
	"context"
	"testing"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/coinbase-samples/ib-usermgr-go/config"
)

func TestAuthClient(t *testing.T) {
	app := config.AppConfig{
		UserPoolId: "poolId",
		ClientId:   "clientId",
	}
	cfg, _ := awsConfig.LoadDefaultConfig(context.Background())

	cognitoClient := InitAuth(&app, cfg)

	if cognitoClient == nil {
		t.Fatal("could not create cognito client")
	}
}
