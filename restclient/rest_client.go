package restclient

import (
	"context"
	"encoding/json"
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

func (m MockResponse[T]) NewRESTClient(method string, url string, response Response[T], err error) *RESTClient {
	m.responses = make(map[string]Tuple[T])
	m.responses[url] = Tuple[T]{
		Method:   method,
		Response: &response,
		Error:    err,
	}
	return &RESTClient{
		test: true,
		mock: m.responses,
	}
}

func (e Execute[T]) Get(url string) (*Response[T], error) {
	var result Response[T]
	if e.RESTClient.test {
		mocks, boxing := e.RESTClient.mock.(map[string]Tuple[T])
		if !boxing {
			return &result, &MockError{Message: "Internal mocking error. "}
		}
		mock := mocks[url]
		return mock.Response, nil
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
