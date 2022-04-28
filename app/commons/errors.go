package commons

import "net/http"

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

func NewServiceUnavailableError(client, message string) *HttpError {
	return &HttpError{
		Code:    http.StatusServiceUnavailable,
		Message: client + ": unavailable service",
		AdditionalInfo: []string{
			message,
		},
	}
}

func NewNotFoundError(client, message string) *HttpError {
	return &HttpError{
		Code:    http.StatusNotFound,
		Message: client + ": resource not found",
		AdditionalInfo: []string{
			message,
		},
	}
}
