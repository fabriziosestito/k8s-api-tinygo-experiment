package quic

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/tls"
	"net"
	"net/netip"
	"sync"
	"sync/atomic"
	"time"
)

type unscaledAckDelay int64

func (d unscaledAckDelay) Duration(ackDelayExponent uint8) time.Duration {
	panic("stub")
}

type ackState struct {
	seen                interface{}
	nextAck             time.Time
	maxRecvTime         time.Time
	maxAckEliciting     packetNumber
	unackedAckEliciting int
}

type atomicBits struct{ bits atomic.Uint32 }

type Config struct {
	TLSConfig                *tls.Config
	MaxBidiRemoteStreams     int64
	MaxUniRemoteStreams      int64
	MaxStreamReadBufferSize  int64
	MaxStreamWriteBufferSize int64
	MaxConnReadBufferSize    int64
}

type ccReno struct {
	maxDatagramSize         int
	congestionWindow        int
	bytesInFlight           int
	slowStartThreshold      int
	recoveryStartTime       time.Time
	congestionPendingAcks   int
	sendOnePacketInRecovery bool
	underutilized           bool
	ackLastLoss             time.Time
	persistentCongestion    [numberSpaceCount]struct{}
}

type Conn struct {
	side                 connSide
	listener             *Listener
	config               *Config
	testHooks            connTestHooks
	peerAddr             netip.AddrPort
	msgc                 chan any
	donec                chan struct{}
	exited               bool
	w                    packetWriter
	acks                 [numberSpaceCount]ackState
	lifetime             lifetimeState
	connIDState          connIDState
	loss                 lossState
	streams              streamsState
	maxIdleTimeout       time.Duration
	idleTimeout          time.Time
	keysInitial          fixedKeyPair
	keysHandshake        fixedKeyPair
	keysAppData          updatingKeyPair
	crypto               [numberSpaceCount]cryptoStream
	tls                  *tls.QUICConn
	handshakeConfirmed   sentVal
	peerAckDelayExponent int8
	testSendPingSpace    numberSpace
	testSendPing         sentVal
}

type connTestHooks interface {
	nextMessage(msgc chan any, nextTimeout time.Time) (now time.Time, message any)
	handleTLSEvent(tls.QUICEvent)
	newConnID(seq int64) ([]byte, error)
	waitUntil(ctx context.Context, until func() bool) error
	timeNow() time.Time
}

type timerEvent struct{}

type wakeEvent struct{}

func (c *Conn) String() string {
	panic("stub")
}

type lifetimeState struct {
	readyc            chan struct{}
	drainingc         chan struct{}
	localErr          error
	finalErr          error
	connCloseSentTime time.Time
	connCloseDelay    time.Duration
	drainEndTime      time.Time
}

func (c *Conn) Close() error {
	panic("stub")
}

func (c *Conn) Wait(ctx context.Context) error {
	panic("stub")
}

func (c *Conn) Abort(err error) {
	panic("stub")
}

type connInflow struct {
	sent      sentVal
	usedLimit int64
	sentLimit int64
	newLimit  int64
	credit    atomic.Int64
}

type connOutflow struct {
	max  int64
	used int64
}

type connIDState struct {
	local                 []connID
	remote                []connID
	nextLocalSeq          int64
	retireRemotePriorTo   int64
	peerActiveConnIDLimit int64
	needSend              bool
}

type connID struct {
	cid     []byte
	seq     int64
	retired bool
	send    sentVal
}

type streamsState struct {
	queue                             interface{}
	streamsMu                         sync.Mutex
	streams                           map[streamID]*Stream
	localLimit                        [streamTypeCount]localStreamLimits
	remoteLimit                       [streamTypeCount]remoteStreamLimits
	peerInitialMaxStreamDataRemote    [streamTypeCount]int64
	peerInitialMaxStreamDataBidiLocal int64
	inflow                            connInflow
	outflow                           connOutflow
	needSend                          atomic.Bool
	sendMu                            sync.Mutex
	queueMeta                         streamRing
	queueData                         streamRing
}

type streamFrameType uint8

type streamRing struct{ head *Stream }

func (c *Conn) AcceptStream(ctx context.Context) (*Stream, error) {
	panic("stub")
}

func (c *Conn) NewStream(ctx context.Context) (*Stream, error) {
	panic("stub")
}

func (c *Conn) NewSendOnlyStream(ctx context.Context) (*Stream, error) {
	panic("stub")
}

type cryptoStream struct {
	in        pipe
	inset     interface{}
	out       pipe
	outunsent interface{}
	outacked  interface{}
}

type datagram struct {
	b    []byte
	addr netip.AddrPort
}

type transportError uint64

type ApplicationError struct {
	Code   uint64
	Reason string
}

type localTransportError transportError

type peerTransportError struct {
	code   transportError
	reason string
}

type StreamErrorCode uint64

func (e transportError) String() string {
	panic("stub")
}

func (e localTransportError) Error() string {
	panic("stub")
}

func (e peerTransportError) Error() string {
	panic("stub")
}

func (e StreamErrorCode) Error() string {
	panic("stub")
}

func (e *ApplicationError) Error() string {
	panic("stub")
}

func (e *ApplicationError) Is(err error) bool {
	panic("stub")
}

type debugFrameStreamDataBlocked struct {
	id  streamID
	max int64
}

type debugFrameStreamsBlocked struct {
	streamType streamType
	max        int64
}

type debugFrameAck struct {
	ackDelay unscaledAckDelay
	ranges   []interface{}
}

type debugFrameStopSending struct {
	id   streamID
	code uint64
}

type debugFramePadding struct{ size int }

type debugFrameResetStream struct {
	id        streamID
	code      uint64
	finalSize int64
}

type debugFrameCrypto struct {
	off  int64
	data []byte
}

type debugFrameNewToken struct{ token []byte }

type debugFrameMaxStreams struct {
	streamType streamType
	max        int64
}

type debugFrameNewConnectionID struct {
	seq           int64
	retirePriorTo int64
	connID        []byte
	token         [16]byte
}

type debugFrame interface {
	String() string
	write(w *packetWriter) bool
}

type debugFrameStream struct {
	id   streamID
	fin  bool
	off  int64
	data []byte
}

type debugFrameRetireConnectionID struct{ seq int64 }

type debugFramePathResponse struct{ data uint64 }

type debugFrameHandshakeDone struct{}

type debugFrameDataBlocked struct{ max int64 }

type debugFrameMaxData struct{ max int64 }

type debugFramePathChallenge struct{ data uint64 }

type debugFramePing struct{}

type debugFrameMaxStreamData struct {
	id  streamID
	max int64
}

type debugFrameConnectionCloseTransport struct {
	code      transportError
	frameType uint64
	reason    string
}

type debugFrameConnectionCloseApplication struct {
	code   uint64
	reason string
}

func (f debugFramePadding) String() string {
	panic("stub")
}

func (f debugFramePing) String() string {
	panic("stub")
}

func (f debugFrameAck) String() string {
	panic("stub")
}

func (f debugFrameResetStream) String() string {
	panic("stub")
}

func (f debugFrameStopSending) String() string {
	panic("stub")
}

func (f debugFrameCrypto) String() string {
	panic("stub")
}

func (f debugFrameNewToken) String() string {
	panic("stub")
}

func (f debugFrameStream) String() string {
	panic("stub")
}

func (f debugFrameMaxData) String() string {
	panic("stub")
}

func (f debugFrameMaxStreamData) String() string {
	panic("stub")
}

func (f debugFrameMaxStreams) String() string {
	panic("stub")
}

func (f debugFrameDataBlocked) String() string {
	panic("stub")
}

func (f debugFrameStreamDataBlocked) String() string {
	panic("stub")
}

func (f debugFrameStreamsBlocked) String() string {
	panic("stub")
}

func (f debugFrameNewConnectionID) String() string {
	panic("stub")
}

func (f debugFrameRetireConnectionID) String() string {
	panic("stub")
}

func (f debugFramePathChallenge) String() string {
	panic("stub")
}

func (f debugFramePathResponse) String() string {
	panic("stub")
}

func (f debugFrameConnectionCloseTransport) String() string {
	panic("stub")
}

func (f debugFrameConnectionCloseApplication) String() string {
	panic("stub")
}

func (f debugFrameHandshakeDone) String() string {
	panic("stub")
}

type gate struct {
	set   chan struct{}
	unset chan struct{}
}

type Listener struct {
	config             *Config
	udpConn            udpConn
	testHooks          connTestHooks
	acceptQueue        interface{}
	connsMu            sync.Mutex
	conns              map[*Conn]struct{}
	closing            bool
	closec             chan struct{}
	connIDUpdateMu     sync.Mutex
	connIDUpdateNeeded atomic.Bool
	connIDUpdates      []connIDUpdate
}

type udpConn interface {
	Close() error
	LocalAddr() net.Addr
	ReadMsgUDPAddrPort(b, control []byte) (n, controln, flags int, _ netip.AddrPort, _ error)
	WriteToUDPAddrPort(b []byte, addr netip.AddrPort) (int, error)
}

type connIDUpdate struct {
	conn    *Conn
	retired bool
	cid     []byte
}

func Listen(network, address string, config *Config) (*Listener, error) {
	panic("stub")
}

func (l *Listener) LocalAddr() netip.AddrPort {
	panic("stub")
}

func (l *Listener) Close(ctx context.Context) error {
	panic("stub")
}

func (l *Listener) Accept(ctx context.Context) (*Conn, error) {
	panic("stub")
}

func (l *Listener) Dial(ctx context.Context, network, address string) (*Conn, error) {
	panic("stub")
}

type ccLimit int

type lossState struct {
	side                         connSide
	handshakeConfirmed           bool
	maxAckDelay                  time.Duration
	timer                        time.Time
	ptoTimerArmed                bool
	ptoExpired                   bool
	ptoBackoffCount              int
	antiAmplificationLimit       int
	rtt                          rttState
	pacer                        pacerState
	cc                           *ccReno
	spaces                       [numberSpaceCount]struct{}
	ackFrameRTT                  time.Duration
	ackFrameContainsAckEliciting bool
}

type pacerState struct {
	bucket           int
	maxBucket        int
	timerGranularity time.Duration
	lastUpdate       time.Time
	nextSend         time.Time
}

type shortPacket struct {
	num     packetNumber
	payload []byte
}

type packetType byte

type longPacket struct {
	ptype     packetType
	version   uint32
	num       packetNumber
	dstConnID []byte
	srcConnID []byte
	payload   []byte
	extra     []byte
}

type genericLongPacket struct {
	version   uint32
	dstConnID []byte
	srcConnID []byte
	data      []byte
}

func (p packetType) String() string {
	panic("stub")
}

type packetNumber int64

type fixedKeyPair struct {
	r
	w fixedKeys
}

type updatingKeys struct {
	suite      uint16
	hdr        headerKey
	pkt        [2]packetKey
	nextSecret []byte
}

type headerProtection interface {
	headerProtection(sample []byte) (mask [5]byte)
}

type aesHeaderProtection struct {
	cipher  cipher.Block
	scratch [aes.BlockSize]byte
}

type chaCha20HeaderProtection struct{ key []byte }

type fixedKeys struct {
	hdr headerKey
	pkt packetKey
}

type headerKey struct{ hp headerProtection }

type updatingKeyPair struct {
	phase        uint8
	updating     bool
	authFailures int64
	minSent      packetNumber
	minReceived  packetNumber
	updateAfter  packetNumber
	r
	w updatingKeys
}

type packetKey struct {
	aead cipher.AEAD
	iv   []byte
}

type packetWriter struct {
	dgramLim int
	pktLim   int
	pktOff   int
	payOff   int
	b        []byte
	sent     *sentPacket
}

type pipe struct {
	start int64
	end   int64
	head  *pipebuf
	tail  *pipebuf
}

type pipebuf struct {
	off  int64
	b    []byte
	next *pipebuf
}

type queue struct {
	gate gate
	err  error
	q    []T
}

type connSide int8

type streamID uint64

type numberSpace byte

type packetFate byte

type streamType uint8

func (s connSide) String() string {
	panic("stub")
}

func (n numberSpace) String() string {
	panic("stub")
}

func (s streamType) String() string {
	panic("stub")
}

type i64range struct {
	start
	end T
}

type rangeset []interface{}

type rttState struct {
	minRTT          time.Duration
	latestRTT       time.Duration
	smoothedRTT     time.Duration
	rttvar          time.Duration
	firstSampleTime time.Time
}

type sentPacket struct {
	num          packetNumber
	size         int
	time         time.Time
	ackEliciting bool
	inFlight     bool
	acked        bool
	lost         bool
	b            []byte
	n            int
}

type sentPacketList struct {
	nextNum packetNumber
	off     int
	size    int
	p       []*sentPacket
}

type sentVal uint64

type streamQueue int

type streamState uint32

type Stream struct {
	id           streamID
	conn         *Conn
	ingate       gate
	in           pipe
	inwin        int64
	insendmax    sentVal
	inmaxbuf     int64
	insize       int64
	inset        interface{}
	inclosed     sentVal
	inresetcode  int64
	outgate      gate
	out          pipe
	outwin       int64
	outmaxsent   int64
	outmaxbuf    int64
	outunsent    interface{}
	outacked     interface{}
	outopened    sentVal
	outclosed    sentVal
	outblocked   sentVal
	outreset     sentVal
	outresetcode uint64
	outdone      chan struct{}
	state        interface{}
	prev
	next *Stream
}

func (s *Stream) IsReadOnly() bool {
	panic("stub")
}

func (s *Stream) IsWriteOnly() bool {
	panic("stub")
}

func (s *Stream) Read(b []byte) (n int, err error) {
	panic("stub")
}

func (s *Stream) ReadContext(ctx context.Context, b []byte) (n int, err error) {
	panic("stub")
}

func (s *Stream) Write(b []byte) (n int, err error) {
	panic("stub")
}

func (s *Stream) WriteContext(ctx context.Context, b []byte) (n int, err error) {
	panic("stub")
}

func (s *Stream) Close() error {
	panic("stub")
}

func (s *Stream) CloseContext(ctx context.Context) error {
	panic("stub")
}

func (s *Stream) CloseRead() {
	panic("stub")
}

func (s *Stream) CloseWrite() {
	panic("stub")
}

func (s *Stream) Reset(code uint64) {
	panic("stub")
}

type localStreamLimits struct {
	gate   gate
	max    int64
	opened int64
}

type remoteStreamLimits struct {
	max     int64
	opened  int64
	closed  int64
	maxOpen int64
	sendMax sentVal
}

type transportParameters struct {
	originalDstConnID              []byte
	maxIdleTimeout                 time.Duration
	statelessResetToken            []byte
	maxUDPPayloadSize              int64
	initialMaxData                 int64
	initialMaxStreamDataBidiLocal  int64
	initialMaxStreamDataBidiRemote int64
	initialMaxStreamDataUni        int64
	initialMaxStreamsBidi          int64
	initialMaxStreamsUni           int64
	ackDelayExponent               int8
	maxAckDelay                    time.Duration
	disableActiveMigration         bool
	preferredAddrV4                netip.AddrPort
	preferredAddrV6                netip.AddrPort
	preferredAddrConnID            []byte
	preferredAddrResetToken        []byte
	activeConnIDLimit              int64
	initialSrcConnID               []byte
	retrySrcConnID                 []byte
}

type Embedme interface{}
