package ui

import (
	"github.com/ddkwork/ux"
	"github.com/ddkwork/ux/languages"
)

func NewScript() ux.Widget {
	path := "plugin/scriptGen/user-mode-memory-allocations.ds"
	// path = "plugin/scriptGen/cpuid.ds"
	// path = "plugin/scriptGen/1.carbon"
	// return ux.NewDsScriptView(path)
	return ux.NewCodeEditor(path, languages.GoKind)
}
