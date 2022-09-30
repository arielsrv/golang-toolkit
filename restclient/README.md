# restclient

# Generic REST Client with connection pool
## ⚡️ Quickstart

```go
package main

import (
	"errors"
	"github.com/arielsrv/golang-toolkit/examples/service"
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
	usersResponse, err := restclient.
		Read[[]UserResponse]{RESTClient: restClient}.
		Get("https://gorest.co.in/public/v2/users", nil)

	if err != nil {
		var restClientError *restclient.APIError
		switch {
		case errors.As(err, &restClientError):
			log.Println(err.Error())
			log.Println(usersResponse.Status)
		default:
			log.Printf("unexpected error: %s\n", err)
		}
	}

	for _, userResponse := range usersResponse.Data {
		log.Printf("User: ID: %d, Name: %s", userResponse.ID, userResponse.Name)
	}

	// Generic post
	userRequest := &service.UserRequest{
		Name: "John Doe",
	}

	result, err := restclient.
		Write[service.UserRequest, UserResponse]{RESTClient: restClient}.
		Post("https://gorest.co.in/public/v2/users", *userRequest, nil)

	if err != nil {
		var restClientError *restclient.APIBadRequestError
		switch {
		case errors.As(err, &restClientError):
			log.Println(err.Error())
			log.Println(result.Status)
		default:
			log.Printf("unexpected error: %s\n", err)
		}
	}

	log.Printf("User: ID: %d, Name: %s", result.Data.ID, result.Data.Name)
}
```
