package textlogger

import (
	"flag"
	"io"

	"k8s.io/klog/v2/internal/verbosity"
)

type Config struct {
	vstate *verbosity.VState
	co     configOptions
}

type ConfigOption func(co *configOptions)

type configOptions struct {
	verbosityFlagName string
	vmoduleFlagName   string
	verbosityDefault  int
	output            io.Writer
}

func (c *Config) Verbosity() flag.Value {
	panic("stub")
}

func (c *Config) VModule() flag.Value {
	panic("stub")
}

func VerbosityFlagName(name string) ConfigOption {
	panic("stub")
}

func VModuleFlagName(name string) ConfigOption {
	panic("stub")
}

func Verbosity(level int) ConfigOption {
	panic("stub")
}

func Output(output io.Writer) ConfigOption {
	panic("stub")
}

func NewConfig(opts ...ConfigOption) *Config {
	panic("stub")
}

func (c *Config) AddFlags(fs *flag.FlagSet) {
	panic("stub")
}

type tlogger struct {
	callDepth int
	prefix    string
	values    []interface{}
	config    *Config
}

type KlogBufferWriter interface {
	WriteKlogBuffer([]byte)
}

func NewLogger(c *Config) interface{} {
	panic("stub")
}

func (l *tlogger) Init(info interface{}) {
	panic("stub")
}

func (l *tlogger) WithCallDepth(depth int) interface{} {
	panic("stub")
}

func (l *tlogger) Enabled(level int) bool {
	panic("stub")
}

func (l *tlogger) Info(level int, msg string, kvList ...interface{}) {
	panic("stub")
}

func (l *tlogger) Error(err error, msg string, kvList ...interface{}) {
	panic("stub")
}

func (l *tlogger) WriteKlogBuffer(data []byte) {
	panic("stub")
}

func (l *tlogger) WithName(name string) interface{} {
	panic("stub")
}

func (l *tlogger) WithValues(kvList ...interface{}) interface{} {
	panic("stub")
}

type Embedme interface{}
