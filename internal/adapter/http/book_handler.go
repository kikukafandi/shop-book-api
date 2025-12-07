package http

import (
	"net/http"
	"strconv"

	"kikukafandi/book-shop-api/internal/helper"
	"kikukafandi/book-shop-api/internal/usecase"

	"github.com/julienschmidt/httprouter"
)

// BookHandler handles HTTP requests for books.
type BookHandler struct {
	bookUsecase *usecase.BookUsecase
}

// NewBookHandler creates a new BookHandler.
func NewBookHandler(bookUsecase *usecase.BookUsecase) *BookHandler {
	return &BookHandler{
		bookUsecase: bookUsecase,
	}
}

// CreateBookRequest is the request body for creating a book.
type CreateBookRequest struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

// UpdateBookRequest is the request body for updating a book.
type UpdateBookRequest struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

// BookResponse is the response body for book operations.
type BookResponse struct {
	ID    uint    `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

// Create handles POST /books.
func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req CreateBookRequest
	if err := helper.ReadJSON(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	input := usecase.CreateBookInput{
		Title: req.Title,
		Price: req.Price,
		Stock: req.Stock,
	}

	output, err := h.bookUsecase.Create(r.Context(), input)
	if err != nil {
		helper.WriteErrorFromDomain(w, err)
		return
	}

	resp := toBookResponse(output)
	helper.WriteJSON(w, http.StatusCreated, helper.Response{
		Code:   http.StatusCreated,
		Status: "success",
		Data:   resp,
	})
}

// FindByID handles GET /books/:id.
func (h *BookHandler) FindByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseUint(ps.ByName("id"), 10, 32)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid book id")
		return
	}

	output, err := h.bookUsecase.FindByID(r.Context(), uint(id))
	if err != nil {
		helper.WriteErrorFromDomain(w, err)
		return
	}

	resp := toBookResponse(output)
	helper.WriteJSON(w, http.StatusOK, helper.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   resp,
	})
}

// FindAll handles GET /books.
func (h *BookHandler) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	outputs, err := h.bookUsecase.FindAll(r.Context())
	if err != nil {
		helper.WriteErrorFromDomain(w, err)
		return
	}

	var responses []BookResponse
	for _, output := range outputs {
		responses = append(responses, toBookResponse(output))
	}

	helper.WriteJSON(w, http.StatusOK, helper.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   responses,
	})
}

// Update handles PUT /books/:id.
func (h *BookHandler) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseUint(ps.ByName("id"), 10, 32)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid book id")
		return
	}

	var req UpdateBookRequest
	if err := helper.ReadJSON(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	input := usecase.UpdateBookInput{
		ID:    uint(id),
		Title: req.Title,
		Price: req.Price,
		Stock: req.Stock,
	}

	output, err := h.bookUsecase.Update(r.Context(), input)
	if err != nil {
		helper.WriteErrorFromDomain(w, err)
		return
	}

	resp := toBookResponse(output)
	helper.WriteJSON(w, http.StatusOK, helper.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   resp,
	})
}

// Delete handles DELETE /books/:id.
func (h *BookHandler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseUint(ps.ByName("id"), 10, 32)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid book id")
		return
	}

	if err := h.bookUsecase.Delete(r.Context(), uint(id)); err != nil {
		helper.WriteErrorFromDomain(w, err)
		return
	}

	helper.WriteJSON(w, http.StatusOK, helper.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   nil,
	})
}

// toBookResponse converts usecase output to HTTP response.
func toBookResponse(output usecase.BookOutput) BookResponse {
	return BookResponse{
		ID:    output.ID,
		Title: output.Title,
		Price: output.Price,
		Stock: output.Stock,
	}
}
