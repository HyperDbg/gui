package HyperDbgSdk

import (
	"github.com/ddkwork/hyperdbgui/SDK/Imports"
)

type (
	HyperDbgSdk interface {
		Ctrl() Imports.HyperDbgCtrlImports
		Script() Imports.HyperDbgScriptImports
		Symbol() Imports.HyperDbgSymImports
	}
	hyperDbgSdk struct {
		ctrl   Imports.HyperDbgCtrlImports
		script Imports.HyperDbgScriptImports
		symbol Imports.HyperDbgSymImports
	}
)

func New() HyperDbgSdk {
	return &hyperDbgSdk{
		ctrl:   Imports.NewHyperDbgCtrlImports(),
		script: Imports.NewHyperDbgScriptImports(),
		symbol: Imports.NewHyperDbgSymImports(),
	}
}

func (h *hyperDbgSdk) Ctrl() Imports.HyperDbgCtrlImports     { return h.ctrl }
func (h *hyperDbgSdk) Script() Imports.HyperDbgScriptImports { return h.script }
func (h *hyperDbgSdk) Symbol() Imports.HyperDbgSymImports    { return h.symbol }
