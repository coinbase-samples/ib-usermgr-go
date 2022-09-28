package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	appConfig "github.com/coinbase-samples/ib-usermgr-go/config"
)

type CognitoClient struct {
	AppClientId string
	UserPoolId  string
	*cip.Client
}

func InitAuth(a *appConfig.AppConfig) *CognitoClient {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	return &CognitoClient{
		a.ClientId,
		a.UserPoolId,
		cip.NewFromConfig(cfg),
	}
}
