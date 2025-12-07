package usecase

import (
	"context"

	"kikukafandi/book-shop-api/internal/domain"
)

// UserUsecase handles all user business logic.
type UserUsecase struct {
	userRepo domain.UserRepository
}

// NewUserUsecase creates a new UserUsecase.
func NewUserUsecase(userRepo domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

// RegisterInput is the input for user registration.
type RegisterInput struct {
	Name     string
	Email    string
	Password string
	Role     string
}

// LoginInput is the input for user login.
type LoginInput struct {
	Email    string
	Password string
}

// UserOutput is the output for user operations.
type UserOutput struct {
	ID    uint
	Name  string
	Email string
	Role  string
}

// Register creates a new user with validation.
func (u *UserUsecase) Register(ctx context.Context, input RegisterInput) (UserOutput, error) {
	// Check if email already exists
	exists, err := u.userRepo.ExistsByEmail(ctx, input.Email)
	if err != nil {
		return UserOutput{}, err
	}
	if exists {
		return UserOutput{}, domain.ErrEmailExists
	}

	// In real app, hash password here
	user := domain.NewUser(input.Name, input.Email, input.Password, input.Role)

	saved, err := u.userRepo.Save(ctx, user)
	if err != nil {
		return UserOutput{}, err
	}

	return toUserOutput(saved), nil
}

// Login authenticates user and returns user data.
func (u *UserUsecase) Login(ctx context.Context, input LoginInput) (UserOutput, error) {
	user, err := u.userRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		return UserOutput{}, domain.ErrInvalidCredential
	}

	// In real app, compare hashed password
	if user.Password != input.Password {
		return UserOutput{}, domain.ErrInvalidCredential
	}

	return toUserOutput(user), nil
}

// FindByID finds a user by ID.
func (u *UserUsecase) FindByID(ctx context.Context, id uint) (UserOutput, error) {
	user, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return UserOutput{}, err
	}

	return toUserOutput(user), nil
}

// FindAll returns all users.
func (u *UserUsecase) FindAll(ctx context.Context) ([]UserOutput, error) {
	users, err := u.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	outputs := make([]UserOutput, len(users))
	for i, user := range users {
		outputs[i] = toUserOutput(user)
	}

	return outputs, nil
}

// toUserOutput converts domain.User to UserOutput.
func toUserOutput(user domain.User) UserOutput {
	return UserOutput{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}
