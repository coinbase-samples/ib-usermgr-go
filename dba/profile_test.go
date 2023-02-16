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

package dba

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/coinbase-samples/ib-usermgr-go/config"
	"github.com/coinbase-samples/ib-usermgr-go/model"
)

func TestReadProfile(t *testing.T) {
	repo := new(MockRepository)
	NewTestDBA(repo)

	resp, err := repo.ReadProfile(ReadProfileFound)

	if err != nil {
		t.Fatal("unexpected error returned from function invocation")
	}

	if resp.Name != "Ted Robinson" {
		t.Fatal("expected name")
	}

	if resp.UserId != ReadProfileFound {
		t.Fatal("expected userId to match")
	}
}

func TestReadProfileNotFound(t *testing.T) {
	repo := new(MockRepository)
	NewTestDBA(repo)

	resp, err := repo.ReadProfile(ReadProfileNotFound)

	if len(resp.Name) > 0 {
		t.Fatal("expected empty name")
	}

	if len(resp.UserId) > 0 {
		t.Fatal("expected userId to be empty")
	}

	if err == errors.New("profile not found") {
		t.Fatal("expected error")
	}
}

func TestUpdateProfile(t *testing.T) {
	repo := new(MockRepository)
	NewTestDBA(repo)

	resp, err := repo.UpdateProfile(UpdateProfile, model.UpdateProfileRequest{
		UserId: UpdateProfile,
		Name:   "Bob Ross",
	})

	if err != nil {
		t.Fatal("expected error returned from function invocation")
	}

	if resp.Name != "Bob Ross" {
		t.Fatal("expected name update")
	}

	if resp.UserId != UpdateProfile {
		t.Fatal("expected userId to match")
	}
}

func TestUpdateProfileNotFound(t *testing.T) {
	repo := new(MockRepository)
	NewTestDBA(repo)

	resp, err := repo.UpdateProfile(UpdateProfileNotFound, model.UpdateProfileRequest{
		UserId: UpdateProfile,
		Name:   "Bob Ross",
	})

	if len(resp.Name) > 0 {
		t.Fatal("expected empty name")
	}

	if len(resp.UserId) > 0 {
		t.Fatal("expected userId to be empty")
	}

	if err == errors.New("profile not found") {
		t.Fatal("expected error")
	}
}

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

func TestReadDynamo(t *testing.T) {
	dynMock := new(DynamoMock)
	app := config.AppConfig{
		ProfileTableName: "Profile",
	}
	repo := &DynamoRepository{
		App: &app,
		Svc: dynMock,
	}
	NewDBA(repo)

	resp, err := Repo.ReadProfile(ReadProfileFound)

	if err != nil {
		t.Fatal("unexpected error")
	}

	if resp.Name != "Ted Robinson" {
		t.Fatal("expected read name")
	}
}

func TestUpdateDynamo(t *testing.T) {
	dynMock := new(DynamoMock)
	app := config.AppConfig{
		ProfileTableName: "Profile",
	}
	repo := &DynamoRepository{
		App: &app,
		Svc: dynMock,
	}
	NewDBA(repo)

	resp, err := Repo.UpdateProfile(UpdateProfile, model.UpdateProfileRequest{
		UserId: UpdateProfile,
		Name:   "Bob Ross",
	})

	if err != nil {
		t.Fatal("unexpected error")
	}

	if resp.Name != "Bob Ross" {
		t.Fatal("expected updated name")
	}
}

type DynamoErrorMock struct {
	dynamodbiface.DynamoDBAPI
}

func (m *DynamoErrorMock) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	return nil, errors.New("some error")
}

func (m *DynamoErrorMock) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	return nil, errors.New("some error")
}

func TestReadErrorDynamo(t *testing.T) {
	dynMock := new(DynamoErrorMock)
	app := config.AppConfig{
		ProfileTableName: "Profile",
	}
	repo := &DynamoRepository{
		App: &app,
		Svc: dynMock,
	}
	NewDBA(repo)

	_, err := Repo.ReadProfile(ReadProfileFound)

	if err == nil {
		t.Fatal("expected error")
	}

}

func TestUpdateErrorDynamo(t *testing.T) {
	dynMock := new(DynamoErrorMock)
	app := config.AppConfig{
		ProfileTableName: "Profile",
	}
	repo := &DynamoRepository{
		App: &app,
		Svc: dynMock,
	}
	NewDBA(repo)

	_, err := Repo.UpdateProfile(UpdateProfile, model.UpdateProfileRequest{
		UserId: UpdateProfile,
		Name:   "Bob Ross",
	})
	if err == nil {
		t.Fatal("expected error")
	}

}
