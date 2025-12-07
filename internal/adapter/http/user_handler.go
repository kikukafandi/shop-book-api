package http

import (
	"net/http"

	"kikukafandi/book-shop-api/internal/helper"
	"kikukafandi/book-shop-api/internal/usecase"

	"github.com/julienschmidt/httprouter"
)

// UserHandler handles HTTP requests for users.
type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

// RegisterRequest is the request body for user registration.
type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// LoginRequest is the request body for user login.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserResponse is the response body for user operations.
type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// Register handles POST /register.
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req RegisterRequest
	if err := helper.ReadJSON(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	input := usecase.RegisterInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	output, err := h.userUsecase.Register(r.Context(), input)
	if err != nil {
		helper.WriteErrorFromDomain(w, err)
		return
	}

	resp := toUserResponse(output)
	helper.WriteJSON(w, http.StatusCreated, helper.Response{
		Code:   http.StatusCreated,
		Status: "success",
		Data:   resp,
	})
}

// Login handles POST /login.
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req LoginRequest
	if err := helper.ReadJSON(r, &req); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	input := usecase.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	}

	output, err := h.userUsecase.Login(r.Context(), input)
	if err != nil {
		helper.WriteErrorFromDomain(w, err)
		return
	}

	resp := toUserResponse(output)
	helper.WriteJSON(w, http.StatusOK, helper.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   resp,
	})
}

// toUserResponse converts usecase output to HTTP response.
func toUserResponse(output usecase.UserOutput) UserResponse {
	return UserResponse{
		ID:    output.ID,
		Name:  output.Name,
		Email: output.Email,
		Role:  output.Role,
	}
}
