ENV=local
SERVICE_NAME=go-kit
POSTGRES_HOST=postgres
POSTGRES_DB=test
POSTGRES_PORT=5432
POSTGRES_USER=user
POSTGRES_PASSWORD=pass
POSTGRES_DRIVER=postgres

ifndef ${TAG}
	TAG := $(shell git --git-dir=$(PWD)/.git rev-parse HEAD)
endif
ifndef ${SHORT_TAG}
	SHORT_TAG := $(shell git --git-dir=$(PWD)/.git rev-parse --short HEAD)
endif
ifndef ${BRANCH}
	BRANCH := $(shell git --git-dir=$(PWD)/.git rev-parse --abbrev-ref HEAD)
endif

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
	@echo "Git app: \t\t${SERVICE_NAME}"
	@echo "Git short tag: \t\t${SHORT_TAG}"
	@echo "\n"


.PHONY: start
start:
	SERVICE_NAME=$(SERVICE_NAME) \
	ENV=$(ENV) \
	POSTGRES_DRIVER=$(POSTGRES_DRIVER) \
	POSTGRES_DB=$(POSTGRES_DB) \
	POSTGRES_HOST=$(POSTGRES_HOST) \
    POSTGRES_PORT=$(POSTGRES_PORT) \
    POSTGRES_USER=$(POSTGRES_USER) \
    POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
	go run src/main.go

.PHONY: format
format:			## format Go code
	echo "  >  Formatting Go code..."
	gofmt -l -s -w .

.PHONY: build
build:  		## build Go code
	echo "  >  building Go code..."
	go build -v ./...


.PHONY: test
test:
	echo "  >  Running vet ..."
	go vet ./...
	echo "  >  Running tests and generating coverage output ..."

	SERVICE_NAME=$(SERVICE_NAME) \
	ENV=$(ENV) \
	POSTGRES_DRIVER=$(POSTGRES_DRIVER) \
	POSTGRES_DB=$(POSTGRES_DB) \
	POSTGRES_HOST=$(POSTGRES_HOST) \
	POSTGRES_PORT=$(POSTGRES_PORT) \
	POSTGRES_USER=$(POSTGRES_USER) \
	POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
	go test ./... -coverprofile coverage.out -covermode count
	sleep 2 # Sleeping to allow for coverage.out file to get generated
	@echo "Current test coverage : $(shell go tool cover -func=coverage.out | grep total | grep -Eo '[0-9]+\.[0-9]+') %"