package service_test

import (
	"github.com/arielsrv/golang-toolkit/examples/service"
	"github.com/arielsrv/golang-toolkit/restclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockUserClient[T any] struct {
	mock.Mock
}

func (m *MockUserClient[T]) GetUsers() ([]service.UserResponse, error) {
	args := m.Called()
	return args.Get(0).([]service.UserResponse), args.Error(1)
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
	return nil, &restclient.Error{Message: "internal server error"}
}
