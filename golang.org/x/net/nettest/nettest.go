package nettest

import (
	"net"
	"testing"
)

type connTester func(t *testing.T, c1, c2 net.Conn)

type MakePipe func() (c1, c2 net.Conn, stop func(), err error)

func TestConn(t *testing.T, mp MakePipe) {
	panic("stub")
}

func SupportsIPv4() bool {
	panic("stub")
}

func SupportsIPv6() bool {
	panic("stub")
}

func SupportsRawSocket() bool {
	panic("stub")
}

func TestableNetwork(network string) bool {
	panic("stub")
}

func TestableAddress(network, address string) bool {
	panic("stub")
}

func NewLocalListener(network string) (net.Listener, error) {
	panic("stub")
}

func NewLocalPacketListener(network string) (net.PacketConn, error) {
	panic("stub")
}

func LocalPath() (string, error) {
	panic("stub")
}

func MulticastSource(network string, ifi *net.Interface) (net.IP, error) {
	panic("stub")
}

func LoopbackInterface() (*net.Interface, error) {
	panic("stub")
}

func RoutedInterface(network string, flags net.Flags) (*net.Interface, error) {
	panic("stub")
}

type Embedme interface{}
