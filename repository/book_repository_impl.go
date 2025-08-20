package repository

import (
	"context"
	"kikukafandi/book-shop-api/model/domain"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
}

func (repository *BookRepositoryImpl) Save(ctx context.Context, db *gorm.DB, book domain.Book) (domain.Book, error) {
	err := db.WithContext(ctx).Save(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func (repository *BookRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB, book domain.Book) []domain.Book {
	err := db.WithContext(ctx).Find(&book).Error
	if err != nil {
		return []domain.Book{}
	}

	return []domain.Book{book}
}
