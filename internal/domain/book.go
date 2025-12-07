package domain

// Book represents the book entity in domain layer.
// No framework imports, no GORM tags - pure domain model.
type Book struct {
	ID    uint
	Title string
	Price float64
	Stock int
}

// NewBook creates a new Book entity.
func NewBook(title string, price float64, stock int) Book {
	return Book{
		Title: title,
		Price: price,
		Stock: stock,
	}
}

// IsAvailable checks if book has stock.
func (b Book) IsAvailable() bool {
	return b.Stock > 0
}

// DecreaseStock decreases book stock by given amount.
func (b *Book) DecreaseStock(amount int) error {
	if b.Stock < amount {
		return ErrInsufficientStock
	}
	b.Stock -= amount
	return nil
}

// IncreaseStock increases book stock by given amount.
func (b *Book) IncreaseStock(amount int) {
	b.Stock += amount
}
