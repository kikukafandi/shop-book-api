package domain

import "time"

// Order represents the order entity in domain layer.
type Order struct {
	ID        uint
	UserID    uint
	BookID    uint
	Quantity  int
	Total     float64
	Status    string
	CreatedAt time.Time
}

// OrderStatus constants.
const (
	OrderStatusPending   = "pending"
	OrderStatusCompleted = "completed"
	OrderStatusCancelled = "cancelled"
)

// NewOrder creates a new Order entity.
func NewOrder(userID, bookID uint, quantity int, total float64) Order {
	return Order{
		UserID:    userID,
		BookID:    bookID,
		Quantity:  quantity,
		Total:     total,
		Status:    OrderStatusPending,
		CreatedAt: time.Now(),
	}
}

// Complete marks order as completed.
func (o *Order) Complete() {
	o.Status = OrderStatusCompleted
}

// Cancel marks order as cancelled.
func (o *Order) Cancel() {
	o.Status = OrderStatusCancelled
}

// IsPending checks if order is still pending.
func (o Order) IsPending() bool {
	return o.Status == OrderStatusPending
}
