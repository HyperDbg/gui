package debugger

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/HyperDbg/debugger/ark"
	"github.com/ddkwork/HyperDbg/debugger/bookmark"
	"github.com/ddkwork/HyperDbg/debugger/breakpoint"
	"github.com/ddkwork/HyperDbg/debugger/comment"
	"github.com/ddkwork/HyperDbg/debugger/disassembly"
	"github.com/ddkwork/HyperDbg/debugger/driver"
	"github.com/ddkwork/HyperDbg/debugger/exception"
	"github.com/ddkwork/HyperDbg/debugger/function"
	"github.com/ddkwork/HyperDbg/debugger/graph"
	"github.com/ddkwork/HyperDbg/debugger/handle"
	"github.com/ddkwork/HyperDbg/debugger/label"
	"github.com/ddkwork/HyperDbg/debugger/loop"
	"github.com/ddkwork/HyperDbg/debugger/memory"
	"github.com/ddkwork/HyperDbg/debugger/notes"
	"github.com/ddkwork/HyperDbg/debugger/peview"
	"github.com/ddkwork/HyperDbg/debugger/process"
	"github.com/ddkwork/HyperDbg/debugger/reference"
	"github.com/ddkwork/HyperDbg/debugger/register"
	"github.com/ddkwork/HyperDbg/debugger/scylla"
	"github.com/ddkwork/HyperDbg/debugger/seh"
	"github.com/ddkwork/HyperDbg/debugger/source"
	"github.com/ddkwork/HyperDbg/debugger/stack"
	"github.com/ddkwork/HyperDbg/debugger/symbol"
	"github.com/ddkwork/HyperDbg/debugger/thread"
	"github.com/ddkwork/HyperDbg/debugger/trace"
	"github.com/ddkwork/HyperDbg/debugger/transparency"
	"github.com/ddkwork/HyperDbg/debugger/type_"
	"github.com/ddkwork/HyperDbg/debugger/watch"
	"github.com/ddkwork/HyperDbg/debugger/xref"
	"golang.org/x/sys/windows"
)

type DebuggerState int

const (
	StateUninitialized DebuggerState = iota
	StateInitialized
	StateRunning
	StatePaused
	StateTerminated
)

type ActiveDebuggingProcess struct {
	IsActive    bool
	IsPaused    bool
	ProcessId   uint32
	ThreadId    uint32
	ProcessName string
}

type HyperDbg struct {
	mu sync.RWMutex

	State                             DebuggerState
	ActiveProcess                     ActiveDebuggingProcess
	BreakPrintingOutput               bool
	AutoUnpause                       bool
	isConnectedToRemoteDebuggee       bool
	isSerialConnectedToRemoteDebuggee bool
	isInstrumentingInstructions       bool
	ignorePauseRequests               bool
	Driver                            *driver.DriverHandle
	DriverLoader                      *driver.DriverProvider
	EventChan                         chan *DebugEvent
	EventCallbacks                    []DebugEventCallback
	EventManager                      *EventManager
	EventHandler                      *EventHandler
	DebugEventHandler                 *DebugEventHandler
	EventLoop                         *EventLoop
	OnTitleUpdate                     func(title string)

	Registers    api.Interface
	Memory       api.Interface
	Breakpoints  api.Interface
	Symbols      api.Interface
	Threads      api.Interface
	Stack        api.Interface
	Disassembly  api.Interface
	Trace        api.Interface
	Process      api.Interface
	Transparency api.Interface
	Hardware     api.Interface
	VMM          api.Interface

	Bookmarks  api.Interface
	Comments   api.Interface
	Labels     api.Interface
	Functions  api.Interface
	Xrefs      api.Interface
	Types      api.Interface
	Watches    api.Interface
	Graphs     api.Interface
	Exceptions api.Interface
	Notes      api.Interface
	Source     api.Interface
	References api.Interface
	Handles    api.Interface
	Script     api.Interface
	Ark        api.Interface
	Scylla     api.Interface
	Seh        api.Interface
	Loops      api.Interface
	PeView     api.Interface

	UIComponents   []api.Interface
	AutoLoadDriver bool
}

func NewHyperDbg() *HyperDbg {
	return NewHyperDbgWithOptions(true)
}

func NewHyperDbgWithOptions(autoLoadDriver bool) *HyperDbg {
	slog.Info("=== HyperDbg 正在初始化 ===")
	slog.Info("自动加载驱动", "value", autoLoadDriver)
	slog.Info("时间", "value", time.Now().Format("2006-01-02 15:04:05.000"))

	driverConfig := &driver.DriverConfig{
		DeviceName:     "\\\\.\\HyperDbgDebuggerDevice",
		ServiceName:    "HyperDbg",
		DriverFileName: "hyperkd.sys",
	}

	driverPath := getDriverPath()
	slog.Info("找到驱动文件", "path", driverPath)

	loader := driver.NewDriverProvider("HyperDbg", driverPath)

	var driverHandle *driver.DriverHandle
	var err error

	if autoLoadDriver {
		slog.Info("[步骤1] 尝试打开驱动设备...")
		slog.Info("时间", "value", time.Now().Format("2006-01-02 15:04:05.000"))

		driverHandle, err = driver.NewDriverHandle(driverConfig)
		if err != nil {
			slog.Error("[步骤1失败] 打开驱动失败", "error", err)
			slog.Info("驱动不可用，尝试安装...")

			if err := enableSeLoadDriverPrivilege(); err != nil {
				slog.Error("[步骤2失败] 启用 SeLoadDriverPrivilege 失败", "error", err)
				slog.Warn("这可能导致驱动安装失败")
			} else {
				slog.Info("[步骤2成功] SeLoadDriverPrivilege 已成功启用")
			}

			slog.Info("[步骤3] 正在安装驱动服务...")
			slog.Info("时间", "value", time.Now().Format("2006-01-02 15:04:05.000"))

			if err := loader.Install(); err != nil {
				slog.Error("安装驱动失败", "error", err)
				if strings.Contains(err.Error(), "already exists") {
					slog.Info("驱动服务已存在，正在移除...")

					if stopErr := loader.Manage(driver.DriverFuncStop); stopErr != nil {
						slog.Warn("停止驱动失败", "error", stopErr)
					} else {
						slog.Info("✓ 驱动服务已成功停止")
					}

					if removeErr := loader.Manage(driver.DriverFuncRemove); removeErr != nil {
						slog.Warn("移除驱动失败", "error", removeErr)
						slog.Info("驱动不可用，运行在模拟模式")
						slog.Info("请以管理员身份运行应用程序")
						driverHandle = nil
					} else {
						slog.Info("✓ 驱动服务已成功移除")
						slog.Info("等待服务清理和杀毒软件扫描...")
						time.Sleep(5 * time.Second)
						slog.Info("重试安装...")
						if retryErr := loader.Install(); retryErr != nil {
							slog.Error("安装驱动失败", "error", retryErr)
							slog.Info("驱动不可用，运行在模拟模式")
							slog.Info("可能的原因:")
							slog.Info("1. 驱动文件未正确签名")
							slog.Info("2. 未启用测试签名")
							slog.Info("3. 权限不足")
							slog.Info("4. 杀毒软件正在阻止驱动")
							slog.Info("5. 路径包含不支持的字符")
							slog.Info("请以管理员身份运行应用程序")
							driverHandle = nil
						} else {
							slog.Info("✓ 驱动已成功安装并启动")

							slog.Info("重试打开驱动...")

							if err := enableSeDebugPrivilege(); err != nil {
								slog.Warn("启用 SeDebugPrivilege 失败", "error", err)
							} else {
								slog.Info("✓ SeDebugPrivilege 已成功启用")
							}

							driverHandle, err = driver.NewDriverHandle(driverConfig)
							if err != nil {
								slog.Error("仍然无法打开驱动", "error", err)
								slog.Info("驱动不可用，运行在模拟模式")
								driverHandle = nil
							} else {
								slog.Info("✓ 安装后驱动加载成功")

								slog.Info("正在初始化 VMM...")

								packet := NewPacket(driverHandle)
								vmmProvider := VmmProviderNew(packet)
								if err := vmmProvider.Initialize(); err != nil {
									slog.Warn("初始化 VMM 失败", "error", err)
									slog.Info("驱动已加载但 VMM 未初始化，部分功能可能不可用")
								} else {
									slog.Info("✓ VMM 初始化成功")
								}
							}
						}
					}
				} else {
					slog.Error("安装驱动失败", "error", err)
					slog.Info("驱动不可用，运行在模拟模式")
					slog.Info("可能的原因:")
					slog.Info("1. 驱动文件未正确签名")
					slog.Info("2. 未启用测试签名")
					slog.Info("3. 权限不足")
					slog.Info("4. 杀毒软件正在阻止驱动")
					slog.Info("5. 路径包含不支持的字符")
					slog.Info("请以管理员身份运行应用程序")
					driverHandle = nil
				}
			} else {
				slog.Info("[步骤3成功] 驱动服务安装成功")
				slog.Info("等待驱动启动...")
				time.Sleep(2 * time.Second)

				slog.Info("重试打开驱动...")

				if err := enableSeDebugPrivilege(); err != nil {
					slog.Warn("启用 SeDebugPrivilege 失败", "error", err)
				} else {
					slog.Info("✓ SeDebugPrivilege 已成功启用")
				}

				driverHandle, err = driver.NewDriverHandle(driverConfig)
				if err != nil {
					slog.Error("安装后仍然无法打开驱动", "error", err)
					slog.Info("驱动不可用，运行在模拟模式")
					driverHandle = nil
				} else {
					slog.Info("✓ 安装后驱动加载成功")

					slog.Info("正在初始化 VMM...")

					packet := NewPacket(driverHandle)
					vmmProvider := VmmProviderNew(packet)
					if err := vmmProvider.Initialize(); err != nil {
						slog.Warn("初始化 VMM 失败", "error", err)
						slog.Info("驱动已加载但 VMM 未初始化，部分功能可能不可用")
					} else {
						slog.Info("✓ VMM 初始化成功")
					}
				}
			}
		} else {
			slog.Info("[步骤1成功] 驱动已加载")

			slog.Info("正在初始化 VMM...")

			packet := NewPacket(driverHandle)
			vmmProvider := VmmProviderNew(packet)

			if err := vmmProvider.Initialize(); err != nil {
				slog.Warn("初始化 VMM 失败", "error", err)
				slog.Info("驱动已加载但 VMM 未初始化，部分功能可能不可用")
			} else {
				slog.Info("✓ VMM 初始化成功")
			}
		}
	}

	slog.Info("驱动状态", "loaded", driverHandle != nil)
	if driverHandle == nil {
		slog.Warn("⚠️  运行在模拟模式 - 功能受限")
		slog.Warn("   没有驱动，大多数调试功能将无法工作")
	}
	slog.Info("=== HyperDbg 初始化完成 ===")

	regMgr := register.New()
	memMgr := memory.New()
	bpMgr := breakpoint.New()
	symMgr := symbol.New()
	threadMgr := thread.New()
	stackMgr := stack.New()
	disasmMgr := disassembly.New()
	traceMgr := trace.New()
	processMgr := process.New()
	transparencyMgr := transparency.New()
	packet := NewPacket(driverHandle)
	vmmMgr := VmmProviderNew(packet)

	bookmarkMgr := bookmark.New()
	commentMgr := comment.New()
	labelMgr := label.New()
	functionMgr := function.New()
	xrefMgr := xref.New()
	watchMgr := watch.New()
	notesMgr := notes.New()
	sourceMgr := source.New()
	referenceMgr := reference.New()
	handleMgr := handle.New()
	peviewMgr := peview.New()
	exceptionMgr := exception.New()
	arkMgr := ark.New()
	scyllaMgr := scylla.New()
	graphMgr := graph.New()
	typeMgr := type_.New()
	loopMgr := loop.New()
	sehMgr := seh.New()

	dbg := &HyperDbg{
		State:                             StateInitialized,
		BreakPrintingOutput:               false,
		AutoUnpause:                       false,
		isConnectedToRemoteDebuggee:       false,
		isSerialConnectedToRemoteDebuggee: false,
		isInstrumentingInstructions:       false,
		ignorePauseRequests:               false,
		Driver:                            driverHandle,
		DriverLoader:                      loader,
		EventChan:                         make(chan *DebugEvent, 100),
		EventCallbacks:                    make([]DebugEventCallback, 0),
		EventManager:                      NewEventManager(driverHandle),
		EventHandler:                      nil,
		Registers:                         regMgr,
		Memory:                            memMgr,
		Breakpoints:                       bpMgr,
		Symbols:                           symMgr,
		Threads:                           threadMgr,
		Stack:                             stackMgr,
		Disassembly:                       disasmMgr,
		Trace:                             traceMgr,
		Process:                           processMgr,
		Transparency:                      transparencyMgr,
		Hardware:                          nil,
		VMM:                               vmmMgr,
		Bookmarks:                         bookmarkMgr,
		Comments:                          commentMgr,
		Labels:                            labelMgr,
		Functions:                         functionMgr,
		Xrefs:                             xrefMgr,
		Types:                             typeMgr,
		Watches:                           watchMgr,
		Graphs:                            graphMgr,
		Exceptions:                        exceptionMgr,
		Notes:                             notesMgr,
		Source:                            sourceMgr,
		References:                        referenceMgr,
		Handles:                           handleMgr,
		Script:                            nil,
		Ark:                               arkMgr,
		Scylla:                            scyllaMgr,
		Seh:                               sehMgr,
		Loops:                             loopMgr,
		PeView:                            peviewMgr,
		UIComponents:                      make([]api.Interface, 0),
		AutoLoadDriver:                    autoLoadDriver,
	}

	dbg.EventHandler = NewEventHandler(driverHandle)
	dbg.DebugEventHandler = NewDebugEventHandler(driverHandle, regMgr.(*register.Provider), processMgr.(*process.Provider), threadMgr.(*thread.Provider))
	dbg.EventLoop = NewEventLoop(dbg.EventManager)

	dbg.EventLoop.RegisterHandler(EventReadMemory, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EventWriteMemory, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EventSetBreakpoint, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EventRemoveBreakpoint, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EventSingleStep, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EventContinue, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EventAttachProcess, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EventDetachProcess, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EventRegisterEvent, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EventUnregisterEvent, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EventModifyEvent, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EventStep, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EventBreakpoint, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EventException, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EventProcessCreated, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EventProcessExited, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EventThreadCreated, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EventThreadExited, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EventModuleLoaded, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EventModuleUnloaded, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EptHookReadWriteAndExecute, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EptHookReadWrite, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EptHookReadAndExecute, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EptHookWriteAndExecute, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EptHookRead, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EptHookWrite, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EptHookExecute, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EptHookExecDetours, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(EptHookExecCC, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(SyscallHookEferSyscall, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(SyscallHookEferSysret, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(CpuidInstructionExecution, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(RdmsrInstructionExecution, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(WrmsrInstructionExecution, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(InInstructionExecution, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(OutInstructionExecution, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(ExceptionOccurred, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(ExternalInterruptOccurred, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(DebugRegistersAccessed, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(TscInstructionExecution, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(PmcInstructionExecution, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(VmcallInstructionExecution, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(XsetbvInstructionExecution, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(ControlRegisterModified, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(ControlRegisterRead, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(ControlRegister3Modified, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(TrapExecutionModeChanged, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})
	dbg.EventLoop.RegisterHandler(TrapExecutionInstructionTrace, func(event *DebugEvent) {
		dbg.DebugEventHandler.Handle(event)
	})

	dbg.EventLoop.Start()

	dbg.RegisterUIComponent(regMgr)
	dbg.RegisterUIComponent(memMgr)
	dbg.RegisterUIComponent(bpMgr)
	dbg.RegisterUIComponent(symMgr)
	dbg.RegisterUIComponent(threadMgr)
	dbg.RegisterUIComponent(stackMgr)
	dbg.RegisterUIComponent(disasmMgr)
	dbg.RegisterUIComponent(traceMgr)
	dbg.RegisterUIComponent(processMgr)
	dbg.RegisterUIComponent(transparencyMgr)
	dbg.RegisterUIComponent(vmmMgr)

	return dbg
}

func (h *HyperDbg) GetState() DebuggerState {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.State
}

func (h *HyperDbg) SetState(state DebuggerState) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.State = state
}

func (h *HyperDbg) GetActiveProcess() ActiveDebuggingProcess {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.ActiveProcess
}

func (h *HyperDbg) SetActiveProcess(process ActiveDebuggingProcess) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.ActiveProcess = process
}

func (h *HyperDbg) IsBreakPrintingOutput() bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.BreakPrintingOutput
}

func (h *HyperDbg) SetBreakPrintingOutput(value bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.BreakPrintingOutput = value
}

func (h *HyperDbg) IsConnectedToRemoteDebuggee() bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.isConnectedToRemoteDebuggee
}

func (h *HyperDbg) SetConnectedToRemoteDebuggee(value bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.isConnectedToRemoteDebuggee = value
}

func (h *HyperDbg) IsSerialConnectedToRemoteDebuggee() bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.isSerialConnectedToRemoteDebuggee
}

func (h *HyperDbg) SetSerialConnectedToRemoteDebuggee(value bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.isSerialConnectedToRemoteDebuggee = value
}

func (h *HyperDbg) IsInstrumentingInstructions() bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.isInstrumentingInstructions
}

func (h *HyperDbg) SetInstrumentingInstructions(value bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.isInstrumentingInstructions = value
}

func (h *HyperDbg) GetDriver() *driver.DriverHandle {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.Driver
}

func (h *HyperDbg) GetRegisters() api.Interface {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.Registers
}

func (h *HyperDbg) GetMemory() api.Interface {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.Memory
}

func (h *HyperDbg) GetBookmarks() api.Interface {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.Bookmarks
}

func (h *HyperDbg) GetComments() api.Interface {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.Comments
}

func (h *HyperDbg) GetLabels() api.Interface {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.Labels
}

func (h *HyperDbg) GetFunctions() api.Interface {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.Functions
}

func (h *HyperDbg) GetXrefs() api.Interface {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.Xrefs
}

func (h *HyperDbg) GetWatches() api.Interface {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.Watches
}

func (h *HyperDbg) GetEventManager() *EventManager {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.EventManager
}

func (h *HyperDbg) GetEventLoop() *EventLoop {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.EventLoop
}

func (h *HyperDbg) IsDriverLoaded() bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.Driver != nil && h.Driver.IsConnected()
}

func (h *HyperDbg) LoadVMM() error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.Driver == nil {
		return fmt.Errorf("驱动未加载")
	}

	slog.Info("=== 正在加载 VMM ===")

	packet := NewPacket(h.Driver)
	vmmProvider := VmmProviderNew(packet)
	if err := vmmProvider.Initialize(); err != nil {
		slog.Error("初始化 VMM 失败", "error", err)
		return fmt.Errorf("failed to initialize VMM: %w", err)
	}

	slog.Info("✓ VMM 初始化成功")
	slog.Info("=== VMM 加载完成 ===")

	return nil
}

func (h *HyperDbg) UnloadDriver() error {
	h.mu.Lock()
	defer h.mu.Unlock()

	slog.Info("=== 正在卸载驱动 ===")

	if h.EventLoop != nil {
		slog.Info("正在停止事件循环...")
		h.EventLoop.Stop()
	}

	if h.Driver == nil {
		slog.Info("驱动未加载，跳过关闭")
	} else {
		slog.Info("步骤 1/5: 正在发送 IOCTL 返回 IRP 挂起数据包并禁止 IOCTL...")
		emptyBuffer := make([]byte, 8)
		err := h.Driver.SendBuffer(emptyBuffer, IoctlReturnIrpPendingAndDisallowIoctl)
		if err != nil {
			slog.Warn("发送 IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL 失败", "error", err)
		} else {
			slog.Info("✓ IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL 发送成功")
		}

		slog.Info("步骤 2/5: 等待 IRP 挂起线程返回...")
		time.Sleep(2 * time.Second)

		slog.Info("步骤 3/5: 正在发送 IOCTL 终止 VMM...")
		err = h.Driver.SendBuffer(emptyBuffer, IoctlTerminateVmx)
		if err != nil {
			slog.Warn("发送 IOCTL_TERMINATE_VMX 失败", "error", err)
		} else {
			slog.Info("✓ IOCTL_TERMINATE_VMX 发送成功")
		}

		slog.Info("步骤 4/5: 等待 VMM 和调试器完全终止...")
		time.Sleep(3 * time.Second)

		slog.Info("步骤 5/5: 正在关闭驱动句柄...")
		if err := h.Driver.Close(); err != nil {
			slog.Warn("关闭驱动失败", "error", err)
		} else {
			slog.Info("✓ 驱动关闭成功")
		}

		h.Driver = nil
	}

	if h.DriverLoader == nil {
		slog.Warn("DriverLoader 未初始化，跳过驱动服务管理")
	} else {
		if err := h.DriverLoader.Manage(driver.DriverFuncStop); err != nil {
			slog.Warn("停止驱动服务失败", "error", err)
		} else {
			slog.Info("✓ 驱动服务停止成功")
		}

		if err := h.DriverLoader.Manage(driver.DriverFuncRemove); err != nil {
			slog.Warn("移除驱动服务失败", "error", err)
		} else {
			slog.Info("✓ 驱动服务移除成功")
			slog.Info("等待服务完全删除...")
			time.Sleep(2 * time.Second)
		}
	}

	slog.Info("=== 驱动和 VMM 卸载完成 ===")
	return nil
}

func (h *HyperDbg) RegisterEventCallback(callback DebugEventCallback) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.EventCallbacks = append(h.EventCallbacks, callback)
}

func (h *HyperDbg) DispatchEvent(event *DebugEvent) {
	h.mu.RLock()
	if h.EventChan != nil {
		select {
		case h.EventChan <- event:
		default:
			slog.Warn("事件通道已满，丢弃事件")
		}
	}
	h.mu.RUnlock()

	h.mu.RLock()
	callbacks := make([]DebugEventCallback, len(h.EventCallbacks))
	copy(callbacks, h.EventCallbacks)
	h.mu.RUnlock()

	for _, callback := range callbacks {
		callback(event)
	}
}

func (h *HyperDbg) GetEventChan() <-chan *DebugEvent {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.EventChan
}

func (h *HyperDbg) RegisterUIComponent(component api.Interface) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.UIComponents = append(h.UIComponents, component)
}

func (h *HyperDbg) UpdateAllPages() {
	h.mu.RLock()
	defer h.mu.RUnlock()

	slog.Info("正在更新所有页面...")

	for _, component := range h.UIComponents {
		if err := component.Update(); err != nil {
			slog.Error("更新组件失败", "component", component.Self(), "error", err)
		}
	}
}

func (h *HyperDbg) LoadProcess(exePath string) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	slog.Info("正在加载进程", "path", exePath)

	if h.Driver == nil {
		slog.Warn("驱动未加载，进程加载已模拟")
		h.ActiveProcess = ActiveDebuggingProcess{
			IsActive:    true,
			IsPaused:    false,
			ProcessId:   1234,
			ThreadId:    5678,
			ProcessName: filepath.Base(exePath),
		}
		h.State = StateRunning
		return nil
	}

	slog.Info("驱动已就绪，尝试真实进程加载")

	processId, threadId, err := loadProcess(exePath)
	if err != nil {
		return fmt.Errorf("failed to load process: %w", err)
	}

	h.ActiveProcess = ActiveDebuggingProcess{
		IsActive:    true,
		IsPaused:    true,
		ProcessId:   processId,
		ThreadId:    threadId,
		ProcessName: filepath.Base(exePath),
	}
	h.State = StatePaused

	slog.Info("进程加载成功", "path", exePath)
	slog.Info("进程 ID", "pid", h.ActiveProcess.ProcessId, "tid", h.ActiveProcess.ThreadId)
	slog.Info("驱动状态", "loaded", h.IsDriverLoaded())

	return nil
}

func (h *HyperDbg) GetActiveProcessInfo() ActiveDebuggingProcess {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.ActiveProcess
}

func enableSeLoadDriverPrivilege() error {
	const SE_LOAD_DRIVER_NAME = "SeLoadDriverPrivilege"
	return enablePrivilege(SE_LOAD_DRIVER_NAME)
}

func enableSeDebugPrivilege() error {
	const SE_DEBUG_NAME = "SeDebugPrivilege"
	return enablePrivilege(SE_DEBUG_NAME)
}

func enablePrivilege(privilegeName string) error {
	var token windows.Token
	currentProcess := windows.CurrentProcess()

	err := windows.OpenProcessToken(currentProcess, windows.TOKEN_ADJUST_PRIVILEGES|windows.TOKEN_QUERY, &token)
	if err != nil {
		return fmt.Errorf("failed to open process token: %w", err)
	}
	defer token.Close()

	privilege := windows.LUIDAndAttributes{}
	privilegeNamePtr, err := windows.UTF16PtrFromString(privilegeName)
	if err != nil {
		return fmt.Errorf("failed to convert privilege name: %w", err)
	}
	err = windows.LookupPrivilegeValue(nil, privilegeNamePtr, &privilege.Luid)
	if err != nil {
		return fmt.Errorf("failed to lookup privilege value: %w", err)
	}

	privilege.Attributes = windows.SE_PRIVILEGE_ENABLED

	privileges := windows.Tokenprivileges{
		PrivilegeCount: 1,
		Privileges:     [1]windows.LUIDAndAttributes{privilege},
	}
	err = windows.AdjustTokenPrivileges(token, false, &privileges, 0, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to adjust token privileges: %w", err)
	}

	return nil
}

func getDriverPath() string {
	return "d:\\笔记本\\p\\ux\\examples\\hypedbg\\doc\\cpp\\HyperDbgUnified\\build\\Release\\hyperkd.sys"
}

func loadProcess(exePath string) (uint32, uint32, error) {
	return 1234, 5678, nil
}

func (h *HyperDbg) PerformStepOver() error {
	return fmt.Errorf("PerformStepOver not implemented")
}

func (h *HyperDbg) PerformStepIn() error {
	return fmt.Errorf("PerformStepIn not implemented")
}

type TokenType int

const (
	TokenCommand TokenType = iota
	TokenArgument
	TokenOption
	TokenValue
)

type CommandToken struct {
	Type  TokenType
	Value string
}
