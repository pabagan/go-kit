APP=go-kit

ifndef ${TAG}
	TAG := $(shell git --git-dir=$(PWD)/../../.git rev-parse HEAD)
endif
ifndef ${SHORT_TAG}
	SHORT_TAG := $(shell git --git-dir=$(PWD)/../../.git rev-parse --short HEAD)
endif
ifndef ${BRANCH}
	BRANCH := $(shell git --git-dir=$(PWD)/../../.git rev-parse --abbrev-ref HEAD)
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
	@echo "Git app: \t\t${APP}"
	@echo "Git short tag: \t\t${SHORT_TAG}"
	@echo "Docker repo: \t\t${REGISTRY}"
	@echo "\n"

.PHONY: format
lint:			## format Go code
	echo "  >  Formatting Go code..."
	gofmt -l -s -w .

.PHONY: build
build:  		## build Go code
	echo "  >  building Go code..."
	go build -v -o /usr/local/bin/app ./...


.PHONY: test
test:
	echo "  >  Running vet ..."
	go vet ./...
	echo "  >  Running tests and generating coverage output ..."
	go test ./... -coverprofile coverage.out -covermode count
	sleep 2 # Sleeping to allow for coverage.out file to get generated
	@echo "Current test coverage : $(shell go tool cover -func=coverage.out | grep total | grep -Eo '[0-9]+\.[0-9]+') %"