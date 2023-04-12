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
	var list model.BookList
	// Формеруем запрос
	stmt, err := r.db.Prepare(`
		SELECT book.id, book.name FROM book 
		JOIN book_author ON book.id = book_author.book_id
		WHERE book_author.author_id = ?
	`)
	if err != nil {
		return nil, err
	}
	// Освобождаем ресурсы
	defer stmt.Close()

	// Выполняем запрос
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	// Освобождаем ресурсы
	defer rows.Close()

	// Прокручиваем, назначаем используя Scan данные столбца полям структуры.
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.ID, &book.Name); err != nil {
			return nil, err
		}
		// добавляем в список
		list = append(list, &book)
	}
	// Проверяем ошибки в запросе
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}
