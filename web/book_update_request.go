package web

type BookUpdateRequest struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}
