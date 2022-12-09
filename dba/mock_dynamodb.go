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
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/coinbase-samples/ib-usermgr-go/config"
	"github.com/coinbase-samples/ib-usermgr-go/model"
)

var TestRepo *MockRepository

type MockRepository struct {
	App *config.AppConfig
	Svc *dynamodb.Client
}

func NewTestDBA(r *MockRepository) {
	TestRepo = r
}

var (
	ReadProfileNotFound   = "4AC6E407-1D8E-4339-BA1C-862ACC58AC5E"
	ReadProfileFound      = "20032259-738B-40A7-AAD7-306B69AF88D4"
	UpdateProfileNotFound = "E6096F2D-C706-42B6-B0E5-D7DD644ED079"
	UpdateProfile         = "AC032259-738B-40A7-AAD7-306B69AAB909"
)

func (m *MockRepository) ReadProfile(id string) (model.ProfileResponse, error) {
	if id == ReadProfileNotFound {
		return model.ProfileResponse{}, errors.New("profile not found")
	}
	return model.ProfileResponse{Name: "Ted Robinson", UserId: id}, nil
}

func (m *MockRepository) UpdateProfile(id string, updateBody model.UpdateProfileRequest) (model.ProfileResponse, error) {
	if id == UpdateProfileNotFound {
		return model.ProfileResponse{}, errors.New("profile not found")
	}
	return model.ProfileResponse{Name: updateBody.Name, UserId: id}, nil
}
