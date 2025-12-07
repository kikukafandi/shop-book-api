package helper

import (
	"errors"
	"net/http"

	"kikukafandi/book-shop-api/internal/domain"
)

// WriteErrorFromDomain maps domain errors to HTTP responses.
func WriteErrorFromDomain(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrBookNotFound):
		WriteError(w, http.StatusNotFound, "book not found")

	case errors.Is(err, domain.ErrUserNotFound):
		WriteError(w, http.StatusNotFound, "user not found")

	case errors.Is(err, domain.ErrOrderNotFound):
		WriteError(w, http.StatusNotFound, "order not found")

	case errors.Is(err, domain.ErrInsufficientStock):
		WriteError(w, http.StatusBadRequest, "insufficient stock")

	case errors.Is(err, domain.ErrInvalidPrice):
		WriteError(w, http.StatusBadRequest, "price must be positive")

	case errors.Is(err, domain.ErrInvalidStock):
		WriteError(w, http.StatusBadRequest, "stock cannot be negative")

	case errors.Is(err, domain.ErrInvalidQuantity):
		WriteError(w, http.StatusBadRequest, "quantity must be positive")

	case errors.Is(err, domain.ErrEmailExists):
		WriteError(w, http.StatusConflict, "email already exists")

	case errors.Is(err, domain.ErrInvalidCredential):
		WriteError(w, http.StatusUnauthorized, "invalid email or password")

	case errors.Is(err, domain.ErrUnauthorized):
		WriteError(w, http.StatusUnauthorized, "unauthorized access")

	default:
		WriteError(w, http.StatusInternalServerError, "internal server error")
	}
}
