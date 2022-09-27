package hashcode_test

import (
	"github.com/arielsrv/golang-toolkit/restclient"
	"github.com/arielsrv/golang-toolkit/restclient/hashcode"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test(t *testing.T) {
	value := hashcode.GetValue("hello world!")
	assert.Equal(t, uint64(0x58735284b97b86bc), value)
}

func TestObject(t *testing.T) {
	mockRequest := restclient.MockRequest{
		Method: http.MethodGet,
		URL:    "api.internal.com",
	}

	actual := mockRequest.GetHashCode()
	assert.Equal(t, uint64(0xda74b5e9b505f619), actual)
}
