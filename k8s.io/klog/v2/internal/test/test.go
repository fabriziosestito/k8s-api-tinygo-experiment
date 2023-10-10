package test

type KMetadataMock struct {
	Name
	NS string
}

type PtrKMetadataMock struct {
	Name
	NS string
}

func (m KMetadataMock) GetName() string {
	panic("stub")
}

func (m KMetadataMock) GetNamespace() string {
	panic("stub")
}

func (m *PtrKMetadataMock) GetName() string {
	panic("stub")
}

func (m *PtrKMetadataMock) GetNamespace() string {
	panic("stub")
}

type Embedme interface{}
