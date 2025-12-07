package db

import (
	"context"

	"kikukafandi/book-shop-api/internal/domain"

	"gorm.io/gorm"
)

// BookModel is the database model for Book.
// GORM tags are only here in adapter layer - domain stays clean.
type BookModel struct {
	ID    uint    `gorm:"primaryKey"`
	Title string  `gorm:"size:255;not null"`
	Price float64 `gorm:"not null"`
	Stock int     `gorm:"not null"`
}

// TableName returns the table name for BookModel.
func (BookModel) TableName() string {
	return "books"
}

// BookRepositoryMySQL implements domain.BookRepository using MySQL/GORM.
type BookRepositoryMySQL struct {
	db *gorm.DB
}

// NewBookRepositoryMySQL creates a new BookRepositoryMySQL.
func NewBookRepositoryMySQL(db *gorm.DB) *BookRepositoryMySQL {
	return &BookRepositoryMySQL{db: db}
}

// Save saves a book to database.
func (r *BookRepositoryMySQL) Save(ctx context.Context, book domain.Book) (domain.Book, error) {
	model := toBookModel(book)

	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return domain.Book{}, err
	}

	return toBookDomain(model), nil
}

// FindByID finds a book by ID.
func (r *BookRepositoryMySQL) FindByID(ctx context.Context, id uint) (domain.Book, error) {
	var model BookModel

	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Book{}, domain.ErrBookNotFound
		}
		return domain.Book{}, err
	}

	return toBookDomain(model), nil
}

// FindAll returns all books.
func (r *BookRepositoryMySQL) FindAll(ctx context.Context) ([]domain.Book, error) {
	var models []BookModel

	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}

	books := make([]domain.Book, len(models))
	for i, model := range models {
		books[i] = toBookDomain(model)
	}

	return books, nil
}

// Update updates a book in database.
func (r *BookRepositoryMySQL) Update(ctx context.Context, book domain.Book) (domain.Book, error) {
	model := toBookModel(book)

	if err := r.db.WithContext(ctx).Save(&model).Error; err != nil {
		return domain.Book{}, err
	}

	return toBookDomain(model), nil
}

// Delete deletes a book from database.
func (r *BookRepositoryMySQL) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&BookModel{}, id).Error; err != nil {
		return err
	}
	return nil
}

// toBookModel converts domain.Book to BookModel.
func toBookModel(book domain.Book) BookModel {
	return BookModel{
		ID:    book.ID,
		Title: book.Title,
		Price: book.Price,
		Stock: book.Stock,
	}
}

// toBookDomain converts BookModel to domain.Book.
func toBookDomain(model BookModel) domain.Book {
	return domain.Book{
		ID:    model.ID,
		Title: model.Title,
		Price: model.Price,
		Stock: model.Stock,
	}
}
