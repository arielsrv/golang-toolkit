package restclient_test

import (
	"github.com/arielsrv/golang-toolkit/restclient"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestMockRequest_GetHashCode(t *testing.T) {
	request := restclient.MockRequest{
		Method: http.MethodGet,
		URL:    "https://www.google.com",
	}

	actual := request.GetHashCode()
	assert.NotEmpty(t, actual)
	assert.Equal(t, uint64(0xb27f66347ca6f9cd), actual)
}

func TestMockResponse_NewRESTClient(t *testing.T) {
	mockResponse := restclient.MockResponse[[]UserResponse]{}.
		NewRESTClient()

	assert.NotNil(t, mockResponse)
	assert.NotNil(t, mockResponse.Responses)
}

func TestMockResponse_AddMockRequest(t *testing.T) {
	mockRequest := restclient.MockRequest{
		Method: http.MethodGet,
		URL:    "https://gorest.co.in/public/v2/users",
	}
	restClient := restclient.MockResponse[[]UserResponse]{}.
		NewRESTClient().
		AddMockRequest(mockRequest, GetUsersResponse(), restclient.NoNetworkError()).
		Build()

	assert.NotNil(t, restClient)
	assert.Nil(t, restClient.HTTPClient)
	assert.NotNil(t, restClient.Mock)
	actual := restClient.Mock.(map[uint64]restclient.Tuple[[]UserResponse]) //nolint:nolintlint,errcheck
	assert.NotNil(t, actual)
	assert.NotNil(t, actual[mockRequest.GetHashCode()])
	assert.NotNil(t, actual[mockRequest.GetHashCode()].Response)
	assert.Equal(t, 1, len(actual[mockRequest.GetHashCode()].Response.Data))
	assert.Equal(t, int64(1), actual[mockRequest.GetHashCode()].Response.Data[0].ID)
	assert.Equal(t, "John Doe", actual[mockRequest.GetHashCode()].Response.Data[0].Name)
}

func TestExecute_GetMock(t *testing.T) {
	mockRequest := restclient.MockRequest{
		Method: http.MethodGet,
		URL:    "https://gorest.co.in/public/v2/users",
	}
	restClient := restclient.MockResponse[[]UserResponse]{}.
		NewRESTClient().
		AddMockRequest(mockRequest, GetUsersResponse(), restclient.NoNetworkError()).
		Build()

	var result restclient.Response[[]UserResponse]
	actual, err := restclient.
		Query[[]UserResponse]{RESTClient: restClient}.
		GetMock(http.MethodGet, "https://gorest.co.in/public/v2/users", &result)

	assert.NoError(t, err)
	assert.NotNil(t, actual)
	assert.NotNil(t, actual.Data)
	assert.Equal(t, 1, len(actual.Data))
	assert.Equal(t, int64(1), actual.Data[0].ID)
	assert.Equal(t, "John Doe", actual.Data[0].Name)
}

func TestExecute_GetMockError(t *testing.T) {
	mockRequest := restclient.MockRequest{
		Method: http.MethodGet,
		URL:    "https://gorest.co.in/public/v2/users",
	}
	restClient := restclient.MockResponse[[]UserResponse]{}.
		NewRESTClient().
		AddMockRequest(mockRequest, GetError(), restclient.NoNetworkError()).
		Build()

	var result restclient.Response[[]UserResponse]
	actual, err := restclient.
		Query[[]UserResponse]{RESTClient: restClient}.
		GetMock(http.MethodGet, "https://gorest.co.in/public/v2/users", &result)

	assert.Error(t, err)
	assert.NotNil(t, actual)
}

func TestExecute_GetMockConversionError(t *testing.T) {
	mockRequest := restclient.MockRequest{
		Method: http.MethodGet,
		URL:    "https://gorest.co.in/public/v2/users",
	}
	restClient := restclient.MockResponse[[]UserResponse]{}.
		NewRESTClient().
		AddMockRequest(mockRequest, GetError(), restclient.NoNetworkError()).
		Build()

	var result restclient.Response[UserResponse]
	actual, err := restclient.
		Query[UserResponse]{RESTClient: restClient}.
		GetMock(http.MethodGet, "https://gorest.co.in/public/v2/users", &result)

	assert.Error(t, err)
	assert.Equal(t, "Internal mocking error. ", err.Error())
	assert.NotNil(t, actual)
}

func TestExecute_GetMockNetworkError(t *testing.T) {
	mockRequest := restclient.MockRequest{
		Method: http.MethodGet,
		URL:    "https://gorest.co.in/public/v2/users",
	}
	restClient := restclient.MockResponse[[]UserResponse]{}.
		NewRESTClient().
		AddMockRequest(mockRequest, GetError(), restclient.NetworkError()).
		Build()

	var result restclient.Response[[]UserResponse]
	actual, err := restclient.
		Query[[]UserResponse]{RESTClient: restClient}.
		GetMock(http.MethodGet, "https://gorest.co.in/public/v2/users", &result)

	assert.Error(t, err)
	assert.NotNil(t, actual)
}

func TestExecute_Intercept_MethodGet(t *testing.T) {
	mockRequest := restclient.MockRequest{
		Method: http.MethodGet,
		URL:    "https://gorest.co.in/public/v2/users",
	}
	restClient := restclient.MockResponse[[]UserResponse]{}.
		NewRESTClient().
		AddMockRequest(mockRequest, GetUsersResponse(), restclient.NoNetworkError()).
		Build()

	actual, err := restclient.
		Query[[]UserResponse]{RESTClient: restClient}.
		Get("https://gorest.co.in/public/v2/users", nil)

	assert.NoError(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, 1, len(actual.Data))
	assert.Equal(t, int64(1), actual.Data[0].ID)
	assert.Equal(t, "John Doe", actual.Data[0].Name)
}

func TestExecute_Intercept_MethodPost(t *testing.T) {
	mockRequest := restclient.MockRequest{
		Method: http.MethodPost,
		URL:    "https://gorest.co.in/public/v2/users",
	}
	restClient := restclient.MockResponse[UserResponse]{}.
		NewRESTClient().
		AddMockRequest(mockRequest, GetUserResponse(), restclient.NoNetworkError()).
		Build()

	userResponse := UserResponse{Name: "John Doe"}

	actual, err := restclient.
		CommandQuery[UserResponse, UserResponse]{RESTClient: restClient}.
		Post("https://gorest.co.in/public/v2/users", userResponse, nil) //nolint:nolintlint,typecheck

	assert.NoError(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, int64(1), actual.Data.ID)
	assert.Equal(t, "John Doe", actual.Data.Name)
}

func GetUserResponse() restclient.Response[UserResponse] {
	userResponse := UserResponse{
		ID:   int64(1),
		Name: "John Doe",
	}
	return restclient.Response[UserResponse]{
		Data:   userResponse,
		Status: http.StatusOK,
	}
}

func GetUsersResponse() restclient.Response[[]UserResponse] {
	userResponse := UserResponse{
		ID:   int64(1),
		Name: "John Doe",
	}
	var result []UserResponse
	result = append(result, userResponse)

	return restclient.Response[[]UserResponse]{
		Data:   result,
		Status: http.StatusOK,
	}
}

func GetError() restclient.Response[[]UserResponse] {
	return restclient.Response[[]UserResponse]{
		Data:   nil,
		Status: http.StatusInternalServerError,
	}
}
