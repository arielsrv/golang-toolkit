package restclient_test

import (
	rest "github.com/arielsrv/golang-toolkit/restclient"
	"net/http"
	"testing"
)

func TestMockup(t *testing.T) {
	defer rest.StopMockupServer()
	rest.StartMockupServer()

	myURL := "http://mytest.com/foo?val1=1&val2=2#fragment"

	myHeaders := make(http.Header)
	myHeaders.Add("Hello", "world")

	mock := rest.Mock{
		URL:          myURL,
		HTTPMethod:   http.MethodGet,
		ReqHeaders:   myHeaders,
		RespHTTPCode: http.StatusOK,
		RespBody:     "foo",
	}

	_ = rest.AddMockups(&mock)

	v := rest.Get(myURL)
	if v.String() != "foo" {
		t.Fatal("Mockup Fail!")
	}
}
