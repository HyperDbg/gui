package main

import (
	"github.com/ddkwork/app"
	"github.com/ddkwork/hyperdbgui/ux"
	"github.com/richardwilkes/unison"
)

func main() {
	return
	app.Run("pe", func(w *unison.Window) {
		ux.LayoutPeView("D:\\workspace\\workspace\\branch\\gui\\bin\\debug\\HPRDBGCTRL.dll", w.Content())
	})
}