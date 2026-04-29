package core

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"unicode"
	"unsafe"

	"github.com/ddkwork/HyperDbg/debugger/common"
	"github.com/ddkwork/HyperDbg/debugger/communication"
	"github.com/ddkwork/HyperDbg/debugger/misc"
	"github.com/ddkwork/HyperDbg/debugger/scriptengine"
	"github.com/ddkwork/HyperDbg/debugger/user_level"
	"github.com/ddkwork/sdk"
	"golang.org/x/sys/windows"
)

const (
	UseImmediateMessagingByDefaultOnEvents = true

	EventForwardingStateNotOpened = 0
	EventForwardingOpened         = 1
	EventForwardingClosed         = 2
)

type EventParsingErrorCause int

const (
	ParsingErrorSuccessfulNoError EventParsingErrorCause = iota
	ParsingErrorScriptSyntaxError
	ParsingErrorNoInput
	ParsingErrorMaximumInputReached
	ParsingErrorOutputNameNotFound
	ParsingErrorOutputSourceAlreadyClosed
	ParsingErrorAllocationError
	ParsingErrorFormatError
	ParsingErrorAttemptToBreakOnVmiMode
	ParsingErrorImmediateMessagingInEventForwardingMode
	ParsingErrorUsingShortCircuitingInPostEvents
)

type Breakpoint struct {
	ID          uint32
	Address     uint64
	ProcessId   uint32
	ThreadId    uint32
	CoreNumber  uint32
	IsEnabled   bool
	Is64Bit     bool
	Instruction [16]byte
}

type BreakpointManager struct {
	breakpoints map[uint32]*Breakpoint
	nextId      uint32
}

func NewBreakpointManager() *BreakpointManager {
	return &BreakpointManager{
		breakpoints: make(map[uint32]*Breakpoint),
		nextId:      1,
	}
}

func (bm *BreakpointManager) Set(address uint64, pid, tid, coreNum uint32) (*Breakpoint, error) {
	bp := &Breakpoint{
		ID:         bm.nextId,
		Address:    address,
		ProcessId:  pid,
		ThreadId:   tid,
		CoreNumber: coreNum,
		IsEnabled:  true,
		Is64Bit:    true,
	}
	bm.breakpoints[bp.ID] = bp
	bm.nextId++
	return bp, nil
}

func (bm *BreakpointManager) Clear(id uint32) error {
	if _, exists := bm.breakpoints[id]; !exists {
		return fmt.Errorf("breakpoint %d not found", id)
	}
	delete(bm.breakpoints, id)
	return nil
}

func (bm *BreakpointManager) Enable(id uint32) error {
	bp, exists := bm.breakpoints[id]
	if !exists {
		return fmt.Errorf("breakpoint %d not found", id)
	}
	bp.IsEnabled = true
	return nil
}

func (bm *BreakpointManager) Disable(id uint32) error {
	bp, exists := bm.breakpoints[id]
	if !exists {
		return fmt.Errorf("breakpoint %d not found", id)
	}
	bp.IsEnabled = false
	return nil
}

func (bm *BreakpointManager) List() []*Breakpoint {
	var result []*Breakpoint
	for _, bp := range bm.breakpoints {
		result = append(result, bp)
	}
	return result
}

func (bm *BreakpointManager) Get(id uint32) (*Breakpoint, bool) {
	bp, exists := bm.breakpoints[id]
	return bp, exists
}

type DebuggerCore struct {
	mu                                sync.RWMutex
	deviceHandle                      windows.Handle
	isSerialConnectedToRemoteDebuggee bool
	isSerialConnectedToRemoteDebugger bool
	isConnectedToRemoteDebuggee       bool
	isConnectedToRemoteDebugger       bool
	isConnectedToHyperDbgLocally      bool
	kernelBaseAddress                 uint64
	eventTag                          uint64
	outputSourceTag                   uint64
	autoUnpause                       bool
	breakPrintingOutput               bool
	eventTraceInitialized             bool
	outputSourcesInitialized          bool
	eventTrace                        common.ListEntry
	outputSources                     common.ListEntry
	activeProcessDebuggingState       user_level.ActiveDebuggingProcess

	forwardingManager *communication.ForwardingManager
	scriptEngine      *scriptengine.ScriptEngineWrapper
	assembler         *misc.Assembler

	sendIoctl                   func(ioctlCode uint32, inputBuffer []byte, outputBuffer []byte) (bool, error)
	deviceIoControl             func(ioctlCode uint32, inputBuffer []byte, inputSize uint32, outputBuffer []byte, outputSize uint32) (bool, uint32, error)
	sendRegisterEventToDebuggee func(event *sdk.DEBUGGER_GENERAL_EVENT_DETAIL, bufferLength uint32) (*sdk.DEBUGGER_EVENT_AND_ACTION_RESULT, error)
	sendAddActionToDebuggee     func(action *sdk.DEBUGGER_GENERAL_ACTION, bufferLength uint32) (*sdk.DEBUGGER_EVENT_AND_ACTION_RESULT, error)
	remoteConnectionSendCommand func(command string, length uint32)
}

func NewDebuggerCore() *DebuggerCore {
	dc := &DebuggerCore{
		eventTag:          uint64(sdk.DebuggerEventTagStartSeed),
		outputSourceTag:   uint64(sdk.DebuggerOutputSourceTagStartSeed),
		autoUnpause:       true,
		forwardingManager: communication.NewForwardingManager(),
		scriptEngine:      scriptengine.NewScriptEngineWrapper(),
		assembler:         misc.NewAssembler(),
	}
	common.InitializeListHead(&dc.eventTrace)
	common.InitializeListHead(&dc.outputSources)
	return dc
}

func (dc *DebuggerCore) SetDeviceHandle(handle windows.Handle) {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	dc.deviceHandle = handle
}

func (dc *DebuggerCore) GetDeviceHandle() windows.Handle {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return dc.deviceHandle
}

func (dc *DebuggerCore) SetSerialConnectedToRemoteDebuggee(connected bool) {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	dc.isSerialConnectedToRemoteDebuggee = connected
}

func (dc *DebuggerCore) SetSerialConnectedToRemoteDebugger(connected bool) {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	dc.isSerialConnectedToRemoteDebugger = connected
}

func (dc *DebuggerCore) SetConnectedToRemoteDebuggee(connected bool) {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	dc.isConnectedToRemoteDebuggee = connected
}

func (dc *DebuggerCore) SetConnectedToRemoteDebugger(connected bool) {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	dc.isConnectedToRemoteDebugger = connected
}

func (dc *DebuggerCore) SetConnectedToHyperDbgLocally(connected bool) {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	dc.isConnectedToHyperDbgLocally = connected
}

func (dc *DebuggerCore) IsConnectedToHyperDbgLocally() bool {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return dc.isConnectedToHyperDbgLocally
}

func (dc *DebuggerCore) SetActiveProcessDebuggingState(state user_level.ActiveDebuggingProcess) {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	dc.activeProcessDebuggingState = state
}

func (dc *DebuggerCore) SetDeviceIoControl(fn func(ioctlCode uint32, inputBuffer []byte, inputSize uint32, outputBuffer []byte, outputSize uint32) (bool, uint32, error)) {
	dc.deviceIoControl = fn
}

func (dc *DebuggerCore) SetSendRegisterEventToDebuggee(fn func(event *sdk.DEBUGGER_GENERAL_EVENT_DETAIL, bufferLength uint32) (*sdk.DEBUGGER_EVENT_AND_ACTION_RESULT, error)) {
	dc.sendRegisterEventToDebuggee = fn
}

func (dc *DebuggerCore) SetSendAddActionToDebuggee(fn func(action *sdk.DEBUGGER_GENERAL_ACTION, bufferLength uint32) (*sdk.DEBUGGER_EVENT_AND_ACTION_RESULT, error)) {
	dc.sendAddActionToDebuggee = fn
}

func (dc *DebuggerCore) SetRemoteConnectionSendCommand(fn func(command string, length uint32)) {
	dc.remoteConnectionSendCommand = fn
}

func (dc *DebuggerCore) SetBreakPrintingOutput(val bool) {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	dc.breakPrintingOutput = val
}

func (dc *DebuggerCore) IsBreakPrintingOutput() bool {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return dc.breakPrintingOutput
}

func (dc *DebuggerCore) SetAutoUnpause(val bool) {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	dc.autoUnpause = val
}

func (dc *DebuggerCore) IsAutoUnpause() bool {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return dc.autoUnpause
}

func showErrorMessage(errorCode uint32) {
	fmt.Printf("%s (0x%x)\n", sdk.DebuggerErrorCode(errorCode).String(), errorCode)
}

func (dc *DebuggerCore) GetNtoskrnlBase() uint64 {
	var ntoskrnlBase uint64

	modules, err := symbolCheckAndAllocateModuleInformation()
	if err != nil {
		fmt.Println("err, unable to get the module list for getting ntoskrnl base address")
		return 0
	}
	defer freeModules(modules)

	for i := uint32(0); i < modules.NumberOfModules; i++ {
		module := modules.Modules[i]
		fileName := getFileNameFromFullPath(module.FullPathName, module.OffsetToFileName)
		if strings.EqualFold(fileName, "ntoskrnl.exe") {
			ntoskrnlBase = uint64(module.ImageBase)
			break
		}
	}

	return ntoskrnlBase
}

func (dc *DebuggerCore) GetKernelBase() uint64 {
	dc.mu.Lock()
	defer dc.mu.Unlock()

	if dc.isSerialConnectedToRemoteDebuggee {
		return dc.kernelBaseAddress
	}

	kernelBase := dc.GetNtoskrnlBase()
	dc.kernelBaseAddress = kernelBase
	return kernelBase
}

func (dc *DebuggerCore) PauseDebuggee() bool {
	output := sdk.DEBUGGER_PAUSE_PACKET_RECEIVED{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPausePacketReceived,
		Output:  structToBytes(&output),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success {
		if err != nil {
			fmt.Printf("ioctl failed with error: %v\n", err)
		} else {
			fmt.Println("ioctl failed")
		}
		return false
	}
	if output.Result == uint32(sdk.DebuggerOperationWasSuccessful) {
		return true
	}
	showErrorMessage(output.Result)
	return false
}

func (dc *DebuggerCore) IsConnectedToAnyInstanceOfDebuggerOrDebuggee() bool {
	dc.mu.RLock()
	defer dc.mu.RUnlock()

	if dc.deviceHandle != 0 && dc.deviceHandle != windows.InvalidHandle {
		return true
	} else if dc.isConnectedToRemoteDebuggee {
		return true
	} else if dc.isConnectedToRemoteDebugger {
		return true
	} else if dc.isSerialConnectedToRemoteDebuggee {
		return true
	} else if dc.isSerialConnectedToRemoteDebugger {
		return true
	}

	return false
}

func (dc *DebuggerCore) ConnectionType() string {
	dc.mu.RLock()
	defer dc.mu.RUnlock()

	if dc.deviceHandle != 0 && dc.deviceHandle != windows.InvalidHandle {
		return "local"
	} else if dc.isConnectedToRemoteDebuggee {
		return "remote-debuggee"
	} else if dc.isConnectedToRemoteDebugger {
		return "remote-debugger"
	} else if dc.isSerialConnectedToRemoteDebuggee {
		return "serial-debuggee"
	} else if dc.isSerialConnectedToRemoteDebugger {
		return "serial-debugger"
	}
	return "disconnected"
}

func (dc *DebuggerCore) IsTagExist(tag uint64) bool {
	dc.mu.RLock()
	defer dc.mu.RUnlock()

	if !dc.eventTraceInitialized {
		return false
	}

	tempList := &dc.eventTrace
	for &dc.eventTrace != tempList.Blink {
		tempList = tempList.Blink
		eventDetail := containerOfEventDetail(tempList)
		if eventDetail.Tag == tag || tag == sdk.DebuggerModifyEventsApplyToAllTag {
			return true
		}
	}

	return false
}

func containerOfEventDetail(entry *common.ListEntry) *sdk.DEBUGGER_GENERAL_EVENT_DETAIL {
	base := uintptr(unsafe.Pointer(entry)) - unsafe.Offsetof(sdk.DEBUGGER_GENERAL_EVENT_DETAIL{}.CommandsEventList)
	return (*sdk.DEBUGGER_GENERAL_EVENT_DETAIL)(unsafe.Pointer(base))
}

func (dc *DebuggerCore) InterpretScript(tokens *[]CommandToken) (bool, *bool, *uint64, *uint32, *uint32, *uint64) {
	isTextVisited := false
	targetBracketString := ""
	var indexesToRemove []int
	index := 0

	scriptSyntaxErrors := true

	for i := range *tokens {
		section := (*tokens)[i]
		index++

		if isTextVisited && isTokenBracketString(section) {
			indexesToRemove = append(indexesToRemove, index)
			targetBracketString = getCaseSensitiveStringFromCommandToken(section)
			isTextVisited = false
			continue
		}

		if compareLowerCaseStrings(section.Value, "script") {
			indexesToRemove = append(indexesToRemove, index)
			isTextVisited = true
			continue
		}
	}

	if len(targetBracketString) == 0 {
		return false, nil, nil, nil, nil, nil
	}

	if strings.HasPrefix(targetBracketString, "file:") {
		filePath := targetBracketString[5:]
		content, err := os.ReadFile(filePath)
		if err != nil || len(content) == 0 {
			fmt.Println("err, either script file is not found or it's empty")
			scriptSyntaxErrors = true
			return true, &scriptSyntaxErrors, nil, nil, nil, nil
		}
		targetBracketString = string(content)
	}

	codeBuffer := dc.scriptEngine.GetEngine().ParseScript(targetBracketString, true)
	if codeBuffer == nil {
		scriptSyntaxErrors = true
		return true, &scriptSyntaxErrors, nil, nil, nil, nil
	}

	scriptSyntaxErrors = false

	var bufferAddress uint64
	var bufferLength uint32
	var pointer uint32
	var scriptCodeBuffer uint64

	if codeBuffer != nil {
		bufferAddress = codeBuffer.Address
		bufferLength = codeBuffer.Length
		pointer = codeBuffer.Pointer
		scriptCodeBuffer = codeBuffer.CodeBuffer
	}

	newIndexToRemove := 0
	for _, indexToRemove := range indexesToRemove {
		newIndexToRemove++
		idx := indexToRemove - newIndexToRemove
		if idx >= 0 && idx < len(*tokens) {
			*tokens = append((*tokens)[:idx], (*tokens)[idx+1:]...)
		}
	}

	return true, &scriptSyntaxErrors, &bufferAddress, &bufferLength, &pointer, &scriptCodeBuffer
}

func (dc *DebuggerCore) InterpretConditionsAndCodes(tokens *[]CommandToken, isConditionBuffer bool, bufferAddress *uint64, bufferLength *uint32) bool {
	isAsm := false
	isTextVisited := false
	targetBracketString := ""
	var indexesToRemove []int
	index := 0

	for i := range *tokens {
		section := (*tokens)[i]
		index++

		if isTextVisited && isTokenBracketString(section) {
			indexesToRemove = append(indexesToRemove, index)
			targetBracketString = getCaseSensitiveStringFromCommandToken(section)
			isTextVisited = false
			continue
		}

		if !isConditionBuffer && compareLowerCaseStrings(section.Value, "code") {
			indexesToRemove = append(indexesToRemove, index)
			isTextVisited = true
			continue
		}

		if isConditionBuffer && compareLowerCaseStrings(section.Value, "condition") {
			indexesToRemove = append(indexesToRemove, index)
			isTextVisited = true
			continue
		}

		if compareLowerCaseStrings(section.Value, "asm") {
			indexesToRemove = append(indexesToRemove, index)
			isAsm = true
			continue
		}
	}

	if len(targetBracketString) == 0 {
		return false
	}

	var localBufferAddress uint64
	var localBufferLength uint32

	if isAsm {
		asmData := dc.assembler.Assemble(targetBracketString + "; ret")
		if asmData == nil || !asmData.IsValid {
			fmt.Printf("err, assemble error\n")
			return false
		}

		finalBuffer := make([]byte, len(asmData.Opcode))
		copy(finalBuffer, asmData.Opcode)

		localBufferAddress = uint64(uintptr(unsafe.Pointer(&finalBuffer[0])))
		localBufferLength = uint32(len(finalBuffer))
	} else {
		targetBracketString += "c3"

		temp := targetBracketString
		if strings.HasPrefix(temp, "0x") || strings.HasPrefix(temp, "0X") || strings.HasPrefix(temp, "\\x") || strings.HasPrefix(temp, "\\X") {
			temp = temp[2:]
		} else if strings.HasPrefix(temp, "x") || strings.HasPrefix(temp, "X") {
			temp = temp[1:]
		}

		temp = strings.ReplaceAll(temp, "\\x", "")

		if len(temp)%2 != 0 {
			temp = "0" + temp
		}

		if !isHexNotation(temp) {
			fmt.Println("please enter condition code in a hex notation")
			return false
		}

		parsedBytes := hexToBytes(temp)
		if parsedBytes == nil {
			return false
		}

		finalBuffer := make([]byte, len(parsedBytes))
		copy(finalBuffer, parsedBytes)

		localBufferAddress = uint64(uintptr(unsafe.Pointer(&finalBuffer[0])))
		localBufferLength = uint32(len(parsedBytes))
	}

	if bufferAddress != nil {
		*bufferAddress = localBufferAddress
	}
	if bufferLength != nil {
		*bufferLength = localBufferLength
	}

	newIndexToRemove := 0
	for _, indexToRemove := range indexesToRemove {
		newIndexToRemove++
		idx := indexToRemove - newIndexToRemove
		if idx >= 0 && idx < len(*tokens) {
			*tokens = append((*tokens)[:idx], (*tokens)[idx+1:]...)
		}
	}

	return true
}

func (dc *DebuggerCore) InterpretConditionsAndCodesBuffer(tokens *[]CommandToken, isConditionBuffer bool) (bool, []byte) {
	isAsm := false
	isTextVisited := false
	targetBracketString := ""
	var indexesToRemove []int
	index := 0

	for i := range *tokens {
		section := (*tokens)[i]
		index++

		if isTextVisited && isTokenBracketString(section) {
			indexesToRemove = append(indexesToRemove, index)
			targetBracketString = getCaseSensitiveStringFromCommandToken(section)
			isTextVisited = false
			continue
		}

		if !isConditionBuffer && compareLowerCaseStrings(section.Value, "code") {
			indexesToRemove = append(indexesToRemove, index)
			isTextVisited = true
			continue
		}

		if isConditionBuffer && compareLowerCaseStrings(section.Value, "condition") {
			indexesToRemove = append(indexesToRemove, index)
			isTextVisited = true
			continue
		}

		if compareLowerCaseStrings(section.Value, "asm") {
			indexesToRemove = append(indexesToRemove, index)
			isAsm = true
			continue
		}
	}

	if len(targetBracketString) == 0 {
		return false, nil
	}

	var resultBuffer []byte

	if isAsm {
		asmData := dc.assembler.Assemble(targetBracketString + "; ret")
		if asmData == nil || !asmData.IsValid {
			fmt.Printf("err, assemble error\n")
			return false, nil
		}

		resultBuffer = make([]byte, len(asmData.Opcode))
		copy(resultBuffer, asmData.Opcode)
	} else {
		targetBracketString += "c3"

		temp := targetBracketString
		if strings.HasPrefix(temp, "0x") || strings.HasPrefix(temp, "0X") || strings.HasPrefix(temp, "\\x") || strings.HasPrefix(temp, "\\X") {
			temp = temp[2:]
		} else if strings.HasPrefix(temp, "x") || strings.HasPrefix(temp, "X") {
			temp = temp[1:]
		}

		temp = strings.ReplaceAll(temp, "\\x", "")

		if len(temp)%2 != 0 {
			temp = "0" + temp
		}

		if !isHexNotation(temp) {
			fmt.Println("please enter condition code in a hex notation")
			return false, nil
		}

		parsedBytes := hexToBytes(temp)
		if parsedBytes == nil {
			return false, nil
		}

		resultBuffer = make([]byte, len(parsedBytes))
		copy(resultBuffer, parsedBytes)
	}

	newIndexToRemove := 0
	for _, indexToRemove := range indexesToRemove {
		newIndexToRemove++
		idx := indexToRemove - newIndexToRemove
		if idx >= 0 && idx < len(*tokens) {
			*tokens = append((*tokens)[:idx], (*tokens)[idx+1:]...)
		}
	}

	return true, resultBuffer
}

func (dc *DebuggerCore) InterpretOutput(tokens *[]CommandToken) (bool, *[]string) {
	isTextVisited := false
	targetBracketString := ""
	var indexesToRemove []int
	index := 0

	for i := range *tokens {
		section := (*tokens)[i]
		index++

		if isTextVisited && isTokenBracketString(section) {
			indexesToRemove = append(indexesToRemove, index)
			targetBracketString = getCaseSensitiveStringFromCommandToken(section)
			isTextVisited = false
			continue
		}

		if compareLowerCaseStrings(section.Value, "output") {
			indexesToRemove = append(indexesToRemove, index)
			isTextVisited = true
			continue
		}
	}

	if len(targetBracketString) == 0 {
		return false, nil
	}

	var inputSources []string
	delimiter := ","

	if strings.Contains(targetBracketString, delimiter) {
		remaining := targetBracketString
		for {
			pos := strings.Index(remaining, delimiter)
			if pos == -1 {
				break
			}
			token := strings.TrimSpace(remaining[:pos])
			if len(token) > 0 {
				inputSources = append(inputSources, token)
			}
			remaining = remaining[pos+1:]
		}
		if len(strings.TrimSpace(remaining)) > 0 {
			inputSources = append(inputSources, strings.TrimSpace(remaining))
		}
	} else {
		inputSources = append(inputSources, targetBracketString)
	}

	newIndexToRemove := 0
	for _, indexToRemove := range indexesToRemove {
		newIndexToRemove++
		idx := indexToRemove - newIndexToRemove
		if idx >= 0 && idx < len(*tokens) {
			*tokens = append((*tokens)[:idx], (*tokens)[idx+1:]...)
		}
	}

	return true, &inputSources
}

func (dc *DebuggerCore) SendEventToKernel(event *sdk.DEBUGGER_GENERAL_EVENT_DETAIL, eventBufferLength uint32) bool {
	var returnedBuffer sdk.DEBUGGER_EVENT_AND_ACTION_RESULT

	if dc.isSerialConnectedToRemoteDebuggee {
		if dc.sendRegisterEventToDebuggee == nil {
			fmt.Println("err, register event callback not initialized for serial connection")
			return false
		}

		tempRegResult, err := dc.sendRegisterEventToDebuggee(event, eventBufferLength)
		if err != nil {
			fmt.Printf("err, failed to send register event to debuggee: %v\n", err)
			return false
		}
		returnedBuffer = *tempRegResult
	} else {
		if dc.deviceIoControl == nil {
			fmt.Println("err, device io control handler not initialized")
			return false
		}

		inputBuffer := serializeEventDetail(event, eventBufferLength)
		outputBuffer := make([]byte, uint32(unsafe.Sizeof(sdk.DEBUGGER_EVENT_AND_ACTION_RESULT{})))

		success, _, err := dc.deviceIoControl(
			sdk.IoctlDebuggerRegisterEvent,
			inputBuffer,
			eventBufferLength,
			outputBuffer,
			uint32(len(outputBuffer)),
		)

		if !success {
			if err != nil {
				fmt.Printf("ioctl failed with error: %v\n", err)
			} else {
				fmt.Println("ioctl failed")
			}
			return false
		}

		returnedBuffer = deserializeEventAndActionResult(outputBuffer)
	}

	if returnedBuffer.IsSuccessful && returnedBuffer.Error == 0 {
		dc.mu.Lock()
		if !dc.isSerialConnectedToRemoteDebuggee && !dc.isSerialConnectedToRemoteDebugger && dc.breakPrintingOutput && dc.autoUnpause {
			dc.breakPrintingOutput = false

			if dc.isConnectedToRemoteDebuggee && dc.remoteConnectionSendCommand != nil {
				dc.remoteConnectionSendCommand("g", uint32(len("g")+1))
			}

			fmt.Println()
		}
		dc.mu.Unlock()

		return true
	}

	if returnedBuffer.Error != 0 {
		showErrorMessage(returnedBuffer.Error)
	}

	return false
}

func (dc *DebuggerCore) RegisterActionToEvent(
	event *sdk.DEBUGGER_GENERAL_EVENT_DETAIL,
	actionBreakToDebugger *sdk.DEBUGGER_GENERAL_ACTION,
	actionBreakToDebuggerLength uint32,
	actionCustomCode *sdk.DEBUGGER_GENERAL_ACTION,
	actionCustomCodeLength uint32,
	actionScript *sdk.DEBUGGER_GENERAL_ACTION,
	actionScriptLength uint32,
) bool {
	var returnedBuffer sdk.DEBUGGER_EVENT_AND_ACTION_RESULT
	_ = returnedBuffer

	if dc.isSerialConnectedToRemoteDebuggee {
		if actionBreakToDebugger != nil {
			if dc.sendAddActionToDebuggee == nil {
				fmt.Println("err, add action callback not initialized for serial connection")
				return false
			}
			tempAddingResult, err := dc.sendAddActionToDebuggee(actionBreakToDebugger, actionBreakToDebuggerLength)
			if err != nil {
				fmt.Printf("err, failed to send break action to debuggee: %v\n", err)
				return false
			}
			returnedBuffer = *tempAddingResult
		}

		if actionCustomCode != nil {
			if dc.sendAddActionToDebuggee == nil {
				return false
			}
			tempAddingResult, err := dc.sendAddActionToDebuggee(actionCustomCode, actionCustomCodeLength)
			if err != nil {
				return false
			}
			returnedBuffer = *tempAddingResult
		}

		if actionScript != nil {
			if dc.sendAddActionToDebuggee == nil {
				return false
			}
			tempAddingResult, err := dc.sendAddActionToDebuggee(actionScript, actionScriptLength)
			if err != nil {
				return false
			}
			returnedBuffer = *tempAddingResult
		}
	} else {
		if dc.deviceIoControl == nil {
			fmt.Println("err, device io control handler not initialized")
			return false
		}

		if actionBreakToDebugger != nil {
			inputBuffer := serializeAction(actionBreakToDebugger, actionBreakToDebuggerLength)
			outputBuffer := make([]byte, uint32(unsafe.Sizeof(sdk.DEBUGGER_EVENT_AND_ACTION_RESULT{})))

			success, _, err := dc.deviceIoControl(
				sdk.IoctlDebuggerAddActionToEvent, inputBuffer,
				actionBreakToDebuggerLength,
				outputBuffer,
				uint32(len(outputBuffer)),
			)
			if !success {
				if err != nil {
					fmt.Printf("ioctl failed with error: %v\n", err)
				}
				return false
			}
			returnedBuffer = deserializeEventAndActionResult(outputBuffer)
		}

		if actionCustomCode != nil {
			inputBuffer := serializeAction(actionCustomCode, actionCustomCodeLength)
			outputBuffer := make([]byte, uint32(unsafe.Sizeof(sdk.DEBUGGER_EVENT_AND_ACTION_RESULT{})))

			success, _, err := dc.deviceIoControl(
				sdk.IoctlDebuggerAddActionToEvent,
				inputBuffer,
				actionCustomCodeLength,
				outputBuffer,
				uint32(len(outputBuffer)),
			)
			if !success {
				if err != nil {
					fmt.Printf("ioctl failed with error: %v\n", err)
				}
				return false
			}
			returnedBuffer = deserializeEventAndActionResult(outputBuffer)
		}

		if actionScript != nil {
			inputBuffer := serializeAction(actionScript, actionScriptLength)
			outputBuffer := make([]byte, uint32(unsafe.Sizeof(sdk.DEBUGGER_EVENT_AND_ACTION_RESULT{})))

			success, _, err := dc.deviceIoControl(
				sdk.IoctlDebuggerAddActionToEvent,
				inputBuffer,
				actionScriptLength,
				outputBuffer,
				uint32(len(outputBuffer)),
			)
			if !success {
				if err != nil {
					fmt.Printf("ioctl failed with error: %v\n", err)
				}
				return false
			}
			returnedBuffer = deserializeEventAndActionResult(outputBuffer)
		}
	}

	dc.FreeEventsAndActionsMemory(nil, actionBreakToDebugger, actionCustomCode, actionScript)

	dc.mu.Lock()
	if !dc.eventTraceInitialized {
		common.InitializeListHead(&dc.eventTrace)
		dc.eventTraceInitialized = true
	}
	common.InsertHeadList(&dc.eventTrace, (*common.ListEntry)(unsafe.Pointer(&event.CommandsEventList)))
	dc.mu.Unlock()

	return true
}

func (dc *DebuggerCore) GetNewDebuggerEventTag() uint64 {
	dc.mu.Lock()
	defer dc.mu.Unlock()

	tag := dc.eventTag
	dc.eventTag++
	return tag
}

func (dc *DebuggerCore) FreeEventsAndActionsMemory(
	event *sdk.DEBUGGER_GENERAL_EVENT_DETAIL,
	actionBreakToDebugger *sdk.DEBUGGER_GENERAL_ACTION,
	actionCustomCode *sdk.DEBUGGER_GENERAL_ACTION,
	actionScript *sdk.DEBUGGER_GENERAL_ACTION,
) {
}

func (dc *DebuggerCore) InterpretGeneralEventAndActionsFields(
	tokens *[]CommandToken,
	eventType uint32,
) (*sdk.DEBUGGER_GENERAL_EVENT_DETAIL, uint32, *sdk.DEBUGGER_GENERAL_ACTION, uint32, *sdk.DEBUGGER_GENERAL_ACTION, uint32, *sdk.DEBUGGER_GENERAL_ACTION, uint32, EventParsingErrorCause) {
	var tempEvent *sdk.DEBUGGER_GENERAL_EVENT_DETAIL
	var tempActionBreak *sdk.DEBUGGER_GENERAL_ACTION
	var tempActionScript *sdk.DEBUGGER_GENERAL_ACTION
	var tempActionCustomCode *sdk.DEBUGGER_GENERAL_ACTION

	callingStage := sdk.VmmCallbackCallingStagePreEventEmulation
	var lengthOfCustomCodeActionBuffer uint32
	var lengthOfScriptActionBuffer uint32
	var lengthOfBreakActionBuffer uint32
	var conditionBufferLength uint32
	var codeBufferLength uint32
	var scriptCodeBuffer uint64
	var scriptBufferLength uint32
	var scriptBufferPointer uint32
	var lengthOfEventBuffer uint32

	var isAShortCircuitingEventByDefault bool
	var indexOfValidSourceTags int
	var isSerialConnected bool
	hasConditionBuffer := false
	hasOutputPath := false
	hasCodeBuffer := false
	hasScript := false
	isNextCommandPid := false
	isNextCommandCoreId := false
	isNextCommandBufferSize := false
	isNextCommandImmediateMessaging := false
	isNextCommandExecutionStage := false
	isNextCommandSc := false
	immediateMessagePassing := UseImmediateMessagingByDefaultOnEvents
	var coreId uint32
	var processId uint32
	var requestBuffer uint32
	var listOfValidSourceTags []uint64
	var indexesToRemove []int
	newIndexToRemove := 0
	index := 0

	if condOk, condBuf := dc.InterpretConditionsAndCodesBuffer(tokens, true); condOk {
		hasConditionBuffer = true
		conditionBufferLength = uint32(len(condBuf))
	} else {
		hasConditionBuffer = false
	}

	if codeOk, codeBuf := dc.InterpretConditionsAndCodesBuffer(tokens, false); codeOk {
		hasCodeBuffer = true
		codeBufferLength = uint32(len(codeBuf))
	} else {
		hasCodeBuffer = false
	}

	scriptExists, scriptSyntaxErrors, _, scriptBufLen, scriptBufPtr, scriptCodeBuf := dc.InterpretScript(tokens)
	if !scriptExists {
		hasScript = false
	} else {
		if scriptSyntaxErrors != nil && *scriptSyntaxErrors {
			reasonForError := ParsingErrorScriptSyntaxError
			return nil, 0, nil, 0, nil, 0, nil, 0, reasonForError
		}
		hasScript = true
		if scriptBufLen != nil {
			scriptBufferLength = *scriptBufLen
		}
		if scriptBufPtr != nil {
			scriptBufferPointer = *scriptBufPtr
		}
		if scriptCodeBuf != nil {
			scriptCodeBuffer = *scriptCodeBuf
		}
	}

	var listOfOutputSources []string
	outputExists, outputSources := dc.InterpretOutput(tokens)
	if !outputExists {
		hasOutputPath = false
	} else {
		if outputSources == nil || len(*outputSources) == 0 {
			fmt.Println("err, no input found")
			return nil, 0, nil, 0, nil, 0, nil, 0, ParsingErrorNoInput
		}

		if uint32(len(*outputSources)) > sdk.DebuggerOutputSourceMaximumRemoteSourceForSingleEvent {
			fmt.Printf("err, based on this build of HyperDbg, the maximum input sources for a single event is 0x%x sources but you entered 0x%x sources\n",
				sdk.DebuggerOutputSourceMaximumRemoteSourceForSingleEvent, len(*outputSources))
			return nil, 0, nil, 0, nil, 0, nil, 0, ParsingErrorMaximumInputReached
		}

		dc.mu.RLock()
		sourcesInitialized := dc.outputSourcesInitialized
		dc.mu.RUnlock()

		if !sourcesInitialized {
			fmt.Println("err, the name you entered, not found. Did you use 'output' command to create it?")
			return nil, 0, nil, 0, nil, 0, nil, 0, ParsingErrorOutputNameNotFound
		}

		listOfOutputSources = *outputSources

		for _, item := range listOfOutputSources {
			outputSourceFound := false

			dc.mu.RLock()
			tempList := dc.outputSources.Flink
			for &dc.outputSources != tempList {
				currentOutputSource := containerOfEventForwarding(tempList)
				if currentOutputSource != nil && strings.EqualFold(currentOutputSource.Name, removeSpaces(item)) {
					if currentOutputSource.State == EventForwardingClosed {
						dc.mu.RUnlock()
						fmt.Println("err, output source already closed")
						return nil, 0, nil, 0, nil, 0, nil, 0, ParsingErrorOutputSourceAlreadyClosed
					}

					if currentOutputSource.State == EventForwardingStateNotOpened {
						fmt.Println("notice: some of the output(s) are not opened, it's not an error, but please ensure to open the output using the 'output' command to forward the results to the target resource")
					}

					outputSourceFound = true
					listOfValidSourceTags = append(listOfValidSourceTags, currentOutputSource.OutputUniqueTag)
					break
				}
				tempList = tempList.Flink
			}
			dc.mu.RUnlock()

			if !outputSourceFound {
				fmt.Println("err, the name you entered, not found. Did you use 'output' command to create it?")
				return nil, 0, nil, 0, nil, 0, nil, 0, ParsingErrorOutputNameNotFound
			}
		}

		hasOutputPath = true
	}

	lengthOfEventBuffer = uint32(unsafe.Sizeof(sdk.DEBUGGER_GENERAL_EVENT_DETAIL{})) + conditionBufferLength

	tempEvent = &sdk.DEBUGGER_GENERAL_EVENT_DETAIL{}
	tempEvent.IsEnabled = true
	tempEvent.Tag = dc.GetNewDebuggerEventTag()
	tempEvent.CoreId = sdk.DebuggerEventApplyToAllCores

	dc.mu.RLock()
	activeProcess := dc.activeProcessDebuggingState
	dc.mu.RUnlock()

	if activeProcess.IsActive {
		fmt.Printf("notice: as you're debugging a user-mode application, this event will only trigger on your current debugging process (pid:%x). If you want the event from the entire system, add 'pid all' to the event\n\n", activeProcess.ProcessId)
		tempEvent.ProcessId = activeProcess.ProcessId
	} else {
		tempEvent.ProcessId = sdk.DebuggerEventApplyToAllProcesses
	}

	tempEvent.EventType = sdk.VMM_EVENT_TYPE_ENUM(eventType)
	tempEvent.CommandStringBuffer = 0

	if hasConditionBuffer {
		tempEvent.ConditionBufferSize = conditionBufferLength
	}

	if hasCodeBuffer {
		lengthOfCustomCodeActionBuffer = uint32(unsafe.Sizeof(sdk.DEBUGGER_GENERAL_ACTION{})) + codeBufferLength
		tempActionCustomCode = &sdk.DEBUGGER_GENERAL_ACTION{}

		tempActionCustomCode.EventTag = tempEvent.Tag
		tempActionCustomCode.ActionType = sdk.RunCustomCode
		tempActionCustomCode.CustomCodeBufferSize = codeBufferLength
		tempEvent.CountOfActions++
	}

	if hasScript {
		lengthOfScriptActionBuffer = uint32(unsafe.Sizeof(sdk.DEBUGGER_GENERAL_ACTION{})) + scriptBufferLength
		tempActionScript = &sdk.DEBUGGER_GENERAL_ACTION{}

		tempActionScript.EventTag = tempEvent.Tag
		tempActionScript.ActionType = sdk.RunScript
		tempActionScript.ScriptBufferSize = scriptBufferLength
		tempActionScript.ScriptBufferPointer = scriptBufferPointer
		tempEvent.CountOfActions++

		dc.scriptEngine.GetEngine().RemoveSymbolBuffer(scriptCodeBuffer)
	}

	if !hasCodeBuffer && !hasScript {
		lengthOfBreakActionBuffer = uint32(unsafe.Sizeof(sdk.DEBUGGER_GENERAL_ACTION{}))
		tempActionBreak = &sdk.DEBUGGER_GENERAL_ACTION{}

		tempActionBreak.EventTag = tempEvent.Tag
		tempActionBreak.ActionType = sdk.BreakToDebugger
		tempEvent.CountOfActions++
	}

	for _, section := range *tokens {
		index++

		if isNextCommandBufferSize {
			parsedBuffer, err := convertTokenToUInt32(section.Value)
			if err != nil {
				fmt.Println("err, buffer size is invalid")
				goto ReturnWithError
			}
			requestBuffer = parsedBuffer

			if tempActionBreak != nil {
				tempActionBreak.PreAllocatedBuffer = requestBuffer
			}
			if tempActionScript != nil {
				tempActionScript.PreAllocatedBuffer = requestBuffer
			}
			if tempActionCustomCode != nil {
				tempActionCustomCode.PreAllocatedBuffer = requestBuffer
			}

			isNextCommandBufferSize = false
			indexesToRemove = append(indexesToRemove, index)
			continue
		}

		if isNextCommandImmediateMessaging {
			if compareLowerCaseStrings(section.Value, "yes") {
				immediateMessagePassing = true
			} else if compareLowerCaseStrings(section.Value, "no") {
				immediateMessagePassing = false
			} else {
				fmt.Println("err, immediate messaging token is invalid")
				goto ReturnWithError
			}

			isNextCommandImmediateMessaging = false
			indexesToRemove = append(indexesToRemove, index)
			continue
		}

		if isNextCommandExecutionStage {
			if compareLowerCaseStrings(section.Value, "pre") {
				callingStage = sdk.VmmCallbackCallingStagePreEventEmulation
			} else if compareLowerCaseStrings(section.Value, "post") {
				callingStage = sdk.VmmCallbackCallingStagePostEventEmulation
			} else if compareLowerCaseStrings(section.Value, "all") {
				callingStage = sdk.VmmCallbackCallingStageAllEventEmulation
			} else {
				fmt.Println("err, the specified execution mode is invalid; you can either choose 'pre' or 'post'")
				goto ReturnWithError
			}

			isNextCommandExecutionStage = false
			indexesToRemove = append(indexesToRemove, index)
			continue
		}

		if isNextCommandSc {
			if compareLowerCaseStrings(section.Value, "on") {
				isAShortCircuitingEventByDefault = true
			} else if compareLowerCaseStrings(section.Value, "off") {
				isAShortCircuitingEventByDefault = false
			} else {
				fmt.Println("err, the specified short-circuiting state is invalid; you can either choose 'on' or 'off'")
				goto ReturnWithError
			}

			isNextCommandSc = false
			indexesToRemove = append(indexesToRemove, index)
			continue
		}

		if isNextCommandPid {
			if compareLowerCaseStrings(section.Value, "all") {
				tempEvent.ProcessId = sdk.DebuggerEventApplyToAllProcesses
			} else {
				parsedPid, err := convertTokenToUInt32(section.Value)
				if err != nil {
					fmt.Println("err, pid is invalid")
					goto ReturnWithError
				}
				processId = parsedPid
				tempEvent.ProcessId = processId
			}

			isNextCommandPid = false
			indexesToRemove = append(indexesToRemove, index)
			continue
		}

		if isNextCommandCoreId {
			parsedCoreId, err := convertTokenToUInt32(section.Value)
			if err != nil {
				fmt.Println("err, core id is invalid")
				goto ReturnWithError
			}
			coreId = parsedCoreId
			tempEvent.CoreId = coreId

			isNextCommandCoreId = false
			indexesToRemove = append(indexesToRemove, index)
			continue
		}

		if compareLowerCaseStrings(section.Value, "pid") {
			isNextCommandPid = true
			indexesToRemove = append(indexesToRemove, index)
			continue
		}

		if compareLowerCaseStrings(section.Value, "core") {
			isNextCommandCoreId = true
			indexesToRemove = append(indexesToRemove, index)
			continue
		}

		if compareLowerCaseStrings(section.Value, "imm") {
			isNextCommandImmediateMessaging = true
			indexesToRemove = append(indexesToRemove, index)
			continue
		}

		if compareLowerCaseStrings(section.Value, "stage") {
			isNextCommandExecutionStage = true
			indexesToRemove = append(indexesToRemove, index)
			continue
		}

		if compareLowerCaseStrings(section.Value, "sc") {
			isNextCommandSc = true
			indexesToRemove = append(indexesToRemove, index)
			continue
		}

		if compareLowerCaseStrings(section.Value, "buffer") {
			isNextCommandBufferSize = true
			indexesToRemove = append(indexesToRemove, index)
			continue
		}
	}

	if isNextCommandCoreId {
		fmt.Println("err, please specify a value for 'core'")
		goto ReturnWithError
	}
	if isNextCommandPid {
		fmt.Println("err, please specify a value for 'pid'")
		goto ReturnWithError
	}
	if isNextCommandBufferSize {
		fmt.Println("err, please specify a value for 'buffer'")
		goto ReturnWithError
	}
	if isNextCommandImmediateMessaging {
		fmt.Println("err, please specify a value for 'imm'")
		goto ReturnWithError
	}
	if isNextCommandExecutionStage {
		fmt.Println("err, please specify a value for 'stage'")
		goto ReturnWithError
	}
	if isNextCommandSc {
		fmt.Println("err, please specify a value for 'sc'")
		goto ReturnWithError
	}

	if (callingStage == sdk.VmmCallbackCallingStagePostEventEmulation ||
		callingStage == sdk.VmmCallbackCallingStageAllEventEmulation) &&
		isAShortCircuitingEventByDefault {
		fmt.Println("err, using the short-circuiting mechanism with 'post' or 'all' stage events doesn't make sense; it's not supported!")
		goto ReturnWithError
	}

	dc.mu.RLock()
	isSerialConnected = dc.isSerialConnectedToRemoteDebuggee
	dc.mu.RUnlock()

	if !isSerialConnected && tempActionBreak != nil {
		fmt.Println("err, the script or assembly code is either not found or invalid. As a result, the default action is to break. However, breaking to the debugger is not possible in the VMI Mode. To achieve full control of the system, you can switch to the Debugger Mode. In the VMI Mode, you can still use scripts and run custom code for local debugging. For more information, please check: https://docs.hyperdbg.org/using-hyperdbg/prerequisites/operation-modes")
		goto ReturnWithError
	}

	if !immediateMessagePassing && hasOutputPath {
		fmt.Println("err, non-immediate message passing is not supported in 'output-forwarding mode'")
		goto ReturnWithError
	}

	if tempActionBreak != nil {
		tempActionBreak.ImmediateMessagePassing = immediateMessagePassing
	}
	if tempActionScript != nil {
		tempActionScript.ImmediateMessagePassing = immediateMessagePassing
	}
	if tempActionCustomCode != nil {
		tempActionCustomCode.ImmediateMessagePassing = immediateMessagePassing
	}

	indexOfValidSourceTags = 0
	for _, item := range listOfValidSourceTags {
		if uint32(indexOfValidSourceTags) < sdk.DebuggerOutputSourceMaximumRemoteSourceForSingleEvent {
			tempEvent.OutputSourceTags[indexOfValidSourceTags] = item
			indexOfValidSourceTags++
		}
	}

	if hasOutputPath {
		tempEvent.HasCustomOutput = true
	}

	if isAShortCircuitingEventByDefault {
		tempEvent.EnableShortCircuiting = true
	}

	tempEvent.EventStage = callingStage

	for _, indexToRemove := range indexesToRemove {
		newIndexToRemove++
		idx := indexToRemove - newIndexToRemove
		if idx >= 0 && idx < len(*tokens) {
			*tokens = append((*tokens)[:idx], (*tokens)[idx+1:]...)
		}
	}

	dc.mu.Lock()
	if !dc.eventTraceInitialized {
		common.InitializeListHead(&dc.eventTrace)
		dc.eventTraceInitialized = true
	}
	dc.mu.Unlock()

	return tempEvent, lengthOfEventBuffer, tempActionBreak, lengthOfBreakActionBuffer, tempActionCustomCode, lengthOfCustomCodeActionBuffer, tempActionScript, lengthOfScriptActionBuffer, ParsingErrorSuccessfulNoError

ReturnWithError:
	return nil, 0, nil, 0, nil, 0, nil, 0, ParsingErrorFormatError
}

// EventForwardingDetail is a Go-specific struct for managing event forwarding state
// This is not from C++ headers - it's used for remote debugging session management
type EventForwardingDetail struct {
	OutputSourcesList common.ListEntry
	Name              string
	OutputUniqueTag   uint64
	State             int
	Type              int
	Description       string
}

func containerOfEventForwarding(entry *common.ListEntry) *EventForwardingDetail {
	offset := uintptr(unsafe.Pointer(entry)) - uintptr(unsafe.Offsetof(EventForwardingDetail{}.OutputSourcesList))
	return (*EventForwardingDetail)(unsafe.Pointer(offset))
}

// RTLProcessModule represents Windows NT RTL_PROCESS_MODULE structure (from DDK/WDK)
// This is a standard Windows kernel structure, not HyperDbg-specific
type RTLProcessModule struct {
	FullPathName     [256]byte
	OffsetToFileName uint16
	ImageBase        uintptr
	ImageSize        uint32
}

// RTLProcessModules represents Windows NT RTL_PROCESS_MODULES structure (from DDK/WDK)
type RTLProcessModules struct {
	NumberOfModules uint32
	Modules         []RTLProcessModule
}

func symbolCheckAndAllocateModuleInformation() (*RTLProcessModules, error) {
	return nil, fmt.Errorf("not implemented on this platform")
}

func freeModules(modules *RTLProcessModules) {
}

func getFileNameFromFullPath(fullPathName [256]byte, offsetToFileName uint16) string {
	start := int(offsetToFileName)
	if start >= len(fullPathName) {
		return ""
	}

	var end int
	for end = start; end < len(fullPathName) && fullPathName[end] != 0; end++ {
	}

	return string(fullPathName[start:end])
}

func isTokenBracketString(token CommandToken) bool {
	return strings.Contains(token.Value, "{") && strings.Contains(token.Value, "}")
}

func getCaseSensitiveStringFromCommandToken(token CommandToken) string {
	if strings.Contains(token.Value, "{") && strings.Contains(token.Value, "}") {
		start := strings.Index(token.Value, "{")
		end := strings.LastIndex(token.Value, "}")
		if start != -1 && end != -1 && start < end {
			return token.Value[start+1 : end]
		}
	}
	return token.Value
}

func compareLowerCaseStrings(s string, target string) bool {
	return strings.EqualFold(s, target)
}

func isHexNotation(s string) bool {
	for _, c := range s {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return len(s) > 0
}

func hexToBytes(hexStr string) []byte {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil
	}
	return bytes
}

func removeSpaces(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

func convertTokenToUInt32(token string) (uint32, error) {
	token = strings.TrimSpace(token)

	if strings.HasPrefix(token, "0x") || strings.HasPrefix(token, "0X") {
		token = token[2:]
		val, err := strconv.ParseUint(token, 16, 32)
		if err != nil {
			return 0, err
		}
		return uint32(val), nil
	} else if strings.HasPrefix(token, "0n") || strings.HasPrefix(token, "0N") {
		token = token[2:]
		val, err := strconv.ParseUint(token, 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(val), nil
	}

	val, err := strconv.ParseUint(token, 16, 32)
	if err != nil {
		val2, err2 := strconv.ParseUint(token, 10, 32)
		if err2 != nil {
			return 0, fmt.Errorf("invalid number format: %s", token)
		}
		return uint32(val2), nil
	}
	return uint32(val), nil
}

func serializeEventDetail(event *sdk.DEBUGGER_GENERAL_EVENT_DETAIL, bufferLength uint32) []byte {
	buf := make([]byte, bufferLength)

	offset := uint32(0)

	binary.LittleEndian.PutUint32(buf[offset:], uint32(event.EventType))
	offset += 4

	binary.LittleEndian.PutUint64(buf[offset:], event.Tag)
	offset += 8

	binary.LittleEndian.PutUint32(buf[offset:], event.ConditionBufferSize)
	offset += 4

	binary.LittleEndian.PutUint32(buf[offset:], event.CoreId)
	offset += 4

	binary.LittleEndian.PutUint32(buf[offset:], event.ProcessId)
	offset += 4

	binary.LittleEndian.PutUint32(buf[offset:], event.CountOfActions)
	offset += 4

	if event.HasCustomOutput {
		buf[offset] = 1
	}
	offset++

	if event.EnableShortCircuiting {
		buf[offset] = 1
	}
	offset++

	binary.LittleEndian.PutUint32(buf[offset:], uint32(event.EventStage))
	offset += 4

	for i := range sdk.DebuggerOutputSourceMaximumRemoteSourceForSingleEvent {
		binary.LittleEndian.PutUint64(buf[offset:], event.OutputSourceTags[i])
		offset += 8
	}

	if event.IsEnabled {
		buf[offset] = 1
	}
	offset++

	return buf
}

func serializeAction(action *sdk.DEBUGGER_GENERAL_ACTION, bufferLength uint32) []byte {
	buf := make([]byte, bufferLength)

	offset := uint32(0)

	binary.LittleEndian.PutUint64(buf[offset:], action.EventTag)
	offset += 8

	binary.LittleEndian.PutUint32(buf[offset:], uint32(action.ActionType))
	offset += 4

	if action.ImmediateMessagePassing {
		buf[offset] = 1
	}
	offset++

	binary.LittleEndian.PutUint32(buf[offset:], action.CustomCodeBufferSize)
	offset += 4

	binary.LittleEndian.PutUint32(buf[offset:], action.ScriptBufferSize)
	offset += 4

	binary.LittleEndian.PutUint32(buf[offset:], action.ScriptBufferPointer)
	offset += 4

	binary.LittleEndian.PutUint32(buf[offset:], action.PreAllocatedBuffer)
	offset += 4

	return buf
}

func deserializeEventAndActionResult(data []byte) sdk.DEBUGGER_EVENT_AND_ACTION_RESULT {
	var result sdk.DEBUGGER_EVENT_AND_ACTION_RESULT
	if len(data) < 8 {
		return result
	}

	if data[0] != 0 {
		result.IsSuccessful = true
	}
	result.Error = binary.LittleEndian.Uint32(data[4:8])
	return result
}

type DebuggerIoctlRequest struct {
	Code    uint32
	Input   []byte
	Output  []byte
	InSize  uint32
	OutSize uint32
}

func structToBytes(v any) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, v)
	return buf.Bytes()
}

func (dc *DebuggerCore) SendIoctlRequest(req DebuggerIoctlRequest) (bool, uint32, error) {
	if dc.deviceIoControl == nil {
		return false, 0, fmt.Errorf("device io control handler not initialized")
	}
	return dc.deviceIoControl(req.Code, req.Input, req.InSize, req.Output, req.OutSize)
}

func (dc *DebuggerCore) AttachToProcess(pid uint32, isStartingNew bool, checkCallback bool, is32Bit bool) error {
	input := sdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS{
		ProcessId: pid,
		Action:    sdk.DebuggerAttachDetachUserModeProcessActionAttach,
	}
	if isStartingNew {
		input.IsStartingNewProcess = true
	}
	if checkCallback {
		input.CheckCallbackAtFirstInstruction = true
	}
	if is32Bit {
		input.Is32Bit = true
	}
	output := sdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerAttachDetachUserModeProcess,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("attach failed: %v", err)
	}
	if output.Result != uint64(sdk.DebuggerOperationWasSuccessful) {
		return fmt.Errorf("attach failed with kernel status: 0x%x", output.Result)
	}
	return nil
}

func (dc *DebuggerCore) DetachFromProcess(pid uint32) error {
	input := sdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS{
		ProcessId: pid,
		Action:    sdk.DebuggerAttachDetachUserModeProcessActionDetach,
	}
	output := sdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerAttachDetachUserModeProcess,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("detach failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) SetBreakpoint(address uint64, pid, tid, coreNum uint32) error {
	input := sdk.DEBUGGEE_BP_PACKET{
		Address: address,
		Pid:     pid,
		Tid:     tid,
		Core:    coreNum,
	}
	output := sdk.DEBUGGEE_BP_PACKET{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlSetBreakpointUserDebugger,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("set breakpoint failed: %v", err)
	}
	if output.Result != uint32(sdk.DebuggerOperationWasSuccessful) {
		return fmt.Errorf("set breakpoint failed with result: 0x%x", output.Result)
	}
	return nil
}

func (dc *DebuggerCore) ContinueDebuggee() error {
	output := sdk.DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlSendSignalExecutionInDebuggeeFinished,
		Output:  structToBytes(&output),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("continue failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) ReadMemory(address uint64, size uint32, pid uint32, memoryType uint8, readingType uint8) ([]byte, error) {
	input := sdk.DEBUGGER_READ_MEMORY{
		Pid:            pid,
		Address:        address,
		Size:           size,
		GetAddressMode: true,
		AddressMode:    sdk.DebuggerReadAddressMode64Bit,
		MemoryType:     sdk.DEBUGGER_READ_MEMORY_TYPE(memoryType),
		ReadingType:    sdk.DEBUGGER_READ_READING_TYPE(readingType),
	}
	headerSize := uint32(unsafe.Sizeof(input))
	inputBuf := make([]byte, headerSize+size)
	copy(inputBuf[:headerSize], structToBytes(&input))
	outputBuf := make([]byte, headerSize+size)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerReadMemory,
		Input:   inputBuf,
		Output:  outputBuf,
		InSize:  headerSize + size,
		OutSize: headerSize + size,
	})
	if !success || err != nil {
		return nil, fmt.Errorf("read memory failed: %v", err)
	}
	return outputBuf[headerSize : headerSize+size], nil
}

func (dc *DebuggerCore) EditMemory(address uint64, data []byte, pid uint32, memoryType uint32, byteSize uint32) error {
	chunkCount := (len(data) + 7) / 8
	if chunkCount == 0 {
		chunkCount = 1
	}
	header := sdk.DEBUGGER_EDIT_MEMORY{
		Address:         address,
		ProcessId:       pid,
		MemoryType:      sdk.DEBUGGER_EDIT_MEMORY_TYPE(memoryType),
		ByteSize:        sdk.DEBUGGER_EDIT_MEMORY_BYTE_SIZE(byteSize),
		CountOf64Chunks: uint32(chunkCount),
	}
	finalSize := uint32(unsafe.Sizeof(header)) + uint32(chunkCount)*8
	header.FinalStructureSize = finalSize
	inputBuf := make([]byte, finalSize)
	copy(inputBuf[:unsafe.Sizeof(header)], structToBytes(&header))
	copy(inputBuf[unsafe.Sizeof(header):], data)
	outputBuf := make([]byte, 4)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerEditMemory,
		Input:   inputBuf,
		Output:  outputBuf,
		InSize:  finalSize,
		OutSize: 4,
	})
	if !success || err != nil {
		return fmt.Errorf("edit memory failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) ReadOrWriteMsr(msr uint32, coreNum uint32, actionType uint32, value uint64) (uint64, error) {
	input := sdk.DEBUGGER_READ_AND_WRITE_ON_MSR{
		Msr:        uint64(msr),
		CoreNumber: coreNum,
		ActionType: sdk.DEBUGGER_MSR_ACTION_TYPE(actionType),
		Value:      value,
	}
	output := sdk.DEBUGGER_READ_AND_WRITE_ON_MSR{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerReadOrWriteMsr,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return 0, fmt.Errorf("msr operation failed: %v", err)
	}
	return output.Value, nil
}

func (dc *DebuggerCore) SearchMemory(address uint64, length uint64, pid uint32, memoryType uint32, byteSize uint32, pattern []byte) ([]uint64, error) {
	chunkCount := (len(pattern) + 7) / 8
	if chunkCount == 0 {
		chunkCount = 1
	}
	header := sdk.DEBUGGER_SEARCH_MEMORY{
		Address:         address,
		Length:          length,
		ProcessId:       pid,
		MemoryType:      sdk.DEBUGGER_SEARCH_MEMORY_TYPE(memoryType),
		ByteSize:        sdk.DEBUGGER_SEARCH_MEMORY_BYTE_SIZE(byteSize),
		CountOf64Chunks: uint32(chunkCount),
	}
	headerSize := uint32(unsafe.Sizeof(header))
	finalSize := headerSize + uint32(chunkCount)*8
	header.FinalStructureSize = finalSize
	inputBuf := make([]byte, finalSize)
	copy(inputBuf[:headerSize], structToBytes(&header))
	copy(inputBuf[headerSize:], pattern)
	outputBuf := make([]byte, 4096)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerSearchMemory,
		Input:   inputBuf,
		Output:  outputBuf,
		InSize:  finalSize,
		OutSize: 4096,
	})
	if !success || err != nil {
		return nil, fmt.Errorf("search memory failed: %v", err)
	}
	var results []uint64
	for i := uint32(0); i < 4096; i += 8 {
		addr := binary.LittleEndian.Uint64(outputBuf[i : i+8])
		if addr == 0 {
			break
		}
		results = append(results, addr)
	}
	return results, nil
}

func (dc *DebuggerCore) BringPagesIn(from, to uint64, pid uint32, pageFaultCode uint32) error {
	input := sdk.DEBUGGER_PAGE_IN_REQUEST{
		VirtualAddressFrom: from,
		VirtualAddressTo:   to,
		ProcessId:          pid,
		PageFaultErrorCode: pageFaultCode,
	}
	output := sdk.DEBUGGER_PAGE_IN_REQUEST{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerBringPagesIn,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("page in failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) PerformActionsOnApic(isLocalApic bool, apicPage []byte) error {
	apicType := sdk.DebuggerApicRequestTypeReadIoApic
	if isLocalApic {
		apicType = sdk.DebuggerApicRequestTypeReadLocalApic
	}
	header := sdk.DEBUGGER_APIC_REQUEST{
		ApicType: apicType,
	}
	headerSize := uint32(unsafe.Sizeof(header))
	requestSize := headerSize + uint32(len(apicPage))
	inputBuf := make([]byte, requestSize)
	copy(inputBuf[:headerSize], structToBytes(&header))
	copy(inputBuf[headerSize:], apicPage)
	outputBuf := make([]byte, 8)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformActionsOnApic,
		Input:   inputBuf,
		Output:  outputBuf,
		InSize:  requestSize,
		OutSize: 8,
	})
	if !success || err != nil {
		return fmt.Errorf("apic operation failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) QueryIdtEntry(idtPacket []byte) error {
	output := make([]byte, len(idtPacket)+4)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlQueryIdtEntry,
		Output:  output,
		OutSize: uint32(len(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("query idt failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) PerformSmiOperation(operationType uint32, address uint64, value uint64, count uint32, processId uint32) error {
	input := sdk.SMI_OPERATION_PACKETS{
		SmiOperationType: sdk.SMI_OPERATION_REQUEST_TYPE(operationType),
		SmiCount:         value,
	}
	output := sdk.SMI_OPERATION_PACKETS{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformSmiOperation,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("smi operation failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) HideTransparentMode(isHide bool, usePid bool, pid uint32, processName string) error {
	input := sdk.DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE{}
	if isHide {
		input.IsHide = true
	}
	if usePid {
		input.TrueIfProcessIdAndFalseIfProcessName = true
	}
	input.ProcId = pid
	nameBytes := []byte(processName)
	if len(nameBytes) > 0 {
		input.LengthOfProcessName = uint32(len(nameBytes))
	}
	output := sdk.DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerHideAndUnhideToTransparentTheDebugger,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("transparent mode failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) FlushLoggingBuffers() (uint32, uint32, error) {
	output := sdk.DEBUGGER_FLUSH_LOGGING_BUFFERS{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerFlushLoggingBuffers,
		Output:  structToBytes(&output),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return 0, 0, fmt.Errorf("flush buffers failed: %v", err)
	}
	return output.CountOfMessagesThatSetAsReadFromVmxRoot, output.CountOfMessagesThatSetAsReadFromVmxNonRoot, nil
}

func (dc *DebuggerCore) GetActiveThreadsAndProcesses() ([]byte, error) {
	output := make([]byte, 65536)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlGetDetailOfActiveThreadsAndProcesses,
		Output:  output,
		OutSize: 65536,
	})
	if !success || err != nil {
		return nil, fmt.Errorf("get threads/processes failed: %v", err)
	}
	return output, nil
}

func (dc *DebuggerCore) GetModules() ([]byte, error) {
	output := make([]byte, 65536)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlGetUserModeModuleDetails,
		Output:  output,
		OutSize: 65536,
	})
	if !success || err != nil {
		return nil, fmt.Errorf("get modules failed: %v", err)
	}
	return output, nil
}

func (dc *DebuggerCore) QueryCountOfActiveProcesses() (uint32, error) {
	input := sdk.DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS{
		QueryType:   sdk.DebuggerQueryActiveProcessesOrThreadsQueryProcessCount,
		QueryAction: sdk.DebuggerQueryActiveProcessesOrThreadsActionQueryCount,
	}
	output := sdk.DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlQueryCountOfActiveProcessesOrThreads,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return 0, fmt.Errorf("query count failed: %v", err)
	}
	return output.Count, nil
}

func (dc *DebuggerCore) PreAllocatePool(poolType uint32, count uint32) error {
	input := sdk.DEBUGGER_PREALLOC_COMMAND{
		Type:  sdk.DEBUGGER_PREALLOC_COMMAND_TYPE(poolType),
		Count: count,
	}
	output := sdk.DEBUGGER_PREALLOC_COMMAND{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlReservePreAllocatedPools,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("preallocate pool failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) PreactivateFunctionality(funcType uint32) error {
	input := sdk.DEBUGGER_PREACTIVATE_COMMAND{
		Type: sdk.DEBUGGER_PREACTIVATE_COMMAND_TYPE(funcType),
	}
	output := sdk.DEBUGGER_PREACTIVATE_COMMAND{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPreactivateFunctionality,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("preactivate functionality failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) Cpuid(leaf uint32) ([4]uint32, error) {
	input := make([]byte, 4)
	binary.LittleEndian.PutUint32(input[0:4], leaf)
	output := make([]byte, 16)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformKernelSideTests,
		Input:   input,
		Output:  output,
		InSize:  4,
		OutSize: 16,
	})
	if !success || err != nil {
		return [4]uint32{}, fmt.Errorf("cpuid failed: %v", err)
	}
	return [4]uint32{
		binary.LittleEndian.Uint32(output[0:4]),
		binary.LittleEndian.Uint32(output[4:8]),
		binary.LittleEndian.Uint32(output[8:12]),
		binary.LittleEndian.Uint32(output[12:16]),
	}, nil
}

func (dc *DebuggerCore) CrRead(crNumber uint32) (uint64, error) {
	input := make([]byte, 8)
	binary.LittleEndian.PutUint32(input[0:4], crNumber)
	output := make([]byte, 16)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformKernelSideTests,
		Input:   input,
		Output:  output,
		InSize:  8,
		OutSize: 16,
	})
	if !success || err != nil {
		return 0, fmt.Errorf("cr read failed: %v", err)
	}
	return binary.LittleEndian.Uint64(output[0:8]), nil
}

func (dc *DebuggerCore) CrWrite(crNumber uint32, value uint64) error {
	input := make([]byte, 16)
	binary.LittleEndian.PutUint32(input[0:4], crNumber)
	binary.LittleEndian.PutUint64(input[8:16], value)
	output := make([]byte, 8)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformKernelSideTests,
		Input:   input,
		Output:  output,
		InSize:  16,
		OutSize: 8,
	})
	if !success || err != nil {
		return fmt.Errorf("cr write failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) DrRead(drNum int) (uint64, error) {
	input := make([]byte, 4)
	binary.LittleEndian.PutUint32(input[0:4], uint32(drNum))
	output := make([]byte, 16)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformKernelSideTests,
		Input:   input,
		Output:  output,
		InSize:  4,
		OutSize: 16,
	})
	if !success || err != nil {
		return 0, fmt.Errorf("dr read failed: %v", err)
	}
	return binary.LittleEndian.Uint64(output[0:8]), nil
}

func (dc *DebuggerCore) DrWrite(drNum int, value uint64) error {
	input := make([]byte, 12)
	binary.LittleEndian.PutUint32(input[0:4], uint32(drNum))
	binary.LittleEndian.PutUint64(input[4:12], value)
	output := make([]byte, 8)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformKernelSideTests,
		Input:   input,
		Output:  output,
		InSize:  12,
		OutSize: 8,
	})
	if !success || err != nil {
		return fmt.Errorf("dr write failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) IoPortIn(port uint16) (uint32, error) {
	input := make([]byte, 2)
	binary.LittleEndian.PutUint16(input[0:2], port)
	output := make([]byte, 8)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformKernelSideTests,
		Input:   input,
		Output:  output,
		InSize:  2,
		OutSize: 8,
	})
	if !success || err != nil {
		return 0, fmt.Errorf("ioport in failed: %v", err)
	}
	return binary.LittleEndian.Uint32(output[0:4]), nil
}

func (dc *DebuggerCore) IoPortOut(port uint16, value uint32) error {
	input := make([]byte, 6)
	binary.LittleEndian.PutUint16(input[0:2], port)
	binary.LittleEndian.PutUint32(input[2:6], value)
	output := make([]byte, 8)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformKernelSideTests,
		Input:   input,
		Output:  output,
		InSize:  6,
		OutSize: 8,
	})
	if !success || err != nil {
		return fmt.Errorf("ioport out failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) PteDetails(virtualAddress uint64, pid uint32) ([]byte, error) {
	input := sdk.DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS{
		VirtualAddress: virtualAddress,
		ProcessId:      pid,
	}
	output := sdk.DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerReadPageTableEntriesDetails,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return nil, fmt.Errorf("pte details failed: %v", err)
	}
	return structToBytes(&output), nil
}

func (dc *DebuggerCore) Va2Pa(virtualAddress uint64, pid uint32) (uint64, error) {
	input := sdk.DEBUGGER_VA2PA_AND_PA2VA_COMMANDS{
		VirtualAddress:     virtualAddress,
		ProcessId:          pid,
		IsVirtual2Physical: true,
	}
	output := sdk.DEBUGGER_VA2PA_AND_PA2VA_COMMANDS{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerVa2paAndPa2vaCommands,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return 0, fmt.Errorf("va2pa failed: %v", err)
	}
	return output.PhysicalAddress, nil
}

func (dc *DebuggerCore) Pa2Va(physicalAddress uint64, pid uint32) (uint64, error) {
	input := sdk.DEBUGGER_VA2PA_AND_PA2VA_COMMANDS{
		PhysicalAddress: physicalAddress,
		ProcessId:       pid,
	}
	output := sdk.DEBUGGER_VA2PA_AND_PA2VA_COMMANDS{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerVa2paAndPa2vaCommands,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return 0, fmt.Errorf("pa2va failed: %v", err)
	}
	return output.VirtualAddress, nil
}

func (dc *DebuggerCore) RevMachineService(address uint64, size uint32, style uint32) (uint64, error) {
	input := sdk.REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST{
		ProcessId: uint32(address),
		Size:      size,
		Mode:      sdk.REVERSING_MACHINE_RECONSTRUCT_MEMORY_MODE(style),
		Type:      sdk.ReversingMachineReconstructMemoryTypeReconstruct,
	}
	output := sdk.REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlRequestRevMachineService,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return 0, fmt.Errorf("rev machine service failed: %v", err)
	}
	return uint64(output.KernelStatus), nil
}

func (dc *DebuggerCore) PciEndpointEnum() ([]byte, error) {
	output := make([]byte, 65536)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPcieEndpointEnum,
		Output:  output,
		OutSize: 65536,
	})
	if !success || err != nil {
		return nil, fmt.Errorf("pci endpoint enum failed: %v", err)
	}
	return output, nil
}

func (dc *DebuggerCore) PciDevInfoEnum(bus, device, function uint8) ([]byte, error) {
	input := [4]byte{bus, device, function}
	output := make([]byte, 4096)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPcidevinfoEnum,
		Input:   input[:],
		Output:  output,
		InSize:  4,
		OutSize: 4096,
	})
	if !success || err != nil {
		return nil, fmt.Errorf("pci devinfo enum failed: %v", err)
	}
	return output[:4096], nil
}

func (dc *DebuggerCore) TscRead() (uint64, error) {
	output := make([]byte, 16)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformKernelSideTests,
		Output:  output,
		OutSize: 16,
	})
	if !success || err != nil {
		return 0, fmt.Errorf("tsc read failed: %v", err)
	}
	return binary.LittleEndian.Uint64(output[0:8]), nil
}

func (dc *DebuggerCore) Vmcall(rax, rcx, rdx uint64) error {
	input := make([]byte, 24)
	binary.LittleEndian.PutUint64(input[0:8], rax)
	binary.LittleEndian.PutUint64(input[8:16], rcx)
	binary.LittleEndian.PutUint64(input[16:24], rdx)
	output := make([]byte, 8)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformKernelSideTests,
		Input:   input,
		Output:  output,
		InSize:  24,
		OutSize: 8,
	})
	if !success || err != nil {
		return fmt.Errorf("vmcall failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) Xsetbv(index uint32, low, high uint64) error {
	input := make([]byte, 20)
	binary.LittleEndian.PutUint32(input[0:4], index)
	binary.LittleEndian.PutUint64(input[4:12], low)
	binary.LittleEndian.PutUint64(input[12:20], high)
	output := make([]byte, 8)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformKernelSideTests,
		Input:   input,
		Output:  output,
		InSize:  20,
		OutSize: 8,
	})
	if !success || err != nil {
		return fmt.Errorf("xsetbv failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) InterruptFire(interruptIndex uint32) error {
	input := make([]byte, 4)
	binary.LittleEndian.PutUint32(input[0:4], interruptIndex)
	output := make([]byte, 8)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformKernelSideTests,
		Input:   input,
		Output:  output,
		InSize:  4,
		OutSize: 8,
	})
	if !success || err != nil {
		return fmt.Errorf("interrupt fire failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) PmcRead(pmcEvent uint32) (uint64, error) {
	input := make([]byte, 4)
	binary.LittleEndian.PutUint32(input[0:4], pmcEvent)
	output := make([]byte, 16)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformKernelSideTests,
		Input:   input,
		Output:  output,
		InSize:  4,
		OutSize: 16,
	})
	if !success || err != nil {
		return 0, fmt.Errorf("pmc read failed: %v", err)
	}
	return binary.LittleEndian.Uint64(output[0:8]), nil
}

func (dc *DebuggerCore) LbrOperation(enable bool) error {
	opType := sdk.HypertraceLbrOperationRequestTypeDisable
	if enable {
		opType = sdk.HypertraceLbrOperationRequestTypeEnable
	}
	packet := sdk.HYPERTRACE_LBR_OPERATION_PACKETS{
		LbrOperationType: opType,
	}
	output := sdk.HYPERTRACE_LBR_OPERATION_PACKETS{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformHypertraceLbrOperation,
		Input:   structToBytes(&packet),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(packet)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("lbr operation failed: %v", err)
	}
	if output.KernelStatus != uint32(sdk.DebuggerOperationWasSuccessful) {
		return fmt.Errorf("lbr operation failed with status: 0x%x", output.KernelStatus)
	}
	return nil
}

func (dc *DebuggerCore) Measure(startAddr, endAddr uint32) (uint64, error) {
	input := make([]byte, 8)
	binary.LittleEndian.PutUint32(input[0:4], startAddr)
	binary.LittleEndian.PutUint32(input[4:8], endAddr)
	output := make([]byte, 16)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformKernelSideTests,
		Input:   input,
		Output:  output,
		InSize:  8,
		OutSize: 16,
	})
	if !success || err != nil {
		return 0, fmt.Errorf("measure failed: %v", err)
	}
	return binary.LittleEndian.Uint64(output[0:8]), nil
}

func (dc *DebuggerCore) PerformKernelTest(size uint32) ([]byte, error) {
	output := make([]byte, size)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformKernelSideTests,
		Output:  output,
		OutSize: size,
	})
	if !success || err != nil {
		return nil, fmt.Errorf("kernel test failed: %v", err)
	}
	return output, nil
}

func (dc *DebuggerCore) HypertraceLbrOperation(opType sdk.HYPERTRACE_LBR_OPERATION_REQUEST_TYPE) error {
	input := sdk.HYPERTRACE_LBR_OPERATION_PACKETS{
		LbrOperationType: opType,
	}
	output := sdk.HYPERTRACE_LBR_OPERATION_PACKETS{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformHypertraceLbrOperation,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("hypertrace lbr operation failed: %v", err)
	}
	if output.KernelStatus != uint32(sdk.DebuggerOperationWasSuccessful) {
		return fmt.Errorf("hypertrace lbr operation failed with status: 0x%x", output.KernelStatus)
	}
	return nil
}

func (dc *DebuggerCore) HypertracePtOperation(opType sdk.HYPERTRACE_PT_OPERATION_REQUEST_TYPE) error {
	input := sdk.HYPERTRACE_PT_OPERATION_PACKETS{
		PtOperationType: opType,
	}
	output := sdk.HYPERTRACE_PT_OPERATION_PACKETS{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformHypertracePtOperation,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("hypertrace pt operation failed: %v", err)
	}
	if output.KernelStatus != uint32(sdk.DebuggerOperationWasSuccessful) {
		return fmt.Errorf("hypertrace pt operation failed with status: 0x%x", output.KernelStatus)
	}
	return nil
}

func (dc *DebuggerCore) TerminateVmx() error {
	output := sdk.DEBUGGER_PERFORM_KERNEL_TESTS{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlTerminateVmx,
		Output:  structToBytes(&output),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("terminate vmx failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) QueryCurrentProcess() ([]byte, error) {
	output := make([]byte, 1024)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlQueryCurrentProcess,
		Output:  output,
		OutSize: 1024,
	})
	if !success || err != nil {
		return nil, fmt.Errorf("query current process failed: %v", err)
	}
	return output, nil
}

func (dc *DebuggerCore) QueryCurrentThread() ([]byte, error) {
	output := make([]byte, 1024)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlQueryCurrentThread,
		Output:  output,
		OutSize: 1024,
	})
	if !success || err != nil {
		return nil, fmt.Errorf("query current thread failed: %v", err)
	}
	return output, nil
}

func (dc *DebuggerCore) ClearBreakpoint(id uint32) error {
	input := sdk.DEBUGGEE_BP_LIST_OR_MODIFY_PACKET{
		BreakpointId: uint64(id),
		Request:      sdk.DebuggeeBreakpointModificationRequestClear,
	}
	output := sdk.DEBUGGEE_BP_LIST_OR_MODIFY_PACKET{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlSetBreakpointUserDebugger,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("clear breakpoint %d failed: %v", id, err)
	}
	return nil
}

func (dc *DebuggerCore) DisableBreakpoint(id uint32) error {
	input := sdk.DEBUGGEE_BP_LIST_OR_MODIFY_PACKET{
		BreakpointId: uint64(id),
		Request:      sdk.DebuggeeBreakpointModificationRequestDisable,
	}
	output := sdk.DEBUGGEE_BP_LIST_OR_MODIFY_PACKET{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlSetBreakpointUserDebugger,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("disable breakpoint %d failed: %v", id, err)
	}
	return nil
}

func (dc *DebuggerCore) EnableBreakpoint(id uint32) error {
	input := sdk.DEBUGGEE_BP_LIST_OR_MODIFY_PACKET{
		BreakpointId: uint64(id),
		Request:      sdk.DebuggeeBreakpointModificationRequestEnable,
	}
	output := sdk.DEBUGGEE_BP_LIST_OR_MODIFY_PACKET{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlSetBreakpointUserDebugger,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("enable breakpoint %d failed: %v", id, err)
	}
	return nil
}

func (dc *DebuggerCore) ListBreakpoints() []*Breakpoint {
	input := make([]byte, 16)
	binary.LittleEndian.PutUint64(input[0:8], 0)
	binary.LittleEndian.PutUint32(input[8:12], uint32(sdk.DebuggeeBreakpointModificationRequestListBreakpoints))
	output := make([]byte, 4096)
	success, _, _ := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlSetBreakpointUserDebugger,
		Input:   input,
		Output:  output,
		InSize:  16,
		OutSize: 4096,
	})
	if !success {
		return []*Breakpoint{}
	}
	var breakpoints []*Breakpoint
	index := 0
	if index+4 <= len(output) {
		count := binary.LittleEndian.Uint32(output[index : index+4])
		index += 4
		for i := uint32(0); i < count && index+48 <= len(output); i++ {
			bp := &Breakpoint{
				ID:         binary.LittleEndian.Uint32(output[index+0 : index+0+4]),
				Address:    binary.LittleEndian.Uint64(output[index+4 : index+4+8]),
				ProcessId:  binary.LittleEndian.Uint32(output[index+12 : index+12+4]),
				ThreadId:   binary.LittleEndian.Uint32(output[index+16 : index+16+4]),
				CoreNumber: binary.LittleEndian.Uint32(output[index+20 : index+20+4]),
				IsEnabled:  binary.LittleEndian.Uint32(output[index+24:index+24+4]) != 0,
				Is64Bit:    binary.LittleEndian.Uint32(output[index+28:index+28+4]) != 0,
			}
			copy(bp.Instruction[:], output[index+32:index+32+16])
			breakpoints = append(breakpoints, bp)
			index += 48
		}
	}
	return breakpoints
}

func (dc *DebuggerCore) KillProcess(pid uint32) error {
	input := sdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS{
		ProcessId: pid,
		Action:    sdk.DebuggerAttachDetachUserModeProcessActionKillProcess,
	}
	output := sdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerAttachDetachUserModeProcess,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("kill process %d failed: %v", pid, err)
	}
	return nil
}

func (dc *DebuggerCore) KillActiveProcess() error {
	return dc.KillProcess(0)
}

func (dc *DebuggerCore) RestartProcess(pid uint32) error {
	input := sdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS{
		ProcessId: pid,
		Action:    sdk.DebuggerAttachDetachUserModeProcessActionRemoveHooks,
	}
	output := sdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerAttachDetachUserModeProcess,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("restart process %d failed: %v", pid, err)
	}
	input.Action = sdk.DebuggerAttachDetachUserModeProcessActionAttach
	success, _, err = dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerAttachDetachUserModeProcess,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("attach process %d failed: %v", pid, err)
	}
	return nil
}

func (dc *DebuggerCore) RestartLastProcess() error {
	return dc.RestartProcess(0)
}

func (dc *DebuggerCore) StartProcess(path, args string) error {
	input := make([]byte, 512)
	copy(input[0:], path)
	if len(args) > 0 {
		copy(input[256:], args)
	}
	output := make([]byte, 8)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerAttachDetachUserModeProcess,
		Input:   input,
		Output:  output,
		InSize:  512,
		OutSize: 8,
	})
	if !success || err != nil {
		return fmt.Errorf("start process %s failed: %v", path, err)
	}
	return nil
}

func (dc *DebuggerCore) SwitchCore(coreNum uint32) error {
	input := sdk.DEBUGGEE_CHANGE_CORE_PACKET{
		NewCore: coreNum,
	}
	output := sdk.DEBUGGEE_CHANGE_CORE_PACKET{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlSendUserDebuggerCommands,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("switch core %d failed: %v", coreNum, err)
	}
	return nil
}

func (dc *DebuggerCore) ShowRegisters() (map[string]uint64, error) {
	result := make(map[string]uint64)
	input := make([]byte, 24)
	binary.LittleEndian.PutUint32(input[0:4], 0xffffffff)
	output := make([]byte, 1024)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlSendUserDebuggerCommands,
		Input:   input,
		Output:  output,
		InSize:  24,
		OutSize: 1024,
	})
	if !success || err != nil {
		return result, fmt.Errorf("failed to read registers: %v", err)
	}
	kernelStatus := binary.LittleEndian.Uint32(output[0:4])
	if kernelStatus != uint32(sdk.DebuggerOperationWasSuccessful) {
		return result, fmt.Errorf("register read failed with kernel status 0x%x", kernelStatus)
	}
	regsOffset := 24
	if regsOffset+168 <= len(output) {
		result["rax"] = binary.LittleEndian.Uint64(output[regsOffset+0 : regsOffset+8])
		result["rbx"] = binary.LittleEndian.Uint64(output[regsOffset+8 : regsOffset+16])
		result["rcx"] = binary.LittleEndian.Uint64(output[regsOffset+16 : regsOffset+24])
		result["rdx"] = binary.LittleEndian.Uint64(output[regsOffset+24 : regsOffset+32])
		result["rsi"] = binary.LittleEndian.Uint64(output[regsOffset+32 : regsOffset+40])
		result["rdi"] = binary.LittleEndian.Uint64(output[regsOffset+40 : regsOffset+48])
		result["rbp"] = binary.LittleEndian.Uint64(output[regsOffset+48 : regsOffset+56])
		result["rsp"] = binary.LittleEndian.Uint64(output[regsOffset+56 : regsOffset+64])
		result["r8"] = binary.LittleEndian.Uint64(output[regsOffset+64 : regsOffset+72])
		result["r9"] = binary.LittleEndian.Uint64(output[regsOffset+72 : regsOffset+80])
		result["r10"] = binary.LittleEndian.Uint64(output[regsOffset+80 : regsOffset+88])
		result["r11"] = binary.LittleEndian.Uint64(output[regsOffset+88 : regsOffset+96])
		result["r12"] = binary.LittleEndian.Uint64(output[regsOffset+96 : regsOffset+104])
		result["r13"] = binary.LittleEndian.Uint64(output[regsOffset+104 : regsOffset+112])
		result["r14"] = binary.LittleEndian.Uint64(output[regsOffset+112 : regsOffset+120])
		result["r15"] = binary.LittleEndian.Uint64(output[regsOffset+120 : regsOffset+128])
		extraOffset := regsOffset + 128
		if extraOffset+40 <= len(output) {
			result["rip"] = binary.LittleEndian.Uint64(output[extraOffset+0 : extraOffset+8])
			result["rflags"] = binary.LittleEndian.Uint64(output[extraOffset+8 : extraOffset+16])
		}
	}
	return result, nil
}

func (dc *DebuggerCore) RegisterEpthook(address uint64, pid, coreId, tag uint32, eventType sdk.VMM_EVENT_TYPE_ENUM) (uint32, error) {
	input := make([]byte, 32)
	binary.LittleEndian.PutUint64(input[0:8], address)
	binary.LittleEndian.PutUint32(input[8:12], pid)
	binary.LittleEndian.PutUint32(input[12:16], coreId)
	binary.LittleEndian.PutUint32(input[16:20], tag)
	binary.LittleEndian.PutUint32(input[20:24], uint32(eventType))
	output := make([]byte, 8)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerRegisterEvent,
		Input:   input,
		Output:  output,
		InSize:  32,
		OutSize: 8,
	})
	if !success || err != nil {
		return 0, fmt.Errorf("epthook failed: %v", err)
	}
	return binary.LittleEndian.Uint32(output[0:4]), nil
}

func (dc *DebuggerCore) RegisterException(exceptionIndex uint32, options uint32) error {
	input := make([]byte, 8)
	binary.LittleEndian.PutUint32(input[0:4], exceptionIndex)
	binary.LittleEndian.PutUint32(input[4:8], options)
	output := make([]byte, 8)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerRegisterEvent,
		Input:   input,
		Output:  output,
		InSize:  8,
		OutSize: 8,
	})
	if !success || err != nil {
		return fmt.Errorf("exception hook failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) RegisterMonitor(address uint64, size, monitorType uint32) error {
	input := make([]byte, 20)
	binary.LittleEndian.PutUint64(input[0:8], address)
	binary.LittleEndian.PutUint32(input[8:12], size)
	binary.LittleEndian.PutUint32(input[12:16], monitorType)
	output := make([]byte, 8)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerRegisterEvent,
		Input:   input,
		Output:  output,
		InSize:  20,
		OutSize: 8,
	})
	if !success || err != nil {
		return fmt.Errorf("monitor failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) RegisterMode(modeType sdk.DEBUGGER_EVENT_MODE_TYPE, pid, coreId uint32) error {
	input := make([]byte, 24)
	binary.LittleEndian.PutUint32(input[0:4], uint32(sdk.TrapExecutionModeChanged))
	binary.LittleEndian.PutUint32(input[4:8], uint32(modeType))
	binary.LittleEndian.PutUint32(input[8:12], pid)
	binary.LittleEndian.PutUint32(input[12:16], coreId)
	output := make([]byte, 8)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerRegisterEvent,
		Input:   input,
		Output:  output,
		InSize:  24,
		OutSize: 8,
	})
	if !success || err != nil {
		return fmt.Errorf("mode trap registration failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) RegisterSyscallSysret(syscallIndex, options uint32) error {
	input := make([]byte, 8)
	binary.LittleEndian.PutUint32(input[0:4], syscallIndex)
	binary.LittleEndian.PutUint32(input[4:8], options)
	output := make([]byte, 8)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerRegisterEvent,
		Input:   input,
		Output:  output,
		InSize:  8,
		OutSize: 8,
	})
	if !success || err != nil {
		return fmt.Errorf("syscall hook failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) RegisterMsrReadEvent(msrIndex, options uint32) error {
	input := make([]byte, 8)
	binary.LittleEndian.PutUint32(input[0:4], msrIndex)
	binary.LittleEndian.PutUint32(input[4:8], options)
	output := make([]byte, 8)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerRegisterEvent,
		Input:   input,
		Output:  output,
		InSize:  8,
		OutSize: 8,
	})
	if !success || err != nil {
		return fmt.Errorf("msr read event failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) RegisterMsrWriteEvent(msrIndex, options uint32) error {
	input := make([]byte, 8)
	binary.LittleEndian.PutUint32(input[0:4], msrIndex)
	binary.LittleEndian.PutUint32(input[4:8], options)
	output := make([]byte, 8)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerRegisterEvent,
		Input:   input,
		Output:  output,
		InSize:  8,
		OutSize: 8,
	})
	if !success || err != nil {
		return fmt.Errorf("msr write event failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) RegisterTscEvent(action, options uint32) error {
	input := make([]byte, 8)
	binary.LittleEndian.PutUint32(input[0:4], action)
	binary.LittleEndian.PutUint32(input[4:8], options)
	output := make([]byte, 8)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerRegisterEvent,
		Input:   input,
		Output:  output,
		InSize:  8,
		OutSize: 8,
	})
	if !success || err != nil {
		return fmt.Errorf("tsc event failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) VmcallExtended(rax, rcx, rdx uint64, options uint32) error {
	input := make([]byte, 32)
	binary.LittleEndian.PutUint64(input[0:8], rax)
	binary.LittleEndian.PutUint64(input[8:16], rcx)
	binary.LittleEndian.PutUint64(input[16:24], rdx)
	binary.LittleEndian.PutUint32(input[24:28], options)
	output := make([]byte, 8)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformKernelSideTests,
		Input:   input,
		Output:  output,
		InSize:  32,
		OutSize: 8,
	})
	if !success || err != nil {
		return fmt.Errorf("vmcall extended failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) MonitorStop(monitorId uint32) error {
	input := sdk.DEBUGGER_MODIFY_EVENTS{
		Tag:          uint64(monitorId),
		TypeOfAction: sdk.DebuggerModifyEventsClear,
	}
	output := sdk.DEBUGGER_MODIFY_EVENTS{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerModifyEvents,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("monitor stop failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) SyscallHook(syscallIndex, options uint32) error {
	input := make([]byte, 8)
	binary.LittleEndian.PutUint32(input[0:4], syscallIndex)
	binary.LittleEndian.PutUint32(input[4:8], options)
	output := make([]byte, 8)
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerRegisterEvent,
		Input:   input,
		Output:  output,
		InSize:  8,
		OutSize: 8,
	})
	if !success || err != nil {
		return fmt.Errorf("syscall hook failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) SyscallUnhook(syscallIndex uint32) error {
	input := sdk.DEBUGGER_MODIFY_EVENTS{
		Tag:          uint64(syscallIndex),
		TypeOfAction: sdk.DebuggerModifyEventsClear,
	}
	output := sdk.DEBUGGER_MODIFY_EVENTS{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlDebuggerModifyEvents,
		Input:   structToBytes(&input),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(input)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("syscall unhook failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) GetLocalApic() (*sdk.LAPIC_PAGE, error) {
	apicBytes := make([]byte, 1024)
	err := dc.PerformActionsOnApic(true, apicBytes)
	if err != nil {
		return nil, err
	}
	apic := &sdk.LAPIC_PAGE{}
	if len(apicBytes) >= int(unsafe.Sizeof(*apic)) {
		buf := bytes.NewReader(apicBytes)
		binary.Read(buf, binary.LittleEndian, apic)
	}
	return apic, nil
}

func (dc *DebuggerCore) GetIdtCount() (uint32, error) {
	idtBytes := make([]byte, 8192)
	err := dc.QueryIdtEntry(idtBytes)
	if err != nil {
		return 0, err
	}
	if len(idtBytes) >= 8 {
		return binary.LittleEndian.Uint32(idtBytes[0:8]), nil
	}
	return 0, nil
}

func (dc *DebuggerCore) GetPteDetailsParsed(virtualAddress uint64) (*sdk.DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS, error) {
	data, err := dc.PteDetails(virtualAddress, 0)
	if err != nil {
		return nil, err
	}
	result := &sdk.DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS{}
	if len(data) >= int(unsafe.Sizeof(*result)) {
		buf := bytes.NewReader(data)
		binary.Read(buf, binary.LittleEndian, result)
	}
	return result, nil
}

func (dc *DebuggerCore) GetPciEndpoints() ([]sdk.PCI_DEV_MINIMAL, error) {
	data, err := dc.PciEndpointEnum()
	if err != nil {
		return nil, err
	}
	var devices []sdk.PCI_DEV_MINIMAL
	entrySize := int(unsafe.Sizeof(sdk.PCI_DEV_MINIMAL{}))
	for i := 0; i+entrySize <= len(data); i += entrySize {
		var entry sdk.PCI_DEV_MINIMAL
		buf := bytes.NewReader(data[i : i+entrySize])
		binary.Read(buf, binary.LittleEndian, &entry)
		devices = append(devices, entry)
	}
	return devices, nil
}

func (dc *DebuggerCore) SendStepPacket(stepType sdk.DEBUGGER_REMOTE_STEPPING_REQUEST) error {
	stepPacket := sdk.DEBUGGEE_STEP_PACKET{
		StepType: stepType,
	}
	output := sdk.DEBUGGEE_STEP_PACKET{}
	success, _, err := dc.SendIoctlRequest(DebuggerIoctlRequest{
		Code:    sdk.IoctlSendUserDebuggerCommands,
		Input:   structToBytes(&stepPacket),
		Output:  structToBytes(&output),
		InSize:  uint32(unsafe.Sizeof(stepPacket)),
		OutSize: uint32(unsafe.Sizeof(output)),
	})
	if !success || err != nil {
		return fmt.Errorf("send step packet failed: %v", err)
	}
	return nil
}

func (dc *DebuggerCore) StepInstrumentationIn(count uint32) error {
	for range count {
		err := dc.SendStepPacket(sdk.DebuggerRemoteSteppingRequestInstrumentationStepIn)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dc *DebuggerCore) StepRegularIn(count uint32) error {
	for range count {
		err := dc.SendStepPacket(sdk.DebuggerRemoteSteppingRequestStepIn)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dc *DebuggerCore) StepOverCmd(count uint32) error {
	for range count {
		err := dc.SendStepPacket(sdk.DebuggerRemoteSteppingRequestStepOver)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dc *DebuggerCore) StepOut(lastInstruction bool) error {
	if lastInstruction {
		return dc.SendStepPacket(sdk.DebuggerRemoteSteppingRequestStepOverForGuLastInstruction)
	}
	return dc.SendStepPacket(sdk.DebuggerRemoteSteppingRequestStepOverForGu)
}
