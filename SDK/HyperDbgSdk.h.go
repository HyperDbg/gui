package HyperDbgSdk

import (
	"github.com/ddkwork/hyperdbgui/SDK/Imports"
)

type (
	HyperDbgSdk interface {
		Imports.HyperDbgCtrlImports
		Imports.HyperDbgScriptImports
		Imports.HyperDbgSymImports
	}
)

//todo add a object merge them ?
