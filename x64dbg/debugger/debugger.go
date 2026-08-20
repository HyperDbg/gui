package debugger

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"syscall"
	"time"

	peparser "github.com/saferwall/pe"
	"golang.org/x/arch/x86/x86asm"

	"github.com/ddkwork/ddk/xed"
	"github.com/ddkwork/ux/app"
	"github.com/ddkwork/x64dbg/debugger/api"
	"github.com/ddkwork/x64dbg/debugger/ark"
	"github.com/ddkwork/x64dbg/debugger/bookmark"
	"github.com/ddkwork/x64dbg/debugger/breakpoint"
	"github.com/ddkwork/x64dbg/debugger/comment"
	"github.com/ddkwork/x64dbg/debugger/disassembly"
	"github.com/ddkwork/x64dbg/debugger/exception"
	"github.com/ddkwork/x64dbg/debugger/function"
	"github.com/ddkwork/x64dbg/debugger/graph"
	"github.com/ddkwork/x64dbg/debugger/imm"
	"github.com/ddkwork/x64dbg/debugger/label"
	"github.com/ddkwork/x64dbg/debugger/loop"
	"github.com/ddkwork/x64dbg/debugger/memory"
	"github.com/ddkwork/x64dbg/debugger/peview"
	"github.com/ddkwork/x64dbg/debugger/register"
	"github.com/ddkwork/x64dbg/debugger/scylla"
	"github.com/ddkwork/x64dbg/debugger/seh"
	"github.com/ddkwork/x64dbg/debugger/stack"
	"github.com/ddkwork/x64dbg/debugger/symbol"
	"github.com/ddkwork/x64dbg/debugger/thread"
	"github.com/ddkwork/x64dbg/debugger/trace"
	"github.com/ddkwork/x64dbg/debugger/type_"
	"github.com/ddkwork/x64dbg/debugger/watch"
	"github.com/ddkwork/x64dbg/debugger/windows"
	"github.com/ddkwork/x64dbg/debugger/xref"
)

type DebuggerState int

const (
	StateInitialized DebuggerState = iota
	StatePaused
	StateRunning
	StateStopped
)

type Debugger struct {
	processHandle    windows.Handle
	processId        uint32
	state            DebuggerState
	entryPoint       uint64
	baseAddress      uint64
	exePath          string
	mainThreadHandle windows.Handle
	mainThreadId     uint32
	lastEvent        *windows.DebugEvent

	breakpoints api.Interface
	memory      api.Interface
	registers   api.Interface
	symbols     api.Interface
	threads     api.Interface
	scylla      api.Interface
	seh         api.Interface
	stack       api.Interface
	trace       api.Interface
	labels      api.Interface
	comments    api.Interface
	functions   api.Interface
	xrefs       api.Interface
	types       api.Interface
	watches     api.Interface
	graphs      api.Interface
	exceptions  api.Interface
	bookmarks   api.Interface
	loops       api.Interface
	disassembly api.Interface
	peview      api.Interface
	ark         api.Interface
	imm         api.Interface

	eventChan        chan *windows.DebugEvent
	stopChan         chan struct{}
	continueChan     chan struct{}
	eventLoopWG      sync.WaitGroup
	detachRequested  bool
	terminatePending bool
	mu               sync.Mutex

	uiComponents    []api.Interface
	pendingStepInto bool
	dumpAddress     uint64
	dumpData        []byte
	OnTitleUpdate   func(title string)
}

type DebugEventCallback func(event *windows.DebugEvent)

func New() *Debugger {
	memMgr := memory.New()
	regMgr := register.New()
	symMgr := symbol.New()
	threadMgr := thread.New()

	dbg := &Debugger{
		state:        StateInitialized,
		breakpoints:  breakpoint.New(),
		memory:       memMgr,
		registers:    regMgr,
		symbols:      symMgr,
		threads:      threadMgr,
		scylla:       scylla.New(),
		seh:          seh.New(),
		stack:        stack.New(),
		trace:        trace.New(),
		labels:       label.New(),
		comments:     comment.New(),
		functions:    function.New(),
		xrefs:        xref.New(),
		types:        type_.New(),
		watches:      watch.New(),
		graphs:       graph.New(),
		exceptions:   exception.New(),
		bookmarks:    bookmark.New(),
		loops:        loop.New(),
		disassembly:  disassembly.New(),
		peview:       peview.New(),
		ark:          ark.New(),
		imm:          imm.New(),
		eventChan:    make(chan *windows.DebugEvent, 100),
		stopChan:     make(chan struct{}),
		continueChan: make(chan struct{}, 1),
	}

	dbg.RegisterUIComponent(dbg.breakpoints)
	dbg.RegisterUIComponent(dbg.memory)
	dbg.RegisterUIComponent(dbg.registers)
	dbg.RegisterUIComponent(dbg.symbols)
	dbg.RegisterUIComponent(dbg.threads)
	dbg.RegisterUIComponent(dbg.scylla)
	dbg.RegisterUIComponent(dbg.seh)
	dbg.RegisterUIComponent(dbg.stack)
	dbg.RegisterUIComponent(dbg.trace)
	dbg.RegisterUIComponent(dbg.labels)
	dbg.RegisterUIComponent(dbg.comments)
	dbg.RegisterUIComponent(dbg.functions)
	dbg.RegisterUIComponent(dbg.xrefs)
	dbg.RegisterUIComponent(dbg.types)
	dbg.RegisterUIComponent(dbg.watches)
	dbg.RegisterUIComponent(dbg.graphs)
	dbg.RegisterUIComponent(dbg.exceptions)
	dbg.RegisterUIComponent(dbg.bookmarks)
	dbg.RegisterUIComponent(dbg.loops)
	dbg.RegisterUIComponent(dbg.disassembly)
	dbg.RegisterUIComponent(dbg.peview)
	dbg.RegisterUIComponent(dbg.ark)
	dbg.RegisterUIComponent(dbg.imm)

	dbg.setupDisassemblyCallback()

	return dbg
}

func (d *Debugger) setupDisassemblyCallback() {
	dasm := d.disassembly.Self().(*disassembly.Disassembler)
	dasm.OnRowSelected = func(data *xed.Disassembly) {
		d.updateImmFromInstruction(data)
	}
	dasm.OnFollowInDump = func(address uint64) {
		d.followInDump(address)
	}

	immMgr := d.imm.Self().(*imm.Imm)
	immMgr.OnFollowInDump = func(address uint64) {
		d.followInDump(address)
	}
}

func (d *Debugger) followInDump(address uint64) {
	if d.processHandle == 0 {
		return
	}

	data, err := d.memory.Self().(*memory.Manager).ReadMemory(address, 256)
	if err != nil {
		return
	}

	d.dumpAddress = address
	d.dumpData = data
}

func (d *Debugger) updateImmFromInstruction(data *xed.Disassembly) {
	if data == nil {
		return
	}
	immMgr := d.imm.Self().(*imm.Imm)
	var rows []imm.ImmData

	if len(data.Opcode) > 0 {
		mode := 64
		inst, err := x86asm.Decode(data.Opcode, mode)
		if err == nil {
			for _, arg := range inst.Args {
				switch a := arg.(type) {
				case x86asm.Imm:
					rows = append(rows, imm.ImmData{
						Reg:     "Imm",
						Address: uint64(a),
						Mem:     "",
					})
				case x86asm.Mem:
					addr := uint64(a.Disp)
					if a.Base != 0 {
						rows = append(rows, imm.ImmData{
							Reg:     a.Base.String(),
							Address: addr,
							Mem:     "",
						})
					} else if addr != 0 {
						rows = append(rows, imm.ImmData{
							Reg:     "Mem",
							Address: addr,
							Mem:     "",
						})
					}
				case x86asm.Reg:
					rows = append(rows, imm.ImmData{
						Reg:     a.String(),
						Address: 0,
						Mem:     "",
					})
				}
			}
		}
	}

	immMgr.UpdateRows(rows)
}

func (d *Debugger) RegisterUIComponent(component api.Interface) {
	d.uiComponents = append(d.uiComponents, component)
}

func (d *Debugger) UpdateAllPages() error {
	d.mu.Lock()
	handle := d.processHandle
	d.mu.Unlock()

	if handle == 0 {
		return nil
	}

	for _, component := range d.uiComponents {
		d.mu.Lock()
		handle := d.processHandle
		d.mu.Unlock()
		if handle == 0 {
			continue
		}

		switch component {
		case d.disassembly:
			var buffer []byte
			var rip uint64

			threadId := d.mainThreadId
			if d.lastEvent != nil {
				threadId = d.lastEvent.ThreadId()
			}
			threadHandle, err := windows.OpenThread(windows.THREAD_GET_CONTEXT, false, threadId)
			if err == nil {
				defer windows.CloseHandle(threadHandle)
				regCtx, err := d.registers.Self().(*register.Manager).GetThreadContext(threadHandle)
				if err == nil {
					rip = regCtx.RIP
					disasm := d.disassembly.Self().(*disassembly.Disassembler)
					if disasm.IsAddressInRange(rip) {
						disasm.SelectAddress(rip)
					} else {
						buffer, _ = d.memory.Self().(*memory.Manager).ReadMemory(rip, 4096)
						if buffer != nil {
							disasm.SetRipBuffer(rip, buffer)
							disasm.Update()
						}
					}
				}
			}
		case d.stack:
			threadId := d.mainThreadId
			if d.lastEvent != nil {
				threadId = d.lastEvent.ThreadId()
			}
			threadHandle, err := windows.OpenThread(windows.THREAD_GET_CONTEXT, false, threadId)
			if err == nil {
				defer windows.CloseHandle(threadHandle)
				regCtx, err := d.registers.Self().(*register.Manager).GetThreadContext(threadHandle)
				if err == nil {
					frames, err := d.stack.Self().(*stack.Manager).WalkStack(
						threadHandle,
						regCtx,
						d.memory.Self().(*memory.Manager).ReadMemory,
						d.symbols.Self().(*symbol.Manager).GetSymbolFromAddress,
						d.symbols.Self().(*symbol.Manager).GetModuleByAddress,
					)
					if err == nil {
						table := d.stack.Self().(*stack.Manager).GetTable()
						table.Root().SetChildren(nil)
						for _, frame := range frames {
							table.Root().AddChild(table.NewNode(*frame))
						}
						table.AirTable.Refresh()
					}
					d.stack.Self().(*stack.Manager).UpdateStack(
						regCtx.RSP,
						16,
						d.memory.Self().(*memory.Manager).ReadMemory,
					)
				}
			}
		case d.seh:
			threadId := d.mainThreadId
			if d.lastEvent != nil {
				threadId = d.lastEvent.ThreadId()
			}
			threadHandle, err := windows.OpenThread(windows.THREAD_GET_CONTEXT, false, threadId)
			if err == nil {
				defer windows.CloseHandle(threadHandle)
				regCtx, err := d.registers.Self().(*register.Manager).GetThreadContext(threadHandle)
				if err == nil {
					handlers, err := d.seh.Self().(*seh.Manager).ScanSEH(
						regCtx,
						d.memory.Self().(*memory.Manager).ReadMemory,
					)
					if err == nil {
						table := d.seh.Self().(*seh.Manager).GetTable()
						table.Root().SetChildren(nil)
						for _, handler := range handlers {
							table.Root().AddChild(table.NewNode(*handler))
						}
						table.AirTable.Refresh()
					}
				}
			}
		case d.scylla:
			if d.baseAddress != 0 {
				region, err := d.memory.Self().(*memory.Manager).QueryMemory(d.baseAddress)
				if err == nil {
					peData, err := d.memory.Self().(*memory.Manager).ReadMemory(d.baseAddress, uint32(region.RegionSize))
					if err == nil {
						moduleName := "main"
						_, err := d.scylla.Self().(*scylla.Manager).ScanImports(
							d.baseAddress,
							peData,
							moduleName,
							region.RegionSize,
						)
						if err == nil {
							table := d.scylla.Self().(*scylla.Manager).GetTable()
							table.Root().SetChildren(nil)
							modules := d.scylla.Self().(*scylla.Manager).GetAllModules()
							for _, module := range modules {
								for _, imp := range module.Imports {
									table.Root().AddChild(table.NewNode(*imp))
								}
							}
							table.AirTable.Refresh()
						}
					}
				}
			}
		case d.registers:
			allThreads := d.threads.Self().(*thread.Manager).GetAllThreads()
			if len(allThreads) > 0 {
				thread := allThreads[0]
				threadHandle, err := windows.OpenThread(windows.THREAD_GET_CONTEXT, false, thread.Id)
				if err == nil {
					defer windows.CloseHandle(threadHandle)
					d.registers.Self().(*register.Manager).UpdateFromThreadHandle(threadHandle)
				}
			}
		case d.trace, d.breakpoints, d.memory, d.imm, d.symbols, d.threads,
			d.labels, d.comments, d.functions, d.xrefs, d.types, d.watches,
			d.graphs, d.exceptions, d.bookmarks, d.loops, d.peview, d.ark:
			component.Update()
		default:
			component.Update()
		}
	}
	app.RequestRedraw()
	return nil
}

func (d *Debugger) ClearAllPages() {
	for _, component := range d.uiComponents {
		component.Clear()
	}
}

func parsePEFile(exePath string) (uint64, uint64, error) {
	data, err := os.ReadFile(exePath)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to read PE file: %w", err)
	}

	pe, err := peparser.NewBytes(data, &peparser.Options{
		Fast: true,
	})
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse PE file: %w", err)
	}
	defer pe.Close()

	err = pe.Parse()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse PE structure: %w", err)
	}

	var imageBase uint64
	var entryPointRVA uint32

	if pe.Is64 {
		oh := pe.NtHeader.OptionalHeader.(peparser.ImageOptionalHeader64)
		imageBase = oh.ImageBase
		entryPointRVA = oh.AddressOfEntryPoint
	} else if pe.Is32 {
		oh := pe.NtHeader.OptionalHeader.(peparser.ImageOptionalHeader32)
		imageBase = uint64(oh.ImageBase)
		entryPointRVA = oh.AddressOfEntryPoint
	} else {
		return 0, 0, fmt.Errorf("unsupported PE architecture")
	}

	// 将入口点 RVA 转换为实际的内存地址
	entryPoint := imageBase + uint64(entryPointRVA)

	return imageBase, entryPoint, nil
}

func (d *Debugger) Attach(pid uint32) error {
	handle, err := windows.DebugActiveProcess(pid)
	if err != nil {
		return err
	}

	d.memory.Self().(*memory.Manager).SetHandle(handle)
	d.symbols.Self().(*symbol.Manager).Initialize(handle)

	d.startEventLoop()

	d.processHandle = handle
	d.processId = pid
	d.state = StatePaused

	return nil
}

func (d *Debugger) Detach() error {
	d.mu.Lock()
	if d.processHandle == 0 {
		d.mu.Unlock()
		return nil
	}

	d.detachRequested = true
	pid := d.processId
	d.mu.Unlock()

	select {
	case d.continueChan <- struct{}{}:
	default:
	}

	windows.DebugActiveProcessStop(pid)

	d.mu.Lock()
	d.processHandle = 0
	d.processId = 0
	d.state = StateStopped
	d.mu.Unlock()

	if d.OnTitleUpdate != nil {
		d.OnTitleUpdate("")
	}

	return nil
}

func (d *Debugger) doDetach() {
	windows.DebugActiveProcessStop(d.processId)
	d.mu.Lock()
	d.processHandle = 0
	d.processId = 0
	d.state = StateStopped
	d.mu.Unlock()
	d.ClearAllPages()
	if d.OnTitleUpdate != nil {
		d.OnTitleUpdate("")
	}
}

func (d *Debugger) CreateProcess(exePath string, cmdLine string) error {
	d.mu.Lock()
	if d.processHandle != 0 {
		d.mu.Unlock()
		d.TerminateProcess(0)
		for range 10 {
			time.Sleep(5 * time.Millisecond)
			d.mu.Lock()
			done := d.processHandle == 0
			d.mu.Unlock()
			if done {
				break
			}
		}
		d.mu.Lock()
	}
	d.mu.Unlock()

	d.mu.Lock()
	d.exePath = exePath
	d.detachRequested = false
	d.terminatePending = false
	d.mu.Unlock()

	imageBase, entryPoint, err := parsePEFile(exePath)
	if err == nil {
		d.baseAddress = imageBase
		d.entryPoint = entryPoint
	}

	d.eventLoopWG.Add(1)
	go func() {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()

		handle, pid, threadHandle, threadId, createErr := windows.CreateProcessDebug(exePath, cmdLine)
		if createErr != nil {
			d.eventLoopWG.Done()
			return
		}

		d.processHandle = handle
		d.processId = pid
		d.state = StatePaused

		d.memory.Self().(*memory.Manager).SetHandle(handle)
		d.symbols.Self().(*symbol.Manager).Initialize(handle)

		d.mainThreadHandle = threadHandle
		d.mainThreadId = threadId

		d.updateWindowTitle()

		d.eventLoop()
	}()

	return nil
}

func (d *Debugger) TerminateProcess(exitCode uint32) error {
	d.mu.Lock()
	if d.processHandle == 0 {
		d.mu.Unlock()
		return nil
	}

	handle := d.processHandle
	d.terminatePending = true
	d.mu.Unlock()

	select {
	case d.continueChan <- struct{}{}:
	default:
	}

	windows.TerminateProcess(handle, exitCode)
	windows.CloseHandle(handle)

	d.mu.Lock()
	d.processHandle = 0
	d.processId = 0
	d.state = StateStopped
	d.mu.Unlock()

	d.ClearAllPages()

	if d.OnTitleUpdate != nil {
		d.OnTitleUpdate("")
	}

	return nil
}

func (d *Debugger) LoadPEFile(exePath string) error {
	imageBase, entryPoint, err := parsePEFile(exePath)
	if err != nil {
		return err
	}

	d.exePath = exePath
	d.baseAddress = imageBase
	d.entryPoint = entryPoint

	d.peview.Self().(*peview.Manager).SetTargetPath(exePath)

	d.UpdateAllPages()

	return nil
}

func (d *Debugger) GetState() DebuggerState {
	// 无锁获取状态
	return d.state
}

func (d *Debugger) GetProcessHandle() windows.Handle {
	// 无锁获取进程句柄
	return d.processHandle
}

func (d *Debugger) GetProcessId() uint32 {
	// 无锁获取进程ID
	return d.processId
}

func (d *Debugger) GetEventChan() <-chan *windows.DebugEvent {
	return d.eventChan
}

func (d *Debugger) GetBreakpoints() *breakpoint.Manager {
	return d.breakpoints.Self().(*breakpoint.Manager)
}

func (d *Debugger) GetMemory() *memory.Manager {
	return d.memory.Self().(*memory.Manager)
}

func (d *Debugger) GetRegisters() *register.Manager {
	return d.registers.Self().(*register.Manager)
}

func (d *Debugger) GetSymbols() *symbol.Manager {
	return d.symbols.Self().(*symbol.Manager)
}

func (d *Debugger) GetThreads() *thread.Manager {
	return d.threads.Self().(*thread.Manager)
}

func (d *Debugger) GetScylla() *scylla.Manager {
	return d.scylla.Self().(*scylla.Manager)
}

func (d *Debugger) GetSEH() *seh.Manager {
	return d.seh.Self().(*seh.Manager)
}

func (d *Debugger) GetStack() *stack.Manager {
	return d.stack.Self().(*stack.Manager)
}

func (d *Debugger) GetTrace() *trace.Manager {
	return d.trace.Self().(*trace.Manager)
}

func (d *Debugger) GetLabels() *label.Manager {
	return d.labels.Self().(*label.Manager)
}

func (d *Debugger) GetComments() *comment.Manager {
	return d.comments.Self().(*comment.Manager)
}

func (d *Debugger) GetFunctions() *function.Manager {
	return d.functions.Self().(*function.Manager)
}

func (d *Debugger) GetXrefs() *xref.Manager {
	return d.xrefs.Self().(*xref.Manager)
}

func (d *Debugger) GetTypes() *type_.TypeManager {
	return d.types.Self().(*type_.TypeManager)
}

func (d *Debugger) GetWatches() *watch.Manager {
	return d.watches.Self().(*watch.Manager)
}

func (d *Debugger) GetGraphs() *graph.Manager {
	return d.graphs.Self().(*graph.Manager)
}

func (d *Debugger) GetExceptions() *exception.Manager {
	return d.exceptions.Self().(*exception.Manager)
}

func (d *Debugger) GetBookmarks() *bookmark.Manager {
	return d.bookmarks.Self().(*bookmark.Manager)
}

func (d *Debugger) GetLoops() *loop.Manager {
	return d.loops.Self().(*loop.Manager)
}

func (d *Debugger) GetDisassembly() api.Interface {
	return d.disassembly
}

func (d *Debugger) GetPeView() *peview.Manager {
	return d.peview.Self().(*peview.Manager)
}

func (d *Debugger) GetArk() *ark.Manager {
	return d.ark.Self().(*ark.Manager)
}

func (d *Debugger) GetImm() api.Interface {
	return d.imm
}

func (d *Debugger) GetEntryPoint() uint64 {
	// 无锁获取入口点
	return d.entryPoint
}

func (d *Debugger) GetBaseAddress() uint64 {
	// 无锁获取基地址
	return d.baseAddress
}

func (d *Debugger) GetExePath() string {
	return d.exePath
}

func (d *Debugger) updateWindowTitle() {
	if d.OnTitleUpdate != nil {
		fileName := d.exePath
		if idx := len(d.exePath) - 1; idx > 0 {
			for i := len(d.exePath) - 1; i >= 0; i-- {
				if d.exePath[i] == '\\' || d.exePath[i] == '/' {
					fileName = d.exePath[i+1:]
					break
				}
			}
		}
		title := fmt.Sprintf("%s - PID: %d", fileName, d.processId)
		d.OnTitleUpdate(title)
	}
}

func (d *Debugger) GetMainThreadContext() (*register.RegisterContext, error) {
	if d.processHandle == 0 {
		return nil, fmt.Errorf("process not attached")
	}

	allThreads := d.threads.Self().(*thread.Manager).GetAllThreads()
	if len(allThreads) == 0 {
		return nil, fmt.Errorf("no threads found")
	}

	mainThread := allThreads[0]
	threadHandle, err := windows.OpenThread(windows.THREAD_GET_CONTEXT, false, mainThread.Id)
	if err != nil {
		return nil, fmt.Errorf("OpenThread failed: %v", err)
	}
	defer windows.CloseHandle(threadHandle)

	return d.registers.Self().(*register.Manager).GetThreadContext(threadHandle)
}

func (d *Debugger) startEventLoop() {
	d.eventLoopWG.Add(1)
	go func() {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		d.eventLoop()
	}()
}

func (d *Debugger) stopEventLoop() {
	select {
	case <-d.stopChan:
	default:
		close(d.stopChan)
	}
	d.stopChan = make(chan struct{})
}

func (d *Debugger) eventLoop() {
	defer d.eventLoopWG.Done()

	for {
		if d.processHandle == 0 {
			return
		}

		if d.detachRequested {
			d.doDetach()
			return
		}

		if d.terminatePending {
			return
		}

		event, err := windows.WaitForDebugEvent()
		if err != nil {
			continue
		}

		if event == nil {
			continue
		}

		d.lastEvent = event
		d.handleDebugEvent(event)

		if d.eventChan != nil {
			select {
			case d.eventChan <- event:
			default:
			}
		}

		continueStatus := uint32(windows.DBG_CONTINUE)

		switch event.DebugEventCode() {
		case windows.LOAD_DLL_DEBUG_EVENT, windows.CREATE_THREAD_DEBUG_EVENT, windows.EXIT_THREAD_DEBUG_EVENT, windows.CREATE_PROCESS_DEBUG_EVENT:
			windows.ContinueDebugEvent(event.ProcessId(), event.ThreadId(), continueStatus)
			d.state = StateRunning
			continue
		case windows.EXCEPTION_DEBUG_EVENT:
			if event.Exception() != nil {
				switch event.Exception().ExceptionRecord.ExceptionCode {
				case windows.EXCEPTION_BREAKPOINT:
					addr := event.Exception().ExceptionRecord.ExceptionAddress
					isUserBreakpoint := d.breakpoints.Self().(*breakpoint.Manager).FindByAddress(addr) != nil

					if d.pendingStepInto && !isUserBreakpoint {
						windows.ContinueDebugEvent(event.ProcessId(), event.ThreadId(), continueStatus)
						d.state = StateRunning
						d.pendingStepInto = false
						continue
					}
				case windows.EXCEPTION_SINGLE_STEP:
					d.pendingStepInto = false
				case 0:
					windows.ContinueDebugEvent(event.ProcessId(), event.ThreadId(), continueStatus)
					d.state = StateRunning
					continue
				}
			} else {
				windows.ContinueDebugEvent(event.ProcessId(), event.ThreadId(), continueStatus)
				d.state = StateRunning
				continue
			}
		case windows.EXIT_PROCESS_DEBUG_EVENT:
			d.state = StateStopped
			d.ClearAllPages()
			windows.ContinueDebugEvent(event.ProcessId(), event.ThreadId(), continueStatus)
			return
		}

		if d.terminatePending || d.detachRequested {
			windows.ContinueDebugEvent(event.ProcessId(), event.ThreadId(), continueStatus)
			if d.detachRequested {
				d.doDetach()
			}
			return
		}

		d.state = StatePaused
		if d.processHandle != 0 {
			d.UpdateAllPages()
		}

		select {
		case <-d.continueChan:
			if d.terminatePending {
				windows.ContinueDebugEvent(event.ProcessId(), event.ThreadId(), continueStatus)
				return
			}
			if d.detachRequested {
				windows.ContinueDebugEvent(event.ProcessId(), event.ThreadId(), continueStatus)
				d.doDetach()
				return
			}
			windows.ContinueDebugEvent(event.ProcessId(), event.ThreadId(), continueStatus)
			d.state = StateRunning
		case <-d.stopChan:
			windows.ContinueDebugEvent(event.ProcessId(), event.ThreadId(), continueStatus)
			return
		}
	}
}

func (d *Debugger) handleDebugEvent(event *windows.DebugEvent) {
	if event == nil || d == nil {
		return
	}

	switch event.DebugEventCode() {
	case windows.EXCEPTION_DEBUG_EVENT:
		if event.Exception() != nil {
			d.handleException(event)
		}
	case windows.CREATE_THREAD_DEBUG_EVENT:
		if event.CreateThread() != nil && d.threads != nil {
			thread := d.threads.Self().(*thread.Manager).AddThread(event.ThreadId(), event.CreateThread().ThreadHandle, event.CreateThread().ThreadLocalBase)
			if thread != nil && d.registers != nil {
				d.registers.Self().(*register.Manager).SetThreadHandle(thread.Id, thread.Handle)
			}
		}
	case windows.EXIT_THREAD_DEBUG_EVENT:
		if d.threads != nil {
			d.threads.Self().(*thread.Manager).RemoveThread(event.ThreadId())
		}
	case windows.CREATE_PROCESS_DEBUG_EVENT:
		if event.CreateProcess() != nil {
			if d.threads != nil && d.mainThreadHandle != 0 {
				thread := d.threads.Self().(*thread.Manager).AddThread(d.mainThreadId, d.mainThreadHandle, 0)
				if thread != nil && d.registers != nil {
					d.registers.Self().(*register.Manager).SetThreadHandle(thread.Id, thread.Handle)
				}
			}
			if d.symbols != nil {
				d.symbols.Self().(*symbol.Manager).LoadModuleForProcess(d.processId, 0, d.baseAddress)
			}
		}
	case windows.EXIT_PROCESS_DEBUG_EVENT:
		d.state = StateStopped
		d.stopEventLoop()
	case windows.LOAD_DLL_DEBUG_EVENT:
		if event.LoadDll() != nil && d.symbols != nil && d.processId != 0 {
			d.symbols.Self().(*symbol.Manager).LoadModule(d.processId, event.LoadDll().FileHandle, event.LoadDll().BaseOfDll)
		}
	case windows.UNLOAD_DLL_DEBUG_EVENT:
		if event.UnloadDll() != nil && d.symbols != nil {
			d.symbols.Self().(*symbol.Manager).UnloadModule(event.UnloadDll().BaseOfDll)
		}
	case windows.OUTPUT_DEBUG_STRING_EVENT:
	}
}

func (d *Debugger) handleException(event *windows.DebugEvent) {
	exception := &event.Exception().ExceptionRecord

	if d.threads != nil && d.registers != nil {
		allThreads := d.threads.Self().(*thread.Manager).GetAllThreads()
		for _, thread := range allThreads {
			threadHandle := d.registers.Self().(*register.Manager).GetThreadHandle(thread.Id)
			if threadHandle != 0 {
				d.registers.Self().(*register.Manager).GetThreadContext(threadHandle)
			}
		}
	}

	isBreakpoint := false

	if exception.ExceptionCode == windows.EXCEPTION_BREAKPOINT {
		addr := exception.ExceptionAddress
		if bp := d.breakpoints.Self().(*breakpoint.Manager).FindByAddress(addr); bp != nil {
			bp.HitCount++
			d.state = StatePaused
			isBreakpoint = true
			if bp.SingleShot {
				d.breakpoints.Self().(*breakpoint.Manager).RemoveSoftwareBreakpoint(d.processHandle, bp)
				d.breakpoints.Self().(*breakpoint.Manager).RemoveBreakpoint(addr)
			}
		}
	}

	if exception.ExceptionCode == windows.EXCEPTION_SINGLE_STEP {
		d.state = StatePaused
		isBreakpoint = true
	}

	if isBreakpoint {
		d.state = StatePaused
	}
}

func (d *Debugger) Continue() error {
	if d.processHandle == 0 {
		return nil
	}

	threads := d.threads.Self().(*thread.Manager).GetAllThreads()
	if len(threads) == 0 {
		threads = []*thread.Thread{{Id: d.mainThreadId}}
	}

	breakpoints := d.breakpoints.Self().(*breakpoint.Manager).GetAllBreakpoints()
	hasHardwareBp := false
	for _, bp := range breakpoints {
		if bp.Enabled && !bp.Active && bp.Type == breakpoint.BreakpointTypeHardware {
			hasHardwareBp = true
			break
		}
	}

	if hasHardwareBp {
		for _, t := range threads {
			threadHandle, err := windows.OpenThread(windows.THREAD_GET_CONTEXT|windows.THREAD_SET_CONTEXT, false, t.Id)
			if err != nil {
				continue
			}

			ctx, err := windows.GetThreadContext(threadHandle)
			if err != nil {
				windows.CloseHandle(threadHandle)
				continue
			}

			for _, bp := range breakpoints {
				if bp.Enabled && !bp.Active && bp.Type == breakpoint.BreakpointTypeHardware {
					d.breakpoints.Self().(*breakpoint.Manager).ApplyHardwareBreakpoint(ctx, bp)
				}
			}

			windows.SetThreadContext(threadHandle, ctx)
			windows.CloseHandle(threadHandle)
		}
	}

	for _, bp := range breakpoints {
		if bp.Enabled && !bp.Active && bp.Type == breakpoint.BreakpointTypeSoftware {
			d.breakpoints.Self().(*breakpoint.Manager).ApplySoftwareBreakpoint(d.processHandle, bp)
		}
	}

	d.continueChan <- struct{}{}
	return nil
}

func (d *Debugger) StepOver() error {
	if d.processHandle == 0 {
		return nil
	}

	threadId := d.mainThreadId
	if d.lastEvent != nil {
		threadId = d.lastEvent.ThreadId()
	}

	threadHandle, err := windows.OpenThread(windows.THREAD_GET_CONTEXT|windows.THREAD_SET_CONTEXT|windows.THREAD_SUSPEND_RESUME, false, threadId)
	if err != nil {
		return fmt.Errorf("OpenThread failed: %v", err)
	}
	defer windows.CloseHandle(threadHandle)

	_, err = windows.SuspendThread(threadHandle)
	if err != nil {
		return fmt.Errorf("SuspendThread failed: %v", err)
	}

	ctx, err := windows.GetThreadContext(threadHandle)
	if err != nil {
		return fmt.Errorf("GetThreadContext failed: %v", err)
	}

	instruction, err := d.memory.Self().(*memory.Manager).ReadMemory(ctx.Rip, 15)
	if err != nil || len(instruction) == 0 {
		windows.ResumeThread(threadHandle)
		return fmt.Errorf("ReadMemory failed: %v", err)
	}

	if isCallInstruction(instruction[0]) {
		bpAddr := ctx.Rip + uint64(len(instruction))
		bp, err := d.breakpoints.Self().(*breakpoint.Manager).SetSoftwareBreakpoint(bpAddr, true)
		if err == nil {
			bp.SingleShot = true
		}
	} else {
		ctx.EFlags |= 0x100
		ctx.EFlags &^= 0x10000
		err = windows.SetThreadContext(threadHandle, ctx)
		if err != nil {
			windows.ResumeThread(threadHandle)
			return fmt.Errorf("SetThreadContext failed: %v", err)
		}
	}

	_, err = windows.ResumeThread(threadHandle)
	if err != nil {
		return fmt.Errorf("ResumeThread failed: %v", err)
	}

	d.continueChan <- struct{}{}
	d.state = StateRunning
	return nil
}

func (d *Debugger) StepInto() error {
	if d.processHandle == 0 {
		return nil
	}

	var exitCode uint32
	err := syscall.GetExitCodeProcess(syscall.Handle(d.processHandle), &exitCode)
	if err != nil {
		return fmt.Errorf("GetExitCodeProcess failed: %v", err)
	}
	if exitCode != 259 {
		return fmt.Errorf("进程已经退出，退出代码=%d", exitCode)
	}

	threadId := d.mainThreadId
	if d.lastEvent != nil {
		threadId = d.lastEvent.ThreadId()
	}

	threadHandle, err := windows.OpenThread(windows.THREAD_GET_CONTEXT|windows.THREAD_SET_CONTEXT|windows.THREAD_SUSPEND_RESUME, false, threadId)
	if err != nil {
		return fmt.Errorf("OpenThread failed: %v", err)
	}
	defer windows.CloseHandle(threadHandle)

	_, err = windows.SuspendThread(threadHandle)
	if err != nil {
		return fmt.Errorf("SuspendThread failed: %v", err)
	}

	ctx, err := windows.GetThreadContext(threadHandle)
	if err != nil {
		return fmt.Errorf("GetThreadContext failed: %v", err)
	}

	ctx.EFlags |= 0x100
	ctx.EFlags &^= 0x10000

	err = windows.SetThreadContext(threadHandle, ctx)
	if err != nil {
		return fmt.Errorf("SetThreadContext failed: %v", err)
	}

	_, err = windows.ResumeThread(threadHandle)
	if err != nil {
		return fmt.Errorf("ResumeThread failed: %v", err)
	}

	d.pendingStepInto = true
	d.continueChan <- struct{}{}
	d.state = StateRunning
	return nil
}

func isCallInstruction(b byte) bool {
	return b == 0xE8 || b == 0xFF || b == 0x9A
}
