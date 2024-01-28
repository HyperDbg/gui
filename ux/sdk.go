package ux

import "github.com/ddkwork/hyperdbgui/sdk"

func NoGui() { //dll load test
	sdk := sdk.New()
	sdk.Ctrl().HyperDbgInstallVmmDriver() //TODO decode error codes
	sdk.Ctrl().HyperDbgLoadVmm()
	sdk.Ctrl().HyperDbgUnloadVmm()

	//sdk.Script().PrintSymbol()
	//sdk.Script().ScriptEngineParse()

	//sdk.Symbol().SymbolInitLoad()
	select {}
}
