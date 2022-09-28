package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/coinbase-samples/ib-usermgr-go/model"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"google.golang.org/grpc"
)

type authMiddleware struct {
	Cip *CognitoClient
}

func (am *authMiddleware) InterceptorNew() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// allow health checks to pass through
		if info.FullMethod == "/grpc.health.v1.Health/Check" || info.FullMethod == "/grpc.health.v1.Health/Watch" {

			return handler(ctx, req)
		}
		l := ctxlogrus.Extract(ctx)

		var newCtx context.Context
		var err error
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		l.Debugln("checking token", info.FullMethod, token, am.Cip.UserPoolId, am.Cip.AppClientId)

		if err != nil {
			return nil, err
		}

		user, err := am.Cip.GetUser(ctx, &cognitoidentityprovider.GetUserInput{
			AccessToken: aws.String(token),
		})

		if err != nil {
			l.Warnf("could not validate token - %v", err)
			return nil, err
		}

		var authedUser = model.User{}
		for _, attr := range user.UserAttributes {
			if *attr.Name == "sub" {
				authedUser.Id = *attr.Value
			} else if *attr.Name == "email" {
				authedUser.Email = *attr.Value
			}
		}
		newCtx = context.WithValue(ctx, model.UserCtxKey, authedUser)

		if err != nil {
			l.Warnf("could not add token to context - %v", err)
			return nil, err
		}
		return handler(newCtx, req)
	}
}
