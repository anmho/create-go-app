

default: api

.PHONY: watch migrate
.PHONY: deps
deps:
	@brew tap golangci/tap
	@brew install golangci/tap/golangci-lint
	@go install github.com/bufbuild/buf/cmd/buf@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
	@go install "github.com/vektra/mockery/v2@v2.49.1"

.PHONY: check
check:
	@go vet ./...
	@golangci-lint run ./...

.PHONY: mocks
mocks: # Generate mocks
	@go generate ./...

.PHONY: clean
clean:
	@rm -rf ./bin
	
.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: test
test:
	@go test -cover ./...

.PHONY: api
api:
	@go build -o ./bin/api ./cmd/api


IMAGE_NAME := anmho/happened
GIT_COMMIT_HASH = $(shell git rev-parse --short HEAD)
TAG ?= $(GIT_COMMIT_HASH)
image:
	@docker build -t $(IMAGE_NAME):$(TAG) .
	@echo docker.io/$(IMAGE_NAME):$(TAG) | tee image.txt

.PHONY: push
push:
	@docker push $(IMAGE_NAME):$(TAG)
	@echo docker.io/$(IMAGE_NAME):$(TAG) | tee image.txt

.PHONY: watch
watch:
	@wgo run ./cmd/api

.PHONY: db
# Database commands
# Startup the development database
db:
	@docker compose up -d

# Applies the specified schema file to the database, similar to Terraform.
.PHONY: migrate
migrate:
	@atlas schema apply --url "postgres://admin:admin@localhost:5433/happened_db?sslmode=disable" --to "file://sql/schema.sql" --dev-url "docker://postgres/15"

# Connects to the local Postgres database
.PHONY: connect
connect:
	@psql "postgres://admin:admin@localhost:5433/happened_db"


