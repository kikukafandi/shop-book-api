package service

import (
	"kikukafandi/book-shop-api/repository"

	"gorm.io/gorm"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	DB             *gorm.DB
}
