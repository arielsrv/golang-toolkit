package restclient

type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return e.Message
}

type APINotFoundError struct {
	StatusCode int
	Message    string
}

func (e *APINotFoundError) Error() string {
	return e.Message
}

type APIBadRequestError struct {
	StatusCode int
	Message    string
}

func (e *APIBadRequestError) Error() string {
	return e.Message
}

type APISecurityError struct {
	StatusCode int
	Message    string
}

func (e *APISecurityError) Error() string {
	return e.Message
}
