// Package objects implements process/thread queries. The C++ counterpart is
// libhyperdbg/code/objects/objects.cpp; it owns:
//   - ObjectShowProcessesOrThreadDetails(IsProcess) — queries the current
//     process or thread details via IOCTL_QUERY_CURRENT_PROCESS /
//     IOCTL_QUERY_CURRENT_THREAD
//   - ObjectShowProcessesOrThreadList(IsProcess, ...) — queries the count of
//     active processes/threads via IOCTL_QUERY_COUNT_OF_ACTIVE_PROCESSES_OR_THREADS,
//     then queries the list itself via IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES
//
// In the Go rewrite the global state from the C side (g_IsKdModuleLoaded,
// g_DeviceHandle) is owned by the Objects struct, which holds a reference to
// the *core.Debugger (for the device handle) and an Output sink. The struct
// is safe for concurrent use; the IOCTLs are serialised by the Debugger's
// own mutex.
//
// Lifecycle:
//
//	obj := objects.New(coreDebugger, out)
//	_ = obj.ShowProcessDetails()
//	procs, _ := obj.ListProcesses()
//	for _, p := range procs { out.Printf("pid=%x cr3=%016x image=%s\n", p.ProcessId, p.Cr3, p.Image) }
package objects

import (
	"encoding/binary"
	"fmt"
	"sync"
	"unsafe"

	"github.com/ddkwork/hyperdbgsdk"
	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
	"github.com/hyperdbg/go-libhyperdbg/debugger/core"
)

// Output abstracts message output so CLI/GUI/MCP can all consume the same
// Objects instance. It mirrors commands.Output / app.Output but is declared
// locally to keep the objects package free of import cycles.
type Output interface {
	Printf(format string, args ...any) error
}

// DebuggerOperationWasSuccessful mirrors DEBUGGER_OPERATION_WAS_SUCCESSFUL.
const DebuggerOperationWasSuccessful uint32 = 0xFFFFFFFF

// Objects owns the process/thread query state. It is bound to one
// *core.Debugger (for the device handle) and one Output sink.
//
// The zero value is not usable; use New.
type Objects struct {
	mu  sync.Mutex
	dev *core.Debugger
	out Output
}

// New constructs an Objects instance bound to the given *core.Debugger. The
// debugger must already be connected (call core.Debugger.Connect or
// app.App.LoadVMM before invoking any Objects method).
func New(d *core.Debugger, out Output) *Objects {
	if out == nil {
		out = discardOutput{}
	}
	return &Objects{dev: d, out: out}
}

// ProcessDetails mirrors the subset of DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET
// that ObjectShowProcessesOrThreadDetails prints: process id, EPROCESS and
// the 16-byte image file name.
type ProcessDetails struct {
	ProcessId   uint32
	Eprocess    uint64
	ProcessName [16]byte
}

// ThreadDetails mirrors the subset of DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET
// that ObjectShowProcessesOrThreadDetails prints: thread id, owning pid,
// ETHREAD, EPROCESS and the 16-byte image file name.
type ThreadDetails struct {
	ThreadId    uint32
	ProcessId   uint32
	Ethread     uint64
	Eprocess    uint64
	ProcessName [16]byte
}

// ProcessListEntry mirrors DEBUGGEE_PROCESS_LIST_DETAILS_ENTRY.
type ProcessListEntry = hyperdbgsdk.DEBUGGEE_PROCESS_LIST_DETAILS_ENTRY

// ThreadListEntry mirrors DEBUGGEE_THREAD_LIST_DETAILS_ENTRY.
type ThreadListEntry = hyperdbgsdk.DEBUGGEE_THREAD_LIST_DETAILS_ENTRY

// ShowProcessDetails mirrors ObjectShowProcessesOrThreadDetails(TRUE). It
// queries the current process via IOCTL_QUERY_CURRENT_PROCESS and prints
// the result to out. Returns the parsed details on success.
func (o *Objects) ShowProcessDetails() (ProcessDetails, error) {
	var pkt hyperdbgsdk.DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET
	buf := make([]byte, int(unsafe.Sizeof(pkt)))
	pktBytes := asBytes(&pkt)
	copy(buf, pktBytes)

	dev, err := o.device()
	if err != nil {
		return ProcessDetails{}, err
	}
	n, err := dev.Ioctl(comm.IOCTL_CODE_QUERY_CURRENT_PROCESS, buf, buf)
	if err != nil {
		return ProcessDetails{}, fmt.Errorf("ShowProcessDetails: IOCTL_QUERY_CURRENT_PROCESS failed: %w", err)
	}
	if n < uint32(len(buf)) {
		return ProcessDetails{}, fmt.Errorf("ShowProcessDetails: short IOCTL response (%d < %d)", n, len(buf))
	}
	copy(pktBytes, buf)
	if pkt.Result != DebuggerOperationWasSuccessful {
		return ProcessDetails{}, fmt.Errorf("ShowProcessDetails: kernel returned status 0x%x", pkt.Result)
	}
	d := ProcessDetails{
		ProcessId:   pkt.ProcessId,
		Eprocess:    pkt.Process,
		ProcessName: pkt.ProcessName,
	}
	o.out.Printf("process id: %x\nprocess (_EPROCESS): %s\nprocess name (16-Byte): %s\n",
		d.ProcessId, separateTo64BitValue(d.Eprocess), cstr(d.ProcessName[:]))
	return d, nil
}

// ShowThreadDetails mirrors ObjectShowProcessesOrThreadDetails(FALSE). It
// queries the current thread via IOCTL_QUERY_CURRENT_THREAD and prints the
// result to out. Returns the parsed details on success.
func (o *Objects) ShowThreadDetails() (ThreadDetails, error) {
	var pkt hyperdbgsdk.DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET
	buf := make([]byte, int(unsafe.Sizeof(pkt)))
	pktBytes := asBytes(&pkt)
	copy(buf, pktBytes)

	dev, err := o.device()
	if err != nil {
		return ThreadDetails{}, err
	}
	n, err := dev.Ioctl(comm.IOCTL_CODE_QUERY_CURRENT_THREAD, buf, buf)
	if err != nil {
		return ThreadDetails{}, fmt.Errorf("ShowThreadDetails: IOCTL_QUERY_CURRENT_THREAD failed: %w", err)
	}
	if n < uint32(len(buf)) {
		return ThreadDetails{}, fmt.Errorf("ShowThreadDetails: short IOCTL response (%d < %d)", n, len(buf))
	}
	copy(pktBytes, buf)
	if pkt.Result != DebuggerOperationWasSuccessful {
		return ThreadDetails{}, fmt.Errorf("ShowThreadDetails: kernel returned status 0x%x", pkt.Result)
	}
	d := ThreadDetails{
		ThreadId:    pkt.ThreadId,
		ProcessId:   pkt.ProcessId,
		Ethread:     pkt.Thread,
		Eprocess:    pkt.Process,
		ProcessName: pkt.ProcessName,
	}
	o.out.Printf("thread id: %x (pid: %x)\nthread (_ETHREAD): %s\nprocess (_EPROCESS): %s\nprocess name (16-Byte): %s\n",
		d.ThreadId, d.ProcessId,
		separateTo64BitValue(d.Ethread),
		separateTo64BitValue(d.Eprocess),
		cstr(d.ProcessName[:]))
	return d, nil
}

// ListProcesses mirrors ObjectShowProcessesOrThreadList(TRUE, ...). It first
// queries the count of active processes, then queries the list itself.
// Returns the parsed list entries. The caller owns the returned slice.
func (o *Objects) ListProcesses() ([]ProcessListEntry, error) {
	return o.listActive(true)
}

// ListThreads mirrors ObjectShowProcessesOrThreadList(FALSE, ...). It first
// queries the count of active threads, then queries the list itself.
// Returns the parsed list entries. The caller owns the returned slice.
func (o *Objects) ListThreads() ([]ThreadListEntry, error) {
	raw, err := o.listActive(false)
	if err != nil {
		return nil, err
	}
	// Reinterpret the []ProcessListEntry as []ThreadListEntry. Both types
	// have the same size on x64 (40 bytes); the cast is sound because the
	// kernel wrote thread entries into the same buffer. We use unsafe.Slice
	// to avoid an extra copy.
	if len(raw) == 0 {
		return nil, nil
	}
	// Bounds-check sizes: the two structs must have equal size for the
	// in-place reinterpretation to be valid.
	if unsafe.Sizeof(ProcessListEntry{}) != unsafe.Sizeof(ThreadListEntry{}) {
		// Fallback: reinterpret element-by-element via unsafe. This branch
		// is taken only if the upstream types diverge in a future revision;
		// the kernel still writes a flat byte buffer so the raw memory layout
		// is authoritative.
		out := make([]ThreadListEntry, len(raw))
		for i := range raw {
			out[i] = *(*ThreadListEntry)(unsafe.Pointer(&raw[i]))
		}
		return out, nil
	}
	out := unsafe.Slice((*ThreadListEntry)(unsafe.Pointer(&raw[0])), len(raw))
	return out, nil
}

// listActive implements the two-step count+list query used by both
// ListProcesses and ListThreads. isProcess selects the query type.
func (o *Objects) listActive(isProcess bool) ([]ProcessListEntry, error) {
	dev, err := o.device()
	if err != nil {
		return nil, err
	}

	var query hyperdbgsdk.DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS
	if isProcess {
		query.QueryType = hyperdbgsdk.DebuggerQueryActiveProcessesOrThreadsQueryProcessCount
	} else {
		query.QueryType = hyperdbgsdk.DebuggerQueryActiveProcessesOrThreadsQueryThreadCount
	}
	query.QueryAction = hyperdbgsdk.DebuggerQueryActiveProcessesOrThreadsActionShowInstantly

	queryBuf := make([]byte, int(unsafe.Sizeof(query)))
	qBytes := asBytes(&query)
	copy(queryBuf, qBytes)

	if _, err := dev.Ioctl(comm.IOCTL_CODE_QUERY_COUNT_OF_ACTIVE_PROCESSES_OR_THREADS, queryBuf, queryBuf); err != nil {
		return nil, fmt.Errorf("listActive: count IOCTL failed: %w", err)
	}
	copy(qBytes, queryBuf)
	if query.Result != uint64(DebuggerOperationWasSuccessful) {
		return nil, fmt.Errorf("listActive: count IOCTL returned status 0x%x", uint32(query.Result))
	}
	if query.Count == 0 {
		return nil, nil
	}
	// Add 5 slots for new processes/threads created between the two IOCTLs,
	// mirroring the C++ code.
	count := query.Count + 5
	entrySize := uint32(unsafe.Sizeof(ProcessListEntry{}))
	listBufSize := count * entrySize
	listBuf := make([]byte, listBufSize)

	// Reuse the query buffer for the second IOCTL, switching the query type
	// to the list variant.
	if isProcess {
		query.QueryType = hyperdbgsdk.DebuggerQueryActiveProcessesOrThreadsQueryProcessList
	} else {
		query.QueryType = hyperdbgsdk.DebuggerQueryActiveProcessesOrThreadsQueryThreadList
	}
	copy(queryBuf, qBytes)

	if _, err := dev.Ioctl(comm.IOCTL_CODE_GET_LIST_OF_THREADS_AND_PROCESSES, queryBuf, listBuf); err != nil {
		return nil, fmt.Errorf("listActive: list IOCTL failed: %w", err)
	}
	copy(qBytes, queryBuf)
	if query.Result != uint64(DebuggerOperationWasSuccessful) {
		return nil, fmt.Errorf("listActive: list IOCTL returned status 0x%x", uint32(query.Result))
	}

	// Reinterpret listBuf as []ProcessListEntry.
	entries := unsafe.Slice((*ProcessListEntry)(unsafe.Pointer(&listBuf[0])), count)
	out := make([]ProcessListEntry, 0, count)
	for i := uint32(0); i < count; i++ {
		if entries[i].Eprocess == 0 {
			continue
		}
		out = append(out, entries[i])
		o.out.Printf("PROCESS\t%llx\n\tProcess Id: %04x\tDirBase (Kernel Cr3): %016llx\tImage: %s\n\n",
			entries[i].Eprocess, entries[i].ProcessId, entries[i].Cr3, cstr(entries[i].ImageFileName[:]))
	}
	return out, nil
}

// device returns the device handle from the bound *core.Debugger. It returns
// an error if the debugger is not connected.
func (o *Objects) device() (*comm.Device, error) {
	o.mu.Lock()
	d := o.dev
	o.mu.Unlock()
	if d == nil {
		return nil, fmt.Errorf("objects: no *core.Debugger bound")
	}
	// The Debugger does not yet expose its *comm.Device directly; when it
	// does (Phase C.3), this will be `return d.Device(), nil`. For the
	// framework stub we return an error so callers know the path is not
	// wired up yet.
	//
	// TODO(Phase C.3): core.Debugger.Device() *comm.Device
	return nil, fmt.Errorf("objects: core.Debugger.Device() not yet exposed (Phase C.3)")
}

// separateTo64BitValue formats a 64-bit value with a backtick separator
// between the high and low 32-bit halves. Mirrors the C++
// SeparateTo64BitValue helper from common.cpp.
func separateTo64BitValue(v uint64) string {
	return fmt.Sprintf("%08x`%08x", uint32(v>>32), uint32(v))
}

// cstr returns the portion of b up to the first NUL byte, as a string.
// Kernel strings are NUL-terminated C strings.
func cstr(b []byte) string {
	for i, c := range b {
		if c == 0 {
			return string(b[:i])
		}
	}
	return string(b)
}

// asBytes returns a byte slice aliasing the memory of v for the duration of
// the call. It is the standard unsafe.Pointer trick used to (de)serialise
// fixed-size C structs without encoding/binary overhead.
func asBytes(v any) []byte {
	switch p := v.(type) {
	case *hyperdbgsdk.DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET:
		const sz = int(unsafe.Sizeof(*p))
		return unsafe.Slice((*byte)(unsafe.Pointer(p)), sz)
	case *hyperdbgsdk.DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET:
		const sz = int(unsafe.Sizeof(*p))
		return unsafe.Slice((*byte)(unsafe.Pointer(p)), sz)
	case *hyperdbgsdk.DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS:
		const sz = int(unsafe.Sizeof(*p))
		return unsafe.Slice((*byte)(unsafe.Pointer(p)), sz)
	}
	return nil
}

// NativeEndian mirrors binary.LittleEndian; used implicitly by the unsafe
// reinterpretation above (the host is little-endian on amd64).
var NativeEndian = binary.LittleEndian

// discardOutput is the default Output when the caller passes nil.
type discardOutput struct{}

func (discardOutput) Printf(format string, args ...any) error {
	_ = args
	_ = format
	return nil
}
