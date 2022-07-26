package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\settings.cpp"
var settingsBuf string

type (
	Settings interface {
		//Fn() (ok bool)
	}
	settings struct{}
)

func Newsettings() Settings { return &settings{} }
