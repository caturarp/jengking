package main

import "net/http"

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func (c *CustomError) Error() string {
	return c.Message
}

var (
	ErrFailedDecode       = NewCustomError(http.StatusInternalServerError, "failed to decode")
	ErrInvalidTypeWebhook = NewCustomError(http.StatusBadRequest, "wrong github webhook payload")
)
