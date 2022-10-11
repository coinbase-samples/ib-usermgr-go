package dba

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/coinbase-samples/ib-usermgr-go/config"
)

// Repo the repository used by dynamo
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	Svc *dynamodb.Client
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	svc := setupService(a)

	return &Repository{
		App: a,
		Svc: svc,
	}
}

// NewDBA sets the repository for the handlers
func NewDBA(r *Repository) {
	Repo = r
}

func setupService(app *config.AppConfig) *dynamodb.Client {
	var cfg aws.Config
	var err error
	if app.IsLocalEnv() {
		cfg, err = awsConfig.LoadDefaultConfig(context.TODO(),
			awsConfig.WithRegion(app.Region),
			awsConfig.WithEndpointResolver(aws.EndpointResolverFunc(
				func(service, region string) (aws.Endpoint, error) {
					return aws.Endpoint{URL: app.DatabaseEndpoint}, nil
				})),
			awsConfig.WithCredentialsProvider(credentials.StaticCredentialsProvider{
				Value: aws.Credentials{
					AccessKeyID: "dummy", SecretAccessKey: "dummy", SessionToken: "dummy",
					Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
				},
			}),
		)
	} else {
		cfg, err = awsConfig.LoadDefaultConfig(context.TODO())
	}
	if err != nil {
		// TODO: should handle retries and health statuses
		fmt.Println("error creating dynamo config", err)
		return nil
	}
	var svc *dynamodb.Client

	if app.IsLocalEnv() {
		svc = dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
			o.EndpointResolver = dynamodb.EndpointResolverFromURL(app.DatabaseEndpoint)
		})
	} else {
		svc = dynamodb.NewFromConfig(cfg)
	}

	return svc
}
