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

package handlers

import (
	"context"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/coinbase-samples/ib-usermgr-go/config"
	"github.com/coinbase-samples/ib-usermgr-go/dba"
	"github.com/coinbase-samples/ib-usermgr-go/model"
	profile "github.com/coinbase-samples/ib-usermgr-go/pkg/pbs/profile/v1"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
)

var (
	ReadProfileNotFound   = "4AC6E407-1D8E-4339-BA1C-862ACC58AC5E"
	ReadProfileFound      = "20032259-738B-40A7-AAD7-306B69AF88D4"
	UpdateProfileNotFound = "E6096F2D-C706-42B6-B0E5-D7DD644ED079"
	UpdateProfile         = "AC032259-738B-40A7-AAD7-306B69AAB909"
)

type DynamoMock struct {
	dynamodbiface.DynamoDBAPI
}

func (m *DynamoMock) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	return &dynamodb.GetItemOutput{
		Item: map[string]types.AttributeValue{
			"UserId": &types.AttributeValueMemberS{Value: ReadProfileFound},
			"Name":   &types.AttributeValueMemberS{Value: "Ted Robinson"},
		},
	}, nil
}

func (m *DynamoMock) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	return &dynamodb.PutItemOutput{
		Attributes: map[string]types.AttributeValue{
			"UserId": &types.AttributeValueMemberS{Value: UpdateProfile},
			"Name":   &types.AttributeValueMemberS{Value: "Bob Ross"},
		},
	}, nil
}

func TestReadHandler(t *testing.T) {
	l := logrus.New()
	entry := logrus.NewEntry(l)
	ctx := context.TODO()
	ctx = ctxlogrus.ToContext(ctx, entry)
	ctx = context.WithValue(ctx, model.UserCtxKey, model.User{Id: "123"})

	dynMock := new(DynamoMock)
	app := config.AppConfig{
		ProfileTableName: "Profile",
	}
	repo := &dba.DynamoRepository{
		App: &app,
		Svc: dynMock,
	}
	dba.NewDBA(repo)

	ps := ProfileServer{}

	resp, err := ps.ReadProfile(ctx, &profile.ReadProfileRequest{
		Id: ReadProfileFound,
	})

	if err != nil {
		t.Fatal("unexpected read profile error")
	}

	if resp.UserId != ReadProfileFound {
		t.Fatal("unexpected user id returned")
	}
}

func TestReadHandlerInvalidId(t *testing.T) {
	l := logrus.New()
	entry := logrus.NewEntry(l)
	ctx := context.TODO()
	ctx = ctxlogrus.ToContext(ctx, entry)
	ctx = context.WithValue(ctx, model.UserCtxKey, model.User{Id: "123"})

	dynMock := new(DynamoMock)
	app := config.AppConfig{
		ProfileTableName: "Profile",
	}
	repo := &dba.DynamoRepository{
		App: &app,
		Svc: dynMock,
	}
	dba.NewDBA(repo)

	ps := ProfileServer{}

	_, err := ps.ReadProfile(ctx, &profile.ReadProfileRequest{
		Id: "short",
	})

	if err == nil {
		t.Fatal("expected read profile error for invalid id")
	}

	validationError := "invalid ReadProfileRequest.Id: value length must be 36 runes"
	if !strings.Contains(err.Error(), validationError) {
		t.Fatalf("incorrect validation, expected - %s got - %s", validationError, err)
	}

}

func TestUpdateHandler(t *testing.T) {
	l := logrus.New()
	entry := logrus.NewEntry(l)
	ctx := context.TODO()
	ctx = ctxlogrus.ToContext(ctx, entry)
	ctx = context.WithValue(ctx, model.UserCtxKey, model.User{Id: "123"})

	dynMock := new(DynamoMock)
	app := config.AppConfig{
		ProfileTableName: "Profile",
	}
	repo := &dba.DynamoRepository{
		App: &app,
		Svc: dynMock,
	}
	dba.NewDBA(repo)

	ps := ProfileServer{}

	resp, err := ps.UpdateProfile(ctx, &profile.UpdateProfileRequest{
		Id:          UpdateProfile,
		Name:        "Bob Ross",
		Email:       "b.ross@coinbase.com",
		LegalName:   "Bob Ross",
		UserName:    "demo0",
		Address:     "123 Happy Way",
		DateOfBirth: "The best day",
	})

	if err != nil {
		t.Fatal("unexpected update profile error")
	}

	if resp.UserId != UpdateProfile {
		t.Fatalf("unexpected user id returned, got %s expected %s", resp.UserId, UpdateProfile)
	}

	if resp.Name != "Bob Ross" {
		t.Fatal("expected name to update")
	}
}

func TestUpdateHandlerEmailError(t *testing.T) {
	l := logrus.New()
	entry := logrus.NewEntry(l)
	ctx := context.TODO()
	ctx = ctxlogrus.ToContext(ctx, entry)
	ctx = context.WithValue(ctx, model.UserCtxKey, model.User{Id: "123"})

	dynMock := new(DynamoMock)
	app := config.AppConfig{
		ProfileTableName: "Profile",
	}
	repo := &dba.DynamoRepository{
		App: &app,
		Svc: dynMock,
	}
	dba.NewDBA(repo)

	ps := ProfileServer{}

	_, err := ps.UpdateProfile(ctx, &profile.UpdateProfileRequest{
		Id:          UpdateProfile,
		Name:        "Bob Ross",
		Email:       "nope",
		LegalName:   "Bob Ross",
		UserName:    "demo0",
		Address:     "123 Happy Way",
		DateOfBirth: "The best day",
	})

	emailError := "invalid UpdateProfileRequest.Email: value must be a valid email address | caused by: mail: missing '@' or angle-addr"
	if !strings.Contains(err.Error(), emailError) {
		t.Fatalf("incorrect validation, expected - %s got - %s", emailError, err)
	}

}
