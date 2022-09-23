# golang-toolkit
[![CI](https://github.com/tj-actions/coverage-badge-go/workflows/CI/badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3ACI)
![Coverage](https://img.shields.io/badge/Coverage-77.1%25-brightgreen)
[![Update release version.](https://github.com/tj-actions/coverage-badge-go/workflows/Update%20release%20version./badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3A%22Update+release+version.%22)

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

    // Generics
    userResponse, err := restclient.
        Execute[[]UserResponse]{RESTClient: restClient}.
        Get("https://gorest.co.in/public/v2/users")

    if err != nil {
        var restClientError *restclient.Error
        switch {
        case errors.As(err, &restClientError):
            log.Println(userResponse.Status)
            log.Print(err.Error())
        default:
            log.Printf("unexpected error: %s\n", err)
        }
    }

    for _, element := range userResponse.Data {
        log.Println("User")
        log.Printf("\tID: %d", element.ID)
        log.Printf("\tName: %s", element.Name)
    }
}
```