package service

import (
	"fmt"
	"github.com/arielsrv/golang-toolkit/restclient"
	httpheader "github.com/go-http-utils/headers"
	"github.com/ldez/mimetype"
	"log"
	"net/http"
	"os"
)

type IUserClient interface {
	GetUsers() ([]UserResponse, error)
	CreateUser(userRequest UserRequest) error
}

type UserClient struct {
	restClient restclient.RESTClient
}

func NewUserClient(restClient restclient.RESTClient) *UserClient {
	return &UserClient{restClient: restClient}
}

func (userClient UserClient) GetUsers() ([]UserResponse, error) {
	response, err := restclient.
		Read[[]UserResponse]{RESTClient: &userClient.restClient}.
		Get("https://gorest.co.in/public/v2/users", nil)

	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (userClient UserClient) CreateUser(userRequest UserRequest) error {
	headers := restclient.Headers{}
	headers.Put(httpheader.Authorization, fmt.Sprintf("Bearer %s", os.Getenv("GOREST_TOKEN")))
	headers.Put(httpheader.ContentType, mimetype.ApplicationJSON)

	response, err := restclient.
		Write[UserRequest, []UserRequest]{RESTClient: &userClient.restClient}.
		Post("https://gorest.co.in/public/v2/users", userRequest, headers)

	if err != nil {
		return err
	}

	if response.Status != http.StatusOK {
		log.Print(response.Status)
		log.Print("error")
	}

	return nil
}
