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

type RESTQuery[TOutput any] interface {
	Get(url string) (*APIResponse[TOutput], error)
}

type RESTCommand[TInput any, TOutput any] interface {
	Post(url string, request TInput) (*APIResponse[TOutput], error)
}

type IClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Read[TOutput any] struct {
	RESTClient *RESTClient
}

type Write[TInput any, TOutput any] struct {
	RESTClient *RESTClient
}

type RESTClient struct {
	HTTPClient  IClient
	restPool    RESTPool
	testingMode bool
	Mock        any
}

type APIResponse[TOutput any] struct {
	Data    TOutput
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

func (e Read[TOutput]) Get(url string, headers Headers) (*APIResponse[TOutput], error) {
	var result *APIResponse[TOutput]
	if e.RESTClient.testingMode {
		return e.GetMock(http.MethodGet, url, result)
	}
	result, err := execute[TOutput](e.RESTClient.HTTPClient, http.MethodGet, url, nil, headers)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (e Write[TInput, TOutput]) Post(url string, request TInput, headers Headers) (*APIResponse[TOutput], error) {
	var result *APIResponse[TOutput]
	if e.RESTClient.testingMode {
		return e.GetMock(http.MethodPost, url, result)
	}
	binary, err := json.Marshal(request)
	if err != nil {
		return result, err
	}
	reader := bytes.NewReader(binary)
	result, err = execute[TOutput](e.RESTClient.HTTPClient, http.MethodPost, url, reader, headers)
	if err != nil {
		return result, err
	}

	return result, nil
}

func execute[TOutput any](client IClient, method string, url string, data io.Reader, headers Headers) (*APIResponse[TOutput], error) {
	request, err := http.NewRequestWithContext(context.Background(), method, url, data)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse APIResponse[TOutput]
	apiResponse.Status = response.StatusCode
	apiResponse.Headers = make(map[string]string)
	for key, values := range response.Header {
		value := strings.Join(values, ",")
		apiResponse.Headers.Put(key, value)
	}

	if !IsOk(response.StatusCode) {
		err = handleError(response, body)
		if err != nil {
			return &apiResponse, err
		}
	}

	err = json.Unmarshal(body, &apiResponse.Data)
	if err != nil {
		return &apiResponse, err
	}

	return &apiResponse, nil
}

func handleError(response *http.Response, body []byte) error {
	switch response.StatusCode {
	case http.StatusNotFound:
		return &APINotFoundError{
			*getAPIError(response.StatusCode, body),
		}
	case http.StatusBadRequest:
		return &APIBadRequestError{
			*getAPIError(response.StatusCode, body),
		}
	case http.StatusUnauthorized, http.StatusForbidden:
		return &APISecurityError{
			*getAPIError(response.StatusCode, body),
		}
	}
	return getAPIError(response.StatusCode, body)
}

func getAPIError(statusCode int, body []byte) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Message:    string(body),
	}
}
