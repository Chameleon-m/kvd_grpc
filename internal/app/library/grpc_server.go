package library

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run library grpc server.
func RunGRPCServer() {
	log.Print("Server grpc started")
	defer log.Print("Server grpc exiting")

	// Создаём контекст, который прослушивает сигнал прерывания от ОС.
	ctxSignal, ctxSignalStop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer ctxSignalStop()

	// Создаём базовый контекст
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	// Слушаем tcp порт
	host := os.Getenv("LIBRARY_SERVER_HOST")
	port := os.Getenv("LIBRARY_SERVER_PORT")
	lsn, err := net.Listen("tcp", net.JoinHostPort(host, port))
	if err != nil {
		log.Fatalf("failde to listen: %v", err)
	}

	// Cоздаём сервер gRPC
	srv := grpc.NewServer()

	// Зарегистрируйте службу отражения на сервере gRPC.
	reflection.Register(srv)

	// Создаём канал для ожидания сигналов завершения
	done := make(chan struct{}, 1)
	defer close(done)
	// Слушаем сигналы
	go func() {
		select {
		// сигнал прерывания от ОС
		case <-ctxSignal.Done():
			log.Print("os signal done")
		// слушаем базовый контекст
		case <-ctx.Done():
			log.Print("app ctx done")
		}
		// В случаем прихода любого из сигналов, пишем в канал
		done <- struct{}{}
	}()

	log.Printf("Starting gRPC listener on host:port %s:%s", host, port)
	go func() {
		// Запускаем сервер
		if err := srv.Serve(lsn); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Блокируемся и ждём сигнала в канале завершения
	<-done

	// Завершаем уже запущенные процессы
	log.Print("Shutting down gracefully, press Ctrl+C again to force")
	srv.GracefulStop()
}
