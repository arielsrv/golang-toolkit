package restclient_test

import (
	"github.com/arielsrv/golang-toolkit/restclient"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Default(t *testing.T) {
	restPool, err := restclient.NewRESTPoolBuilder().
		MakeDefault().
		Build()

	assert.NoError(t, err)
	assert.NotNil(t, restPool)
	assert.Equal(t, time.Second, restPool.Timeout)
	assert.Equal(t, time.Second, restPool.IdleConnectionTimeout)
	assert.Equal(t, 50, restPool.MaxIdleConnections)
	assert.Equal(t, 50, restPool.MaxConnectionsPerHost)
	assert.Equal(t, 100, restPool.MaxIdleConnectionsPerHost)
}

func Test_Config(t *testing.T) {
	restPool, err := restclient.NewRESTPoolBuilder().
		WithName("__default__").
		WithTimeout(time.Millisecond * 500).
		WithIdleConnectionTimeout(time.Millisecond * 500).
		WithMaxIdleConnections(20).
		WithMaxConnectionsPerHost(20).
		WithMaxIdleConnectionsPerHost(20).
		Build()

	assert.NoError(t, err)
	assert.NotNil(t, restPool)
	assert.Equal(t, time.Millisecond*500, restPool.Timeout)
	assert.Equal(t, time.Millisecond*500, restPool.IdleConnectionTimeout)
	assert.Equal(t, 20, restPool.MaxIdleConnections)
	assert.Equal(t, 20, restPool.MaxConnectionsPerHost)
	assert.Equal(t, 20, restPool.MaxIdleConnectionsPerHost)
}
