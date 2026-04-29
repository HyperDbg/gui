package user_level

import (
	"fmt"

	"github.com/ddkwork/HyperDbg/debugger/misc"
	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/sdk"
	"golang.org/x/sys/windows"
)

func HandleUserDebuggerPausing(
	pausePacket *sdk.DEBUGGEE_UD_PAUSED_PACKET,
	activeProcess *ActiveDebuggingProcess,
	syncObjects *[MaximumSyncObjects]SyncronizationEventState,
) {
	switch pausePacket.PausingReason {
	case sdk.DebuggeePausingReasonDebuggeeStartingModuleLoaded:
		mylog.Info("the target module is loaded and a breakpoint is set to the entrypoint")
		mylog.Info("press 'g' to reach to the entrypoint of the main module...")
	case sdk.DebuggeePausingReasonDebuggeeGeneralThreadIntercepted:
	default:
		break
	}

	if pausePacket.ReadInstructionLen != MaximumInstrSize {
		length, _ := misc.GetInstructionLength(pausePacket.InstructionBytesOnRip[:], func() misc.DisasmMode {
			if pausePacket.Is32Bit {
				return misc.Mode32
			}
			return misc.Mode64
		}())
		if length > uint32(pausePacket.ReadInstructionLen) {
			mylog.Warning("there might be a misinterpretation in disassembling the current instruction")
		}
	}

	if !activeProcess.IsPaused || pausePacket.ThreadId == activeProcess.ThreadId {
		activeProcess.IsActive = true
		activeProcess.ProcessId = pausePacket.ProcessId
		activeProcess.ThreadId = pausePacket.ThreadId
		activeProcess.Is32Bit = pausePacket.Is32Bit
		activeProcess.Rip = pausePacket.Rip
		activeProcess.ProcessDebuggingToken = pausePacket.ProcessDebuggingToken
		copy(activeProcess.InstructionBytesOnRip[:], pausePacket.InstructionBytesOnRip[:])
		activeProcess.IsPaused = true

		if !pausePacket.Is32Bit {
			fmt.Println()
			Disassemble64(pausePacket.InstructionBytesOnRip[:], pausePacket.Rip, MaximumInstrSize, 1, true, &pausePacket.Rflags)
		} else {
			fmt.Println()
			Disassemble32(pausePacket.InstructionBytesOnRip[:], pausePacket.Rip, MaximumInstrSize, 1, true, &pausePacket.Rflags)
		}
	}

	if syncObjects[SyncObjectUserDebuggerIsDebuggerRunning].IsOnWaitingState {
		windows.SetEvent(syncObjects[SyncObjectUserDebuggerIsDebuggerRunning].EventHandle)
	}
}

func Disassemble64(instructionBytes []byte, rip uint64, maxInstrSize, count uint32, showRflags bool, rflags *uint64) {
	for i := uint32(0); i < count && i*maxInstrSize < uint32(len(instructionBytes)); i++ {
		offset := i * maxInstrSize
		end := min(int(offset)+int(maxInstrSize), len(instructionBytes))
		inst, err := misc.DisassembleOne(instructionBytes[offset:end], rip+uint64(offset), misc.Mode64)
		if err != nil {
			continue
		}
		fmt.Printf("%s: %s\n", fmt.Sprintf("%x", inst.Address), misc.FormatDisassembly(inst))
	}
}

func Disassemble32(instructionBytes []byte, rip uint64, maxInstrSize, count uint32, showRflags bool, rflags *uint64) {
	for i := uint32(0); i < count && i*maxInstrSize < uint32(len(instructionBytes)); i++ {
		offset := i * maxInstrSize
		end := min(int(offset)+int(maxInstrSize), len(instructionBytes))
		inst, err := misc.DisassembleOne(instructionBytes[offset:end], rip+uint64(offset), misc.Mode32)
		if err != nil {
			continue
		}
		fmt.Printf("%s: %s\n", fmt.Sprintf("%x", inst.Address), misc.FormatDisassembly(inst))
	}
}
