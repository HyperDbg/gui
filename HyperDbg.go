package main

import ( //
	"cogentcore.org/core/gi"
	"cogentcore.org/core/goosi/driver/desktop"
	"cogentcore.org/core/grr"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/states"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/units"
	"embed"
	_ "embed"
	"github.com/go-gl/glfw/v3.3/glfw"
	"io/fs"
	"os"
	"path/filepath"
)

////go:embed Fleet.ico
//var icon []byte

//go:generate go get cogentcore.org/core@main
//go:generate core generate
//go:generate core build -v -t android/arm64
//go:generate core build -v -t windows/amd64

//go:embed ICON/*.svg
var myIcons embed.FS

//go:embed pageIco/*.svg
var pageIco embed.FS

func main() {
	icons.AddFS(grr.Log1(fs.Sub(myIcons, "ICON")))
	icons.AddFS(grr.Log1(fs.Sub(pageIco, "pageIco")))

	b := gi.NewAppBody("HyperDbg")
	//b.App().SetIconBytes(icon)
	b.AddAppBar(func(tb *gi.Toolbar) {
		//o := gi.NewButton(tb).SetText("open").SetIcon(icons.FolderOpen)
		//o.On(events.MouseEnter, func(e events.Event) {
		//	println("MouseEnter")
		//	o.Styles.Font.Face.Size += 3 //todo放大字体和刷新
		//	o.ApplyStylePrefs()
		//})
		//o.On(events.MouseLeave, func(e events.Event) {
		//	println("MouseLeave")
		//	o.Styles.Font.Face.Size -= 3
		//	o.ApplyStylePrefs()
		//})

		o := gi.NewButton(tb).SetTooltip("open").SetIcon("1")
		o.OnWidgetAdded(func(w gi.Widget) {
			if lb, ok := w.(*gi.Label); ok {
				lb.Style(func(s *styles.Style) {
					if o.StateIs(states.Hovered) {
						s.Font.Size = units.Dp(17)
					} else {
						s.Font.Size = units.Dp(14)
					}
				})
			}
		})

		gi.NewButton(tb).SetTooltip("reload").SetIcon("2")
		gi.NewButton(tb).SetTooltip("stop").SetIcon("3")
		gi.NewButton(tb).SetTooltip("f7").SetIcon("4")
		gi.NewButton(tb).SetTooltip("f8").SetIcon("5")
		gi.NewButton(tb).SetTooltip("f9").SetIcon("6")
		gi.NewButton(tb).SetTooltip("scylla").SetIcon("7")
		gi.NewButton(tb).SetTooltip("path").SetIcon("8")
		gi.NewButton(tb).SetTooltip("calc").SetIcon("9")
		gi.NewButton(tb).SetTooltip("setting").SetIcon("10")
		gi.NewButton(tb).SetTooltip("about").SetIcon("11")
	})

	tabs := gi.NewTabs(b)

	pageCpu(tabs.NewTab("cpu", "Cpu.svg"))

	pageLog(tabs.NewTab("log", "log"))
	pageNotes(tabs.NewTab("notes", "notes"))

	pageBreak(tabs.NewTab("break", "breakpoint"))
	pageMemory(tabs.NewTab("memory", "memory"))
	pageCallStack(tabs.NewTab("stack", "stack"))
	pageCallSeh(tabs.NewTab("seh", "seh"))
	pageScript(tabs.NewTab("script", "script"))
	pageSymbol(tabs.NewTab("symbol", "Symbols"))
	pageSource(tabs.NewTab("source", "source"))

	tabs.NewTab("references")
	pageThread(tabs.NewTab("thread", "thread"))

	tabs.NewTab("handle", "handles")
	tabs.NewTab("trace", "traceinto")

	c := gi.NewTextField(b).SetPlaceholder("command")
	c.Style(func(s *styles.Style) {
		s.Min.X.Pw(300)
		//s.Max.X.Pw(300)//todo 取屏幕宽度为最大宽度
		//s.Display = styles.Grid
	})
	w := b.NewWindow().Run()
	win := w.MainMgr.RenderWin.GoosiWin
	ww, ok := win.(*desktop.Window)
	if ok {
		ww.Glw.SetDropCallback(func(w *glfw.Window, names []string) {
			for _, name := range names {
				println(name)
				ext := filepath.Ext(name)
				switch ext {
				case ".go", ".js":
					file, err := os.ReadFile(name)
					if err != nil {
						panic(err.Error())
					}
					println(file)
					//txbuf.SetText(file)
					//txed1.SetBuf(txbuf)
					//txed2.SetBuf(txbuf)
				}
			}
		})
	}
	w.Wait()
}
