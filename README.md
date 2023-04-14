# kvd_grpc

# Запуск gRPC сервера и DB

Настроить переменные окружения, смотри раздел ENV
Настроить hosts, смотри в конце

## Вариант 1 docker-compose
```
docker-compose up -d
// Накатываем миграции (подробней смотри раздел "Запускаем миграции")
make migrate-up
// Накатываем данные
docker exec -i db sh -c 'exec mysql -uroot -p"$MYSQL_ROOT_PASSWORD"' < testdata/data.sql
```

## Вариант 2 docker
```
// Создаём сеть
docker network create netApplication
// Заускаем mysql
docker run --name=db --network=netApplication --hostname=db -p 3306:3306 -v ./deployments/docker/mysql/conf.d:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=library -e MYSQL_USER=librarian -e MYSQL_PASSWORD=librarianpassword -d mysql:8.0.32
// Собираем сервер
make build
// Накатываем миграции (подробней смотри раздел "Запускаем миграции")
make migrate-up
// Накатываем данные
docker exec -i db sh -c 'exec mysql -uroot -p"$MYSQL_ROOT_PASSWORD"' < testdata/data.sql
// Заупскаем сервер
./bin/library_grpc_server
```

## Запускаем миграции
```
go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
// -tags 'mysql'
```

Linux
```
DB_URI='mysql://librarian:librarianpassword@tcp(db:3306)/library' make migrate-up
or
DB_URI='mysql://librarian:librarianpassword@tcp(db:3306)/library' make migrate-up-docker
```
Windows
```
$env:DB_URI = 'mysql://librarian:librarianpassword@tcp(db:3306)/library';make migrate-up
or
$env:DB_URI = 'mysql://librarian:librarianpassword@tcp(db:3306)/library';make migrate-up-docker
```
127.0.0.1 db

## Тесты
```
make test
```

Установить 
```
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```
```
grpcurl -plaintext 127.0.0.1:50051 list
grpcurl -plaintext 127.0.0.1:50051 grpc.health.v1.Health.Check
grpcurl -plaintext -d '{\"id\":1}' 127.0.0.1:50051 kvado.Author.getByBook
grpcurl -plaintext -d '{\"id\":1}' 127.0.0.1:50051 kvado.Book.getByAuthor
```

# Нагрузочное тестирование
Инструмент - https://ghz.sh/
https://github.com/bojand/ghz
```
go install github.com/bojand/ghz/cmd/ghz@latest
```
```
ghz --insecure --call kvado.Author.getByBook -d '{\"id\":1}' -n 100000 -c150 0.0.0.0:50051
ghz --insecure --call kvado.Book.getByAuthor -d '{\"id\":1}' -n 100000 -c150 0.0.0.0:50051
or
ghz --insecure --config testdata/ghz/kvado.Author.getByBook.json
ghz --insecure --config testdata/ghz/kvado.Book.getByAuthor.json
```

Summary:
  Count:        100000
  Total:        7.74 s
  Slowest:      106.71 ms
  Fastest:      1.03 ms
  Average:      14.56 ms
  Requests/sec: 12918.19

# MAIN COMMANDS
```
make build
make test
make lints
make migrate-up
make build-protoc
make build-mocks
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
2) Добавить путь в env PATH


```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/golang/mock/mockgen@latest
```

Генерация pb
```
make build-protoc
```
Генерация моков
```
make build-mocks
```
Создание миграции
```
make migrate-create-sql name="init"
```

## Update hostnames
Once the replica set is up, you will need to update hostnames in local /etc/hosts file.
```
127.0.0.1 db library_server
# ...
```
**NOTE**: In windows, the hosts file is located at C:\Windows\System32\drivers\etc\hosts