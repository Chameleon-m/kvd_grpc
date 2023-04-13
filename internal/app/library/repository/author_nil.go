package repository

import (
	"context"

	"github.com/Chameleon-m/kvd_grpc/internal/app/library/model"
)

// AuthorNilRepository ...
type AuthorNilRepository struct {
	Places model.AuthorList
}

// NewAuthorNilRepository созаём репозиторий
func NewAuthorNilRepository() *AuthorNilRepository {
	return &AuthorNilRepository{}
}

// FindAllByBook список авторов по книге
func (r *AuthorNilRepository) FindAllByBook(ctx context.Context, id uint64) (model.AuthorList, error) {
	var list model.AuthorList
	return list, nil
}
