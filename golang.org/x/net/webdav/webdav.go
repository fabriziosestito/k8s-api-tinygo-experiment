package webdav

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	ixml "golang.org/x/net/webdav/internal/xml"
)

type File interface {
}

type memFile struct {
	n                *memFSNode
	nameSnapshot     string
	childrenSnapshot []os.FileInfo
	pos              int
}

type memFSNode struct {
	children  map[string]*memFSNode
	mu        sync.Mutex
	data      []byte
	mode      os.FileMode
	modTime   time.Time
	deadProps map[xml.Name]Property
}

type memFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

type FileSystem interface {
	Mkdir(ctx context.Context, name string, perm os.FileMode) error
	OpenFile(ctx context.Context, name string, flag int, perm os.FileMode) (File, error)
	RemoveAll(ctx context.Context, name string) error
	Rename(ctx context.Context, oldName, newName string) error
	Stat(ctx context.Context, name string) (os.FileInfo, error)
}

type Dir string

type memFS struct {
	mu   sync.Mutex
	root memFSNode
}

func (d Dir) Mkdir(ctx context.Context, name string, perm os.FileMode) error {
	panic("stub")
}

func (d Dir) OpenFile(ctx context.Context, name string, flag int, perm os.FileMode) (File, error) {
	panic("stub")
}

func (d Dir) RemoveAll(ctx context.Context, name string) error {
	panic("stub")
}

func (d Dir) Rename(ctx context.Context, oldName, newName string) error {
	panic("stub")
}

func (d Dir) Stat(ctx context.Context, name string) (os.FileInfo, error) {
	panic("stub")
}

func NewMemFS() FileSystem {
	panic("stub")
}

func (fs *memFS) Mkdir(ctx context.Context, name string, perm os.FileMode) error {
	panic("stub")
}

func (fs *memFS) OpenFile(ctx context.Context, name string, flag int, perm os.FileMode) (File, error) {
	panic("stub")
}

func (fs *memFS) RemoveAll(ctx context.Context, name string) error {
	panic("stub")
}

func (fs *memFS) Rename(ctx context.Context, oldName, newName string) error {
	panic("stub")
}

func (fs *memFS) Stat(ctx context.Context, name string) (os.FileInfo, error) {
	panic("stub")
}

func (n *memFSNode) DeadProps() (map[xml.Name]Property, error) {
	panic("stub")
}

func (n *memFSNode) Patch(patches []Proppatch) ([]Propstat, error) {
	panic("stub")
}

func (f *memFileInfo) Name() string {
	panic("stub")
}

func (f *memFileInfo) Size() int64 {
	panic("stub")
}

func (f *memFileInfo) Mode() os.FileMode {
	panic("stub")
}

func (f *memFileInfo) ModTime() time.Time {
	panic("stub")
}

func (f *memFileInfo) IsDir() bool {
	panic("stub")
}

func (f *memFileInfo) Sys() interface{} {
	panic("stub")
}

func (f *memFile) DeadProps() (map[xml.Name]Property, error) {
	panic("stub")
}

func (f *memFile) Patch(patches []Proppatch) ([]Propstat, error) {
	panic("stub")
}

func (f *memFile) Close() error {
	panic("stub")
}

func (f *memFile) Read(p []byte) (int, error) {
	panic("stub")
}

func (f *memFile) Readdir(count int) ([]os.FileInfo, error) {
	panic("stub")
}

func (f *memFile) Seek(offset int64, whence int) (int64, error) {
	panic("stub")
}

func (f *memFile) Stat() (os.FileInfo, error) {
	panic("stub")
}

func (f *memFile) Write(p []byte) (int, error) {
	panic("stub")
}

type ifList struct {
	resourceTag string
	conditions  []Condition
}

type ifHeader struct{ lists []ifList }

type memLS struct {
	mu       sync.Mutex
	byName   map[string]*memLSNode
	byToken  map[string]*memLSNode
	gen      uint64
	byExpiry byExpiry
}

type memLSNode struct {
	details       LockDetails
	token         string
	refCount      int
	expiry        time.Time
	byExpiryIndex int
	held          bool
}

type LockSystem interface {
	Confirm(now time.Time, name0, name1 string, conditions ...Condition) (release func(), err error)
	Create(now time.Time, details LockDetails) (token string, err error)
	Refresh(now time.Time, token string, duration time.Duration) (LockDetails, error)
	Unlock(now time.Time, token string) error
}

type Condition struct {
	Not   bool
	Token string
	ETag  string
}

type LockDetails struct {
	Root      string
	Duration  time.Duration
	OwnerXML  string
	ZeroDepth bool
}

type byExpiry []*memLSNode

func NewMemLS() LockSystem {
	panic("stub")
}

func (m *memLS) Confirm(now time.Time, name0, name1 string, conditions ...Condition) (func(), error) {
	panic("stub")
}

func (m *memLS) Create(now time.Time, details LockDetails) (string, error) {
	panic("stub")
}

func (m *memLS) Refresh(now time.Time, token string, duration time.Duration) (LockDetails, error) {
	panic("stub")
}

func (m *memLS) Unlock(now time.Time, token string) error {
	panic("stub")
}

func (b *byExpiry) Len() int {
	panic("stub")
}

func (b *byExpiry) Less(i, j int) bool {
	panic("stub")
}

func (b *byExpiry) Swap(i, j int) {
	panic("stub")
}

func (b *byExpiry) Push(x interface{}) {
	panic("stub")
}

func (b *byExpiry) Pop() interface{} {
	panic("stub")
}

type ContentTyper interface {
	ContentType(ctx context.Context) (string, error)
}

type ETager interface {
	ETag(ctx context.Context) (string, error)
}

type DeadPropsHolder interface {
	DeadProps() (map[xml.Name]Property, error)
	Patch([]Proppatch) ([]Propstat, error)
}

type Proppatch struct {
	Remove bool
	Props  []Property
}

type Propstat struct {
	Props               []Property
	Status              int
	XMLError            string
	ResponseDescription string
}

type Handler struct {
	Prefix     string
	FileSystem FileSystem
	LockSystem LockSystem
	Logger     func(*http.Request, error)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("stub")
}

func StatusText(code int) string {
	panic("stub")
}

type propfindProps []xml.Name

type Property struct {
	XMLName  xml.Name
	Lang     string
	InnerXML []byte
}

type ixmlProperty struct {
	XMLName  ixml.Name
	Lang     string
	InnerXML []byte
}

type multistatusWriter struct {
	responseDescription string
	w                   http.ResponseWriter
	enc                 *ixml.Encoder
}

type ixmlPropstat struct {
	Prop                []ixmlProperty
	Status              string
	Error               *xmlError
	ResponseDescription string
}

type propertyupdate struct {
	XMLName   ixml.Name
	Lang      string
	SetRemove []setRemove
}

type propstat struct {
	Prop                []Property
	Status              string
	Error               *xmlError
	ResponseDescription string
}

type response struct {
	XMLName             ixml.Name
	Href                []string
	Propstat            []propstat
	Status              string
	Error               *xmlError
	ResponseDescription string
}

type proppatchProps []Property

type lockInfo struct {
	XMLName   ixml.Name
	Exclusive *struct{}
	Shared    *struct{}
	Write     *struct{}
	Owner     owner
}

type propfind struct {
	XMLName  ixml.Name
	Allprop  *struct{}
	Propname *struct{}
	Prop     propfindProps
	Include  propfindProps
}

type xmlError struct {
	XMLName  ixml.Name
	InnerXML []byte
}

type setRemove struct {
	XMLName ixml.Name
	Lang    string
	Prop    proppatchProps
}

type owner struct{ InnerXML string }

type countingReader struct {
	n int
	r io.Reader
}

type xmlValue []byte

func (c *countingReader) Read(p []byte) (int, error) {
	panic("stub")
}

func (pn *propfindProps) UnmarshalXML(d *ixml.Decoder, start ixml.StartElement) error {
	panic("stub")
}

func (ps propstat) MarshalXML(e *ixml.Encoder, start ixml.StartElement) error {
	panic("stub")
}

func (v *xmlValue) UnmarshalXML(d *ixml.Decoder, start ixml.StartElement) error {
	panic("stub")
}

func (ps *proppatchProps) UnmarshalXML(d *ixml.Decoder, start ixml.StartElement) error {
	panic("stub")
}

type Embedme interface{}
