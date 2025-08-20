package domain

type Book struct {
	ID    uint `gorm:"primaryKey"`
	Title string
	Price float64
	Stock int
}
