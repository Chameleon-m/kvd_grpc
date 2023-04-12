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
	var list model.AuthorList
	// Подготавливаем запрос
	stmt, err := r.db.Prepare(`
		SELECT author.id, author.name FROM author 
		JOIN book_author ON author.id = book_author.author_id
		WHERE book_author.book_id = ?
	`)
	if err != nil {
		return nil, err
	}
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
		var author model.Author
		if err := rows.Scan(&author.ID, &author.Name); err != nil {
			return nil, err
		}
		// добавляем в список
		list = append(list, &author)
	}
	// Проверяем ошибки в запросе
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}
