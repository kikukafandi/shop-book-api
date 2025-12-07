package usecase

import (
	"context"

	"kikukafandi/book-shop-api/internal/domain"
)

// BookUsecase handles all book business logic.
type BookUsecase struct {
	bookRepo domain.BookRepository
}

// NewBookUsecase creates a new BookUsecase.
func NewBookUsecase(bookRepo domain.BookRepository) *BookUsecase {
	return &BookUsecase{
		bookRepo: bookRepo,
	}
}

// CreateBookInput is the input for creating a book.
type CreateBookInput struct {
	Title string
	Price float64
	Stock int
}

// UpdateBookInput is the input for updating a book.
type UpdateBookInput struct {
	ID    uint
	Title string
	Price float64
	Stock int
}

// BookOutput is the output for book operations.
type BookOutput struct {
	ID    uint
	Title string
	Price float64
	Stock int
}

// Create creates a new book with validation.
func (u *BookUsecase) Create(ctx context.Context, input CreateBookInput) (BookOutput, error) {
	// Business rule: price must be positive
	if input.Price <= 0 {
		return BookOutput{}, domain.ErrInvalidPrice
	}

	// Business rule: stock cannot be negative
	if input.Stock < 0 {
		return BookOutput{}, domain.ErrInvalidStock
	}

	book := domain.NewBook(input.Title, input.Price, input.Stock)

	saved, err := u.bookRepo.Save(ctx, book)
	if err != nil {
		return BookOutput{}, err
	}

	return toBookOutput(saved), nil
}

// FindByID finds a book by ID.
func (u *BookUsecase) FindByID(ctx context.Context, id uint) (BookOutput, error) {
	book, err := u.bookRepo.FindByID(ctx, id)
	if err != nil {
		return BookOutput{}, err
	}

	return toBookOutput(book), nil
}

// FindAll returns all books.
func (u *BookUsecase) FindAll(ctx context.Context) ([]BookOutput, error) {
	books, err := u.bookRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	outputs := make([]BookOutput, len(books))
	for i, book := range books {
		outputs[i] = toBookOutput(book)
	}

	return outputs, nil
}

// Update updates an existing book.
func (u *BookUsecase) Update(ctx context.Context, input UpdateBookInput) (BookOutput, error) {
	// Business rule: price must be positive
	if input.Price <= 0 {
		return BookOutput{}, domain.ErrInvalidPrice
	}

	// Business rule: stock cannot be negative
	if input.Stock < 0 {
		return BookOutput{}, domain.ErrInvalidStock
	}

	// Check if book exists
	_, err := u.bookRepo.FindByID(ctx, input.ID)
	if err != nil {
		return BookOutput{}, err
	}

	book := domain.Book{
		ID:    input.ID,
		Title: input.Title,
		Price: input.Price,
		Stock: input.Stock,
	}

	updated, err := u.bookRepo.Update(ctx, book)
	if err != nil {
		return BookOutput{}, err
	}

	return toBookOutput(updated), nil
}

// Delete deletes a book by ID.
func (u *BookUsecase) Delete(ctx context.Context, id uint) error {
	// Check if book exists
	_, err := u.bookRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return u.bookRepo.Delete(ctx, id)
}

// toBookOutput converts domain.Book to BookOutput.
func toBookOutput(book domain.Book) BookOutput {
	return BookOutput{
		ID:    book.ID,
		Title: book.Title,
		Price: book.Price,
		Stock: book.Stock,
	}
}
