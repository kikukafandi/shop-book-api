package web

type BookCreateRequest struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}
