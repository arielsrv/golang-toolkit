## ⚡️ Services

```go
package main

import (
    "github.com/arielsrv/golang-toolkit/examples/service"
    "github.com/arielsrv/golang-toolkit/restclient"
    "log"
    "time"
)

func main() {
    restPool, err := restclient.
        NewRESTPoolBuilder().
        WithName("users").
        WithTimeout(time.Millisecond * 1000).
        WithMaxConnectionsPerHost(20).
        WithMaxIdleConnectionsPerHost(20).
        Build()

    if err != nil {
        log.Fatalln(err)
    }

    restClient := restclient.NewRESTClient(*restPool)
    userClient := service.NewUserClient(*restClient)
    userService := service.NewUserService(userClient)

    usersDto, err := userService.GetUsers()

    if err != nil {
        log.Fatal(err)
    }

    for _, userResponse := range usersDto {
        log.Printf("User: ID: %d, FullName: %s", userResponse.ID, userResponse.FullName)
    }
}
```

### UserClient

```go
package service

import (
    "github.com/arielsrv/golang-toolkit/restclient"
)

type IUserClient interface {
    GetUsers() ([]UserResponse, error)
}

type UserClient struct {
    restClient restclient.RESTClient
}

func NewUserClient(restClient restclient.RESTClient) *UserClient {
    return &UserClient{restClient: restClient}
}

func (userClient UserClient) GetUsers() ([]UserResponse, error) {
    response, err := restclient.
        Execute[[]UserResponse]{RESTClient: &userClient.restClient}.
        Get("https://gorest.co.in/public/v2/users")

    if err != nil {
        return nil, err
    }

    return response.Data, nil
}
```

### UserTest

```go
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
```