package book

import (
	"context"
	"log"

	"github.com/Chameleon-m/kvd_grpc/internal/app/library/model"
	. "github.com/Chameleon-m/kvd_grpc/internal/app/library/transport/grpc/handlers"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ServiceInterface ...
type ServiceInterface interface {
	GetListByAuthor(ctx context.Context, id uint64) (model.BookList, error)
}

type Book struct {
	UnimplementedBookServer

	ctx     context.Context
	service ServiceInterface
}

func NewHandler(
	ctx context.Context,
	service ServiceInterface,
) *Book {
	return &Book{
		ctx:     ctx,
		service: service,
	}
}

// GetByAuthor реализует kvado.Book.getByAuthor
func (s *Book) GetByAuthor(ctx context.Context, in *AuthorRequest) (*BookListResponse, error) {

	// Проверяем Deadline запроса
	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("RPC has reached deadline exceeded state : %s", ctx.Err())
		return nil, status.Error(codes.DeadlineExceeded, ctx.Err().Error())
	}

	// Провреяем ID
	// TODO можно валидировать через плагин https://github.com/grpc-ecosystem/go-grpc-middleware/tree/main/validator
	if in.GetId() == 0 {
		return nil, StatusInvalidArgumentDetails(&ErrBadRequeestFieldId)
	}

	// Список книг по автору
	bookList, err := s.service.GetListByAuthor(ctx, in.GetId())
	if err != nil {
		log.Printf("service.GetListByAuthor error: %s", err.Error())
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	// Собираем ответ
	return &BookListResponse{Items: makeBookList(bookList)}, status.New(codes.OK, "").Err()
}

// Преобразуем с моделей в grpc структуры
func makeBookList(list model.BookList) []*BookItem {
	items := make([]*BookItem, 0, len(list))
	for _, v := range list {
		items = append(items, &BookItem{Id: v.ID, Name: v.Name})
	}
	return items
}
