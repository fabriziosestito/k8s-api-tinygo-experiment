package test

import (
	"flag"
	"io"
	"testing"
)

type faultyStringer struct{}

type faultyMarshaler struct{}

type recursiveMarshaler struct{}

type stringer struct{ s string }

type OutputConfig struct {
	NewLogger             func(out io.Writer, v int, vmodule string) interface{}
	AsBackend             bool
	ExpectedOutputMapping map[string]string
	SupportsVModule       bool
}

type testcase struct {
	withHelper     bool
	withNames      []string
	withValues     []interface{}
	moreValues     []interface{}
	evenMoreValues []interface{}
	v              int
	vmodule        string
	text           string
	values         []interface{}
	err            error
	expectedOutput string
}

type kmeta struct {
	Name
	Namespace string
}

type customErrorJSON struct{ s string }

type faultyError struct{}

type typeMeta struct{ Kind string }

type myConfig struct {
	typeMeta
	RealField int
}

type myList struct {
	Value int
	Next  *myList
}

func InitKlog(tb testing.TB) *flag.FlagSet {
	panic("stub")
}

func Output(t *testing.T, config OutputConfig) {
	panic("stub")
}

func Benchmark(b *testing.B, config OutputConfig) {
	panic("stub")
}

func (k kmeta) GetName() string {
	panic("stub")
}

func (k kmeta) GetNamespace() string {
	panic("stub")
}

func (e *customErrorJSON) Error() string {
	panic("stub")
}

func (e *customErrorJSON) MarshalJSON() ([]byte, error) {
	panic("stub")
}

func (s *stringer) String() string {
	panic("stub")
}

func (f faultyStringer) String() string {
	panic("stub")
}

func (f faultyMarshaler) MarshalLog() interface{} {
	panic("stub")
}

func (r recursiveMarshaler) MarshalLog() interface{} {
	panic("stub")
}

func (f faultyError) Error() string {
	panic("stub")
}

func (t typeMeta) String() string {
	panic("stub")
}

func (t typeMeta) MarshalLog() interface{} {
	panic("stub")
}

func ZaprOutputMappingDirect() map[string]string {
	panic("stub")
}

func ZaprOutputMappingIndirect() map[string]string {
	panic("stub")
}

type Embedme interface{}
