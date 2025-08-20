package repository

import (
	"context"
	"database/sql"
	"kikukafandi/book-shop-api/model/domain"
)

type BookRepository interface {
	Save(ctx context.Context, tx sql.Tx, book domain.Book) domain.Book
	FindAll(ctx context.Context, tx sql.Tx, book domain.Book) []domain.Book
	FindById(ctx context.Context, tx sql.Tx, bookId uint) domain.Book
	Update(ctx context.Context, tx sql.Tx, book domain.Book)
	Delete(ctx context.Context, tx sql.Tx, book domain.Book)
}
