package restclient

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"strings"
)

type Error struct {
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

type MockError struct {
	Message string
}

func (e *MockError) Error() string {
	return e.Message
}

type REST[T any] interface {
	Get(url string) (Response[T], error)
}

type IClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Execute[T any] struct {
	RESTClient *RESTClient
}

type RESTClient struct {
	HTTPClient IClient
	restPool   RESTPool
	test       bool
	mock       any
}

type Response[T any] struct {
	Data    T
	Status  int
	Headers Headers
}

type Headers map[string]string

func (h Headers) Get(key string) string {
	return h[key]
}

func (h Headers) Put(key string, value string) {
	h[key] = value
}

func NewRESTClient(restPool RESTPool) *RESTClient {
	httpClient := http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   restPool.SocketTimeout,
				KeepAlive: restPool.SocketKeepAlive,
			}).DialContext,
			MaxIdleConns:        restPool.MaxIdleConnections,
			MaxIdleConnsPerHost: restPool.MaxIdleConnectionsPerHost,
			MaxConnsPerHost:     restPool.MaxConnectionsPerHost,
			IdleConnTimeout:     restPool.IdleConnectionTimeout,
			TLSHandshakeTimeout: restPool.TLSHandshakeTimeout,
		},
		Timeout: restPool.Timeout,
	}
	return &RESTClient{
		HTTPClient: &httpClient,
		restPool:   restPool,
	}
}

type Tuple[T any] struct {
	Method   string
	Response *Response[T]
	Error    error
}

type MockResponse[T any] struct {
	responses map[string]Tuple[T]
}

func NoNetworkError() error {
	return nil
}

func NetworkError() error {
	return errors.New("network error")
}

func (m MockResponse[T]) NewRESTClient() *MockResponse[T] {
	m.responses = make(map[string]Tuple[T])
	return &m
}

type MockRequest struct {
	Method string
	URL    string
}

func (m MockResponse[T]) Add(mockedRequest MockRequest, response Response[T], err error) *MockResponse[T] {
	hash := GetHash(mockedRequest)
	m.responses[hash] = Tuple[T]{
		Method:   mockedRequest.Method,
		Response: &response,
		Error:    err,
	}
	return &m
}

func GetHash(mockedRequest MockRequest) string {
	return mockedRequest.Method + mockedRequest.URL
}

func (m MockResponse[T]) Build() *RESTClient {
	return &RESTClient{
		test: true,
		mock: m.responses,
	}
}

func (e Execute[T]) Get(url string) (*Response[T], error) {
	var result Response[T]
	if e.RESTClient.test {
		return e.GetMock(http.MethodGet, url, result)
	}
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	response, err := e.RESTClient.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(response.Body)

	result.Status = response.StatusCode
	result.Headers = make(map[string]string)
	for key, values := range response.Header {
		value := strings.Join(values, ",")
		result.Headers.Put(key, value)
	}

	if response.StatusCode != http.StatusOK {
		return &result, &Error{Message: string(body)}
	}

	err = json.Unmarshal(body, &result.Data)
	if err != nil {
		return &result, err
	}

	return &result, nil
}

func (e Execute[T]) GetMock(method string, url string, result Response[T]) (*Response[T], error) {
	mocks, boxing := e.RESTClient.mock.(map[string]Tuple[T])
	if !boxing {
		return &result, &MockError{Message: "Internal mocking error. "}
	}
	hash := GetHash(MockRequest{
		Method: method,
		URL:    url,
	})
	mock := mocks[hash]
	return mock.Response, mock.Error
}
