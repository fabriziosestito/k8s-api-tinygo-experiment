package html

import (
	"io"

	"golang.org/x/net/html/atom"
)

func EscapeString(s string) string {
	panic("stub")
}

func UnescapeString(s string) string {
	panic("stub")
}

type NodeType uint32

type nodeStack []*Node

type insertionModeStack []insertionMode

type Node struct {
	Parent
	FirstChild
	LastChild
	PrevSibling
	NextSibling *Node
	Type        NodeType
	DataAtom    atom.Atom
	Data        string
	Namespace   string
	Attr        []Attribute
}

func (n *Node) InsertBefore(newChild, oldChild *Node) {
	panic("stub")
}

func (n *Node) AppendChild(c *Node) {
	panic("stub")
}

func (n *Node) RemoveChild(c *Node) {
	panic("stub")
}

type ParseOption func(p *parser)

type parser struct {
	tokenizer           *Tokenizer
	tok                 Token
	hasSelfClosingToken bool
	doc                 *Node
	oe
	afe nodeStack
	head
	form *Node
	scripting
	framesetOK      bool
	templateStack   insertionModeStack
	im              insertionMode
	originalIM      insertionMode
	fosterParenting bool
	quirks          bool
	fragment        bool
	context         *Node
}

type insertionMode func(*parser) bool

type scope int

func Parse(r io.Reader) (*Node, error) {
	panic("stub")
}

func ParseFragment(r io.Reader, context *Node) ([]*Node, error) {
	panic("stub")
}

func ParseOptionEnableScripting(enable bool) ParseOption {
	panic("stub")
}

func ParseWithOptions(r io.Reader, opts ...ParseOption) (*Node, error) {
	panic("stub")
}

func ParseFragmentWithOptions(r io.Reader, context *Node, opts ...ParseOption) ([]*Node, error) {
	panic("stub")
}

type writer interface {
	WriteString(string) (int, error)
}

func Render(w io.Writer, n *Node) error {
	panic("stub")
}

type TokenType uint32

type Tokenizer struct {
	r             io.Reader
	tt            TokenType
	err           error
	readErr       error
	raw           span
	buf           []byte
	maxBuf        int
	data          span
	pendingAttr   [2]span
	attr          [][2]span
	nAttrReturned int
	rawTag        string
	textIsRaw     bool
	convertNUL    bool
	allowCDATA    bool
}

type Token struct {
	Type     TokenType
	DataAtom atom.Atom
	Data     string
	Attr     []Attribute
}

type span struct {
	start
	end int
}

type Attribute struct {
	Namespace
	Key
	Val string
}

func (t TokenType) String() string {
	panic("stub")
}

func (t Token) String() string {
	panic("stub")
}

func (z *Tokenizer) AllowCDATA(allowCDATA bool) {
	panic("stub")
}

func (z *Tokenizer) NextIsNotRawText() {
	panic("stub")
}

func (z *Tokenizer) Err() error {
	panic("stub")
}

func (z *Tokenizer) Buffered() []byte {
	panic("stub")
}

func (z *Tokenizer) Next() TokenType {
	panic("stub")
}

func (z *Tokenizer) Raw() []byte {
	panic("stub")
}

func (z *Tokenizer) Text() []byte {
	panic("stub")
}

func (z *Tokenizer) TagName() (name []byte, hasAttr bool) {
	panic("stub")
}

func (z *Tokenizer) TagAttr() (key, val []byte, moreAttr bool) {
	panic("stub")
}

func (z *Tokenizer) Token() Token {
	panic("stub")
}

func (z *Tokenizer) SetMaxBuf(n int) {
	panic("stub")
}

func NewTokenizer(r io.Reader) *Tokenizer {
	panic("stub")
}

func NewTokenizerFragment(r io.Reader, contextTag string) *Tokenizer {
	panic("stub")
}

type Embedme interface{}
