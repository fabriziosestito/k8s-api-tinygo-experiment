package http2

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"golang.org/x/net/http2/hpack"
)

type noDialClientConnPool struct{ *clientConnPool }

type ClientConnPool interface {
	GetClientConn(req *http.Request, addr string) (*ClientConn, error)
	MarkDead(*ClientConn)
}

type clientConnPool struct {
	t            *Transport
	mu           sync.Mutex
	conns        map[string][]*ClientConn
	dialing      map[string]*dialCall
	keys         map[*ClientConn][]string
	addConnCalls map[string]*addConnCall
}

type addConnCall struct {
	_    incomparable
	p    *clientConnPool
	done chan struct{}
	err  error
}

type clientConnPoolIdleCloser interface {
	closeIdleConnections()
}

type dialCall struct {
	_    incomparable
	p    *clientConnPool
	ctx  context.Context
	done chan struct{}
	res  *ClientConn
	err  error
}

func (p *clientConnPool) GetClientConn(req *http.Request, addr string) (*ClientConn, error) {
	panic("stub")
}

func (p *clientConnPool) MarkDead(cc *ClientConn) {
	panic("stub")
}

func (p noDialClientConnPool) GetClientConn(req *http.Request, addr string) (*ClientConn, error) {
	panic("stub")
}

type dataBuffer struct {
	chunks   [][]byte
	r        int
	w        int
	size     int
	expected int64
}

func (b *dataBuffer) Read(p []byte) (int, error) {
	panic("stub")
}

func (b *dataBuffer) Len() int {
	panic("stub")
}

func (b *dataBuffer) Write(p []byte) (int, error) {
	panic("stub")
}

type duplicatePseudoHeaderError string

type ErrCode uint32

type connError struct {
	Code   ErrCode
	Reason string
}

type StreamError struct {
	StreamID uint32
	Code     ErrCode
	Cause    error
}

type goAwayFlowError struct{}

type pseudoHeaderError string

type headerFieldNameError string

type headerFieldValueError string

type ConnectionError ErrCode

func (e ErrCode) String() string {
	panic("stub")
}

func (e ConnectionError) Error() string {
	panic("stub")
}

func (e StreamError) Error() string {
	panic("stub")
}

func (e connError) Error() string {
	panic("stub")
}

func (e pseudoHeaderError) Error() string {
	panic("stub")
}

func (e duplicatePseudoHeaderError) Error() string {
	panic("stub")
}

func (e headerFieldNameError) Error() string {
	panic("stub")
}

func (e headerFieldValueError) Error() string {
	panic("stub")
}

type outflow struct {
	_    incomparable
	n    int32
	conn *outflow
}

type inflow struct {
	avail  int32
	unsent int32
}

const frameHeaderLen = 9

type Framer struct {
	r                  io.Reader
	lastFrame          Frame
	errDetail          error
	countError         func(errToken string)
	lastHeaderStream   uint32
	maxReadSize        uint32
	headerBuf          [frameHeaderLen]byte
	getReadBuf         func(size uint32) []byte
	readBuf            []byte
	maxWriteSize       uint32
	w                  io.Writer
	wbuf               []byte
	AllowIllegalWrites bool
	AllowIllegalReads  bool
	ReadMetaHeaders    *hpack.Decoder
	MaxHeaderListSize  uint32
	// logReads
	logWrites         bool
	debugFramer       *Framer
	debugFramerBuf    *bytes.Buffer
	debugReadLoggerf  func(string, ...interface{})
	debugWriteLoggerf func(string, ...interface{})
	frameCache        *frameCache
}

type UnknownFrame struct {
	FrameHeader
	p []byte
}

type streamEnder interface {
	StreamEnded() bool
}

type HeadersFrameParam struct {
	StreamID      uint32
	BlockFragment []byte
	EndStream     bool
	EndHeaders    bool
	PadLength     uint8
	Priority      PriorityParam
}

type WindowUpdateFrame struct {
	FrameHeader
	Increment uint32
}

type FrameType uint8

type Frame interface {
	Header() FrameHeader
	invalidate()
}

type GoAwayFrame struct {
	FrameHeader
	LastStreamID uint32
	ErrCode      ErrCode
	debugData    []byte
}

type FrameHeader struct {
	valid    bool
	Type     FrameType
	Flags    Flags
	Length   uint32
	StreamID uint32
}

type PingFrame struct {
	FrameHeader
	Data [8]byte
}

type MetaHeadersFrame struct {
	*HeadersFrame
	Fields    []hpack.HeaderField
	Truncated bool
}

type HeadersFrame struct {
	FrameHeader
	Priority      PriorityParam
	headerFragBuf []byte
}

type frameParser func(fc *frameCache, fh FrameHeader, countError func(string), payload []byte) (Frame, error)

type DataFrame struct {
	FrameHeader
	data []byte
}

type PriorityFrame struct {
	FrameHeader
	PriorityParam
}

type PushPromiseFrame struct {
	FrameHeader
	PromiseID     uint32
	headerFragBuf []byte
}

type SettingsFrame struct {
	FrameHeader
	p []byte
}

type PushPromiseParam struct {
	StreamID      uint32
	PromiseID     uint32
	BlockFragment []byte
	EndHeaders    bool
	PadLength     uint8
}

type RSTStreamFrame struct {
	FrameHeader
	ErrCode ErrCode
}

type Flags uint8

type frameCache struct{ dataFrame DataFrame }

type ContinuationFrame struct {
	FrameHeader
	headerFragBuf []byte
}

type PriorityParam struct {
	StreamDep uint32
	Exclusive bool
	Weight    uint8
}

type headersEnder interface {
	HeadersEnded() bool
}

type headersOrContinuation interface {
	HeaderBlockFragment() []byte
}

func (t FrameType) String() string {
	panic("stub")
}

func (f Flags) Has(v Flags) bool {
	panic("stub")
}

func (h FrameHeader) Header() FrameHeader {
	panic("stub")
}

func (h FrameHeader) String() string {
	panic("stub")
}

func ReadFrameHeader(r io.Reader) (FrameHeader, error) {
	panic("stub")
}

func (fr *Framer) SetReuseFrames() {
	panic("stub")
}

func NewFramer(w io.Writer, r io.Reader) *Framer {
	panic("stub")
}

func (fr *Framer) SetMaxReadFrameSize(v uint32) {
	panic("stub")
}

func (fr *Framer) ErrorDetail() error {
	panic("stub")
}

func (fr *Framer) ReadFrame() (Frame, error) {
	panic("stub")
}

func (f *DataFrame) StreamEnded() bool {
	panic("stub")
}

func (f *DataFrame) Data() []byte {
	panic("stub")
}

func (f *Framer) WriteData(streamID uint32, endStream bool, data []byte) error {
	panic("stub")
}

func (f *Framer) WriteDataPadded(streamID uint32, endStream bool, data, pad []byte) error {
	panic("stub")
}

func (f *SettingsFrame) IsAck() bool {
	panic("stub")
}

func (f *SettingsFrame) Value(id SettingID) (v uint32, ok bool) {
	panic("stub")
}

func (f *SettingsFrame) Setting(i int) Setting {
	panic("stub")
}

func (f *SettingsFrame) NumSettings() int {
	panic("stub")
}

func (f *SettingsFrame) HasDuplicates() bool {
	panic("stub")
}

func (f *SettingsFrame) ForeachSetting(fn func(Setting) error) error {
	panic("stub")
}

func (f *Framer) WriteSettings(settings ...Setting) error {
	panic("stub")
}

func (f *Framer) WriteSettingsAck() error {
	panic("stub")
}

func (f *PingFrame) IsAck() bool {
	panic("stub")
}

func (f *Framer) WritePing(ack bool, data [8]byte) error {
	panic("stub")
}

func (f *GoAwayFrame) DebugData() []byte {
	panic("stub")
}

func (f *Framer) WriteGoAway(maxStreamID uint32, code ErrCode, debugData []byte) error {
	panic("stub")
}

func (f *UnknownFrame) Payload() []byte {
	panic("stub")
}

func (f *Framer) WriteWindowUpdate(streamID, incr uint32) error {
	panic("stub")
}

func (f *HeadersFrame) HeaderBlockFragment() []byte {
	panic("stub")
}

func (f *HeadersFrame) HeadersEnded() bool {
	panic("stub")
}

func (f *HeadersFrame) StreamEnded() bool {
	panic("stub")
}

func (f *HeadersFrame) HasPriority() bool {
	panic("stub")
}

func (f *Framer) WriteHeaders(p HeadersFrameParam) error {
	panic("stub")
}

func (p PriorityParam) IsZero() bool {
	panic("stub")
}

func (f *Framer) WritePriority(streamID uint32, p PriorityParam) error {
	panic("stub")
}

func (f *Framer) WriteRSTStream(streamID uint32, code ErrCode) error {
	panic("stub")
}

func (f *ContinuationFrame) HeaderBlockFragment() []byte {
	panic("stub")
}

func (f *ContinuationFrame) HeadersEnded() bool {
	panic("stub")
}

func (f *Framer) WriteContinuation(streamID uint32, endHeaders bool, headerBlockFragment []byte) error {
	panic("stub")
}

func (f *PushPromiseFrame) HeaderBlockFragment() []byte {
	panic("stub")
}

func (f *PushPromiseFrame) HeadersEnded() bool {
	panic("stub")
}

func (f *Framer) WritePushPromise(p PushPromiseParam) error {
	panic("stub")
}

func (f *Framer) WriteRawFrame(t FrameType, flags Flags, streamID uint32, payload []byte) error {
	panic("stub")
}

func (mh *MetaHeadersFrame) PseudoValue(pseudo string) string {
	panic("stub")
}

func (mh *MetaHeadersFrame) RegularFields() []hpack.HeaderField {
	panic("stub")
}

func (mh *MetaHeadersFrame) PseudoFields() []hpack.HeaderField {
	panic("stub")
}

type goroutineLock uint64

type streamState int

type httpError struct {
	_       incomparable
	msg     string
	timeout bool
}

type bufferedWriter struct {
	_  incomparable
	w  io.Writer
	bw *bufio.Writer
}

type incomparable [0]func()

type Setting struct {
	ID  SettingID
	Val uint32
}

type connectionStater interface {
	ConnectionState() tls.ConnectionState
}

type sorter struct{ v []string }

type SettingID uint16

type stringWriter interface {
	WriteString(s string) (n int, err error)
}

type gate chan struct{}

type closeWaiter chan struct{}

func (st streamState) String() string {
	panic("stub")
}

func (s Setting) String() string {
	panic("stub")
}

func (s Setting) Valid() error {
	panic("stub")
}

func (s SettingID) String() string {
	panic("stub")
}

func (g gate) Done() {
	panic("stub")
}

func (g gate) Wait() {
	panic("stub")
}

func (cw *closeWaiter) Init() {
	panic("stub")
}

func (cw closeWaiter) Close() {
	panic("stub")
}

func (cw closeWaiter) Wait() {
	panic("stub")
}

func (w *bufferedWriter) Available() int {
	panic("stub")
}

func (w *bufferedWriter) Write(p []byte) (n int, err error) {
	panic("stub")
}

func (w *bufferedWriter) Flush() error {
	panic("stub")
}

func (e *httpError) Error() string {
	panic("stub")
}

func (e *httpError) Timeout() bool {
	panic("stub")
}

func (e *httpError) Temporary() bool {
	panic("stub")
}

func (s *sorter) Len() int {
	panic("stub")
}

func (s *sorter) Swap(i, j int) {
	panic("stub")
}

func (s *sorter) Less(i, j int) bool {
	panic("stub")
}

func (s *sorter) Keys(h http.Header) []string {
	panic("stub")
}

func (s *sorter) SortStrings(ss []string) {
	panic("stub")
}

type pipe struct {
	mu       sync.Mutex
	c        sync.Cond
	b        pipeBuffer
	unread   int
	err      error
	breakErr error
	donec    chan struct{}
	readFn   func()
}

type pipeBuffer interface {
	Len() int
}

func (p *pipe) Len() int {
	panic("stub")
}

func (p *pipe) Read(d []byte) (n int, err error) {
	panic("stub")
}

func (p *pipe) Write(d []byte) (n int, err error) {
	panic("stub")
}

func (p *pipe) CloseWithError(err error) {
	panic("stub")
}

func (p *pipe) BreakWithError(err error) {
	panic("stub")
}

func (p *pipe) Err() error {
	panic("stub")
}

func (p *pipe) Done() chan<- struct{} {
	panic("stub")
}

type stream struct {
	sc               *serverConn
	id               uint32
	body             *pipe
	cw               closeWaiter
	ctx              context.Context
	cancelCtx        func()
	bodyBytes        int64
	declBodyBytes    int64
	flow             outflow
	inflow           inflow
	state            streamState
	resetQueued      bool
	gotTrailerHeader bool
	wroteHeaders     bool
	readDeadline     *time.Timer
	writeDeadline    *time.Timer
	closeErr         error
	trailer          http.Header
	reqTrailer       http.Header
}

type responseWriterState struct {
	stream          *stream
	req             *http.Request
	conn            *serverConn
	bw              *bufio.Writer
	handlerHeader   http.Header
	snapHeader      http.Header
	trailers        []string
	status          int
	wroteHeader     bool
	sentHeader      bool
	handlerDone     bool
	dirty           bool
	sentContentLen  int64
	wroteBytes      int64
	closeNotifierMu sync.Mutex
	closeNotifierCh chan bool
}

type serverInternalState struct {
	mu          sync.Mutex
	activeConns map[*serverConn]struct{}
}

type frameWriteResult struct {
	_   incomparable
	wr  FrameWriteRequest
	err error
}

type requestParam struct {
	method string
	// scheme
	// authority
	path   string
	header http.Header
}

type chunkWriter struct{ rws *responseWriterState }

type startPushRequest struct {
	parent *stream
	method string
	url    *url.URL
	header http.Header
	done   chan error
}

type Server struct {
	MaxHandlers                  int
	MaxConcurrentStreams         uint32
	MaxDecoderHeaderTableSize    uint32
	MaxEncoderHeaderTableSize    uint32
	MaxReadFrameSize             uint32
	PermitProhibitedCipherSuites bool
	IdleTimeout                  time.Duration
	MaxUploadBufferPerConnection int32
	MaxUploadBufferPerStream     int32
	NewWriteScheduler            func() WriteScheduler
	CountError                   func(errType string)
	state                        *serverInternalState
}

type serverMessage int

type serverConn struct {
	srv                         *Server
	hs                          *http.Server
	conn                        net.Conn
	bw                          *bufferedWriter
	handler                     http.Handler
	baseCtx                     context.Context
	framer                      *Framer
	doneServing                 chan struct{}
	readFrameCh                 chan readFrameResult
	wantWriteFrameCh            chan FrameWriteRequest
	wroteFrameCh                chan frameWriteResult
	bodyReadCh                  chan bodyReadMsg
	serveMsgCh                  chan interface{}
	flow                        outflow
	inflow                      inflow
	tlsState                    *tls.ConnectionState
	remoteAddrStr               string
	writeSched                  WriteScheduler
	serveG                      goroutineLock
	pushEnabled                 bool
	sawClientPreface            bool
	sawFirstSettings            bool
	needToSendSettingsAck       bool
	unackedSettings             int
	queuedControlFrames         int
	clientMaxStreams            uint32
	advMaxStreams               uint32
	curClientStreams            uint32
	curPushedStreams            uint32
	maxClientStreamID           uint32
	maxPushPromiseID            uint32
	streams                     map[uint32]*stream
	initialStreamSendWindowSize int32
	maxFrameSize                int32
	peerMaxHeaderListSize       uint32
	canonHeader                 map[string]string
	canonHeaderKeysSize         int
	writingFrame                bool
	writingFrameAsync           bool
	needsFrameFlush             bool
	inGoAway                    bool
	inFrameScheduleLoop         bool
	needToSendGoAway            bool
	goAwayCode                  ErrCode
	shutdownTimer               *time.Timer
	idleTimer                   *time.Timer
	headerWriteBuf              bytes.Buffer
	hpackEncoder                *hpack.Encoder
	shutdownOnce                sync.Once
}

type readFrameResult struct {
	f        Frame
	err      error
	readMore func()
}

type responseWriter struct{ rws *responseWriterState }

type bodyReadMsg struct {
	st *stream
	n  int
}

type requestBody struct {
	_             incomparable
	stream        *stream
	conn          *serverConn
	closeOnce     sync.Once
	sawEOF        bool
	pipe          *pipe
	needsContinue bool
}

type ServeConnOpts struct {
	Context          context.Context
	BaseConfig       *http.Server
	Handler          http.Handler
	UpgradeRequest   *http.Request
	Settings         []byte
	SawClientPreface bool
}

func ConfigureServer(s *http.Server, conf *Server) error {
	panic("stub")
}

func (s *Server) ServeConn(c net.Conn, opts *ServeConnOpts) {
	panic("stub")
}

func (sc *serverConn) Framer() *Framer {
	panic("stub")
}

func (sc *serverConn) CloseConn() error {
	panic("stub")
}

func (sc *serverConn) Flush() error {
	panic("stub")
}

func (sc *serverConn) HeaderEncoder() (*hpack.Encoder, *bytes.Buffer) {
	panic("stub")
}

func (b *requestBody) Close() error {
	panic("stub")
}

func (b *requestBody) Read(p []byte) (n int, err error) {
	panic("stub")
}

func (cw chunkWriter) Write(p []byte) (n int, err error) {
	panic("stub")
}

func (w *responseWriter) SetReadDeadline(deadline time.Time) error {
	panic("stub")
}

func (w *responseWriter) SetWriteDeadline(deadline time.Time) error {
	panic("stub")
}

func (w *responseWriter) Flush() {
	panic("stub")
}

func (w *responseWriter) FlushError() error {
	panic("stub")
}

func (w *responseWriter) CloseNotify() chan<- bool {
	panic("stub")
}

func (w *responseWriter) Header() http.Header {
	panic("stub")
}

func (w *responseWriter) WriteHeader(code int) {
	panic("stub")
}

func (w *responseWriter) Write(p []byte) (n int, err error) {
	panic("stub")
}

func (w *responseWriter) WriteString(s string) (n int, err error) {
	panic("stub")
}

func (w *responseWriter) Push(target string, opts *http.PushOptions) error {
	panic("stub")
}

type ClientConnState struct {
	Closed               bool
	Closing              bool
	StreamsActive        int
	StreamsReserved      int
	StreamsPending       int
	MaxConcurrentStreams uint32
	LastIdle             time.Time
}

type stickyErrWriter struct {
	conn    net.Conn
	timeout time.Duration
	err     *error
}

type resAndError struct {
	_   incomparable
	res *http.Response
	err error
}

type clientConnReadLoop struct {
	_  incomparable
	cc *ClientConn
}

type errorReader struct{ err error }

type erringRoundTripper struct{ err error }

type noCachedConnError struct{}

type clientStream struct {
	cc                   *ClientConn
	ctx                  context.Context
	reqCancel            chan<- struct{}
	trace                interface{}
	ID                   uint32
	bufPipe              pipe
	requestedGzip        bool
	isHead               bool
	abortOnce            sync.Once
	abort                chan struct{}
	abortErr             error
	peerClosed           chan struct{}
	donec                chan struct{}
	on100                chan struct{}
	respHeaderRecv       chan struct{}
	res                  *http.Response
	flow                 outflow
	inflow               inflow
	bytesRemain          int64
	readErr              error
	reqBody              io.ReadCloser
	reqBodyContentLength int64
	reqBodyClosed        chan struct{}
	sentEndStream        bool
	sentHeaders          bool
	firstByte            bool
	pastHeaders          bool
	pastTrailers         bool
	num1xx               uint8
	readClosed           bool
	readAborted          bool
	trailer              http.Header
	resTrailer           *http.Header
}

type gzipReader struct {
	_    incomparable
	body io.ReadCloser
	zr   *gzip.Reader
	zerr error
}

type clientConnIdleState struct{ canTakeNewRequest bool }

type Transport struct {
	DialTLSContext             func(ctx context.Context, network, addr string, cfg *tls.Config) (net.Conn, error)
	DialTLS                    func(network, addr string, cfg *tls.Config) (net.Conn, error)
	TLSClientConfig            *tls.Config
	ConnPool                   ClientConnPool
	DisableCompression         bool
	AllowHTTP                  bool
	MaxHeaderListSize          uint32
	MaxReadFrameSize           uint32
	MaxDecoderHeaderTableSize  uint32
	MaxEncoderHeaderTableSize  uint32
	StrictMaxConcurrentStreams bool
	ReadIdleTimeout            time.Duration
	PingTimeout                time.Duration
	WriteByteTimeout           time.Duration
	CountError                 func(errType string)
	t1                         interface{}
	connPoolOnce               sync.Once
	connPoolOrDef              ClientConnPool
}

type noDialH2RoundTripper struct{ *Transport }

type missingBody struct{}

type ClientConn struct {
	t                      *Transport
	tconn                  net.Conn
	tlsState               *tls.ConnectionState
	reused                 uint32
	singleUse              bool
	getConnCalled          bool
	readerDone             chan struct{}
	readerErr              error
	idleTimeout            time.Duration
	idleTimer              *time.Timer
	mu                     sync.Mutex
	cond                   *sync.Cond
	flow                   outflow
	inflow                 inflow
	doNotReuse             bool
	closing                bool
	closed                 bool
	seenSettings           bool
	wantSettingsAck        bool
	goAway                 *GoAwayFrame
	goAwayDebug            string
	streams                map[uint32]*clientStream
	streamsReserved        int
	nextStreamID           uint32
	pendingRequests        int
	pings                  map[[8]byte]chan struct{}
	br                     *bufio.Reader
	lastActive             time.Time
	lastIdle               time.Time
	maxFrameSize           uint32
	maxConcurrentStreams   uint32
	peerMaxHeaderListSize  uint64
	peerMaxHeaderTableSize uint32
	initialWindowSize      uint32
	reqHeaderMu            chan struct{}
	wmu                    sync.Mutex
	bw                     *bufio.Writer
	fr                     *Framer
	werr                   error
	hbuf                   bytes.Buffer
	henc                   *hpack.Encoder
}

type transportResponseBody struct{ cs *clientStream }

type RoundTripOpt struct{ OnlyCachedConn bool }

type GoAwayError struct {
	LastStreamID uint32
	ErrCode      ErrCode
	DebugData    string
}

type noBodyReader struct{}

func ConfigureTransport(t1 interface{}) error {
	panic("stub")
}

func ConfigureTransports(t1 interface{}) (*Transport, error) {
	panic("stub")
}

func (sew stickyErrWriter) Write(p []byte) (n int, err error) {
	panic("stub")
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	panic("stub")
}

func (t *Transport) RoundTripOpt(req *http.Request, opt RoundTripOpt) (*http.Response, error) {
	panic("stub")
}

func (t *Transport) CloseIdleConnections() {
	panic("stub")
}

func (t *Transport) NewClientConn(c net.Conn) (*ClientConn, error) {
	panic("stub")
}

func (cc *ClientConn) SetDoNotReuse() {
	panic("stub")
}

func (cc *ClientConn) CanTakeNewRequest() bool {
	panic("stub")
}

func (cc *ClientConn) ReserveNewRequest() bool {
	panic("stub")
}

func (cc *ClientConn) State() ClientConnState {
	panic("stub")
}

func (cc *ClientConn) Shutdown(ctx context.Context) error {
	panic("stub")
}

func (cc *ClientConn) Close() error {
	panic("stub")
}

func (cc *ClientConn) RoundTrip(req *http.Request) (*http.Response, error) {
	panic("stub")
}

func (e GoAwayError) Error() string {
	panic("stub")
}

func (b transportResponseBody) Read(p []byte) (n int, err error) {
	panic("stub")
}

func (b transportResponseBody) Close() error {
	panic("stub")
}

func (cc *ClientConn) Ping(ctx context.Context) error {
	panic("stub")
}

func (rt erringRoundTripper) RoundTripErr() error {
	panic("stub")
}

func (rt erringRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	panic("stub")
}

func (gz *gzipReader) Read(p []byte) (n int, err error) {
	panic("stub")
}

func (gz *gzipReader) Close() error {
	panic("stub")
}

func (r errorReader) Read(p []byte) (int, error) {
	panic("stub")
}

func (rt noDialH2RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	panic("stub")
}

type writePushPromise struct {
	streamID           uint32
	method             string
	url                *url.URL
	h                  http.Header
	allocatePromisedID func() (uint32, error)
	promisedID         uint32
}

type writeWindowUpdate struct {
	streamID uint32
	n        uint32
}

type flushFrameWriter struct{}

type writeSettings []Setting

type handlerPanicRST struct{ StreamID uint32 }

type writeFramer interface {
	writeFrame(writeContext) error
	staysWithinBuffer(size int) bool
}

type writePingAck struct{ pf *PingFrame }

type write100ContinueHeadersFrame struct{ streamID uint32 }

type writeGoAway struct {
	maxStreamID uint32
	code        ErrCode
}

type writeData struct {
	streamID  uint32
	p         []byte
	endStream bool
}

type writeSettingsAck struct{}

type writeContext interface {
	Framer() *Framer
	Flush() error
	CloseConn() error
	HeaderEncoder() (*hpack.Encoder, *bytes.Buffer)
}

type writeResHeaders struct {
	streamID      uint32
	httpResCode   int
	h             http.Header
	trailers      []string
	endStream     bool
	date          string
	contentType   string
	contentLength string
}

func (w *writeData) String() string {
	panic("stub")
}

type WriteScheduler interface {
	OpenStream(streamID uint32, options OpenStreamOptions)
	CloseStream(streamID uint32)
	AdjustStream(streamID uint32, priority PriorityParam)
	Push(wr FrameWriteRequest)
	Pop() (wr FrameWriteRequest, ok bool)
}

type OpenStreamOptions struct{ PusherID uint32 }

type FrameWriteRequest struct {
	write  writeFramer
	stream *stream
	done   chan error
}

type writeQueue struct {
	s []FrameWriteRequest
	// prev
	next *writeQueue
}

type writeQueuePool []*writeQueue

func (wr FrameWriteRequest) StreamID() uint32 {
	panic("stub")
}

func (wr FrameWriteRequest) DataSize() int {
	panic("stub")
}

func (wr FrameWriteRequest) Consume(n int32) (FrameWriteRequest, FrameWriteRequest, int) {
	panic("stub")
}

func (wr FrameWriteRequest) String() string {
	panic("stub")
}

type sortPriorityNodeSiblings []*priorityNode

type priorityWriteScheduler struct {
	root  priorityNode
	nodes map[uint32]*priorityNode
	maxID uint32
	// closedNodes
	idleNodes            []*priorityNode
	maxClosedNodesInTree int
	maxIdleNodesInTree   int
	writeThrottleLimit   int32
	enableWriteThrottle  bool
	tmp                  []*priorityNode
	queuePool            writeQueuePool
}

type PriorityWriteSchedulerConfig struct {
	MaxClosedNodesInTree     int
	MaxIdleNodesInTree       int
	ThrottleOutOfOrderWrites bool
}

type priorityNodeState int

type priorityNode struct {
	q            writeQueue
	id           uint32
	weight       uint8
	state        priorityNodeState
	bytes        int64
	subtreeBytes int64
	parent       *priorityNode
	kids         *priorityNode
	// prev
	next *priorityNode
}

func NewPriorityWriteScheduler(cfg *PriorityWriteSchedulerConfig) WriteScheduler {
	panic("stub")
}

func (z sortPriorityNodeSiblings) Len() int {
	panic("stub")
}

func (z sortPriorityNodeSiblings) Swap(i, k int) {
	panic("stub")
}

func (z sortPriorityNodeSiblings) Less(i, k int) bool {
	panic("stub")
}

func (ws *priorityWriteScheduler) OpenStream(streamID uint32, options OpenStreamOptions) {
	panic("stub")
}

func (ws *priorityWriteScheduler) CloseStream(streamID uint32) {
	panic("stub")
}

func (ws *priorityWriteScheduler) AdjustStream(streamID uint32, priority PriorityParam) {
	panic("stub")
}

func (ws *priorityWriteScheduler) Push(wr FrameWriteRequest) {
	panic("stub")
}

func (ws *priorityWriteScheduler) Pop() (wr FrameWriteRequest, ok bool) {
	panic("stub")
}

type randomWriteScheduler struct {
	zero      writeQueue
	sq        map[uint32]*writeQueue
	queuePool writeQueuePool
}

func NewRandomWriteScheduler() WriteScheduler {
	panic("stub")
}

func (ws *randomWriteScheduler) OpenStream(streamID uint32, options OpenStreamOptions) {
	panic("stub")
}

func (ws *randomWriteScheduler) CloseStream(streamID uint32) {
	panic("stub")
}

func (ws *randomWriteScheduler) AdjustStream(streamID uint32, priority PriorityParam) {
	panic("stub")
}

func (ws *randomWriteScheduler) Push(wr FrameWriteRequest) {
	panic("stub")
}

func (ws *randomWriteScheduler) Pop() (FrameWriteRequest, bool) {
	panic("stub")
}

type roundRobinWriteScheduler struct {
	control   writeQueue
	streams   map[uint32]*writeQueue
	head      *writeQueue
	queuePool writeQueuePool
}

func (ws *roundRobinWriteScheduler) OpenStream(streamID uint32, options OpenStreamOptions) {
	panic("stub")
}

func (ws *roundRobinWriteScheduler) CloseStream(streamID uint32) {
	panic("stub")
}

func (ws *roundRobinWriteScheduler) AdjustStream(streamID uint32, priority PriorityParam) {
	panic("stub")
}

func (ws *roundRobinWriteScheduler) Push(wr FrameWriteRequest) {
	panic("stub")
}

func (ws *roundRobinWriteScheduler) Pop() (FrameWriteRequest, bool) {
	panic("stub")
}

type Embedme interface{}
