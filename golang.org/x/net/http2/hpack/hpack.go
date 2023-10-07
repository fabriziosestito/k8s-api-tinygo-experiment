package hpack

import (
	"bytes"
	"io"
)

type Encoder struct {
	dynTab          dynamicTable
	minSize         uint32
	maxSizeLimit    uint32
	tableSizeUpdate bool
	w               io.Writer
	buf             []byte
}

func NewEncoder(w io.Writer) *Encoder {
	panic("stub")
}

func (e *Encoder) WriteField(f HeaderField) error {
	panic("stub")
}

func (e *Encoder) SetMaxDynamicTableSize(v uint32) {
	panic("stub")
}

func (e *Encoder) MaxDynamicTableSize() (v uint32) {
	panic("stub")
}

func (e *Encoder) SetMaxDynamicTableSizeLimit(v uint32) {
	panic("stub")
}

type HeaderField struct {
	Name
	Value     string
	Sensitive bool
}

type dynamicTable struct {
	table          headerFieldTable
	size           uint32
	maxSize        uint32
	allowedMaxSize uint32
}

type indexType int

type DecodingError struct{ Err error }

type Decoder struct {
	dynTab      dynamicTable
	emit        func(f HeaderField)
	emitEnabled bool
	maxStrLen   int
	buf         []byte
	saveBuf     bytes.Buffer
	firstField  bool
}

type undecodedString struct {
	isHuff bool
	b      []byte
}

type InvalidIndexError int

func (de DecodingError) Error() string {
	panic("stub")
}

func (e InvalidIndexError) Error() string {
	panic("stub")
}

func (hf HeaderField) IsPseudo() bool {
	panic("stub")
}

func (hf HeaderField) String() string {
	panic("stub")
}

func (hf HeaderField) Size() uint32 {
	panic("stub")
}

func NewDecoder(maxDynamicTableSize uint32, emitFunc func(f HeaderField)) *Decoder {
	panic("stub")
}

func (d *Decoder) SetMaxStringLength(n int) {
	panic("stub")
}

func (d *Decoder) SetEmitFunc(emitFunc func(f HeaderField)) {
	panic("stub")
}

func (d *Decoder) SetEmitEnabled(v bool) {
	panic("stub")
}

func (d *Decoder) EmitEnabled() bool {
	panic("stub")
}

func (d *Decoder) SetMaxDynamicTableSize(v uint32) {
	panic("stub")
}

func (d *Decoder) SetAllowedMaxDynamicTableSize(v uint32) {
	panic("stub")
}

func (d *Decoder) DecodeFull(p []byte) ([]HeaderField, error) {
	panic("stub")
}

func (d *Decoder) Close() error {
	panic("stub")
}

func (d *Decoder) Write(p []byte) (n int, err error) {
	panic("stub")
}

type incomparable [0]func()

type node struct {
	_        incomparable
	children *[256]*node
	codeLen  uint8
	sym      byte
}

func HuffmanDecode(w io.Writer, v []byte) (int, error) {
	panic("stub")
}

func HuffmanDecodeToString(v []byte) (string, error) {
	panic("stub")
}

func AppendHuffmanString(dst []byte, s string) []byte {
	panic("stub")
}

func HuffmanEncodeLength(s string) uint64 {
	panic("stub")
}

type pairNameValue struct {
	name
	value string
}

type headerFieldTable struct {
	ents        []HeaderField
	evictCount  uint64
	byName      map[string]uint64
	byNameValue map[pairNameValue]uint64
}

type Embedme interface{}
