package errors

import (
	"errors"
	"net/http"
)

type RestError struct {
	Message string `json: "message"`
	Status  int    `json: "status"`
	Error   string `json: "error"`
}

func New(msg string) error {
	return errors.New(msg)
}

func NewBadRequest(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFound(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
