// Package transparency implements the HyperDbg anti-detection module.
//
// This file mirrors HyperDbg/hyperdbg/libhyperdbg/code/debugger/transparency/
// transparency.cpp (hypervisor-presence measurement) and the
// hide.cpp / unhide.cpp extension-command handlers (IOCTL to enable/disable
// transparent mode). It provides:
//
//   - Measurement helpers (RdtscDiffVmexit, RdtscVmexitTracing,
//     CpuidTimeStampCounter, RdtscEmulationDetection, CheckHypervisorPresence,
//     CheckRdtscpVmexit) that detect a hypervisor by timing rdtsc+cpuid+rdtsc
//     and rdtsc+rdtsc sequences.
//   - IOCTL helpers (HideDebugger, HideDebuggerByName, HideDebuggerEx,
//     UnhideDebugger) that send IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_
//     THE_DEBUGGER to the VMM driver to enable/disable kernel-side
//     transparent-mode simulation.
//   - A combined entry point (TransparentHypervisor) that runs both detection
//     checks and reports a single result.
//   - CPUID/RDTSC simulation helpers (SimulateCpuidResult, SimulateRdtscValue)
//     that add Gaussian noise to a measured value, mirroring what the kernel
//     does once transparent mode is enabled.
//
// API design follows the go-libhyperdbg conventions: no global state, context
// propagation on blocking calls, an Output interface for CLI/GUI/MCP reuse,
// errors returned via `error`, and concurrent-safe methods.
package transparency

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"unsafe"

	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
	"github.com/hyperdbg/go-libhyperdbg/types"
)

// ----------------------------------------------------------------------------
// Public constants — mirror C++ #defines used by the transparency code.
// ----------------------------------------------------------------------------

// Transparent-mode evade-mask bits. Mirror TRANSPARENT_EVADE_MASK_* in
// include/SDK/headers/Constants.h.
const (
	EvadeMaskSyscallHook             uint32 = 0x00000001 // TRANSPARENT_EVADE_MASK_SYSCALL_HOOK
	EvadeMaskCpuid                   uint32 = 0x00000002 // TRANSPARENT_EVADE_MASK_CPUID
	EvadeMaskMsr                     uint32 = 0x00000004 // TRANSPARENT_EVADE_MASK_MSR
	EvadeMaskTrapFlag                uint32 = 0x00000008 // TRANSPARENT_EVADE_MASK_TRAP_FLAG
	EvadeCheckNonLongModeRipOverflow uint32 = 0x00000010 // TRANSPARENT_EVADE_CHECK_NON_LONG_MODE_RIP_OVERFLOW

	// EvadeMaskAll is the bitwise-OR of all evade-mask bits, matching
	// TRANSPARENT_EVADE_MASK_ALL.
	EvadeMaskAll uint32 = EvadeMaskSyscallHook | EvadeMaskCpuid | EvadeMaskMsr | EvadeMaskTrapFlag | EvadeCheckNonLongModeRipOverflow

	// EvadeMaskDefault matches TRANSPARENT_EVADE_MASK_DEFAULT (== EvadeMaskAll).
	EvadeMaskDefault uint32 = EvadeMaskAll
)

// Kernel status codes returned in DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE.
// .KernelStatus. Mirror include/SDK/headers/ErrorCodes.h.
const (
	// KernelStatusSuccess is DEBUGGER_OPERATION_WAS_SUCCESSFUL (0xFFFFFFFF).
	KernelStatusSuccess uint32 = 0xFFFFFFFF
	// KernelStatusUnableToHideOrUnhideDebugger is
	// DEBUGGER_ERROR_UNABLE_TO_HIDE_OR_UNHIDE_DEBUGGER (0xc0000009).
	KernelStatusUnableToHideOrUnhideDebugger uint32 = 0xc0000009
)

// TestCount is the number of samples taken per measurement, matching the C++
// #define TestCount 1000 in transparency.h.
const TestCount = 1000

// ----------------------------------------------------------------------------
// Public types
// ----------------------------------------------------------------------------

// Output abstracts command output. CLI implementations write to stdout, GUI
// implementations to a widget, MCP implementations to a JSON channel. This
// mirrors the API design convention used across go-libhyperdbg so the
// transparency module is reusable by all three front-ends.
type Output interface {
	Printf(format string, args ...any) error
}

// nopOutput discards all output. Used as the default when the caller passes a
// nil Output to NewTransparency.
type nopOutput struct{}

func (nopOutput) Printf(format string, args ...any) error { return nil }

// HypervisorResult is the combined outcome of a TransparentHypervisor check.
// It holds the per-test statistics (average / standard deviation / median in
// TSC ticks) and the boolean detection verdicts.
type HypervisorResult struct {
	// HypervisorDetected is true if the rdtsc+cpuid+rdtsc measurement
	// indicates a hypervisor VM-exit on CPUID (average outside (0, 1000)).
	HypervisorDetected bool

	// RdtscpEmulated is true if the rdtsc+rdtsc measurement indicates
	// rdtsc/p emulation (average outside (0, 750)).
	RdtscpEmulated bool

	// CPUID test statistics ( TransparentModeCpuidTimeStampCounter ).
	CpuidAvg    uint64
	CpuidStddev uint64
	CpuidMedian uint64

	// RDTSC test statistics ( TransparentModeRdtscEmulationDetection ).
	RdtscAvg    uint64
	RdtscStddev uint64
	RdtscMedian uint64
}

// SyscallNumberResolver resolves a Windows native syscall number by name (e.g.
// "NtQuerySystemInformation"). It is used by HideDebuggerEx when the
// EvadeMaskSyscallHook bit is set, mirroring C++ CommandHideFillSystemCalls
// (which uses PeGetSyscallNumber). Callers that want syscall-hook transparency
// must supply a resolver — typically backed by the pe-parser module — since
// the transparency package itself has no PE dependency.
//
// Returning an error for any name aborts the hide request, matching the C++
// behaviour where a failed PeGetSyscallNumber only logs a warning but the
// overall request still proceeds (the kernel treats a zero syscall number as
// "do not hook this syscall"). To replicate that lax behaviour, return 0 with
// a nil error for unresolved names instead.
type SyscallNumberResolver func(name string) (uint32, error)

// Transparency holds anti-detection state. It is safe for concurrent use:
// all methods take a sync.Mutex when touching shared state.
//
// A nil *Transparency is never valid; always construct one with
// NewTransparency.
type Transparency struct {
	dev    *comm.Device
	output Output

	mu sync.Mutex // guards rng and any future mutable state

	rng *GaussianRng // for adding Gaussian noise to simulated CPUID/RDTSC values

	// syscallResolver, if non-nil, is consulted by HideDebuggerEx to fill the
	// SYSTEM_CALL_NUMBERS_INFORMATION block when EvadeMaskSyscallHook is set.
	// If nil, HideDebuggerEx returns an error for any request that includes
	// the syscall-hook bit.
	syscallResolver SyscallNumberResolver
}

// NewTransparency creates a Transparency instance bound to the given device
// and output.
//
//   - dev is used to send IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_
//     DEBUGGER. It may be nil for users that only want the measurement helpers
//     (RdtscDiffVmexit, CheckHypervisorPresence, …); the IOCTL methods will
//     return an error in that case.
//   - output receives human-readable status messages (matching the C++
//     ShowMessages calls). A nil Output is replaced with a no-op writer.
//
// The GaussianRng is constructed with mean=0, stddev=1 — the same defaults the
// C++ code uses for noise generation. Callers can replace it via SetRng if
// they need different statistics for the simulation helpers.
func NewTransparency(dev *comm.Device, output Output) *Transparency {
	if output == nil {
		output = nopOutput{}
	}
	return &Transparency{
		dev:    dev,
		output: output,
		rng:    NewGaussianRng(0, 1),
	}
}

// SetRng replaces the Gaussian noise generator used by the simulation helpers
// (SimulateCpuidResult, SimulateRdtscValue). This is intended for tests that
// need deterministic noise. Most callers should leave the default (mean=0,
// stddev=1, MSVCRT seed 1) which matches the C++ Randn(0, 1) baseline.
func (t *Transparency) SetRng(rng *GaussianRng) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.rng = rng
}

// SetSyscallResolver installs a syscall-number resolver for use by
// HideDebuggerEx when the EvadeMaskSyscallHook bit is set. Pass nil to
// disable syscall-hook transparency.
func (t *Transparency) SetSyscallResolver(r SyscallNumberResolver) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.syscallResolver = r
}

// ----------------------------------------------------------------------------
// Measurement helpers — mirror C++ transparency.cpp.
// ----------------------------------------------------------------------------

// RdtscDiffVmexit measures the cycle-count difference of an
// rdtsc → cpuid(leaf 0) → rdtsc sequence. Mirrors C++
// TransparentModeRdtscDiffVmexit.
//
// When a hypervisor is present, CPUID forces a VM-exit, inflating the delta to
// thousands of cycles. Native CPUID is ~hundreds of cycles.
func (t *Transparency) RdtscDiffVmexit() uint64 {
	ret := rdtsc()
	// C++: CpuCpuId(CpuidResult, 0); — forces VM-exit when hypervisor present.
	_ = cpuidLeaf0()
	ret2 := rdtsc()
	return ret2 - ret
}

// RdtscVmexitTracing measures the cycle-count difference of an
// rdtsc → rdtsc sequence. Mirrors C++ TransparentModeRdtscVmexitTracing.
//
// Used to detect rdtsc/p emulation: a hypervisor that intercepts rdtsc
// inflates this delta noticeably compared to native execution.
func (t *Transparency) RdtscVmexitTracing() uint64 {
	ret := rdtsc()
	ret2 := rdtsc()
	return ret2 - ret
}

// CpuidTimeStampCounter runs TestCount iterations of RdtscDiffVmexit and
// computes the average, standard deviation (with +5 for variance, per C++),
// and median via GaussianGenerateRandom. Mirrors C++ TransparentMode-
// CpuidTimeStampCounter.
//
// Returns detected=true when the simple arithmetic mean is outside (0, 1000)
// — i.e. a hypervisor VM-exit on CPUID is likely.
func (t *Transparency) CpuidTimeStampCounter() (detected bool, avg, stddev, median uint64) {
	results := make([]float64, TestCount)
	var sum float64
	for i := 0; i < TestCount; i++ {
		m := float64(t.RdtscDiffVmexit())
		results[i] = m
		sum += m
	}
	avg, stddev, median = GaussianGenerateRandom(results)
	avgF := sum / TestCount
	// C++: return (Avg < 1000 && Avg > 0) ? FALSE : TRUE;
	detected = !(avgF < 1000 && avgF > 0)
	return
}

// RdtscEmulationDetection runs TestCount iterations of RdtscVmexitTracing and
// computes the average, standard deviation (with +5), and median via
// GaussianGenerateRandom. Mirrors C++ TransparentModeRdtscEmulationDetection.
//
// Returns detected=true when the simple arithmetic mean is outside (0, 750) —
// i.e. rdtsc/p emulation is likely.
func (t *Transparency) RdtscEmulationDetection() (detected bool, avg, stddev, median uint64) {
	results := make([]float64, TestCount)
	var sum float64
	for i := 0; i < TestCount; i++ {
		m := float64(t.RdtscVmexitTracing())
		results[i] = m
		sum += m
	}
	avg, stddev, median = GaussianGenerateRandom(results)
	avgF := sum / TestCount
	// C++: return (Avg < 750 && Avg > 0) ? FALSE : TRUE;
	detected = !(avgF < 750 && avgF > 0)
	return
}

// CheckHypervisorPresence runs CpuidTimeStampCounter and prints the verdict.
// Mirrors C++ TransparentModeCheckHypervisorPresence.
//
// ctx is checked for cancellation before the (CPU-bound) measurement loop
// starts; the loop itself runs for ~TestCount iterations and is not
// interruptible mid-flight.
func (t *Transparency) CheckHypervisorPresence(ctx context.Context) (detected bool, avg, stddev, median uint64, err error) {
	if err = ctx.Err(); err != nil {
		return
	}
	detected, avg, stddev, median = t.CpuidTimeStampCounter()
	if detected {
		t.output.Printf("hypervisor detected\n")
	} else {
		t.output.Printf("hypervisor not detected\n")
	}
	return
}

// CheckRdtscpVmexit runs RdtscEmulationDetection and prints the verdict.
// Mirrors C++ TransparentModeCheckRdtscpVmexit.
func (t *Transparency) CheckRdtscpVmexit(ctx context.Context) (detected bool, avg, stddev, median uint64, err error) {
	if err = ctx.Err(); err != nil {
		return
	}
	detected, avg, stddev, median = t.RdtscEmulationDetection()
	if detected {
		t.output.Printf("rdtsc/p emulation detected\n")
	} else {
		t.output.Printf("rdtsc/p emulation not detected\n")
	}
	return
}

// ----------------------------------------------------------------------------
// Combined entry point
// ----------------------------------------------------------------------------

// TransparentHypervisor is the main anti-detection entry point. It runs both
// the CPUID-based hypervisor-presence check and the rdtsc/p-emulation check,
// returning a single combined result. It is the Go equivalent of calling the
// C++ TransparentModeCheckHypervisorPresence followed by
// TransparentModeCheckRdtscpVmexit.
//
// This method is safe for concurrent use, but the underlying measurement
// helpers read the TSC and execute CPUID directly; running them in parallel
// from multiple goroutines will skew the timings. Callers that want accurate
// measurements should serialise calls to TransparentHypervisor (and to the
// individual Check* methods).
func (t *Transparency) TransparentHypervisor(ctx context.Context) (HypervisorResult, error) {
	var r HypervisorResult
	var err error
	r.HypervisorDetected, r.CpuidAvg, r.CpuidStddev, r.CpuidMedian, err = t.CheckHypervisorPresence(ctx)
	if err != nil {
		return r, err
	}
	r.RdtscpEmulated, r.RdtscAvg, r.RdtscStddev, r.RdtscMedian, err = t.CheckRdtscpVmexit(ctx)
	return r, err
}

// ----------------------------------------------------------------------------
// IOCTL helpers — mirror C++ hide.cpp / unhide.cpp.
// ----------------------------------------------------------------------------

// HideDebugger enables transparent mode for the given process ID. Mirrors C++
// HyperDbgEnableTransparentMode(ProcessId, NULL, TRUE) with the default evade
// mask (TRANSPARENT_EVADE_MASK_DEFAULT).
func (t *Transparency) HideDebugger(ctx context.Context, processID uint32) error {
	return t.HideDebuggerEx(ctx, processID, "", true, EvadeMaskDefault)
}

// HideDebuggerByName enables transparent mode for processes matching the
// given name. Mirrors C++ HyperDbgEnableTransparentMode(NULL, name, FALSE)
// with the default evade mask.
func (t *Transparency) HideDebuggerByName(ctx context.Context, processName string) error {
	return t.HideDebuggerEx(ctx, 0, processName, false, EvadeMaskDefault)
}

// HideDebuggerEx enables transparent mode with full parameters. Mirrors C++
// HyperDbgEnableTransparentModeEx(ProcessId, ProcessName, IsProcessId,
// EvadeMask).
//
//   - When isProcessID is true, processID identifies the target by PID and
//     processName is ignored.
//   - When isProcessID is false, processName identifies the target by image
//     name (case-sensitive, NUL-terminated). The name is appended to the IOCTL
//     input buffer past the struct, exactly as in C++ hide.cpp.
//   - evadeMask selects which transparency features to enable. A zero mask is
//     replaced with EvadeMaskDefault.
//
// On success the kernel reports KernelStatusSuccess and the method returns
// nil. On failure the method returns an error wrapping the kernel status.
func (t *Transparency) HideDebuggerEx(ctx context.Context, processID uint32, processName string, isProcessID bool, evadeMask uint32) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if t.dev == nil {
		return errors.New("transparency: nil device; cannot send hide IOCTL")
	}

	if evadeMask == 0 {
		evadeMask = EvadeMaskDefault
	}
	// C++: if ((EffectiveEvadeMask & ~TRANSPARENT_EVADE_MASK_ALL) != 0)
	if evadeMask&^EvadeMaskAll != 0 {
		return fmt.Errorf("transparency: unknown evade-mask bits 0x%x", evadeMask)
	}

	// Build the request struct. Its Go layout (types.DEBUGGER_HIDE_AND_
	// TRANSPARENT_DEBUGGER_MODE) mirrors the C ABI 1:1 — IsHide,
	// TrueIfProcessIdAndFalseIfProcessName, [2]byte padding, ProcId,
	// LengthOfProcessName, SystemCallNumbersInformation, KernelStatus,
	// EvadeMask.
	var req types.DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE
	req.IsHide = true
	req.TrueIfProcessIdAndFalseIfProcessName = isProcessID
	req.EvadeMask = evadeMask

	// If the syscall-hook bit is set, fill the SYSTEM_CALL_NUMBERS_INFORMATION
	// block. C++ hide.cpp does this via CommandHideFillSystemCalls which calls
	// PeGetSyscallNumber for each Nt* routine. We delegate to the configured
	// SyscallNumberResolver.
	if evadeMask&EvadeMaskSyscallHook != 0 {
		t.mu.Lock()
		resolver := t.syscallResolver
		t.mu.Unlock()
		if resolver == nil {
			return errors.New("transparency: EvadeMaskSyscallHook set but no SyscallNumberResolver configured")
		}
		if err := fillSystemCalls(&req.SystemCallNumbersInformation, resolver, t.output); err != nil {
			return fmt.Errorf("transparency: failed to resolve syscall numbers: %w", err)
		}
	}

	structSize := uint32(unsafe.Sizeof(req))

	var (
		inSize uint32
		inBuf  []byte
	)
	if isProcessID {
		// C++: HideRequest.ProcId = (UINT32)ProcessId;
		//      RequestBufferSize  = sizeof(DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE);
		req.ProcId = processID
		inSize = structSize
		inBuf = structAsBytes(unsafe.Pointer(&req), structSize)
	} else {
		// C++: HideRequest.LengthOfProcessName = (UINT32)strlen(ProcessName) + 1;
		//      RequestBufferSize = sizeof(...) + HideRequest.LengthOfProcessName;
		if len(processName) == 0 {
			return errors.New("transparency: empty process name")
		}
		nameLen := uint32(len(processName) + 1) // include NUL terminator
		req.LengthOfProcessName = nameLen
		inSize = structSize + nameLen
		inBuf = make([]byte, inSize)
		copy(inBuf[:structSize], structAsBytes(unsafe.Pointer(&req), structSize))
		copy(inBuf[structSize:structSize+uint32(len(processName))], processName)
		// inBuf[structSize + len(processName)] is already 0 (NUL) from make.
	}

	// Output buffer is the struct itself; the driver writes KernelStatus.
	// C++: Status = PlatformDeviceIoControl(
	//         g_DeviceHandle,
	//         IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER,
	//         FinalRequestBuffer, RequestBufferSize,
	//         FinalRequestBuffer, SIZEOF_DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE,
	//         &ReturnedLength, NULL);
	outBuf := make([]byte, structSize)
	if _, err := t.dev.Ioctl(ctx,
		comm.IOCTL_CODE_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER,
		inBuf, outBuf); err != nil {
		return fmt.Errorf("hide IOCTL failed: %w", err)
	}

	// Parse the kernel status from the output buffer.
	var resp types.DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE
	bytesIntoStruct(unsafe.Pointer(&resp), outBuf)

	switch resp.KernelStatus {
	case KernelStatusSuccess:
		t.output.Printf("transparent debugging successfully enabled :)\n")
		return nil
	case KernelStatusUnableToHideOrUnhideDebugger:
		t.output.Printf("unable to hide the debugger (transparent-debugging) :(\n")
		return fmt.Errorf("transparency: kernel unable to hide debugger (status 0x%x)", resp.KernelStatus)
	default:
		t.output.Printf("unknown error occurred :(\n")
		return fmt.Errorf("transparency: unknown kernel status 0x%x", resp.KernelStatus)
	}
}

// UnhideDebugger disables transparent mode. Mirrors C++ HyperDbgDisable-
// TransparentMode: sends the same IOCTL with IsHide=FALSE and an empty
// request, the kernel reverts all transparency measures.
func (t *Transparency) UnhideDebugger(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if t.dev == nil {
		return errors.New("transparency: nil device; cannot send unhide IOCTL")
	}

	// C++: DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE UnhideRequest = {0};
	//      UnhideRequest.IsHide = FALSE;
	var req types.DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE
	// IsHide stays false (zero value).

	structSize := uint32(unsafe.Sizeof(req))
	inBuf := structAsBytes(unsafe.Pointer(&req), structSize)
	outBuf := make([]byte, structSize)

	// C++: same IOCTL, input/output both SIZEOF_DEBUGGER_HIDE_AND_TRANSPARENT_
	// DEBUGGER_MODE.
	if _, err := t.dev.Ioctl(ctx,
		comm.IOCTL_CODE_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER,
		inBuf, outBuf); err != nil {
		return fmt.Errorf("unhide IOCTL failed: %w", err)
	}

	var resp types.DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE
	bytesIntoStruct(unsafe.Pointer(&resp), outBuf)

	switch resp.KernelStatus {
	case KernelStatusSuccess:
		t.output.Printf("transparent debugging successfully disabled :)\n")
		return nil
	default:
		// C++: ShowErrorMessage(UnhideRequest.KernelStatus);
		return fmt.Errorf("transparency: unhide failed with kernel status 0x%x", resp.KernelStatus)
	}
}

// fillSystemCalls populates the SYSTEM_CALL_NUMBERS_INFORMATION block by
// resolving each Nt* syscall number via the given resolver. Mirrors C++
// CommandHideFillSystemCalls in hide.cpp, which uses PeGetSyscallNumber.
//
// Unlike the C++ version (which only logs warnings on individual failures
// and continues), this Go port returns an error on the first resolver
// failure — callers that want the lax C++ behaviour should return (0, nil)
// from their resolver for unresolved names.
func fillSystemCalls(info *types.SYSTEM_CALL_NUMBERS_INFORMATION, resolve SyscallNumberResolver, out Output) error {
	type slot struct {
		name string
		dst  *uint32
	}
	slots := []slot{
		{"NtQuerySystemInformation", &info.SysNtQuerySystemInformation},
		{"NtQuerySystemInformationEx", &info.SysNtQuerySystemInformationEx},
		{"NtSystemDebugControl", &info.SysNtSystemDebugControl},
		{"NtQueryAttributesFile", &info.SysNtQueryAttributesFile},
		{"NtOpenDirectoryObject", &info.SysNtOpenDirectoryObject},
		{"NtQueryDirectoryObject", &info.SysNtQueryDirectoryObject},
		{"NtQueryInformationProcess", &info.SysNtQueryInformationProcess},
		{"NtSetInformationProcess", &info.SysNtSetInformationProcess},
		{"NtQueryInformationThread", &info.SysNtQueryInformationThread},
		{"NtSetInformationThread", &info.SysNtSetInformationThread},
		{"NtOpenFile", &info.SysNtOpenFile},
		{"NtOpenKey", &info.SysNtOpenKey},
		{"NtOpenKeyEx", &info.SysNtOpenKeyEx},
		{"NtQueryValueKey", &info.SysNtQueryValueKey},
		{"NtEnumerateKey", &info.SysNtEnumerateKey},
	}
	for _, s := range slots {
		n, err := resolve(s.name)
		if err != nil {
			return fmt.Errorf("resolve %s: %w", s.name, err)
		}
		*s.dst = n
		if n == 0 {
			out.Printf("warning, failed to get %s syscall number for transparent-mode\n", s.name)
		}
	}
	return nil
}

// ----------------------------------------------------------------------------
// CPUID / RDTSC simulation helpers (kernel-side emulation mirrors).
// ----------------------------------------------------------------------------

// SimulateCpuidResult adds Gaussian noise to a measured CPUID timing, mirroring
// what the kernel-side transparent-mode hook does once HideDebugger has been
// called: the kernel intercepts CPUID, fakes a "no hypervisor" result, and
// adjusts the TSC delta by sampling from a Gaussian distribution whose mean
// and stddev were computed from the pre-hide measurement (see
// CpuidTimeStampCounter).
//
// mean and stddev should come from a prior CpuidTimeStampCounter call
// (avg and stddev respectively). The returned value is the simulated CPUID
// timing in TSC ticks, cast to uint64 to match the (UINT64)Randn(...) pattern
// in the C++ source comments.
//
// The default GaussianRng is configured with mu=0, sigma=1 (a standard
// normal). This method draws a standard-normal sample and applies the linear
// transformation N(mean, stddev) = mean + stddev * N(0, 1) to produce the
// requested distribution without mutating the rng's configured parameters or
// copying its mutex-guarded state.
func (t *Transparency) SimulateCpuidResult(mean, stddev float64) uint64 {
	t.mu.Lock()
	rng := t.rng
	t.mu.Unlock()
	// rng is *GaussianRng (pointer, not a copy). Next() returns a sample from
	// N(rng.mu, rng.sigma) == N(0, 1) by default; transform to N(mean, stddev).
	sample := rng.Next()
	return uint64(mean + stddev*sample)
}

// SimulateRdtscValue adds Gaussian noise to a measured rdtsc+rdtsc timing,
// mirroring the kernel-side rdtsc/p emulation. As with SimulateCpuidResult,
// mean and stddev should come from a prior RdtscEmulationDetection call.
func (t *Transparency) SimulateRdtscValue(mean, stddev float64) uint64 {
	return t.SimulateCpuidResult(mean, stddev) // same noise model
}

// NoiseSample returns a single Gaussian sample from the configured rng, useful
// for tests that want to inspect the noise distribution directly.
func (t *Transparency) NoiseSample() float64 {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.rng.Next()
}

// ----------------------------------------------------------------------------
// unsafe helpers — struct ↔ []byte without allocation.
// ----------------------------------------------------------------------------

// structAsBytes returns a byte slice that aliases the struct's memory. The
// slice is only valid for the lifetime of the struct pointer; callers must
// not retain it. Used to pass a struct to comm.Device.Ioctl without copying.
func structAsBytes(p unsafe.Pointer, n uint32) []byte {
	if n == 0 {
		return nil
	}
	return unsafe.Slice((*byte)(p), n)
}

// bytesIntoStruct copies a byte slice into a struct. The slice length must be
// >= the struct size; extra bytes are ignored. Used to parse the IOCTL output
// buffer back into a struct.
func bytesIntoStruct(p unsafe.Pointer, b []byte) {
	if len(b) == 0 {
		return
	}
	dst := unsafe.Slice((*byte)(p), len(b))
	copy(dst, b)
}
