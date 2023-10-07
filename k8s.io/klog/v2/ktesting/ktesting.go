package ktesting

import (
	"context"
	"flag"
	"strings"
	"sync"
	"time"

	"k8s.io/klog/v2/internal/serialize"
	"k8s.io/klog/v2/internal/verbosity"
)

type configOptions struct {
	anyToString       serialize.AnyToStringFunc
	verbosityFlagName string
	vmoduleFlagName   string
	verbosityDefault  int
	bufferLogs        bool
}

type Config struct {
	vstate *verbosity.VState
	co     configOptions
}

type ConfigOption func(co *configOptions)

func (c *Config) Verbosity() flag.Value {
	panic("stub")
}

func (c *Config) VModule() flag.Value {
	panic("stub")
}

func AnyToString(anyToString func(value interface{}) string) ConfigOption {
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

func BufferLogs(enabled bool) ConfigOption {
	panic("stub")
}

func NewConfig(opts ...ConfigOption) *Config {
	panic("stub")
}

func (c *Config) AddFlags(fs *flag.FlagSet) {
	panic("stub")
}

func NewTestContext(tl TL) (interface{}, context.Context) {
	panic("stub")
}

type BufferTL struct{ strings.Builder }

type Log []LogEntry

type LogType string

type Underlier interface {
	GetUnderlying() TL
	GetBuffer() Buffer
}

type tloggerShared struct {
	mutex                sync.Mutex
	t                    TL
	goroutineWarningDone bool
	formatter            serialize.Formatter
	testName             string
	config               *Config
	buffer               logBuffer
	callDepth            int
}

type TL interface {
	Helper()
	Log(args ...interface{})
}

type logBuffer struct {
	mutex sync.Mutex
	text  strings.Builder
	log   Log
}

type tlogger struct {
	shared *tloggerShared
	prefix string
	values []interface{}
}

type NopTL struct{}

type Buffer interface {
	String() string
	Data() Log
}

type LogEntry struct {
	Timestamp       time.Time
	Type            LogType
	Prefix          string
	Message         string
	Verbosity       int
	Err             error
	WithKVList      []interface{}
	ParameterKVList []interface{}
}

func (n NopTL) Helper() {
	panic("stub")
}

func (n NopTL) Log(args ...interface{}) {
	panic("stub")
}

func (n *BufferTL) Helper() {
	panic("stub")
}

func (n *BufferTL) Log(args ...interface{}) {
	panic("stub")
}

func NewLogger(t TL, c *Config) interface{} {
	panic("stub")
}

func (l Log) DeepCopy() Log {
	panic("stub")
}

func (b *logBuffer) String() string {
	panic("stub")
}

func (b *logBuffer) Data() Log {
	panic("stub")
}

func (l tlogger) Init(info interface{}) {
	panic("stub")
}

func (l tlogger) GetCallStackHelper() func() {
	panic("stub")
}

func (l tlogger) Info(level int, msg string, kvList ...interface{}) {
	panic("stub")
}

func (l tlogger) Enabled(level int) bool {
	panic("stub")
}

func (l tlogger) Error(err error, msg string, kvList ...interface{}) {
	panic("stub")
}

func (l tlogger) WithName(name string) interface{} {
	panic("stub")
}

func (l tlogger) WithValues(kvList ...interface{}) interface{} {
	panic("stub")
}

func (l tlogger) GetUnderlying() TL {
	panic("stub")
}

func (l tlogger) GetBuffer() Buffer {
	panic("stub")
}

type Embedme interface{}
