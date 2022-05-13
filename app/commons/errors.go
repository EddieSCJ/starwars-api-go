package commons

import (
	"net/http"
)

type HTTPError struct {
	Code           int      `json:"code"`
	Message        string   `json:"message"`
	AdditionalInfo []string `json:"additionalInfo"`
}

func NewBadRequest(message string) *HTTPError {
	return &HTTPError{
		Code:    http.StatusBadRequest,
		Message: "invalid request",
		AdditionalInfo: []string{
			message,
		},
	}
}

func NewInternalServerError() *HTTPError {
	return &HTTPError{
		Code:    http.StatusInternalServerError,
		Message: "an unknown error occurred",
	}
}
