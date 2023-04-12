# kvd_grpc

# MAIN COMMANDS
```
make build
make test
make lints
make build-protoc
# ... смотри Makefile
```

# ENV
Смотри .env.example

Linux, добавить в ~/.profile
```
ENV_MODE=dev
# ... from .env.example
```

Windows, через GUI, Система->Дополнительные параметрысистемы->Переменные среды

# DEV 

## Устанавливаем protoc, protoc-gen-go, grpc, mockgen

```
sudo apt install protobuf-compiler
```
Для win
1) Скачать https://github.com/protocolbuffers/protobuf/releases 
2) Добавить в enc PATH


```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/golang/mock/mockgen@latest
```

```
make build-protoc
```