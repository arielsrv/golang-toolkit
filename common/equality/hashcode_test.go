package equality_test

import (
	"github.com/arielsrv/golang-toolkit/common/equality"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test(t *testing.T) {
	value := equality.GetValue("hello world!")
	assert.Equal(t, uint64(0x58735284b97b86bc), value)
}

func TestObject(t *testing.T) {
	mockRequest := GetObject()

	actual := mockRequest.GetHashCode()
	assert.Equal(t, uint64(0xda74b5e9b505f619), actual)
}

type MockRequest struct {
	Method string
	URL    string
}

func (m MockRequest) GetHashCode() uint64 {
	hash := uint64(7)
	hash = uint64(31)*hash + equality.GetValue(m.Method)
	hash = uint64(31)*hash + equality.GetValue(m.URL)
	return hash
}

func GetObject() *MockRequest {
	return &MockRequest{
		Method: http.MethodGet,
		URL:    "api.internal.com",
	}
}
