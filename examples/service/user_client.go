package service

import (
	"fmt"
	"github.com/arielsrv/golang-toolkit/restclient"
	httpheader "github.com/go-http-utils/headers"
	"github.com/ldez/mimetype"
	"os"
)

type IUserClient interface {
	GetUsers() ([]UserResponse, error)
	CreateUser(userRequest UserRequest) (int64, error)
	GetUser(id int64) (*UserResponse, error)
}

type UserClient struct {
	baseURL    string
	restClient restclient.RESTClient
}

func NewUserClient(restClient restclient.RESTClient) *UserClient {
	return &UserClient{
		baseURL:    "https://gorest.co.in/public/v2",
		restClient: restClient,
	}
}

func (userClient UserClient) GetUsers() ([]UserResponse, error) {
	apiURL := fmt.Sprintf("%s/users", userClient.baseURL)
	response, err := restclient.
		Read[[]UserResponse]{RESTClient: &userClient.restClient}.
		Get(apiURL, nil)

	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (userClient UserClient) GetUser(userID int64) (*UserResponse, error) {
	apiURL := fmt.Sprintf("%s/users/%d", userClient.baseURL, userID)
	response, err := restclient.
		Read[UserResponse]{RESTClient: &userClient.restClient}.
		Get(apiURL, nil)

	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}

func (userClient UserClient) CreateUser(userRequest UserRequest) (int64, error) {
	headers := restclient.Headers{}
	headers.Put(httpheader.Authorization, fmt.Sprintf("Bearer %s", os.Getenv("GOREST_TOKEN")))
	headers.Put(httpheader.ContentType, mimetype.ApplicationJSON)
	apiURL := fmt.Sprintf("%s/users", userClient.baseURL)
	response, err := restclient.
		Write[UserRequest, UserResponse]{RESTClient: &userClient.restClient}.
		Post(apiURL, userRequest, headers)

	if err != nil {
		return 0, err
	}

	return response.Data.ID, nil
}
