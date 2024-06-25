package symbol

type (
	InterfaceWin32u interface{}
	objectWin32u    struct{}
)

func NewGeneratedFileWin32u() InterfaceWin32u { return &objectWin32u{} }
