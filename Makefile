REGION ?= us-east-1
PROFILE ?= sa-infra
ENV_NAME ?= dev

ACCOUNT_ID := $(shell aws sts get-caller-identity --profile $(PROFILE) --query 'Account' --output text)

.PHONY: docker-build
docker-build:
	@docker build --platform linux/amd64 --build-arg REGION=$(REGION) --build-arg ENV_NAME=$(ENV_NAME) --build-arg ACCOUNT_ID=$(ACCOUNT_ID) .

.PHONY: docker-build-local
docker-build-local:
	@docker build --tag ib-usermgr-go:local --build-arg REGION=$(REGION) --build-arg ENV_NAME=local --build-arg ACCOUNT_ID=$(ACCOUNT_ID) .

.PHONY: docker-start-local
docker-start-local:
	@docker run --net ib-system_default -p 8450:8450 -p 8451:8451 --env-file .env -d ib-usermgr-go:local


.PHONY: start
start:
	go run cmd/server/*.go
