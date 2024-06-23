package main

import (
	"github.com/ddkwork/app"
	"github.com/ddkwork/hyperdbgui/ux"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("asm", func(w *unison.Window) {
		ux.LayoutDisassemblyTable("ux\\bin\\debug\\hyperdbg-cli.exe", w.Content())
	})
}
