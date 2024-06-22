package main

import (
	"github.com/ddkwork/app"
	"github.com/ddkwork/hyperdbgui/ux"
	"github.com/richardwilkes/unison"
)

func main() {
	return
	app.Run("asm", func(w *unison.Window) {
		ux.LayoutDisassemblyTable("D:\\workspace\\workspace\\branch\\gui\\bin\\debug\\hyperdbg-cli.exe", w.Content())
	})
}
