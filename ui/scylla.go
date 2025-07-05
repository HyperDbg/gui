package ui

import (
	"github.com/ddkwork/ux"
	"github.com/ddkwork/ux/languages"
)

// https://github.com/NtQuery/Scylla
func LayoutScylla() ux.Widget {
	return ux.NewCodeEditor("log.log", languages.GoKind)
}
