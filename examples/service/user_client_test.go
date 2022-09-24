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
		Add(
			http.MethodGet,
			"https://gorest.co.in/public/v2/users",
			GetUserResponse(),
			restclient.NoNetworkError()).
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
		Add(http.MethodGet, "https://gorest.co.in/public/v2/users", GetUserResponse(), restclient.NetworkError()).
		Build()

	assert.NotNil(t, restClient)

	userClient := service.NewUserClient(*restClient)

	actual, err := userClient.GetUsers()
	assert.Error(t, err)
	assert.Nil(t, actual)
}

func GetUserResponse() restclient.Response[[]service.UserResponse] {
	userResponse := service.UserResponse{
		ID:   int64(1),
		Name: "John Doe",
	}
	var result []service.UserResponse
	result = append(result, userResponse)

	var response restclient.Response[[]service.UserResponse]
	response.Data = result
	response.Status = http.StatusOK
	return response
}
