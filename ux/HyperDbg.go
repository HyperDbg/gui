package ux

import (
	"cogentcore.org/core/gi"
	"cogentcore.org/core/grr"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/styles"
	"embed"
	"github.com/ddkwork/golibrary/stream"
	"github.com/ddkwork/golibrary/widget"
	"io/fs"
	"path/filepath"
)

////go:embed Fleet.ico
//var icon []byte

//go:generate go get cogentcore.org/core@main
//go:generate core generate
//go:generate core build -v -t android/arm64
//go:generate core build -v -t windows/amd64

//go:embed asserts/ico/ico_aaamain.svg
var mainIcons []byte

//go:embed asserts/bar/*.svg
var bar embed.FS

//go:embed asserts/pageico/*.svg
var pageIco embed.FS

func Run() {
	icons.AddFS(grr.Log1(fs.Sub(bar, "asserts/bar")))
	icons.AddFS(grr.Log1(fs.Sub(pageIco, "asserts/pageico")))
	gi.TheApp.SetIconBytes(mainIcons)
	b := gi.NewBody("HyperDbg")
	b.AddAppBar(func(tb *gi.Toolbar) {
		widget.NewButton(tb).SetTooltip("open").SetIcon("open")
		widget.NewButton(tb).SetTooltip("restart").SetIcon("restart")
		widget.NewButton(tb).SetTooltip("close").SetIcon("close")
		widget.NewButton(tb).SetTooltip("run").SetIcon("run") //.SetShortcut()//todo
		widget.NewButton(tb).SetTooltip("runthread").SetIcon("runthread")
		widget.NewButton(tb).SetTooltip("pause").SetIcon("pause")
		gi.NewSeparator(tb)
		widget.NewButton(tb).SetTooltip("stepin").SetIcon("stepin")
		widget.NewButton(tb).SetTooltip("stepover").SetIcon("stepover")
		widget.NewButton(tb).SetTooltip("stepin").SetIcon("stepin")
		widget.NewButton(tb).SetTooltip("trin").SetIcon("trin")
		widget.NewButton(tb).SetTooltip("trover").SetIcon("trover")
		widget.NewButton(tb).SetTooltip("tillret").SetIcon("tillret")
		widget.NewButton(tb).SetTooltip("tilluser").SetIcon("tilluser")
		gi.NewSeparator(tb)
		widget.NewButton(tb).SetTooltip("log").SetIcon("log")
		widget.NewButton(tb).SetTooltip("modules").SetIcon("modules")
		widget.NewButton(tb).SetTooltip("windows").SetIcon("windows")
		widget.NewButton(tb).SetTooltip("threads").SetIcon("threads")
		widget.NewButton(tb).SetTooltip("cpu").SetIcon("cpu")
		widget.NewButton(tb).SetTooltip("search").SetIcon("search")
		widget.NewButton(tb).SetTooltip("trace").SetIcon("trace")
		gi.NewSeparator(tb)
		widget.NewButton(tb).SetTooltip("bpoints").SetIcon("bpoints")
		widget.NewButton(tb).SetTooltip("bpmem").SetIcon("bpmem")
		widget.NewButton(tb).SetTooltip("bphard").SetIcon("bphard")
		widget.NewButton(tb).SetTooltip("options").SetIcon("options")
		widget.NewButton(tb).SetTooltip("scylla").SetIcon("terminal")
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
	pageSymbol(tabs.NewTab("symbol", "symbols"))
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
