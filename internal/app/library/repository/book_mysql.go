package repository

import (
	"context"
	"database/sql"

	"github.com/Chameleon-m/kvd_grpc/internal/app/library/model"
)

// BookMysqlRepository ...
type BookMysqlRepository struct {
	db *sql.DB
}

// NewBookMysqlRepository созаём репозиторий
func NewBookMysqlRepository(db *sql.DB) *BookMysqlRepository {
	return &BookMysqlRepository{
		db: db,
	}
}

// FindAllByAuthor список книг по автору
func (r *BookMysqlRepository) FindAllByAuthor(ctx context.Context, id uint64) (model.BookList, error) {

	return nil, nil
}
