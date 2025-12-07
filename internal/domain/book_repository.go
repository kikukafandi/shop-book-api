package domain

import "context"

// BookRepository is the port (interface) for book persistence.
// This interface lives in domain - implementations live in adapter/db.
type BookRepository interface {
	Save(ctx context.Context, book Book) (Book, error)
	FindByID(ctx context.Context, id uint) (Book, error)
	FindAll(ctx context.Context) ([]Book, error)
	Update(ctx context.Context, book Book) (Book, error)
	Delete(ctx context.Context, id uint) error
}
