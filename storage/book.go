package storage

import (
	"database/sql"
	"time"
)

type DBManager struct {
	db *sql.DB
}

func NewDBManager(db *sql.DB) *DBManager {
	return &DBManager{db: db}
}

type Book struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	AuthorName string    `json:"author_name"`
	Price      float32   `json:"amount"`
	Amount     int       `json:"price"`
	CreatedAt  time.Time `json:"created_at"`
}

func (b *DBManager) CreateBook(book *Book) (*Book, error) {

	query := `
		INSERT INTO books (
			title,
			author_name,
			price,
			amount
		) VALUES ($1, $2, $3, $4)
		RETURNING id, title, author_name, price, amount, created_at
	`
	row := b.db.QueryRow(
		query,
		book.Title,
		book.AuthorName,
		book.Price,
		book.Amount,
	)

	var res Book
	err := row.Scan(
		&res.Id,
		&res.Title,
		&res.AuthorName,
		&res.Price,
		&res.Amount,
		&res.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &res, err
}

func (m *DBManager) GetBook(id int) (*Book, error) {
	var res Book

	query := `
		SELECT
			id,
			title,
			author_name,
			price,
			amount,
			created_at
		FROM books
		WHERE id=$1
	`

	row := m.db.QueryRow(
		query,
		id,
	)

	err := row.Scan(
		&res.Id,
		&res.Title,
		&res.AuthorName,
		&res.Price,
		&res.Amount,
		&res.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
