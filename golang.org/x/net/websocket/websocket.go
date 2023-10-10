package websocket

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type DialError struct {
	*Config
	Err error
}

func (e *DialError) Error() string {
	panic("stub")
}

func NewConfig(server, origin string) (config *Config, err error) {
	panic("stub")
}

func NewClient(config *Config, rwc io.ReadWriteCloser) (ws *Conn, err error) {
	panic("stub")
}

func Dial(url_, protocol, origin string) (ws *Conn, err error) {
	panic("stub")
}

func DialConfig(config *Config) (ws *Conn, err error) {
	panic("stub")
}

type hybiFrameHeader struct {
	Fin        bool
	Rsv        [3]bool
	OpCode     byte
	Length     int64
	MaskingKey []byte
	data       *bytes.Buffer
}

type hybiFrameWriterFactory struct {
	*bufio.Writer
	needMaskingKey bool
}

type hybiFrameHandler struct {
	conn        *Conn
	payloadType byte
}

type hybiFrameReaderFactory struct{ *bufio.Reader }

type hybiFrameWriter struct {
	writer *bufio.Writer
	header *hybiFrameHeader
}

type hybiServerHandshaker struct {
	*Config
	accept []byte
}

type hybiFrameReader struct {
	reader io.Reader
	header hybiFrameHeader
	pos    int64
	length int
}

func (frame *hybiFrameReader) Read(msg []byte) (n int, err error) {
	panic("stub")
}

func (frame *hybiFrameReader) PayloadType() byte {
	panic("stub")
}

func (frame *hybiFrameReader) HeaderReader() io.Reader {
	panic("stub")
}

func (frame *hybiFrameReader) TrailerReader() io.Reader {
	panic("stub")
}

func (frame *hybiFrameReader) Len() (n int) {
	panic("stub")
}

func (buf hybiFrameReaderFactory) NewFrameReader() (frame frameReader, err error) {
	panic("stub")
}

func (frame *hybiFrameWriter) Write(msg []byte) (n int, err error) {
	panic("stub")
}

func (frame *hybiFrameWriter) Close() error {
	panic("stub")
}

func (buf hybiFrameWriterFactory) NewFrameWriter(payloadType byte) (frame frameWriter, err error) {
	panic("stub")
}

func (handler *hybiFrameHandler) HandleFrame(frame frameReader) (frameReader, error) {
	panic("stub")
}

func (handler *hybiFrameHandler) WriteClose(status int) (err error) {
	panic("stub")
}

func (handler *hybiFrameHandler) WritePong(msg []byte) (n int, err error) {
	panic("stub")
}

func (c *hybiServerHandshaker) ReadHandshake(buf *bufio.Reader, req *http.Request) (code int, err error) {
	panic("stub")
}

func Origin(config *Config, req *http.Request) (*url.URL, error) {
	panic("stub")
}

func (c *hybiServerHandshaker) AcceptHandshake(buf *bufio.Writer) (err error) {
	panic("stub")
}

func (c *hybiServerHandshaker) NewServerConn(buf *bufio.ReadWriter, rwc io.ReadWriteCloser, request *http.Request) *Conn {
	panic("stub")
}

type Server struct {
	Config
	Handshake func(*Config, *http.Request) error
	Handler
}

type Handler func(*Conn)

func (s Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	panic("stub")
}

func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	panic("stub")
}

type ProtocolError struct{ ErrorString string }

type frameWriter interface {
}

type Conn struct {
	config  *Config
	request *http.Request
	buf     *bufio.ReadWriter
	rwc     io.ReadWriteCloser
	rio     sync.Mutex
	frameReaderFactory
	frameReader
	wio sync.Mutex
	frameWriterFactory
	frameHandler
	PayloadType        byte
	defaultCloseStatus int
	MaxPayloadBytes    int
}

type Codec struct {
	Marshal   func(v interface{}) (data []byte, payloadType byte, err error)
	Unmarshal func(data []byte, payloadType byte, v interface{}) (err error)
}

type serverHandshaker interface {
	ReadHandshake(buf *bufio.Reader, req *http.Request) (code int, err error)
	AcceptHandshake(buf *bufio.Writer) (err error)
	NewServerConn(buf *bufio.ReadWriter, rwc io.ReadWriteCloser, request *http.Request) (conn *Conn)
}

type Addr struct{ *url.URL }

type frameHandler interface {
	HandleFrame(frame frameReader) (r frameReader, err error)
	WriteClose(status int) (err error)
}

type Config struct {
	Location      *url.URL
	Origin        *url.URL
	Protocol      []string
	Version       int
	TlsConfig     *tls.Config
	Header        http.Header
	Dialer        *net.Dialer
	handshakeData map[string]string
}

type frameWriterFactory interface {
	NewFrameWriter(payloadType byte) (w frameWriter, err error)
}

type frameReader interface {
	PayloadType() byte
	HeaderReader() io.Reader
	TrailerReader() io.Reader
	Len() int
}

type frameReaderFactory interface {
	NewFrameReader() (r frameReader, err error)
}

func (err *ProtocolError) Error() string {
	panic("stub")
}

func (addr *Addr) Network() string {
	panic("stub")
}

func (ws *Conn) Read(msg []byte) (n int, err error) {
	panic("stub")
}

func (ws *Conn) Write(msg []byte) (n int, err error) {
	panic("stub")
}

func (ws *Conn) Close() error {
	panic("stub")
}

func (ws *Conn) IsClientConn() bool {
	panic("stub")
}

func (ws *Conn) IsServerConn() bool {
	panic("stub")
}

func (ws *Conn) LocalAddr() net.Addr {
	panic("stub")
}

func (ws *Conn) RemoteAddr() net.Addr {
	panic("stub")
}

func (ws *Conn) SetDeadline(t time.Time) error {
	panic("stub")
}

func (ws *Conn) SetReadDeadline(t time.Time) error {
	panic("stub")
}

func (ws *Conn) SetWriteDeadline(t time.Time) error {
	panic("stub")
}

func (ws *Conn) Config() *Config {
	panic("stub")
}

func (ws *Conn) Request() *http.Request {
	panic("stub")
}

func (cd Codec) Send(ws *Conn, v interface{}) (err error) {
	panic("stub")
}

func (cd Codec) Receive(ws *Conn, v interface{}) (err error) {
	panic("stub")
}

type Embedme interface{}
