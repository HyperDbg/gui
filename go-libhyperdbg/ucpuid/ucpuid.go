// Package ucpuid implements the `ucpuid` command — read and decode USER-
// specified CPUID information. The C++ counterpart is
// libhyperdbg/ucpuid.cpp; it owns:
//   - CommandUserCpuidHelp() — print the command help
//   - CommandShowUserCpuidMessage(FunctionId, SubFunctionId, CpuidRequest) —
//     decode and print the CPUID result for the given leaf/sub-leaf
//   - CommandCpuidRequestCpuid(FunctionId, SubFunctionId) — send
//     IOCTL_DEBUGGER_CPUID to the kernel and dispatch the result to the
//     decoder
//   - CommandUserCpuid(CommandTokens, Command) — parse the command line and
//     dispatch to CommandCpuidRequestCpuid
//
// In the Go rewrite the global state from the C side (g_IsKdModuleLoaded,
// g_DeviceHandle, g_IsSerialConnectedToRemoteDebuggee) is owned by the
// Handler struct, which holds a reference to the *core.Debugger (for the
// device handle) and an Output sink. The serial-debuggee path is not yet
// wired up; only the local IOCTL path is implemented.
//
// Lifecycle:
//
//	h := ucpuid.New(coreDebugger, out)
//	_ = h.RequestCpuid(ctx, 0x1, 0x0)          // raw IOCTL
//	_ = h.ParseAndExecute(ctx, []string{"ucpuid", "1"}) // full command
package ucpuid

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"unsafe"

	"github.com/ddkwork/hyperdbgsdk"
	"github.com/hyperdbg/go-libhyperdbg/common"
	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
	"github.com/hyperdbg/go-libhyperdbg/debugger/core"
)

// Output abstracts message output so CLI/GUI/MCP can all consume the same
// Handler instance. It mirrors commands.Output / app.Output but is declared
// locally to keep the ucpuid package free of import cycles.
type Output interface {
	Printf(format string, args ...any) error
}

// DebuggerOperationWasSuccessful mirrors DEBUGGER_OPERATION_WAS_SUCCESSFUL.
const DebuggerOperationWasSuccessful uint32 = 0xFFFFFFFF

// Handler owns the ucpuid command state. It is bound to one *core.Debugger
// (for the device handle) and one Output sink.
//
// The zero value is not usable; use New.
type Handler struct {
	mu  sync.Mutex
	dev *core.Debugger
	out Output

	// isSerialConnectedToRemoteDebuggee mirrors the C++ global of the same
	// name. When true, RequestCpuid would forward the request to the remote
	// debuggee over the serial KD channel instead of issuing a local IOCTL.
	// The serial path is not yet ported (Phase C.3); the flag is kept on
	// the struct so the API matches the C++ behaviour.
	isSerialConnectedToRemoteDebuggee bool
}

// New constructs a Handler instance bound to the given *core.Debugger. The
// debugger must already be connected (call core.Debugger.Connect or
// app.App.LoadVMM before invoking any Handler method).
func New(d *core.Debugger, out Output) *Handler {
	if out == nil {
		out = discardOutput{}
	}
	return &Handler{dev: d, out: out}
}

// SetSerialConnected mirrors setting g_IsSerialConnectedToRemoteDebuggee.
// When true, RequestCpuid will (once the serial KD path is ported) forward
// CPUID requests to the remote debuggee instead of issuing a local IOCTL.
func (h *Handler) SetSerialConnected(connected bool) {
	h.mu.Lock()
	h.isSerialConnectedToRemoteDebuggee = connected
	h.mu.Unlock()
}

// Help mirrors CommandUserCpuidHelp. It prints the command syntax and a few
// example invocations to out.
func (h *Handler) Help() {
	h.out.Printf("ucpuid : reads CPUID information on the target debuggee.\n\n")
	h.out.Printf("syntax : \tucpuid [Function (hex)] [SubFunction (hex)]\n\n")
	h.out.Printf("\n")
	h.out.Printf("\t\te.g : ucpuid 1\n")
	h.out.Printf("\t\te.g : ucpuid 4 2\n")
	h.out.Printf("\t\te.g : ucpuid D 0\n")
	h.out.Printf("\t\te.g : ucpuid 0x80000008 0\n\n")
	h.out.Printf("\n")
	h.out.Printf("note 1: use `ucpuid 0` to see the maximum supported CPUID leaf\n")
	h.out.Printf("note 2: use `ucpuid 0x80000000` to see the maximum supported extended leaf\n")
}

// ParseAndExecute mirrors CommandUserCpuid. It parses the command tokens
// (the first token is the command name "ucpuid" and is consumed), validates
// the argument count, extracts the (optional) function id and sub-function
// id, and dispatches to RequestCpuid.
//
// tokens is the full token list including the command name. Returns an error
// if parsing fails or if the IOCTL fails.
func (h *Handler) ParseAndExecute(ctx context.Context, tokens []string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if len(tokens) == 0 {
		return fmt.Errorf("ucpuid: empty token list")
	}
	// Drop the command name (tokens[0]).
	args := tokens[1:]
	if len(args) > 2 {
		h.out.Printf("incorrect use of the '%s'\n\n", tokens[0])
		h.Help()
		return fmt.Errorf("ucpuid: too many arguments (%d)", len(args))
	}

	var (
		functionId    uint32
		subFunctionId uint32
		setFunctionId bool
	)
	for _, section := range args {
		if !setFunctionId {
			v, ok := common.ConvertStringToUInt32(section)
			if !ok {
				h.out.Printf("please specify a correct hex value for function id\n\n")
				h.Help()
				return fmt.Errorf("ucpuid: invalid function id %q", section)
			}
			functionId = v
			setFunctionId = true
			continue
		}
		v, ok := common.ConvertStringToUInt32(section)
		if !ok {
			h.out.Printf("please specify a correct hex value for sub-function id\n\n")
			h.Help()
			return fmt.Errorf("ucpuid: invalid sub-function id %q", section)
		}
		subFunctionId = v
	}
	if !setFunctionId {
		h.out.Printf("please specify a cpuid function id\n\n")
		h.Help()
		return fmt.Errorf("ucpuid: missing function id")
	}
	return h.RequestCpuid(ctx, functionId, subFunctionId)
}

// RequestCpuid mirrors CommandCpuidRequestCpuid. It sends IOCTL_DEBUGGER_CPUID
// to the kernel with the (FunctionId, SubFunctionId) pair and dispatches the
// result to ShowCpuidMessage on success.
//
// The serial-debuggee path (KdSendUserCpuidPacketToDebuggee) is not yet
// ported; if SetSerialConnected(true) was called, RequestCpuid returns an
// error indicating the path is not yet available.
func (h *Handler) RequestCpuid(ctx context.Context, functionId, subFunctionId uint32) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	h.mu.Lock()
	serial := h.isSerialConnectedToRemoteDebuggee
	h.mu.Unlock()
	if serial {
		// TODO(Phase C.3): KdSendUserCpuidPacketToDebuggee(functionId, subFunctionId)
		return fmt.Errorf("ucpuid: serial-debuggee CPUID path not yet implemented (Phase C.3)")
	}

	dev, err := h.device()
	if err != nil {
		return err
	}

	var req hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE
	req.FunctionId = functionId
	req.SubFunctionId = subFunctionId

	buf := asBytes(&req)
	if _, err := dev.Ioctl(ctx, comm.IOCTL_CODE_DEBUGGER_CPUID, buf, buf); err != nil {
		h.out.Printf("ioctl failed with code 0x%x\n", err)
		return fmt.Errorf("ucpuid: IOCTL_DEBUGGER_CPUID failed: %w", err)
	}
	if req.KernelStatus != DebuggerOperationWasSuccessful {
		h.out.Printf("Receiving CPUID result was not successful :(\n")
		return fmt.Errorf("ucpuid: kernel returned status 0x%x", req.KernelStatus)
	}
	h.ShowCpuidMessage(functionId, subFunctionId, &req)
	return nil
}

// ShowCpuidMessage mirrors CommandShowUserCpuidMessage. It dispatches on the
// CPUID leaf (function id) to a leaf-specific decoder. Leaves that HyperDbg
// does not yet decode are printed as raw dwords, matching the C++ default
// case.
func (h *Handler) ShowCpuidMessage(functionId, subFunctionId uint32, req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	if req == nil {
		h.out.Printf("  NULL value!\n")
		return
	}
	switch functionId {
	case 0x0:
		h.showLeaf0(req)
	case 0x1:
		h.showLeaf1(req)
	case 0x2:
		h.showLeafRaw(0x2, "Legacy Cache Descriptor",
			"This leaf is deprecated and returns legacy cache information.\n  Use CPUID.(EAX=04H) for deterministic cache parameters instead.", req)
	case 0x3:
		h.showLeafRaw(0x3, "",
			"This leaf is reserved/not implemented on modern processors.\n  (Was previously used for Processor Serial Number on older CPUs)", req)
	case 0x4:
		h.showLeaf4(subFunctionId, req)
	case 0x5:
		h.showLeaf5(req)
	case 0x6:
		h.showLeaf6(req)
	case 0x7:
		h.showLeaf7(subFunctionId, req)
	case 0x8:
		h.showLeafRaw(0x8, "",
			"This leaf is reserved/not implemented on modern processors.\n  (Was previously used for Processor Serial Number on some CPUs)", req)
	case 0x9:
		h.showLeaf9(req)
	case 0xA:
		h.showLeafA(req)
	case 0xB:
		h.showLeafB(subFunctionId, req)
	case 0xC:
		h.showLeafRaw(0xC, "",
			"This leaf is reserved (not implemented).", req)
	case 0xD:
		h.showLeafD(subFunctionId, req)
	case 0xE:
		h.showLeafRaw(0xE, "",
			"This leaf is reserved (not implemented).", req)
	case 0xF:
		h.showLeafRaw(0xF, "Intel RDT Monitoring",
			"This leaf is not yet implemented in HyperDbg.\n  (Intel Resource Director Technology - Monitoring)", req)
	case 0x10:
		h.showLeaf10(subFunctionId, req)
	case 0x11:
		h.showLeafRaw(0x11, "",
			"This leaf is reserved (not implemented).", req)
	case 0x12:
		h.showLeaf12(subFunctionId, req)
	case 0x13:
		h.showLeafRaw(0x13, "",
			"This leaf is reserved (not implemented).", req)
	case 0x14:
		h.showLeaf14(subFunctionId, req)
	case 0x15:
		h.showLeaf15(req)
	case 0x16:
		h.showLeaf16(req)
	case 0x17:
		h.showLeaf17(subFunctionId, req)
	case 0x18:
		h.showLeaf18(subFunctionId, req)
	case 0x80000000:
		h.showLeaf80000000(req)
	case 0x80000001:
		h.showLeaf80000001(req)
	case 0x80000002, 0x80000003, 0x80000004:
		h.showLeaf80000002to04(req)
	case 0x80000005:
		h.out.Printf("EAX = 0x80000005: not implemented.\n")
	case 0x80000006:
		h.showLeaf80000006(req)
	case 0x80000007:
		h.showLeaf80000007(req)
	case 0x80000008:
		h.showLeaf80000008(req)
	default:
		h.out.Printf("==== CPUID.(EAX=%08XH) ====\n\n", functionId)
		h.out.Printf("  CPUID leaf 0x%08X is not implemented in HyperDbg.\n", functionId)
		h.out.Printf("  You can decode the raw values yourself:\n")
		h.out.Printf("    EAX: 0x%08X\n", req.EAX)
		h.out.Printf("    EBX: 0x%08X\n", req.EBX)
		h.out.Printf("    ECX: 0x%08X\n", req.ECX)
		h.out.Printf("    EDX: 0x%08X\n\n", req.EDX)
	}
}

// device returns the device handle from the bound *core.Debugger. It returns
// an error if the debugger is not connected or does not yet expose its
// device handle.
func (h *Handler) device() (*comm.Device, error) {
	h.mu.Lock()
	d := h.dev
	h.mu.Unlock()
	if d == nil {
		return nil, fmt.Errorf("ucpuid: no *core.Debugger bound")
	}
	// TODO(Phase C.3): core.Debugger.Device() *comm.Device
	return nil, fmt.Errorf("ucpuid: core.Debugger.Device() not yet exposed (Phase C.3)")
}

// asBytes returns a byte slice aliasing the memory of req for the duration
// of the call. Used to (de)serialise the fixed-size C struct without
// encoding/binary overhead.
func asBytes(req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) []byte {
	const sz = int(unsafe.Sizeof(*req))
	return unsafe.Slice((*byte)(unsafe.Pointer(req)), sz)
}

// ---------------------------------------------------------------------------
// Bitfield helpers. These mirror the CPUID_* macros in ia32.h. Each helper
// extracts a contiguous bit range [hi:lo] from a uint32 register value.
// ---------------------------------------------------------------------------

// bits extracts bits [hi:lo] (inclusive) from v, returning them right-
// shifted to bit 0. hi must be >= lo and < 32.
func bits(v uint32, hi, lo uint) uint32 {
	if hi < lo {
		return 0
	}
	mask := uint32(((1 << (hi - lo + 1)) - 1) << lo)
	return (v & mask) >> lo
}

// bit reports whether bit `n` of v is set. Equivalent to bits(v, n, n) == 1.
func bit(v uint32, n uint) bool {
	return v&(1<<n) != 0
}

// btoi converts a bool to the "TRUE"/"FALSE" string used throughout the C++
// decoder.
func btoi(b bool) string {
	if b {
		return "TRUE"
	}
	return "FALSE"
}

// vendorString extracts the 12-byte vendor string from EBX/EDX/ECX of leaf 0.
func vendorString(ebx, edx, ecx uint32) string {
	var b [12]byte
	for i := 0; i < 4; i++ {
		b[i] = byte(ebx >> (i * 8))
		b[4+i] = byte(edx >> (i * 8))
		b[8+i] = byte(ecx >> (i * 8))
	}
	return strings.TrimRight(string(b[:]), "\x00")
}

// brandString returns the NUL-terminated brand string from the BrandString
// field of the request struct (filled in by the kernel for leaves
// 0x80000002-0x80000004).
func brandString(req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) string {
	for i, c := range req.BrandString {
		if c == 0 {
			return string(req.BrandString[:i])
		}
	}
	return string(req.BrandString[:])
}

// noSubleavesBanner prints the "LEAF N HAS NO SUBLEAVES" banner the C++
// decoder emits for leaves that ignore the sub-leaf input.
func (h *Handler) noSubleavesBanner(leafHex string) {
	h.out.Printf("  *******************************************************\n")
	h.out.Printf("  *               LEAF %s HAS NO SUBLEAVES               *\n", leafHex)
	h.out.Printf("  *       ANY SUBLEAF YOU ENTER WILL DEFAULT TO 0       *\n")
	h.out.Printf("  *      AND THE PROCESSOR RETURNS UNDEFINED VALUES     *\n")
	h.out.Printf("  *******************************************************\n\n")
}

// maxSubleafBanner prints the "Max NumberOfSubLeaves = N" banner.
func (h *Handler) maxSubleafBanner(maxSubleaf uint32) {
	h.out.Printf("  *******************************************************\n")
	h.out.Printf("  *             Max NumberOfSubLeaves = %u               *\n", maxSubleaf)
	h.out.Printf("  *******************************************************\n\n")
}

// showLeafRaw prints the standard "raw dwords" block used for reserved /
// not-implemented leaves.
func (h *Handler) showLeafRaw(leaf uint32, title, body string, req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	if title != "" {
		h.out.Printf("==== CPUID.(EAX=%XH) %s ====\n\n", leaf, title)
	} else {
		h.out.Printf("==== CPUID.(EAX=%XH) ====\n\n", leaf)
	}
	if body != "" {
		h.out.Printf("  %s\n", body)
	}
	h.out.Printf("  Raw data:\n")
	h.out.Printf("  EAX: 0x%08X\n", req.EAX)
	h.out.Printf("  EBX: 0x%08X\n", req.EBX)
	h.out.Printf("  ECX: 0x%08X\n", req.ECX)
	h.out.Printf("  EDX: 0x%08X\n\n", req.EDX)
}

// ---------------------------------------------------------------------------
// Leaf-specific decoders. Each method mirrors the corresponding `case 0xN:`
// block in CommandShowUserCpuidMessage.
// ---------------------------------------------------------------------------

// showLeaf0 decodes CPUID.0: vendor string + max basic leaf.
func (h *Handler) showLeaf0(req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.noSubleavesBanner("0")
	h.out.Printf("  Vendor : %s\n", vendorString(req.EBX, req.EDX, req.ECX))
	h.out.Printf("  Maximum supported basic leaf : %u\n\n", req.EAX)
}

// showLeaf1 decodes CPUID.1: version, additional and feature information.
func (h *Handler) showLeaf1(req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=01H) Version / Additional / Feature Information ====\n\n")
	h.noSubleavesBanner("1")

	// EAX: Version Information
	h.out.Printf("-- EAX: Version Information --\n\n")
	h.out.Printf("  SteppingId       = %u\n", bits(req.EAX, 3, 0))
	h.out.Printf("  Model            = %u\n", bits(req.EAX, 7, 4))
	h.out.Printf("  FamilyId         = %u\n", bits(req.EAX, 11, 8))
	h.out.Printf("  ProcessorType    = %u\n", bits(req.EAX, 13, 12))
	h.out.Printf("  ExtendedModelId  = %u\n", bits(req.EAX, 19, 16))
	h.out.Printf("  ExtendedFamilyId = %u\n\n", bits(req.EAX, 27, 20))

	// EBX: Additional Information
	h.out.Printf("-- EBX: Additional Information --\n\n")
	h.out.Printf("  BrandIndex        = %u\n", bits(req.EBX, 7, 0))
	h.out.Printf("  ClflushLineSize   = %u (cache line = %u bytes)\n",
		bits(req.EBX, 15, 8), bits(req.EBX, 15, 8)*8)
	h.out.Printf("  MaxAddressableIds = %u\n", bits(req.EBX, 23, 16))
	h.out.Printf("  InitialApicId     = %u\n\n", bits(req.EBX, 31, 24))

	// ECX: Feature Information
	h.out.Printf("-- ECX: Feature Information --\n\n")
	ecx := req.ECX
	h.out.Printf("  SSE3                  = %s\n", btoi(bit(ecx, 0)))
	h.out.Printf("  PCLMULQDQ             = %s\n", btoi(bit(ecx, 1)))
	h.out.Printf("  DTES64                = %s\n", btoi(bit(ecx, 2)))
	h.out.Printf("  MONITOR/MWAIT         = %s\n", btoi(bit(ecx, 3)))
	h.out.Printf("  CPL Qualified DS      = %s\n", btoi(bit(ecx, 4)))
	h.out.Printf("  VMX                   = %s\n", btoi(bit(ecx, 5)))
	h.out.Printf("  SMX                   = %s\n", btoi(bit(ecx, 6)))
	h.out.Printf("  EIST (SpeedStep)      = %s\n", btoi(bit(ecx, 7)))
	h.out.Printf("  TM2                   = %s\n", btoi(bit(ecx, 8)))
	h.out.Printf("  SSSE3                 = %s\n", btoi(bit(ecx, 9)))
	h.out.Printf("  L1 Context ID         = %s\n", btoi(bit(ecx, 10)))
	h.out.Printf("  Silicon Debug         = %s\n", btoi(bit(ecx, 11)))
	h.out.Printf("  FMA                   = %s\n", btoi(bit(ecx, 12)))
	h.out.Printf("  CMPXCHG16B            = %s\n", btoi(bit(ecx, 13)))
	h.out.Printf("  xTPR Update Control   = %s\n", btoi(bit(ecx, 14)))
	h.out.Printf("  PDCM                  = %s\n", btoi(bit(ecx, 15)))
	h.out.Printf("  PCID                  = %s\n", btoi(bit(ecx, 17)))
	h.out.Printf("  DCA                   = %s\n", btoi(bit(ecx, 18)))
	h.out.Printf("  SSE4.1                = %s\n", btoi(bit(ecx, 19)))
	h.out.Printf("  SSE4.2                = %s\n", btoi(bit(ecx, 20)))
	h.out.Printf("  x2APIC                = %s\n", btoi(bit(ecx, 21)))
	h.out.Printf("  MOVBE                 = %s\n", btoi(bit(ecx, 22)))
	h.out.Printf("  POPCNT                = %s\n", btoi(bit(ecx, 23)))
	h.out.Printf("  TSC-Deadline          = %s\n", btoi(bit(ecx, 24)))
	h.out.Printf("  AESNI                 = %s\n", btoi(bit(ecx, 25)))
	h.out.Printf("  XSAVE/XRSTOR          = %s\n", btoi(bit(ecx, 26)))
	h.out.Printf("  OSXSAVE               = %s\n", btoi(bit(ecx, 27)))
	h.out.Printf("  AVX                   = %s\n", btoi(bit(ecx, 28)))
	h.out.Printf("  F16C                  = %s\n", btoi(bit(ecx, 29)))
	h.out.Printf("  RDRAND                = %s\n\n", btoi(bit(ecx, 30)))

	// EDX: Feature Information
	h.out.Printf("-- EDX: Feature Information --\n\n")
	edx := req.EDX
	h.out.Printf("  FPU                   = %s\n", btoi(bit(edx, 0)))
	h.out.Printf("  VME                   = %s\n", btoi(bit(edx, 1)))
	h.out.Printf("  DE                    = %s\n", btoi(bit(edx, 2)))
	h.out.Printf("  PSE                   = %s\n", btoi(bit(edx, 3)))
	h.out.Printf("  TSC                   = %s\n", btoi(bit(edx, 4)))
	h.out.Printf("  MSR (RDMSR/WRMSR)     = %s\n", btoi(bit(edx, 5)))
	h.out.Printf("  PAE                   = %s\n", btoi(bit(edx, 6)))
	h.out.Printf("  MCE                   = %s\n", btoi(bit(edx, 7)))
	h.out.Printf("  CX8 (CMPXCHG8B)       = %s\n", btoi(bit(edx, 8)))
	h.out.Printf("  APIC On-Chip          = %s\n", btoi(bit(edx, 9)))
	h.out.Printf("  SEP (SYSENTER/EXIT)   = %s\n", btoi(bit(edx, 11)))
	h.out.Printf("  MTRR                  = %s\n", btoi(bit(edx, 12)))
	h.out.Printf("  PGE                   = %s\n", btoi(bit(edx, 13)))
	h.out.Printf("  MCA                   = %s\n", btoi(bit(edx, 14)))
	h.out.Printf("  CMOV                  = %s\n", btoi(bit(edx, 15)))
	h.out.Printf("  PAT                   = %s\n", btoi(bit(edx, 16)))
	h.out.Printf("  PSE-36                = %s\n", btoi(bit(edx, 17)))
	h.out.Printf("  PSN                   = %s\n", btoi(bit(edx, 18)))
	h.out.Printf("  CLFSH                 = %s\n", btoi(bit(edx, 19)))
	h.out.Printf("  DS (Debug Store)      = %s\n", btoi(bit(edx, 21)))
	h.out.Printf("  ACPI (Thermal/Clock)  = %s\n", btoi(bit(edx, 22)))
	h.out.Printf("  MMX                   = %s\n", btoi(bit(edx, 23)))
	h.out.Printf("  FXSR (FXSAVE/FXRSTOR) = %s\n", btoi(bit(edx, 24)))
	h.out.Printf("  SSE                   = %s\n", btoi(bit(edx, 25)))
	h.out.Printf("  SSE2                  = %s\n", btoi(bit(edx, 26)))
	h.out.Printf("  SS (Self Snoop)       = %s\n", btoi(bit(edx, 27)))
	h.out.Printf("  HTT                   = %s\n", btoi(bit(edx, 28)))
	h.out.Printf("  TM (Thermal Monitor)  = %s\n", btoi(bit(edx, 29)))
	h.out.Printf("  PBE                   = %s\n\n", btoi(bit(edx, 31)))
}

// showLeaf4 decodes CPUID.4: deterministic cache parameters.
func (h *Handler) showLeaf4(subFunctionId uint32, req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=04H) Deterministic Cache Parameters ====\n\n")
	h.maxSubleafBanner(req.Leaf4MaxSubLeaf)
	h.out.Printf("---- CPUID.(EAX=04H, ECX=%u) ----\n\n", subFunctionId)

	// EAX
	cacheType := bits(req.EAX, 4, 0)
	typeName := "Reserved"
	switch cacheType {
	case 0:
		typeName = "Null (no more caches)"
	case 1:
		typeName = "Data Cache"
	case 2:
		typeName = "Instruction Cache"
	case 3:
		typeName = "Unified Cache"
	}
	h.out.Printf("-- EAX --\n\n")
	h.out.Printf("  CacheTypeField                   = %u (%s)\n", cacheType, typeName)
	if cacheType == 0 {
		h.out.Printf("  (no more caches; stopping enumeration)\n\n")
		return
	}
	h.out.Printf("  CacheLevel                       = %u\n", bits(req.EAX, 7, 5))
	h.out.Printf("  SelfInitializingCacheLevel       = %s\n", btoi(bit(req.EAX, 8)))
	fa := bit(req.EAX, 9)
	h.out.Printf("  FullyAssociativeCache            = %s%s\n", btoi(fa), func() string {
		if fa {
			return " (fully associative)"
		}
		return ""
	}())
	maxIdsLogical := bits(req.EAX, 25, 14)
	maxIdsCores := bits(req.EAX, 31, 26)
	h.out.Printf("  MaxAddressableIds(LogicalProcs)  (raw) = %u -> actual = %u (raw + 1)\n", maxIdsLogical, maxIdsLogical+1)
	h.out.Printf("  MaxAddressableIds(Cores)         (raw) = %u -> actual = %u (raw + 1)\n\n", maxIdsCores, maxIdsCores+1)

	// EBX
	lineSize := bits(req.EBX, 11, 0)
	partitions := bits(req.EBX, 21, 12)
	ways := bits(req.EBX, 31, 22)
	h.out.Printf("-- EBX --\n\n")
	h.out.Printf("  SystemCoherencyLineSize          (raw) = %u -> actual = %u bytes (raw + 1)\n", lineSize, lineSize+1)
	h.out.Printf("  PhysicalLinePartitions           (raw) = %u -> actual = %u (raw + 1)\n", partitions, partitions+1)
	h.out.Printf("  WaysOfAssociativity              (raw) = %u -> actual = %u (raw + 1)\n\n", ways, ways+1)

	// ECX
	sets := req.ECX
	h.out.Printf("-- ECX --\n\n")
	h.out.Printf("  NumberOfSets                     (raw) = %u -> actual = %u (raw + 1)\n\n", sets, sets+1)

	// EDX
	h.out.Printf("-- EDX --\n\n")
	h.out.Printf("  WriteBackInvalidate              = %s\n", btoi(bit(req.EDX, 0)))
	incl := bit(req.EDX, 1)
	h.out.Printf("  CacheInclusiveness               = %s%s\n", btoi(incl), func() string {
		if incl {
			return " (inclusive of lower levels)"
		}
		return ""
	}())
	complex := bit(req.EDX, 2)
	h.out.Printf("  ComplexCacheIndexing             = %s%s\n\n", btoi(complex), func() string {
		if complex {
			return " (complex/hashed indexing)"
		}
		return " (direct mapped)"
	}())

	// Cache size
	cacheSize := uint64(ways+1) * uint64(partitions+1) * uint64(lineSize+1) * uint64(sets+1)
	h.out.Printf("-- Cache Size --\n\n")
	h.out.Printf("  Cache Size (per spec formula)    = %llu bytes (%llu KB, %llu MB)\n\n",
		cacheSize, cacheSize/1024, cacheSize/(1024*1024))
}

// showLeaf5 decodes CPUID.5: MONITOR/MWAIT.
func (h *Handler) showLeaf5(req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=05H) MONITOR/MWAIT Leaf ====\n\n")
	h.noSubleavesBanner("5")

	h.out.Printf("-- EAX --\n\n")
	h.out.Printf("  SmallestMonitorLineSize = %u bytes\n\n", bits(req.EAX, 15, 0))

	h.out.Printf("-- EBX --\n\n")
	h.out.Printf("  LargestMonitorLineSize  = %u bytes\n\n", bits(req.EBX, 15, 0))

	h.out.Printf("-- ECX --\n\n")
	h.out.Printf("  EnumerationOfMonitorMwaitExtensions             = %s\n", btoi(bit(req.ECX, 0)))
	h.out.Printf("  SupportsTreatingInterruptsAsBreakEventForMwait  = %s\n\n", btoi(bit(req.ECX, 1)))

	h.out.Printf("-- EDX --\n\n")
	for i := uint(0); i <= 7; i++ {
		h.out.Printf("  NumberOfC%dSubCStates = %u\n", i, bits(req.EDX, i*4+3, i*4))
	}
	h.out.Printf("\n")
}

// showLeaf6 decodes CPUID.6: thermal and power management.
func (h *Handler) showLeaf6(req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=06H) Thermal and Power Management Leaf ====\n\n")
	h.noSubleavesBanner("6")

	h.out.Printf("-- EAX --\n\n")
	eax := req.EAX
	h.out.Printf("  TemperatureSensorSupported              = %s\n", btoi(bit(eax, 0)))
	h.out.Printf("  IntelTurboBoostTechnologyAvailable       = %s\n", btoi(bit(eax, 1)))
	h.out.Printf("  ARAT (ApicTimerAlwaysRunning)            = %s\n", btoi(bit(eax, 2)))
	h.out.Printf("  PLN (PowerLimitNotification)             = %s\n", btoi(bit(eax, 3)))
	h.out.Printf("  ECMD (ClockModulationDuty)               = %s\n", btoi(bit(eax, 4)))
	h.out.Printf("  PTM (PackageThermalManagement)           = %s\n", btoi(bit(eax, 6)))
	h.out.Printf("  HWP Base Registers                       = %s\n", btoi(bit(eax, 7)))
	h.out.Printf("  HWP_Notification                         = %s\n", btoi(bit(eax, 8)))
	h.out.Printf("  HWP_Activity_Window                      = %s\n", btoi(bit(eax, 9)))
	h.out.Printf("  HWP_Energy_Performance_Preference        = %s\n", btoi(bit(eax, 10)))
	h.out.Printf("  HWP_Package_Level_Request                = %s\n", btoi(bit(eax, 11)))
	h.out.Printf("  HDC                                      = %s\n", btoi(bit(eax, 13)))
	h.out.Printf("  Intel Turbo Boost Max Technology 3.0     = %s\n", btoi(bit(eax, 14)))
	h.out.Printf("  HWP Capabilities                         = %s\n", btoi(bit(eax, 15)))
	h.out.Printf("  HWP PECI Override                        = %s\n", btoi(bit(eax, 16)))
	h.out.Printf("  Flexible HWP                             = %s\n", btoi(bit(eax, 17)))
	h.out.Printf("  Fast Access Mode for HWP Request MSR     = %s\n", btoi(bit(eax, 18)))
	h.out.Printf("  Ignoring Idle Logical Proc HWP Request   = %s\n", btoi(bit(eax, 19)))
	h.out.Printf("  Intel Thread Director                    = %s\n\n", btoi(bit(eax, 20)))

	h.out.Printf("-- EBX --\n\n")
	h.out.Printf("  NumberOfInterruptThresholdsInThermalSensor = %u\n\n", bits(req.EBX, 3, 0))

	h.out.Printf("-- ECX --\n\n")
	h.out.Printf("  HardwareCoordinationFeedbackCapability   = %s\n", btoi(bit(req.ECX, 0)))
	h.out.Printf("  NumberOfIntelThreadDirectorClasses (bit) = %s\n", btoi(bit(req.ECX, 1)))
	h.out.Printf("  PerformanceEnergyBiasPreference          = %u\n\n", bits(req.ECX, 3, 0))

	h.out.Printf("-- EDX --\n\n")
	h.out.Printf("  Reserved                                 = 0x%08X\n\n", req.EDX)
}

// showLeaf7 decodes CPUID.7: structured extended feature flags.
func (h *Handler) showLeaf7(subFunctionId uint32, req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=07H) Structured Extended Feature Flags ====\n\n")
	h.maxSubleafBanner(req.LeafEaxMaxSubleaf)

	if subFunctionId != 0 {
		// Sub-leaves > 0 have no bitfield macros defined in ia32.h; print raw.
		h.out.Printf("  No bitfield macros are defined for these in ia32.h, so their\n")
		h.out.Printf("  raw dwords are shown without decoding:\n")
		h.out.Printf("    ECX=%u: EAX=0x%08X EBX=0x%08X ECX=0x%08X EDX=0x%08X\n\n",
			subFunctionId, req.EAX, req.EBX, req.ECX, req.EDX)
		return
	}

	h.out.Printf("---- CPUID.(EAX=07H, ECX=%u) ----\n\n", subFunctionId)

	// EAX
	h.out.Printf("-- EAX --\n\n")
	h.out.Printf("  NumberOfSubLeaves (max ECX input) = %u\n\n", req.EAX)

	// EBX
	h.out.Printf("-- EBX --\n\n")
	ebx := req.EBX
	h.out.Printf("  FSGSBASE                       = %s\n", btoi(bit(ebx, 0)))
	h.out.Printf("  IA32_TSC_ADJUST MSR            = %s\n", btoi(bit(ebx, 1)))
	h.out.Printf("  SGX                            = %s\n", btoi(bit(ebx, 2)))
	h.out.Printf("  BMI1                           = %s\n", btoi(bit(ebx, 3)))
	h.out.Printf("  HLE                            = %s\n", btoi(bit(ebx, 4)))
	h.out.Printf("  AVX2                           = %s\n", btoi(bit(ebx, 5)))
	h.out.Printf("  FDP_EXCPTN_ONLY                = %s\n", btoi(bit(ebx, 6)))
	h.out.Printf("  SMEP                           = %s\n", btoi(bit(ebx, 7)))
	h.out.Printf("  BMI2                           = %s\n", btoi(bit(ebx, 8)))
	h.out.Printf("  Enhanced REP MOVSB/STOSB       = %s\n", btoi(bit(ebx, 9)))
	h.out.Printf("  INVPCID                        = %s\n", btoi(bit(ebx, 10)))
	h.out.Printf("  RTM                            = %s\n", btoi(bit(ebx, 11)))
	h.out.Printf("  RDT-M (Monitoring)             = %s\n", btoi(bit(ebx, 12)))
	h.out.Printf("  Deprecates FPU CS/DS           = %s\n", btoi(bit(ebx, 13)))
	h.out.Printf("  MPX                            = %s\n", btoi(bit(ebx, 14)))
	h.out.Printf("  RDT-A (Allocation)             = %s\n", btoi(bit(ebx, 15)))
	h.out.Printf("  AVX512F                        = %s\n", btoi(bit(ebx, 16)))
	h.out.Printf("  AVX512DQ                       = %s\n", btoi(bit(ebx, 17)))
	h.out.Printf("  RDSEED                         = %s\n", btoi(bit(ebx, 18)))
	h.out.Printf("  ADX                            = %s\n", btoi(bit(ebx, 19)))
	h.out.Printf("  SMAP                           = %s\n", btoi(bit(ebx, 20)))
	h.out.Printf("  AVX512_IFMA                    = %s\n", btoi(bit(ebx, 21)))
	h.out.Printf("  CLFLUSHOPT                     = %s\n", btoi(bit(ebx, 23)))
	h.out.Printf("  CLWB                           = %s\n", btoi(bit(ebx, 24)))
	h.out.Printf("  Intel Processor Trace          = %s\n", btoi(bit(ebx, 25)))
	h.out.Printf("  AVX512PF (Xeon Phi only)       = %s\n", btoi(bit(ebx, 26)))
	h.out.Printf("  AVX512ER (Xeon Phi only)       = %s\n", btoi(bit(ebx, 27)))
	h.out.Printf("  AVX512CD                       = %s\n", btoi(bit(ebx, 28)))
	h.out.Printf("  SHA                            = %s\n", btoi(bit(ebx, 29)))
	h.out.Printf("  AVX512BW                       = %s\n", btoi(bit(ebx, 30)))
	h.out.Printf("  AVX512VL                       = %s\n\n", btoi(bit(ebx, 31)))

	// ECX
	h.out.Printf("-- ECX --\n\n")
	ecx := req.ECX
	h.out.Printf("  PREFETCHWT1 (Xeon Phi only)    = %s\n", btoi(bit(ecx, 0)))
	h.out.Printf("  AVX512_VBMI                    = %s\n", btoi(bit(ecx, 1)))
	h.out.Printf("  UMIP                           = %s\n", btoi(bit(ecx, 2)))
	h.out.Printf("  PKU                            = %s\n", btoi(bit(ecx, 3)))
	h.out.Printf("  OSPKE                          = %s\n", btoi(bit(ecx, 4)))
	h.out.Printf("  WAITPKG                        = %s\n", btoi(bit(ecx, 5)))
	h.out.Printf("  AVX512_VBMI2                   = %s\n", btoi(bit(ecx, 6)))
	h.out.Printf("  CET_SS (shadow stack)          = %s\n", btoi(bit(ecx, 7)))
	h.out.Printf("  GFNI                           = %s\n", btoi(bit(ecx, 8)))
	h.out.Printf("  VAES                           = %s\n", btoi(bit(ecx, 9)))
	h.out.Printf("  VPCLMULQDQ                     = %s\n", btoi(bit(ecx, 10)))
	h.out.Printf("  AVX512_VNNI                    = %s\n", btoi(bit(ecx, 11)))
	h.out.Printf("  AVX512_BITALG                  = %s\n", btoi(bit(ecx, 12)))
	h.out.Printf("  TME_EN                         = %s\n", btoi(bit(ecx, 13)))
	h.out.Printf("  AVX512_VPOPCNTDQ               = %s\n", btoi(bit(ecx, 14)))
	h.out.Printf("  LA57 (5-level paging)          = %s\n", btoi(bit(ecx, 16)))
	h.out.Printf("  MAWAU (BNDLDX/BNDSTX)          = %u (NOT BOOLEAN)\n", bits(ecx, 21, 17))
	h.out.Printf("  RDPID                          = %s\n", btoi(bit(ecx, 22)))
	h.out.Printf("  KL (Key Locker)                = %s\n", btoi(bit(ecx, 23)))
	h.out.Printf("  CLDEMOTE                       = %s\n", btoi(bit(ecx, 25)))
	h.out.Printf("  MOVDIRI                        = %s\n", btoi(bit(ecx, 27)))
	h.out.Printf("  MOVDIR64B                      = %s\n", btoi(bit(ecx, 28)))
	h.out.Printf("  SGX_LC (Launch Config)         = %s\n", btoi(bit(ecx, 30)))
	h.out.Printf("  PKS                            = %s\n\n", btoi(bit(ecx, 31)))

	// EDX
	h.out.Printf("-- EDX --\n\n")
	edx := req.EDX
	h.out.Printf("  AVX512_4VNNIW (Xeon Phi only)  = %s\n", btoi(bit(edx, 2)))
	h.out.Printf("  AVX512_4FMAPS (Xeon Phi only)  = %s\n", btoi(bit(edx, 3)))
	h.out.Printf("  Fast Short REP MOV             = %s\n", btoi(bit(edx, 4)))
	h.out.Printf("  AVX512_VP2INTERSECT            = %s\n", btoi(bit(edx, 8)))
	h.out.Printf("  MD_CLEAR                       = %s\n", btoi(bit(edx, 10)))
	h.out.Printf("  SERIALIZE                      = %s\n", btoi(bit(edx, 14)))
	h.out.Printf("  Hybrid part                    = %s\n", btoi(bit(edx, 15)))
	h.out.Printf("  PCONFIG                        = %s\n", btoi(bit(edx, 18)))
	h.out.Printf("  CET_IBT (branch tracking)      = %s\n", btoi(bit(edx, 20)))
	h.out.Printf("  IBRS/IBPB                      = %s\n", btoi(bit(edx, 26)))
	h.out.Printf("  STIBP                          = %s\n", btoi(bit(edx, 27)))
	h.out.Printf("  L1D_FLUSH                      = %s\n", btoi(bit(edx, 28)))
	h.out.Printf("  IA32_ARCH_CAPABILITIES MSR     = %s\n", btoi(bit(edx, 29)))
	h.out.Printf("  IA32_CORE_CAPABILITIES MSR     = %s\n", btoi(bit(edx, 30)))
	h.out.Printf("  SSBD                           = %s\n\n", btoi(bit(edx, 31)))
}

// showLeaf9 decodes CPUID.9: direct cache access information.
func (h *Handler) showLeaf9(req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=09H) Direct Cache Access Information ====\n\n")
	h.noSubleavesBanner("9")
	h.out.Printf("  IA32_PLATFORM_DCA_CAP (EAX, mirrors MSR 1F8H) = 0x%08X\n", req.EAX)
	h.out.Printf("  Reserved (EBX) = 0x%08X\n", req.EBX)
	h.out.Printf("  Reserved (ECX) = 0x%08X\n", req.ECX)
	h.out.Printf("  Reserved (EDX) = 0x%08X\n\n", req.EDX)
}

// showLeafA decodes CPUID.A: architectural performance monitoring.
func (h *Handler) showLeafA(req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=0AH) Architectural Performance Monitoring Leaf ====\n\n")
	h.noSubleavesBanner("A")

	versionId := bits(req.EAX, 7, 0)
	if versionId == 0 {
		h.out.Printf("  VersionId = 0 -> architectural performance monitoring not supported;\n")
		h.out.Printf("  remaining fields in this leaf are not architecturally defined.\n\n")
		return
	}

	h.out.Printf("-- EAX --\n\n")
	h.out.Printf("  VersionId                               = %u\n", versionId)
	h.out.Printf("  NumberOfCountersPerLogicalProcessor     = %u\n", bits(req.EAX, 15, 8))
	h.out.Printf("  BitWidthOfPerformanceMonitoringCounter  = %u\n", bits(req.EAX, 23, 16))
	h.out.Printf("  EbxBitVectorLength                      = %u\n\n", bits(req.EAX, 31, 24))

	h.out.Printf("-- EBX (bit = 1 means the event is NOT available) --\n\n")
	h.out.Printf("  CoreCycleEventNotAvailable                = %u\n", bit(req.EBX, 0))
	h.out.Printf("  InstructionRetiredEventNotAvailable       = %u\n", bit(req.EBX, 1))
	h.out.Printf("  ReferenceCyclesEventNotAvailable          = %u\n", bit(req.EBX, 2))
	h.out.Printf("  LastLevelCacheReferenceEventNotAvailable  = %u\n", bit(req.EBX, 3))
	h.out.Printf("  LastLevelCacheMissesEventNotAvailable     = %u\n", bit(req.EBX, 4))
	h.out.Printf("  BranchInstructionRetiredEventNotAvailable = %u\n", bit(req.EBX, 5))
	h.out.Printf("  BranchMispredictRetiredEventNotAvailable  = %u\n\n", bit(req.EBX, 6))

	h.out.Printf("-- ECX --\n\n")
	h.out.Printf("  Reserved                                   = 0x%08X\n\n", req.ECX)

	h.out.Printf("-- EDX --\n\n")
	if versionId > 1 {
		h.out.Printf("  NumberOfFixedFunctionPerformanceCounters   = %u\n", bits(req.EDX, 4, 0))
		h.out.Printf("  BitWidthOfFixedFunctionPerformanceCounters = %u\n\n", bits(req.EDX, 12, 5))
	} else {
		h.out.Printf("  NOTE: VersionId=%u (<=1): fixed-function counter fields are not\n", versionId)
		h.out.Printf("        architecturally defined; showing raw macro output anyway:\n")
		h.out.Printf("  NumberOfFixedFunctionPerformanceCounters   = %u (undefined)\n", bits(req.EDX, 4, 0))
		h.out.Printf("  BitWidthOfFixedFunctionPerformanceCounters = %u (undefined)\n", bits(req.EDX, 12, 5))
	}
	h.out.Printf("  AnyThreadDeprecation                       = %s\n\n", btoi(bit(req.EDX, 15+0)))
}

// showLeafB decodes CPUID.B: extended topology enumeration.
func (h *Handler) showLeafB(subFunctionId uint32, req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	if !req.LeafBSupported {
		h.out.Printf("==== CPUID.(EAX=0BH) Extended Topology Enumeration ====\n")
		h.out.Printf("  Leaf presence check failed: CPUID.0BH:EBX[15:0] == 0 at sub-leaf 0.\n")
		h.out.Printf("  This processor does not implement leaf 0BH (consider leaf 1FH instead).\n\n")
		return
	}
	h.out.Printf("==== CPUID.(EAX=0BH) Extended Topology Enumeration ====\n\n")
	h.maxSubleafBanner(req.LeafBMaxSubleaf)

	levelType := bits(req.ECX, 15, 8)
	typeName := "Reserved"
	switch levelType {
	case 0:
		typeName = "Invalid"
	case 1:
		typeName = "SMT (Hyper-Threading)"
	case 2:
		typeName = "Core"
	}
	h.out.Printf("---- CPUID.(EAX=0BH, ECX=%u) ----\n\n", subFunctionId)

	h.out.Printf("-- ECX --\n\n")
	h.out.Printf("  LevelNumber (echoes ECX input) = %u\n", bits(req.ECX, 7, 0))
	h.out.Printf("  LevelType                      = %u (%s)\n\n", levelType, typeName)

	h.out.Printf("-- EAX --\n\n")
	h.out.Printf("  X2ApicIdToUniqueTopologyIdShift = %u\n\n", bits(req.EAX, 4, 0))

	h.out.Printf("-- EBX --\n\n")
	h.out.Printf("  NumberOfLogicalProcessorsAtThisLevelType = %u (DIAGNOSTIC ONLY - do not use for topology enumeration)\n\n", bits(req.EBX, 15, 0))

	h.out.Printf("-- EDX --\n\n")
	h.out.Printf("  X2ApicId (current logical processor) = %u\n\n", req.EDX)
}

// showLeafD decodes CPUID.D: processor extended state enumeration.
func (h *Handler) showLeafD(subFunctionId uint32, req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=0DH) Processor Extended State Enumeration ====\n")
	h.out.Printf("  *******************************************************\n")
	h.out.Printf("  *             Max NumberOfSubLeaves = 62              *\n")
	h.out.Printf("  *******************************************************\n\n")

	switch {
	case subFunctionId == 0:
		h.out.Printf("-- EAX (XCR0 lower 32 bits) --\n\n")
		eax := req.EAX
		h.out.Printf("  X87State (bit 0)               = %s\n", btoi(bit(eax, 0)))
		h.out.Printf("  SSEState (bit 1)               = %s\n", btoi(bit(eax, 1)))
		h.out.Printf("  AVXState (bit 2)               = %s\n", btoi(bit(eax, 2)))
		h.out.Printf("  MPXState (bits 4:3)            = %u\n", bits(eax, 4, 3))
		h.out.Printf("  AVX512State (bits 7:5)         = %u\n", bits(eax, 7, 5))
		h.out.Printf("  UsedForIa32Xss1 (bit 8)        = %s\n", btoi(bit(eax, 8)))
		h.out.Printf("  PKRUState (bit 9)              = %s\n", btoi(bit(eax, 9)))
		h.out.Printf("  UsedForIa32Xss2 (bit 13)       = %s\n\n", btoi(bit(eax, 13)))

		h.out.Printf("-- EBX --\n\n")
		h.out.Printf("  MaxSizeRequiredByEnabledFeaturesInXcr0 = %u bytes\n\n", req.EBX)

		h.out.Printf("-- ECX --\n\n")
		h.out.Printf("  MaxSizeOfXsaveXrstorSaveArea           = %u bytes\n\n", req.ECX)

		h.out.Printf("-- EDX (XCR0 upper 32 bits) --\n\n")
		h.out.Printf("  Xcr0SupportedBits (bits 63:32 of XCR0) = 0x%08X\n\n", req.EDX)

	case subFunctionId == 1:
		h.out.Printf("-- EAX --\n\n")
		eax := req.EAX
		h.out.Printf("  SupportsXsavecAndCompactedXrstor (XSAVEC) = %s\n", btoi(bit(eax, 0)))
		h.out.Printf("  SupportsXgetbvWithEcx1                     = %s\n", btoi(bit(eax, 1)))
		h.out.Printf("  SupportsXsaveXrstorAndIa32Xss (XSAVES)     = %s\n\n", btoi(bit(eax, 3)))

		h.out.Printf("-- EBX --\n\n")
		h.out.Printf("  SizeOfXsaveArea (XCR0 | IA32_XSS enabled) = %u bytes\n", req.EBX)

		h.out.Printf("-- ECX (IA32_XSS lower 32 bits) --\n\n")
		ecx := req.ECX
		h.out.Printf("  UsedForXcr01 (bits 7:0)              = 0x%02X\n", bits(ecx, 7, 0))
		h.out.Printf("  PtState (bit 8)                      = %s\n", btoi(bit(ecx, 8)))
		h.out.Printf("  UsedForXcr02 (bit 9)                 = %s\n", btoi(bit(ecx, 9)))
		h.out.Printf("  CetUserState (bit 11)                = %s\n", btoi(bit(ecx, 11)))
		h.out.Printf("  CetSupervisorState (bit 12)          = %s\n", btoi(bit(ecx, 12)))
		h.out.Printf("  HdcState (bit 13)                    = %s\n", btoi(bit(ecx, 13)))
		h.out.Printf("  LbrState (bit 15)                    = %s\n", btoi(bit(ecx, 15)))
		h.out.Printf("  HwpState (bit 16)                    = %s\n\n", btoi(bit(ecx, 16)))

		h.out.Printf("-- EDX (IA32_XSS upper 32 bits) --\n\n")
		h.out.Printf("  SupportedUpperIa32XssBits (bits 63:32) = 0x%08X\n\n", req.EDX)

	default: // subFunctionId >= 2
		validBits := req.XCR0Vector | req.IA32_XSS_Vector
		if (validBits>>subFunctionId)&1 == 0 {
			h.out.Printf("  Sub-leaf %u is NOT supported on this CPU\n", subFunctionId)
			h.out.Printf("  Valid sub-leaves are those with bits set in:\n")
			h.out.Printf("    XCR0 vector:     0x%016llX\n", req.XCR0Vector)
			h.out.Printf("    IA32_XSS vector: 0x%016llX\n", req.IA32_XSS_Vector)
			h.out.Printf("    Combined vector: 0x%016llX\n\n", validBits)
			return
		}
		h.out.Printf("-- State Component Information --\n\n")
		h.out.Printf("  SaveAreaSize (EAX)   = %u bytes\n", req.EAX)
		h.out.Printf("  SaveAreaOffset (EBX) = %u bytes\n", req.EBX)
		managed := bit(req.ECX, 0)
		h.out.Printf("  ManagedViaIa32Xss (ECX bit 0) = %u (%s)\n", managed, func() string {
			if managed {
				return "IA32_XSS"
			}
			return "XCR0"
		}())
		aligned := bit(req.ECX, 1)
		h.out.Printf("  Aligned64ByteBoundary (ECX bit 1) = %u%s\n", aligned, func() string {
			if aligned {
				return " (next 64-byte boundary)"
			}
			return " (immediately following)"
		}())
		h.out.Printf("  Reserved (EDX) = 0x%08X\n\n", req.EDX)
	}
}

// showLeaf10 decodes CPUID.10: Intel RDT allocation information.
func (h *Handler) showLeaf10(subFunctionId uint32, req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=10H, ECX=%u) Intel RDT Allocation Information ====\n\n", subFunctionId)
	h.out.Printf("  *******************************************************\n")
	h.out.Printf("  *             Max NumberOfSubLeaves = 3               *\n")
	h.out.Printf("  *******************************************************\n\n")

	switch subFunctionId {
	case 0:
		h.out.Printf("-- EAX --\n\n")
		h.out.Printf("  Ia32PlatformDcaCap = 0x%08X\n\n", req.EAX)
		h.out.Printf("-- EBX (Supported Allocation Types) --\n\n")
		h.out.Printf("  Raw value = 0x%08X\n", req.EBX)
		h.out.Printf("  L3 Cache Allocation (bit 1) = %s\n", btoi(bit(req.EBX, 1)))
		h.out.Printf("  L2 Cache Allocation (bit 2) = %s\n", btoi(bit(req.EBX, 2)))
		h.out.Printf("  Memory Bandwidth Allocation (bit 3) = %s\n\n", btoi(bit(req.EBX, 3)))
		h.out.Printf("-- ECX --\n\n")
		h.out.Printf("  Reserved = 0x%08X\n\n", req.ECX)
		h.out.Printf("-- EDX --\n\n")
		h.out.Printf("  Reserved = 0x%08X\n\n", req.EDX)

	case 1:
		h.out.Printf("-- EAX --\n\n")
		h.out.Printf("  LengthOfCapacityBitMask (minus-one) = %u (actual = %u bits)\n\n", bits(req.EAX, 11, 0), bits(req.EAX, 11, 0)+1)
		h.out.Printf("-- EBX --\n\n")
		h.out.Printf("  Bit-granular map = 0x%08X\n\n", req.EBX)
		h.out.Printf("-- ECX --\n\n")
		cdp := bit(req.ECX, 2)
		h.out.Printf("  CodeAndDataPriorizationTechnologySupported = %s%s\n\n", btoi(cdp), func() string {
			if cdp {
				return " (CDP supported)"
			}
			return ""
		}())
		h.out.Printf("-- EDX --\n\n")
		cos := bits(req.EDX, 15, 0)
		h.out.Printf("  HighestCosNumberSupported = %u (actual = %u COS)\n\n", cos, cos+1)

	case 2:
		h.out.Printf("-- EAX --\n\n")
		h.out.Printf("  LengthOfCapacityBitMask (minus-one) = %u (actual = %u bits)\n\n", bits(req.EAX, 11, 0), bits(req.EAX, 11, 0)+1)
		h.out.Printf("-- EBX --\n\n")
		h.out.Printf("  Bit-granular map = 0x%08X\n\n", req.EBX)
		h.out.Printf("-- ECX --\n\n")
		h.out.Printf("  Reserved = 0x%08X\n\n", req.ECX)
		h.out.Printf("-- EDX --\n\n")
		cos := bits(req.EDX, 15, 0)
		h.out.Printf("  HighestCosNumberSupported = %u (actual = %u COS)\n\n", cos, cos+1)

	case 3:
		h.out.Printf("-- EAX --\n\n")
		mba := bits(req.EAX, 11, 0)
		h.out.Printf("  MaxMbaThrottlingValue (minus-one) = %u (actual = %u)\n\n", mba, mba+1)
		h.out.Printf("-- EBX --\n\n")
		h.out.Printf("  Reserved = 0x%08X\n\n", req.EBX)
		h.out.Printf("-- ECX --\n\n")
		linear := bit(req.ECX, 2)
		h.out.Printf("  ResponseOfDelayIsLinear = %s%s\n", btoi(linear), func() string {
			if linear {
				return " (linear response)"
			}
			return " (non-linear)\n\n"
		}())
		h.out.Printf("-- EDX --\n\n")
		cos := bits(req.EDX, 15, 0)
		h.out.Printf("  HighestCosNumberSupported = %u (actual = %u COS)\n\n", cos, cos+1)

	default:
		h.out.Printf("  Sub-leaf %u is NOT supported on this CPU\n", subFunctionId)
		h.out.Printf("  raw bytes: EAX = %u\n  EBX = %u\n  ECX = %u\n  EDX = %u\n\n",
			req.EAX, req.EBX, req.ECX, req.EDX)
	}
}

// showLeaf12 decodes CPUID.12: Intel SGX information.
func (h *Handler) showLeaf12(subFunctionId uint32, req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=12H) Intel SGX Information ====\n\n")
	if !req.Leaf12Supported {
		h.out.Printf("  SGX is not supported on this CPU (CPUID.7H:EBX[2] = 0).\n\n")
		return
	}
	h.maxSubleafBanner(req.Leaf12MaxSubLeaf)

	switch {
	case subFunctionId == 0:
		h.out.Printf("-- EAX (SGX Capabilities) --\n\n")
		eax := req.EAX
		h.out.Printf("  SGX1 Support (bit 0)           = %s\n", btoi(bit(eax, 0)))
		h.out.Printf("  SGX2 Support (bit 1)           = %s\n", btoi(bit(eax, 1)))
		h.out.Printf("  ENCLV Advanced (bit 5)         = %s\n", btoi(bit(eax, 5)))
		h.out.Printf("  ENCLS Advanced (bit 6)         = %s\n\n", btoi(bit(eax, 6)))

		h.out.Printf("-- EBX (MISCSELECT - Extended SGX Features) --\n\n")
		h.out.Printf("  Miscselect = 0x%08X\n\n", req.EBX)
		h.out.Printf("-- ECX --\n\n")
		h.out.Printf("  Reserved = 0x%08X\n\n", req.ECX)

		h.out.Printf("-- EDX (Maximum Enclave Sizes) --\n\n")
		maxNot64 := bits(req.EDX, 7, 0)
		max64 := bits(req.EDX, 15, 8)
		h.out.Printf("  MaxEnclaveSizeNot64 = %u (2^%u = %llu bytes)\n", maxNot64, maxNot64, uint64(1)<<maxNot64)
		h.out.Printf("  MaxEnclaveSize64     = %u (2^%u = %llu bytes)\n\n", max64, max64, uint64(1)<<max64)

	case subFunctionId == 1:
		h.out.Printf("-- EAX (SECS.ATTRIBUTES[31:0]) --\n\n")
		h.out.Printf("  ValidSecsAttributes0 = 0x%08X\n\n", req.EAX)
		h.out.Printf("-- EBX (SECS.ATTRIBUTES[63:32]) --\n\n")
		h.out.Printf("  ValidSecsAttributes1 = 0x%08X\n\n", req.EBX)
		h.out.Printf("-- ECX (SECS.ATTRIBUTES[95:64]) --\n\n")
		h.out.Printf("  ValidSecsAttributes2 = 0x%08X\n\n", req.ECX)
		h.out.Printf("-- EDX (SECS.ATTRIBUTES[127:96]) --\n\n")
		h.out.Printf("  ValidSecsAttributes3 = 0x%08X\n\n", req.EDX)

	default:
		subLeafType := bits(req.EAX, 3, 0)
		switch subLeafType {
		case 0:
			h.out.Printf("-- EAX --\n\n")
			h.out.Printf("  SubLeafType = %u (Invalid - no EPC section)\n", subLeafType)
			h.out.Printf("\n-- EBX --\n")
			h.out.Printf("  Zero = 0x%08X\n", req.EBX)
			h.out.Printf("\n-- ECX --\n")
			h.out.Printf("  Zero = 0x%08X\n", req.ECX)
			h.out.Printf("\n-- EDX --\n")
			h.out.Printf("  Zero = 0x%08X\n", req.EDX)
		case 1:
			baseLow := uint64(bits(req.EAX, 31, 12))
			baseHigh := uint64(bits(req.EBX, 19, 0))
			baseAddr := (baseHigh << 32) | (baseLow << 12)
			sizeLow := uint64(bits(req.ECX, 31, 12))
			sizeHigh := uint64(bits(req.EDX, 19, 0))
			sizeBytes := (sizeHigh << 32) | (sizeLow << 12)

			h.out.Printf("-- EAX --\n\n")
			h.out.Printf("  SubLeafType = %u (EPC Section)\n", subLeafType)
			h.out.Printf("  EpcBasePhysicalAddress1 (bits 31:12) = 0x%05X\n\n", bits(req.EAX, 31, 12))
			h.out.Printf("-- EBX --\n\n")
			h.out.Printf("  EpcBasePhysicalAddress2 (bits 51:32) = 0x%05X\n\n", bits(req.EBX, 19, 0))
			h.out.Printf("-- ECX --\n\n")
			prop := bits(req.ECX, 3, 0)
			h.out.Printf("  EpcSectionProperty = %u", prop)
			switch prop {
			case 0:
				h.out.Printf(" (No confidentiality/integrity)\n\n")
			case 1:
				h.out.Printf(" (Confidentiality and integrity protection)\n\n")
			default:
				h.out.Printf(" (Reserved)\n\n")
			}
			h.out.Printf("  EpcSize1 (bits 31:12) = 0x%05X\n\n", bits(req.ECX, 31, 12))
			h.out.Printf("-- EDX --\n\n")
			h.out.Printf("  EpcSize2 (bits 51:32) = 0x%05X\n\n", bits(req.EDX, 19, 0))
			h.out.Printf("-- EPC Section Information --\n\n")
			h.out.Printf("  Base Address = 0x%016llX\n", baseAddr)
			h.out.Printf("  Size         = 0x%016llX bytes (%llu MB)\n\n", sizeBytes, sizeBytes/(1024*1024))
		default:
			h.out.Printf("-- Raw Data for Sub-leaf %u (Type %u) --\n\n", subFunctionId, subLeafType)
			h.out.Printf("  EAX = 0x%08X\n", req.EAX)
			h.out.Printf("  EBX = 0x%08X\n", req.EBX)
			h.out.Printf("  ECX = 0x%08X\n", req.ECX)
			h.out.Printf("  EDX = 0x%08X\n", req.EDX)
			h.out.Printf("  Note: Reserved sub-leaf type. Please refer to the latest Intel SDM.\n\n")
		}
	}
}

// showLeaf14 decodes CPUID.14: Intel Processor Trace information.
func (h *Handler) showLeaf14(subFunctionId uint32, req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=14H) Intel Processor Trace Information ====\n\n")
	h.maxSubleafBanner(req.LeafEaxMaxSubleaf)
	h.out.Printf("==== CPUID.(EAX=14H, ECX=%u) Intel Processor Trace Information ====\n\n", subFunctionId)

	switch subFunctionId {
	case 0:
		h.out.Printf("-- EAX --\n\n")
		h.out.Printf("  MaxSubLeaf = %u\n\n", bits(req.EAX, 31, 0))

		h.out.Printf("-- EBX (Supported Features) --\n\n")
		ebx := req.EBX
		h.out.Printf("  CR3Filter support (bit 0)                    = %s\n", btoi(bit(ebx, 0)))
		h.out.Printf("  Configurable PSB & Cycle-Accurate (bit 1)    = %s\n", btoi(bit(ebx, 1)))
		h.out.Printf("  IP Filtering & TraceStop (bit 2)             = %s\n", btoi(bit(ebx, 2)))
		h.out.Printf("  MTC & COFI suppression (bit 3)               = %s\n", btoi(bit(ebx, 3)))
		h.out.Printf("  PTWRITE support (bit 4)                      = %s\n", btoi(bit(ebx, 4)))
		h.out.Printf("  Power Event Trace (bit 5)                    = %s\n", btoi(bit(ebx, 5)))
		h.out.Printf("  PSB/PMI preservation (bit 6)                 = %s\n", btoi(bit(ebx, 6)))
		h.out.Printf("  Event Trace (bit 7)                          = %s\n", btoi(bit(ebx, 7)))
		h.out.Printf("  Disable TNT (bit 8)                          = %s\n\n", btoi(bit(ebx, 8)))

		h.out.Printf("-- ECX (Output Schemes) --\n\n")
		ecx := req.ECX
		h.out.Printf("  ToPA output scheme (bit 0)                   = %s\n", btoi(bit(ecx, 0)))
		h.out.Printf("  ToPA tables with any entries (bit 1)         = %s\n", btoi(bit(ecx, 1)))
		h.out.Printf("  Single-Range Output scheme (bit 2)           = %s\n", btoi(bit(ecx, 2)))
		h.out.Printf("  Trace Transport subsystem (bit 3)            = %s\n", btoi(bit(ecx, 3)))
		h.out.Printf("  LIP values include CS base (bit 31)          = %s\n\n", btoi(bit(ecx, 31)))

		h.out.Printf("-- EDX --\n\n")
		h.out.Printf("  Reserved = 0x%08X\n\n", req.EDX)

	case 1:
		h.out.Printf("-- EAX --\n\n")
		h.out.Printf("  NumberOfConfigurableAddressRangesForFiltering = %u\n", bits(req.EAX, 2, 0))
		h.out.Printf("  BitmapOfSupportedMtcPeriodEncodings          = 0x%04X\n\n", bits(req.EAX, 31, 16))

		h.out.Printf("-- EBX --\n\n")
		h.out.Printf("  BitmapOfSupportedCycleThresholdValueEncodings = 0x%04X\n", bits(req.EBX, 15, 0))
		h.out.Printf("  BitmapOfSupportedConfigurablePsbFrequencyEncodings = 0x%04X\n\n", bits(req.EBX, 31, 16))

		h.out.Printf("-- ECX --\n\n")
		h.out.Printf("  Reserved = 0x%08X\n\n", req.ECX)
		h.out.Printf("-- EDX --\n\n")
		h.out.Printf("  Reserved = 0x%08X\n\n", req.EDX)

	default:
		h.out.Printf("-- Raw Data for Sub-leaf %u --\n\n", subFunctionId)
		h.out.Printf("  EAX = 0x%08X\n", req.EAX)
		h.out.Printf("  EBX = 0x%08X\n", req.EBX)
		h.out.Printf("  ECX = 0x%08X\n", req.ECX)
		h.out.Printf("  EDX = 0x%08X\n", req.EDX)
		h.out.Printf("  Note: This sub-leaf may be for future Intel PT features.\n")
		h.out.Printf("  Please refer to the latest Intel SDM for interpretation.\n\n")
	}
}

// showLeaf15 decodes CPUID.15: TSC and core crystal clock information.
func (h *Handler) showLeaf15(req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=15H) TSC and Core Crystal Clock Information ====\n\n")
	h.noSubleavesBanner("15h")

	denominator := req.EAX
	numerator := req.EBX
	nominalFreq := req.ECX

	h.out.Printf("-- EAX --\n\n")
	h.out.Printf("  Denominator = %u\n\n", denominator)
	h.out.Printf("-- EBX --\n\n")
	h.out.Printf("  Numerator = %u\n\n", numerator)
	h.out.Printf("-- ECX --\n\n")
	h.out.Printf("  NominalFrequency = %u Hz", nominalFreq)
	if nominalFreq > 0 {
		h.out.Printf(" (%u MHz, %u GHz)\n", nominalFreq/1000000, nominalFreq/1000000000)
	} else {
		h.out.Printf("  (not enumerated)\n\n")
	}
	h.out.Printf("\n-- EDX --\n")
	h.out.Printf("  Reserved = 0x%08X\n", req.EDX)

	h.out.Printf("-- TSC Frequency Information --\n\n")
	if denominator == 0 || numerator == 0 {
		h.out.Printf("  TSC ratio not enumerated (denominator or numerator is 0)\n\n")
		return
	}
	h.out.Printf("  TSC / Core Crystal Clock ratio = %u / %u\n", numerator, denominator)
	h.out.Printf("  TSC frequency = Core Crystal Clock * (%u/%u)\n", numerator, denominator)
	if nominalFreq > 0 {
		tscFreq := uint64(nominalFreq) * uint64(numerator) / uint64(denominator)
		h.out.Printf("  TSC frequency = %llu Hz (%llu MHz, %llu GHz)\n\n",
			tscFreq, tscFreq/1000000, tscFreq/1000000000)
	} else {
		h.out.Printf("  TSC frequency cannot be calculated (nominal frequency not enumerated)\n\n")
	}
}

// showLeaf16 decodes CPUID.16: processor frequency information.
func (h *Handler) showLeaf16(req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=16H) Processor Frequency Information ====\n\n")
	h.noSubleavesBanner("16h")

	baseFreq := req.EAX
	maxFreq := req.EBX
	busFreq := req.ECX

	h.out.Printf("-- EAX --\n\n")
	if baseFreq == 0 {
		h.out.Printf("  ProcessorBaseFrequencyMhz = 0 (not supported)\n\n")
	} else {
		h.out.Printf("  ProcessorBaseFrequencyMhz = %u MHz\n", baseFreq)
	}

	h.out.Printf("-- EBX --\n\n")
	if maxFreq == 0 {
		h.out.Printf("  ProcessorMaximumFrequencyMhz = 0 (not supported)\n\n")
	} else {
		h.out.Printf("  ProcessorMaximumFrequencyMhz = %u MHz\n", maxFreq)
	}

	h.out.Printf("-- ECX --\n\n")
	if busFreq == 0 {
		h.out.Printf("  BusFrequencyMhz = 0 (not supported)\n\n")
	} else {
		h.out.Printf("  BusFrequencyMhz = %u MHz\n\n", busFreq)
	}

	h.out.Printf("-- EDX --\n\n")
	h.out.Printf("  Reserved = 0x%08X\n\n", req.EDX)

	h.out.Printf("-- Frequency Summary --\n\n")
	h.out.Printf("  Base Frequency:     %s\n", freqOrNotSupported(baseFreq))
	h.out.Printf("  Maximum Frequency:  %s\n", freqOrNotSupported(maxFreq))
	h.out.Printf("  Bus Frequency:      %s\n\n", freqOrNotSupported(busFreq))

	if baseFreq > 0 && maxFreq > 0 && maxFreq > baseFreq {
		turbo := maxFreq - baseFreq
		pct := float64(turbo) / float64(baseFreq) * 100.0
		h.out.Printf("  Turbo Boost:        %u MHz above base (%.1f%% increase)\n\n", turbo, pct)
	}

	h.out.Printf("  *******************************************************\n")
	h.out.Printf("  *  NOTE: These values are for display purposes only   *\n")
	h.out.Printf("  *  They do not reflect actual running frequencies     *\n")
	h.out.Printf("  *  Actual frequencies depend on workload, power, etc. *\n")
	h.out.Printf("  *******************************************************\n\n")
}

// freqOrNotSupported returns "%u MHz" for non-zero frequencies and
// "Not Supported" for zero, matching the C++ leaf 16 summary block.
func freqOrNotSupported(mhz uint32) string {
	if mhz == 0 {
		return "Not Supported"
	}
	return fmt.Sprintf("%d MHz", mhz)
}

// showLeaf17 decodes CPUID.17: SoC vendor information.
func (h *Handler) showLeaf17(subFunctionId uint32, req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=17H) SoC Vendor Information ====\n\n")
	h.maxSubleafBanner(bits(req.EAX, 31, 0))
	h.out.Printf("==== CPUID.(EAX=17H, ECX=%u) SoC Vendor Information ====\n\n", subFunctionId)

	switch {
	case subFunctionId == 0:
		h.out.Printf("-- EAX --\n\n")
		h.out.Printf("  MaxSocIdIndex = %u\n\n", bits(req.EAX, 31, 0))
		h.out.Printf("-- EBX --\n\n")
		h.out.Printf("  SocVendorId = 0x%04X\n", bits(req.EBX, 15, 0))
		h.out.Printf("  IsVendorScheme = %s\n\n", func() string {
			if bit(req.EBX, 16) {
				return "TRUE (Industry Standard)"
			}
			return "FALSE (Intel Assigned)"
		}())
		h.out.Printf("-- ECX --\n\n")
		h.out.Printf("  ProjectId = 0x%08X\n\n", req.ECX)
		h.out.Printf("-- EDX --\n\n")
		h.out.Printf("  SteppingId = 0x%08X\n\n", req.EDX)

	case subFunctionId >= 1 && subFunctionId <= 3:
		printBrandReg := func(name string, v uint32) {
			h.out.Printf("  %s = 0x%08X (%c%c%c%c)\n\n", name, v,
				byte(v&0xFF), byte((v>>8)&0xFF), byte((v>>16)&0xFF), byte((v>>24)&0xFF))
		}
		h.out.Printf("-- EAX --\n\n")
		printBrandReg("SocVendorBrandString[0..3]", req.EAX)
		h.out.Printf("-- EBX --\n\n")
		printBrandReg("SocVendorBrandString[4..7]", req.EBX)
		h.out.Printf("-- ECX --\n\n")
		printBrandReg("SocVendorBrandString[8..11]", req.ECX)
		h.out.Printf("-- EDX --\n\n")
		printBrandReg("SocVendorBrandString[12..15]", req.EDX)

	default:
		h.out.Printf("-- EAX --\n\n")
		h.out.Printf("  Reserved = 0x%08X\n\n", req.EAX)
		h.out.Printf("-- EBX --\n\n")
		h.out.Printf("  Reserved = 0x%08X\n\n", req.EBX)
		h.out.Printf("-- ECX --\n\n")
		h.out.Printf("  Reserved = 0x%08X\n\n", req.ECX)
		h.out.Printf("-- EDX --\n\n")
		h.out.Printf("  Reserved = 0x%08X\n\n", req.EDX)
	}
}

// showLeaf18 decodes CPUID.18: deterministic address translation parameters.
func (h *Handler) showLeaf18(subFunctionId uint32, req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=18H) Deterministic Address Translation Parameters ====\n\n")
	h.maxSubleafBanner(req.LeafEaxMaxSubleaf)
	h.out.Printf("==== CPUID.(EAX=18H, ECX=%u) Deterministic Address Translation Parameters ====\n\n", subFunctionId)

	tlbTypeName := func(t uint32) string {
		switch t {
		case 0:
			return "Null (sub-leaf not valid)"
		case 1:
			return "Data TLB"
		case 2:
			return "Instruction TLB"
		case 3:
			return "Unified TLB"
		default:
			return "Reserved"
		}
	}

	if subFunctionId == 0 {
		h.out.Printf("-- EAX --\n\n")
		h.out.Printf("  MaxSubLeaf = %u\n\n", bits(req.EAX, 31, 0))

		h.out.Printf("-- EBX --\n\n")
		ebx := req.EBX
		h.out.Printf("  PageEntries4KbSupported    = %s\n", btoi(bit(ebx, 0)))
		h.out.Printf("  PageEntries2MbSupported    = %s\n", btoi(bit(ebx, 1)))
		h.out.Printf("  PageEntries4MbSupported    = %s\n", btoi(bit(ebx, 2)))
		h.out.Printf("  PageEntries1GbSupported    = %s\n", btoi(bit(ebx, 3)))
		part := bits(ebx, 15, 8)
		h.out.Printf("  Partitioning               = %u%s\n", part, func() string {
			if part == 0 {
				return " (soft partitioning)"
			}
			return ""
		}())
		h.out.Printf("  WaysOfAssociativity (W)    = %u\n\n", bits(ebx, 31, 16))

		h.out.Printf("-- ECX --\n\n")
		h.out.Printf("  NumberOfSets               = %u\n\n", req.ECX)

		h.out.Printf("-- EDX --\n\n")
		tlbType := bits(req.EDX, 4, 0)
		h.out.Printf("  TranslationCacheTypeField  = %u (%s)\n", tlbType, tlbTypeName(tlbType))
		h.out.Printf("  TranslationCacheLevel      = %u\n", bits(req.EDX, 7, 5))
		fa := bit(req.EDX, 9)
		h.out.Printf("  FullyAssociativeStructure  = %s%s\n", btoi(fa), func() string {
			if fa {
				return " (fully associative)"
			}
			return ""
		}())
		maxIds := bits(req.EDX, 23, 16)
		h.out.Printf("  MaxAddressableIdsForLogicalProcessors (raw) = %u -> actual = %u (raw + 1)\n", maxIds, maxIds+1)
		return
	}

	// subFunctionId >= 1
	if bits(req.EDX, 4, 0) == 0 {
		h.out.Printf("  Sub-leaf %u is invalid (TranslationCacheTypeField = 0)\n", subFunctionId)
		h.out.Printf("  All registers should be zero according to spec:\n")
		h.out.Printf("  EAX = 0x%08X\n", req.EAX)
		h.out.Printf("  EBX = 0x%08X\n", req.EBX)
		h.out.Printf("  ECX = 0x%08X\n", req.ECX)
		h.out.Printf("  EDX = 0x%08X\n", req.EDX)
		return
	}

	h.out.Printf("-- EAX --\n\n")
	h.out.Printf("  Reserved = 0x%08X\n\n", req.EAX)

	h.out.Printf("-- EBX --\n\n")
	ebx := req.EBX
	h.out.Printf("  PageEntries4KbSupported    = %s\n", btoi(bit(ebx, 0)))
	h.out.Printf("  PageEntries2MbSupported    = %s\n", btoi(bit(ebx, 1)))
	h.out.Printf("  PageEntries4MbSupported    = %s\n", btoi(bit(ebx, 2)))
	h.out.Printf("  PageEntries1GbSupported    = %s\n", btoi(bit(ebx, 3)))
	part := bits(ebx, 15, 8)
	h.out.Printf("  Partitioning               = %u%s\n", part, func() string {
		if part == 0 {
			return " (soft partitioning)"
		}
		return ""
	}())
	h.out.Printf("  WaysOfAssociativity (W)    = %u\n\n", bits(ebx, 31, 16))

	h.out.Printf("-- ECX --\n\n")
	h.out.Printf("  NumberOfSets               = %u\n\n", req.ECX)

	h.out.Printf("-- EDX --\n\n")
	tlbType := bits(req.EDX, 4, 0)
	h.out.Printf("  TranslationCacheTypeField  = %u (%s)\n", tlbType, tlbTypeName(tlbType))
	h.out.Printf("  TranslationCacheLevel      = %u\n", bits(req.EDX, 7, 5))
	fa := bit(req.EDX, 9)
	h.out.Printf("  FullyAssociativeStructure  = %s%s\n", btoi(fa), func() string {
		if fa {
			return " (fully associative)"
		}
		return ""
	}())
	maxIds := bits(req.EDX, 23, 16)
	h.out.Printf("  MaxAddressableIdsForLogicalProcessors (raw) = %u -> actual = %u (raw + 1)\n", maxIds, maxIds+1)
}

// ---------------------------------------------------------------------------
// Extended leaves (0x80000000 - 0x80000008)
// ---------------------------------------------------------------------------

// showLeaf80000000 decodes CPUID.80000000: extended function information.
func (h *Handler) showLeaf80000000(req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=80000000H) Extended Function Information ====\n\n")
	h.out.Printf("  *******************************************************\n")
	h.out.Printf("  *           LEAF 0x80000000 HAS NO SUBLEAVES          *\n")
	h.out.Printf("  *       ANY SUBLEAF YOU ENTER WILL DEFAULT TO 0       *\n")
	h.out.Printf("  *      AND THE PROCESSOR RETURNS UNDEFINED VALUES     *\n")
	h.out.Printf("  *******************************************************\n\n")

	h.out.Printf("-- EAX --\n\n")
	h.out.Printf("  MaxExtendedFunctions = 0x%08X (%u)\n\n", req.EAX, req.EAX)
	h.out.Printf("-- EBX --\n\n")
	h.out.Printf("  Reserved = 0x%08X\n\n", req.EBX)
	h.out.Printf("-- ECX --\n\n")
	h.out.Printf("  Reserved = 0x%08X\n\n", req.ECX)
	h.out.Printf("-- EDX --\n\n")
	h.out.Printf("  Reserved = 0x%08X\n\n", req.EDX)
}

// showLeaf80000001 decodes CPUID.80000001: extended CPU signature.
func (h *Handler) showLeaf80000001(req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=80000001H) Extended CPU Signature ====\n\n")
	h.out.Printf("  *******************************************************\n")
	h.out.Printf("  *           LEAF 0x80000001 HAS NO SUBLEAVES          *\n")
	h.out.Printf("  *       ANY SUBLEAF YOU ENTER WILL DEFAULT TO 0       *\n")
	h.out.Printf("  *      AND THE PROCESSOR RETURNS UNDEFINED VALUES     *\n")
	h.out.Printf("  *******************************************************\n\n")

	h.out.Printf("-- EAX --\n\n")
	h.out.Printf("  Reserved = 0x%08X\n\n", req.EAX)
	h.out.Printf("-- EBX --\n\n")
	h.out.Printf("  Reserved = 0x%08X\n\n", req.EBX)

	h.out.Printf("-- ECX --\n\n")
	ecx := req.ECX
	h.out.Printf("  LAHF/SAHF Available in 64-bit Mode = %s\n", btoi(bit(ecx, 0)))
	h.out.Printf("  LZCNT                              = %s\n", btoi(bit(ecx, 5)))
	h.out.Printf("  PREFETCHW                          = %s\n\n", btoi(bit(ecx, 8)))

	h.out.Printf("-- EDX --\n\n")
	edx := req.EDX
	h.out.Printf("  SYSCALL/SYSRET Available in 64-bit Mode = %s\n", btoi(bit(edx, 11)))
	h.out.Printf("  Execute Disable Bit Available            = %s\n", btoi(bit(edx, 20)))
	h.out.Printf("  1-GByte Pages Available                 = %s\n", btoi(bit(edx, 26)))
	h.out.Printf("  RDTSCP Available                        = %s\n", btoi(bit(edx, 27)))
	h.out.Printf("  Intel 64 Architecture Available         = %s\n\n", btoi(bit(edx, 29)))
}

// showLeaf80000002to04 decodes CPUID.80000002-4: processor brand string.
func (h *Handler) showLeaf80000002to04(req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=80000002H-80000004H) Processor Information ====\n\n")
	h.out.Printf("  *******************************************************\n")
	h.out.Printf("  *         LEAF 0x80000002-4 HAS NO SUBLEAVES          *\n")
	h.out.Printf("  *       ANY SUBLEAF YOU ENTER WILL DEFAULT TO 0       *\n")
	h.out.Printf("  *      AND THE PROCESSOR RETURNS UNDEFINED VALUES     *\n")
	h.out.Printf("  *******************************************************\n\n")
	h.out.Printf("  Brand String = \"%s\"\n\n", brandString(req))
}

// showLeaf80000006 decodes CPUID.80000006: extended cache information.
func (h *Handler) showLeaf80000006(req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=80000006H) Extended Cache Information ====\n\n")
	h.out.Printf("  *******************************************************\n")
	h.out.Printf("  *          LEAF 0x80000006 HAS NO SUBLEAVES           *\n")
	h.out.Printf("  *       ANY SUBLEAF YOU ENTER WILL DEFAULT TO 0       *\n")
	h.out.Printf("  *      AND THE PROCESSOR RETURNS UNDEFINED VALUES     *\n")
	h.out.Printf("  *******************************************************\n\n")

	h.out.Printf("-- EAX --\n\n")
	h.out.Printf("  Reserved = 0x%08X\n\n", req.EAX)
	h.out.Printf("-- EBX --\n\n")
	h.out.Printf("  Reserved = 0x%08X\n\n", req.EBX)

	h.out.Printf("-- ECX (L2 Cache Information) --\n")
	lineSize := bits(req.ECX, 7, 0)
	assoc := bits(req.ECX, 15, 12)
	cacheSize := bits(req.ECX, 31, 16)
	assocName := "Reserved"
	switch assoc {
	case 0x00:
		assocName = "Disabled"
	case 0x01:
		assocName = "Direct mapped"
	case 0x02:
		assocName = "2-way"
	case 0x04:
		assocName = "4-way"
	case 0x06:
		assocName = "8-way"
	case 0x08:
		assocName = "16-way"
	case 0x0F:
		assocName = "Fully associative"
	}
	h.out.Printf("  CacheLineSizeInBytes       = %u bytes\n", lineSize)
	h.out.Printf("  L2AssociativityField       = 0x%02X (%s)\n", assoc, assocName)
	h.out.Printf("  CacheSizeIn1KUnits         = %u KB (%u MB)\n", cacheSize, cacheSize/1024)

	h.out.Printf("-- EDX --\n\n")
	h.out.Printf("  Reserved = 0x%08X\n\n", req.EDX)
}

// showLeaf80000007 decodes CPUID.80000007: extended TSC information.
func (h *Handler) showLeaf80000007(req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=80000007H) Extended Time Stamp Counter ====\n\n")
	h.out.Printf("  *******************************************************\n")
	h.out.Printf("  *          LEAF 0x80000007 HAS NO SUBLEAVES           *\n")
	h.out.Printf("  *       ANY SUBLEAF YOU ENTER WILL DEFAULT TO 0       *\n")
	h.out.Printf("  *      AND THE PROCESSOR RETURNS UNDEFINED VALUES     *\n")
	h.out.Printf("  *******************************************************\n\n")

	h.out.Printf("-- EAX --\n\n")
	h.out.Printf("  Reserved = 0x%08X\n\n", req.EAX)
	h.out.Printf("-- EBX --\n\n")
	h.out.Printf("  Reserved = 0x%08X\n\n", req.EBX)
	h.out.Printf("-- ECX --\n\n")
	h.out.Printf("  Reserved = 0x%08X\n\n", req.ECX)
	h.out.Printf("-- EDX --\n\n")
	inv := bit(req.EDX, 8)
	h.out.Printf("  InvariantTscAvailable = %u%s\n\n", inv, func() string {
		if inv {
			return " (TSC runs at constant rate)"
		}
		return ""
	}())
}

// showLeaf80000008 decodes CPUID.80000008: virtual & physical address sizes.
func (h *Handler) showLeaf80000008(req *hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE) {
	h.out.Printf("==== CPUID.(EAX=80000008H) Virtual & Physical Address Sizes ====\n\n")
	h.out.Printf("  *******************************************************\n")
	h.out.Printf("  *          LEAF 0x80000008 HAS NO SUBLEAVES           *\n")
	h.out.Printf("  *       ANY SUBLEAF YOU ENTER WILL DEFAULT TO 0       *\n")
	h.out.Printf("  *      AND THE PROCESSOR RETURNS UNDEFINED VALUES     *\n")
	h.out.Printf("  *******************************************************\n\n")

	h.out.Printf("-- EAX --\n\n")
	physicalBits := bits(req.EAX, 7, 0)
	linearBits := bits(req.EAX, 15, 8)
	h.out.Printf("  NumberOfPhysicalAddressBits = %u (max physical address = 2^%u = %llu bytes)\n",
		physicalBits, physicalBits, uint64(1)<<physicalBits)
	h.out.Printf("  NumberOfLinearAddressBits  = %u (max linear address = 2^%u = %llu bytes)\n",
		linearBits, linearBits, uint64(1)<<linearBits)

	h.out.Printf("-- EBX --\n\n")
	h.out.Printf("  Reserved = 0x%08X\n\n", req.EBX)
	h.out.Printf("-- ECX --\n\n")
	h.out.Printf("  Reserved = 0x%08X\n\n", req.ECX)
	h.out.Printf("-- EDX --\n\n")
	h.out.Printf("  Reserved = 0x%08X\n\n", req.EDX)
}

// discardOutput is the default Output when the caller passes nil.
type discardOutput struct{}

func (discardOutput) Printf(format string, args ...any) error {
	_ = args
	_ = format
	return nil
}
