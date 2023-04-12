package service

import (
	"context"

	"github.com/Chameleon-m/kvd_grpc/internal/app/library/model"
)

// AuthorRepositoryInterface ...
type AuthorRepositoryInterface interface {
	FindAllByBook(ctx context.Context, id uint64) (model.AuthorList, error)
}

// AuthorService ...
type AuthorService struct {
	authorRepo AuthorRepositoryInterface
}

// NewAuthorService создаём сервис авторов
func NewAuthorService(authorRepo AuthorRepositoryInterface) *AuthorService {
	return &AuthorService{
		authorRepo: authorRepo,
	}
}

// Список авторов по книге
func (s AuthorService) GetListByBook(ctx context.Context, id uint64) (model.AuthorList, error) {
	return s.authorRepo.FindAllByBook(ctx, id)
}
