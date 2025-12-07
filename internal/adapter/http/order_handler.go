package http

import (
	"net/http"
	"strconv"

	"kikukafandi/book-shop-api/internal/helper"
	"kikukafandi/book-shop-api/internal/usecase"

	"github.com/julienschmidt/httprouter"
)

// OrderHandler handles HTTP requests for orders.
type OrderHandler struct {
	orderUsecase *usecase.OrderUsecase
}

// NewOrderHandler creates a new OrderHandler.
func NewOrderHandler(orderUsecase *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{
		orderUsecase: orderUsecase,
	}
}

// CreateOrderRequest is the request body for creating an order.
type CreateOrderRequest struct {
	UserID   uint `json:"user_id"`
	BookID   uint `json:"book_id"`
	Quantity int  `json:"quantity"`
}

// OrderResponse is the response body for order operations.
type OrderResponse struct {
	ID       uint    `json:"id"`
	UserID   uint    `json:"user_id"`
	BookID   uint    `json:"book_id"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
	Status   string  `json:"status"`
}

// Create handles POST /orders.
func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req CreateOrderRequest
	if err := helper.ReadJSON(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	input := usecase.CreateOrderInput{
		UserID:   req.UserID,
		BookID:   req.BookID,
		Quantity: req.Quantity,
	}

	output, err := h.orderUsecase.Create(r.Context(), input)
	if err != nil {
		helper.WriteErrorFromDomain(w, err)
		return
	}

	resp := toOrderResponse(output)
	helper.WriteJSON(w, http.StatusCreated, helper.Response{
		Code:   http.StatusCreated,
		Status: "success",
		Data:   resp,
	})
}

// FindByID handles GET /orders/:id.
func (h *OrderHandler) FindByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseUint(ps.ByName("id"), 10, 32)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid order id")
		return
	}

	output, err := h.orderUsecase.FindByID(r.Context(), uint(id))
	if err != nil {
		helper.WriteErrorFromDomain(w, err)
		return
	}

	resp := toOrderResponse(output)
	helper.WriteJSON(w, http.StatusOK, helper.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   resp,
	})
}

// FindAll handles GET /orders.
func (h *OrderHandler) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	outputs, err := h.orderUsecase.FindAll(r.Context())
	if err != nil {
		helper.WriteErrorFromDomain(w, err)
		return
	}

	var responses []OrderResponse
	for _, output := range outputs {
		responses = append(responses, toOrderResponse(output))
	}

	helper.WriteJSON(w, http.StatusOK, helper.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   responses,
	})
}

// FindByUserID handles GET /users/:userId/orders.
func (h *OrderHandler) FindByUserID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, err := strconv.ParseUint(ps.ByName("userId"), 10, 32)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	outputs, err := h.orderUsecase.FindByUserID(r.Context(), uint(userID))
	if err != nil {
		helper.WriteErrorFromDomain(w, err)
		return
	}

	var responses []OrderResponse
	for _, output := range outputs {
		responses = append(responses, toOrderResponse(output))
	}

	helper.WriteJSON(w, http.StatusOK, helper.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   responses,
	})
}

// toOrderResponse converts usecase output to HTTP response.
func toOrderResponse(output usecase.OrderOutput) OrderResponse {
	return OrderResponse{
		ID:       output.ID,
		UserID:   output.UserID,
		BookID:   output.BookID,
		Quantity: output.Quantity,
		Total:    output.Total,
		Status:   output.Status,
	}
}
