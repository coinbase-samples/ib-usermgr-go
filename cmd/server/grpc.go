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

package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/coinbase-samples/ib-usermgr-go/auth"
	"github.com/coinbase-samples/ib-usermgr-go/config"
	"github.com/coinbase-samples/ib-usermgr-go/handlers"
	"github.com/coinbase-samples/ib-usermgr-go/log"
	v1 "github.com/coinbase-samples/ib-usermgr-go/pkg/pbs/profile/v1"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func gRPCListen(app config.AppConfig, aw auth.Middleware) {

	// if local expose both grpc and http endpoints
	activePort := app.Port
	if app.IsLocalEnv() {
		activePort = app.GrpcPort
	}

	//setup conn
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", activePort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	grpcOptions := setupGrpcOptions(app, aw)
	s := grpc.NewServer(grpcOptions...)

	//register grpc handlers
	v1.RegisterProfileServiceServer(s, &handlers.ProfileServer{})
	registerHealth(s)
	reflection.Register(s)

	log.Debugf("gRPC Server starting on port %s\n", activePort)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to listen for gRPC: %v", err)
		}
	}()

	//if local, start http as an interface
	var gwServer *http.Server
	if app.IsLocalEnv() {
		gwServer, err = setupHttp(app, s)
		if err != nil {
			log.Errorf("issues setting up http server: %v", err)
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	if gwServer != nil {
		gwServer.Shutdown(ctx)
	}

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Debug("stopping")
	os.Exit(0)
}

func registerHealth(s *grpc.Server) {
	healthServer := health.NewServer()
	healthServer.SetServingStatus("grpc.health.v1.Health", grpc_health_v1.HealthCheckResponse_SERVING)
	grpc_health_v1.RegisterHealthServer(s, healthServer)
}

func setupGrpcOptions(app config.AppConfig, aw auth.Middleware) []grpc.ServerOption {
	// Logrus entry is used, allowing pre-definition of certain fields by the user.
	// See example setup here https://github.com/grpc-ecosystem/go-grpc-middleware/blob/master/logging/logrus/examples_test.go
	opts := []grpc_logrus.Option{
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
		grpc_logrus.WithDecider(func(fullMethodName string, err error) bool {
			// will not log gRPC calls if it was a call to healthcheck and no error was raised
			if err == nil && fullMethodName == "/grpc.health.v1.Health/Check" {
				return false
			}

			// by default everything will be logged
			return true
		}),
	}

	entry := log.NewEntry()

	grpcOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(entry.GetUnderneath(), opts...),
			aw.InterceptorNew(),
			grpc_validator.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	}

	if !app.IsLocalEnv() {
		// load tls for grpc
		tlsCredentials, err := loadCredentials()
		if err != nil {
			log.Fatalf("Cannot load TLS credentials: %v", err)
		}

		grpcOptions = append(grpcOptions, grpc.Creds(tlsCredentials))
	}

	return grpcOptions
}

func loadCredentials() (credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		return nil, err
	}

	return credentials.NewTLS(
		&tls.Config{
			Certificates: []tls.Certificate{cert},
			ClientAuth:   tls.NoClientCert,
		},
	), nil
}
