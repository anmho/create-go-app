all: cli

.PHONY: cli
cli:
	@go build -o ./bin/cli .

release:
	@goreleaser release --clean


# example tag
# tag:
#     @git tag -a v1.0.0 -m "Release version 1.0.0"