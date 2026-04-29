package kernel_level

import (
	"encoding/binary"
	"fmt"
	"sync"
	"unsafe"

	"github.com/ddkwork/sdk"
)

type KernelSyncObjectState struct {
	IsSignaled bool
	mu         sync.Mutex
}

const MaximumSyncObjects = 256

type KernelListeningManager struct {
	syncObjects                             [MaximumSyncObjects]KernelSyncObjectState
	isSerialConnectedToRemoteDebuggee       bool
	isDebuggeeRunning                       bool
	ignoreNewLoggingMessages                bool
	currentRemoteCore                       uint32
	kernelBaseAddress                       uint64
	currentRunningInstruction               [sdk.MaximumInstrSize]byte
	isRunningInstruction32Bit               bool
	sharedEventStatus                       bool
	resultOfEvaluatedExpression             uint64
	errorStateOfResultOfEvaluatedExpression uint32
	debuggeeResultOfRegisteringEvent        sdk.DEBUGGER_EVENT_AND_ACTION_RESULT
	debuggeeResultOfAddingActionsToEvent    sdk.DEBUGGER_EVENT_AND_ACTION_RESULT
	mu                                      sync.RWMutex
	kd                                      *KernelDebugger
}

func NewKernelListeningManager(kd *KernelDebugger) *KernelListeningManager {
	return &KernelListeningManager{
		syncObjects: [MaximumSyncObjects]KernelSyncObjectState{},
		kd:          kd,
	}
}

func (klm *KernelListeningManager) ListeningSerialPortInDebugger() bool {
	for {
		bufferToReceive := make([]byte, int(sdk.MaxSerialPacketSize))
		var lengthReceived uint32

		received, err := klm.kdReceivePacketFromDebuggee(bufferToReceive, &lengthReceived)
		if err != nil {
			fmt.Printf("err, error receiving packet: %v\n", err)
			continue
		}

		if !received {
			if lengthReceived == 0 && bufferToReceive[0] == 0 {
				fmt.Println("the remote connection is closed")

				if klm.isSerialConnectedToRemoteDebuggee {
					klm.isSerialConnectedToRemoteDebuggee = false
				}

				klm.kdCloseConnection()
				return false
			}
			fmt.Println("err, invalid buffer received")
			continue
		}

		if lengthReceived == 1 && bufferToReceive[0] == 0 {
			continue
		}

		if len(bufferToReceive) < int(lengthReceived) {
			continue
		}

		packet := bufferToReceive[:lengthReceived]

		if len(packet) < 24 {
			continue
		}

		indicator := binary.LittleEndian.Uint64(packet[8:16])
		if indicator != sdk.IndicatorOfHyperdbgPacket {
			continue
		}

		expectedChecksum := ComputeDataChecksum(packet[1:])
		actualChecksum := packet[0]
		if expectedChecksum != actualChecksum {
			fmt.Println("\nerr, checksum is invalid")
			continue
		}

		packetType := binary.LittleEndian.Uint32(packet[16:20])
		if packetType != 3 {
			fmt.Println("\nerr, unknown packet received from the debuggee")
			continue
		}

		requestedAction := binary.LittleEndian.Uint32(packet[20:24])
		payload := packet[24:]

		klm.handleRequestedAction(sdk.DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION(requestedAction), payload)
	}
}

func (klm *KernelListeningManager) handleRequestedAction(action sdk.DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION, payload []byte) {
	switch action {
	case sdk.DebuggerRemotePacketPingAndSendSupportedVersion:
		klm.handlePingPacket()

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeStarted:
		klm.handleDebuggeeStartedPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeLoggingMechanism:
		klm.handleLoggingPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeePausedAndCurrentInstruction:
		klm.handlePausedPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfChangingCore:
		klm.handleChangeCoreResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfChangingProcess:
		klm.handleChangeProcessResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeReloadSearchQuery:
		klm.handleSearchResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfChangingThread:
		klm.handleChangeThreadResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfFlush:
		klm.handleFlushResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfCallstack:
		klm.handleCallstackResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultTestQuery:
		klm.handleTestQueryResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfRunningScript:
		klm.handleScriptResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfFormats:
		klm.handleFormatsResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfRegisteringEvent:
		klm.handleRegisterEventResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfAddingActionToEvent:
		klm.handleAddActionToEventResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfQueryAndModifyEvent:
		klm.handleQueryAndModifyEventResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeReloadSymbolFinished:
		klm.handleSymbolReloadFinishedPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfReadingRegisters:
		klm.handleReadRegistersResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfWriteRegister:
		klm.handleWriteRegisterResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfApicRequests:
		klm.handleApicResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfQueryIdtEntriesRequests:
		klm.handleIdtEntriesResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfReadingMemory:
		klm.handleReadMemoryResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfEditingMemory:
		klm.handleEditMemoryResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfBp:
		klm.handleBpResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfShortCircuitingState:
		klm.handleShortCircuitingResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfPte:
		klm.handlePteResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultSmiOperationRequests:
		klm.handleSmiOperationResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultHypertraceLbrOperationRequests:
		klm.handleHyperTraceLbrOperationResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultHypertracePtOperationRequests:
		klm.handleHyperTracePtOperationResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfBringingPagesIn:
		klm.handlePageInResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfVa2paAndPa2va:
		klm.handleVa2paPa2vaResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfListOrModifyBreakpoints:
		klm.handleListOrModifyBreakpointsResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeUpdateSymbolInfo:
		klm.handleUpdateSymbolInfoPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfPcitree:
		klm.handlePcitreeResultPacket(payload)

	case sdk.DebuggerRemotePacketRequestedActionDebuggeeResultOfPcidevinfo:
		klm.handlePcidevinfoResultPacket(payload)

	default:
		fmt.Println("err, unknown packet action received from the debuggee")
	}
}

func (klm *KernelListeningManager) handlePingPacket() {
	fmt.Println("ping received, sending response...")
	klm.kdSendResponseOfPingPacket()
}

func (klm *KernelListeningManager) handleDebuggeeStartedPacket(payload []byte) {
	if len(payload) < 8+int(sdk.MaximumCharacterForOsName) {
		return
	}

	initPacket := sdk.DEBUGGER_PREPARE_DEBUGGEE{}
	initPacket.KernelBaseAddress = binary.LittleEndian.Uint64(payload[0:8])
	osNameBytes := payload[8 : 8+sdk.MaximumCharacterForOsName]
	for i := range osNameBytes {
		initPacket.OsName[i] = int8(osNameBytes[i])
	}

	klm.mu.Lock()
	klm.kernelBaseAddress = initPacket.KernelBaseAddress
	klm.mu.Unlock()

	osNameByteSlice := make([]byte, len(initPacket.OsName))
	for i, v := range initPacket.OsName {
		osNameByteSlice[i] = byte(v)
	}
	osName := string(osNameByteSlice)
	nullIdx := 0
	for i, b := range osName {
		if b == 0 {
			nullIdx = i
			break
		}
	}
	if nullIdx > 0 {
		osName = osName[:nullIdx]
	}

	fmt.Printf("connected to debuggee %s\n", osName)

	klm.signalEvent(int(SyncObjectKernelDebuggerStartedPacketReceived))
}

func (klm *KernelListeningManager) handleLoggingPacket(payload []byte) {
	if !klm.ignoreNewLoggingMessages {
		message := extractMessageFromPayload(payload)
		fmt.Printf("%s", message)
	}
}

func (klm *KernelListeningManager) handlePausedPacket(payload []byte) {
	klm.mu.Lock()
	klm.ignoreNewLoggingMessages = true
	klm.isDebuggeeRunning = false
	klm.mu.Unlock()

	pausePacket := parsePausePacket(payload)

	klm.mu.Lock()
	klm.currentRemoteCore = pausePacket.CurrentCore
	copy(klm.currentRunningInstruction[:], pausePacket.InstructionBytesOnRip[:])
	klm.isRunningInstruction32Bit = pausePacket.IsProcessorOn32BitMode
	klm.mu.Unlock()

	switch pausePacket.PausingReason {
	case sdk.DebuggeePausingReasonDebuggeeSoftwareBreakpointHit:
		if pausePacket.EventTag != 0 {
			fmt.Printf("breakpoint 0x%x hit\n", pausePacket.EventTag)
		}

	case sdk.DebuggeePausingReasonDebuggeeHardwareDebugRegisterHit:
		break

	case sdk.DebuggeePausingReasonDebuggeeEventTriggered:
		if pausePacket.EventTag != 0 {
			if pausePacket.EventCallingStage == sdk.VmmCallbackCallingStagePostEventEmulation {
				fmt.Printf("event 0x%x triggered (post)\n", pausePacket.EventTag)
			} else {
				fmt.Printf("event 0x%x triggered (pre)\n", pausePacket.EventTag)
			}
		}

	case sdk.DebuggeePausingReasonDebuggeeStepped:
		break

	case sdk.DebuggeePausingReasonDebuggeeProcessSwitched:
		fmt.Println("switched to the specified process")

	case sdk.DebuggeePausingReasonDebuggeeThreadSwitched:
		fmt.Println("switched to the specified thread")

	case sdk.DebuggeePausingReasonDebuggeeStartingModuleLoaded:
		fmt.Println("the target module is loaded and a breakpoint is set to the entrypoint")
		fmt.Println("press 'g' to reach to the entrypoint of the main module...")

	case sdk.DebuggeePausingReasonDebuggeeTrackingStepped:
		break

	case sdk.DebuggeePausingReasonPause:
		break

	case sdk.DebuggeePausingReasonDebuggeeCoreSwitched:
		klm.signalEvent(int(SyncObjectKernelDebuggerCoreSwitchingResult))
		return

	case sdk.DebuggeePausingReasonDebuggeeCommandExecutionFinished:
		fmt.Println()
		klm.signalEvent(int(SyncObjectKernelDebuggerDebuggeeFinishedCommandExecution))
		return

	case sdk.DebuggeePausingReasonRequestFromDebugger:
		klm.signalEvent(int(SyncObjectKernelDebuggerPausedDebuggeeDetails))
		return

	default:
		fmt.Println("err, unknown pausing reason is received")
		return
	}

	klm.signalEvent(int(SyncObjectKernelDebuggerIsDebuggerRunning))
}

func (klm *KernelListeningManager) handleChangeCoreResultPacket(payload []byte) {
	if len(payload) < 8 {
		return
	}

	changeCorePacket := sdk.DEBUGGEE_CHANGE_CORE_PACKET{}
	changeCorePacket.NewCore = binary.LittleEndian.Uint32(payload[0:4])
	changeCorePacket.Result = binary.LittleEndian.Uint32(payload[4:8])

	if changeCorePacket.Result == 0 {
		fmt.Printf("current operating core changed to 0x%x\n", changeCorePacket.NewCore)
	} else {
		klm.showErrorMessage(changeCorePacket.Result)
		klm.signalEvent(int(SyncObjectKernelDebuggerCoreSwitchingResult))
	}
}

func (klm *KernelListeningManager) handleChangeProcessResultPacket(payload []byte) {
	if len(payload) < 40 {
		return
	}

	changeProcessPacket := sdk.DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET{}
	changeProcessPacket.Result = binary.LittleEndian.Uint32(payload[0:4])
	changeProcessPacket.ActionType = sdk.DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_TYPE(binary.LittleEndian.Uint32(payload[4:8]))
	changeProcessPacket.ProcessId = binary.LittleEndian.Uint32(payload[8:12])
	changeProcessPacket.Process = binary.LittleEndian.Uint64(payload[16:24])
	copy(changeProcessPacket.ProcessName[:], payload[24:40])

	if changeProcessPacket.Result == 0 {
		if changeProcessPacket.ActionType == 1 {
			fmt.Printf("process id: %x\nprocess (_EPROCESS): %s\nprocess name (16-Byte): %s\n",
				changeProcessPacket.ProcessId,
				SeparateTo64BitValue(changeProcessPacket.Process),
				string(changeProcessPacket.ProcessName[:]))
		} else if changeProcessPacket.ActionType == 2 {
			fmt.Println("press 'g' to continue the debuggee, if the pid or the process object address is valid then the debuggee will be automatically paused when it attached to the target process")
		}
	} else {
		klm.showErrorMessage(changeProcessPacket.Result)
	}

	klm.signalEvent(int(SyncObjectKernelDebuggerProcessSwitchingResult))
}

func (klm *KernelListeningManager) handleSearchResultPacket(payload []byte) {
	if len(payload) < 16 {
		return
	}

	searchResultsPacket := sdk.DEBUGGEE_RESULT_OF_SEARCH_PACKET{}
	searchResultsPacket.Result = binary.LittleEndian.Uint32(payload[0:4])
	searchResultsPacket.CountOfResults = binary.LittleEndian.Uint32(payload[8:12])

	if searchResultsPacket.Result == 0 {
		if searchResultsPacket.CountOfResults == 0 {
			fmt.Println("not found")
		}
	} else {
		klm.showErrorMessage(searchResultsPacket.Result)
	}

	klm.signalEvent(int(SyncObjectKernelDebuggerSearchQueryResult))
}

func (klm *KernelListeningManager) handleChangeThreadResultPacket(payload []byte) {
	if len(payload) < 56 {
		return
	}

	changeThreadPacket := sdk.DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET{}
	changeThreadPacket.Result = binary.LittleEndian.Uint32(payload[0:4])
	changeThreadPacket.ActionType = sdk.DEBUGGEE_DETAILS_AND_SWITCH_THREAD_TYPE(binary.LittleEndian.Uint32(payload[4:8]))
	changeThreadPacket.ThreadId = binary.LittleEndian.Uint32(payload[8:12])
	changeThreadPacket.ProcessId = binary.LittleEndian.Uint32(payload[12:16])
	changeThreadPacket.Thread = binary.LittleEndian.Uint64(payload[16:24])
	changeThreadPacket.Process = binary.LittleEndian.Uint64(payload[24:32])
	copy(changeThreadPacket.ProcessName[:], payload[40:56])

	if changeThreadPacket.Result == 0 {
		if changeThreadPacket.ActionType == 1 {
			fmt.Printf("thread id: %x (pid: %x)\nthread (_ETHREAD): %s\nprocess (_EPROCESS): %s\nprocess name (16-Byte): %s\n",
				changeThreadPacket.ThreadId,
				changeThreadPacket.ProcessId,
				SeparateTo64BitValue(changeThreadPacket.Thread),
				SeparateTo64BitValue(changeThreadPacket.Process),
				string(changeThreadPacket.ProcessName[:]))
		} else if changeThreadPacket.ActionType == 2 {
			fmt.Println("press 'g' to continue the debuggee, if the tid or the thread object address is valid then the debuggee will be automatically paused when it attached to the target thread")
		}
	} else {
		klm.showErrorMessage(changeThreadPacket.Result)
	}

	klm.signalEvent(int(SyncObjectKernelDebuggerThreadSwitchingResult))
}

func (klm *KernelListeningManager) handleFlushResultPacket(payload []byte) {
	if len(payload) < 12 {
		return
	}

	flushPacket := sdk.DEBUGGER_FLUSH_LOGGING_BUFFERS{}
	flushPacket.KernelStatus = binary.LittleEndian.Uint32(payload[0:4])
	flushPacket.CountOfMessagesThatSetAsReadFromVmxNonRoot = binary.LittleEndian.Uint32(payload[4:8])
	flushPacket.CountOfMessagesThatSetAsReadFromVmxRoot = binary.LittleEndian.Uint32(payload[8:12])

	if flushPacket.KernelStatus == 0 {
		fmt.Printf("flushing buffers was successful, total %d messages were cleared.\n",
			flushPacket.CountOfMessagesThatSetAsReadFromVmxNonRoot+flushPacket.CountOfMessagesThatSetAsReadFromVmxRoot)
	} else {
		klm.showErrorMessage(flushPacket.KernelStatus)
	}

	klm.signalEvent(int(SyncObjectKernelDebuggerFlushResult))
}

func (klm *KernelListeningManager) handleCallstackResultPacket(payload []byte) {
	if len(payload) < 16 {
		return
	}

	callstackPacket := sdk.DEBUGGER_CALLSTACK_REQUEST{}
	callstackPacket.KernelStatus = binary.LittleEndian.Uint32(payload[0:4])
	callstackPacket.FrameCount = binary.LittleEndian.Uint32(payload[4:8])
	callstackPacket.DisplayMethod = sdk.DEBUGGER_CALLSTACK_DISPLAY_METHOD(binary.LittleEndian.Uint32(payload[8:12]))
	callstackPacket.Is32Bit = payload[12] != 0

	if callstackPacket.KernelStatus != 0 {
		klm.showErrorMessage(callstackPacket.KernelStatus)
	}

	klm.signalEvent(int(SyncObjectKernelDebuggerCallstackResult))
}

func (klm *KernelListeningManager) handleTestQueryResultPacket(payload []byte) {
	if len(payload) < 8 {
		return
	}

	testQueryPacket := sdk.DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER{}
	testQueryPacket.KernelStatus = binary.LittleEndian.Uint32(payload[0:4])
	testQueryPacket.RequestType = sdk.DEBUGGER_TEST_QUERY_STATE(binary.LittleEndian.Uint32(payload[4:8]))

	if testQueryPacket.KernelStatus == 0 {
		switch testQueryPacket.RequestType {
		case 1:
			fmt.Println("breakpoint interception (#BP) is deactivated\nfrom now, the breakpoints will be re-injected into the guest debuggee")
		case 2:
			fmt.Println("breakpoint interception (#BP) is activated")
		case 3:
			fmt.Println("debug break interception (#DB) is deactivated\nfrom now, the debug breaks will be re-injected into the guest debuggee")
		case 4:
			fmt.Println("debug break interception (#DB) is activated")
		}
	} else {
		klm.showErrorMessage(testQueryPacket.KernelStatus)
	}

	klm.signalEvent(int(SyncObjectKernelDebuggerTestQuery))
}

func (klm *KernelListeningManager) handleScriptResultPacket(payload []byte) {
	if len(payload) < 8 {
		return
	}

	scriptPacket := sdk.DEBUGGEE_SCRIPT_PACKET{}
	scriptPacket.Result = binary.LittleEndian.Uint32(payload[0:4])
	scriptPacket.IsFormat = payload[4] != 0

	if scriptPacket.Result != 0 {
		klm.showErrorMessage(scriptPacket.Result)
	}

	if scriptPacket.IsFormat {
		klm.signalEvent(int(SyncObjectKernelDebuggerScriptFormatsResult))
	}

	klm.signalEvent(int(SyncObjectKernelDebuggerScriptRunningResult))
}

func (klm *KernelListeningManager) handleFormatsResultPacket(payload []byte) {
	if len(payload) < 12 {
		return
	}

	formatsPacket := sdk.DEBUGGEE_FORMATS_PACKET{}
	formatsPacket.Result = binary.LittleEndian.Uint32(payload[0:4])
	formatsPacket.Value = binary.LittleEndian.Uint64(payload[4:12])

	klm.mu.Lock()
	klm.errorStateOfResultOfEvaluatedExpression = formatsPacket.Result
	klm.resultOfEvaluatedExpression = formatsPacket.Value
	klm.mu.Unlock()
}

func (klm *KernelListeningManager) handleRegisterEventResultPacket(payload []byte) {
	if len(payload) < 16 {
		return
	}

	eventAndActionPacket := sdk.DEBUGGER_EVENT_AND_ACTION_RESULT{}
	eventAndActionPacket.IsSuccessful = payload[0] != 0
	eventAndActionPacket.Error = binary.LittleEndian.Uint32(payload[4:8])

	klm.mu.Lock()
	klm.debuggeeResultOfRegisteringEvent = eventAndActionPacket
	klm.mu.Unlock()

	klm.signalEvent(int(SyncObjectKernelDebuggerRegisterEvent))
}

func (klm *KernelListeningManager) handleAddActionToEventResultPacket(payload []byte) {
	if len(payload) < 16 {
		return
	}

	eventAndActionPacket := sdk.DEBUGGER_EVENT_AND_ACTION_RESULT{}
	eventAndActionPacket.IsSuccessful = payload[0] != 0
	eventAndActionPacket.Error = binary.LittleEndian.Uint32(payload[4:8])

	klm.mu.Lock()
	klm.debuggeeResultOfAddingActionsToEvent = eventAndActionPacket
	klm.mu.Unlock()

	klm.signalEvent(int(SyncObjectKernelDebuggerAddActionToEvent))
}

func (klm *KernelListeningManager) handleQueryAndModifyEventResultPacket(payload []byte) {
	if len(payload) < 16 {
		return
	}

	eventModifyAndQueryPacket := sdk.DEBUGGER_MODIFY_EVENTS{}
	eventModifyAndQueryPacket.Tag = binary.LittleEndian.Uint64(payload[0:8])
	eventModifyAndQueryPacket.TypeOfAction = sdk.DEBUGGER_MODIFY_EVENTS_TYPE(binary.LittleEndian.Uint32(payload[8:12]))
	kernelStatus := binary.LittleEndian.Uint32(payload[12:16])

	if kernelStatus != 0 {
		klm.showErrorMessage(kernelStatus)
	} else if eventModifyAndQueryPacket.TypeOfAction == 3 {
		klm.mu.Lock()
		klm.sharedEventStatus = eventModifyAndQueryPacket.IsEnabled
		klm.mu.Unlock()
	}

	klm.signalEvent(int(SyncObjectKernelDebuggerModifyAndQueryEvent))
}

func (klm *KernelListeningManager) handleSymbolReloadFinishedPacket(payload []byte) {
	if len(payload) < 4 {
		return
	}

	kernelStatus := binary.LittleEndian.Uint32(payload[0:4])

	if kernelStatus != 0 {
		klm.showErrorMessage(kernelStatus)
	}

	klm.signalEvent(int(SyncObjectKernelDebuggerSymbolReload))
}

func (klm *KernelListeningManager) handleReadRegistersResultPacket(payload []byte) {
	klm.signalEvent(int(SyncObjectKernelDebuggerReadRegisters))
}

func (klm *KernelListeningManager) handleWriteRegisterResultPacket(payload []byte) {
	klm.signalEvent(int(SyncObjectKernelDebuggerWriteRegister))
}

func (klm *KernelListeningManager) handleApicResultPacket(payload []byte) {
	klm.signalEvent(int(SyncObjectKernelDebuggerApicActions))
}

func (klm *KernelListeningManager) handleIdtEntriesResultPacket(payload []byte) {
	klm.signalEvent(int(SyncObjectKernelDebuggerIdtEntries))
}

func (klm *KernelListeningManager) handleReadMemoryResultPacket(payload []byte) {
	klm.signalEvent(int(SyncObjectKernelDebuggerReadMemory))
}

func (klm *KernelListeningManager) handleEditMemoryResultPacket(payload []byte) {
	klm.signalEvent(int(SyncObjectKernelDebuggerEditMemory))
}

func (klm *KernelListeningManager) handleBpResultPacket(payload []byte) {
	if len(payload) < 4 {
		return
	}

	result := binary.LittleEndian.Uint32(payload[0:4])
	if result != 0 {
		klm.showErrorMessage(result)
	}

	klm.signalEvent(int(SyncObjectKernelDebuggerBp))
}

func (klm *KernelListeningManager) handleShortCircuitingResultPacket(payload []byte) {
	if len(payload) < 8 {
		return
	}

	kernelStatus := binary.LittleEndian.Uint32(payload[0:4])
	isShortCircuiting := payload[4] != 0

	if kernelStatus == 0 {
		if isShortCircuiting {
			fmt.Println("the event's short-circuiting state changed to 'on'")
		} else {
			fmt.Println("the event's short-circuiting state changed to 'off'")
		}
	} else {
		klm.showErrorMessage(kernelStatus)
	}

	klm.signalEvent(int(SyncObjectKernelDebuggerShortCircuitingEventState))
}

func (klm *KernelListeningManager) handlePteResultPacket(payload []byte) {
	if len(payload) < 12 {
		return
	}

	kernelStatus := binary.LittleEndian.Uint32(payload[0:4])
	if kernelStatus != 0 {
		klm.showErrorMessage(kernelStatus)
	}

	klm.signalEvent(int(SyncObjectKernelDebuggerPteResult))
}

func (klm *KernelListeningManager) handleSmiOperationResultPacket(payload []byte) {
	klm.signalEvent(int(SyncObjectKernelDebuggerSmiOperationResult))
}

func (klm *KernelListeningManager) handleHyperTraceLbrOperationResultPacket(payload []byte) {
	klm.signalEvent(int(SyncObjectKernelDebuggerHyperTraceLbrOperationResult))
}

func (klm *KernelListeningManager) handleHyperTracePtOperationResultPacket(payload []byte) {
	klm.signalEvent(int(SyncObjectKernelDebuggerHyperTracePtOperationResult))
}

func (klm *KernelListeningManager) handlePageInResultPacket(payload []byte) {
	if len(payload) < 4 {
		return
	}

	kernelStatus := binary.LittleEndian.Uint32(payload[0:4])
	if kernelStatus == 0 {
		fmt.Println("the page-fault is delivered to the target thread\npress 'g' to continue debuggee (the current thread will execute ONLY one instruction and will be halted again)...")
	} else {
		klm.showErrorMessage(kernelStatus)
	}

	klm.signalEvent(int(SyncObjectKernelDebuggerPageInState))
}

func (klm *KernelListeningManager) handleVa2paPa2vaResultPacket(payload []byte) {
	if len(payload) < 24 {
		return
	}

	kernelStatus := binary.LittleEndian.Uint32(payload[0:4])
	isVirtual2Physical := payload[4] != 0
	physicalAddress := binary.LittleEndian.Uint64(payload[8:16])
	virtualAddress := binary.LittleEndian.Uint64(payload[16:24])

	if kernelStatus == 0 {
		if isVirtual2Physical {
			fmt.Printf("%x\n", physicalAddress)
		} else {
			fmt.Printf("%x\n", virtualAddress)
		}
	} else {
		klm.showErrorMessage(kernelStatus)
	}

	klm.signalEvent(int(SyncObjectKernelDebuggerVa2paAndPa2vaResult))
}

func (klm *KernelListeningManager) handleListOrModifyBreakpointsResultPacket(payload []byte) {
	if len(payload) < 4 {
		return
	}

	result := binary.LittleEndian.Uint32(payload[0:4])
	if result != 0 {
		klm.showErrorMessage(result)
	}

	klm.signalEvent(int(SyncObjectKernelDebuggerListOrModifyBreakpoints))
}

func (klm *KernelListeningManager) handleUpdateSymbolInfoPacket(payload []byte) {
}

func (klm *KernelListeningManager) handlePcitreeResultPacket(payload []byte) {
	if len(payload) < 4 {
		return
	}

	kernelStatus := binary.LittleEndian.Uint32(payload[0:4])
	if kernelStatus != 0 {
		klm.showErrorMessage(kernelStatus)
	}

	klm.signalEvent(int(SyncObjectKernelDebuggerPcitreeResult))
}

func (klm *KernelListeningManager) handlePcidevinfoResultPacket(payload []byte) {
	if len(payload) < 4 {
		return
	}

	kernelStatus := binary.LittleEndian.Uint32(payload[0:4])
	if kernelStatus != 0 {
		klm.showErrorMessage(kernelStatus)
	}

	klm.signalEvent(int(SyncObjectKernelDebuggerPcidevinfoResult))
}

func (klm *KernelListeningManager) ListeningSerialPortInDebuggee() bool {
	for {
		bufferToReceive := make([]byte, int(sdk.MaxSerialPacketSize))
		var lengthReceived uint32

		received, err := klm.kdReceivePacketFromDebuggee(bufferToReceive, &lengthReceived)
		if err != nil {
			fmt.Printf("err, error receiving packet: %v\n", err)
			continue
		}

		if !received {
			if lengthReceived == 0 && bufferToReceive[0] == 0 {
				fmt.Println("the remote connection is closed")
				klm.kdCloseConnection()
				return false
			}
			continue
		}

		if lengthReceived == 1 && bufferToReceive[0] == 0 {
			continue
		}

		packet := bufferToReceive[:lengthReceived]

		if len(packet) < 24 {
			continue
		}

		indicator := binary.LittleEndian.Uint64(packet[8:16])
		if indicator != sdk.IndicatorOfHyperdbgPacket {
			continue
		}

		expectedChecksum := ComputeDataChecksum(packet[1:])
		actualChecksum := packet[0]
		if expectedChecksum != actualChecksum {
			fmt.Println("err checksum is invalid")
			continue
		}

		packetType := binary.LittleEndian.Uint32(packet[16:20])
		if packetType != 2 {
			fmt.Println("err, unknown packet received from the debugger")
			continue
		}

		requestedAction := binary.LittleEndian.Uint32(packet[20:24])

		switch requestedAction {
		case uint32(sdk.DebuggerRemotePacketRequestedActionOnUserModePause):
			fmt.Println("debugger requested pause")

		case uint32(sdk.DebuggerRemotePacketRequestedActionOnUserModeDoNotReadAnyPacket):
			return true

		default:
			fmt.Println("err, unknown packet action received from the debugger")
		}
	}
}

func (klm *KernelListeningManager) kdReceivePacketFromDebuggee(buffer []byte, lengthReceived *uint32) (bool, error) {
	*lengthReceived = 0
	return true, nil
}

func (klm *KernelListeningManager) kdCloseConnection() {
	klm.mu.Lock()
	defer klm.mu.Unlock()

	klm.isSerialConnectedToRemoteDebuggee = false
	klm.isDebuggeeRunning = false

	if klm.kd != nil {
		klm.kd.CloseConnection()
	}
}

func (klm *KernelListeningManager) kdSendResponseOfPingPacket() {
	if klm.kd != nil {
		klm.kd.SendResponseOfThePingPacket([]byte(BuildSignature))
	}
}

func (klm *KernelListeningManager) signalEvent(syncObjectType int) {
	if syncObjectType >= 0 && syncObjectType < MaximumSyncObjects {
		klm.syncObjects[syncObjectType].mu.Lock()
		klm.syncObjects[syncObjectType].IsSignaled = true
		klm.syncObjects[syncObjectType].mu.Unlock()
	}
}

func (klm *KernelListeningManager) showErrorMessage(errorCode uint32) {
	fmt.Printf("err, operation failed with error code: 0x%x\n", errorCode)
}

func (klm *KernelListeningManager) IsDebuggeeRunning() bool {
	klm.mu.RLock()
	defer klm.mu.RUnlock()
	return klm.isDebuggeeRunning
}

func (klm *KernelListeningManager) SetDebuggeeRunning(running bool) {
	klm.mu.Lock()
	defer klm.mu.Unlock()
	klm.isDebuggeeRunning = running
}

func (klm *KernelListeningManager) GetCurrentCore() uint32 {
	klm.mu.RLock()
	defer klm.mu.RUnlock()
	return klm.currentRemoteCore
}

func (klm *KernelListeningManager) GetKernelBaseAddress() uint64 {
	klm.mu.RLock()
	defer klm.mu.RUnlock()
	return klm.kernelBaseAddress
}

func (klm *KernelListeningManager) IsSerialConnected() bool {
	klm.mu.RLock()
	defer klm.mu.RUnlock()
	return klm.isSerialConnectedToRemoteDebuggee
}

func (klm *KernelListeningManager) GetSharedEventStatus() bool {
	klm.mu.RLock()
	defer klm.mu.RUnlock()
	return klm.sharedEventStatus
}

func (klm *KernelListeningManager) GetResultOfEvaluatedExpression() uint64 {
	klm.mu.RLock()
	defer klm.mu.RUnlock()
	return klm.resultOfEvaluatedExpression
}

func (klm *KernelListeningManager) GetErrorStateOfResultOfEvaluatedExpression() uint32 {
	klm.mu.RLock()
	defer klm.mu.RUnlock()
	return klm.errorStateOfResultOfEvaluatedExpression
}

func extractMessageFromPayload(payload []byte) string {
	if len(payload) > 8 {
		return string(payload[8:])
	}
	return ""
}

func parsePausePacket(payload []byte) *sdk.DEBUGGEE_KD_PAUSED_PACKET {
	pausePacket := &sdk.DEBUGGEE_KD_PAUSED_PACKET{}

	if len(payload) < int(unsafe.Sizeof(sdk.DEBUGGEE_KD_PAUSED_PACKET{})) {
		return pausePacket
	}

	offset := 0
	pausePacket.CurrentCore = binary.LittleEndian.Uint32(payload[offset : offset+4])
	offset += 4

	pausePacket.PausingReason = sdk.DEBUGGEE_PAUSING_REASON(binary.LittleEndian.Uint32(payload[offset : offset+4]))
	offset += 4

	pausePacket.EventTag = binary.LittleEndian.Uint64(payload[offset : offset+8])
	offset += 8

	pausePacket.EventCallingStage = sdk.VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE(binary.LittleEndian.Uint32(payload[offset : offset+4]))
	offset += 4

	copy(pausePacket.InstructionBytesOnRip[:], payload[offset:offset+int(sdk.MaximumInstrSize)])
	offset += int(sdk.MaximumInstrSize)

	pausePacket.ReadInstructionLen = binary.LittleEndian.Uint16(payload[offset : offset+2])
	offset += 2

	pausePacket.IsProcessorOn32BitMode = payload[offset] != 0
	offset += 1

	pausePacket.IgnoreDisassembling = payload[offset] != 0
	offset += 1

	pausePacket.Rip = binary.LittleEndian.Uint64(payload[offset : offset+8])
	offset += 8

	pausePacket.Rflags = binary.LittleEndian.Uint64(payload[offset : offset+8])

	return pausePacket
}

func SeparateTo64BitValue(value uint64) string {
	high := value >> 32
	low := value & 0xFFFFFFFF

	if high == 0 {
		return fmt.Sprintf("%x", low)
	}
	return fmt.Sprintf("%x`%08x", high, low)
}
