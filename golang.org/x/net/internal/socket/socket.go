package socket

import (
	"net"
	"sync"
	"syscall"
	"time"
)

type mmsghdrs []mmsghdr

type mmsghdrsPacker struct {
	hs        mmsghdrs
	sockaddrs []byte
	vs        []iovec
}

type syscaller struct {
	n              int
	operr          error
	hs             mmsghdrs
	flags          int
	boundRecvmmsgF func(uintptr) bool
	boundSendmmsgF func(uintptr) bool
}

type mmsgTmps struct {
	packer    mmsghdrsPacker
	syscaller syscaller
}

type mmsgTmpsPool struct{ p sync.Pool }

func (p *mmsgTmpsPool) Get() *mmsgTmps {
	panic("stub")
}

func (p *mmsgTmpsPool) Put(tmps *mmsgTmps) {
	panic("stub")
}

type Conn struct {
	network string
	c       syscall.RawConn
}

type tcpConn interface {
	SyscallConn() (syscall.RawConn, error)
	SetLinger(int) error
}

type udpConn interface {
	SyscallConn() (syscall.RawConn, error)
	ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *net.UDPAddr, err error)
}

type ipConn interface {
	SyscallConn() (syscall.RawConn, error)
	ReadMsgIP(b, oob []byte) (n, oobn, flags int, addr *net.IPAddr, err error)
}

func NewConn(c net.Conn) (*Conn, error) {
	panic("stub")
}

type ControlMessage []byte

type Message struct {
	Buffers [][]byte
	OOB     []byte
	Addr    net.Addr
	N       int
	NN      int
	Flags   int
}

type Option struct {
	Level int
	Name  int
	Len   int
}

func (o *Option) Get(c *Conn, b []byte) (int, error) {
	panic("stub")
}

func (o *Option) GetInt(c *Conn) (int, error) {
	panic("stub")
}

func (o *Option) Set(c *Conn, b []byte) error {
	panic("stub")
}

func (o *Option) SetInt(c *Conn, v int) error {
	panic("stub")
}

func ControlMessageSpace(dataLen int) int {
	panic("stub")
}

func (m ControlMessage) Data(dataLen int) []byte {
	panic("stub")
}

func (m ControlMessage) Next(dataLen int) ControlMessage {
	panic("stub")
}

func (m ControlMessage) MarshalHeader(lvl, typ, dataLen int) error {
	panic("stub")
}

func (m ControlMessage) ParseHeader() (lvl, typ, dataLen int, err error) {
	panic("stub")
}

func (m ControlMessage) Marshal(lvl, typ int, data []byte) (ControlMessage, error) {
	panic("stub")
}

func (m ControlMessage) Parse() ([]ControlMessage, error) {
	panic("stub")
}

func NewControlMessage(dataLen []int) ControlMessage {
	panic("stub")
}

func (c *Conn) RecvMsg(m *Message, flags int) error {
	panic("stub")
}

func (c *Conn) SendMsg(m *Message, flags int) error {
	panic("stub")
}

func (c *Conn) RecvMsgs(ms []Message, flags int) (int, error) {
	panic("stub")
}

func (c *Conn) SendMsgs(ms []Message, flags int) (int, error) {
	panic("stub")
}

type ipv6ZoneCache struct {
	sync.RWMutex
	lastFetched time.Time
	toIndex     map[string]int
	toName      map[int]string
}

type Embedme interface{}
