package main

import (
	"flag"

	"github.com/ddkwork/HyperDbg/debugger"
	_ "github.com/ddkwork/HyperDbg/debugger/log"
	"github.com/ddkwork/HyperDbg/ui"
)

func main() {
	driverOnly := flag.Bool("driver-only", false, "Only load driver, don't start UI")
	flag.Parse()

	ui.SetEnableConsoleLog(true)

	if *driverOnly {
		ui.RunDriverOnly()
	} else {
		ui.Run(func(dbg *debugger.HyperDbg) {
		})
	}
}
