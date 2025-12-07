package usecase

import (
	"context"

	"kikukafandi/book-shop-api/internal/domain"
)

// OrderUsecase handles all order business logic.
type OrderUsecase struct {
	orderRepo domain.OrderRepository
	bookRepo  domain.BookRepository
	userRepo  domain.UserRepository
}

// NewOrderUsecase creates a new OrderUsecase.
func NewOrderUsecase(
	orderRepo domain.OrderRepository,
	bookRepo domain.BookRepository,
	userRepo domain.UserRepository,
) *OrderUsecase {
	return &OrderUsecase{
		orderRepo: orderRepo,
		bookRepo:  bookRepo,
		userRepo:  userRepo,
	}
}

// CreateOrderInput is the input for creating an order.
type CreateOrderInput struct {
	UserID   uint
	BookID   uint
	Quantity int
}

// OrderOutput is the output for order operations.
type OrderOutput struct {
	ID       uint
	UserID   uint
	BookID   uint
	Quantity int
	Total    float64
	Status   string
}

// Create creates a new order with business validations.
func (u *OrderUsecase) Create(ctx context.Context, input CreateOrderInput) (OrderOutput, error) {
	// Business rule: quantity must be positive
	if input.Quantity <= 0 {
		return OrderOutput{}, domain.ErrInvalidQuantity
	}

	// Check user exists
	_, err := u.userRepo.FindByID(ctx, input.UserID)
	if err != nil {
		return OrderOutput{}, domain.ErrUserNotFound
	}

	// Check book exists and has stock
	book, err := u.bookRepo.FindByID(ctx, input.BookID)
	if err != nil {
		return OrderOutput{}, domain.ErrBookNotFound
	}

	// Business rule: check stock availability
	if err := book.DecreaseStock(input.Quantity); err != nil {
		return OrderOutput{}, err
	}

	// Calculate total
	total := book.Price * float64(input.Quantity)

	// Create order
	order := domain.NewOrder(input.UserID, input.BookID, input.Quantity, total)

	// Update book stock
	_, err = u.bookRepo.Update(ctx, book)
	if err != nil {
		return OrderOutput{}, err
	}

	// Save order
	saved, err := u.orderRepo.Save(ctx, order)
	if err != nil {
		return OrderOutput{}, err
	}

	return toOrderOutput(saved), nil
}

// FindByID finds an order by ID.
func (u *OrderUsecase) FindByID(ctx context.Context, id uint) (OrderOutput, error) {
	order, err := u.orderRepo.FindByID(ctx, id)
	if err != nil {
		return OrderOutput{}, err
	}

	return toOrderOutput(order), nil
}

// FindByUserID finds all orders by user ID.
func (u *OrderUsecase) FindByUserID(ctx context.Context, userID uint) ([]OrderOutput, error) {
	orders, err := u.orderRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	outputs := make([]OrderOutput, len(orders))
	for i, order := range orders {
		outputs[i] = toOrderOutput(order)
	}

	return outputs, nil
}

// FindAll returns all orders.
func (u *OrderUsecase) FindAll(ctx context.Context) ([]OrderOutput, error) {
	orders, err := u.orderRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	outputs := make([]OrderOutput, len(orders))
	for i, order := range orders {
		outputs[i] = toOrderOutput(order)
	}

	return outputs, nil
}

// toOrderOutput converts domain.Order to OrderOutput.
func toOrderOutput(order domain.Order) OrderOutput {
	return OrderOutput{
		ID:       order.ID,
		UserID:   order.UserID,
		BookID:   order.BookID,
		Quantity: order.Quantity,
		Total:    order.Total,
		Status:   order.Status,
	}
}
