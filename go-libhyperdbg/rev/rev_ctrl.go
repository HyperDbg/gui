// Package rev implements the reversing-machine service request. The C++
// counterpart is libhyperdbg/code/rev/rev-ctrl.cpp; it owns:
//   - RevRequestService(RevRequest) — sends an
//     IOCTL_REQUEST_REV_MACHINE_SERVICE to the kernel with a
//     REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST payload and reports the
//     kernel's status back to the user.
//
// The reversing machine reconstructs memory regions (e.g. unpack Themida-
// packed code) by triggering the driver to walk page tables and snapshot
// the target process's address space. The user-mode side is a thin IOCTL
// wrapper; the heavy lifting happens in the kernel.
//
// In the Go rewrite the global state from the C side (g_IsVmmModuleLoaded,
// g_DeviceHandle) is owned by the Ctrl struct, which holds a reference to
// the *core.Debugger (for the device handle) and an Output sink.
//
// Lifecycle:
//
//	r := rev.New(coreDebugger, out)
//	req := rev.ReconstructMemoryRequest{ProcessId: pid, Size: 0x1000,
//	    Mode: rev.ModeReconstruct, Type: rev.TypeMemory}
//	_ = r.RequestService(ctx, &req)
package rev

import (
	"context"
	"fmt"
	"sync"
	"unsafe"

	"github.com/ddkwork/hyperdbgsdk"
	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
	"github.com/hyperdbg/go-libhyperdbg/debugger/core"
)

// Output abstracts message output so CLI/GUI/MCP can all consume the same
// Ctrl instance. It mirrors commands.Output / app.Output but is declared
// locally to keep the rev package free of import cycles.
type Output interface {
	Printf(format string, args ...any) error
}

// DebuggerOperationWasSuccessful mirrors DEBUGGER_OPERATION_WAS_SUCCESSFUL.
const DebuggerOperationWasSuccessful uint32 = 0xFFFFFFFF

// ReconstructMemoryMode mirrors REVERSING_MACHINE_RECONSTRUCT_MEMORY_MODE.
// The concrete values are defined in types/sdk.go; we re-export the type
// alias here for caller convenience.
type ReconstructMemoryMode = hyperdbgsdk.REVERSING_MACHINE_RECONSTRUCT_MEMORY_MODE

// ReconstructMemoryType mirrors REVERSING_MACHINE_RECONSTRUCT_MEMORY_TYPE.
type ReconstructMemoryType = hyperdbgsdk.REVERSING_MACHINE_RECONSTRUCT_MEMORY_TYPE

// ReconstructMemoryRequest mirrors REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST.
// It is the payload sent to the kernel via IOCTL_REQUEST_REV_MACHINE_SERVICE.
type ReconstructMemoryRequest = hyperdbgsdk.REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST

// Ctrl owns the reversing-machine service state. It is bound to one
// *core.Debugger (for the device handle) and one Output sink.
//
// The zero value is not usable; use New.
type Ctrl struct {
	mu  sync.Mutex
	dev *core.Debugger
	out Output
}

// New constructs a Ctrl instance bound to the given *core.Debugger. The
// debugger must already have the VMM module loaded (call app.App.LoadVMM
// before invoking any Ctrl method).
func New(d *core.Debugger, out Output) *Ctrl {
	if out == nil {
		out = discardOutput{}
	}
	return &Ctrl{dev: d, out: out}
}

// RequestService mirrors RevRequestService. It sends the request to the
// kernel via IOCTL_REQUEST_REV_MACHINE_SERVICE and reports the result. The
// request struct is updated in place with the kernel's status
// (KernelStatus field).
//
// Returns nil if the kernel reports DEBUGGER_OPERATION_WAS_SUCCESSFUL;
// otherwise returns an error describing the status code.
func (c *Ctrl) RequestService(ctx context.Context, req *ReconstructMemoryRequest) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if req == nil {
		return fmt.Errorf("rev.RequestService: nil request")
	}
	dev, err := c.device()
	if err != nil {
		return err
	}
	buf := asBytes(req)
	if buf == nil {
		return fmt.Errorf("rev.RequestService: cannot serialise request")
	}
	if _, err := dev.Ioctl(ctx, comm.IOCTL_CODE_REQUEST_REV_MACHINE_SERVICE, buf, buf); err != nil {
		return fmt.Errorf("rev.RequestService: IOCTL failed: %w", err)
	}
	if req.KernelStatus != DebuggerOperationWasSuccessful {
		return fmt.Errorf("rev.RequestService: kernel returned status 0x%x", req.KernelStatus)
	}
	c.out.Printf("the reversing machine service request was successful!\n")
	return nil
}

// device returns the device handle from the bound *core.Debugger. It returns
// an error if the debugger is not connected or does not yet expose its
// device handle.
func (c *Ctrl) device() (*comm.Device, error) {
	c.mu.Lock()
	d := c.dev
	c.mu.Unlock()
	if d == nil {
		return nil, fmt.Errorf("rev: no *core.Debugger bound")
	}
	// TODO(Phase C.3): core.Debugger.Device() *comm.Device
	return nil, fmt.Errorf("rev: core.Debugger.Device() not yet exposed (Phase C.3)")
}

// asBytes returns a byte slice aliasing the memory of req for the duration
// of the call.
func asBytes(req *ReconstructMemoryRequest) []byte {
	const sz = int(unsafe.Sizeof(*req))
	return unsafe.Slice((*byte)(unsafe.Pointer(req)), sz)
}

// discardOutput is the default Output when the caller passes nil.
type discardOutput struct{}

func (discardOutput) Printf(format string, args ...any) error {
	_ = args
	_ = format
	return nil
}
