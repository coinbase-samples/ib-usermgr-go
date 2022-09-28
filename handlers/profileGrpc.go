package handlers

import (
	"context"

	"github.com/coinbase-samples/ib-usermgr-go/conversions"
	"github.com/coinbase-samples/ib-usermgr-go/dba"
	"github.com/coinbase-samples/ib-usermgr-go/model"
	profile "github.com/coinbase-samples/ib-usermgr-go/pkg/pbs/profile/v1"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type ProfileServer struct {
	profile.UnimplementedProfileServiceServer
	Tracer trace.Tracer
}

func (o *ProfileServer) ReadProfile(ctx context.Context, req *profile.ReadProfileRequest) (*profile.ReadProfileResponse, error) {
	l := ctxlogrus.Extract(ctx)
	authedUser := ctx.Value(model.UserCtxKey).(model.User)
	if err := req.Validate(); err != nil {
		l.Debugln("invalid request", err)
		return nil, err
	}
	l.Debugln("starting trace")
	_, span := o.Tracer.Start(ctx, "readProfile",
		trace.WithAttributes(attribute.String("UserId", authedUser.Id)))
	defer span.End()

	l.Warnf("fetching user - %s - %s", authedUser.Id, req.Id)
	body, err := dba.Repo.ReadProfile(authedUser.Id)

	if err != nil {
		l.Debugln("error reading profile with dynamo", err)
	}

	response := conversions.ConvertReadProfileToProto(body)

	return &response, nil
}

func (o *ProfileServer) UpdateProfile(ctx context.Context, req *profile.UpdateProfileRequest) (*profile.UpdateProfileResponse, error) {
	l := ctxlogrus.Extract(ctx)
	authedUser := ctx.Value(model.UserCtxKey).(model.User)
	if err := req.Validate(); err != nil {
		l.Debugln("invalid request", err)
		return nil, err
	}
	l.Debugln("starting trace")
	_, span := o.Tracer.Start(ctx, "updateProfile",
		trace.WithAttributes(attribute.String("UserId", authedUser.Id)))
	defer span.End()

	updateBody := conversions.ConvertUpdateProfileToModel(req)

	l.Debugln("updating user", authedUser.Id, req.Id)
	_, err := dba.Repo.UpdateProfile(authedUser.Id, updateBody)

	if err != nil {
		l.Debugln("error updating profile with dynamo", err)
	}

	body, err := dba.Repo.ReadProfile((authedUser.Id))
	if err != nil {
		l.Debugln("error reading profile after update with dynamo", err)
	}
	response := conversions.ConvertUpdateProfileToProto(body)

	return &response, nil
}

/*
func readProfile(w http.ResponseWriter, r *http.Request, id string) {
	body, err := dba.Repo.ReadProfile(id)

	if err != nil {
		fmt.Println("error reading profile with dynamo", err)
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(body)
	io.WriteString(w, string(response))
}

func updateProfile(w http.ResponseWriter, r *http.Request, id string) {
	var updateBody model.UpdateProfileRequest
	err := json.NewDecoder(r.Body).Decode(&updateBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	body, err := dba.Repo.UpdateProfile(id, updateBody)

	if err != nil {
		fmt.Println("error updating profile with dynamo", err)
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(body)
	io.WriteString(w, string(response))
}
*/
