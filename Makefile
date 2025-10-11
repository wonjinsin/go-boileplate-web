PACKAGE = pikachu
CUSTOM_OS = ${GOOS}
BASE_PATH = $(shell pwd)
BIN = $(BASE_PATH)/bin
BINARY_NAME = bin/$(PACKAGE)
MAIN = $(BASE_PATH)/main.go
GOLINT = $(BIN)/golint
GOBIN = $(shell go env GOPATH)/bin
MOCK = $(GOBIN)/mockgen
PKG_LIST = $(shell cd $(BASE_PATH) && cat pkg.list)


ifneq (, $(CUSTOM_OS))
	OS ?= $(CUSTOM_OS)
else
	OS ?= $(shell uname | awk '{print tolower($0)}')
endif

tool:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/mikefarah/yq/v4@latest
	go install go.uber.org/mock/mockgen@latest

build:
	GOOS=$(OS) go build -o $(BINARY_NAME) $(MAIN)

.PHONY: vet
vet:
	go vet

.PHONY: fmt
fmt:
	go fmt

.PHONY: lint
lint: golangci-lint run

.PHONY: test
test: build-mocks
	go test -v -cover ./...

test-all: test vet fmt lint

build-mocks:
	$(MOCK) -source=service/service.go -destination=mock/mock_service.go -package=mock
	$(MOCK) -source=repository/repository.go -destination=mock/mock_repository.go -package=mock

.PHONY: init
init: 
	go mod init pikachu

.PHONY: tidy
tidy: 
	go mod tidy

.PHONY: vendor
vendor: build-mocks
	go mod vendor

migrate-up:
ifndef ENV
	$(error ENV is not set. Please specify environment e.g., 'make migrate-up ENV=dev')
endif
	@echo "Running migrations for $(ENV) environment using $(ENV_FILE)"
	$(MIGRATE) up

migrate-down:
ifndef ENV
	$(error ENV is not set. Please specify environment e.g., 'make migrate-up ENV=dev')
endif
	@echo "Running migrations for $(ENV) environment using $(ENV_FILE)"
	$(MIGRATE) down 1

start:
	@$(BIN)/$(PACKAGE)

all: tool init tidy vendor build

clean:; $(info cleaning…) @ 
	@rm -rf vendor mock bin
	@rm -rf go.mod go.sum pkg.list
