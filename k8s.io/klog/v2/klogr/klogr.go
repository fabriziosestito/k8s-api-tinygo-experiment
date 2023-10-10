package klogr

type Format string

type klogger struct {
	level     int
	callDepth int
	prefix    string
	values    []interface{}
	format    Format
}

type Option func(*klogger)

func WithFormat(format Format) Option {
	panic("stub")
}

func New() interface{} {
	panic("stub")
}

func NewWithOptions(options ...Option) interface{} {
	panic("stub")
}

func (l *klogger) Init(info interface{}) {
	panic("stub")
}

func (l *klogger) Info(level int, msg string, kvList ...interface{}) {
	panic("stub")
}

func (l *klogger) Enabled(level int) bool {
	panic("stub")
}

func (l *klogger) Error(err error, msg string, kvList ...interface{}) {
	panic("stub")
}

func (l klogger) WithName(name string) interface{} {
	panic("stub")
}

func (l klogger) WithValues(kvList ...interface{}) interface{} {
	panic("stub")
}

func (l klogger) WithCallDepth(depth int) interface{} {
	panic("stub")
}

type Embedme interface{}
