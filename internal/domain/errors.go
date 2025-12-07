package domain

import "errors"

// Domain errors - these are business rule violations.
var (
	ErrBookNotFound      = errors.New("book not found")
	ErrUserNotFound      = errors.New("user not found")
	ErrOrderNotFound     = errors.New("order not found")
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrInvalidPrice      = errors.New("price must be positive")
	ErrInvalidStock      = errors.New("stock cannot be negative")
	ErrInvalidQuantity   = errors.New("quantity must be positive")
	ErrEmailExists       = errors.New("email already exists")
	ErrInvalidCredential = errors.New("invalid email or password")
	ErrUnauthorized      = errors.New("unauthorized access")
)
