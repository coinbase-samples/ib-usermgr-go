REGION ?= us-east-1
PROFILE ?= cb-stp-sa-dev
ENV_NAME ?= dev

ACCOUNT_ID := $(shell aws sts get-caller-identity --profile $(PROFILE) --query 'Account' --output text)

docker-build:
	@docker build --platform linux/amd64 --build-arg REGION=$(REGION) --build-arg ENV_NAME=$(ENV_NAME) --build-arg ACCOUNT_ID=$(ACCOUNT_ID) .

docker-local-build:
	@docker build --tag ib-usermgr-go:local --build-arg REGION=$(REGION) --build-arg ENV_NAME=local --build-arg ACCOUNT_ID=$(ACCOUNT_ID) .

docker-local-start:
	@docker run --net ib-system_default -p 8450:8450 -p 8451:8451 --env-file .env -d ib-usermgr-go:local

docker-setup:
	@docker compose up -d \
	&& @sleep 20 \
	&& ./setupDynamo.sh 

seed-db:
	./setupDynamo.sh

docker-setup-down:
	@docker compose down

start-full: docker-setup start

start:
	@go run cmd/server/*.go

test.unit:
	@go test -race -cover ./...