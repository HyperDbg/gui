package main

import (
	"github.com/ddkwork/app"
	"github.com/ddkwork/hyperdbgui/ux"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("pe", func(w *unison.Window) {
		ux.LayoutPeView("ux\\bin\\debug\\HPRDBGCTRL.dll", w.Content())
	})
}
