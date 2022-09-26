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

type Tuple[TOutput any] struct {
	Method   string
	Response *Response[TOutput]
	Error    error
}

func NoNetworkError() error {
	return nil
}

func NetworkError() error {
	return errors.New("network error")
}

type MockResponse[TOutput any] struct {
	Responses map[uint64]Tuple[TOutput]
}

func (mockResponse MockResponse[TOutput]) NewRESTClient() *MockResponse[TOutput] {
	mockResponse.Responses = make(map[uint64]Tuple[TOutput])
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

func (mockResponse MockResponse[TOutput]) AddMockRequest(mockRequest MockRequest, response Response[TOutput], err error) *MockResponse[TOutput] {
	hash := mockRequest.GetHashCode()
	mockResponse.Responses[hash] = Tuple[TOutput]{
		Method:   mockRequest.Method,
		Response: &response,
		Error:    err,
	}
	return &mockResponse
}

func (mockResponse MockResponse[TOutput]) Build() *RESTClient {
	return &RESTClient{
		testingMode: true,
		Mock:        mockResponse.Responses,
	}
}

func get[TOutput any](reference any, method string, url string) (*Tuple[TOutput], error) {
	mocks, boxing := reference.(map[uint64]Tuple[TOutput])
	if !boxing {
		return nil, &MockError{Message: "Internal mocking error. "}
	}
	mockedRequest := MockRequest{
		Method: method,
		URL:    url,
	}
	mock := mocks[mockedRequest.GetHashCode()]
	if mock.Response.Status != http.StatusOK {
		return nil, &APIError{Message: "mocked api error"}
	}

	return &mock, nil
}

func (e Read[TOutput]) GetMock(method string, url string, result *Response[TOutput]) (*Response[TOutput], error) {
	mock, err := get[TOutput](e.RESTClient.Mock, method, url)
	if err != nil {
		return result, &MockError{Message: "Internal mocking error. "}
	}
	return mock.Response, mock.Error
}

func (e Write[TInput, TOutput]) GetMock(method string, url string, result *Response[TOutput]) (*Response[TOutput], error) {
	mock, err := get[TOutput](e.RESTClient.Mock, method, url)
	if err != nil {
		return result, &MockError{Message: "Internal mocking error. "}
	}
	return mock.Response, mock.Error
}
