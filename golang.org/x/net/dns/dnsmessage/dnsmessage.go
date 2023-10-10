package dnsmessage

type Class uint16

type header struct {
	id          uint16
	bits        uint16
	questions   uint16
	answers     uint16
	authorities uint16
	additionals uint16
}

type CNAMEResource struct{ CNAME Name }

type section uint8

type SRVResource struct {
	Priority uint16
	Weight   uint16
	Port     uint16
	Target   Name
}

type AResource struct{ A [4]byte }

type OpCode uint16

type ResourceBody interface {
	pack(msg []byte, compression map[string]uint16, compressionOff int) ([]byte, error)
	realType() Type
	GoString() string
}

type Builder struct {
	msg         []byte
	section     section
	header      header
	start       int
	compression map[string]uint16
}

type Parser struct {
	msg             []byte
	header          header
	section         section
	off             int
	index           int
	resHeaderValid  bool
	resHeaderOffset int
	resHeaderType   Type
	resHeaderLength uint16
}

type Type uint16

type RCode uint16

type Header struct {
	ID                 uint16
	Response           bool
	OpCode             OpCode
	Authoritative      bool
	Truncated          bool
	RecursionDesired   bool
	RecursionAvailable bool
	AuthenticData      bool
	CheckingDisabled   bool
	RCode              RCode
}

type MXResource struct {
	Pref uint16
	MX   Name
}

type OPTResource struct{ Options []Option }

type TXTResource struct{ TXT []string }

type Option struct {
	Code uint16
	Data []byte
}

type Resource struct {
	Header ResourceHeader
	Body   ResourceBody
}

type Name struct {
	Data   [255]byte
	Length uint8
}

type Message struct {
	Header
	Questions   []Question
	Answers     []Resource
	Authorities []Resource
	Additionals []Resource
}

type UnknownResource struct {
	Type Type
	Data []byte
}

type Question struct {
	Name  Name
	Type  Type
	Class Class
}

type NSResource struct{ NS Name }

type ResourceHeader struct {
	Name   Name
	Type   Type
	Class  Class
	TTL    uint32
	Length uint16
}

type PTRResource struct{ PTR Name }

type SOAResource struct {
	NS      Name
	MBox    Name
	Serial  uint32
	Refresh uint32
	Retry   uint32
	Expire  uint32
	MinTTL  uint32
}

type AAAAResource struct{ AAAA [16]byte }

type nestedError struct {
	s   string
	err error
}

func (t Type) String() string {
	panic("stub")
}

func (t Type) GoString() string {
	panic("stub")
}

func (c Class) String() string {
	panic("stub")
}

func (c Class) GoString() string {
	panic("stub")
}

func (o OpCode) GoString() string {
	panic("stub")
}

func (r RCode) String() string {
	panic("stub")
}

func (r RCode) GoString() string {
	panic("stub")
}

func (e *nestedError) Error() string {
	panic("stub")
}

func (m *Header) GoString() string {
	panic("stub")
}

func (r *Resource) GoString() string {
	panic("stub")
}

func (p *Parser) Start(msg []byte) (Header, error) {
	panic("stub")
}

func (p *Parser) Question() (Question, error) {
	panic("stub")
}

func (p *Parser) AllQuestions() ([]Question, error) {
	panic("stub")
}

func (p *Parser) SkipQuestion() error {
	panic("stub")
}

func (p *Parser) SkipAllQuestions() error {
	panic("stub")
}

func (p *Parser) AnswerHeader() (ResourceHeader, error) {
	panic("stub")
}

func (p *Parser) Answer() (Resource, error) {
	panic("stub")
}

func (p *Parser) AllAnswers() ([]Resource, error) {
	panic("stub")
}

func (p *Parser) SkipAnswer() error {
	panic("stub")
}

func (p *Parser) SkipAllAnswers() error {
	panic("stub")
}

func (p *Parser) AuthorityHeader() (ResourceHeader, error) {
	panic("stub")
}

func (p *Parser) Authority() (Resource, error) {
	panic("stub")
}

func (p *Parser) AllAuthorities() ([]Resource, error) {
	panic("stub")
}

func (p *Parser) SkipAuthority() error {
	panic("stub")
}

func (p *Parser) SkipAllAuthorities() error {
	panic("stub")
}

func (p *Parser) AdditionalHeader() (ResourceHeader, error) {
	panic("stub")
}

func (p *Parser) Additional() (Resource, error) {
	panic("stub")
}

func (p *Parser) AllAdditionals() ([]Resource, error) {
	panic("stub")
}

func (p *Parser) SkipAdditional() error {
	panic("stub")
}

func (p *Parser) SkipAllAdditionals() error {
	panic("stub")
}

func (p *Parser) CNAMEResource() (CNAMEResource, error) {
	panic("stub")
}

func (p *Parser) MXResource() (MXResource, error) {
	panic("stub")
}

func (p *Parser) NSResource() (NSResource, error) {
	panic("stub")
}

func (p *Parser) PTRResource() (PTRResource, error) {
	panic("stub")
}

func (p *Parser) SOAResource() (SOAResource, error) {
	panic("stub")
}

func (p *Parser) TXTResource() (TXTResource, error) {
	panic("stub")
}

func (p *Parser) SRVResource() (SRVResource, error) {
	panic("stub")
}

func (p *Parser) AResource() (AResource, error) {
	panic("stub")
}

func (p *Parser) AAAAResource() (AAAAResource, error) {
	panic("stub")
}

func (p *Parser) OPTResource() (OPTResource, error) {
	panic("stub")
}

func (p *Parser) UnknownResource() (UnknownResource, error) {
	panic("stub")
}

func (m *Message) Unpack(msg []byte) error {
	panic("stub")
}

func (m *Message) Pack() ([]byte, error) {
	panic("stub")
}

func (m *Message) AppendPack(b []byte) ([]byte, error) {
	panic("stub")
}

func (m *Message) GoString() string {
	panic("stub")
}

func NewBuilder(buf []byte, h Header) Builder {
	panic("stub")
}

func (b *Builder) EnableCompression() {
	panic("stub")
}

func (b *Builder) StartQuestions() error {
	panic("stub")
}

func (b *Builder) StartAnswers() error {
	panic("stub")
}

func (b *Builder) StartAuthorities() error {
	panic("stub")
}

func (b *Builder) StartAdditionals() error {
	panic("stub")
}

func (b *Builder) Question(q Question) error {
	panic("stub")
}

func (b *Builder) CNAMEResource(h ResourceHeader, r CNAMEResource) error {
	panic("stub")
}

func (b *Builder) MXResource(h ResourceHeader, r MXResource) error {
	panic("stub")
}

func (b *Builder) NSResource(h ResourceHeader, r NSResource) error {
	panic("stub")
}

func (b *Builder) PTRResource(h ResourceHeader, r PTRResource) error {
	panic("stub")
}

func (b *Builder) SOAResource(h ResourceHeader, r SOAResource) error {
	panic("stub")
}

func (b *Builder) TXTResource(h ResourceHeader, r TXTResource) error {
	panic("stub")
}

func (b *Builder) SRVResource(h ResourceHeader, r SRVResource) error {
	panic("stub")
}

func (b *Builder) AResource(h ResourceHeader, r AResource) error {
	panic("stub")
}

func (b *Builder) AAAAResource(h ResourceHeader, r AAAAResource) error {
	panic("stub")
}

func (b *Builder) OPTResource(h ResourceHeader, r OPTResource) error {
	panic("stub")
}

func (b *Builder) UnknownResource(h ResourceHeader, r UnknownResource) error {
	panic("stub")
}

func (b *Builder) Finish() ([]byte, error) {
	panic("stub")
}

func (h *ResourceHeader) GoString() string {
	panic("stub")
}

func (h *ResourceHeader) SetEDNS0(udpPayloadLen int, extRCode RCode, dnssecOK bool) error {
	panic("stub")
}

func (h *ResourceHeader) DNSSECAllowed() bool {
	panic("stub")
}

func (h *ResourceHeader) ExtendedRCode(rcode RCode) RCode {
	panic("stub")
}

func NewName(name string) (Name, error) {
	panic("stub")
}

func MustNewName(name string) Name {
	panic("stub")
}

func (n Name) String() string {
	panic("stub")
}

func (n *Name) GoString() string {
	panic("stub")
}

func (q *Question) GoString() string {
	panic("stub")
}

func (r *CNAMEResource) GoString() string {
	panic("stub")
}

func (r *MXResource) GoString() string {
	panic("stub")
}

func (r *NSResource) GoString() string {
	panic("stub")
}

func (r *PTRResource) GoString() string {
	panic("stub")
}

func (r *SOAResource) GoString() string {
	panic("stub")
}

func (r *TXTResource) GoString() string {
	panic("stub")
}

func (r *SRVResource) GoString() string {
	panic("stub")
}

func (r *AResource) GoString() string {
	panic("stub")
}

func (r *AAAAResource) GoString() string {
	panic("stub")
}

func (o *Option) GoString() string {
	panic("stub")
}

func (r *OPTResource) GoString() string {
	panic("stub")
}

func (r *UnknownResource) GoString() string {
	panic("stub")
}

type Embedme interface{}
