package author

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
	GetListByBook(ctx context.Context, id uint64) (model.AuthorList, error)
}

type Author struct {
	UnimplementedAuthorServer

	ctx     context.Context
	service ServiceInterface
}

func NewHandler(
	ctx context.Context,
	service ServiceInterface,
) *Author {
	return &Author{
		ctx:     ctx,
		service: service,
	}
}

// GetByAuthor реализует kvado.Auhor.getByBook
func (s *Author) GetByBook(ctx context.Context, in *BookRequest) (*AuthorListResponse, error) {

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

	// Список авторов по книге
	authorList, err := s.service.GetListByBook(ctx, in.GetId())
	if err != nil {
		log.Printf("service.GetListByBook error: %s", err.Error())
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	// Собираем ответ
	return &AuthorListResponse{Items: makeBookList(model.BookList(authorList))}, status.New(codes.OK, "").Err()
}

// Преобразуем с моделей в grpc структуры
func makeBookList(list model.BookList) []*AuthorItem {
	items := make([]*AuthorItem, 0, len(list))
	for _, v := range list {
		items = append(items, &AuthorItem{Id: v.ID, Name: v.Name})
	}
	return items
}
