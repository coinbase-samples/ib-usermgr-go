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

package handlers

import (
	"context"
	"errors"
	"fmt"

	"github.com/coinbase-samples/ib-usermgr-go/conversions"
	"github.com/coinbase-samples/ib-usermgr-go/dba"
	"github.com/coinbase-samples/ib-usermgr-go/model"
	profile "github.com/coinbase-samples/ib-usermgr-go/pkg/pbs/profile/v1"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
)

type ProfileServer struct {
	profile.UnimplementedProfileServiceServer
}

func (o *ProfileServer) ReadProfile(ctx context.Context, req *profile.ReadProfileRequest) (*profile.ReadProfileResponse, error) {
	l := ctxlogrus.Extract(ctx)
	authedUser := ctx.Value(model.UserCtxKey).(model.User)
	if err := req.ValidateAll(); err != nil {
		l.Debugln("invalid read profile request", err)
		return nil, fmt.Errorf("profile handler could not validate request: %w", err)
	}

	l.Debugf("fetching user - %s - %s", authedUser.Id, req.Id)
	body, err := dba.Repo.ReadProfile(authedUser.Id)

	if err != nil {
		l.Debugln("error reading profile with dynamo", err)
		return nil, fmt.Errorf("profile handler could not read profile: %w", err)
	}

	response := conversions.ConvertReadProfileToProto(body)

	l.Debugf("returning read profile response - %v", &response)
	return &response, nil
}

func (o *ProfileServer) UpdateProfile(ctx context.Context, req *profile.UpdateProfileRequest) (*profile.UpdateProfileResponse, error) {
	l := ctxlogrus.Extract(ctx)
	authedUser := ctx.Value(model.UserCtxKey).(model.User)
	if err := req.ValidateAll(); err != nil {
		l.Debugln("invalid update profile request", err)
		return nil, fmt.Errorf("profile handler could not validate request: %w", err)
	}

	updateBody := conversions.ConvertUpdateProfileToModel(req)

	l.Debugln("updating user", authedUser.Id, req.Id)
	body, err := dba.Repo.UpdateProfile(authedUser.Id, updateBody)

	if err != nil {
		l.Debugln("error updating profile with dynamo", err)
		return nil, fmt.Errorf("profile handler could not update profile: %w", err)
	}

	response := conversions.ConvertUpdateProfileToProto(body)

	l.Debugf("returning update profile response - %v", &response)
	return &response, nil
}

func (o *ProfileServer) CreateProfile(ctx context.Context, req *profile.CreateProfileRequest) (*profile.CreateProfileResponse, error) {
	return nil, errors.New("method not implemented")
}
