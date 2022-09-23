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
	response, err := restclient.
		Execute[[]UserResponse]{RESTClient: restClient}.
		Get("https://gorest.co.in/public/v2/users2")

	if err != nil {
		var restClientError *restclient.Error
		switch {
		case errors.As(err, &restClientError):
			log.Println(err.Error())
			log.Println(response.Status)
		default:
			log.Printf("unexpected error: %s\n", err)
		}
	}

	for _, element := range response.Data {
		log.Printf("User: ID: %d, Name: %s", element.ID, element.Name)
	}
}
