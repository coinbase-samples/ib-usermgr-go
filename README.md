### User Manager

This is the application code for the User Manager implementation.

This application is the simplest example of having a grpc-gateway service that can read/write to dynamodb.

The service is authenticated by AWS Cognito

### Local Environment Setup

- [Docker](https://docs.docker.com/get-docker/) - Containers are used to run the Localstack for DynamoDb Database locally.
  Installation instructions unnecessary as you probably have this already installed. If not, follow the link and download Docker.

Copy the sample.env to .env and fill appropriately

To insert the database configuration and start the application layer run:

```
make start-full
```

This does the following:

- Runs docker-compose up - default localstack is localhost:4566
- Runs setupDynamo to give default profiles in the database
- Starts the application server - default http port is 8450 and grpc port is 8451

## License

This library is licensed under the Apache 2.0 License. See the [LICENSE](LICENSE) file.
