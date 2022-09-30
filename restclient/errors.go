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

type APIBadRequestError struct {
	APIError
}

type APISecurityError struct {
	APIError
}

func IsOk(statusCode int) bool {
	return statusCode < http.StatusBadRequest
}
