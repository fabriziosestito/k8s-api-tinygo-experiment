package bpf

func Assemble(insts []Instruction) ([]RawInstruction, error) {
	panic("stub")
}

func Disassemble(raw []RawInstruction) (insts []Instruction, allDecoded bool) {
	panic("stub")
}

type jumpOp uint16

type opOperand uint16

type Register uint16

type Extension int

type JumpTest uint16

type ALUOp uint16

type LoadScratch struct {
	Dst Register
	N   int
}

type LoadMemShift struct{ Off uint32 }

type Jump struct{ Skip uint32 }

type JumpIf struct {
	Cond      JumpTest
	Val       uint32
	SkipTrue  uint8
	SkipFalse uint8
}

type TAX struct{}

type RawInstruction struct {
	Op uint16
	Jt uint8
	Jf uint8
	K  uint32
}

type ALUOpConstant struct {
	Op  ALUOp
	Val uint32
}

type LoadIndirect struct {
	Off  uint32
	Size int
}

type LoadAbsolute struct {
	Off  uint32
	Size int
}

type LoadExtension struct{ Num Extension }

type NegateA struct{}

type RetConstant struct{ Val uint32 }

type TXA struct{}

type LoadConstant struct {
	Dst Register
	Val uint32
}

type StoreScratch struct {
	Src Register
	N   int
}

type ALUOpX struct{ Op ALUOp }

type JumpIfX struct {
	Cond      JumpTest
	SkipTrue  uint8
	SkipFalse uint8
}

type RetA struct{}

type Instruction interface {
	Assemble() (RawInstruction, error)
}

func (ri RawInstruction) Assemble() (RawInstruction, error) {
	panic("stub")
}

func (ri RawInstruction) Disassemble() Instruction {
	panic("stub")
}

func (a LoadConstant) Assemble() (RawInstruction, error) {
	panic("stub")
}

func (a LoadConstant) String() string {
	panic("stub")
}

func (a LoadScratch) Assemble() (RawInstruction, error) {
	panic("stub")
}

func (a LoadScratch) String() string {
	panic("stub")
}

func (a LoadAbsolute) Assemble() (RawInstruction, error) {
	panic("stub")
}

func (a LoadAbsolute) String() string {
	panic("stub")
}

func (a LoadIndirect) Assemble() (RawInstruction, error) {
	panic("stub")
}

func (a LoadIndirect) String() string {
	panic("stub")
}

func (a LoadMemShift) Assemble() (RawInstruction, error) {
	panic("stub")
}

func (a LoadMemShift) String() string {
	panic("stub")
}

func (a LoadExtension) Assemble() (RawInstruction, error) {
	panic("stub")
}

func (a LoadExtension) String() string {
	panic("stub")
}

func (a StoreScratch) Assemble() (RawInstruction, error) {
	panic("stub")
}

func (a StoreScratch) String() string {
	panic("stub")
}

func (a ALUOpConstant) Assemble() (RawInstruction, error) {
	panic("stub")
}

func (a ALUOpConstant) String() string {
	panic("stub")
}

func (a ALUOpX) Assemble() (RawInstruction, error) {
	panic("stub")
}

func (a ALUOpX) String() string {
	panic("stub")
}

func (a NegateA) Assemble() (RawInstruction, error) {
	panic("stub")
}

func (a NegateA) String() string {
	panic("stub")
}

func (a Jump) Assemble() (RawInstruction, error) {
	panic("stub")
}

func (a Jump) String() string {
	panic("stub")
}

func (a JumpIf) Assemble() (RawInstruction, error) {
	panic("stub")
}

func (a JumpIf) String() string {
	panic("stub")
}

func (a JumpIfX) Assemble() (RawInstruction, error) {
	panic("stub")
}

func (a JumpIfX) String() string {
	panic("stub")
}

func (a RetA) Assemble() (RawInstruction, error) {
	panic("stub")
}

func (a RetA) String() string {
	panic("stub")
}

func (a RetConstant) Assemble() (RawInstruction, error) {
	panic("stub")
}

func (a RetConstant) String() string {
	panic("stub")
}

func (a TXA) Assemble() (RawInstruction, error) {
	panic("stub")
}

func (a TXA) String() string {
	panic("stub")
}

func (a TAX) Assemble() (RawInstruction, error) {
	panic("stub")
}

func (a TAX) String() string {
	panic("stub")
}

type Setter interface {
	SetBPF(filter []RawInstruction) error
}

type VM struct{ filter []Instruction }

func NewVM(filter []Instruction) (*VM, error) {
	panic("stub")
}

func (v *VM) Run(in []byte) (int, error) {
	panic("stub")
}

type Embedme interface{}
