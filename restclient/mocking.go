package restclient

import (
	"errors"
	"github.com/arielsrv/golang-toolkit/common/equality"
)

type MockError struct {
	Message string
}

func (e *MockError) Error() string {
	return e.Message
}

type Tuple[TOutput any] struct {
	Method   string
	Response *APIResponse[TOutput]
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
	hash = uint64(31)*hash + equality.GetValue(mockRequest.Method)
	hash = uint64(31)*hash + equality.GetValue(mockRequest.URL)
	return hash
}

func (mockResponse MockResponse[TOutput]) AddMockRequest(mockRequest MockRequest, response APIResponse[TOutput], err error) *MockResponse[TOutput] {
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
	hashCode := mockedRequest.GetHashCode()
	mock := mocks[hashCode]
	if !IsOk(mock.Response.Status) {
		return nil, &APIError{Message: "mocked api error"}
	}

	return &mock, nil
}

func (e Read[TOutput]) GetMock(method string, url string, result *APIResponse[TOutput]) (*APIResponse[TOutput], error) {
	mock, err := get[TOutput](e.RESTClient.Mock, method, url)
	if err != nil {
		return result, err
	}
	return mock.Response, mock.Error
}

func (e Write[TInput, TOutput]) GetMock(method string, url string, result *APIResponse[TOutput]) (*APIResponse[TOutput], error) {
	mock, err := get[TOutput](e.RESTClient.Mock, method, url)
	if err != nil {
		return result, err
	}
	return mock.Response, mock.Error
}
