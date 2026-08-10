// Package hwdbg implements the hardware-debugger packet interpreter. The C++
// counterpart is libhyperdbg/code/hwdbg/hwdbg-interpreter.cpp; it owns:
//   - g_HwdbgInstanceInfo / g_HwdbgInstanceInfoIsValid — the cached instance
//     info received from the debuggee on connect
//   - g_HwdbgPortConfiguration — the per-port configuration words
//   - HwdbgInterpretPacket(buf, len) — parse an incoming hwdbg packet
//   - HwdbgParseStringMemoryLine / HwdbgInterpreterFillMemoryFromFile /
//     HwdbgInterpreterFillFileFromMemory — memory-dump file I/O
//   - HwdbgComputeNumberOfFlipFlopsNeeded — FPGA resource estimator
//   - HwdbgInterpreterSendPacketAndBufferToHwdbg — build+send a packet
//   - HwdbgShowIntanceInfo — display the instance info
//   - HwdbgReadInstanceInfoFromFile / HwdbgLoadInstanceInfo — load instance
//     info from a memory-dump file
//
// In the Go rewrite the global state from the C side is owned by the
// Interpreter struct so that multiple debugger instances can coexist
// (GUI/MCP requirement, see API design spec). All mutable fields are
// guarded by a mutex.
//
// Lifecycle:
//
//	intp := hwdbg.NewInterpreter(out)
//	_ = intp.LoadInstanceInfo(ctx, "instance.bin", 4096)
//	intp.ShowInstanceInfo()
//	_ = intp.SendPacket(ctx, instanceInfo, "out.bin",
//	    types.DebuggerRemotePacketTypeDebuggerToDebuggeeHardwareLevel,
//	    types.HwdbgActionSendInstanceInfo, nil)
package hwdbg

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"unsafe"

	"github.com/hyperdbg/go-libhyperdbg/common"
	"github.com/hyperdbg/go-libhyperdbg/types"
)

// Output abstracts message output so CLI/GUI/MCP can all consume the same
// Interpreter instance. It mirrors commands.Output / app.Output but is
// declared locally to keep the hwdbg package free of import cycles.
type Output interface {
	Printf(format string, args ...any) error
}

// IndicatorOfHyperdbgPacket mirrors INDICATOR_OF_HYPERDBG_PACKET ("HYPERDBG"
// as a little-endian uint64 = 0x4859504552444247).
const IndicatorOfHyperdbgPacket uint64 = 0x4859504552444247

// DefaultInitialDebuggeeToDebuggerOffset mirrors
// DEFAULT_INITIAL_DEBUGGEE_TO_DEBUGGER_OFFSET. The C++ header defines it as
// a small constant used before the instance info is received. We use the
// same value the driver expects (0x100 = 256 bytes).
const DefaultInitialDebuggeeToDebuggerOffset = 0x100

// Interpreter owns the hwdbg packet interpreter state. All fields are
// guarded by mu. The zero value is not usable; use NewInterpreter.
type Interpreter struct {
	mu sync.Mutex

	out Output

	// instanceInfo mirrors g_HwdbgInstanceInfo. Valid only when
	// instanceInfoValid is true.
	instanceInfo      types.HWDBG_INSTANCE_INFORMATION
	instanceInfoValid bool

	// portConfiguration mirrors g_HwdbgPortConfiguration. Indexed by port
	// number; populated from the instance-info packet.
	portConfiguration []uint32
}

// NewInterpreter constructs an Interpreter writing diagnostics to out. Pass
// nil to silence diagnostics.
func NewInterpreter(out Output) *Interpreter {
	if out == nil {
		out = discardOutput{}
	}
	return &Interpreter{out: out}
}

// InstanceInfo returns a copy of the cached instance info and whether it is
// valid. Safe to call before LoadInstanceInfo; returns zero+false in that
// case.
func (i *Interpreter) InstanceInfo() (types.HWDBG_INSTANCE_INFORMATION, bool) {
	i.mu.Lock()
	defer i.mu.Unlock()
	return i.instanceInfo, i.instanceInfoValid
}

// PortConfiguration returns a copy of the cached port configuration slice.
func (i *Interpreter) PortConfiguration() []uint32 {
	i.mu.Lock()
	defer i.mu.Unlock()
	out := make([]uint32, len(i.portConfiguration))
	copy(out, i.portConfiguration)
	return out
}

// SetInstanceInfo installs a known instance info directly, bypassing the
// packet path. Intended for tests and for the script-engine integration
// that receives the instance info out-of-band.
func (i *Interpreter) SetInstanceInfo(info types.HWDBG_INSTANCE_INFORMATION, ports []uint32) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.instanceInfo = info
	i.instanceInfoValid = true
	i.portConfiguration = append([]uint32(nil), ports...)
}

// InterpretPacket mirrors HwdbgInterpretPacket. It parses an incoming hwdbg
// packet from buf, dispatches on the requested action, and returns true if
// the packet was handled.
//
// The packet layout is:
//
//	[ optional debuggee-area offset ]
//	DEBUGGER_REMOTE_PACKET (Indicator, TypeOfThePacket, RequestedActionOfThePacket, Checksum)
//	[ optional payload ]
//
// The debuggee-area offset is g_HwdbgInstanceInfo.DebuggeeAreaOffset once
// the instance info is known, otherwise DefaultInitialDebuggeeToDebuggerOffset.
func (i *Interpreter) InterpretPacket(buf []byte) bool {
	if len(buf) < int(unsafe.Sizeof(types.DEBUGGER_REMOTE_PACKET{})) {
		i.out.Printf("err, hwdbg packet too short (%d bytes)\n", len(buf))
		return false
	}
	i.mu.Lock()
	valid := i.instanceInfoValid
	debuggeeOffset := DefaultInitialDebuggeeToDebuggerOffset
	if valid {
		debuggeeOffset = int(i.instanceInfo.DebuggeeAreaOffset)
	}
	i.mu.Unlock()

	if debuggeeOffset > len(buf) {
		i.out.Printf("err, hwdbg packet offset %d beyond buffer length %d\n", debuggeeOffset, len(buf))
		return false
	}
	packetBuf := buf[debuggeeOffset:]
	if len(packetBuf) < int(unsafe.Sizeof(types.DEBUGGER_REMOTE_PACKET{})) {
		i.out.Printf("err, hwdbg packet truncated after offset\n")
		return false
	}
	pkt := (*types.DEBUGGER_REMOTE_PACKET)(unsafe.Pointer(&packetBuf[0]))
	if pkt.Indicator != IndicatorOfHyperdbgPacket {
		// Not a HyperDbg packet; silently ignore (mirrors the C++ fallthrough).
		return false
	}
	if pkt.TypeOfThePacket != types.DebuggerRemotePacketTypeDebuggeeToDebuggerHardwareLevel {
		i.out.Printf("err, unknown packet received from the debuggee\n")
		return false
	}

	switch types.HWDBG_RESPONSE_ENUMS(pkt.RequestedActionOfThePacket) {
	case types.HwdbgResponseSuccessOrErrorMessage:
		// TODO: implement success/error message handling.
		return true

	case types.HwdbgResponseInstanceInfo:
		// Payload follows the packet header: HWDBG_INSTANCE_INFORMATION
		// then numberOfPorts uint32 values.
		infoOffset := int(unsafe.Sizeof(types.DEBUGGER_REMOTE_PACKET{}))
		if len(packetBuf)-infoOffset < int(unsafe.Sizeof(types.HWDBG_INSTANCE_INFORMATION{})) {
			i.out.Printf("err, hwdbg instance info payload too short\n")
			return false
		}
		info := (*types.HWDBG_INSTANCE_INFORMATION)(unsafe.Pointer(&packetBuf[infoOffset]))
		portsOffset := infoOffset + int(unsafe.Sizeof(types.HWDBG_INSTANCE_INFORMATION{}))
		portsNeeded := int(info.NumberOfPorts)
		if len(packetBuf)-portsOffset < portsNeeded*4 {
			i.out.Printf("err, hwdbg port configuration payload too short\n")
			return false
		}
		ports := make([]uint32, portsNeeded)
		for j := 0; j < portsNeeded; j++ {
			ports[j] = *(*uint32)(unsafe.Pointer(&packetBuf[portsOffset+j*4]))
		}
		i.mu.Lock()
		i.instanceInfo = *info
		i.instanceInfoValid = true
		i.portConfiguration = ports
		i.mu.Unlock()
		return true

	default:
		i.out.Printf("err, unknown packet request received from the debuggee\n")
		return false
	}
}

// ParseStringMemoryLine mirrors HwdbgParseStringMemoryLine. It parses a
// single line of a memory-dump file (format: "address: hexword hexword ...")
// and returns the list of 32-bit words.
func ParseStringMemoryLine(line string) []uint32 {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil
	}
	colon := strings.Index(line, ":")
	if colon < 0 {
		return nil
	}
	rest := strings.Fields(line[colon+1:])
	out := make([]uint32, 0, len(rest))
	for _, tok := range rest {
		if len(tok) != 8 {
			continue
		}
		v, err := strconv.ParseUint(tok, 16, 32)
		if err != nil {
			continue
		}
		out = append(out, uint32(v))
	}
	return out
}

// FillMemoryFromFile mirrors HwdbgInterpreterFillMemoryFromFile. It reads
// fileName line by line, parses each line with ParseStringMemoryLine, and
// writes the resulting uint32 values into memoryBuffer. Returns an error if
// the file cannot be opened or if the buffer overflows.
func (i *Interpreter) FillMemoryFromFile(fileName string, memoryBuffer []uint32) error {
	f, err := os.Open(fileName)
	if err != nil {
		i.out.Printf("err, unable to open file %s\n", fileName)
		return fmt.Errorf("hwdbg.FillMemoryFromFile: %w", err)
	}
	defer f.Close()

	idx := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		values := ParseStringMemoryLine(scanner.Text())
		for _, v := range values {
			if idx >= len(memoryBuffer) {
				i.out.Printf("err, buffer overflow, file contains more data than buffer can hold\n")
				return fmt.Errorf("hwdbg.FillMemoryFromFile: buffer overflow")
			}
			memoryBuffer[idx] = v
			idx++
		}
	}
	return scanner.Err()
}

// FillFileFromMemory mirrors HwdbgInterpreterFillFileFromMemory. It writes
// memoryBuffer to fileName in the memory-dump format (one 32-bit word per
// line, with a comment indicating the address and the field meaning for the
// header words). requestedAction is the action code embedded in the comment
// of the 6th word.
func (i *Interpreter) FillFileFromMemory(instanceInfo *types.HWDBG_INSTANCE_INFORMATION, fileName string, memoryBuffer []uint32, requestedAction types.HWDBG_ACTION_ENUMS) error {
	f, err := os.Create(fileName)
	if err != nil {
		i.out.Printf("err, unable to open file %s\n", fileName)
		return fmt.Errorf("hwdbg.FillFileFromMemory: %w", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	address := uint64(0)
	for idx, v := range memoryBuffer {
		_, _ = fmt.Fprintf(w, "%08x ; +0x%x", v, address)
		switch idx {
		case 0, 1:
			_, _ = fmt.Fprintf(w, "   | Checksum")
		case 2, 3:
			_, _ = fmt.Fprintf(w, "   | Indicator")
		case 4:
			_, _ = fmt.Fprintf(w, "  | TypeOfThePacket - DEBUGGER_TO_DEBUGGEE_HARDWARE_LEVEL (0x4)")
		case 5:
			_, _ = fmt.Fprintf(w, "  | RequestedActionOfThePacket - Value (0x%x)", uint32(requestedAction))
		case 6:
			_, _ = fmt.Fprintf(w, "  | Start of Optional Data")
		}
		_, _ = fmt.Fprintln(w)
		address += 4
	}

	// Pad with zeros up to sharedMemorySize so the file matches the shared
	// memory layout expected by the hardware debugger.
	if instanceInfo != nil {
		for address < uint64(instanceInfo.SharedMemorySize) {
			_, _ = fmt.Fprintf(w, "%08x ; +0x%x", 0, address)
			address += 4
			if address < uint64(instanceInfo.SharedMemorySize) {
				_, _ = fmt.Fprintln(w)
			}
		}
	}
	return nil
}

// ComputeNumberOfFlipFlopsNeeded mirrors
// HwdbgComputeNumberOfFlipFlopsNeeded. It estimates the number of flip-
// flops the target FPGA needs to hold the script state, based on the
// instance info and the number of script stages.
func (i *Interpreter) ComputeNumberOfFlipFlopsNeeded(instanceInfo *types.HWDBG_INSTANCE_INFORMATION, numberOfStages uint32) uint64 {
	if instanceInfo == nil {
		return 0
	}
	const sizeofShortSymbol = int(unsafe.Sizeof(types.HWDBG_SHORT_SYMBOL{}))
	const sizeofUint64 = int(unsafe.Sizeof(uint64(0)))
	shortSymbolPerUint64 := sizeofShortSymbol / sizeofUint64

	var total uint64
	// operator (GET + SET)
	total += uint64(numberOfStages) *
		uint64(instanceInfo.MaximumNumberOfSupportedGetScriptOperators+instanceInfo.MaximumNumberOfSupportedSetScriptOperators) *
		uint64(instanceInfo.ScriptVariableLength) *
		uint64(shortSymbolPerUint64)
	// main operator (/2 because Type is not inferred)
	total += uint64(numberOfStages) * uint64(instanceInfo.ScriptVariableLength) * uint64(shortSymbolPerUint64) / 2
	// local + global variables
	total += uint64(numberOfStages) * uint64(instanceInfo.NumberOfSupportedLocalAndGlobalVariables) * uint64(instanceInfo.ScriptVariableLength)
	// temporary variables
	total += uint64(numberOfStages) * uint64(instanceInfo.NumberOfSupportedTemporaryVariables) * uint64(instanceInfo.ScriptVariableLength)
	// stage index register + targetStage (*2)
	stagesShift := common.Log2Ceil(instanceInfo.MaximumNumberOfStages *
		(instanceInfo.MaximumNumberOfSupportedGetScriptOperators +
			instanceInfo.MaximumNumberOfSupportedSetScriptOperators + 1))
	total += uint64(numberOfStages) * uint64(stagesShift) * 2
	// stage enable flip-flop
	total += uint64(numberOfStages)
	// input => output flip-flop
	total += uint64(numberOfStages) * uint64(instanceInfo.NumberOfPins)
	return total
}

// SendPacketAndBufferToHwdbg mirrors
// HwdbgInterpreterSendPacketAndBufferToHwdbg. It builds a
// DEBUGGER_REMOTE_PACKET with the given type and requested action, appends
// the optional payload buffer, applies the debuggee-area offset, and writes
// the result to fileName in the memory-dump format.
//
// Returns an error if the final buffer exceeds the maximum command size
// allowed by the instance info.
func (i *Interpreter) SendPacketAndBufferToHwdbg(instanceInfo *types.HWDBG_INSTANCE_INFORMATION, fileName string, packetType types.DEBUGGER_REMOTE_PACKET_TYPE, requestedAction types.HWDBG_ACTION_ENUMS, payload []byte) error {
	i.mu.Lock()
	valid := i.instanceInfoValid
	i.mu.Unlock()

	var commandMaxSize uint64
	var offset uint64
	if valid && instanceInfo != nil {
		commandMaxSize = uint64(instanceInfo.DebuggeeAreaOffset) - uint64(instanceInfo.DebuggerAreaOffset)
		offset = uint64(instanceInfo.DebuggerAreaOffset)
	} else {
		commandMaxSize = DefaultInitialDebuggeeToDebuggerOffset - DefaultInitialDebuggeeToDebuggerOffset
		offset = DefaultInitialDebuggeeToDebuggerOffset
	}

	payloadLen := uint64(len(payload))
	finalBufferSize := uint64(unsafe.Sizeof(types.DEBUGGER_REMOTE_PACKET{})) + payloadLen
	if finalBufferSize > commandMaxSize {
		i.out.Printf("err, buffer is above the maximum buffer size that can be sent to hwdbg (%d > %d)\n", finalBufferSize, commandMaxSize)
		return fmt.Errorf("hwdbg.SendPacketAndBufferToHwdbg: payload exceeds maximum size")
	}
	finalBufferSize += offset

	finalBuf := make([]byte, finalBufferSize)
	pkt := (*types.DEBUGGER_REMOTE_PACKET)(unsafe.Pointer(&finalBuf[offset]))
	pkt.Indicator = IndicatorOfHyperdbgPacket
	pkt.TypeOfThePacket = packetType
	pkt.RequestedActionOfThePacket = types.DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION(requestedAction)
	// Checksum is computed over the packet body (excluding the checksum byte
	// itself) plus the payload. We use a simple sum mod 256, matching the
	// C++ KdComputeDataChecksum.
	pkt.Checksum = computeChecksum(finalBuf[offset+1 : offset+uint64(unsafe.Sizeof(types.DEBUGGER_REMOTE_PACKET{}))])
	pkt.Checksum += computeChecksum(payload)
	if payloadLen > 0 {
		copy(finalBuf[offset+uint64(unsafe.Sizeof(types.DEBUGGER_REMOTE_PACKET{})):], payload)
	}

	// Reinterpret finalBuf as []uint32 for FillFileFromMemory. The buffer
	// size must be a multiple of 4; pad up if necessary.
	if len(finalBuf)%4 != 0 {
		pad := make([]byte, 4-(len(finalBuf)%4))
		finalBuf = append(finalBuf, pad...)
	}
	memBuf := unsafe.Slice((*uint32)(unsafe.Pointer(&finalBuf[0])), len(finalBuf)/4)

	var infoPtr *types.HWDBG_INSTANCE_INFORMATION
	if instanceInfo != nil {
		infoPtr = instanceInfo
	}
	return i.FillFileFromMemory(infoPtr, fileName, memBuf, requestedAction)
}

// ShowInstanceInfo mirrors HwdbgShowIntanceInfo. It prints the cached
// instance info and port configuration to out.
func (i *Interpreter) ShowInstanceInfo() {
	i.mu.Lock()
	info := i.instanceInfo
	valid := i.instanceInfoValid
	ports := append([]uint32(nil), i.portConfiguration...)
	i.mu.Unlock()
	if !valid {
		i.out.Printf("err, instance info is not valid\n")
		return
	}
	i.out.Printf("Debuggee Version: 0x%x\n", info.Version)
	i.out.Printf("Debuggee Maximum Number Of Stages: 0x%x\n", info.MaximumNumberOfStages)
	i.out.Printf("Debuggee Script Variable Length: 0x%x\n", info.ScriptVariableLength)
	i.out.Printf("Debuggee Number of Supported Local (and global) Variables: 0x%x\n", info.NumberOfSupportedLocalAndGlobalVariables)
	i.out.Printf("Debuggee Number of Supported Temporary Variables: 0x%x\n", info.NumberOfSupportedTemporaryVariables)
	i.out.Printf("Debuggee Maximum Number Of Supported GET Script Operators: 0x%x\n", info.MaximumNumberOfSupportedGetScriptOperators)
	i.out.Printf("Debuggee Maximum Number Of Supported SET Script Operators: 0x%x\n", info.MaximumNumberOfSupportedSetScriptOperators)
	i.out.Printf("Debuggee Shared Memory Size: 0x%x\n", info.SharedMemorySize)
	i.out.Printf("Debuggee Debugger Area Offset: 0x%x\n", info.DebuggerAreaOffset)
	i.out.Printf("Debuggee Debuggee Area Offset: 0x%x\n", info.DebuggeeAreaOffset)
	i.out.Printf("Debuggee Script Capabilities Mask: 0x%llx\n", info.ScriptCapabilities)
	i.out.Printf("Debuggee Number Of Pins: 0x%x\n", info.NumberOfPins)
	i.out.Printf("Debuggee Number Of Ports: 0x%x\n", info.NumberOfPorts)
	i.out.Printf("Debuggee BRAM Address Width: 0x%x\n", info.BramAddrWidth)
	i.out.Printf("Debuggee BRAM Data Width: 0x%x (%d bit)\n", info.BramDataWidth, info.BramDataWidth)
	for portNum, cfg := range ports {
		i.out.Printf("Port number %d ($hw_port%d): 0x%x\n", portNum, portNum, cfg)
	}
}

// ReadInstanceInfoFromFile mirrors HwdbgReadInstanceInfoFromFile. It reads
// the instance-info packet from fileName into memoryBuffer and prints the
// raw words for verification. Returns nil on success.
func (i *Interpreter) ReadInstanceInfoFromFile(fileName string, memoryBuffer []uint32) error {
	if err := i.FillMemoryFromFile(fileName, memoryBuffer); err != nil {
		return err
	}
	for _, w := range memoryBuffer {
		i.out.Printf("%08x \n", w)
	}
	return nil
}

// LoadInstanceInfo mirrors HwdbgLoadInstanceInfo. It reads the instance-info
// packet from instanceFilePathToRead, interprets it, and (on success) prints
// the instance info via ShowInstanceInfo.
//
// initialBramBufferSize is the number of uint32 words to allocate for the
// read buffer; it must be large enough to hold the entire instance-info
// packet (header + instance info + port configuration).
func (i *Interpreter) LoadInstanceInfo(ctx context.Context, instanceFilePathToRead string, initialBramBufferSize uint32) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	memBuf := make([]uint32, initialBramBufferSize)
	if err := i.ReadInstanceInfoFromFile(instanceFilePathToRead, memBuf); err != nil {
		i.out.Printf("err, unable to read instance info packet of the debuggee")
		return err
	}
	if !i.InterpretPacket(uint32SliceToBytes(memBuf)) {
		i.out.Printf("err, unable to interpret instance info packet of the debuggee")
		return fmt.Errorf("hwdbg.LoadInstanceInfo: interpret failed")
	}
	i.out.Printf("instance info interpreted successfully\n")
	i.ShowInstanceInfo()
	return nil
}

// computeChecksum mirrors KdComputeDataChecksum: sum of all bytes mod 256.
func computeChecksum(buf []byte) uint8 {
	var sum uint32
	for _, b := range buf {
		sum += uint32(b)
	}
	return uint8(sum & 0xFF)
}

// uint32SliceToBytes reinterprets a []uint32 as a []byte without copying.
// The returned slice aliases the same memory.
func uint32SliceToBytes(s []uint32) []byte {
	if len(s) == 0 {
		return nil
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(&s[0])), len(s)*4)
}

// discardOutput is the default Output when the caller passes nil.
type discardOutput struct{}

func (discardOutput) Printf(format string, args ...any) error {
	_ = args
	_ = format
	return nil
}
