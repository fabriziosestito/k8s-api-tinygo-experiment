package proxy

import (
	"context"
	"net"
	"net/url"
	"sync"
)

type ContextDialer interface {
	DialContext(ctx context.Context, network, address string) (net.Conn, error)
}

func Dial(ctx context.Context, network, address string) (net.Conn, error) {
	panic("stub")
}

type direct struct{}

type PerHost struct {
	def
	bypass         Dialer
	bypassNetworks []*net.IPNet
	bypassIPs      []net.IP
	bypassZones    []string
	bypassHosts    []string
}

func NewPerHost(defaultDialer, bypass Dialer) *PerHost {
	panic("stub")
}

func (p *PerHost) Dial(network, addr string) (c net.Conn, err error) {
	panic("stub")
}

func (p *PerHost) DialContext(ctx context.Context, network, addr string) (c net.Conn, err error) {
	panic("stub")
}

func (p *PerHost) AddFromString(s string) {
	panic("stub")
}

func (p *PerHost) AddIP(ip net.IP) {
	panic("stub")
}

func (p *PerHost) AddNetwork(net *net.IPNet) {
	panic("stub")
}

func (p *PerHost) AddZone(zone string) {
	panic("stub")
}

func (p *PerHost) AddHost(host string) {
	panic("stub")
}

type Auth struct {
	User
	Password string
}

type envOnce struct {
	names []string
	once  sync.Once
	val   string
}

type Dialer interface {
	Dial(network, addr string) (c net.Conn, err error)
}

func FromEnvironment() Dialer {
	panic("stub")
}

func FromEnvironmentUsing(forward Dialer) Dialer {
	panic("stub")
}

func RegisterDialerType(scheme string, f func(*url.URL, Dialer) (Dialer, error)) {
	panic("stub")
}

func FromURL(u *url.URL, forward Dialer) (Dialer, error) {
	panic("stub")
}

func (e *envOnce) Get() string {
	panic("stub")
}

func SOCKS5(network, address string, auth *Auth, forward Dialer) (Dialer, error) {
	panic("stub")
}

type Embedme interface{}
