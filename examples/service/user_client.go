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

func (u UserClient) GetUsers() ([]UserResponse, error) {
	response, err := restclient.
		Execute[[]UserResponse]{RESTClient: &u.restClient}.
		Get("https://gorest.co.in/public/v2/users")

	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
