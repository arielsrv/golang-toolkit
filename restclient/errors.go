package restclient

import "net/http"

type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return e.Message
}

type APINotFoundError struct {
	APIError
}

func (e *APINotFoundError) Error() string {
	return e.Message
}

type APIBadRequestError struct {
	APIError
}

func (e *APIBadRequestError) Error() string {
	return e.Message
}

type APISecurityError struct {
	APIError
}

func (e *APISecurityError) Error() string {
	return e.Message
}

func IsOk(statusCode int) bool {
	return statusCode < http.StatusBadRequest
}
