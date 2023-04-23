ifneq (test, ${ENV_MODE})
    migrateArgs := -source file://migrations -database "${DB_URI}?x-tls-insecure-skip-verify=false" -verbose # x-statement-timeout
else
    migrateArgs := -source file://migrations -database "${DB_URI_TEST}?x-tls-insecure-skip-verify=false" -verbose
endif

# $(CURDIR) fix old docker version for Windows
dockerMigrate := docker run --name migrate --rm -i --volume="$(CURDIR)/migrations:/migrations" --network netApplication migrate/migrate:v4.15.2

GIT_HASH ?= $(shell git log --format="%h" -n 1)

# TODO ENV
DOCKER_USERNAME ?= korolevd
APPLICATION_SERVER ?= ${DOCKER_USERNAME}/kvado-library-grpc-server
APPLICATION_SERVER_MIGRATION ?= ${DOCKER_USERNAME}/kvado-library-server-migration

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
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -tags "${ENV_MODE} netgo" -installsuffix netgo -o $(BIN_DIR)/linux_amd64/library_grpc_server cmd/library_grpc_server/main.go

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
	go test -tags unit,integration ./...
test-unit:
	go test -tags unit ./...
test-integration:
	go test -tags integration ./...
test-race:
	go test -tags unit,integration -race -vet=off testing ./...
test-coverage:
	go test -tags unit,integration ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

lints:
	golangci-lint run ./...

build-docker: build-docker-server build-docker-migration
build-docker-server: docker-build-push-server docker-build-push-server-latest
build-docker-migration: docker-build-push-migration docker-build-push-migration-latest

docker-build-push-server:
	docker build --tag ${APPLICATION_SERVER}:${GIT_HASH} -f deployments/docker/library_server/Dockerfile .
	docker push ${APPLICATION_SERVER}:${GIT_HASH}

docker-build-push-server-latest:
	docker build --tag ${APPLICATION_SERVER}:latest -f deployments/docker/library_server/Dockerfile .
	docker tag ${APPLICATION_SERVER}:latest ${APPLICATION_SERVER}:latest
	docker push ${APPLICATION_SERVER}:latest

docker-build-push-migration:
	docker build --tag ${APPLICATION_SERVER_MIGRATION}:${GIT_HASH} -f deployments/docker/library_server_migration/Dockerfile .
	docker push ${APPLICATION_SERVER_MIGRATION}:${GIT_HASH}
	

docker-build-push-migration-latest:
	docker build --tag ${APPLICATION_SERVER_MIGRATION}:latest -f deployments/docker/library_server_migration/Dockerfile .
	docker tag  ${APPLICATION_SERVER_MIGRATION}:latest ${APPLICATION_SERVER_MIGRATION}:latest
	docker push ${APPLICATION_SERVER_MIGRATION}:latest

dev-kube: dev-kube-start dev-kube-apply

dev-kube-start:
	minikube start
	minikube tunnel

dev-kube-apply:
	kubectl apply -f deployments/kubernetes/library_server/dev/namespace.yaml
	kubectl apply -f deployments/kubernetes/library_server/dev/config.yaml
	kubectl apply -f deployments/kubernetes/library_server/dev/secret.yaml
	kubectl apply -f deployments/kubernetes/library_server/dev/secret_tls.yaml
	kubectl apply -f deployments/kubernetes/library_server/dev/db.yaml
	kubectl apply -f deployments/kubernetes/library_server/dev/migration_up_job.yaml
	kubectl apply -f deployments/kubernetes/library_server/dev/server_grpc.yaml
	kubectl apply -f deployments/kubernetes/library_server/dev/ingress.yaml