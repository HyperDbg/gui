package main

import (
	"github.com/ddkwork/app"
	"github.com/ddkwork/hyperdbgui/ux"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("script", func(w *unison.Window) {
		libhyperdbg.LayoutScript(w.Content())
	})
}
