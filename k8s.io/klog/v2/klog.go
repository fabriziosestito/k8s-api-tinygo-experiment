package klog

import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"io"
	"os"
	"time"

	stdLog "log"
	"sync"

	"k8s.io/klog/v2/internal/clock"
	"k8s.io/klog/v2/internal/severity"
)

type loggerOptions struct {
	contextualLogger bool
	flush            func()
	writeKlogBuffer  func([]byte)
}

type logWriter struct {
	Logger
	writeKlogBuffer func([]byte)
}

type LoggerOption func(o *loggerOptions)

func SetLogger(logger interface{}) {
	panic("stub")
}

func SetLoggerWithOptions(logger interface{}, opts ...LoggerOption) {
	panic("stub")
}

func ContextualLogger(enabled bool) LoggerOption {
	panic("stub")
}

func FlushLogger(flush func()) LoggerOption {
	panic("stub")
}

func WriteKlogBuffer(write func([]byte)) LoggerOption {
	panic("stub")
}

func ClearLogger() {
	panic("stub")
}

func EnableContextualLogging(enabled bool) {
	panic("stub")
}

func FromContext(ctx context.Context) Logger {
	panic("stub")
}

func TODO() Logger {
	panic("stub")
}

func Background() Logger {
	panic("stub")
}

func LoggerWithValues(logger Logger, kv ...interface{}) Logger {
	panic("stub")
}

func LoggerWithName(logger Logger, name string) Logger {
	panic("stub")
}

func NewContext(ctx context.Context, logger Logger) context.Context {
	panic("stub")
}

func FlushAndExit(flushTimeout time.Duration, exitCode int) {
	panic("stub")
}

type formatAny struct{ Object interface{} }

func Format(obj interface{}) interface{} {
	panic("stub")
}

func (f formatAny) String() string {
	panic("stub")
}

func (f formatAny) MarshalLog() interface{} {
	panic("stub")
}

type Logger interface{}

type LogSink interface{}

type RuntimeInfo interface{}

type kobjSlice struct{ arg interface{} }

type ObjectRef struct {
	Name      string
	Namespace string
}

type KMetadata interface {
	GetName() string
	GetNamespace() string
}

func (ref ObjectRef) String() string {
	panic("stub")
}

func (ref ObjectRef) WriteText(out *bytes.Buffer) {
	panic("stub")
}

func (ref ObjectRef) MarshalLog() interface{} {
	panic("stub")
}

func KObj(obj KMetadata) ObjectRef {
	panic("stub")
}

func KRef(namespace, name string) ObjectRef {
	panic("stub")
}

func KObjs(arg interface{}) []ObjectRef {
	panic("stub")
}

func KObjSlice(arg interface{}) interface{} {
	panic("stub")
}

func (ks kobjSlice) String() string {
	panic("stub")
}

func (ks kobjSlice) MarshalLog() interface{} {
	panic("stub")
}

func (ks kobjSlice) WriteText(out *bytes.Buffer) {
	panic("stub")
}

type OutputStats struct {
	lines int64
	bytes int64
}

type Level int32

type settings struct {
	contextualLoggingEnabled bool
	logger                   *logWriter
	loggerOptions            loggerOptions
	toStderr                 bool
	alsoToStderr             bool
	stderrThreshold          severityValue
	file                     [severity.NumSeverity]flushSyncWriter
	flushInterval            time.Duration
	filterLength             int32
	traceLocation            traceLocation
	vmodule                  moduleSpec
	verbosity                Level
	logDir                   string
	logFile                  string
	logFileMaxSizeMB         uint64
	skipHeaders              bool
	skipLogHeaders           bool
	addDirHeader             bool
	oneOutput                bool
	filter                   LogFilter
}

type State interface {
	Restore()
}

type Verbose struct {
	enabled bool
	logger  *logWriter
}

type modulePat struct {
	pattern string
	literal bool
	level   Level
}

type logBridge severity.Severity

type severityValue struct{ severity.Severity }

type flushSyncWriter interface {
	Flush() error
	Sync() error
}

type LogFilter interface {
	Filter(args []interface{}) []interface{}
	FilterF(format string, args []interface{}) (string, []interface{})
	FilterS(msg string, keysAndValues []interface{}) (string, []interface{})
}

type loggingT struct {
	settings
	flushD *flushDaemon
	mu     sync.Mutex
	pcs    [1]uintptr
	vmap   map[uintptr]Level
}

type redirectBuffer struct{ w io.Writer }

type syncBuffer struct {
	logger *loggingT
	*bufio.Writer
	file     *os.File
	sev      severity.Severity
	nbytes   uint64
	maxbytes uint64
}

type moduleSpec struct{ filter []modulePat }

type state struct {
	settings
	flushDRunning bool
	maxSize       uint64
}

type traceLocation struct {
	file string
	line int
}

type flushDaemon struct {
	mu       sync.Mutex
	clock    clock.WithTicker
	flush    func()
	stopC    chan struct{}
	stopDone chan struct{}
}

func (s *severityValue) String() string {
	panic("stub")
}

func (s *severityValue) Get() interface{} {
	panic("stub")
}

func (s *severityValue) Set(value string) error {
	panic("stub")
}

func (s *OutputStats) Lines() int64 {
	panic("stub")
}

func (s *OutputStats) Bytes() int64 {
	panic("stub")
}

func (l *Level) String() string {
	panic("stub")
}

func (l *Level) Get() interface{} {
	panic("stub")
}

func (l *Level) Set(value string) error {
	panic("stub")
}

func (m *moduleSpec) String() string {
	panic("stub")
}

func (m *moduleSpec) Get() interface{} {
	panic("stub")
}

func (m *moduleSpec) Set(value string) error {
	panic("stub")
}

func (t *traceLocation) String() string {
	panic("stub")
}

func (t *traceLocation) Get() interface{} {
	panic("stub")
}

func (t *traceLocation) Set(value string) error {
	panic("stub")
}

func InitFlags(flagset *flag.FlagSet) {
	panic("stub")
}

func Flush() {
	panic("stub")
}

func CaptureState() State {
	panic("stub")
}

func (s *state) Restore() {
	panic("stub")
}

func (rb *redirectBuffer) Sync() error {
	panic("stub")
}

func (rb *redirectBuffer) Flush() error {
	panic("stub")
}

func (rb *redirectBuffer) Write(bytes []byte) (n int, err error) {
	panic("stub")
}

func SetOutput(w io.Writer) {
	panic("stub")
}

func SetOutputBySeverity(name string, w io.Writer) {
	panic("stub")
}

func LogToStderr(stderr bool) {
	panic("stub")
}

func (sb *syncBuffer) Sync() error {
	panic("stub")
}

func CalculateMaxSize() uint64 {
	panic("stub")
}

func (sb *syncBuffer) Write(p []byte) (n int, err error) {
	panic("stub")
}

func StopFlushDaemon() {
	panic("stub")
}

func StartFlushDaemon(interval time.Duration) {
	panic("stub")
}

func CopyStandardLogTo(name string) {
	panic("stub")
}

func NewStandardLogger(name string) *stdLog.Logger {
	panic("stub")
}

func (lb logBridge) Write(b []byte) (n int, err error) {
	panic("stub")
}

func V(level Level) Verbose {
	panic("stub")
}

func VDepth(depth int, level Level) Verbose {
	panic("stub")
}

func (v Verbose) Enabled() bool {
	panic("stub")
}

func (v Verbose) Info(args ...interface{}) {
	panic("stub")
}

func (v Verbose) InfoDepth(depth int, args ...interface{}) {
	panic("stub")
}

func (v Verbose) Infoln(args ...interface{}) {
	panic("stub")
}

func (v Verbose) InfolnDepth(depth int, args ...interface{}) {
	panic("stub")
}

func (v Verbose) Infof(format string, args ...interface{}) {
	panic("stub")
}

func (v Verbose) InfofDepth(depth int, format string, args ...interface{}) {
	panic("stub")
}

func (v Verbose) InfoS(msg string, keysAndValues ...interface{}) {
	panic("stub")
}

func InfoSDepth(depth int, msg string, keysAndValues ...interface{}) {
	panic("stub")
}

func (v Verbose) InfoSDepth(depth int, msg string, keysAndValues ...interface{}) {
	panic("stub")
}

func (v Verbose) Error(err error, msg string, args ...interface{}) {
	panic("stub")
}

func (v Verbose) ErrorS(err error, msg string, keysAndValues ...interface{}) {
	panic("stub")
}

func Info(args ...interface{}) {
	panic("stub")
}

func InfoDepth(depth int, args ...interface{}) {
	panic("stub")
}

func Infoln(args ...interface{}) {
	panic("stub")
}

func InfolnDepth(depth int, args ...interface{}) {
	panic("stub")
}

func Infof(format string, args ...interface{}) {
	panic("stub")
}

func InfofDepth(depth int, format string, args ...interface{}) {
	panic("stub")
}

func InfoS(msg string, keysAndValues ...interface{}) {
	panic("stub")
}

func Warning(args ...interface{}) {
	panic("stub")
}

func WarningDepth(depth int, args ...interface{}) {
	panic("stub")
}

func Warningln(args ...interface{}) {
	panic("stub")
}

func WarninglnDepth(depth int, args ...interface{}) {
	panic("stub")
}

func Warningf(format string, args ...interface{}) {
	panic("stub")
}

func WarningfDepth(depth int, format string, args ...interface{}) {
	panic("stub")
}

func Error(args ...interface{}) {
	panic("stub")
}

func ErrorDepth(depth int, args ...interface{}) {
	panic("stub")
}

func Errorln(args ...interface{}) {
	panic("stub")
}

func ErrorlnDepth(depth int, args ...interface{}) {
	panic("stub")
}

func Errorf(format string, args ...interface{}) {
	panic("stub")
}

func ErrorfDepth(depth int, format string, args ...interface{}) {
	panic("stub")
}

func ErrorS(err error, msg string, keysAndValues ...interface{}) {
	panic("stub")
}

func ErrorSDepth(depth int, err error, msg string, keysAndValues ...interface{}) {
	panic("stub")
}

func Fatal(args ...interface{}) {
	panic("stub")
}

func FatalDepth(depth int, args ...interface{}) {
	panic("stub")
}

func Fatalln(args ...interface{}) {
	panic("stub")
}

func FatallnDepth(depth int, args ...interface{}) {
	panic("stub")
}

func Fatalf(format string, args ...interface{}) {
	panic("stub")
}

func FatalfDepth(depth int, format string, args ...interface{}) {
	panic("stub")
}

func Exit(args ...interface{}) {
	panic("stub")
}

func ExitDepth(depth int, args ...interface{}) {
	panic("stub")
}

func Exitln(args ...interface{}) {
	panic("stub")
}

func ExitlnDepth(depth int, args ...interface{}) {
	panic("stub")
}

func Exitf(format string, args ...interface{}) {
	panic("stub")
}

func ExitfDepth(depth int, format string, args ...interface{}) {
	panic("stub")
}

func SetLogFilter(filter LogFilter) {
	panic("stub")
}

type klogger struct {
	level     int
	callDepth int
	prefix    string
	values    []interface{}
}

func NewKlogr() Logger {
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
