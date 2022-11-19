package SDK

import (
	"github.com/ddkwork/hyperdbgui/SDK/Imports"
)

type (
	HyperDbgSdk interface {
		Ctrl() Imports.Ctrl
		Script() Imports.Script
		Symbol() Imports.Symbol
	}
	hyperDbgSdk struct {
		ctrl   Imports.Ctrl
		script Imports.Script
		symbol Imports.Symbol
	}
)

func New() HyperDbgSdk {
	return &hyperDbgSdk{
		ctrl:   Imports.NewCtrl(),
		script: Imports.NewScript(),
		symbol: Imports.NewSymbol(),
	}
}

func (h *hyperDbgSdk) Ctrl() Imports.Ctrl     { return h.ctrl }
func (h *hyperDbgSdk) Script() Imports.Script { return h.script }
func (h *hyperDbgSdk) Symbol() Imports.Symbol { return h.symbol }
