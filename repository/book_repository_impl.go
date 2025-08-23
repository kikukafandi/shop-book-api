package repository

import (
	"context"
	"kikukafandi/book-shop-api/helper"
	"kikukafandi/book-shop-api/model/domain"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
}

func (repository *BookRepositoryImpl) Save(ctx context.Context, db *gorm.DB, book domain.Book) (domain.Book, error) {
	err := db.WithContext(ctx).Save(&book).Error
	helper.PanicIfError(err)

	return book, nil
}

func (repository *BookRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]domain.Book, error) {
	var books []domain.Book
	err := db.WithContext(ctx).Find(&books).Error
	helper.PanicIfError(err)
	return books, nil
}

func (repository *BookRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, bookId uint) (domain.Book, error) {
	var book domain.Book
	err := db.WithContext(ctx).First(&book, bookId).Error
	helper.PanicIfError(err)
	return book, nil
}
func (repository *BookRepositoryImpl) Update(ctx context.Context, db *gorm.DB, book domain.Book) (domain.Book, error) {
	err := db.WithContext(ctx).Save(&book).Error
	helper.PanicIfError(err)
	return book, nil
}
func (repository *BookRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, book domain.Book) {
	err := db.WithContext(ctx).Delete(&book).Error
	helper.PanicIfError(err)
}
