package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/ghduuep/litly/internal/domain"
)

type PostgresBookRepository struct {
	db *sql.DB
}

func NewPostgresBookRepository(db *sql.DB) *PostgresBookRepository {
	return &PostgresBookRepository{
		db: db,
	}
}

func (r *PostgresBookRepository) Create(ctx context.Context, u *domain.Book) (*domain.Book, error) {
	query := `INSERT INTO books (title, isbn, description, published_at, genre, pages) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, title, isbn, description, published_at, genre, pages, created_at = NOW(), updated_at = NOW()`

	newBook := &domain.Book{}
	err := r.db.QueryRowContext(ctx, query, u.Title, u.ISBN, u.Description, u.PublishedAt, u.Genre, u.Pages).Scan(
		&newBook.ID,
		&newBook.Title,
		&newBook.ISBN,
		&newBook.Description,
		&newBook.PublishedAt,
		&newBook.Genre,
		&newBook.Pages,
		&newBook.CreatedAt,
		&newBook.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return newBook, nil
}

func (r *PostgresBookRepository) FindByID(ctx context.Context, id int64) (*domain.Book, error) {
	query := `SELECT id, title, isbn, description, published_at, genre, pages, created_at, updated_at FROM books WHERE id = $1`

	book := &domain.Book{}
	err := r.db.QueryRowContext(ctx, query).Scan(
		&book.ID,
		&book.Title,
		&book.ISBN,
		&book.Description,
		&book.PublishedAt,
		&book.Genre,
		&book.Pages,
		&book.CreatedAt,
		&book.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("book not found with id %d", id)
		}
		return nil, err
	}

	return book, nil
}

func (r *PostgresBookRepository) Update(ctx context.Context, id int64, u *domain.User) (*domain.User, error) {
	var setParts []string
	var args []any
	argsPos := 1

	if u.Username != nil {
		setParts = append(setParts, fmt.Sprintf("username = $%d", argsPos))
		args = append(args, u.Username)
		argsPos++
	}
}

func (r *PostgresBookRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM books WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}

func (r *PostgresBookRepository) findAll(ctx context.Context) ([]*domain.User, error) {
	return nil, nil
}
