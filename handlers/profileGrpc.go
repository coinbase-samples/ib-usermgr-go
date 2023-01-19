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
	"github.com/coinbase-samples/ib-usermgr-go/log"
	"github.com/coinbase-samples/ib-usermgr-go/model"
	profile "github.com/coinbase-samples/ib-usermgr-go/pkg/pbs/profile/v1"
)

type ProfileServer struct {
	profile.UnimplementedProfileServiceServer
}

func (o *ProfileServer) ReadProfile(ctx context.Context, req *profile.ReadProfileRequest) (*profile.ReadProfileResponse, error) {
	authedUser := ctx.Value(model.UserCtxKey).(model.User)
	if err := req.ValidateAll(); err != nil {
		log.DebugfCtx(ctx, "invalid read profile request: %v", err)
		return nil, fmt.Errorf("profile handler could not validate request: %w", err)
	}

	log.DebugfCtx(ctx, "fetching user - %s - %s", authedUser.Id, req.Id)
	body, err := dba.Repo.ReadProfile(authedUser.Id)

	if err != nil {
		log.DebugfCtx(ctx, "error reading profile with dynamo: %v", err)
		return nil, fmt.Errorf("profile handler could not read profile: %w", err)
	}

	response := conversions.ConvertReadProfileToProto(body)

	log.DebugfCtx(ctx, "returning read profile response - %v", &response)
	return &response, nil
}

func (o *ProfileServer) UpdateProfile(ctx context.Context, req *profile.UpdateProfileRequest) (*profile.UpdateProfileResponse, error) {
	authedUser := ctx.Value(model.UserCtxKey).(model.User)
	if err := req.ValidateAll(); err != nil {
		log.DebugfCtx(ctx, "invalid update profile request: %v", err)
		return nil, fmt.Errorf("profile handler could not validate request: %w", err)
	}

	updateBody := conversions.ConvertUpdateProfileToModel(req)

	log.DebugfCtx(ctx, "updating user: %s - %s", authedUser.Id, req.Id)
	body, err := dba.Repo.UpdateProfile(authedUser.Id, updateBody)

	if err != nil {
		log.DebugfCtx(ctx, "error updating profile with dynamo", err)
		return nil, fmt.Errorf("profile handler could not update profile: %w", err)
	}

	response := conversions.ConvertUpdateProfileToProto(body)

	log.DebugfCtx(ctx, "returning update profile response - %v", &response)
	return &response, nil
}

func (o *ProfileServer) CreateProfile(ctx context.Context, req *profile.CreateProfileRequest) (*profile.CreateProfileResponse, error) {
	return nil, errors.New("method not implemented")
}
