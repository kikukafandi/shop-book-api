package db

import (
	"context"
	"time"

	"kikukafandi/book-shop-api/internal/domain"

	"gorm.io/gorm"
)

// OrderModel is the database model for Order.
type OrderModel struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null;index"`
	BookID    uint      `gorm:"not null;index"`
	Quantity  int       `gorm:"not null"`
	Total     float64   `gorm:"not null"`
	Status    string    `gorm:"size:50;not null"`
	CreatedAt time.Time `gorm:"not null"`
}

// TableName returns the table name for OrderModel.
func (OrderModel) TableName() string {
	return "orders"
}

// OrderRepositoryMySQL implements domain.OrderRepository using MySQL/GORM.
type OrderRepositoryMySQL struct {
	db *gorm.DB
}

// NewOrderRepositoryMySQL creates a new OrderRepositoryMySQL.
func NewOrderRepositoryMySQL(db *gorm.DB) *OrderRepositoryMySQL {
	return &OrderRepositoryMySQL{db: db}
}

// Save saves an order to database.
func (r *OrderRepositoryMySQL) Save(ctx context.Context, order domain.Order) (domain.Order, error) {
	model := toOrderModel(order)

	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return domain.Order{}, err
	}

	return toOrderDomain(model), nil
}

// FindByID finds an order by ID.
func (r *OrderRepositoryMySQL) FindByID(ctx context.Context, id uint) (domain.Order, error) {
	var model OrderModel

	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Order{}, domain.ErrOrderNotFound
		}
		return domain.Order{}, err
	}

	return toOrderDomain(model), nil
}

// FindByUserID finds all orders by user ID.
func (r *OrderRepositoryMySQL) FindByUserID(ctx context.Context, userID uint) ([]domain.Order, error) {
	var models []OrderModel

	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&models).Error; err != nil {
		return nil, err
	}

	orders := make([]domain.Order, len(models))
	for i, model := range models {
		orders[i] = toOrderDomain(model)
	}

	return orders, nil
}

// FindAll returns all orders.
func (r *OrderRepositoryMySQL) FindAll(ctx context.Context) ([]domain.Order, error) {
	var models []OrderModel

	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}

	orders := make([]domain.Order, len(models))
	for i, model := range models {
		orders[i] = toOrderDomain(model)
	}

	return orders, nil
}

// Update updates an order in database.
func (r *OrderRepositoryMySQL) Update(ctx context.Context, order domain.Order) (domain.Order, error) {
	model := toOrderModel(order)

	if err := r.db.WithContext(ctx).Save(&model).Error; err != nil {
		return domain.Order{}, err
	}

	return toOrderDomain(model), nil
}

// Delete deletes an order from database.
func (r *OrderRepositoryMySQL) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&OrderModel{}, id).Error; err != nil {
		return err
	}
	return nil
}

// toOrderModel converts domain.Order to OrderModel.
func toOrderModel(order domain.Order) OrderModel {
	return OrderModel{
		ID:        order.ID,
		UserID:    order.UserID,
		BookID:    order.BookID,
		Quantity:  order.Quantity,
		Total:     order.Total,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
	}
}

// toOrderDomain converts OrderModel to domain.Order.
func toOrderDomain(model OrderModel) domain.Order {
	return domain.Order{
		ID:        model.ID,
		UserID:    model.UserID,
		BookID:    model.BookID,
		Quantity:  model.Quantity,
		Total:     model.Total,
		Status:    model.Status,
		CreatedAt: model.CreatedAt,
	}
}
