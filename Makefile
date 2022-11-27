MAKEPATH := $(abspath $(lastword $(MAKEFILE_LIST)))
PWDSLASH := $(dir $(MAKEPATH))
PWD := $(realpath -s $(PWDSLASH))
PROJECT_ID=test
APP=go-kit
REGISTRY=europe-west4-docker.pkg.dev/$(PROJECT_ID)/$(PROJECT_ID)
ENV=dev

# -------------------------------------------
#
# Git info
#
# -------------------------------------------
ifndef ${TAG}
	TAG := $(shell git --git-dir=$(PWD)/../../.git rev-parse HEAD)
endif
ifndef ${SHORT_TAG}
	SHORT_TAG := $(shell git --git-dir=$(PWD)/../../.git rev-parse --short HEAD)
endif
ifndef ${BRANCH}
	BRANCH := $(shell git --git-dir=$(PWD)/../../.git rev-parse --abbrev-ref HEAD)
endif

# -------------------------------------------
#
# Help
#
# -------------------------------------------
.DEFAULT_GOAL := help
.PHONY: help
help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

# -------------------------------------------
#
# Project info run: make info
#
# -------------------------------------------
.PHONY: info
info:			## Useful information
	@echo "--------------------------------------------------------------"
	@echo "Commands:"
	@echo " - make docker-build\t Docker build all containers for production"
	@echo " - make deploy     \t Compile source, lint, test, docker push and terraform deployment"
	@echo "--------------------------------------------------------------"
	@echo "Workdir (pwd): \t\t${PWD}"
	@echo "Git branch: \t\t${BRANCH}"
	@echo "Git tag: \t\t${TAG}"
	@echo "Git short tag: \t\t${SHORT_TAG}"
	@echo "Docker repo: \t\t${REGISTRY}"
	@echo "\n"

# -------------------------------------------
#
# Docker
#
# -------------------------------------------
docker-build:			## Build production images
	docker build --target prod \
		-t $(REGISTRY)/$(APP):$(TAG) \
		-t $(REGISTRY)/$(APP):latest --build-arg BUILDKIT_INLINE_CACHE=1 --cache-from $(REGISTRY)/$(APP):latest $(PWD)

.PHONY: docker-push
docker-push: docker-build
docker-push:			## pushing docker images to registry
	echo "TODO"
#	docker push $(REGISTRY)/$(APP):$(TAG)
#	docker push $(REGISTRY)/$(APP):latest

# -------------------------------------------
#
# TypeOrm Develop
#
# -------------------------------------------
.PHONY: fixtures
fixtures:
	echo "TODO"

# -------------------------------------------
#
# Docker Develop
#
# -------------------------------------------
start:
	docker-compose up

stop:
	docker-compose down

# -------------------------------------------
#
# Compile source code
#
# -------------------------------------------
.PHONY: build
build:  ## compilint typescript
	go build -v -o /usr/local/bin/app ./...

# -------------------------------------------
#
# Format
#
# -------------------------------------------
.PHONY: format
lint:			## Linting app & resources
	gofmt -l -s -w .

# -------------------------------------------
#
# test
#
# -------------------------------------------
.PHONY: test
test:						## Test app & resources
	go test ./...

.PHONY: test-coverage
test-coverage:	## Test coverage app & resources
	go test -coverpkg=./... ./...

.PHONY: test-e2e
test-e2e:				## Test coverage app & resources
	echo "TODO"

# -------------------------------------------
#
# Terraform
#
# -------------------------------------------
.PHONY: terraform-format
terraform-format:				## Terraform format
	echo "TODO"
#terraform -chdir=terraform  fmt --recursive


.PHONY: terraform-destroy
terraform-destroy:			## Terraform destroy
	echo "TODO"
	#terraform -chdir=terraform destroy -var-file="env.$(ENV).tfvars" -var="git_commit_hash=$(TAG)"

.PHONY: terraform-plan
terraform-plan:					## Terraform plan
	echo "TODO"
#	terraform -chdir=terraform init -backend-config="bucket=$(PROJECT_ID)" -backend-config="prefix=$(APP)"
#	terraform -chdir=terraform plan -var-file="env.$(ENV).tfvars" -var="git_commit_hash=$(TAG)"

.PHONY: terraform-deploy-trace
terraform-deploy-trace:	## Terraform apply infra with trace (user for debugging)
	echo "TODO"
#	terraform -chdir=terraform init -backend-config="bucket=$(PROJECT_ID)" -backend-config="prefix=$(APP)"
#	TF_LOG=TRACE terraform -chdir=terraform plan -var-file="env.$(ENV).tfvars" -var="git_commit_hash=$(TAG)"
#	TF_LOG=TRACE terraform -chdir=terraform apply -auto-approve -var-file="env.$(ENV).tfvars" -var="git_commit_hash=$(TAG)"

.PHONY: terraform-apply
terraform-apply:				## Terraform apply infra
	echo "TODO"
#	@make -s terraform-format
#	@make -s terraform-plan
#	terraform -chdir=terraform apply -auto-approve -var-file="env.$(ENV).tfvars" -var="git_commit_hash=$(TAG)"

# -------------------------------------------
#
# Put all the actions together to deploy
#
# -------------------------------------------
.PHONY: pre-deploy
pre-deploy:
	@make -s format
	@make -s build
	@make -s test-coverage

.PHONY: deploy
deploy:
	@make -s pre-deploy
	@make -s docker-push
	@make -s terraform-apply
