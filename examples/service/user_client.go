package service

import "github.com/arielsrv/golang-toolkit/restclient"

type IUserClient interface {
	GetUsers() ([]UserResponse, error)
}

type UserClient struct {
	restClient restclient.RESTClient
}

func NewUserClient(restClient restclient.RESTClient) *UserClient {
	return &UserClient{restClient: restClient}
}
