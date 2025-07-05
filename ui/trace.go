package ui

import (
	"github.com/ddkwork/ux"
	"github.com/ddkwork/ux/languages"
)

func NewTrace() ux.Widget {
	return ux.NewCodeEditor("log.log", languages.GoKind)
}
