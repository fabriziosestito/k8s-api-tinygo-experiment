package sockstest

import (
	"io"
	"net"

	"golang.org/x/net/internal/socks"
)

type CmdRequest struct {
	Version int
	Cmd     socks.Command
	Addr    socks.Addr
}

type AuthRequest struct {
	Version int
	Methods []socks.AuthMethod
}

type Server struct{ ln net.Listener }

func ParseAuthRequest(b []byte) (*AuthRequest, error) {
	panic("stub")
}

func MarshalAuthReply(ver int, m socks.AuthMethod) ([]byte, error) {
	panic("stub")
}

func ParseCmdRequest(b []byte) (*CmdRequest, error) {
	panic("stub")
}

func MarshalCmdReply(ver int, reply socks.Reply, a *socks.Addr) ([]byte, error) {
	panic("stub")
}

func (s *Server) Addr() net.Addr {
	panic("stub")
}

func (s *Server) TargetAddr() net.Addr {
	panic("stub")
}

func (s *Server) Close() error {
	panic("stub")
}

func NewServer(authFunc, cmdFunc func(io.ReadWriter, []byte) error) (*Server, error) {
	panic("stub")
}

func NoAuthRequired(rw io.ReadWriter, b []byte) error {
	panic("stub")
}

func NoProxyRequired(rw io.ReadWriter, b []byte) error {
	panic("stub")
}

type Embedme interface{}
