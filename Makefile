ifneq (test, ${ENV_MODE})
    migrateArgs := -source file://migrations -database "${DB_URI}?x-tls-insecure-skip-verify=false" -verbose # x-statement-timeout
else
    migrateArgs := -source file://migrations -database "${DB_URI_TEST}?x-tls-insecure-skip-verify=false" -verbose
endif

# $(CURDIR) fix old docker version for Windows
dockerMigrate := docker run --name migrate --rm -i --volume="$(CURDIR)/migrations:/migrations" --network netApplication migrate/migrate:v4.15.2

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

build-mocks:
	@mockgen -source internal/app/library/transport/grpc/handlers/author/author.go -destination internal/app/library/transport/grpc/handlers/author/author_mock.go -package author
	@mockgen -source internal/app/library/transport/grpc/handlers/book/book.go -destination internal/app/library/transport/grpc/handlers/book/book_mock.go -package book
	@mockgen -source internal/app/library/service/author.go -destination internal/app/library/service/author_mock.go -package service
	@mockgen -source internal/app/library/service/book.go -destination internal/app/library/service/book_mock.go -package service

migrate-up:
	migrate $(migrateArgs) up $(if $n,$n,)
migrate-down:
	migrate $(migrateArgs) down $(if $n,$n,)
migrate-goto:
	migrate $(migrateArgs) goto $(v)
migrate-force:
	migrate $(migrateArgs) force $(v)
migrate-drop:
	migrate $(migrateArgs) drop
migrate-version:
	migrate $(migrateArgs) version
migrate-create-sql:
	migrate $(migrateArgs) create -ext sql -dir migrations $(name)

migrate-up-docker: 
	$(dockerMigrate) $(migrateArgs) up $(if $n,$n,)
migrate-down-docker:
	$(dockerMigrate) $(migrateArgs) down $(if $n,$n,)
migrate-goto-docker:
	$(dockerMigrate) $(migrateArgs) goto $(v)
migrate-force-docker:
	$(dockerMigrate) $(migrateArgs) force $(v)
migrate-drop-docker:
	$(dockerMigrate) $(migrateArgs) drop
migrate-version-docker:
	$(dockerMigrate) $(migrateArgs) version
migrate-create-sql-docker:
	$(dockerMigrate) $(migrateArgs) create -ext sql -dir migrations $(name)

test:
	go test -tags testing ./...
test-race:
	go test -tags -race -vet=off testing ./...
test-coverage:
	go test -tags testing ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

lints:
	golangci-lint run ./...