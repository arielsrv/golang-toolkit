package service_test

import (
	"github.com/arielsrv/golang-toolkit/examples/service"
	"github.com/arielsrv/golang-toolkit/restclient"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestOk(t *testing.T) {
	restClient := restclient.MockResponse[[]service.UserResponse]{}.
		NewRESTClient().
		AddMockRequest(restclient.MockRequest{
			Method: http.MethodGet,
			URL:    "https://gorest.co.in/public/v2/users",
		}, GetAPIResponse(), restclient.NoNetworkError()).
		Build()

	assert.NotNil(t, restClient)

	userClient := service.NewUserClient(*restClient)

	actual, err := userClient.GetUsers()
	assert.NoError(t, err)
	assert.NotNil(t, actual)
}

func TestNetworkError(t *testing.T) {
	restClient := restclient.MockResponse[[]service.UserResponse]{}.
		NewRESTClient().
		AddMockRequest(restclient.MockRequest{
			Method: http.MethodGet,
			URL:    "https://gorest.co.in/public/v2/users",
		}, GetAPIResponse(), restclient.NetworkError()).
		Build()

	assert.NotNil(t, restClient)

	userClient := service.NewUserClient(*restClient)

	actual, err := userClient.GetUsers()
	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestApiError(t *testing.T) {
	restClient := restclient.MockResponse[[]service.UserResponse]{}.
		NewRESTClient().
		AddMockRequest(restclient.MockRequest{
			Method: http.MethodGet,
			URL:    "https://gorest.co.in/public/v2/users",
		}, GetAPIError(), restclient.NoNetworkError()).
		Build()

	assert.NotNil(t, restClient)

	userClient := service.NewUserClient(*restClient)

	actual, err := userClient.GetUsers()
	assert.Error(t, err)
	assert.Nil(t, actual)
}

func GetAPIResponse() restclient.Response[[]service.UserResponse] {
	userResponse := service.UserResponse{
		ID:   int64(1),
		Name: "John Doe",
	}
	var result []service.UserResponse
	result = append(result, userResponse)

	return restclient.Response[[]service.UserResponse]{
		Data:   result,
		Status: http.StatusOK,
	}
}

func GetAPIError() restclient.Response[[]service.UserResponse] {
	return restclient.Response[[]service.UserResponse]{
		Data:   nil,
		Status: http.StatusInternalServerError,
	}
}
