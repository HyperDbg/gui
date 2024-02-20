package sdk

import (
	"github.com/ddkwork/hyperdbgui/sdk/Imports"
)

type (
	Interface interface {
		Ctrl() Imports.Ctrl
		Script() Imports.Script
		Symbol() Imports.Symbol
	}
	object struct {
		ctrl   Imports.Ctrl
		script Imports.Script
		symbol Imports.Symbol
	}
)

func New() Interface {
	return &object{
		ctrl:   Imports.NewCtrl(),
		script: Imports.NewScript(),
		symbol: Imports.NewSymbol(),
	}
}

func (s *object) Ctrl() Imports.Ctrl     { return s.ctrl }
func (s *object) Script() Imports.Script { return s.script }
func (s *object) Symbol() Imports.Symbol { return s.symbol }
