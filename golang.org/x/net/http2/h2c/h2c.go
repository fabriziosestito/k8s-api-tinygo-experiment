package h2c

import (
	"bufio"
	"net"
	"net/http"

	"golang.org/x/net/http2"
)

type h2cHandler struct {
	Handler http.Handler
	s       *http2.Server
}

type bufConn struct {
	net.Conn
	*bufio.Reader
}

func NewHandler(h http.Handler, s *http2.Server) http.Handler {
	panic("stub")
}

func (s h2cHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("stub")
}

func (c *bufConn) Read(p []byte) (int, error) {
	panic("stub")
}

type Embedme interface{}
