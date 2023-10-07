package verbosity

import "sync"

type Value interface {
	Type() string
}

type Level int32

type levelSpec struct {
	vs *VState
	l  Level
}

type VState struct {
	mu           sync.Mutex
	vmodule      moduleSpec
	verbosity    levelSpec
	pcs          [1]uintptr
	vmap         map[uintptr]Level
	filterLength int32
}

type moduleSpec struct {
	vs     *VState
	filter []modulePat
}

type modulePat struct {
	pattern string
	literal bool
	level   Level
}

func New() *VState {
	panic("stub")
}

func (vs *VState) V() Value {
	panic("stub")
}

func (vs *VState) VModule() Value {
	panic("stub")
}

func (l *levelSpec) String() string {
	panic("stub")
}

func (l *levelSpec) Get() interface{} {
	panic("stub")
}

func (l *levelSpec) Type() string {
	panic("stub")
}

func (l *levelSpec) Set(value string) error {
	panic("stub")
}

func (m *moduleSpec) String() string {
	panic("stub")
}

func (m *moduleSpec) Get() interface{} {
	panic("stub")
}

func (m *moduleSpec) Type() string {
	panic("stub")
}

func (m *moduleSpec) Set(value string) error {
	panic("stub")
}

func (vs *VState) Enabled(level Level, depth int) bool {
	panic("stub")
}

type Embedme interface{}
