## ⚡️ Services

```go
package main

import (
	"fmt"
	"github.com/arielsrv/golang-toolkit/examples/service"
	"github.com/arielsrv/golang-toolkit/restclient"
	"github.com/tjarratt/babble"
	"log"
	"strings"
	"time"
)

func main() {
	restPool, err := restclient.
		NewRESTPoolBuilder().
		WithName("users").
		WithTimeout(time.Millisecond * 5000).
		WithMaxConnectionsPerHost(20).
		WithMaxIdleConnectionsPerHost(20).
		Build()

	if err != nil {
		log.Fatalln(err)
	}

	restClient := restclient.NewRESTClient(*restPool)
	userClient := service.NewUserClient(*restClient)
	userService := service.NewUserService(userClient)

	name := strings.ToLower(babble.NewBabbler().Babble())
	userDto, err := userService.CreateUser(service.UserDto{
		FullName: name,
		Email:    fmt.Sprintf("%s@github.com", name),
		Gender:   "male",
		Status:   "active",
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("User: ID: %d, Name: %s, Email: %s",
		userDto.ID,
		userDto.FullName,
		userDto.Email)

	search, err := userService.GetUser(userDto.ID)

	if err != nil {
		log.Fatalf("Error User %d, %s", userDto.ID, err)
	}

	log.Printf("User: ID: %d, Name: %s, Email: %s",
		search.ID,
		search.FullName,
		search.Email)
}
```

### ⚡️ Unit Test
```go
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
		}, GetAPIResponse(), restclient.NoNetworkError()).
		Build()

	assert.NotNil(t, restClient)

	userClient := service.NewUserClient(*restClient)

	actual, err := userClient.GetUsers()
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
```
