package xml

import (
	"bufio"
	"bytes"
	"io"
	"reflect"
)

type Marshaler interface {
	MarshalXML(e *Encoder, start StartElement) error
}

type MarshalerAttr interface {
	MarshalXMLAttr(name Name) (Attr, error)
}

type Encoder struct{ p printer }

type printerPrefix struct {
	prefix string
	url    string
	mark   bool
}

type parentStack struct {
	p       *printer
	xmlns   string
	parents []string
}

type UnsupportedTypeError struct{ Type reflect.Type }

type printer struct {
	*bufio.Writer
	encoder    *Encoder
	seq        int
	indent     string
	prefix     string
	depth      int
	indentedIn bool
	putNewline bool
	defaultNS  string
	attrNS     map[string]string
	attrPrefix map[string]string
	prefixes   []printerPrefix
	tags       []Name
}

func Marshal(v interface{}) ([]byte, error) {
	panic("stub")
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	panic("stub")
}

func NewEncoder(w io.Writer) *Encoder {
	panic("stub")
}

func (enc *Encoder) Indent(prefix, indent string) {
	panic("stub")
}

func (enc *Encoder) Encode(v interface{}) error {
	panic("stub")
}

func (enc *Encoder) EncodeElement(v interface{}, start StartElement) error {
	panic("stub")
}

func (enc *Encoder) EncodeToken(t Token) error {
	panic("stub")
}

func (enc *Encoder) Flush() error {
	panic("stub")
}

func (e *UnsupportedTypeError) Error() string {
	panic("stub")
}

type UnmarshalError string

type Unmarshaler interface {
	UnmarshalXML(d *Decoder, start StartElement) error
}

type UnmarshalerAttr interface {
	UnmarshalXMLAttr(attr Attr) error
}

func Unmarshal(data []byte, v interface{}) error {
	panic("stub")
}

func (d *Decoder) Decode(v interface{}) error {
	panic("stub")
}

func (d *Decoder) DecodeElement(v interface{}, start *StartElement) error {
	panic("stub")
}

func (e UnmarshalError) Error() string {
	panic("stub")
}

func (d *Decoder) Skip() error {
	panic("stub")
}

type typeInfo struct {
	xmlname *fieldInfo
	fields  []fieldInfo
}

type fieldInfo struct {
	idx     []int
	name    string
	xmlns   string
	flags   fieldFlags
	parents []string
}

type TagPathError struct {
	Struct reflect.Type
	Field1
	Tag1 string
	Field2
	Tag2 string
}

type fieldFlags int

func (e *TagPathError) Error() string {
	panic("stub")
}

type stack struct {
	next *stack
	kind int
	name Name
	ok   bool
}

type SyntaxError struct {
	Msg  string
	Line int
}

type Token interface {
}

type EndElement struct{ Name Name }

type Directive []byte

type Comment []byte

type Name struct {
	Space
	Local string
}

type CharData []byte

type ProcInst struct {
	Target string
	Inst   []byte
}

type Attr struct {
	Name  Name
	Value string
}

type StartElement struct {
	Name Name
	Attr []Attr
}

type Decoder struct {
	Strict         bool
	AutoClose      []string
	Entity         map[string]string
	CharsetReader  func(charset string, input io.Reader) (io.Reader, error)
	DefaultSpace   string
	r              io.ByteReader
	buf            bytes.Buffer
	saved          *bytes.Buffer
	stk            *stack
	free           *stack
	needClose      bool
	toClose        Name
	nextToken      Token
	nextByte       int
	ns             map[string]string
	err            error
	line           int
	offset         int64
	unmarshalDepth int
}

func (e *SyntaxError) Error() string {
	panic("stub")
}

func (e StartElement) Copy() StartElement {
	panic("stub")
}

func (e StartElement) End() EndElement {
	panic("stub")
}

func (c CharData) Copy() CharData {
	panic("stub")
}

func (c Comment) Copy() Comment {
	panic("stub")
}

func (p ProcInst) Copy() ProcInst {
	panic("stub")
}

func (d Directive) Copy() Directive {
	panic("stub")
}

func CopyToken(t Token) Token {
	panic("stub")
}

func NewDecoder(r io.Reader) *Decoder {
	panic("stub")
}

func (d *Decoder) Token() (t Token, err error) {
	panic("stub")
}

func (d *Decoder) RawToken() (Token, error) {
	panic("stub")
}

func (d *Decoder) InputOffset() int64 {
	panic("stub")
}

func EscapeText(w io.Writer, s []byte) error {
	panic("stub")
}

func (p *printer) EscapeString(s string) {
	panic("stub")
}

func Escape(w io.Writer, s []byte) {
	panic("stub")
}

type Embedme interface{}
