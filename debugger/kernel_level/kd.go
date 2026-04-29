package kernel_level

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sync"
	"unsafe"

	"github.com/ddkwork/HyperDbg/debugger/core"
	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/sdk"
)

const (
	BuildSignature = "HyperDbg"
)

type SyncObjectType int

const (
	SyncObjectKernelDebuggerStartedPacketReceived SyncObjectType = iota
	SyncObjectKernelDebuggerPausedDebuggeeDetails
	SyncObjectKernelDebuggerIsDebuggerRunning
	SyncObjectKernelDebuggerCoreSwitchingResult
	SyncObjectKernelDebuggerShortCircuitingEventState
	SyncObjectKernelDebuggerModifyAndQueryEvent
	SyncObjectKernelDebuggerFlushResult
	SyncObjectKernelDebuggerCallstackResult
	SyncObjectKernelDebuggerTestQuery
	SyncObjectKernelDebuggerSymbolReload
	SyncObjectKernelDebuggerReadRegisters
	SyncObjectKernelDebuggerWriteRegister
	SyncObjectKernelDebuggerReadMemory
	SyncObjectKernelDebuggerEditMemory
	SyncObjectKernelDebuggerRegisterEvent
	SyncObjectKernelDebuggerAddActionToEvent
	SyncObjectKernelDebuggerProcessSwitchingResult
	SyncObjectKernelDebuggerThreadSwitchingResult
	SyncObjectKernelDebuggerPteResult
	SyncObjectKernelDebuggerPageInState
	SyncObjectKernelDebuggerVa2paAndPa2vaResult
	SyncObjectKernelDebuggerApicActions
	SyncObjectKernelDebuggerSmiOperationResult
	SyncObjectKernelDebuggerHyperTraceLbrOperationResult
	SyncObjectKernelDebuggerHyperTracePtOperationResult
	SyncObjectKernelDebuggerIdtEntries
	SyncObjectKernelDebuggerBp
	SyncObjectKernelDebuggerListOrModifyBreakpoints
	SyncObjectKernelDebuggerScriptRunningResult
	SyncObjectKernelDebuggerScriptFormatsResult
	SyncObjectKernelDebuggerDebuggeeFinishedCommandExecution
	SyncObjectKernelDebuggerSearchQueryResult
	SyncObjectKernelDebuggerPcitreeResult
	SyncObjectKernelDebuggerPcidevinfoResult
)

const KernelDebuggerMaxSyncObjects = 42

type KernelDebugger struct {
	mu                                sync.Mutex
	currentRemoteCore                 uint32
	isSerialConnectedToRemoteDebuggee bool
	isSerialConnectedToRemoteDebugger bool
	isDebuggeeRunning                 bool
	ignoreNewLoggingMessages          bool
	ignorePauseRequests               bool
	isDebuggeeInHandshakingPhase      bool
	serialConnectionAlreadyClosed     bool
	shouldPreviousCommandBeContinued  bool
	isRunningInstruction32Bit         bool
	kernelBaseAddress                 uint64
	currentRunningInstruction         [sdk.MaximumInstrSize]byte
	sharedEventStatus                 bool
	sendCommandToDebuggee             func(uint32, sdk.DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION, []byte) bool
	sendCommandAndBufferToDebuggee    func(uint32, sdk.DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION, []byte) bool
	onInterpretPausedDebuggee         func()
	core                              *core.Core
	syncObjects                       [KernelDebuggerMaxSyncObjects]KernelSyncObjectState
}

func NewKernelDebugger(
	sendCommand func(uint32, sdk.DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION, []byte) bool,
	onInterpretPaused func(),
) *KernelDebugger {
	kd := &KernelDebugger{
		currentRemoteCore:         uint32(sdk.DebuggerRemotePacketRequestedActionDebuggeePausedAndCurrentInstruction),
		sendCommandToDebuggee:     sendCommand,
		onInterpretPausedDebuggee: onInterpretPaused,
		core: core.NewCore(func(req core.SteppingRequest) bool {
			return sendCommand(1, sdk.DebuggerRemotePacketRequestedActionOnVmxRootModeStep, nil)
		}),
	}
	return kd
}

func CheckForTheEndOfBuffer(currentLoopIndex *uint32, buffer []byte) bool {
	actualBufferLength := *currentLoopIndex

	if actualBufferLength <= 3 {
		return false
	}

	if buffer[actualBufferLength] == byte(sdk.SerialEndOfBufferChar4) &&
		buffer[actualBufferLength-1] == byte(sdk.SerialEndOfBufferChar3) &&
		buffer[actualBufferLength-2] == byte(sdk.SerialEndOfBufferChar2) &&
		buffer[actualBufferLength-3] == byte(sdk.SerialEndOfBufferChar1) {

		buffer[actualBufferLength-3] = 0
		buffer[actualBufferLength-2] = 0
		buffer[actualBufferLength-1] = 0
		buffer[actualBufferLength] = 0

		*currentLoopIndex = actualBufferLength - 3
		return true
	}
	return false
}

func ComputeDataChecksum(buffer []byte) byte {
	var checksum byte = 0
	for _, b := range buffer {
		checksum += b
	}
	return checksum
}

func (kd *KernelDebugger) InterpretPausedDebuggee() {
	kd.waitForKernelResponse(SyncObjectKernelDebuggerPausedDebuggeeDetails)
}

func (kd *KernelDebugger) SendContinuePacketToDebuggee() bool {
	kd.currentRemoteCore = uint32(sdk.DebuggerRemotePacketRequestedActionDebuggeePausedAndCurrentInstruction)

	if !kd.commandPacketToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootModeContinue) {
		return false
	}
	return true
}

func (kd *KernelDebugger) SendSwitchCorePacketToDebuggee(newCore uint32) bool {
	if newCore == kd.currentRemoteCore {
		mylog.Info("the current operating core is", "core", fmt.Sprintf("%x", newCore), "(not changed)")
		return false
	}

	buf := bytes.NewBuffer(nil)
	changeCore := sdk.DEBUGGEE_CHANGE_CORE_PACKET{NewCore: newCore}
	binary.Write(buf, binary.LittleEndian, &changeCore)

	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootModeChangeCore,
		buf.Bytes()) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerCoreSwitchingResult)
	kd.currentRemoteCore = newCore
	return true
}

func (kd *KernelDebugger) SendShortCircuitingEventToDebuggee(isEnabled bool) bool {
	packet := sdk.DEBUGGER_SHORT_CIRCUITING_EVENT{IsShortCircuiting: isEnabled}

	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.LittleEndian, &packet)

	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootSetShortCircuitingState,
		buf.Bytes()) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerShortCircuitingEventState)
	return true
}

func (kd *KernelDebugger) SendEventQueryAndModifyPacketToDebuggee(tag uint64, typeOfAction uint32, isEnabled *bool) bool {
	kd.sharedEventStatus = false

	packet := sdk.DEBUGGER_MODIFY_EVENTS{
		Tag:          tag,
		TypeOfAction: sdk.DEBUGGER_MODIFY_EVENTS_TYPE(typeOfAction),
	}

	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.LittleEndian, &packet)

	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootQueryAndModifyEvent,
		buf.Bytes()) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerModifyAndQueryEvent)

	if typeOfAction == 3 {
		if isEnabled != nil {
			*isEnabled = kd.sharedEventStatus
		}
	}

	return true
}

func (kd *KernelDebugger) SendFlushPacketToDebuggee() bool {
	packet := sdk.DEBUGGER_FLUSH_LOGGING_BUFFERS{}

	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.LittleEndian, &packet)

	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootModeFlushBuffers,
		buf.Bytes()) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerFlushResult)
	return true
}

func (kd *KernelDebugger) SendCallStackPacketToDebuggee(baseAddress uint64, size uint32, displayMethod uint32, is32Bit bool) bool {
	if size == 0 {
		return false
	}

	var frameCount uint32
	if is32Bit {
		frameCount = size / 4
	} else {
		frameCount = size / 8
	}

	callstackRequestSize := uint64(unsafe.Sizeof(sdk.DEBUGGER_CALLSTACK_REQUEST{})) + uint64(frameCount)*uint64(unsafe.Sizeof(uint64(0)))*2

	packet := sdk.DEBUGGER_CALLSTACK_REQUEST{
		BaseAddress:   baseAddress,
		Is32Bit:       is32Bit,
		Size:          size,
		BufferSize:    callstackRequestSize,
		FrameCount:    frameCount,
		DisplayMethod: sdk.DEBUGGER_CALLSTACK_DISPLAY_METHOD(displayMethod),
	}

	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.LittleEndian, &packet)

	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootModeCallstack,
		buf.Bytes()) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerCallstackResult)
	return true
}

func (kd *KernelDebugger) SendTestQueryPacketToDebuggee(requestType uint32) bool {
	packet := sdk.DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER{RequestType: sdk.DEBUGGER_TEST_QUERY_STATE(requestType)}

	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.LittleEndian, &packet)

	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootModeTestQuery,
		buf.Bytes()) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerTestQuery)
	return true
}

func (kd *KernelDebugger) SendTestQueryPacketWithContextToDebuggee(requestType uint32, context uint64) bool {
	packet := sdk.DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER{
		RequestType: sdk.DEBUGGER_TEST_QUERY_STATE(requestType),
		Context:     context,
	}

	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.LittleEndian, &packet)

	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootModeTestQuery,
		buf.Bytes()) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerTestQuery)
	return true
}

func (kd *KernelDebugger) SendSymbolReloadPacketToDebuggee(processId uint32) bool {
	packet := struct {
		ProcessId uint32
	}{ProcessId: processId}

	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.LittleEndian, &packet)

	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootSymbolReload,
		buf.Bytes()) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerSymbolReload)
	return true
}

func (kd *KernelDebugger) SendReadRegisterPacketToDebuggee(regDes []byte) bool {
	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootReadRegisters,
		regDes) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerReadRegisters)
	return true
}

func (kd *KernelDebugger) SendWriteRegisterPacketToDebuggee(regDes []byte) bool {
	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootWriteRegister,
		regDes) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerWriteRegister)
	return true
}

func (kd *KernelDebugger) SendReadMemoryPacketToDebuggee(readMem []byte) bool {
	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootReadMemory,
		readMem) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerReadMemory)
	return true
}

func (kd *KernelDebugger) SendEditMemoryPacketToDebuggee(editMem []byte) bool {
	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootEditMemory,
		editMem) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerEditMemory)
	return true
}

func (kd *KernelDebugger) SendRegisterEventPacketToDebuggee(event []byte, eventBufferLength uint32) []byte {
	headerSize := uint32(unsafe.Sizeof(sdk.DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET{}))
	len := eventBufferLength + headerSize

	headerBuf := make([]byte, len)
	binary.LittleEndian.PutUint32(headerBuf[0:4], eventBufferLength)
	copy(headerBuf[headerSize:], event)

	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootRegisterEvent,
		headerBuf) {
		return nil
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerRegisterEvent)
	return nil
}

func (kd *KernelDebugger) SendAddActionToEventPacketToDebuggee(generalAction []byte, generalActionLength uint32) []byte {
	headerSize := uint32(unsafe.Sizeof(sdk.DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET{}))
	len := generalActionLength + headerSize

	headerBuf := make([]byte, len)
	binary.LittleEndian.PutUint32(headerBuf[0:4], generalActionLength)
	copy(headerBuf[headerSize:], generalAction)

	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootAddActionToEvent,
		headerBuf) {
		return nil
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerAddActionToEvent)
	return nil
}

func (kd *KernelDebugger) SendSwitchProcessPacketToDebuggee(actionType uint32, newPid uint32, newProcess uint64, setChangeByClockInterrupt bool) bool {
	packet := sdk.DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET{
		ActionType:        sdk.DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_TYPE(actionType),
		ProcessId:         newPid,
		Process:           newProcess,
		IsSwitchByClkIntr: setChangeByClockInterrupt,
	}

	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.LittleEndian, &packet)

	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootModeChangeProcess,
		buf.Bytes()) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerProcessSwitchingResult)
	return true
}

func (kd *KernelDebugger) SendSwitchThreadPacketToDebuggee(actionType uint32, newTid uint32, newThread uint64, checkByClockInterrupt bool) bool {
	packet := sdk.DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET{
		ActionType:            sdk.DEBUGGEE_DETAILS_AND_SWITCH_THREAD_TYPE(actionType),
		ThreadId:              newTid,
		Thread:                newThread,
		CheckByClockInterrupt: checkByClockInterrupt,
	}

	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.LittleEndian, &packet)

	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootModeChangeThread,
		buf.Bytes()) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerThreadSwitchingResult)
	return true
}

func (kd *KernelDebugger) SendPtePacketToDebuggee(ptePacket []byte) bool {
	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootSymbolQueryPte,
		ptePacket) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerPteResult)
	return true
}

func (kd *KernelDebugger) SendPageinPacketToDebuggee(pageinPacket []byte) bool {
	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootInjectPageFault,
		pageinPacket) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerPageInState)
	return true
}

func (kd *KernelDebugger) SendVa2paAndPa2vaPacketToDebuggee(va2paPa2vaPacket []byte) bool {
	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootQueryPa2vaAndVa2pa,
		va2paPa2vaPacket) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerVa2paAndPa2vaResult)
	return true
}

func (kd *KernelDebugger) SendApicActionPacketsToDebuggee(apicRequest []byte) bool {
	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootPerformActionsOnApic,
		apicRequest) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerApicActions)
	return true
}

func (kd *KernelDebugger) SendSmiPacketsToDebuggee(smiOperationRequest []byte) bool {
	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootPerformSmiOperation,
		smiOperationRequest) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerSmiOperationResult)
	return true
}

func (kd *KernelDebugger) SendHyperTraceLbrPacketsToDebuggee(hyperTraceLbrOperationRequest []byte) bool {
	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootPerformHypertraceLbrOperation,
		hyperTraceLbrOperationRequest) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerHyperTraceLbrOperationResult)
	return true
}

func (kd *KernelDebugger) SendHyperTracePtPacketsToDebuggee(hyperTracePtOperationRequest []byte) bool {
	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootPerformHypertracePtOperation,
		hyperTracePtOperationRequest) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerHyperTracePtOperationResult)
	return true
}

func (kd *KernelDebugger) SendQueryIdtPacketsToDebuggee(idtRequest []byte) bool {
	if !kd.commandPacketToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootReadIdtEntries) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerIdtEntries)
	return true
}

func (kd *KernelDebugger) SendBpPacketToDebuggee(bpPacket []byte) bool {
	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootBp,
		bpPacket) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerBp)
	return true
}

func (kd *KernelDebugger) SendListOrModifyPacketToDebuggee(listOrModifyPacket []byte) bool {
	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootListOrModifyBreakpoints,
		listOrModifyPacket) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerListOrModifyBreakpoints)
	return true
}

func (kd *KernelDebugger) SendScriptPacketToDebuggee(bufferAddress uint64, bufferLength uint32, pointer uint32, isFormat bool) bool {
	scriptPacket := sdk.DEBUGGEE_SCRIPT_PACKET{
		ScriptBufferSize:    bufferLength,
		ScriptBufferPointer: pointer,
		IsFormat:            isFormat,
	}

	headerBuf := bytes.NewBuffer(nil)
	binary.Write(headerBuf, binary.LittleEndian, &scriptPacket)

	scriptData := unsafe.Slice((*byte)(unsafe.Pointer(uintptr(bufferAddress))), bufferLength)
	headerBuf.Write(scriptData)

	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootRunScript,
		headerBuf.Bytes()) {
		return false
	}

	if isFormat {
		kd.waitForKernelResponse(SyncObjectKernelDebuggerScriptFormatsResult)
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerScriptRunningResult)
	return true
}

func (kd *KernelDebugger) SendUserInputPacketToDebuggee(sendBuf []byte, ignoreBreakingAgain bool) bool {
	userInputPkt := sdk.DEBUGGEE_USER_INPUT_PACKET{
		CommandLen:           uint32(len(sendBuf)),
		IgnoreFinishedSignal: ignoreBreakingAgain,
	}

	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.LittleEndian, &userInputPkt)
	buf.Write(sendBuf)

	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootUserInputBuffer,
		buf.Bytes()) {
		return false
	}

	if !ignoreBreakingAgain {
		kd.waitForKernelResponse(SyncObjectKernelDebuggerDebuggeeFinishedCommandExecution)
	}

	return true
}

func (kd *KernelDebugger) SendSearchRequestPacketToDebuggee(searchRequestBuffer []byte) bool {
	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootSearchQuery,
		searchRequestBuffer) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerSearchQueryResult)
	return true
}

func (kd *KernelDebugger) SendStepPacketToDebuggee(stepRequestType uint32) bool {
	stepPacket := sdk.DEBUGGEE_STEP_PACKET{
		StepType: sdk.DEBUGGER_REMOTE_STEPPING_REQUEST(stepRequestType),
	}

	if stepRequestType == 3 || stepRequestType == 6 {
		callInstructionSize := uint32(0)
		isCall := kd.checkWhetherCurrentInstructionIsCall(&callInstructionSize)
		if isCall {
			stepPacket.IsCurrentInstructionACall = true
			stepPacket.CallLength = callInstructionSize
		}
	}

	if stepRequestType == 3 || stepRequestType == 6 || stepRequestType == 2 {
		kd.mu.Lock()
		kd.isDebuggeeRunning = true
		kd.mu.Unlock()
	}

	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.LittleEndian, &stepPacket)

	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootModeStep,
		buf.Bytes()) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerIsDebuggerRunning)
	return true
}

func (kd *KernelDebugger) SendPausePacketToDebuggee() bool {
	if !kd.commandPacketToDebuggee(
		2,
		sdk.DebuggerRemotePacketRequestedActionOnUserModePause) {
		return false
	}

	kd.InterpretPausedDebuggee()
	return true
}

func (kd *KernelDebugger) BreakControlCheckAndPauseDebugger(signalRunningFlag bool) {
	kd.mu.Lock()
	running := kd.isDebuggeeRunning
	kd.mu.Unlock()

	if running {
		if !kd.SendPausePacketToDebuggee() {
			fmt.Println("err, unable to pause the debuggee")
		}

		if signalRunningFlag {
			kd.receivedKernelResponse(SyncObjectKernelDebuggerIsDebuggerRunning)
		}
	}
}

func (kd *KernelDebugger) SetStatusAndWaitForPause() {
	kd.mu.Lock()
	kd.isDebuggeeRunning = true
	kd.mu.Unlock()

	kd.theRemoteSystemIsRunning()
}

func (kd *KernelDebugger) BreakControlCheckAndContinueDebugger() {
	kd.mu.Lock()
	running := kd.isDebuggeeRunning
	kd.mu.Unlock()

	if !running {
		if kd.SendContinuePacketToDebuggee() {
			kd.SetStatusAndWaitForPause()
		} else {
			fmt.Println("err, unable to continue the debuggee")
		}
	}
}

func (kd *KernelDebugger) SendResponseOfThePingPacket(buildSignature []byte) bool {
	if !kd.commandPacketAndBufferToDebuggee(
		2,
		sdk.DebuggerRemotePacketRequestedActionOnUserModeDebuggerVersion,
		buildSignature) {
		fmt.Println("err, unable to send response to the ping packet")
		return false
	}
	return true
}

func (kd *KernelDebugger) theRemoteSystemIsRunning() {
	fmt.Println("debuggee is running...")
	kd.waitForKernelResponse(SyncObjectKernelDebuggerIsDebuggerRunning)
}

func (kd *KernelDebugger) CloseConnection() bool {
	kd.mu.Lock()
	if kd.serialConnectionAlreadyClosed {
		kd.mu.Unlock()
		return true
	}
	kd.serialConnectionAlreadyClosed = true
	kd.mu.Unlock()

	if kd.isSerialConnectedToRemoteDebugger {
		fmt.Println("unloading debugger vmm module on debuggee...")
	} else if kd.isSerialConnectedToRemoteDebuggee {
		if kd.commandPacketToDebuggee(
			1,
			sdk.DebuggerRemotePacketRequestedActionOnVmxRootModeCloseAndUnloadDebuggee) {
			fmt.Println("unloading debugger vmm module on debuggee...")

			kd.commandPacketToDebuggee(
				2,
				sdk.DebuggerRemotePacketRequestedActionOnUserModeDoNotReadAnyPacket)
		}
	} else {
		fmt.Println("err, start packet not received but the target machine closed the connection")
		kd.receivedKernelResponse(SyncObjectKernelDebuggerStartedPacketReceived)
	}

	kd.uninitializeConnection()
	return true
}

func (kd *KernelDebugger) SendPcitreePacketToDebuggee(pcitreePacket []byte) bool {
	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootQueryPcitree,
		pcitreePacket) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerPcitreeResult)
	return true
}

func (kd *KernelDebugger) SendPcidevinfoPacketToDebuggee(pcidevinfoPacket []byte) bool {
	if !kd.commandPacketAndBufferToDebuggee(
		1,
		sdk.DebuggerRemotePacketRequestedActionOnVmxRootQueryPcidevinfo,
		pcidevinfoPacket) {
		return false
	}

	kd.waitForKernelResponse(SyncObjectKernelDebuggerPcidevinfoResult)
	return true
}

func (kd *KernelDebugger) commandPacketToDebuggee(packetType uint32, requestedAction sdk.DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION) bool {
	packet := sdk.DEBUGGER_REMOTE_PACKET{
		Indicator:                  sdk.IndicatorOfHyperdbgPacket,
		TypeOfThePacket:            sdk.DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnVmxRoot,
		RequestedActionOfThePacket: requestedAction,
	}

	checksumData := make([]byte, 0)
	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.LittleEndian, &packet)
	data := buf.Bytes()
	checksumData = append(checksumData, data[1:]...)
	packet.Checksum = ComputeDataChecksum(checksumData)

	buf.Reset()
	binary.Write(buf, binary.LittleEndian, &packet)

	if kd.sendCommandToDebuggee != nil {
		return kd.sendCommandToDebuggee(packetType, requestedAction, buf.Bytes())
	}
	return false
}

func (kd *KernelDebugger) commandPacketAndBufferToDebuggee(packetType uint32, requestedAction sdk.DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION, buffer []byte) bool {
	packet := sdk.DEBUGGER_REMOTE_PACKET{
		Indicator:                  sdk.IndicatorOfHyperdbgPacket,
		TypeOfThePacket:            sdk.DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnVmxRoot,
		RequestedActionOfThePacket: requestedAction,
	}

	checksumData := make([]byte, 0)
	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.LittleEndian, &packet)
	data := buf.Bytes()
	checksumData = append(checksumData, data[1:]...)
	packet.Checksum = ComputeDataChecksum(checksumData)
	packet.Checksum += ComputeDataChecksum(buffer)

	buf.Reset()
	binary.Write(buf, binary.LittleEndian, &packet)

	if kd.sendCommandAndBufferToDebuggee != nil {
		return kd.sendCommandAndBufferToDebuggee(packetType, requestedAction, buf.Bytes())
	}

	if kd.sendCommandToDebuggee != nil {
		combined := append(buf.Bytes(), buffer...)
		return kd.sendCommandToDebuggee(packetType, requestedAction, combined)
	}
	return false
}

func (kd *KernelDebugger) waitForKernelResponse(syncObject SyncObjectType) {
	if int(syncObject) >= 0 && int(syncObject) < KernelDebuggerMaxSyncObjects {
		kd.syncObjects[syncObject].mu.Lock()
		kd.syncObjects[syncObject].IsSignaled = false
		kd.syncObjects[syncObject].mu.Unlock()

		for {
			kd.syncObjects[syncObject].mu.Lock()
			signaled := kd.syncObjects[syncObject].IsSignaled
			kd.syncObjects[syncObject].mu.Unlock()

			if signaled {
				break
			}
		}
	}
}

func (kd *KernelDebugger) receivedKernelResponse(syncObject SyncObjectType) {
	if int(syncObject) >= 0 && int(syncObject) < KernelDebuggerMaxSyncObjects {
		kd.syncObjects[syncObject].mu.Lock()
		kd.syncObjects[syncObject].IsSignaled = true
		kd.syncObjects[syncObject].mu.Unlock()
	}
}

func (kd *KernelDebugger) uninitializeConnection() {
	kd.mu.Lock()
	defer kd.mu.Unlock()

	kd.currentRemoteCore = uint32(sdk.DebuggerRemotePacketRequestedActionDebuggeePausedAndCurrentInstruction)
	kd.isDebuggeeInHandshakingPhase = false
	kd.isSerialConnectedToRemoteDebugger = false
	kd.isSerialConnectedToRemoteDebuggee = false
	kd.isDebuggeeRunning = false
	kd.shouldPreviousCommandBeContinued = false
	kd.ignoreNewLoggingMessages = false
	kd.serialConnectionAlreadyClosed = true

	for i := range KernelDebuggerMaxSyncObjects {
		kd.syncObjects[i].mu.Lock()
		kd.syncObjects[i].IsSignaled = true
		kd.syncObjects[i].mu.Unlock()
	}
}

func (kd *KernelDebugger) checkWhetherCurrentInstructionIsCall(callInstructionSize *uint32) bool {
	if callInstructionSize == nil {
		return false
	}

	instr := kd.currentRunningInstruction[:]
	for i := range instr {
		b := instr[i]
		if b == 0xE8 {
			*callInstructionSize = 5
			return true
		}
		if b == 0x9A {
			*callInstructionSize = 7
			return true
		}
		if b == 0xFF {
			if i+1 < len(instr) {
				modrm := instr[i+1]
				reg := (modrm >> 3) & 0x7
				if reg == 2 {
					*callInstructionSize = 2
					return true
				}
			}
		}
		if b == 0x0F {
			break
		}
		if (b >= 0x50 && b <= 0x5F) || (b >= 0x70 && b <= 0x7F) || b == 0xC2 || b == 0xC3 || b == 0xCB || b == 0xCC || b == 0xCD {
			break
		}
	}
	return false
}

func (kd *KernelDebugger) IsDebuggeeRunning() bool {
	kd.mu.Lock()
	defer kd.mu.Unlock()
	return kd.isDebuggeeRunning
}

func (kd *KernelDebugger) SetDebuggeeRunning(running bool) {
	kd.mu.Lock()
	defer kd.mu.Unlock()
	kd.isDebuggeeRunning = running
}

func (kd *KernelDebugger) GetCurrentCore() uint32 {
	kd.mu.Lock()
	defer kd.mu.Unlock()
	return kd.currentRemoteCore
}

func (kd *KernelDebugger) GetKernelBaseAddress() uint64 {
	kd.mu.Lock()
	defer kd.mu.Unlock()
	return kd.kernelBaseAddress
}

func (kd *KernelDebugger) IsSerialConnected() bool {
	kd.mu.Lock()
	defer kd.mu.Unlock()
	return kd.isSerialConnectedToRemoteDebuggee
}
