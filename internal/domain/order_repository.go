package domain

import "context"

// OrderRepository is the port (interface) for order persistence.
type OrderRepository interface {
	Save(ctx context.Context, order Order) (Order, error)
	FindByID(ctx context.Context, id uint) (Order, error)
	FindByUserID(ctx context.Context, userID uint) ([]Order, error)
	FindAll(ctx context.Context) ([]Order, error)
	Update(ctx context.Context, order Order) (Order, error)
	Delete(ctx context.Context, id uint) error
}
