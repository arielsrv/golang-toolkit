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
	panic("implement me")
}

type IUserService interface {
	GetUsers() ([]UserDto, error)
}

type UserDto struct {
	ID       int64  `json:"id,omitempty"`
	FullName string `json:"name,omitempty"`
}

type UserResponse struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type UserService struct {
	userClient IUserClient
}

func NewUserService(userClient IUserClient) *UserService {
	return &UserService{userClient: userClient}
}

func (service UserService) GetUsers() ([]UserDto, error) {
	response, err := service.userClient.GetUsers()

	if err != nil {
		return nil, err
	}

	var result []UserDto
	for _, userResponse := range response {
		userDto := UserDto{
			ID:       userResponse.ID,
			FullName: userResponse.Name,
		}
		result = append(result, userDto)
	}

	return result, nil
}
