package domain

import "context"

// UserRepository is the port (interface) for user persistence.
type UserRepository interface {
	Save(ctx context.Context, user User) (User, error)
	FindByID(ctx context.Context, id uint) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	FindAll(ctx context.Context) ([]User, error)
	Update(ctx context.Context, user User) (User, error)
	Delete(ctx context.Context, id uint) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}
