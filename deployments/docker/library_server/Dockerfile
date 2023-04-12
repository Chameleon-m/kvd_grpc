# Многоэтапная сборка

# Этап 1
FROM golang:1.20.3

ENV location=/go/src/github.com/Chameleon-m/kvd_grpc
WORKDIR ${location}

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd/library_grpc_server/main.go ./cmd/library_grpc_server/main.go
COPY ./internal ./internal

WORKDIR ${location}/cmd/library_grpc_server/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o library_grpc_server .

# Этап 2
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/Chameleon-m/kvd_grpc/cmd/library_grpc_server/library_grpc_server .

RUN GRPC_HEALTH_PROBE_VERSION=v0.4.17 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

ENTRYPOINT [ "/root/library_grpc_server" ]
EXPOSE 50051