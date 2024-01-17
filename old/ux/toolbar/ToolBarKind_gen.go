package toolbar

//Code generated by ToolbarHelper - DO NOT EDIT.
import (
	_ "embed"
	"github.com/ddkwork/GolandProjects/ui/widget"

	"github.com/richardwilkes/unison"
	"strings"
)

type ToolbarKind byte

const (
	InvalidToolbarKind ToolbarKind = iota
	AppearKind
	BreaksKind
	CloseKind
	CpuKind
	GotoKind
	HandlesKind
	HelpKind
	LogKind
	MemoryKind
	ModulesKind
	OpenKind
	OptionsKind
	PatchesKind
	PauseKind
	RefersKind
	RestartKind
	RunKind
	SourceKind
	StackKind
	StepinKind
	StepoverKind
	ThreadsKind
	TillretKind
	TraceKind
	TrinKind
	TroverKind
	WindowsKind
)

func (t ToolbarKind) String() string {
	switch t {
	case AppearKind:
		return "Appear"
	case BreaksKind:
		return "Breaks"
	case CloseKind:
		return "Close"
	case CpuKind:
		return "Cpu"
	case GotoKind:
		return "Goto"
	case HandlesKind:
		return "Handles"
	case HelpKind:
		return "Help"
	case LogKind:
		return "Log"
	case MemoryKind:
		return "Memory"
	case ModulesKind:
		return "Modules"
	case OpenKind:
		return "Open"
	case OptionsKind:
		return "Options"
	case PatchesKind:
		return "Patches"
	case PauseKind:
		return "Pause"
	case RefersKind:
		return "Refers"
	case RestartKind:
		return "Restart"
	case RunKind:
		return "Run"
	case SourceKind:
		return "Source"
	case StackKind:
		return "Stack"
	case StepinKind:
		return "Stepin"
	case StepoverKind:
		return "Stepover"
	case ThreadsKind:
		return "Threads"
	case TillretKind:
		return "Tillret"
	case TraceKind:
		return "Trace"
	case TrinKind:
		return "Trin"
	case TroverKind:
		return "Trover"
	case WindowsKind:
		return "Windows"
	}
	return "invalid ToolbarKind"
}
func (t ToolbarKind) Keys() []string {
	return []string{
		"Appear",
		"Breaks",
		"Close",
		"Cpu",
		"Goto",
		"Handles",
		"Help",
		"Log",
		"Memory",
		"Modules",
		"Open",
		"Options",
		"Patches",
		"Pause",
		"Refers",
		"Restart",
		"Run",
		"Source",
		"Stack",
		"Stepin",
		"Stepover",
		"Threads",
		"Tillret",
		"Trace",
		"Trin",
		"Trover",
		"Windows",
	}
}
func (t ToolbarKind) FileName() string {
	switch t {
	case AppearKind:
		return "APPEAR.ico"
	case BreaksKind:
		return "BREAKS.ico"
	case CloseKind:
		return "CLOSE.ico"
	case CpuKind:
		return "CPU.ico"
	case GotoKind:
		return "GOTO.ico"
	case HandlesKind:
		return "HANDLES.ico"
	case HelpKind:
		return "HELP.ico"
	case LogKind:
		return "LOG.ico"
	case MemoryKind:
		return "MEMORY.ico"
	case ModulesKind:
		return "MODULES.ico"
	case OpenKind:
		return "OPEN.ico"
	case OptionsKind:
		return "OPTIONS.ico"
	case PatchesKind:
		return "PATCHES.ico"
	case PauseKind:
		return "PAUSE.ico"
	case RefersKind:
		return "REFERS.ico"
	case RestartKind:
		return "RESTART.ico"
	case RunKind:
		return "RUN.ico"
	case SourceKind:
		return "SOURCE.ico"
	case StackKind:
		return "STACK.ico"
	case StepinKind:
		return "STEPIN.ico"
	case StepoverKind:
		return "STEPOVER.ico"
	case ThreadsKind:
		return "THREADS.ico"
	case TillretKind:
		return "TILLRET.ico"
	case TraceKind:
		return "TRACE.ico"
	case TrinKind:
		return "TRIN.ico"
	case TroverKind:
		return "TROVER.ico"
	case WindowsKind:
		return "WINDOWS.ico"
	}
	return "invalid ToolbarKind"
}
func (t ToolbarKind) AssertKind(name string) ToolbarKind {
	for _, kind := range t.Kinds() {
		if strings.Contains(strings.ToLower(name), kind.String()) {
			return kind
		}
	}
	return InvalidToolbarKind
}
func (t ToolbarKind) Kinds() []ToolbarKind {
	return []ToolbarKind{
		AppearKind,
		BreaksKind,
		CloseKind,
		CpuKind,
		GotoKind,
		HandlesKind,
		HelpKind,
		LogKind,
		MemoryKind,
		ModulesKind,
		OpenKind,
		OptionsKind,
		PatchesKind,
		PauseKind,
		RefersKind,
		RestartKind,
		RunKind,
		SourceKind,
		StackKind,
		StepinKind,
		StepoverKind,
		ThreadsKind,
		TillretKind,
		TraceKind,
		TrinKind,
		TroverKind,
		WindowsKind,
	}
}
func (t ToolbarKind) Image() *unison.Image {
	switch t {
	case AppearKind:
		return widget.MustImage(appearBuf)
	case BreaksKind:
		return widget.MustImage(breaksBuf)
	case CloseKind:
		return widget.MustImage(closeBuf)
	case CpuKind:
		return widget.MustImage(cpuBuf)
	case GotoKind:
		return widget.MustImage(gotoBuf)
	case HandlesKind:
		return widget.MustImage(handlesBuf)
	case HelpKind:
		return widget.MustImage(helpBuf)
	case LogKind:
		return widget.MustImage(logBuf)
	case MemoryKind:
		return widget.MustImage(memoryBuf)
	case ModulesKind:
		return widget.MustImage(modulesBuf)
	case OpenKind:
		return widget.MustImage(openBuf)
	case OptionsKind:
		return widget.MustImage(optionsBuf)
	case PatchesKind:
		return widget.MustImage(patchesBuf)
	case PauseKind:
		return widget.MustImage(pauseBuf)
	case RefersKind:
		return widget.MustImage(refersBuf)
	case RestartKind:
		return widget.MustImage(restartBuf)
	case RunKind:
		return widget.MustImage(runBuf)
	case SourceKind:
		return widget.MustImage(sourceBuf)
	case StackKind:
		return widget.MustImage(stackBuf)
	case StepinKind:
		return widget.MustImage(stepinBuf)
	case StepoverKind:
		return widget.MustImage(stepoverBuf)
	case ThreadsKind:
		return widget.MustImage(threadsBuf)
	case TillretKind:
		return widget.MustImage(tillretBuf)
	case TraceKind:
		return widget.MustImage(traceBuf)
	case TrinKind:
		return widget.MustImage(trinBuf)
	case TroverKind:
		return widget.MustImage(troverBuf)
	case WindowsKind:
		return widget.MustImage(windowsBuf)

	}
	return nil
}

func (t ToolbarKind) Button() *unison.Button {
	switch t {
	case AppearKind:
		return widget.CreateToolBarButton(AppearKind.Buffer(), AppearKind.String())
	case BreaksKind:
		return widget.CreateToolBarButton(BreaksKind.Buffer(), BreaksKind.String())
	case CloseKind:
		return widget.CreateToolBarButton(CloseKind.Buffer(), CloseKind.String())
	case CpuKind:
		return widget.CreateToolBarButton(CpuKind.Buffer(), CpuKind.String())
	case GotoKind:
		return widget.CreateToolBarButton(GotoKind.Buffer(), GotoKind.String())
	case HandlesKind:
		return widget.CreateToolBarButton(HandlesKind.Buffer(), HandlesKind.String())
	case HelpKind:
		return widget.CreateToolBarButton(HelpKind.Buffer(), HelpKind.String())
	case LogKind:
		return widget.CreateToolBarButton(LogKind.Buffer(), LogKind.String())
	case MemoryKind:
		return widget.CreateToolBarButton(MemoryKind.Buffer(), MemoryKind.String())
	case ModulesKind:
		return widget.CreateToolBarButton(ModulesKind.Buffer(), ModulesKind.String())
	case OpenKind:
		return widget.CreateToolBarButton(OpenKind.Buffer(), OpenKind.String())
	case OptionsKind:
		return widget.CreateToolBarButton(OptionsKind.Buffer(), OptionsKind.String())
	case PatchesKind:
		return widget.CreateToolBarButton(PatchesKind.Buffer(), PatchesKind.String())
	case PauseKind:
		return widget.CreateToolBarButton(PauseKind.Buffer(), PauseKind.String())
	case RefersKind:
		return widget.CreateToolBarButton(RefersKind.Buffer(), RefersKind.String())
	case RestartKind:
		return widget.CreateToolBarButton(RestartKind.Buffer(), RestartKind.String())
	case RunKind:
		return widget.CreateToolBarButton(RunKind.Buffer(), RunKind.String())
	case SourceKind:
		return widget.CreateToolBarButton(SourceKind.Buffer(), SourceKind.String())
	case StackKind:
		return widget.CreateToolBarButton(StackKind.Buffer(), StackKind.String())
	case StepinKind:
		return widget.CreateToolBarButton(StepinKind.Buffer(), StepinKind.String())
	case StepoverKind:
		return widget.CreateToolBarButton(StepoverKind.Buffer(), StepoverKind.String())
	case ThreadsKind:
		return widget.CreateToolBarButton(ThreadsKind.Buffer(), ThreadsKind.String())
	case TillretKind:
		return widget.CreateToolBarButton(TillretKind.Buffer(), TillretKind.String())
	case TraceKind:
		return widget.CreateToolBarButton(TraceKind.Buffer(), TraceKind.String())
	case TrinKind:
		return widget.CreateToolBarButton(TrinKind.Buffer(), TrinKind.String())
	case TroverKind:
		return widget.CreateToolBarButton(TroverKind.Buffer(), TroverKind.String())
	case WindowsKind:
		return widget.CreateToolBarButton(WindowsKind.Buffer(), WindowsKind.String())
	}
	return unison.NewButton()
}
func (t ToolbarKind) Label() *unison.Label {
	label := unison.NewLabel()
	switch t {
	case AppearKind:
		label.Drawable = widget.SizedDrawable(AppearKind.Buffer())
	case BreaksKind:
		label.Drawable = widget.SizedDrawable(BreaksKind.Buffer())
	case CloseKind:
		label.Drawable = widget.SizedDrawable(CloseKind.Buffer())
	case CpuKind:
		label.Drawable = widget.SizedDrawable(CpuKind.Buffer())
	case GotoKind:
		label.Drawable = widget.SizedDrawable(GotoKind.Buffer())
	case HandlesKind:
		label.Drawable = widget.SizedDrawable(HandlesKind.Buffer())
	case HelpKind:
		label.Drawable = widget.SizedDrawable(HelpKind.Buffer())
	case LogKind:
		label.Drawable = widget.SizedDrawable(LogKind.Buffer())
	case MemoryKind:
		label.Drawable = widget.SizedDrawable(MemoryKind.Buffer())
	case ModulesKind:
		label.Drawable = widget.SizedDrawable(ModulesKind.Buffer())
	case OpenKind:
		label.Drawable = widget.SizedDrawable(OpenKind.Buffer())
	case OptionsKind:
		label.Drawable = widget.SizedDrawable(OptionsKind.Buffer())
	case PatchesKind:
		label.Drawable = widget.SizedDrawable(PatchesKind.Buffer())
	case PauseKind:
		label.Drawable = widget.SizedDrawable(PauseKind.Buffer())
	case RefersKind:
		label.Drawable = widget.SizedDrawable(RefersKind.Buffer())
	case RestartKind:
		label.Drawable = widget.SizedDrawable(RestartKind.Buffer())
	case RunKind:
		label.Drawable = widget.SizedDrawable(RunKind.Buffer())
	case SourceKind:
		label.Drawable = widget.SizedDrawable(SourceKind.Buffer())
	case StackKind:
		label.Drawable = widget.SizedDrawable(StackKind.Buffer())
	case StepinKind:
		label.Drawable = widget.SizedDrawable(StepinKind.Buffer())
	case StepoverKind:
		label.Drawable = widget.SizedDrawable(StepoverKind.Buffer())
	case ThreadsKind:
		label.Drawable = widget.SizedDrawable(ThreadsKind.Buffer())
	case TillretKind:
		label.Drawable = widget.SizedDrawable(TillretKind.Buffer())
	case TraceKind:
		label.Drawable = widget.SizedDrawable(TraceKind.Buffer())
	case TrinKind:
		label.Drawable = widget.SizedDrawable(TrinKind.Buffer())
	case TroverKind:
		label.Drawable = widget.SizedDrawable(TroverKind.Buffer())
	case WindowsKind:
		label.Drawable = widget.SizedDrawable(WindowsKind.Buffer())
	}
	return label
}
func (t ToolbarKind) Buffer() []byte {
	switch t {
	case AppearKind:
		return appearBuf
	case BreaksKind:
		return breaksBuf
	case CloseKind:
		return closeBuf
	case CpuKind:
		return cpuBuf
	case GotoKind:
		return gotoBuf
	case HandlesKind:
		return handlesBuf
	case HelpKind:
		return helpBuf
	case LogKind:
		return logBuf
	case MemoryKind:
		return memoryBuf
	case ModulesKind:
		return modulesBuf
	case OpenKind:
		return openBuf
	case OptionsKind:
		return optionsBuf
	case PatchesKind:
		return patchesBuf
	case PauseKind:
		return pauseBuf
	case RefersKind:
		return refersBuf
	case RestartKind:
		return restartBuf
	case RunKind:
		return runBuf
	case SourceKind:
		return sourceBuf
	case StackKind:
		return stackBuf
	case StepinKind:
		return stepinBuf
	case StepoverKind:
		return stepoverBuf
	case ThreadsKind:
		return threadsBuf
	case TillretKind:
		return tillretBuf
	case TraceKind:
		return traceBuf
	case TrinKind:
		return trinBuf
	case TroverKind:
		return troverBuf
	case WindowsKind:
		return windowsBuf
	}
	return nil
}

var (
	//go:embed 2052/BITMAP/APPEAR.ico
	appearBuf []byte

	//go:embed 2052/BITMAP/BREAKS.ico
	breaksBuf []byte

	//go:embed 2052/BITMAP/CLOSE.ico
	closeBuf []byte

	//go:embed 2052/BITMAP/CPU.ico
	cpuBuf []byte

	//go:embed 2052/BITMAP/GOTO.ico
	gotoBuf []byte

	//go:embed 2052/BITMAP/HANDLES.ico
	handlesBuf []byte

	//go:embed 2052/BITMAP/HELP.ico
	helpBuf []byte

	//go:embed 2052/BITMAP/LOG.ico
	logBuf []byte

	//go:embed 2052/BITMAP/MEMORY.ico
	memoryBuf []byte

	//go:embed 2052/BITMAP/MODULES.ico
	modulesBuf []byte

	//go:embed 2052/BITMAP/OPEN.ico
	openBuf []byte

	//go:embed 2052/BITMAP/OPTIONS.ico
	optionsBuf []byte

	//go:embed 2052/BITMAP/PATCHES.ico
	patchesBuf []byte

	//go:embed 2052/BITMAP/PAUSE.ico
	pauseBuf []byte

	//go:embed 2052/BITMAP/REFERS.ico
	refersBuf []byte

	//go:embed 2052/BITMAP/RESTART.ico
	restartBuf []byte

	//go:embed 2052/BITMAP/RUN.ico
	runBuf []byte

	//go:embed 2052/BITMAP/SOURCE.ico
	sourceBuf []byte

	//go:embed 2052/BITMAP/STACK.ico
	stackBuf []byte

	//go:embed 2052/BITMAP/STEPIN.ico
	stepinBuf []byte

	//go:embed 2052/BITMAP/STEPOVER.ico
	stepoverBuf []byte

	//go:embed 2052/BITMAP/THREADS.ico
	threadsBuf []byte

	//go:embed 2052/BITMAP/TILLRET.ico
	tillretBuf []byte

	//go:embed 2052/BITMAP/TRACE.ico
	traceBuf []byte

	//go:embed 2052/BITMAP/TRIN.ico
	trinBuf []byte

	//go:embed 2052/BITMAP/TROVER.ico
	troverBuf []byte

	//go:embed 2052/BITMAP/WINDOWS.ico
	windowsBuf []byte
)
