package service

import (
	"context"

	"github.com/Chameleon-m/kvd_grpc/internal/app/library/model"
)

// BookRepositoryInterface ...
type BookRepositoryInterface interface {
	FindAllByAuthor(ctx context.Context, id uint64) (model.BookList, error)
}

// BookService ...
type BookService struct {
	bookRepo BookRepositoryInterface
}

// NewBookService создаём сервис книг
func NewBookService(bookRepo BookRepositoryInterface) *BookService {
	return &BookService{
		bookRepo: bookRepo,
	}
}

// Список книг по автору
func (s *BookService) GetListByAuthor(ctx context.Context, id uint64) (model.BookList, error) {
	if id == 0 {
		return nil, ErrInvalidId
	}
	return s.bookRepo.FindAllByAuthor(ctx, id)
}
