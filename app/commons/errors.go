package commons

import (
	"net/http"
)

type HttpError struct {
	Code           int      `json:"code"`
	Message        string   `json:"message"`
	AdditionalInfo []string `json:"additionalInfo"`
}

func NewBadGatewayError(client, message string) *HttpError {
	return &HttpError{
		Code:    http.StatusBadGateway,
		Message: client + ": error while requesting data",
		AdditionalInfo: []string{
			message,
		},
	}
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

func NewServiceUnavailableError(message string) *HttpError {
	return &HttpError{
		Code:    http.StatusServiceUnavailable,
		Message: "unavailable service",
		AdditionalInfo: []string{
			message,
		},
	}
}

func NewNotFoundError(message string) *HttpError {
	return &HttpError{
		Code:    http.StatusNotFound,
		Message: "resource not found",
		AdditionalInfo: []string{
			message,
		},
	}
}

func NewInternalServerError(message string) *HttpError {
	return &HttpError{
		Code:    http.StatusInternalServerError,
		Message: "an unknown error occurred",
		AdditionalInfo: []string{
			message,
		},
	}
}
