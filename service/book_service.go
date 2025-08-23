package service

import (
	"context"
	"kikukafandi/book-shop-api/web"
)

type BookService interface {
	Create(ctx context.Context, request web.BookCreateRequest) web.BookResponse
	Update(ctx context.Context, request web.BookUpdateRequest) web.BookResponse
	Delete(ctx context.Context, Bookid int) web.BookResponse
	FindById(ctx context.Context, BookId int)
	FindAll(ctx context.Context) []web.BookResponse
}
