all: cli

.PHONY: cli
cli:
	@go build -o ./bin/cli .

release:
	@goreleaser release --clean