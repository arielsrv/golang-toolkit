package core

import (
	"errors"
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

func (builder *RESTPoolBuilder) WithName(name string) *RESTPoolBuilder {
	builder.Name = name
	return builder
}

func (builder *RESTPoolBuilder) WithTimeout(timeout time.Duration) *RESTPoolBuilder {
	builder.Timeout = timeout
	return builder
}

func (builder *RESTPoolBuilder) WithSocketTimeout(socketTimeout time.Duration) *RESTPoolBuilder {
	builder.SocketTimeout = socketTimeout
	return builder
}

func (builder *RESTPoolBuilder) WithSocketKeepAlive(socketKeepAlive time.Duration) *RESTPoolBuilder {
	builder.SocketKeepAlive = socketKeepAlive
	return builder
}

func (builder *RESTPoolBuilder) WithTLSHandshakeTimeout(tlsHandshakeTimeout time.Duration) *RESTPoolBuilder {
	builder.TLSHandshakeTimeout = tlsHandshakeTimeout
	return builder
}

func (builder *RESTPoolBuilder) WithIdleConnectionTimeout(idleConnectionTimeout time.Duration) *RESTPoolBuilder {
	builder.IdleConnectionTimeout = idleConnectionTimeout
	return builder
}

func (builder *RESTPoolBuilder) WithMaxIdleConnections(maxIdleConnections int) *RESTPoolBuilder {
	builder.MaxIdleConnections = maxIdleConnections
	return builder
}

func (builder *RESTPoolBuilder) WithMaxConnectionsPerHost(maxConnectionsPerHost int) *RESTPoolBuilder {
	builder.MaxConnectionsPerHost = maxConnectionsPerHost
	return builder
}

func (builder *RESTPoolBuilder) WithMaxIdleConnectionsPerHost(maxIdleConnectionsPerHost int) *RESTPoolBuilder {
	builder.MaxIdleConnectionsPerHost = maxIdleConnectionsPerHost
	return builder
}

func (builder *RESTPoolBuilder) MakeDefault() *RESTPoolBuilder {
	builder.Name = "__default__"
	builder.MaxIdleConnections = MaxIdleConnections
	builder.MaxConnectionsPerHost = MaxConnectionsPerHost
	builder.MaxIdleConnectionsPerHost = MaxIdleConnectionsPerHost
	builder.Timeout = Timeout
	builder.IdleConnectionTimeout = IdleConnectionTimeout
	builder.TLSHandshakeTimeout = TLSHandshakeTimeout
	builder.SocketTimeout = SocketTimeout
	builder.SocketKeepAlive = SocketKeepAlive
	return builder
}

func (builder *RESTPoolBuilder) Build() (*RESTPool, error) {
	if builder.Name == "" {
		return nil, errors.New("builder.Name cannot be empty. ")
	}
	return &RESTPool{
		Name:                      builder.Name,
		MaxIdleConnections:        builder.MaxIdleConnections,
		MaxConnectionsPerHost:     builder.MaxConnectionsPerHost,
		MaxIdleConnectionsPerHost: builder.MaxIdleConnectionsPerHost,
		Timeout:                   builder.Timeout,
		IdleConnectionTimeout:     builder.IdleConnectionTimeout,
		TLSHandshakeTimeout:       builder.TLSHandshakeTimeout,
		SocketTimeout:             builder.SocketTimeout,
		SocketKeepAlive:           builder.SocketKeepAlive,
	}, nil
}
