//go:build integration

package author

import (
	"context"
	"net"

	"log"
	"testing"
	"time"

	"github.com/Chameleon-m/kvd_grpc/internal/app/library/repository"
	"github.com/Chameleon-m/kvd_grpc/internal/app/library/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/test/bufconn"
)

const (
	address = "localhost:50051"
	bufSize = 1024 * 1024
)

var listenerBuf *bufconn.Listener

// Интеграционный тест, который запускает сервер gRPC, а клиент проверяет службу с помощью RPC.
func TestIntegrationServer_GetByBook(t *testing.T) {
	// Запуск обычного сервера gRPC работает на HTTP2
	srv := initGRPCServerHTTP2()
	defer srv.Stop()
	// Настраиваем соединение с сервером.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// Закрываем соединение
	defer conn.Close()

	// Создаём запрос и его параметры
	book := BookRequest{Id: 1}
	// Создаём клиент
	client := NewAuthorClient(conn)

	// Создаём контекст с таймаутом в секунду
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// Отправляем запрос
	respose, err := client.GetByBook(ctx, &book)
	if err != nil {
		log.Fatalf("could not GetByBook: %v", err)
	}
	log.Printf("Res %s", respose.Items)
}

// Интеграционный тест написан с использованием Buffconn, т.е без заупска сервера
func TestIntegrationServer_GetByBookBufConn(t *testing.T) {
	ctx := context.Background()
	initGRPCServerBuffConn()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(getBufDialer(listenerBuf)), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Создаём запрос и его параметры
	book := BookRequest{Id: 1}
	// Создаём клиент
	client := NewAuthorClient(conn)

	// Создаём контекст с таймаутом в секунду
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// Отправляем запрос
	respose, err := client.GetByBook(ctx, &book)
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Res %s", respose.Items)
}

// Инициализация сервера для тестов
func initGRPCServerHTTP2() *grpc.Server {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return initServer(listener)
}

func getBufDialer(listener *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, url string) (net.Conn, error) {
		return listener.Dial()
	}
}

// Инициализация BufConn
// Пакет bufconn предоставляет реализацию буфера для net.Conn, и связанные с ним функции прослушивания.
func initGRPCServerBuffConn() {
	listenerBuf = bufconn.Listen(bufSize)
	initServer(listenerBuf)
}

// TODO инициализация бд, тестовые данные, общие способы инициализации
func initServer(listener net.Listener) *grpc.Server {
	// Cоздаём сервер gRPC
	srv := grpc.NewServer()
	// Используем репозиторий заглушку
	nilAuthorRepository := repository.NewAuthorNilRepository()
	authorService := service.NewAuthorService(nilAuthorRepository)
	// Регистрируем обработчики авторов
	RegisterAuthorServer(srv, NewHandler(context.Background(), authorService))
	// Регистрируем службу отражения на сервере gRPC.
	reflection.Register(srv)
	go func() {
		if err := srv.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return srv
}
