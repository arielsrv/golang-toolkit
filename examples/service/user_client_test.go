package service_test

import (
	"github.com/arielsrv/golang-toolkit/examples/service"
	"github.com/arielsrv/golang-toolkit/restclient"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestOkGet(t *testing.T) {
	restClient := restclient.MockResponse[[]service.UserResponse]{}.
		NewRESTClient().
		AddMockRequest(restclient.MockRequest{
			Method: http.MethodGet,
			URL:    "https://gorest.co.in/public/v2/users",
		}, GetAPICollectionResponse(), restclient.NoNetworkError()).
		Build()

	assert.NotNil(t, restClient)

	userClient := service.NewUserClient(*restClient)

	actual, err := userClient.GetUsers()
	assert.NoError(t, err)
	assert.NotNil(t, actual)
}

func TestOkGetUser(t *testing.T) {
	restClient := restclient.MockResponse[service.UserResponse]{}.
		NewRESTClient().
		AddMockRequest(restclient.MockRequest{
			Method: http.MethodGet,
			URL:    "https://gorest.co.in/public/v2/users/1",
		}, restclient.Response[service.UserResponse]{
			Data:    GetAPICollectionResponse().Data[0],
			Status:  http.StatusOK,
			Headers: nil,
		}, restclient.NoNetworkError()).
		Build()

	assert.NotNil(t, restClient)

	userClient := service.NewUserClient(*restClient)

	actual, err := userClient.GetUser(int64(1))
	assert.NoError(t, err)
	assert.NotNil(t, actual)
}

func TestOkPost(t *testing.T) {
	restClient := restclient.MockResponse[service.UserResponse]{}.
		NewRESTClient().
		AddMockRequest(restclient.MockRequest{
			Method: http.MethodPost,
			URL:    "https://gorest.co.in/public/v2/users",
		}, GetAPIPostResponse(), restclient.NoNetworkError()).
		Build()

	assert.NotNil(t, restClient)

	userClient := service.NewUserClient(*restClient)

	userRequest := &service.UserRequest{
		Name: "John Doe",
	}

	actual, err := userClient.CreateUser(*userRequest)
	assert.NoError(t, err)
	assert.NotNil(t, actual)
}

func TestNetworkError(t *testing.T) {
	restClient := restclient.MockResponse[[]service.UserResponse]{}.
		NewRESTClient().
		AddMockRequest(restclient.MockRequest{
			Method: http.MethodGet,
			URL:    "https://gorest.co.in/public/v2/users",
		}, GetAPICollectionResponse(), restclient.NetworkError()).
		Build()

	assert.NotNil(t, restClient)

	userClient := service.NewUserClient(*restClient)

	actual, err := userClient.GetUsers()
	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestNetworkErrorGetUser(t *testing.T) {
	restClient := restclient.MockResponse[service.UserResponse]{}.
		NewRESTClient().
		AddMockRequest(restclient.MockRequest{
			Method: http.MethodGet,
			URL:    "https://gorest.co.in/public/v2/users/1",
		}, GetAPIResponse(), restclient.NetworkError()).
		Build()

	assert.NotNil(t, restClient)

	userClient := service.NewUserClient(*restClient)

	actual, err := userClient.GetUser(int64(1))
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

func GetAPICollectionResponse() restclient.Response[[]service.UserResponse] {
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

func GetAPIResponse() restclient.Response[service.UserResponse] {
	return restclient.Response[service.UserResponse]{
		Data:   GetAPICollectionResponse().Data[0],
		Status: http.StatusOK,
	}
}

func GetAPIPostResponse() restclient.Response[service.UserResponse] {
	userResponse := service.UserResponse{
		ID:   int64(1),
		Name: "John Doe",
	}

	return restclient.Response[service.UserResponse]{
		Data:   userResponse,
		Status: http.StatusOK,
	}
}

func GetAPIError() restclient.Response[[]service.UserResponse] {
	return restclient.Response[[]service.UserResponse]{
		Data:   nil,
		Status: http.StatusInternalServerError,
	}
}
