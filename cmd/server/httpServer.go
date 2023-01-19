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

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/coinbase-samples/ib-usermgr-go/config"
	"github.com/coinbase-samples/ib-usermgr-go/log"
	v1 "github.com/coinbase-samples/ib-usermgr-go/pkg/pbs/profile/v1"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func getProfileConnAddress(app config.AppConfig) string {
	if app.IsLocalEnv() {
		return fmt.Sprintf("%s:%s", "0.0.0.0", app.GrpcPort)
	}
	return fmt.Sprintf("%s:%s", app.InternalApiHostname, app.GrpcPort)
}

func profileConn(app config.AppConfig) (*grpc.ClientConn, error) {
	dialProfileConn := getProfileConnAddress(app)
	log.Debugf("connecting to profile localhost grpc: %s", dialProfileConn)
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		dialProfileConn,
		//grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	return conn, err
}

func setupHttp(app config.AppConfig, grpcServer *grpc.Server) (*http.Server, error) {
	log.Debug("dialing profile")
	pConn, err := profileConn(app)
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	log.Debug("Connected to profile")

	gwmux := runtime.NewServeMux(runtime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
		md := make(map[string]string)
		if method, ok := runtime.RPCMethod(ctx); ok {
			md["method"] = method // /grpc.gateway.examples.internal.proto.examplepb.LoginService/Login
		}
		if pattern, ok := runtime.HTTPPathPattern(ctx); ok {
			md["pattern"] = pattern // /v1/example/login
		}
		return metadata.New(md)
	}))

	gwmux.HandlePath("GET", "/health", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		log.Debug("responding to health check")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "ok\n")
	})

	// Register Service Handlers
	err = v1.RegisterProfileServiceHandler(context.Background(), gwmux, pConn)

	if err != nil {
		log.Fatalf("Failed to register profile: %v", err)
	}

	gwServer := &http.Server{
		Handler:      makeHttpHandler(gwmux, app),
		Addr:         fmt.Sprintf(":%s", app.Port),
		WriteTimeout: 40 * time.Second,
		ReadTimeout:  40 * time.Second,
	}

	log.Debugf("started http gRPC-Gateway on - %v", app.Port)

	go func() {
		if app.Env == "local" {
			if err := gwServer.ListenAndServe(); err != nil {
				log.Fatalf("ListenAndServe: %v", err)
			}
			log.Debugf("started http")
		} else {
			if err := gwServer.ListenAndServeTLS("server.crt", "server.key"); err != nil {
				log.Fatalf("ListenAndServeTLS: %v", err)
			}
			log.Debugf("started https")
		}
	}()

	return gwServer, nil
}

func makeHttpHandler(gwmux http.Handler, app config.AppConfig) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	origins := []string{
		fmt.Sprintf("https://localhost:%s", app.Port),
		fmt.Sprintf("http://localhost:%s", app.Port),
		"http://localhost:4200",
		"https://localhost:4200",
	}

	originsOk := handlers.AllowedOrigins(origins)

	log.Debugf("starting http - %v - %v - %v", originsOk, headersOk, methodsOk)
	return handlers.CORS(originsOk, headersOk, methodsOk)(gwmux)
}
