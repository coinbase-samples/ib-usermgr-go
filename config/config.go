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

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// TODO move to util pkg
type BaseConfig struct {
	Env       string `mapstructure:"ENV_NAME"`
	LogLevel  string `mapstructure:"LOG_LEVEL"`
	LogToFile string `mapstructure:"LOG_TO_FILE"`
	Region    string `mapstructure:"REGION"`
	Port      string `mapstructure:"PORT"`
	GrpcPort  string `mapstructure:"GRPC_PORT"`
}

func UnmarshalBase(b *BaseConfig) {
	err := viper.Unmarshal(&b)
	if err != nil {
		fmt.Printf("Cannot parse env file %v\n", err)
	}
}

type AppConfig struct {
	BaseConfig
	ClientId            string `mapstructure:"COGNITO_APP_CLIENT_ID"`
	UserPoolId          string `mapstructure:"COGNITO_USER_POOL_ID"`
	DatabaseEndpoint    string `mapstructure:"DB_ENDPOINT"`
	ProfileTableName    string `mapstructure:"PROFILE_TABLE"`
	InternalApiHostname string `mapstructure:"INTERNAL_API_HOSTNAME"`
}

func (a AppConfig) IsLocalEnv() bool {
	return a.Env == "local"
}

func Setup(app *AppConfig) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.AllowEmptyEnv(true)
	// set defaults
	viper.SetDefault("ENV_NAME", "local")
	viper.SetDefault("LOG_LEVEL", "warning")
	viper.SetDefault("AWS_REGION", "us-east-1")

	viper.SetDefault("PORT", "8451")
	viper.SetDefault("GRPC_PORT", "8452")

	viper.SetDefault("COGNITO_APP_CLIENT_ID", "local")
	viper.SetDefault("COGNITO_USER_POOL_ID", "local")
	viper.SetDefault("DB_ENDPOINT", "http://localhost:4566")
	viper.SetDefault("PROFILE_TABLE", "Profile")
	viper.SetDefault("INTERNAL_API_HOSTNAME", "NOT_SET")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Missing env file %v\n", err)
	}

	if err = viper.Unmarshal(&app); err != nil {
		fmt.Printf("Cannot parse env file %v\n", err)
	}
	UnmarshalBase(&app.BaseConfig)
}
