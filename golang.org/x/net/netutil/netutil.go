package netutil

import (
	"net"
	"sync"
)

type limitListener struct {
	net.Listener
	sem       chan struct{}
	closeOnce sync.Once
	done      chan struct{}
}

type limitListenerConn struct {
	net.Conn
	releaseOnce sync.Once
	release     func()
}

func LimitListener(l net.Listener, n int) net.Listener {
	panic("stub")
}

func (l *limitListener) Accept() (net.Conn, error) {
	panic("stub")
}

func (l *limitListener) Close() error {
	panic("stub")
}

func (l *limitListenerConn) Close() error {
	panic("stub")
}

type Embedme interface{}
