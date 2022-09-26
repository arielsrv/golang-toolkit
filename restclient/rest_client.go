package restclient

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"strings"
)

type REST[T any] interface {
	Get(url string) (*Response[T], error)
	Post(url string, request T) (*Response[T], error)
}

type IClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Execute[T any] struct {
	RESTClient *RESTClient
}

type RESTClient struct {
	HTTPClient  IClient
	restPool    RESTPool
	testingMode bool
	Mock        any
}

type Response[T any] struct {
	Data    T
	Status  int
	Headers Headers
}

type Headers map[string]string

func (headers Headers) Get(key string) string {
	return headers[key]
}

func (headers Headers) Put(key string, value string) {
	headers[key] = value
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

func (e Execute[T]) Get(url string, headers Headers) (*Response[T], error) {
	var result *Response[T]
	if e.RESTClient.testingMode {
		return e.GetMock(http.MethodGet, url, result)
	}
	result, err := e.call(http.MethodGet, url, nil, headers)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (e Execute[T]) Post(url string, request T, headers Headers) (*Response[T], error) {
	var result *Response[T]
	if e.RESTClient.testingMode {
		return e.GetMock(http.MethodPost, url, result)
	}
	binary, err := json.Marshal(request)
	if err != nil {
		return result, err
	}
	reader := bytes.NewReader(binary)
	result, err = e.call(http.MethodPost, url, reader, headers)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (e Execute[T]) call(method string, url string, data io.Reader, headers Headers) (*Response[T], error) {
	var result Response[T]
	request, err := http.NewRequestWithContext(context.Background(), method, url, data)
	if err != nil {
		return nil, err
	}
	
	for key, value := range headers {
		request.Header.Set(key, value)
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

	if response.StatusCode > http.StatusBadRequest {
		err = e.handleError(response, body)
		if err != nil {
			return &result, err
		}
	}

	err = json.Unmarshal(body, &result.Data)
	if err != nil {
		return &result, err
	}

	return &result, nil
}

func (e Execute[T]) handleError(response *http.Response, body []byte) error {
	switch response.StatusCode {
	case http.StatusNotFound:
		return &APINotFoundError{
			StatusCode: response.StatusCode,
			Message:    string(body),
		}
	case http.StatusBadRequest:
		return &APIBadRequestError{
			StatusCode: response.StatusCode,
			Message:    string(body),
		}
	case http.StatusUnauthorized:
	case http.StatusForbidden:
		return &APISecurityError{
			StatusCode: response.StatusCode,
			Message:    string(body),
		}
	}
	return &APIError{
		StatusCode: response.StatusCode,
		Message:    string(body),
	}
}
