package restclient_test

import (
	"bytes"
	"encoding/json"
	"github.com/personal/golang-toolkit/restclient"
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
	return args.Get(0).(*http.Response), nil
}

func Test(t *testing.T) {
	httpClient := new(MockClient)
	httpClient.
		On("Do").
		Return(Get())

	restClient := restclient.
		RESTClient{HTTPClient: httpClient}

	userResponse, err := restclient.
		Execute[UserResponse]{RESTClient: &restClient}.
		Get("api.internal.iskaypet.com/users")

	assert.NoError(t, err)
	assert.NotNil(t, userResponse)

	assert.NotNil(t, userResponse.Data)
	assert.Equal(t, http.StatusOK, userResponse.Status)
	assert.Equal(t, int64(1), userResponse.Data.ID)
	assert.Equal(t, "John Doe", userResponse.Data.Name)
}

type UserResponse struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func Get() (*http.Response, error) {
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
	}, nil
}
