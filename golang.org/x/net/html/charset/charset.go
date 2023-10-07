package charset

import "io"

type htmlEncoding struct{ Embedme }

func Lookup(label string) (e interface{}, name string) {
	panic("stub")
}

func (h *htmlEncoding) NewEncoder() interface{} {
	panic("stub")
}

func DetermineEncoding(content []byte, contentType string) (e interface{}, name string, certain bool) {
	panic("stub")
}

func NewReader(r io.Reader, contentType string) (io.Reader, error) {
	panic("stub")
}

func NewReaderLabel(label string, input io.Reader) (io.Reader, error) {
	panic("stub")
}

type Embedme interface{}
