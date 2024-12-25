all: cli

.PHONY: cli
cli:
	@go build -o ./bin/cli .
