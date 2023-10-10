package severity

type Severity int32

func ByName(s string) (Severity, bool) {
	panic("stub")
}

type Embedme interface{}

const (
	InfoLog Severity = iota
	WarningLog
	ErrorLog
	FatalLog
	NumSeverity = 4
)
