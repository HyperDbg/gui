package main

import "github.com/ddkwork/hyperdbgui/SDK"

func NoGui() { //dll load test
	sdk := SDK.New()
	sdk.Ctrl().HyperDbgInstallVmmDriver() //TODO decode error codes
	sdk.Ctrl().HyperDbgLoadVmm()
	sdk.Ctrl().HyperDbgUnloadVmm()

	//sdk.Script().PrintSymbol()
	//sdk.Script().ScriptEngineParse()

	//sdk.Symbol().SymbolInitLoad()
	select {}
}
