package icmp

import (
	"net"
	"time"

	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

type DstUnreach struct {
	Data       []byte
	Extensions []Extension
}

func (p *DstUnreach) Len(proto int) int {
	panic("stub")
}

func (p *DstUnreach) Marshal(proto int) ([]byte, error) {
	panic("stub")
}

type Echo struct {
	ID   int
	Seq  int
	Data []byte
}

type ExtendedEchoRequest struct {
	ID         int
	Seq        int
	Local      bool
	Extensions []Extension
}

type ExtendedEchoReply struct {
	ID     int
	Seq    int
	State  int
	Active bool
	IPv4   bool
	IPv6   bool
}

func (p *Echo) Len(proto int) int {
	panic("stub")
}

func (p *Echo) Marshal(proto int) ([]byte, error) {
	panic("stub")
}

func (p *ExtendedEchoRequest) Len(proto int) int {
	panic("stub")
}

func (p *ExtendedEchoRequest) Marshal(proto int) ([]byte, error) {
	panic("stub")
}

func (p *ExtendedEchoReply) Len(proto int) int {
	panic("stub")
}

func (p *ExtendedEchoReply) Marshal(proto int) ([]byte, error) {
	panic("stub")
}

type PacketConn struct {
	c  net.PacketConn
	p4 *ipv4.PacketConn
	p6 *ipv6.PacketConn
}

func (c *PacketConn) IPv4PacketConn() *ipv4.PacketConn {
	panic("stub")
}

func (c *PacketConn) IPv6PacketConn() *ipv6.PacketConn {
	panic("stub")
}

func (c *PacketConn) ReadFrom(b []byte) (int, net.Addr, error) {
	panic("stub")
}

func (c *PacketConn) WriteTo(b []byte, dst net.Addr) (int, error) {
	panic("stub")
}

func (c *PacketConn) Close() error {
	panic("stub")
}

func (c *PacketConn) LocalAddr() net.Addr {
	panic("stub")
}

func (c *PacketConn) SetDeadline(t time.Time) error {
	panic("stub")
}

func (c *PacketConn) SetReadDeadline(t time.Time) error {
	panic("stub")
}

func (c *PacketConn) SetWriteDeadline(t time.Time) error {
	panic("stub")
}

type RawExtension struct{ Data []byte }

type Extension interface {
	Len(proto int) int
	Marshal(proto int) ([]byte, error)
}

func (p *RawExtension) Len(proto int) int {
	panic("stub")
}

func (p *RawExtension) Marshal(proto int) ([]byte, error) {
	panic("stub")
}

type InterfaceIdent struct {
	Class int
	Type  int
	Name  string
	Index int
	AFI   int
	Addr  []byte
}

type InterfaceInfo struct {
	Class     int
	Type      int
	Interface *net.Interface
	Addr      *net.IPAddr
}

func (ifi *InterfaceInfo) Len(proto int) int {
	panic("stub")
}

func (ifi *InterfaceInfo) Marshal(proto int) ([]byte, error) {
	panic("stub")
}

func (ifi *InterfaceIdent) Len(_ int) int {
	panic("stub")
}

func (ifi *InterfaceIdent) Marshal(proto int) ([]byte, error) {
	panic("stub")
}

func ParseIPv4Header(b []byte) (*ipv4.Header, error) {
	panic("stub")
}

func IPv6PseudoHeader(src, dst net.IP) []byte {
	panic("stub")
}

func ListenPacket(network, address string) (*PacketConn, error) {
	panic("stub")
}

type Message struct {
	Type     Type
	Code     int
	Checksum int
	Body     MessageBody
}

type Type interface {
	Protocol() int
}

func (m *Message) Marshal(psh []byte) ([]byte, error) {
	panic("stub")
}

func ParseMessage(proto int, b []byte) (*Message, error) {
	panic("stub")
}

type DefaultMessageBody RawBody

type MessageBody interface {
	Len(proto int) int
	Marshal(proto int) ([]byte, error)
}

type RawBody struct{ Data []byte }

func (p *RawBody) Len(proto int) int {
	panic("stub")
}

func (p *RawBody) Marshal(proto int) ([]byte, error) {
	panic("stub")
}

type MPLSLabelStack struct {
	Class  int
	Type   int
	Labels []MPLSLabel
}

type MPLSLabel struct {
	Label int
	TC    int
	S     bool
	TTL   int
}

func (ls *MPLSLabelStack) Len(proto int) int {
	panic("stub")
}

func (ls *MPLSLabelStack) Marshal(proto int) ([]byte, error) {
	panic("stub")
}

type PacketTooBig struct {
	MTU  int
	Data []byte
}

func (p *PacketTooBig) Len(proto int) int {
	panic("stub")
}

func (p *PacketTooBig) Marshal(proto int) ([]byte, error) {
	panic("stub")
}

type ParamProb struct {
	Pointer    uintptr
	Data       []byte
	Extensions []Extension
}

func (p *ParamProb) Len(proto int) int {
	panic("stub")
}

func (p *ParamProb) Marshal(proto int) ([]byte, error) {
	panic("stub")
}

type TimeExceeded struct {
	Data       []byte
	Extensions []Extension
}

func (p *TimeExceeded) Len(proto int) int {
	panic("stub")
}

func (p *TimeExceeded) Marshal(proto int) ([]byte, error) {
	panic("stub")
}

type Embedme interface{}
