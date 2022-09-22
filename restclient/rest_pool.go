package restclient

import (
	"time"
)

type RESTPool struct {
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
