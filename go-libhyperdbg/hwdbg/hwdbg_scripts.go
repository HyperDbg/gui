// Package hwdbg also implements the hardware-script runner. The C++
// counterpart is libhyperdbg/code/hwdbg/hwdbg-scripts.cpp; it owns:
//   - HwdbgScriptPrintScriptBuffer(buf, len) — dump the raw script bytes
//   - HwdbgScriptCompressScriptBuffer(InstanceInfo, ScriptBuffer, ...) —
//     convert SYMBOL buffer to HWDBG_SHORT_SYMBOL and compress it to the
//     target BRAM data width
//   - HwdbgScriptPrintFinalScriptBufferAndHardwareDetails(...) — dump the
//     compressed script + FPGA resource usage
//   - HwdbgScriptWriteScriptConfigurationPacketIntoFile(...) — wrap the
//     compressed script in an HWDBG_SCRIPT_BUFFER header and write it via
//     HwdbgInterpreterSendPacketAndBufferToHwdbg
//   - HwdbgScriptCreateHwdbgScript(buf, len, path) — top-level orchestrator
//     that calls the script-engine checker, the compressor, the printer and
//     the file writer
//   - HwdbgScriptGetScriptBufferFromRawString(script, ...) — run the script
//     engine parser to obtain a SYMBOL buffer
//   - HwdbgScriptSendScriptPacket(...) — assemble header+payload and forward
//     to HwdbgInterpreterSendPacketAndBufferToHwdbg
//   - HwdbgScriptRunScript(script, instancePath, savePath, bramSize) — top-
//     level entry that loads instance info, parses the script, prints it and
//     creates the hwdbg script file
//
// The C++ code relies on the script-engine wrapper (ScriptEngineParseWrapper,
// HardwareScriptInterpreterConvertSymbolToHwdbgShortSymbolBuffer,
// HardwareScriptInterpreterCompressBuffer,
// HardwareScriptInterpreterCheckScriptBufferWithScriptCapabilities,
// HardwareScriptInterpreterFreeHwdbgShortSymbolBuffer) which is not yet
// ported to Go. The methods here expose the API shape and delegate to a
// ScriptEngine interface so that the eventual Go script engine can be plugged
// in without changing call sites. Until that port lands, the runner returns
// ErrScriptEngineNotAvailable from the affected entry points.
package hwdbg

import (
	"context"
	"fmt"
	"sync"
	"unsafe"

	"github.com/ddkwork/hyperdbgsdk"
)

// ErrScriptEngineNotAvailable is returned by the script-runner entry points
// when no ScriptEngine implementation has been wired up. The C++ code paths
// this guards (ScriptEngineParseWrapper,
// HardwareScriptInterpreterConvertSymbolToHwdbgShortSymbolBuffer, ...)
// are not yet ported to Go.
var ErrScriptEngineNotAvailable = fmt.Errorf("hwdbg: script engine not yet available (Phase C.3)")

// ScriptEngine abstracts the script-engine operations needed by the hwdbg
// script runner. The C++ counterpart lives in libhyperdbg's script-engine
// wrapper layer; the Go port is tracked under Phase C.3.
//
// Implementations must be safe for concurrent use; the runner serialises
// calls itself but the engine may be shared across runners.
type ScriptEngine interface {
	// ParseScript mirrors ScriptEngineParseWrapper. It parses scriptText
	// and returns the resulting SYMBOL buffer (the code buffer) plus the
	// head address, byte length and pointer the C++ API exposes.
	ParseScript(scriptText string) (codeBuffer []byte, headAddress uint64, bufferLength uint32, pointer uint32, err error)

	// CheckScriptCapabilities mirrors
	// HardwareScriptInterpreterCheckScriptBufferWithScriptCapabilities.
	// It validates that the script (represented as a SYMBOL buffer) is
	// supported by the instance info and returns the number of stages, the
	// number of operands the script needs and the number of operands the
	// instance implements.
	CheckScriptCapabilities(instanceInfo *hyperdbgsdk.HWDBG_INSTANCE_INFORMATION, scriptBuffer []byte) (numberOfStages, numberOfOperandsForScript, numberOfOperandsImplemented uint32, err error)

	// ConvertToShortSymbols mirrors
	// HardwareScriptInterpreterConvertSymbolToHwdbgShortSymbolBuffer.
	// It converts the SYMBOL buffer to a HWDBG_SHORT_SYMBOL buffer sized
	// for the instance info.
	ConvertToShortSymbols(instanceInfo *hyperdbgsdk.HWDBG_INSTANCE_INFORMATION, scriptBuffer []byte, numberOfStagesForScript uint32) (newScriptBuffer []hyperdbgsdk.HWDBG_SHORT_SYMBOL, err error)

	// CompressBuffer mirrors HardwareScriptInterpreterCompressBuffer. It
	// compresses newScriptBuffer in place to the target BRAM data width
	// and returns the compressed size (in bytes) and the number of bytes
	// per chunk.
	CompressBuffer(newScriptBuffer []hyperdbgsdk.HWDBG_SHORT_SYMBOL, scriptVariableLength, bramDataWidth uint32) (compressedSize, numberOfBytesPerChunk uint32, err error)

	// FreeShortSymbolBuffer mirrors
	// HardwareScriptInterpreterFreeHwdbgShortSymbolBuffer. Go is garbage
	// collected so this is a no-op for the default engine, but the method
	// is kept on the interface so that engines backed by C allocators can
	// release memory deterministically.
	FreeShortSymbolBuffer(newScriptBuffer []hyperdbgsdk.HWDBG_SHORT_SYMBOL)
}

// ScriptRunner owns the hwdbg script-runner state. It holds a reference to
// the *Interpreter (for instance info, packet send and FPGA resource
// estimation) and an optional ScriptEngine. All mutable fields are guarded
// by mu. The zero value is not usable; use NewScriptRunner.
type ScriptRunner struct {
	mu sync.Mutex

	out    Output
	intp   *Interpreter
	engine ScriptEngine
}

// NewScriptRunner constructs a ScriptRunner bound to the given *Interpreter.
// engine may be nil, in which case methods that need the script engine
// return ErrScriptEngineNotAvailable.
func NewScriptRunner(intp *Interpreter, engine ScriptEngine, out Output) *ScriptRunner {
	if out == nil {
		out = discardOutput{}
	}
	return &ScriptRunner{out: out, intp: intp, engine: engine}
}

// SetScriptEngine replaces the script engine. Pass nil to disable the
// script-dependent entry points. Safe to call concurrently with method
// invocations; in-flight calls use the previous engine.
func (r *ScriptRunner) SetScriptEngine(engine ScriptEngine) {
	r.mu.Lock()
	r.engine = engine
	r.mu.Unlock()
}

// PrintScriptBuffer mirrors HwdbgScriptPrintScriptBuffer. It dumps the raw
// script bytes to out, prefixed with the size and the number of flip-flops
// (script bits) the buffer represents.
func (r *ScriptRunner) PrintScriptBuffer(scriptBuffer []byte) {
	r.out.Printf("\nHyperDbg (general) script buffer (size=%d, flip-flops (just script)=%d):\n\n",
		len(scriptBuffer), len(scriptBuffer)*8)
	for _, b := range scriptBuffer {
		r.out.Printf("%02X ", b)
	}
	r.out.Printf("\n")
}

// CompressScriptBuffer mirrors HwdbgScriptCompressScriptBuffer. It converts
// scriptBuffer (a SYMBOL buffer) to HWDBG_SHORT_SYMBOL and compresses it to
// the BRAM data width. Returns the new short-symbol buffer (aliased by the
// returned slice), the compressed size in bytes and the number of bytes per
// chunk.
//
// Returns ErrScriptEngineNotAvailable when no ScriptEngine is wired up.
func (r *ScriptRunner) CompressScriptBuffer(instanceInfo *hyperdbgsdk.HWDBG_INSTANCE_INFORMATION, scriptBuffer []byte, numberOfStagesForScript uint32) (newScriptBuffer []hyperdbgsdk.HWDBG_SHORT_SYMBOL, compressedSize, numberOfBytesPerChunk uint32, err error) {
	engine, intp, ok := r.snapshot()
	if !ok {
		return nil, 0, 0, ErrScriptEngineNotAvailable
	}
	_, valid := intp.InstanceInfo()
	if !valid {
		r.out.Printf("err, the instance info is not valid\n")
		return nil, 0, 0, fmt.Errorf("hwdbg.CompressScriptBuffer: instance info is not valid")
	}
	if instanceInfo == nil {
		return nil, 0, 0, fmt.Errorf("hwdbg.CompressScriptBuffer: nil instance info")
	}
	// Mirror the C++ guard: scriptVariableLength must be at least 8 bits.
	if instanceInfo.ScriptVariableLength < 8 {
		r.out.Printf("err, the script variable length should be at least 8 bits (1 byte)\n")
		return nil, 0, 0, fmt.Errorf("hwdbg.CompressScriptBuffer: script variable length < 8 bits")
	}

	short, err := engine.ConvertToShortSymbols(instanceInfo, scriptBuffer, numberOfStagesForScript)
	if err != nil {
		r.out.Printf("err, unable to convert the script buffer to short symbol buffer\n")
		return nil, 0, 0, fmt.Errorf("hwdbg.CompressScriptBuffer: %w", err)
	}
	compressed, perChunk, err := engine.CompressBuffer(short, instanceInfo.ScriptVariableLength, instanceInfo.BramDataWidth)
	if err != nil {
		engine.FreeShortSymbolBuffer(short)
		return nil, 0, 0, fmt.Errorf("hwdbg.CompressScriptBuffer: %w", err)
	}
	return short, compressed, perChunk, nil
}

// PrintFinalScriptBufferAndHardwareDetails mirrors
// HwdbgScriptPrintFinalScriptBufferAndHardwareDetails. It dumps the
// compressed script buffer along with the FPGA resource usage computed by
// Interpreter.ComputeNumberOfFlipFlopsNeeded.
func (r *ScriptRunner) PrintFinalScriptBufferAndHardwareDetails(instanceInfo *hyperdbgsdk.HWDBG_INSTANCE_INFORMATION, compressedSize, numberOfStagesForScript, numberOfOperandsForScript, numberOfOperandsImplemented, numberOfBytesPerChunk uint32, newScriptBuffer []hyperdbgsdk.HWDBG_SHORT_SYMBOL) {
	r.out.Printf("\n---------------------------------------------------------\n")

	numberOfNeededFlipFlops := r.intp.ComputeNumberOfFlipFlopsNeeded(instanceInfo, numberOfStagesForScript)

	pct := float64(0)
	if numberOfOperandsImplemented != 0 {
		pct = (float64(numberOfOperandsForScript) / float64(numberOfOperandsImplemented)) * 100
	}
	r.out.Printf("hwdbg script buffer (buffer size=%d, stages=%d, operands needed: %d - operands used: %d (%.2f%%), total used flip-flops=%d, number of bytes per chunk: %d):\n\n",
		compressedSize, numberOfStagesForScript,
		numberOfOperandsImplemented, numberOfOperandsForScript, pct,
		numberOfNeededFlipFlops, numberOfBytesPerChunk)

	// Reinterpret the short-symbol buffer as raw bytes for the hex dump.
	if len(newScriptBuffer) > 0 {
		raw := unsafe.Slice((*byte)(unsafe.Pointer(&newScriptBuffer[0])), compressedSize)
		for _, b := range raw {
			r.out.Printf("%02X ", b)
		}
	}
}

// WriteScriptConfigurationPacketIntoFile mirrors
// HwdbgScriptWriteScriptConfigurationPacketIntoFile. It assembles the
// HWDBG_SCRIPT_BUFFER header followed by the compressed script buffer and
// forwards the result to Interpreter.SendPacketAndBufferToHwdbg with the
// hwdbgActionConfigureScriptBuffer action.
//
// numberOfSymbols should be (numberOfStages + numberOfOperands - 1) per the
// C++ call site; the caller is responsible for the arithmetic.
func (r *ScriptRunner) WriteScriptConfigurationPacketIntoFile(instanceInfo *hyperdbgsdk.HWDBG_INSTANCE_INFORMATION, fileName string, numberOfSymbols uint32, newScriptBuffer []hyperdbgsdk.HWDBG_SHORT_SYMBOL, compressedSize uint32) error {
	r.out.Printf("\n\nwriting script configuration packet into the file\n")

	// Build [HWDBG_SCRIPT_BUFFER header][payload bytes].
	var header hyperdbgsdk.HWDBG_SCRIPT_BUFFER
	header.ScriptNumberOfSymbols = numberOfSymbols

	headerSize := uint32(unsafe.Sizeof(header))
	totalSize := headerSize + compressedSize
	finalBuf := make([]byte, totalSize)
	copy(finalBuf[:headerSize], asHeaderBytes(&header))
	if compressedSize > 0 && len(newScriptBuffer) > 0 {
		payload := unsafe.Slice((*byte)(unsafe.Pointer(&newScriptBuffer[0])), compressedSize)
		copy(finalBuf[headerSize:], payload)
	}

	if err := r.intp.SendPacketAndBufferToHwdbg(instanceInfo, fileName,
		hyperdbgsdk.DebuggerRemotePacketTypeDebuggerToDebuggeeHardwareLevel,
		hyperdbgsdk.HwdbgActionConfigureScriptBuffer, finalBuf); err != nil {
		r.out.Printf("err, unable to write script buffer\n")
		return err
	}
	r.out.Printf("\n[*] script buffer successfully written into file: %s\n", fileName)
	return nil
}

// CreateHwdbgScript mirrors HwdbgScriptCreateHwdbgScript. It validates the
// script against the instance info's capabilities, compresses it, prints the
// final buffer and hardware details, and writes the configuration packet to
// hardwareScriptFilePathToSave.
//
// Returns ErrScriptEngineNotAvailable when no ScriptEngine is wired up.
func (r *ScriptRunner) CreateHwdbgScript(scriptBuffer []byte, hardwareScriptFilePathToSave string) error {
	engine, intp, ok := r.snapshot()
	if !ok {
		return ErrScriptEngineNotAvailable
	}
	instanceInfo, valid := intp.InstanceInfo()
	if !valid {
		r.out.Printf("err, the instance info is not valid\n")
		return fmt.Errorf("hwdbg.CreateHwdbgScript: instance info is not valid")
	}

	numberOfStagesForScript, numberOfOperandsForScript, numberOfOperandsImplemented, err := engine.CheckScriptCapabilities(&instanceInfo, scriptBuffer)
	if err != nil {
		r.out.Printf("\n[-] target script is NOT supported by this instance of hwdbg!\n")
		return fmt.Errorf("hwdbg.CreateHwdbgScript: %w", err)
	}
	r.out.Printf("\n[+] target script is supported by this instance of hwdbg!\n")

	newScriptBuffer, compressedSize, numberOfBytesPerChunk, err := r.CompressScriptBuffer(&instanceInfo, scriptBuffer, numberOfStagesForScript)
	if err != nil {
		r.out.Printf("err, unable to compress the script buffer\n")
		return err
	}
	defer engine.FreeShortSymbolBuffer(newScriptBuffer)

	r.PrintFinalScriptBufferAndHardwareDetails(&instanceInfo, compressedSize, numberOfStagesForScript, numberOfOperandsForScript, numberOfOperandsImplemented, numberOfBytesPerChunk, newScriptBuffer)

	// Number of symbols = Number of stages + Number of operands - 1
	numberOfSymbols := numberOfStagesForScript + numberOfOperandsImplemented - 1
	return r.WriteScriptConfigurationPacketIntoFile(&instanceInfo, hardwareScriptFilePathToSave, numberOfSymbols, newScriptBuffer, compressedSize)
}

// GetScriptBufferFromRawString mirrors HwdbgScriptGetScriptBufferFromRawString.
// It runs the script engine parser on scriptText and returns the resulting
// code buffer plus the head address, byte length and pointer.
//
// Returns ErrScriptEngineNotAvailable when no ScriptEngine is wired up.
func (r *ScriptRunner) GetScriptBufferFromRawString(scriptText string) (codeBuffer []byte, headAddress uint64, bufferLength, pointer uint32, err error) {
	engine, _, ok := r.snapshot()
	if !ok {
		return nil, 0, 0, 0, ErrScriptEngineNotAvailable
	}
	return engine.ParseScript(scriptText)
}

// SendScriptPacket mirrors HwdbgScriptSendScriptPacket. It assembles the
// HWDBG_SCRIPT_BUFFER header followed by the short-symbol buffer and
// forwards the result to Interpreter.SendPacketAndBufferToHwdbg with the
// hwdbgActionConfigureScriptBuffer action.
func (r *ScriptRunner) SendScriptPacket(instanceInfo *hyperdbgsdk.HWDBG_INSTANCE_INFORMATION, fileName string, numberOfSymbols uint32, buffer []hyperdbgsdk.HWDBG_SHORT_SYMBOL, bufferLength uint32) error {
	var header hyperdbgsdk.HWDBG_SCRIPT_BUFFER
	header.ScriptNumberOfSymbols = numberOfSymbols

	headerSize := uint32(unsafe.Sizeof(header))
	totalSize := headerSize + bufferLength
	finalBuf := make([]byte, totalSize)
	copy(finalBuf[:headerSize], asHeaderBytes(&header))
	if bufferLength > 0 && len(buffer) > 0 {
		payload := unsafe.Slice((*byte)(unsafe.Pointer(&buffer[0])), bufferLength)
		copy(finalBuf[headerSize:], payload)
	}

	return r.intp.SendPacketAndBufferToHwdbg(instanceInfo, fileName,
		hyperdbgsdk.DebuggerRemotePacketTypeDebuggerToDebuggeeHardwareLevel,
		hyperdbgsdk.HwdbgActionConfigureScriptBuffer, finalBuf)
}

// RunScript mirrors HwdbgScriptRunScript. It is the top-level entry that
// loads the instance info from instanceFilePathToRead, parses the script,
// prints the script buffer, and creates the hwdbg script file at
// hardwareScriptFilePathToSave.
//
// initialBramBufferSize is the number of uint32 words to allocate for the
// instance-info read buffer (mirrors the C++ parameter).
//
// Returns ErrScriptEngineNotAvailable when no ScriptEngine is wired up.
func (r *ScriptRunner) RunScript(ctx context.Context, script, instanceFilePathToRead, hardwareScriptFilePathToSave string, initialBramBufferSize uint32) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if r.intp == nil {
		return fmt.Errorf("hwdbg.RunScript: no *Interpreter bound")
	}
	if err := r.intp.LoadInstanceInfo(ctx, instanceFilePathToRead, initialBramBufferSize); err != nil {
		return err
	}

	codeBuffer, _, bufferLength, _, err := r.GetScriptBufferFromRawString(script)
	if err != nil {
		return fmt.Errorf("hwdbg.RunScript: %w", err)
	}
	r.PrintScriptBuffer(codeBuffer[:bufferLength])

	if err := r.CreateHwdbgScript(codeBuffer[:bufferLength], hardwareScriptFilePathToSave); err != nil {
		r.out.Printf("err, unable to create hwdbg script\n")
		return err
	}
	return nil
}

// snapshot returns the engine and interpreter under the lock. ok is false
// when no script engine is wired up.
func (r *ScriptRunner) snapshot() (ScriptEngine, *Interpreter, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.engine == nil {
		return nil, nil, false
	}
	return r.engine, r.intp, true
}

// asHeaderBytes returns a byte slice aliasing the memory of header for the
// duration of the call. Used to serialise the HWDBG_SCRIPT_BUFFER header
// without encoding/binary overhead.
func asHeaderBytes(header *hyperdbgsdk.HWDBG_SCRIPT_BUFFER) []byte {
	const sz = int(unsafe.Sizeof(*header))
	return unsafe.Slice((*byte)(unsafe.Pointer(header)), sz)
}
