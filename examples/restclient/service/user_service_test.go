package service_test

import (
	"errors"
	"github.com/arielsrv/golang-toolkit/examples/restclient/service"
	"github.com/arielsrv/golang-toolkit/restclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockUserClient[TOutput any] struct {
	mock.Mock
}

func (m *MockUserClient[TOutput]) GetUsers() ([]service.UserResponse, error) {
	args := m.Called()
	return args.Get(0).([]service.UserResponse), args.Error(1)
}

func (m *MockUserClient[TOutput]) GetUser(int64) (*service.UserResponse, error) {
	args := m.Called()
	return args.Get(0).(*service.UserResponse), args.Error(1)
}

func (m *MockUserClient[TOutput]) CreateUser(service.UserRequest) (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func Test(t *testing.T) {
	userClient := new(MockUserClient[[]service.UserResponse])
	userClient.On("GetUsers").Return(Ok())
	userService := service.NewUserService(userClient)

	actual, err := userService.GetUsers()

	assert.NoError(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, 1, len(actual))
	assert.Equal(t, int64(1), actual[0].ID)
	assert.Equal(t, "John Doe", actual[0].FullName)
}

func TestGetUser(t *testing.T) {
	userClient := new(MockUserClient[service.UserResponse])
	userClient.On("GetUser").Return(GetUserResponse())
	userService := service.NewUserService(userClient)

	actual, err := userService.GetUser(int64(1))

	assert.NoError(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, int64(1), actual.ID)
	assert.Equal(t, "John Doe", actual.FullName)
}

func TestGetError(t *testing.T) {
	userClient := new(MockUserClient[service.UserResponse])
	userClient.On("GetUser").Return(GetError())
	userService := service.NewUserService(userClient)

	actual, err := userService.GetUser(int64(1))

	assert.Error(t, err)
	assert.Equal(t, "client error", err.Error())
	assert.Nil(t, actual)
}

func GetError() (*service.UserResponse, error) {
	return nil, errors.New("client error")
}

func GetUserResponse() (*service.UserResponse, error) {
	return &service.UserResponse{
		ID:   int64(1),
		Name: "John Doe",
	}, nil
}

func TestCreate(t *testing.T) {
	userClient := new(MockUserClient[service.UserResponse])
	userDto := &service.UserDto{FullName: "John Doe"}
	userClient.On("CreateUser").Return(GetUser())
	userService := service.NewUserService(userClient)

	actual, err := userService.CreateUser(*userDto)

	assert.NoError(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, int64(1), actual.ID)
	assert.Equal(t, "John Doe", userDto.FullName)
}

func GetUser() (int64, error) {
	return int64(1), nil
}

func TestError(t *testing.T) {
	userClient := new(MockUserClient[[]service.UserResponse])
	userClient.On("GetUsers").Return(Error())
	userService := service.NewUserService(userClient)

	actual, err := userService.GetUsers()

	assert.Error(t, err)
	assert.Nil(t, actual)
}

func Ok() ([]service.UserResponse, error) {
	userResponse := service.UserResponse{
		ID:   int64(1),
		Name: "John Doe",
	}
	var result []service.UserResponse
	result = append(result, userResponse)
	return result, nil
}

func Error() ([]service.UserResponse, error) {
	return nil, &restclient.APIError{Message: "api server error"}
}
