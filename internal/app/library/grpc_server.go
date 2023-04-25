package library

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
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

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred: ", err)
			os.Exit(1)
		}
	}()

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
	dbConfig.Addr = getEnvReq("DB_ADDR")
	dbConfig.User = getEnvReq("DB_USER")
	dbConfig.Passwd = getEnvReq("DB_PASSWORD")
	dbConfig.DBName = getEnvReq("DB_NAME")
	// dbConfig.Timeout = time.Second*1
	// dbConfig.InterpolateParams = true
	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		log.Panic(err)
	}

	dbConnMaxLifeTime := getEnvReqInt("DB_CONN_MAX_LIFE_TIME_SEC")
	dbConnMaxIdleTime := getEnvReqInt("DB_CONN_MAX_IDLE_TIME_SEC")
	dbMaxOpenConss := getEnvReqInt("DB_MAX_OPEN_CONNNS")
	dbMaxIdleConns := getEnvReqInt("DB_MAX_IDLE_CONNS")

	// Время жизни соединения
	db.SetConnMaxLifetime(time.Second * time.Duration(dbConnMaxLifeTime))
	// Время ожидания в пуле
	db.SetConnMaxIdleTime(time.Second * time.Duration(dbConnMaxIdleTime))
	// Максимальное количество соединений
	db.SetMaxOpenConns(dbMaxOpenConss)
	// Ограничение размера пула
	db.SetMaxIdleConns(dbMaxIdleConns)
	// Закрываем соединение
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("close error:", err)
		}
	}()
	// Проверяем соединение с базой данных
	if err := db.PingContext(ctx); err != nil {
		log.Panic(err)
	}
	// Передаём логер драйверу mysql
	if err := mysql.SetLogger(log.Default()); err != nil {
		log.Panic(err)
	}

	// Слушаем tcp порт
	addr := getEnvReq("LIBRARY_SERVER_ADDR")
	lsn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Panicf("failde to listen: %v", err)
	}

	// Cоздаём сервер gRPC
	srv := grpc.NewServer()
	defer func() {
		// Завершаем уже запущенные процессы
		log.Print("Shutting down gracefully gRPC server")
		srv.GracefulStop()
		srv.Stop()
	}()
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

	log.Printf("Starting gRPC listener on address %s", addr)
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
}

func getEnvReq(env string) string {
	e := os.Getenv(env)
	if e == "" {
		panic(fmt.Errorf("Env %s not set", env))
	}
	return e
}

func getEnvReqInt(env string) int {
	e := getEnvReq(env)
	i, err := strconv.Atoi(e)
	if err != nil {
		panic(fmt.Errorf("Env %s error %v", env, err))
	}
	return i
}
