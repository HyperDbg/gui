package symbol

type (
	InterfaceNtdll interface{}
	objectNtdll    struct{}
)

func NewGeneratedFileNtdll() InterfaceNtdll { return &objectNtdll{} }
