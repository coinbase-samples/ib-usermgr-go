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
	"errors"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type MockCognito struct{}

func (m *MockCognito) GetUser(
	ctx context.Context,
	params *cognitoidentityprovider.GetUserInput,
	optFns ...func(*cognitoidentityprovider.Options),
) (*cognitoidentityprovider.GetUserOutput, error) {
	token, _ := grpc_auth.AuthFromMD(ctx, "bearer")
	if token == "goodToken" {
		return &cognitoidentityprovider.GetUserOutput{
			UserAttributes: []types.AttributeType{},
		}, nil
	}
	if token == "badToken" {
		return nil, errors.New("Request unauthenticated with bearer")
	}
	return nil, nil
}

func TestMiddlewareEmptyContext(t *testing.T) {
	cip := MockCognito{}
	aw := Middleware{Cip: &cip}
	intercepter := aw.InterceptorNew()
	info := grpc.UnaryServerInfo{FullMethod: ""}
	resp, err := intercepter(
		context.Background(),
		&struct{}{},
		&info,
		func(_ context.Context, _ interface{}) (interface{}, error) {
			return &struct{}{}, nil
		},
	)

	if err == nil {
		t.Fatal("expected an error with empty context")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
}

func TestMiddlewareHealthCheckContext(t *testing.T) {
	cip := MockCognito{}
	aw := Middleware{Cip: &cip}
	intercepter := aw.InterceptorNew()
	info := grpc.UnaryServerInfo{FullMethod: "/grpc.health.v1.Health/Check"}
	resp, err := intercepter(
		context.Background(),
		&struct{}{},
		&info,
		func(_ context.Context, _ interface{}) (interface{}, error) {
			return &struct{}{}, nil
		},
	)

	if err != nil {
		t.Fatal("expected an error with empty context")
	}
	if resp == nil {
		t.Fatal("expected valid response")
	}
}

func TestMiddlewareUserContext(t *testing.T) {
	cip := MockCognito{}
	aw := Middleware{Cip: &cip}
	intercepter := aw.InterceptorNew()
	info := grpc.UnaryServerInfo{FullMethod: "/anyMethod"}
	md := metadata.Pairs("authorization", "bearer goodToken")
	ctx := metautils.NiceMD(md).ToIncoming(context.Background()) //.ToOutgoing(context.Background())
	resp, err := intercepter(
		ctx,
		&struct{}{},
		&info,
		func(_ context.Context, _ interface{}) (interface{}, error) {
			return &struct{}{}, nil
		},
	)

	if err != nil {
		t.Fatal("unexpected error with goodToken")
	}
	if resp == nil {
		t.Fatal("expected valid response")
	}
}

func TestMiddlewareBadUserContext(t *testing.T) {
	cip := MockCognito{}
	aw := Middleware{Cip: &cip}
	intercepter := aw.InterceptorNew()
	info := grpc.UnaryServerInfo{FullMethod: "/anyMethod"}
	md := metadata.Pairs("authorization", "bearer badToken")
	ctx := metautils.NiceMD(md).ToIncoming(context.Background()) //.ToOutgoing(context.Background())
	resp, err := intercepter(
		ctx,
		&struct{}{},
		&info,
		func(_ context.Context, _ interface{}) (interface{}, error) {
			return &struct{}{}, nil
		},
	)

	if !strings.Contains(err.Error(), "Request unauthenticated with bearer") {
		t.Fatal("expected error with badToken")
	}
	if resp != nil {
		t.Fatal("expected valid response")
	}
}
