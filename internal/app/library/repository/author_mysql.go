package repository

import (
	"context"
	"database/sql"

	"github.com/Chameleon-m/kvd_grpc/internal/app/library/model"
)

// AuthorMysqlRepository ...
type AuthorMysqlRepository struct {
	db *sql.DB
}

// NewAuthorMysqlRepository созаём репозиторий
func NewAuthorMysqlRepository(db *sql.DB) *AuthorMysqlRepository {
	return &AuthorMysqlRepository{
		db: db,
	}
}

// FindAllByBook список авторов по книге
func (r *AuthorMysqlRepository) FindAllByBook(ctx context.Context, id uint64) (model.AuthorList, error) {

	return nil, nil
}
