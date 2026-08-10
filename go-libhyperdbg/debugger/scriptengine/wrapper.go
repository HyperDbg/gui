package scriptengine

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"

	astencoder "github.com/hyperdbg/go-bridge/ast"
	"github.com/hyperdbg/go-bridge/protocol"
	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
	"github.com/hyperdbg/go-libhyperdbg/types"
)

// maxUint32 is the largest value that fits in DEBUGGER_GENERAL_ACTION's
// ScriptBufferSize field (a uint32). The driver's RUN_SCRIPT path rejects
// ScriptBufferSize == 0 with DEBUGGER_ERROR_ACTION_BUFFER_SIZE_IS_ZERO, so
// any non-empty payload above this bound cannot be sent at all.
const maxUint32 = 1<<32 - 1

// Wrapper bridges Go source callbacks and the HyperDbg kernel driver. It
// compiles Go source to binary AST (via go-bridge/ast), builds the
// DEBUGGER_GENERAL_EVENT_DETAIL and DEBUGGER_GENERAL_ACTION structures the
// driver expects, and dispatches them via the
// IOCTL_DEBUGGER_REGISTER_EVENT + IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT IOCTLs.
//
// A Wrapper is the Go counterpart to a small subset of the C++ wrapper at
// HyperDbg/hyperdbg/libhyperdbg/code/debugger/script-engine/script-engine-wrapper.cpp.
// The C++ wrapper drives the legacy .ds bytecode VM (ScriptEngineParse +
// ScriptEngineExecute + SYMBOL_BUFFER); this Wrapper drives the new Go-subset
// interpreter (go-bridge/ast → binary AST → hyperkd/code/go-interp). Both
// paths coexist in the kernel — see HOOK_FLAG_GO_AST in
// hyperkd/code/go-interp/interp_stub.c.
//
// A Wrapper is safe for concurrent use: Compile consults a thread-safe
// AstCache, nextTag is incremented atomically, and the pending-scripts map
// is guarded by a mutex. A single Wrapper instance can serve multiple
// Debugger instances (or a Debugger + an MCP layer) without external locking.
type Wrapper struct {
	cache *AstCache

	// nextTag is the next event tag to hand out. Tags identify event+action
	// pairs in the kernel. Tag 0 is reserved as "no tag", so the first
	// issued tag is 1. Accessed with sync/atomic.
	nextTag uint64

	// mu guards scripts. We use a plain Mutex rather than RWMutex because
	// the only read path (RegisterHook's takeScriptBytes) also deletes.
	mu sync.Mutex

	// scripts holds the script bytes passed to BuildAction so RegisterHook
	// can append them to the action IOCTL payload without taking scriptBytes
	// as an additional parameter (the task's API signature does not allow
	// it). Entries are removed by RegisterHook; callers that abandon an
	// action between BuildAction and RegisterHook will leak the entry
	// until the Wrapper itself is garbage-collected.
	scripts map[*types.DEBUGGER_GENERAL_ACTION][]byte
}

// NewWrapper creates a Wrapper with a fresh AST cache and a tag counter
// starting at 1 (tag 0 is reserved by the driver as "no tag").
func NewWrapper() *Wrapper {
	return &Wrapper{
		cache:   NewAstCache(),
		nextTag: 1,
		scripts: make(map[*types.DEBUGGER_GENERAL_ACTION][]byte),
	}
}

// Cache returns the Wrapper's AST cache (never nil). Callers may Clear it to
// force re-compilation after a go-bridge upgrade, or to reclaim memory.
func (w *Wrapper) Cache() *AstCache { return w.cache }

// Compile encodes callbackSrc to the binary AST wire format defined by
// go-bridge/protocol. The returned []byte is suitable for embedding in a
// DEBUGGER_GENERAL_ACTION payload sent via IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT.
//
// Compile consults the cache first; on a miss it invokes astencoder.Encode
// (go/parser → subset validator → wire-format serialiser) and stores the
// result. The cached payload is copied on return, so callers may freely
// mutate it.
//
// Error classification:
//   - errors.Is(err, ErrSubnetViolation) — callbackSrc failed parsing or
//     uses constructs outside the supported Go subset. The wrapped error is
//     typically a *protocol.ValidationError whose Pos field carries the
//     source location (e.g. "hook.go:42:5").
//   - errors.Is(err, ErrScriptCompile) — unexpected encoder failure (e.g.
//     empty payload returned). Indicates a go-bridge bug, not a user error.
//
// Compile is safe for concurrent use. On simultaneous cache misses for the
// same source, both goroutines run astencoder.Encode; the second Put overwrites
// the first with identical bytes (a benign race).
func (w *Wrapper) Compile(callbackSrc string) ([]byte, error) {
	if cached, ok := w.cache.Get(callbackSrc); ok {
		return cached, nil
	}
	payload, err := astencoder.Encode(callbackSrc)
	if err != nil {
		// go-bridge wraps every user-facing failure as a *protocol.ValidationError
		// (parse errors, subset violations, encoder limitations all share this
		// type). errors.As unwraps any future wrapping, so callers that want
		// to distinguish "user's fault" (subset) from "our fault" (encoder
		// bug) can use errors.Is(err, ErrSubnetViolation).
		var ve *protocol.ValidationError
		if errors.As(err, &ve) {
			return nil, fmt.Errorf("%w: %v", ErrSubnetViolation, err)
		}
		return nil, fmt.Errorf("%w: %v", ErrScriptCompile, err)
	}
	if len(payload) == 0 {
		return nil, fmt.Errorf("%w: encoder returned empty payload", ErrScriptCompile)
	}
	w.cache.Put(callbackSrc, payload)
	return payload, nil
}

// BuildAction constructs a DEBUGGER_GENERAL_ACTION describing how the kernel
// should react when the associated event fires. The script bytes are
// remembered internally by the Wrapper and attached to the IOCTL payload
// automatically by RegisterHook — the caller does not need to splice them
// into the buffer.
//
// Action type rules:
//   - types.RunScript: scriptBytes must be non-empty (the kernel rejects
//     ScriptBufferSize == 0 with DEBUGGER_ERROR_ACTION_BUFFER_SIZE_IS_ZERO).
//   - types.BreakToDebugger: scriptBytes is ignored and may be nil. The
//     action simply breaks into the debugger when the event fires.
//   - types.RunCustomCode: not supported by this Wrapper. Custom-code
//     actions have a different buffer layout (CustomCodeBufferPointer +
//     trailing custom-code bytes, with different driver-side handling) and
//     belong on a dedicated path; see the C libhyperdbg's
//     RegisterActionToEvent for the reference implementation.
//
// The returned *DEBUGGER_GENERAL_ACTION must be passed to RegisterHook to
// release the Wrapper's internal reference to scriptBytes. Failing to call
// RegisterHook leaks the entry until the Wrapper is garbage-collected.
//
// ScriptBufferPointer is left as 0: the kernel computes the script address
// as &action + sizeof(DEBUGGER_GENERAL_ACTION) (see
// hyperkd/code/debugger/core/Debugger.c:3428), so this field is unused on
// the new Go-AST path. It is non-zero only on the legacy .ds-bytecode path
// where the script lives in a separate allocation.
func (w *Wrapper) BuildAction(scriptBytes []byte, actionType types.DEBUGGER_EVENT_ACTION_TYPE_ENUM) (*types.DEBUGGER_GENERAL_ACTION, error) {
	switch actionType {
	case types.RunScript:
		if len(scriptBytes) == 0 {
			return nil, fmt.Errorf("%w: RunScript requires non-empty script bytes", ErrActionBuild)
		}
		if uint64(len(scriptBytes)) > maxUint32 {
			return nil, fmt.Errorf("%w: script buffer %d bytes exceeds uint32 limit (%d)",
				ErrActionBuild, len(scriptBytes), maxUint32)
		}
	case types.BreakToDebugger:
		// No script bytes needed; scriptBytes is ignored.
	case types.RunCustomCode:
		return nil, fmt.Errorf("%w: RunCustomCode not supported by this wrapper (use a dedicated path)", ErrActionBuild)
	default:
		return nil, fmt.Errorf("%w: unknown action type %d (%s)",
			ErrActionBuild, actionType, actionType.String())
	}
	action := &types.DEBUGGER_GENERAL_ACTION{
		ActionType:          actionType,
		ScriptBufferSize:    uint32(len(scriptBytes)),
		ScriptBufferPointer: 0,
	}
	if len(scriptBytes) > 0 {
		w.mu.Lock()
		w.scripts[action] = scriptBytes
		w.mu.Unlock()
	}
	return action, nil
}

// BuildEvent constructs a DEBUGGER_GENERAL_EVENT_DETAIL for the given event
// type and options. A unique Tag is allocated atomically; ProcessId is set
// to 0 (matches the C convention "0 = all processes"); EventStage defaults
// to VmmCallbackCallingStagePreEventEmulation, matching the C libhyperdbg
// default in debugger.cpp:1754.
//
// The caller should populate options.OptionalParam1..6 as appropriate for
// the chosen event type (e.g. for HiddenHookExecDetours, OptionalParam1 is
// the hook target address).
//
// BuildEvent does not currently fail (any VMM_EVENT_TYPE_ENUM value is
// accepted; the kernel performs final validation). ErrEventBuild is
// reserved for future parameter checks.
func (w *Wrapper) BuildEvent(eventType types.VMM_EVENT_TYPE_ENUM, options types.DEBUGGER_EVENT_OPTIONS) (*types.DEBUGGER_GENERAL_EVENT_DETAIL, error) {
	tag := atomic.AddUint64(&w.nextTag, 1) - 1
	return &types.DEBUGGER_GENERAL_EVENT_DETAIL{
		ProcessId:       0, // 0 = all processes
		IsEnabled:       true,
		EventStage:      types.VmmCallbackCallingStagePreEventEmulation,
		HasCustomOutput: false,
		Tag:             tag,
		EventType:       eventType,
		Options:         options,
		CountOfActions:  0,
	}, nil
}

// RegisterHook sends the event and action to the kernel.
//
// Steps:
//  1. Send the event struct via IOCTL_DEBUGGER_REGISTER_EVENT. The driver
//     assigns the event to event.Tag and returns
//     DEBUGGER_EVENT_AND_ACTION_RESULT with IsSuccessful/Error set.
//  2. Attach the action via IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT. The
//     action's EventTag is set to event.Tag (the kernel requires them to
//     match), and the script bytes (if any) are appended to the IOCTL input
//     buffer immediately after the DEBUGGER_GENERAL_ACTION struct — exactly
//     as the C libhyperdbg does in debugger.cpp:2235-2257.
//
// Returns the event tag (event.Tag) on success.
//
// RegisterHook consumes the pending script bytes associated with action
// (see BuildAction); calling RegisterHook twice with the same action is
// an error (the second call will send an empty script buffer).
//
// Context cancellation is checked before each IOCTL; the IOCTLs themselves
// are synchronous (METHOD_BUFFERED) and cannot be interrupted mid-call.
func (w *Wrapper) RegisterHook(ctx context.Context, dev *comm.Device, event *types.DEBUGGER_GENERAL_EVENT_DETAIL, action *types.DEBUGGER_GENERAL_ACTION) (uint64, error) {
	if dev == nil {
		return 0, fmt.Errorf("%w: nil device", ErrIoctlRegister)
	}
	if event == nil {
		return 0, fmt.Errorf("%w: nil event", ErrIoctlRegister)
	}
	if action == nil {
		return 0, fmt.Errorf("%w: nil action", ErrIoctlRegister)
	}

	// 1. Register the event.
	eventBuf, err := structToBytes(event)
	if err != nil {
		return 0, fmt.Errorf("%w: serialise event: %v", ErrIoctlRegister, err)
	}
	resultSize := unsafe.Sizeof(types.DEBUGGER_EVENT_AND_ACTION_RESULT{})
	resultBuf := make([]byte, resultSize)
	if _, err := dev.Ioctl(ctx, comm.IOCTL_CODE_DEBUGGER_REGISTER_EVENT, eventBuf, resultBuf); err != nil {
		return 0, fmt.Errorf("%w: REGISTER_EVENT IOCTL: %v", ErrIoctlRegister, err)
	}
	var result types.DEBUGGER_EVENT_AND_ACTION_RESULT
	if err := bytesToStruct(resultBuf, &result); err != nil {
		return 0, fmt.Errorf("%w: deserialise REGISTER_EVENT result: %v", ErrIoctlRegister, err)
	}
	if !result.IsSuccessful {
		return 0, fmt.Errorf("%w: kernel rejected event (error=%d, see DEBUGGER_ERROR_* in SDK)",
			ErrIoctlRegister, result.Error)
	}

	// 2. Attach the action.
	//    The script bytes (if any) follow the action struct in the same
	//    IOCTL input buffer; the kernel reads them at offset
	//    sizeof(DEBUGGER_GENERAL_ACTION) from the start of the input
	//    buffer (see Debugger.c:3428).
	action.EventTag = event.Tag
	scriptBytes := w.takeScriptBytes(action) // may be nil for BreakToDebugger

	actionBuf, err := structToBytes(action)
	if err != nil {
		return 0, fmt.Errorf("%w: serialise action: %v", ErrIoctlRegister, err)
	}
	payload := make([]byte, len(actionBuf)+len(scriptBytes))
	copy(payload, actionBuf)
	copy(payload[len(actionBuf):], scriptBytes)

	actionResultBuf := make([]byte, resultSize)
	if _, err := dev.Ioctl(ctx, comm.IOCTL_CODE_DEBUGGER_ADD_ACTION_TO_EVENT, payload, actionResultBuf); err != nil {
		return 0, fmt.Errorf("%w: ADD_ACTION_TO_EVENT IOCTL: %v", ErrIoctlRegister, err)
	}
	var actionResult types.DEBUGGER_EVENT_AND_ACTION_RESULT
	if err := bytesToStruct(actionResultBuf, &actionResult); err != nil {
		return 0, fmt.Errorf("%w: deserialise ADD_ACTION_TO_EVENT result: %v", ErrIoctlRegister, err)
	}
	if !actionResult.IsSuccessful {
		return 0, fmt.Errorf("%w: kernel rejected action (error=%d, see DEBUGGER_ERROR_* in SDK)",
			ErrIoctlRegister, actionResult.Error)
	}
	return event.Tag, nil
}

// takeScriptBytes removes and returns the script bytes associated with
// action. Returns nil if no script bytes were registered (e.g., for
// BreakToDebugger actions or if BuildAction was never called with a
// non-empty scriptBytes).
func (w *Wrapper) takeScriptBytes(action *types.DEBUGGER_GENERAL_ACTION) []byte {
	w.mu.Lock()
	defer w.mu.Unlock()
	sb := w.scripts[action]
	delete(w.scripts, action)
	return sb
}

// structToBytes serialises a pointer-to-struct into a byte slice whose
// layout matches the C ABI. The types package preserves field alignment
// via explicit padding fields (e.g. `_ [N]byte`), so the resulting bytes
// can be sent directly to the driver via DeviceIoControl without any
// packing transformation.
//
// Returns a copy: the caller may freely mutate the result without affecting
// the source struct, and vice versa.
func structToBytes(ptr any) ([]byte, error) {
	switch s := ptr.(type) {
	case *types.DEBUGGER_GENERAL_EVENT_DETAIL:
		sz := unsafe.Sizeof(*s)
		src := unsafe.Slice((*byte)(unsafe.Pointer(s)), sz)
		out := make([]byte, sz)
		copy(out, src)
		return out, nil
	case *types.DEBUGGER_GENERAL_ACTION:
		sz := unsafe.Sizeof(*s)
		src := unsafe.Slice((*byte)(unsafe.Pointer(s)), sz)
		out := make([]byte, sz)
		copy(out, src)
		return out, nil
	case *types.DEBUGGER_EVENT_AND_ACTION_RESULT:
		sz := unsafe.Sizeof(*s)
		src := unsafe.Slice((*byte)(unsafe.Pointer(s)), sz)
		out := make([]byte, sz)
		copy(out, src)
		return out, nil
	}
	return nil, fmt.Errorf("structToBytes: unsupported type %T (expected *types.DEBUGGER_GENERAL_EVENT_DETAIL, *types.DEBUGGER_GENERAL_ACTION or *types.DEBUGGER_EVENT_AND_ACTION_RESULT)", ptr)
}

// bytesToStruct deserialises a byte slice into a pointer-to-struct. The
// inverse of structToBytes; same ABI-compatibility notes apply.
func bytesToStruct(buf []byte, ptr any) error {
	switch s := ptr.(type) {
	case *types.DEBUGGER_EVENT_AND_ACTION_RESULT:
		sz := unsafe.Sizeof(*s)
		if uintptr(len(buf)) < sz {
			return fmt.Errorf("bytesToStruct: buffer %d bytes < struct %d bytes", len(buf), sz)
		}
		dst := unsafe.Slice((*byte)(unsafe.Pointer(s)), sz)
		copy(dst, buf[:sz])
		return nil
	}
	return fmt.Errorf("bytesToStruct: unsupported type %T (expected *types.DEBUGGER_EVENT_AND_ACTION_RESULT)", ptr)
}
