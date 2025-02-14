.PHONY: lint
lint:
	go mod tidy && \
	go fmt ./... && \
	golangci-lint run --enable-only gci --fix && \
	gofumpt -l -w . && \
	golangci-lint run --fix
