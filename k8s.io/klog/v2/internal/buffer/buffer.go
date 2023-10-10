package buffer

import (
	"bytes"
	"time"

	"k8s.io/klog/v2/internal/severity"
)

type Buffer struct {
	bytes.Buffer
	Tmp  [64]byte
	next *Buffer
}

func GetBuffer() *Buffer {
	panic("stub")
}

func PutBuffer(b *Buffer) {
	panic("stub")
}

func (buf *Buffer) FormatHeader(s severity.Severity, file string, line int, now time.Time) {
	panic("stub")
}

func (buf *Buffer) SprintHeader(s severity.Severity, now time.Time) string {
	panic("stub")
}

type Embedme interface{}
