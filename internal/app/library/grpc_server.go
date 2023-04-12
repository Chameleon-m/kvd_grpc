package library

import (
	"context"
	"database/sql"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Chameleon-m/kvd_grpc/internal/app/library/repository"
	"github.com/Chameleon-m/kvd_grpc/internal/app/library/service"
	"github.com/Chameleon-m/kvd_grpc/internal/app/library/transport/grpc/handlers/author"
	"github.com/Chameleon-m/kvd_grpc/internal/app/library/transport/grpc/handlers/book"
	healthprobe "github.com/Chameleon-m/kvd_grpc/internal/app/library/transport/grpc/handlers/health_probe"

	mysql "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
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

	// Настраиваем соединение с DB
	dbConfig := mysql.NewConfig()
	dbConfig.Net = "tcp"
	dbConfig.Addr = os.Getenv("DB_ADDR")
	dbConfig.User = os.Getenv("DB_USER")
	dbConfig.Passwd = os.Getenv("DB_PASSWORD")
	dbConfig.DBName = os.Getenv("DB_NAME")
	// dbConfig.Timeout = time.Second*1
	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		log.Println(err)
		return
	}

	// Время жизни соединения
	db.SetConnMaxLifetime(time.Minute * 5)
	// Время ожидания в пуле
	db.SetConnMaxIdleTime(time.Second * 5)
	// Максимальное количество соединений
	db.SetMaxOpenConns(10)
	// Ограничение размера пула
	db.SetMaxIdleConns(10)
	// Закрываем соединение
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("close error:", err)
		}
	}()
	// Проверяем соединение с базой данных
	if err := db.PingContext(ctx); err != nil {
		log.Println(err)
		return
	}
	// Передаём логер драйверу mysql
	if err := mysql.SetLogger(log.Default()); err != nil {
		log.Println(err)
		return
	}

	// Слушаем tcp порт
	host := os.Getenv("LIBRARY_SERVER_HOST")
	port := os.Getenv("LIBRARY_SERVER_PORT")
	lsn, err := net.Listen("tcp", net.JoinHostPort(host, port))
	if err != nil {
		log.Printf("failde to listen: %v", err)
		return
	}

	// Cоздаём сервер gRPC
	srv := grpc.NewServer()
	// Регистрируем обработчики книг
	mysqlBookRepository := repository.NewBookMysqlRepository(db)
	bookService := service.NewBookService(mysqlBookRepository)
	book.RegisterBookServer(srv, book.NewHandler(ctx, bookService))
	// Регистрируем обработчики авторов
	mysqlAuthorRepository := repository.NewAuthorMysqlRepository(db)
	authorService := service.NewAuthorService(mysqlAuthorRepository)
	author.RegisterAuthorServer(srv, author.NewHandler(ctx, authorService))
	// Регистрируем обработчик health probe
	grpc_health_v1.RegisterHealthServer(srv, healthprobe.NewChecker())
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
			log.Printf("failed to serve: %v", err)
			// Завершаем контекст
			ctxCancel()
		}
	}()

	// Блокируемся и ждём сигнала в канале завершения
	<-done

	// Завершаем уже запущенные процессы
	log.Print("Shutting down gracefully, press Ctrl+C again to force")
	srv.GracefulStop()
}
