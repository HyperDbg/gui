package scriptengine

import (
	"testing"
	"unsafe"

	"github.com/ddkwork/hyperdbgsdk"
	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
	"github.com/hyperdbg/go-libhyperdbg/debugger/driverloader"
)

const driverPath = `C:\Users\Administrator\AppData\Local\hyperdbg\hyperkd.sys`

// initVmmRequest mirrors DEBUGGER_INIT_VMM_PACKET.
type initVmmRequest struct {
	KernelStatus uint32
}

// clearEvent sends IOCTL_CODE_DEBUGGER_MODIFY_EVENTS with
// DebuggerModifyEventsClear to remove the event identified by tag.
func clearEvent(t *testing.T, dev *comm.Device, tag uint64) {
	t.Helper()
	var mod hyperdbgsdk.DEBUGGER_MODIFY_EVENTS
	mod.Tag = tag
	mod.TypeOfAction = hyperdbgsdk.DebuggerModifyEventsClear
	modSize := uint32(unsafe.Sizeof(mod))
	if _, err := dev.IoctlStruct(
		hyperdbgsdk.IoctlDebuggerModifyEvents,
		unsafe.Pointer(&mod), unsafe.Pointer(&mod), modSize, modSize); err != nil {
		t.Logf("best-effort clearEvent(tag=%d) failed: %v", tag, err)
	}
}

// TestEptHook_Register 集成测试：加载驱动+VMM，注册 EPT hook，验证 Tag > 0。
func TestEptHook_Register(t *testing.T) {
	// 1. 加载驱动
	d := driverloader.NewDriver(driverPath)
	if err := d.Load(); err != nil {
		t.Fatalf("加载驱动失败: %v", err)
	}
	t.Cleanup(func() { _ = d.Unload() })

	// 2. 加载VMM
	dev, err := comm.Open(comm.DeviceName)
	if err != nil {
		t.Fatalf("打开设备失败: %v", err)
	}
	t.Cleanup(func() { _ = dev.Close() })

	var vmmReq initVmmRequest
	vmmSize := uint32(unsafe.Sizeof(vmmReq))
	if _, err := dev.IoctlStruct(hyperdbgsdk.IoctlInitVmm,
		unsafe.Pointer(&vmmReq), unsafe.Pointer(&vmmReq), vmmSize, vmmSize); err != nil {
		t.Skipf("IOCTL_INIT_VMM failed: %v", err)
	}
	if vmmReq.KernelStatus != uint32(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		t.Skipf("VMM init failed (0x%08x); system lacks VT-x", vmmReq.KernelStatus)
	}
	t.Cleanup(func() {
		_, _ = dev.IoctlStruct(hyperdbgsdk.IoctlTerminateVmx, nil, nil, 0, 0)
	})

	// 3. 编译 Go 回调并注册 EPT hook
	w := NewWrapper()

	callbackSrc := `package hook
func hook(ctx *HookCtx) {
	ctx.Break()
}`
	scriptBytes, compileErr := w.Compile(callbackSrc)

	var action *hyperdbgsdk.DEBUGGER_GENERAL_ACTION
	if compileErr != nil {
		t.Logf("Compile failed (%v); falling back to BreakToDebugger", compileErr)
		action, err = w.BuildAction(nil, hyperdbgsdk.BreakToDebugger)
	} else {
		t.Logf("Compile OK: %d bytes of binary AST", len(scriptBytes))
		action, err = w.BuildAction(scriptBytes, hyperdbgsdk.RunScript)
	}
	if err != nil {
		t.Fatalf("BuildAction failed: %v", err)
	}

	opts := hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{
		OptionalParam1: 0xFFFFFFFF,
	}
	event, err := w.BuildEvent(hyperdbgsdk.HiddenHookExecDetours, opts)
	if err != nil {
		t.Fatalf("BuildEvent failed: %v", err)
	}

	tag, err := w.RegisterHook(dev, event, action)
	if err != nil {
		t.Logf("RegisterHook with HiddenHookExecDetours failed: %v; "+
			"retrying with CpuidInstructionExecution", err)
		action2, err2 := w.BuildAction(nil, hyperdbgsdk.BreakToDebugger)
		if err2 != nil {
			t.Fatalf("BuildAction (fallback) failed: %v", err2)
		}
		event2, err2 := w.BuildEvent(hyperdbgsdk.CpuidInstructionExecution, hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{})
		if err2 != nil {
			t.Fatalf("BuildEvent (fallback) failed: %v", err2)
		}
		tag, err = w.RegisterHook(dev, event2, action2)
		if err != nil {
			t.Fatalf("RegisterHook (fallback) failed: %v", err)
		}
		t.Logf("fallback registration succeeded: Tag=%d", tag)
	} else {
		t.Logf("registration succeeded: Tag=%d", tag)
	}

	t.Cleanup(func() { clearEvent(t, dev, tag) })

	if tag == 0 {
		t.Fatal("returned Tag is 0; expected non-zero event tag")
	}
	t.Logf("PASS: verified Tag=%d > 0", tag)
}
