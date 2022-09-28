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
