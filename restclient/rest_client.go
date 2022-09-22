package restclient

import (
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
)

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
}

type Response[T any] struct {
	Data    T
	Status  int
	Headers map[string][]string
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

func (execute Execute[T]) Get(url string) (Response[T], error) {
	var result Response[T]
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return result, err
	}
	response, err := execute.RESTClient.HTTPClient.Do(request)
	if err != nil {
		return result, err
	}

	body, err := io.ReadAll(response.Body)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(response.Body)
	err = json.Unmarshal(body, &result.Data)

	result.Status = response.StatusCode
	result.Headers = response.Header

	return result, nil
}
