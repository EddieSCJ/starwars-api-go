package commons

import (
	"net/http"
)

type HttpError struct {
	Code           int      `json:"code"`
	Message        string   `json:"message"`
	AdditionalInfo []string `json:"additionalInfo"`
}

func NewBadRequest(message string) *HttpError {
	return &HttpError{
		Code:    http.StatusBadRequest,
		Message: "invalid request",
		AdditionalInfo: []string{
			message,
		},
	}
}

func NewInternalServerError() *HttpError {
	return &HttpError{
		Code:    http.StatusInternalServerError,
		Message: "an unknown error occurred",
	}
}
