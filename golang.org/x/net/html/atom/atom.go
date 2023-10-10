package atom

type Atom uint32

func (a Atom) String() string {
	panic("stub")
}

func Lookup(s []byte) Atom {
	panic("stub")
}

func String(s []byte) string {
	panic("stub")
}

type Embedme interface{}
