//Code generated from mapPather - DO NOT EDIT.

package tabs

import (
	_ "embed"
	"github.com/ddkwork/golibrary/skiaLib/widget"
	"github.com/richardwilkes/unison"
)

type kind byte

var name kind = 0

const (
	invalidKind kind = iota
	scriptCodeKind
	stackKind
	logKind
	notesKind
	processorCpuKind
	sehChainKind
	vtKind
	vt1Kind
	downloadSymbolsKind
	sourceKind
	traceKind
	traceintoKind
	breakpointKind
	handlesKind
	memoryMapKind
	threadSwitchKind
)

func (k kind) String() string {
	switch k {
	case breakpointKind:
		return "breakpointKind"
	case handlesKind:
		return "handlesKind"
	case memoryMapKind:
		return "memoryMapKind"
	case threadSwitchKind:
		return "threadSwitchKind"
	case scriptCodeKind:
		return "scriptCodeKind"
	case stackKind:
		return "stackKind"
	case logKind:
		return "logKind"
	case notesKind:
		return "notesKind"
	case processorCpuKind:
		return "processorCpuKind"
	case sehChainKind:
		return "sehChainKind"
	case vtKind:
		return "vtKind"
	case vt1Kind:
		return "vt1Kind"
	case downloadSymbolsKind:
		return "downloadSymbolsKind"
	case sourceKind:
		return "sourceKind"
	case traceKind:
		return "traceKind"
	case traceintoKind:
		return "traceintoKind"

	}
	return "invalid bmp kind"
}

func (k kind) Name() string {
	switch k {
	case breakpointKind:
		return "breakpoint"
	case handlesKind:
		return "handles"
	case memoryMapKind:
		return "memoryMap"
	case threadSwitchKind:
		return "threadSwitch"
	case scriptCodeKind:
		return "scriptCode"
	case stackKind:
		return "stack"
	case logKind:
		return "log"
	case notesKind:
		return "notes"
	case processorCpuKind:
		return "processorCpu"
	case sehChainKind:
		return "sehChain"
	case vtKind:
		return "vt"
	case vt1Kind:
		return "vt1"
	case downloadSymbolsKind:
		return "downloadSymbols"
	case sourceKind:
		return "source"
	case traceKind:
		return "trace"
	case traceintoKind:
		return "traceinto"

	}
	return "invalid bmp name"
}

func (k kind) Image() *unison.Image {
	switch k {
	case breakpointKind:
		return widget.MustImage(pngBreakpoint)
	case handlesKind:
		return widget.MustImage(pngHandles)
	case memoryMapKind:
		return widget.MustImage(pngMemoryMap)
	case threadSwitchKind:
		return widget.MustImage(pngThreadSwitch)
	case scriptCodeKind:
		return widget.MustImage(pngScriptCode)
	case stackKind:
		return widget.MustImage(pngStack)
	case logKind:
		return widget.MustImage(pngLog)
	case notesKind:
		return widget.MustImage(pngNotes)
	case processorCpuKind:
		return widget.MustImage(pngProcessorCpu)
	case sehChainKind:
		return widget.MustImage(pngSehChain)
	case vtKind:
		return widget.MustImage(pngVt)
	case vt1Kind:
		return widget.MustImage(pngVt1)
	case downloadSymbolsKind:
		return widget.MustImage(pngDownloadSymbols)
	case sourceKind:
		return widget.MustImage(pngSource)
	case traceKind:
		return widget.MustImage(pngTrace)
	case traceintoKind:
		return widget.MustImage(pngTraceinto)

	}
	return nil
}

var (
	//go:embed pageIco/scriptCode.ico
	pngScriptCode []byte

	//go:embed pageIco/stack.ico
	pngStack []byte

	//go:embed pageIco/log.ico
	pngLog []byte

	//go:embed pageIco/notes.ico
	pngNotes []byte

	//go:embed pageIco/processorCpu.ico
	pngProcessorCpu []byte

	//go:embed pageIco/sehChain.ico
	pngSehChain []byte

	//go:embed pageIco/vt.ico
	pngVt []byte

	//go:embed pageIco/vt1.ico
	pngVt1 []byte

	//go:embed pageIco/downloadSymbols.ico
	pngDownloadSymbols []byte

	//go:embed pageIco/source.ico
	pngSource []byte

	//go:embed pageIco/trace.ico
	pngTrace []byte

	//go:embed pageIco/traceinto.ico
	pngTraceinto []byte

	//go:embed pageIco/breakpoint.ico
	pngBreakpoint []byte

	//go:embed pageIco/handles.ico
	pngHandles []byte

	//go:embed pageIco/memoryMap.ico
	pngMemoryMap []byte

	//go:embed pageIco/threadSwitch.ico
	pngThreadSwitch []byte
)
