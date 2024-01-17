//Code generated from mapPather - DO NOT EDIT.

package toolbar

import (
	_ "embed"
	"github.com/ddkwork/golibrary/skiaLib/widget"
	"github.com/richardwilkes/unison"
)

type kind byte

var name kind = 0

const (
	invalidKind kind = iota
	windowsKind
	appearKind
	gotoKind
	modulesKind
	memoryKind
	runKind
	sourceKind
	handlesKind
	traceKind
	logKind
	troverKind
	pauseKind
	stackKind
	tillretKind
	patchesKind
	refersKind
	stepoverKind
	trinKind
	breaksKind
	helpKind
	restartKind
	threadsKind
	closeKind
	openKind
	cpuKind
	optionsKind
	stepinKind
)

func (k kind) String() string {
	switch k {
	case helpKind:
		return "helpKind"
	case stepoverKind:
		return "stepoverKind"
	case trinKind:
		return "trinKind"
	case breaksKind:
		return "breaksKind"
	case openKind:
		return "openKind"
	case restartKind:
		return "restartKind"
	case threadsKind:
		return "threadsKind"
	case closeKind:
		return "closeKind"
	case stepinKind:
		return "stepinKind"
	case cpuKind:
		return "cpuKind"
	case optionsKind:
		return "optionsKind"
	case modulesKind:
		return "modulesKind"
	case windowsKind:
		return "windowsKind"
	case appearKind:
		return "appearKind"
	case gotoKind:
		return "gotoKind"
	case memoryKind:
		return "memoryKind"
	case runKind:
		return "runKind"
	case sourceKind:
		return "sourceKind"
	case handlesKind:
		return "handlesKind"
	case traceKind:
		return "traceKind"
	case logKind:
		return "logKind"
	case tillretKind:
		return "tillretKind"
	case troverKind:
		return "troverKind"
	case pauseKind:
		return "pauseKind"
	case stackKind:
		return "stackKind"
	case patchesKind:
		return "patchesKind"
	case refersKind:
		return "refersKind"

	}
	return "invalid bmp kind"
}

func (k kind) Image() *unison.Image {
	switch k {
	case windowsKind:
		return widget.MustImage(bmpWindows)
	case appearKind:
		return widget.MustImage(bmpAppear)
	case gotoKind:
		return widget.MustImage(bmpGoto)
	case modulesKind:
		return widget.MustImage(bmpModules)
	case memoryKind:
		return widget.MustImage(bmpMemory)
	case runKind:
		return widget.MustImage(bmpRun)
	case sourceKind:
		return widget.MustImage(bmpSource)
	case handlesKind:
		return widget.MustImage(bmpHandles)
	case traceKind:
		return widget.MustImage(bmpTrace)
	case logKind:
		return widget.MustImage(bmpLog)
	case troverKind:
		return widget.MustImage(bmpTrover)
	case pauseKind:
		return widget.MustImage(bmpPause)
	case stackKind:
		return widget.MustImage(bmpStack)
	case tillretKind:
		return widget.MustImage(bmpTillret)
	case patchesKind:
		return widget.MustImage(bmpPatches)
	case refersKind:
		return widget.MustImage(bmpRefers)
	case stepoverKind:
		return widget.MustImage(bmpStepover)
	case trinKind:
		return widget.MustImage(bmpTrin)
	case breaksKind:
		return widget.MustImage(bmpBreaks)
	case helpKind:
		return widget.MustImage(bmpHelp)
	case restartKind:
		return widget.MustImage(bmpRestart)
	case threadsKind:
		return widget.MustImage(bmpThreads)
	case closeKind:
		return widget.MustImage(bmpClose)
	case openKind:
		return widget.MustImage(bmpOpen)
	case cpuKind:
		return widget.MustImage(bmpCpu)
	case optionsKind:
		return widget.MustImage(bmpOptions)
	case stepinKind:
		return widget.MustImage(bmpStepin)

	}
	return nil
}

var (
	//go:embed 2052/BITMAP/TRACE.ico
	bmpTrace []byte

	//go:embed 2052/BITMAP/LOG.ico
	bmpLog []byte

	//go:embed 2052/BITMAP/TILLRET.ico
	bmpTillret []byte

	//go:embed 2052/BITMAP/TROVER.ico
	bmpTrover []byte

	//go:embed 2052/BITMAP/PAUSE.ico
	bmpPause []byte

	//go:embed 2052/BITMAP/STACK.ico
	bmpStack []byte

	//go:embed 2052/BITMAP/PATCHES.ico
	bmpPatches []byte

	//go:embed 2052/BITMAP/REFERS.ico
	bmpRefers []byte

	//go:embed 2052/BITMAP/HELP.ico
	bmpHelp []byte

	//go:embed 2052/BITMAP/STEPOVER.ico
	bmpStepover []byte

	//go:embed 2052/BITMAP/TRIN.ico
	bmpTrin []byte

	//go:embed 2052/BITMAP/BREAKS.ico
	bmpBreaks []byte

	//go:embed 2052/BITMAP/OPEN.ico
	bmpOpen []byte

	//go:embed 2052/BITMAP/RESTART.ico
	bmpRestart []byte

	//go:embed 2052/BITMAP/THREADS.ico
	bmpThreads []byte

	//go:embed 2052/BITMAP/CLOSE.ico
	bmpClose []byte

	//go:embed 2052/BITMAP/STEPIN.ico
	bmpStepin []byte

	//go:embed 2052/BITMAP/CPU.ico
	bmpCpu []byte

	//go:embed 2052/BITMAP/OPTIONS.ico
	bmpOptions []byte

	//go:embed 2052/BITMAP/MODULES.ico
	bmpModules []byte

	//go:embed 2052/BITMAP/WINDOWS.ico
	bmpWindows []byte

	//go:embed 2052/BITMAP/APPEAR.ico
	bmpAppear []byte

	//go:embed 2052/BITMAP/GOTO.ico
	bmpGoto []byte

	//go:embed 2052/BITMAP/MEMORY.ico
	bmpMemory []byte

	//go:embed 2052/BITMAP/RUN.ico
	bmpRun []byte

	//go:embed 2052/BITMAP/SOURCE.ico
	bmpSource []byte

	//go:embed 2052/BITMAP/HANDLES.ico
	bmpHandles []byte
)
