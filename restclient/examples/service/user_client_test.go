package service_test

import (
	"github.com/arielsrv/golang-toolkit/restclient/core"
	"github.com/arielsrv/golang-toolkit/restclient/examples/service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestOkGet(t *testing.T) {
	restClient := core.MockResponse[[]service.UserResponse]{}.
		NewRESTClient().
		AddMockRequest(core.MockRequest{
			Method: http.MethodGet,
			URL:    "https://gorest.co.in/public/v2/users",
		}, GetAPICollectionResponse(), core.NoNetworkError()).
		Build()

	assert.NotNil(t, restClient)

	userClient := service.NewUserClient(*restClient)

	actual, err := userClient.GetUsers()
	assert.NoError(t, err)
	assert.NotNil(t, actual)
}

func TestOkGetUser(t *testing.T) {
	restClient := core.MockResponse[service.UserResponse]{}.
		NewRESTClient().
		AddMockRequest(core.MockRequest{
			Method: http.MethodGet,
			URL:    "https://gorest.co.in/public/v2/users/1",
		}, core.APIResponse[service.UserResponse]{
			Data:    GetAPICollectionResponse().Data[0],
			Status:  http.StatusOK,
			Headers: nil,
		}, core.NoNetworkError()).
		Build()

	assert.NotNil(t, restClient)

	userClient := service.NewUserClient(*restClient)

	actual, err := userClient.GetUser(int64(1))
	assert.NoError(t, err)
	assert.NotNil(t, actual)
}

func TestOkPost(t *testing.T) {
	restClient := core.MockResponse[service.UserResponse]{}.
		NewRESTClient().
		AddMockRequest(core.MockRequest{
			Method: http.MethodPost,
			URL:    "https://gorest.co.in/public/v2/users",
		}, GetAPIPostResponse(), core.NoNetworkError()).
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
	restClient := core.MockResponse[[]service.UserResponse]{}.
		NewRESTClient().
		AddMockRequest(core.MockRequest{
			Method: http.MethodGet,
			URL:    "https://gorest.co.in/public/v2/users",
		}, GetAPICollectionResponse(), core.NetworkError()).
		Build()

	assert.NotNil(t, restClient)

	userClient := service.NewUserClient(*restClient)

	actual, err := userClient.GetUsers()
	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestNetworkErrorGetUser(t *testing.T) {
	restClient := core.MockResponse[service.UserResponse]{}.
		NewRESTClient().
		AddMockRequest(core.MockRequest{
			Method: http.MethodGet,
			URL:    "https://gorest.co.in/public/v2/users/1",
		}, GetAPIResponse(), core.NetworkError()).
		Build()

	assert.NotNil(t, restClient)

	userClient := service.NewUserClient(*restClient)

	actual, err := userClient.GetUser(int64(1))
	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestApiError(t *testing.T) {
	restClient := core.MockResponse[[]service.UserResponse]{}.
		NewRESTClient().
		AddMockRequest(core.MockRequest{
			Method: http.MethodGet,
			URL:    "https://gorest.co.in/public/v2/users",
		}, GetAPIError(), core.NoNetworkError()).
		Build()

	assert.NotNil(t, restClient)

	userClient := service.NewUserClient(*restClient)

	actual, err := userClient.GetUsers()
	assert.Error(t, err)
	assert.Nil(t, actual)
}

func GetAPICollectionResponse() core.APIResponse[[]service.UserResponse] {
	userResponse := service.UserResponse{
		ID:   int64(1),
		Name: "John Doe",
	}
	var result []service.UserResponse
	result = append(result, userResponse)

	return core.APIResponse[[]service.UserResponse]{
		Data:   result,
		Status: http.StatusOK,
	}
}

func GetAPIResponse() core.APIResponse[service.UserResponse] {
	return core.APIResponse[service.UserResponse]{
		Data:   GetAPICollectionResponse().Data[0],
		Status: http.StatusOK,
	}
}

func GetAPIPostResponse() core.APIResponse[service.UserResponse] {
	userResponse := service.UserResponse{
		ID:   int64(1),
		Name: "John Doe",
	}

	return core.APIResponse[service.UserResponse]{
		Data:   userResponse,
		Status: http.StatusOK,
	}
}

func GetAPIError() core.APIResponse[[]service.UserResponse] {
	return core.APIResponse[[]service.UserResponse]{
		Data:   nil,
		Status: http.StatusInternalServerError,
	}
}
