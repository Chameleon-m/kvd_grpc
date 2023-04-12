
BIN_DIR = $(PWD)/bin

.PHONY: build

$(VERBOSE).SILENT:

clean:
	rm -rf bin/*

dependencies:
	go mod download

build: dependencies build-server

build-server: 
	go build -tags ${ENV_MODE} -o $(BIN_DIR)/library_grpc_server cmd/library_grpc_server/main.go

linux-binaries:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "${ENV_MODE} netgo" -installsuffix netgo -o $(BIN_DIR)/library_grpc_server cmd/library_grpc_server/main.go

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

test:
	go test -tags testing ./...
test-race:
	go test -tags -race -vet=off testing ./...
test-coverage:
	go test -tags testing ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

lints:
	golangci-lint run ./...