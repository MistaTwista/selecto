.PHONY: install-linter
install-linter:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: lint
lint: install-linter
	@golangci-lint run

.PHONY: build
build: build-gena build-reedo ## Build binaries
	go build -o bin/selecto main.go

.PHONY: install
install: build ## Build and install to user bin
	cp bin/selecto ~/.bin/selecto

.PHONY: build-gena
build-gena:
	${MAKE} -C gena build

.PHONY: build-reedo
build-reedo:
	${MAKE} -C reedo build

exec: build ## Build and exec binary
	./bin/selecto

treedo: build ## Build and run test with reedo
	./bin/gena -d 4s -e 100ms | bin/reedo

test: build ## Build and run test with selecto
	./bin/gena -d 1s -e 10ms | bin/selecto --stdin | cat

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
