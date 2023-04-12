
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

build-protoc:
	@protoc --proto_path=internal/app/library/proto --go_out=internal/app/library/transport/grpc/handlers/author --go_opt=paths=source_relative --go-grpc_out=internal/app/library/transport/grpc/handlers/author --go-grpc_opt=paths=source_relative internal/app/library/proto/author.proto
	@protoc --proto_path=internal/app/library/proto --go_out=internal/app/library/transport/grpc/handlers/book --go_opt=paths=source_relative --go-grpc_out=internal/app/library/transport/grpc/handlers/book --go-grpc_opt=paths=source_relative internal/app/library/proto/book.proto

test:
	go test -tags testing ./...
test-race:
	go test -tags -race -vet=off testing ./...
test-coverage:
	go test -tags testing ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

lints:
	golangci-lint run ./...