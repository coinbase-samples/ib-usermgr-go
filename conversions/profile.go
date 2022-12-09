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

package conversions

import (
	"github.com/coinbase-samples/ib-usermgr-go/model"
	profile "github.com/coinbase-samples/ib-usermgr-go/pkg/pbs/profile/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertReadProfileToProto(p model.ProfileResponse) profile.ReadProfileResponse {
	return profile.ReadProfileResponse{
		UserId:      p.UserId,
		Email:       p.Email,
		Name:        p.Name,
		LegalName:   p.LegalName,
		UserName:    p.UserName,
		Roles:       p.Roles,
		Address:     p.Address,
		DateOfBirth: p.DateOfBirth,
		CreatedAt:   timestamppb.New(p.CreatedAt),
		UpdatedAt:   timestamppb.New(p.UpdatedAt),
	}
}

func ConvertUpdateProfileToModel(p *profile.UpdateProfileRequest) model.UpdateProfileRequest {
	return model.UpdateProfileRequest{
		UserId:      p.Id,
		Email:       p.Email,
		Name:        p.Name,
		LegalName:   p.LegalName,
		UserName:    p.UserName,
		Address:     p.Address,
		DateOfBirth: p.DateOfBirth,
	}
}

func ConvertUpdateProfileToProto(p model.ProfileResponse) profile.UpdateProfileResponse {
	return profile.UpdateProfileResponse{
		UserId:      p.UserId,
		Email:       p.Email,
		Name:        p.Name,
		LegalName:   p.LegalName,
		UserName:    p.UserName,
		Roles:       p.Roles,
		Address:     p.Address,
		DateOfBirth: p.DateOfBirth,
		CreatedAt:   timestamppb.New(p.CreatedAt),
		UpdatedAt:   timestamppb.New(p.UpdatedAt),
	}
}
