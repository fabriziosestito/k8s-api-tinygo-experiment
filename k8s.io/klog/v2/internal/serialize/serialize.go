package serialize

import (
	"bytes"
	"fmt"
)

type textWriter interface {
	WriteText(*bytes.Buffer)
}

type Formatter struct{ AnyToStringHook AnyToStringFunc }

type AnyToStringFunc func(v interface{}) string

func WithValues(oldKV, newKV []interface{}) []interface{} {
	panic("stub")
}

func MergeKVs(first, second []interface{}) []interface{} {
	panic("stub")
}

func (f Formatter) MergeAndFormatKVs(b *bytes.Buffer, first, second []interface{}) {
	panic("stub")
}

func MergeAndFormatKVs(b *bytes.Buffer, first, second []interface{}) {
	panic("stub")
}

func (f Formatter) KVListFormat(b *bytes.Buffer, keysAndValues ...interface{}) {
	panic("stub")
}

func KVListFormat(b *bytes.Buffer, keysAndValues ...interface{}) {
	panic("stub")
}

func (f Formatter) KVFormat(b *bytes.Buffer, k, v interface{}) {
	panic("stub")
}

func KVFormat(b *bytes.Buffer, k, v interface{}) {
	panic("stub")
}

func StringerToString(s fmt.Stringer) (ret string) {
	panic("stub")
}

func MarshalerToValue(m interface{}) (ret interface{}) {
	panic("stub")
}

func ErrorToString(err error) (ret string) {
	panic("stub")
}

type Embedme interface{}
