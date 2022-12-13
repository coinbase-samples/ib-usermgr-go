### User Manager

This is the application code for the User Manager implementation.

### Local Environment Setup

- [Docker](https://docs.docker.com/get-docker/) - Containers are used to run the Localstack for DynamoDb Database locally.
  Installation instructions unnecessary as you probably have this already installed. If not, follow the link and download Docker.

To spin up the container to run dynamodb:
`docker-compose up -d`

This will start up the database running at localhost:4566

To insert the database configuration and start the application layer run:

```
make start-full
```
