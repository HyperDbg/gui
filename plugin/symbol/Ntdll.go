package symbol

type (
	InterfaceNtdll any
	objectNtdll    struct{}
)

func NewGeneratedFileNtdll() InterfaceNtdll { return &objectNtdll{} }
