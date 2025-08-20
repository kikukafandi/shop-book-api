package repository

import (
	"context"
	"kikukafandi/book-shop-api/model/domain"

	"gorm.io/gorm"
)

type BookRepository interface {
	Save(ctx context.Context, db *gorm.DB, book domain.Book) domain.Book
	FindAll(ctx context.Context, db *gorm.DB, book domain.Book) []domain.Book
	FindById(ctx context.Context, db *gorm.DB, bookId uint) domain.Book
	Update(ctx context.Context, db *gorm.DB, book domain.Book)
	Delete(ctx context.Context, db *gorm.DB, book domain.Book)
}
