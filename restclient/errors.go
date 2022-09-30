package restclient

import "net/http"

type APIError struct {
	StatusCode int
	Message    string
}

func NewAPIError(statusCode int, message string) *APIError {
	return &APIError{StatusCode: statusCode, Message: message}
}

func (e *APIError) Error() string {
	return e.Message
}

type APINotFoundError struct {
	APIError
}

type APIBadRequestError struct {
	APIError
}

type APISecurityError struct {
	APIError
}

func IsOk(statusCode int) bool {
	return statusCode < http.StatusBadRequest
}
