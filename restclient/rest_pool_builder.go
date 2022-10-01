package restclient

import (
	"errors"
	"github.com/arielsrv/golang-toolkit/common"
	"time"
)

const MaxIdleConnections = 50
const MaxConnectionsPerHost = 50
const MaxIdleConnectionsPerHost = 100
const Timeout = time.Millisecond * 1000
const IdleConnectionTimeout = time.Millisecond * 1000
const TLSHandshakeTimeout = time.Second * 1000
const SocketTimeout = time.Millisecond * 5000
const SocketKeepAlive = time.Millisecond * 5000

type RESTPoolBuilder struct {
	Name                      string
	MaxConnectionsPerHost     int
	MaxIdleConnections        int
	MaxIdleConnectionsPerHost int
	Timeout                   time.Duration
	IdleConnectionTimeout     time.Duration
	TLSHandshakeTimeout       time.Duration
	SocketTimeout             time.Duration
	SocketKeepAlive           time.Duration
}

func NewRESTPoolBuilder() *RESTPoolBuilder {
	return &RESTPoolBuilder{
		MaxConnectionsPerHost:     MaxConnectionsPerHost,
		MaxIdleConnections:        MaxIdleConnections,
		MaxIdleConnectionsPerHost: MaxConnectionsPerHost,
		Timeout:                   Timeout,
		IdleConnectionTimeout:     IdleConnectionTimeout,
		TLSHandshakeTimeout:       TLSHandshakeTimeout,
		SocketTimeout:             SocketTimeout,
		SocketKeepAlive:           SocketKeepAlive,
	}
}

func (restPoolBuilder *RESTPoolBuilder) WithName(name string) *RESTPoolBuilder {
	restPoolBuilder.Name = name
	return restPoolBuilder
}

func (restPoolBuilder *RESTPoolBuilder) WithTimeout(timeout time.Duration) *RESTPoolBuilder {
	restPoolBuilder.Timeout = timeout
	return restPoolBuilder
}

func (restPoolBuilder *RESTPoolBuilder) WithSocketTimeout(socketTimeout time.Duration) *RESTPoolBuilder {
	restPoolBuilder.SocketTimeout = socketTimeout
	return restPoolBuilder
}

func (restPoolBuilder *RESTPoolBuilder) WithSocketKeepAlive(socketKeepAlive time.Duration) *RESTPoolBuilder {
	restPoolBuilder.SocketKeepAlive = socketKeepAlive
	return restPoolBuilder
}

func (restPoolBuilder *RESTPoolBuilder) WithTLSHandshakeTimeout(tlsHandshakeTimeout time.Duration) *RESTPoolBuilder {
	restPoolBuilder.TLSHandshakeTimeout = tlsHandshakeTimeout
	return restPoolBuilder
}

func (restPoolBuilder *RESTPoolBuilder) WithIdleConnectionTimeout(idleConnectionTimeout time.Duration) *RESTPoolBuilder {
	restPoolBuilder.IdleConnectionTimeout = idleConnectionTimeout
	return restPoolBuilder
}

func (restPoolBuilder *RESTPoolBuilder) WithMaxIdleConnections(maxIdleConnections int) *RESTPoolBuilder {
	restPoolBuilder.MaxIdleConnections = maxIdleConnections
	return restPoolBuilder
}

func (restPoolBuilder *RESTPoolBuilder) WithMaxConnectionsPerHost(maxConnectionsPerHost int) *RESTPoolBuilder {
	restPoolBuilder.MaxConnectionsPerHost = maxConnectionsPerHost
	return restPoolBuilder
}

func (restPoolBuilder *RESTPoolBuilder) WithMaxIdleConnectionsPerHost(maxIdleConnectionsPerHost int) *RESTPoolBuilder {
	restPoolBuilder.MaxIdleConnectionsPerHost = maxIdleConnectionsPerHost
	return restPoolBuilder
}

func (restPoolBuilder *RESTPoolBuilder) MakeDefault() *RESTPoolBuilder {
	restPoolBuilder.Name = "__default__"
	restPoolBuilder.MaxIdleConnections = MaxIdleConnections
	restPoolBuilder.MaxConnectionsPerHost = MaxConnectionsPerHost
	restPoolBuilder.MaxIdleConnectionsPerHost = MaxIdleConnectionsPerHost
	restPoolBuilder.Timeout = Timeout
	restPoolBuilder.IdleConnectionTimeout = IdleConnectionTimeout
	restPoolBuilder.TLSHandshakeTimeout = TLSHandshakeTimeout
	restPoolBuilder.SocketTimeout = SocketTimeout
	restPoolBuilder.SocketKeepAlive = SocketKeepAlive
	return restPoolBuilder
}

func (restPoolBuilder *RESTPoolBuilder) Build() (*RESTPool, error) {
	err := common.NotEmpty(restPoolBuilder.Name)
	if err != nil {
		return nil, errors.New("restPoolBuilder.Name cannot be empty. ")
	}
	return &RESTPool{
		Name:                      restPoolBuilder.Name,
		MaxIdleConnections:        restPoolBuilder.MaxIdleConnections,
		MaxConnectionsPerHost:     restPoolBuilder.MaxConnectionsPerHost,
		MaxIdleConnectionsPerHost: restPoolBuilder.MaxIdleConnectionsPerHost,
		Timeout:                   restPoolBuilder.Timeout,
		IdleConnectionTimeout:     restPoolBuilder.IdleConnectionTimeout,
		TLSHandshakeTimeout:       restPoolBuilder.TLSHandshakeTimeout,
		SocketTimeout:             restPoolBuilder.SocketTimeout,
		SocketKeepAlive:           restPoolBuilder.SocketKeepAlive,
	}, nil
}
