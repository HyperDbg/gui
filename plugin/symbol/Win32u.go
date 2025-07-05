package symbol

type (
	InterfaceWin32u any
	objectWin32u    struct{}
)

func NewGeneratedFileWin32u() InterfaceWin32u { return &objectWin32u{} }
