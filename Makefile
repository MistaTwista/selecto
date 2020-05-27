.PHONY: install-linter
install-linter:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: lint
lint: install-linter
	@golangci-lint run

.PHONY: build
build: ## Build a binary
	go build -o bin/selecto main.go
	go build -o bin/gena gena/gena.go
	go build -o bin/reedo reedo/reedo.go


exec: build ## Build and exec binary
	./bin/selecto

treedo: build ## Build and run test with reedo
	./bin/gena -d 4s -e 100ms | bin/reedo

test: build ## Build and run test with selecto
	./bin/gena -d 1s -e 10ms | bin/selecto --stdin | cat

ttt: ## TTT
	grep -E 'Host [^*]' ~/.ssh/config* | awk 'BEGIN {FS=":Host "}; {printf "%s @ %s\n", $2, $1}' | ./bin/selecto | awk 'BEGIN {FS=" @ "}; {printf "ssh %s -F %s\n", $1, $2}' | echo

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
