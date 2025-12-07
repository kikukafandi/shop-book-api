package helper

import (
	"encoding/json"
	"net/http"
)

// Response is the standard API response structure.
type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

// ErrorResponse is the error response structure.
type ErrorResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// WriteJSON writes JSON response to http.ResponseWriter.
func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// WriteError writes error response to http.ResponseWriter.
func WriteError(w http.ResponseWriter, statusCode int, message string) {
	resp := ErrorResponse{
		Code:    statusCode,
		Status:  "error",
		Message: message,
	}
	WriteJSON(w, statusCode, resp)
}

// ReadJSON reads JSON from request body into target.
func ReadJSON(r *http.Request, target interface{}) error {
	return json.NewDecoder(r.Body).Decode(target)
}
