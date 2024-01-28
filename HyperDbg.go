package main

import (
	"cogentcore.org/core/gi"
	"cogentcore.org/core/styles"
	"github.com/ddkwork/golibrary/stream"
	"github.com/ddkwork/golibrary/widget"
	"path/filepath"
)

////go:embed Fleet.ico
//var icon []byte

//go:generate go get cogentcore.org/core@main
//go:generate core generate
//go:generate core build -v -t android/arm64
//go:generate core build -v -t windows/amd64

////go:embed SND/ICO/ICO_AAAMAIN.svg
//var mainIcons []byte

////go:embed SND/ICO/*.svg
//var myIcons embed.FS

////go:embed png/*.svg
//var mypngs embed.FS

////go:embed SND/pageIco/*.svg
//var pageIco embed.FS

func main() {
	//icons.AddFS(grr.Log1(fs.Sub(myIcons, "SND/ICO")))
	//icons.AddFS(grr.Log1(fs.Sub(myIcons, "png")))
	//icons.AddFS(grr.Log1(fs.Sub(pageIco, "SND/pageIco")))
	//gi.TheApp.SetIconBytes(mainIcons)
	b := gi.NewBody("HyperDbg")
	b.AddAppBar(func(tb *gi.Toolbar) {
		widget.NewButton(tb).SetTooltip("open").SetIcon("openopen")
		widget.NewButton(tb).SetTooltip("RESTART").SetIcon("restart")
		widget.NewButton(tb).SetTooltip("CLOSE").SetIcon("close")
		widget.NewButton(tb).SetTooltip("RUN").SetIcon("run") //.SetShortcut()//todo
		widget.NewButton(tb).SetTooltip("RUNTHREAD").SetIcon("runthread")
		widget.NewButton(tb).SetTooltip("PAUSE").SetIcon("pause")
		gi.NewSeparator(tb)
		widget.NewButton(tb).SetTooltip("STEPIN").SetIcon("stepin")
		widget.NewButton(tb).SetTooltip("STEPOVER").SetIcon("stepover")
		widget.NewButton(tb).SetTooltip("STEPIN").SetIcon("STEPIN")
		widget.NewButton(tb).SetTooltip("TRIN").SetIcon("TRIN")
		widget.NewButton(tb).SetTooltip("TROVER").SetIcon("TROVER")
		widget.NewButton(tb).SetTooltip("TILLRET").SetIcon("TILLRET")
		widget.NewButton(tb).SetTooltip("TILLUSER").SetIcon("TILLUSER")
		gi.NewSeparator(tb)
		widget.NewButton(tb).SetTooltip("LOG").SetIcon("LOG")
		widget.NewButton(tb).SetTooltip("MODULES").SetIcon("MODULES")
		widget.NewButton(tb).SetTooltip("WINDOWS").SetIcon("WINDOWS")
		widget.NewButton(tb).SetTooltip("THREADS").SetIcon("THREADS")
		widget.NewButton(tb).SetTooltip("CPU").SetIcon("cpu")
		widget.NewButton(tb).SetTooltip("SEARCH").SetIcon("SEARCH")
		widget.NewButton(tb).SetTooltip("TRACE").SetIcon("TRACE")
		gi.NewSeparator(tb)
		widget.NewButton(tb).SetTooltip("BPOINTS").SetIcon("BPOINTS")
		widget.NewButton(tb).SetTooltip("BPMEM").SetIcon("BPMEM")
		widget.NewButton(tb).SetTooltip("BPHARD").SetIcon("BPHARD")
		widget.NewButton(tb).SetTooltip("OPTIONS").SetIcon("OPTIONS")
		widget.NewButton(tb).SetTooltip("scylla").SetIcon("scylla")
		widget.NewButton(tb).SetTooltip("about").SetIcon("about")
	})

	tabs := gi.NewTabs(b)

	pageCpu(tabs.NewTab("cpu", "cpu"))

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

	widget.NewWindowRunAndWait(b, func(names []string) {
		for _, name := range names {
			if filepath.Ext(name) == ".json" {
				stream.NewReadFile(name)
			}
		}
	})
}
