package restclient

import (
	"errors"
	"github.com/arielsrv/golang-toolkit/restclient/hashcode"
	"net/http"
)

type MockError struct {
	Message string
}

func (e *MockError) Error() string {
	return e.Message
}

type Tuple[T any] struct {
	Method   string
	Response *Response[T]
	Error    error
}

func NoNetworkError() error {
	return nil
}

func NetworkError() error {
	return errors.New("network error")
}

type MockResponse[T any] struct {
	Responses map[uint64]Tuple[T]
}

func (mockResponse MockResponse[T]) NewRESTClient() *MockResponse[T] {
	mockResponse.Responses = make(map[uint64]Tuple[T])
	return &mockResponse
}

type MockRequest struct {
	Method string
	URL    string
}

func (mockRequest MockRequest) GetHashCode() uint64 {
	hash := uint64(7)
	hash = uint64(31)*hash + hashcode.String(mockRequest.Method)
	hash = uint64(31)*hash + hashcode.String(mockRequest.URL)
	return hash
}

func (mockResponse MockResponse[T]) AddMockRequest(mockRequest MockRequest, response Response[T], err error) *MockResponse[T] {
	hash := mockRequest.GetHashCode()
	mockResponse.Responses[hash] = Tuple[T]{
		Method:   mockRequest.Method,
		Response: &response,
		Error:    err,
	}
	return &mockResponse
}

func (mockResponse MockResponse[T]) Build() *RESTClient {
	return &RESTClient{
		testingMode: true,
		Mock:        mockResponse.Responses,
	}
}

func (e Execute[T]) GetMock(method string, url string, result Response[T]) (*Response[T], error) {
	mocks, boxing := e.RESTClient.Mock.(map[uint64]Tuple[T])
	if !boxing {
		return &result, &MockError{Message: "Internal mocking error. "}
	}
	mockedRequest := MockRequest{
		Method: method,
		URL:    url,
	}
	mock := mocks[mockedRequest.GetHashCode()]
	if mock.Response.Status != http.StatusOK {
		return &result, &Error{Message: "mocked api error"}
	}

	return mock.Response, mock.Error
}
