package httpproxy

import (
	"net"
	"net/url"
)

type matcher interface {
	match(host, port string, ip net.IP) bool
}

type domainMatch struct {
	host      string
	port      string
	matchHost bool
}

type Config struct {
	HTTPProxy  string
	HTTPSProxy string
	NoProxy    string
	CGI        bool
}

type allMatch struct{}

type config struct {
	Config
	httpsProxy     *url.URL
	httpProxy      *url.URL
	ipMatchers     []matcher
	domainMatchers []matcher
}

type cidrMatch struct{ cidr *net.IPNet }

type ipMatch struct {
	ip   net.IP
	port string
}

func FromEnvironment() *Config {
	panic("stub")
}

func (cfg *Config) ProxyFunc() func(reqURL *url.URL) (*url.URL, error) {
	panic("stub")
}

type Embedme interface{}
