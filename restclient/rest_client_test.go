package restclient_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/arielsrv/golang-toolkit/examples/service"
	"github.com/arielsrv/golang-toolkit/restclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"testing"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) Do(*http.Request) (*http.Response, error) {
	args := m.Called()
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestGetOk(t *testing.T) {
	httpClient := new(MockClient)
	httpClient.
		On("Do").
		Return(Ok())

	restClient := restclient.
		RESTClient{HTTPClient: httpClient}

	userResponse, err := restclient.
		Execute[UserResponse]{RESTClient: &restClient}.
		Get("api.internal.iskaypet.com/users", nil)

	assert.NoError(t, err)
	assert.NotNil(t, userResponse)

	assert.NotNil(t, userResponse.Data)
	assert.NotNil(t, userResponse.Headers)
	assert.Equal(t, "abc,def", userResponse.Headers.Get("custom-header"))
	assert.Equal(t, http.StatusOK, userResponse.Status)
	assert.Equal(t, int64(1), userResponse.Data.ID)
	assert.Equal(t, "John Doe", userResponse.Data.Name)
}

func TestPostOk(t *testing.T) {
	httpClient := new(MockClient)
	httpClient.
		On("Do").
		Return(Ok())

	restClient := restclient.
		RESTClient{HTTPClient: httpClient}

	userRequest := service.UserRequest{
		Name: "John Doe",
	}

	userResponse, err := restclient.
		Execute[service.UserRequest]{RESTClient: &restClient}.
		Post("api.internal.iskaypet.com/users", userRequest, nil)

	assert.NoError(t, err)
	assert.NotNil(t, userResponse)

	assert.NotNil(t, userResponse.Data)
	assert.NotNil(t, userResponse.Headers)
	assert.Equal(t, "abc,def", userResponse.Headers.Get("custom-header"))
	assert.Equal(t, http.StatusOK, userResponse.Status)
	assert.Equal(t, int64(1), userResponse.Data.ID)
	assert.Equal(t, "John Doe", userResponse.Data.Name)
}

func TestGetNotFound(t *testing.T) {
	httpClient := new(MockClient)
	httpClient.
		On("Do").
		Return(NotFound())

	restClient := restclient.
		RESTClient{HTTPClient: httpClient}

	userRequest := service.UserRequest{
		Name: "John Doe",
	}

	userResponse, err := restclient.
		Execute[service.UserRequest]{RESTClient: &restClient}.
		Post("api.internal.iskaypet.com/users", userRequest, nil)

	assert.Error(t, err)
	assert.Equal(t, "not found", err.Error())
	var restClientError *restclient.APINotFoundError
	assert.True(t, errors.As(err, &restClientError))
	assert.NotNil(t, userResponse)
	assert.Equal(t, http.StatusNotFound, userResponse.Status)
}

func TestGetSecurityError(t *testing.T) {
	httpClient := new(MockClient)
	httpClient.
		On("Do").
		Return(Unauthorized())

	restClient := restclient.
		RESTClient{HTTPClient: httpClient}

	userRequest := service.UserRequest{
		Name: "John Doe",
	}

	userResponse, err := restclient.
		Execute[service.UserRequest]{RESTClient: &restClient}.
		Post("api.internal.iskaypet.com/users", userRequest, nil)

	assert.Error(t, err)
	assert.Equal(t, "unauthorized", err.Error())
	var restClientError *restclient.APISecurityError
	assert.True(t, errors.As(err, &restClientError))
	assert.NotNil(t, userResponse)
	assert.Equal(t, http.StatusUnauthorized, userResponse.Status)
}

func TestPostNotFound(t *testing.T) {
	httpClient := new(MockClient)
	httpClient.
		On("Do").
		Return(NotFound())

	restClient := restclient.
		RESTClient{HTTPClient: httpClient}

	userResponse, err := restclient.
		Execute[UserResponse]{RESTClient: &restClient}.
		Get("api.internal.iskaypet.com/users", nil)

	assert.Error(t, err)
	assert.Equal(t, "not found", err.Error())
	var restClientError *restclient.APINotFoundError
	assert.True(t, errors.As(err, &restClientError))
	assert.NotNil(t, userResponse)
	assert.Equal(t, http.StatusNotFound, userResponse.Status)
}

func TestNewRestClient(t *testing.T) {
	restPool, err := restclient.NewRESTPoolBuilder().MakeDefault().Build()
	assert.NoError(t, err)

	restClient := restclient.NewRESTClient(*restPool)
	assert.NotNil(t, restClient)
	assert.NotNil(t, restClient.HTTPClient)
}

func TestParsingError(t *testing.T) {
	httpClient := new(MockClient)
	httpClient.
		On("Do").
		Return(Ok())

	restClient := restclient.
		RESTClient{HTTPClient: httpClient}

	_, err := restclient.
		Execute[[]UserResponse]{RESTClient: &restClient}.
		Get("api.internal.iskaypet.com/users", nil)

	assert.Error(t, err)
}

func TestInvalidScheme(t *testing.T) {
	httpClient := new(MockClient)
	httpClient.
		On("Do").
		Return(Error("invalid url"))

	restClient := restclient.
		RESTClient{HTTPClient: httpClient}

	response, err := restclient.
		Execute[UserResponse]{RESTClient: &restClient}.
		Get("mailto://\\n", nil)

	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestInvalidRequest(t *testing.T) {
	httpClient := new(MockClient)
	httpClient.
		On("Do").
		Return(Error("invalid request"))

	restClient := restclient.
		RESTClient{HTTPClient: httpClient}

	response, err := restclient.
		Execute[UserResponse]{RESTClient: &restClient}.
		Get("api.internal.com", nil)

	assert.Error(t, err)
	assert.Equal(t, "invalid request", err.Error())
	assert.Nil(t, response)
}

type UserResponse struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func Ok() (*http.Response, error) {
	userResponse := UserResponse{
		ID:   int64(1),
		Name: "John Doe",
	}
	binary, err := json.Marshal(userResponse)
	if err != nil {
		return nil, err
	}
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(binary)),
		Header: map[string][]string{
			"custom-header": {"abc", "def"},
		},
	}, nil
}

func NotFound() (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       io.NopCloser(bytes.NewBuffer([]byte("not found"))),
	}, nil
}

func Unauthorized() (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusUnauthorized,
		Body:       io.NopCloser(bytes.NewBuffer([]byte("unauthorized"))),
	}, nil
}

func Error(message string) (*http.Response, error) {
	return nil, errors.New(message)
}
