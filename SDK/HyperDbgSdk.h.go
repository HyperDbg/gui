package SDK

import (
	"github.com/ddkwork/hyperdbgui/SDK/Imports"
)

type (
	SDK interface {
		Ctrl() Imports.Ctrl
		Script() Imports.Script
		Symbol() Imports.Symbol
	}
	sdk struct {
		ctrl   Imports.Ctrl
		script Imports.Script
		symbol Imports.Symbol
	}
)

func New() SDK {
	return &sdk{
		ctrl:   Imports.NewCtrl(),
		script: Imports.NewScript(),
		symbol: Imports.NewSymbol(),
	}
}

func (s *sdk) Ctrl() Imports.Ctrl     { return s.ctrl }
func (s *sdk) Script() Imports.Script { return s.script }
func (s *sdk) Symbol() Imports.Symbol { return s.symbol }
