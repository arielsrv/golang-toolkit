# golang-toolkit
[![CI](https://github.com/tj-actions/coverage-badge-go/workflows/CI/badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3ACI)
![Coverage](https://img.shields.io/badge/Coverage-94.4%25-brightgreen)
[![Update release version.](https://github.com/tj-actions/coverage-badge-go/workflows/Update%20release%20version./badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3A%22Update+release+version.%22)

# Generic REST Client with connection pool
## ⚡️ Quickstart

```go
package main

import (
    "errors"
    "github.com/arielsrv/golang-toolkit/restclient"
    "log"
    "time"
)

type UserResponse struct {
    ID   int64  `json:"id,omitempty"`
    Name string `json:"name,omitempty"`
}

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
    
    restClient := restclient.
        NewRESTClient(*restPool)

    // Generic get
    response, err := restclient.
        Read[[]UserResponse]{RESTClient: restClient}.
        Get("https://gorest.co.in/public/v2/users2")

    if err != nil {
        var restClientError *restclient.APIError
        switch {
        case errors.As(err, &restClientError):
            log.Println(err.Error())
            log.Println(response.Status)
        default:
            log.Printf("unexpected error: %s\n", err)
        }
    }

    for _, userResponse := range response.Data {
        log.Printf("User: ID: %d, Name: %s", userResponse.ID, userResponse.Name)
    }

    // Generic post
    userRequest := &UserResponse{
        Name: "John Doe"
    }

    response, err := restclient.
        Write[UserResponse, UserResponse]{RESTClient: restClient}.
        Post("https://gorest.co.in/public/v2/users2", userRequest)

    if err != nil {
        var restClientError *restclient.APIBadRequestError
        switch {
        case errors.As(err, &restClientError):
            log.Println(err.Error())
            log.Println(response.Status)
        default:
            log.Printf("unexpected error: %s\n", err)
        }
    }

    log.Printf("User: ID: %d, Name: %s", userRequest.ID, userRequest.Name)
}
```