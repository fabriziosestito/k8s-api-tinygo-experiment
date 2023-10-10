package socks

import (
	"context"
	"io"
	"net"
)

type Reply int

type Conn struct {
	net.Conn
	boundAddr net.Addr
}

type Dialer struct {
	cmd          Command
	proxyNetwork string
	proxyAddress string
	ProxyDial    func(context.Context, string, string) (net.Conn, error)
	AuthMethods  []AuthMethod
	Authenticate func(context.Context, io.ReadWriter, AuthMethod) error
}

type Command int

type UsernamePassword struct {
	Username string
	Password string
}

type AuthMethod int

type Addr struct {
	Name string
	IP   net.IP
	Port int
}

func (cmd Command) String() string {
	panic("stub")
}

func (code Reply) String() string {
	panic("stub")
}

func (a *Addr) Network() string {
	panic("stub")
}

func (a *Addr) String() string {
	panic("stub")
}

func (c *Conn) BoundAddr() net.Addr {
	panic("stub")
}

func (d *Dialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	panic("stub")
}

func (d *Dialer) DialWithConn(ctx context.Context, c net.Conn, network, address string) (net.Addr, error) {
	panic("stub")
}

func (d *Dialer) Dial(network, address string) (net.Conn, error) {
	panic("stub")
}

func NewDialer(network, address string) *Dialer {
	panic("stub")
}

func (up *UsernamePassword) Authenticate(ctx context.Context, rw io.ReadWriter, auth AuthMethod) error {
	panic("stub")
}

type Embedme interface{}
