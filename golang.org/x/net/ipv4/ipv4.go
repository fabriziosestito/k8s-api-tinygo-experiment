package ipv4

import (
	"net"
	"sync"
	"time"

	"golang.org/x/net/bpf"
	"golang.org/x/net/internal/socket"
)

type Message socket.Message

func (c *payloadHandler) ReadBatch(ms []Message, flags int) (int, error) {
	panic("stub")
}

func (c *payloadHandler) WriteBatch(ms []Message, flags int) (int, error) {
	panic("stub")
}

func (c *packetHandler) ReadBatch(ms []Message, flags int) (int, error) {
	panic("stub")
}

func (c *packetHandler) WriteBatch(ms []Message, flags int) (int, error) {
	panic("stub")
}

type rawOpt struct {
	sync.RWMutex
	cflags ControlFlags
}

type ControlMessage struct {
	TTL     int
	Src     net.IP
	Dst     net.IP
	IfIndex int
}

type ctlOpt struct {
	name    int
	length  int
	marshal func([]byte, *ControlMessage) []byte
	parse   func(*ControlMessage, []byte)
}

type ControlFlags uint

func (cm *ControlMessage) String() string {
	panic("stub")
}

func (cm *ControlMessage) Marshal() []byte {
	panic("stub")
}

func (cm *ControlMessage) Parse(b []byte) error {
	panic("stub")
}

func NewControlMessage(cf ControlFlags) []byte {
	panic("stub")
}

func (c *dgramOpt) MulticastTTL() (int, error) {
	panic("stub")
}

func (c *dgramOpt) SetMulticastTTL(ttl int) error {
	panic("stub")
}

func (c *dgramOpt) MulticastInterface() (*net.Interface, error) {
	panic("stub")
}

func (c *dgramOpt) SetMulticastInterface(ifi *net.Interface) error {
	panic("stub")
}

func (c *dgramOpt) MulticastLoopback() (bool, error) {
	panic("stub")
}

func (c *dgramOpt) SetMulticastLoopback(on bool) error {
	panic("stub")
}

func (c *dgramOpt) JoinGroup(ifi *net.Interface, group net.Addr) error {
	panic("stub")
}

func (c *dgramOpt) LeaveGroup(ifi *net.Interface, group net.Addr) error {
	panic("stub")
}

func (c *dgramOpt) JoinSourceSpecificGroup(ifi *net.Interface, group, source net.Addr) error {
	panic("stub")
}

func (c *dgramOpt) LeaveSourceSpecificGroup(ifi *net.Interface, group, source net.Addr) error {
	panic("stub")
}

func (c *dgramOpt) ExcludeSourceSpecificGroup(ifi *net.Interface, group, source net.Addr) error {
	panic("stub")
}

func (c *dgramOpt) IncludeSourceSpecificGroup(ifi *net.Interface, group, source net.Addr) error {
	panic("stub")
}

func (c *dgramOpt) ICMPFilter() (*ICMPFilter, error) {
	panic("stub")
}

func (c *dgramOpt) SetICMPFilter(f *ICMPFilter) error {
	panic("stub")
}

func (c *dgramOpt) SetBPF(filter []bpf.RawInstruction) error {
	panic("stub")
}

type dgramOpt struct{ *socket.Conn }

type RawConn struct {
	genericOpt
	dgramOpt
	packetHandler
}

type Conn struct{ genericOpt }

type genericOpt struct{ *socket.Conn }

type PacketConn struct {
	genericOpt
	dgramOpt
	payloadHandler
}

func NewConn(c net.Conn) *Conn {
	panic("stub")
}

func (c *PacketConn) SetControlMessage(cf ControlFlags, on bool) error {
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

func (c *PacketConn) Close() error {
	panic("stub")
}

func NewPacketConn(c net.PacketConn) *PacketConn {
	panic("stub")
}

func (c *RawConn) SetControlMessage(cf ControlFlags, on bool) error {
	panic("stub")
}

func (c *RawConn) SetDeadline(t time.Time) error {
	panic("stub")
}

func (c *RawConn) SetReadDeadline(t time.Time) error {
	panic("stub")
}

func (c *RawConn) SetWriteDeadline(t time.Time) error {
	panic("stub")
}

func (c *RawConn) Close() error {
	panic("stub")
}

func NewRawConn(c net.PacketConn) (*RawConn, error) {
	panic("stub")
}

func (c *genericOpt) TOS() (int, error) {
	panic("stub")
}

func (c *genericOpt) SetTOS(tos int) error {
	panic("stub")
}

func (c *genericOpt) TTL() (int, error) {
	panic("stub")
}

func (c *genericOpt) SetTTL(ttl int) error {
	panic("stub")
}

type Header struct {
	Version  int
	Len      int
	TOS      int
	TotalLen int
	ID       int
	Flags    HeaderFlags
	FragOff  int
	TTL      int
	Protocol int
	Checksum int
	Src      net.IP
	Dst      net.IP
	Options  []byte
}

type HeaderFlags int

func (h *Header) String() string {
	panic("stub")
}

func (h *Header) Marshal() ([]byte, error) {
	panic("stub")
}

func (h *Header) Parse(b []byte) error {
	panic("stub")
}

func ParseHeader(b []byte) (*Header, error) {
	panic("stub")
}

type ICMPType int

type ICMPFilter struct{ icmpFilter }

func (typ ICMPType) String() string {
	panic("stub")
}

func (typ ICMPType) Protocol() int {
	panic("stub")
}

func (f *ICMPFilter) Accept(typ ICMPType) {
	panic("stub")
}

func (f *ICMPFilter) Block(typ ICMPType) {
	panic("stub")
}

func (f *ICMPFilter) SetAll(block bool) {
	panic("stub")
}

func (f *ICMPFilter) WillBlock(typ ICMPType) bool {
	panic("stub")
}

type packetHandler struct {
	*net.IPConn
	*socket.Conn
	rawOpt
}

func (c *packetHandler) ReadFrom(b []byte) (h *Header, p []byte, cm *ControlMessage, err error) {
	panic("stub")
}

func (c *packetHandler) WriteTo(h *Header, p []byte, cm *ControlMessage) error {
	panic("stub")
}

type payloadHandler struct {
	net.PacketConn
	*socket.Conn
	rawOpt
}

func (c *payloadHandler) ReadFrom(b []byte) (n int, cm *ControlMessage, src net.Addr, err error) {
	panic("stub")
}

func (c *payloadHandler) WriteTo(b []byte, cm *ControlMessage, dst net.Addr) (n int, err error) {
	panic("stub")
}

type sockOpt struct {
	socket.Option
	typ int
}

type Embedme interface{}
