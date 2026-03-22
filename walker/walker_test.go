package walker

import (
	"fmt"
	"testing"
	"time"

	"github.com/ddkwork/HyperDbg/debugger"
	"github.com/ddkwork/golibrary/std/mylog"
)

func TestWalker(t *testing.T) {
	t.Run("IoctlCodeToString", func(t *testing.T) {
		ioctlCodes := map[uint32]string{
			0x806: "IOCTL_DEBUGGER_REGISTER_EVENT",
			0x807: "IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT",
			0x808: "IOCTL_DEBUGGER_MODIFY_EVENTS",
		}
		for code, name := range ioctlCodes {
			t.Logf("0x%04X: %s", code, name)
		}
	})

	t.Run("SyscallNumberLookup", func(t *testing.T) {
		syscalls := map[uint32]string{
			0x00: "NtAllocateVirtualMemory",
			0x07: "NtDeviceIoControlFile",
			0x12: "NtAllocateVirtualMemory",
		}
		for num, name := range syscalls {
			t.Logf("0x%02X: %s", num, name)
		}
	})

	t.Run("IoControlCodeParsing", func(t *testing.T) {
		testCodes := []uint32{0x00220000, 0x80002000, 0x80002004}
		for _, code := range testCodes {
			t.Logf("0x%08X -> %s", code, IoControlCodeToString(code))
		}
	})

	t.Run("ProcessBehaviorLogger", func(t *testing.T) {
		dbg := debugger.NewUserDebug()
		defer dbg.UnloadDriver()

		mylog.Call(func() {
			targetPid := uint32(0)
			walker := NewProcessBehaviorLogger(targetPid)
			logger := walker.(*ProcessBehaviorLogger)
			mylog.Info("Process Behavior Logger initialized")
			mylog.Info("Target PID: ", targetPid, " (0 means all processes)")

			eventChan := dbg.Packet().GetEventChan()
			dbg.HookSyscall(0)

			mylog.Info("All syscalls hooked for behavior logging...")
			timeout := time.After(5 * time.Second)

			for {
				select {
				case event := <-eventChan:
					if event == nil {
						continue
					}
					logger.Maiter(event, func() {
						t.Logf("method:syscall, pid: %x, tid: %x, syscall: %x",
							event.ProcessId, event.ThreadId, event.RAX)
					})
				case <-timeout:
					mylog.Info("Process behavior logging finished")
					t.Log(logger.Summary())
					return
				}
			}
		})
	})

	t.Run("MemoryAllocationMonitor", func(t *testing.T) {
		dbg := debugger.NewUserDebug()
		defer dbg.UnloadDriver()

		mylog.Call(func() {
			targetPid := uint32(0)
			walker := NewMemoryAllocationMonitor(targetPid)
			monitor := walker.(*MemoryAllocationMonitor)
			mylog.Info("Memory Allocation Monitor initialized")

			eventChan := dbg.Packet().GetEventChan()
			dbg.HookSyscall(SyscallNtAllocateVirtualMemory)

			mylog.Info("NtAllocateVirtualMemory hooked...")
			timeout := time.After(5 * time.Second)

			for {
				select {
				case event := <-eventChan:
					if event == nil {
						continue
					}
					monitor.Maiter(event, func() {
						t.Logf("[SYSCALL] NtAllocateVirtualMemory, pid: %x | size: %x",
							event.ProcessId, event.R9)
					})
				case <-timeout:
					mylog.Info("Memory allocation monitoring finished")
					t.Log(monitor.Summary())
					return
				}
			}
		})
	})

	t.Run("ComprehensiveMonitor", func(t *testing.T) {
		dbg := debugger.NewUserDebug()
		defer dbg.UnloadDriver()

		mylog.Call(func() {
			targetPid := uint32(0)
			walker := NewComprehensiveMonitor(targetPid)
			monitor := walker.(*ComprehensiveMonitor)
			mylog.Info("Comprehensive Monitor initialized")

			eventChan := dbg.Packet().GetEventChan()
			dbg.HookSyscall(0)

			mylog.Info("All syscalls hooked for comprehensive monitoring...")
			timeout := time.After(5 * time.Second)

			for {
				select {
				case event := <-eventChan:
					monitor.Maiter(event, func() {
						t.Logf("syscall: %x, pid: %x", event.RAX, event.ProcessId)
					})
				case <-timeout:
					mylog.Info("Comprehensive monitoring finished")
					t.Log(monitor.Summary())
					return
				}
			}
		})
	})

	t.Run("DFIRMonitor", func(t *testing.T) {
		dbg := debugger.NewUserDebug()
		defer dbg.UnloadDriver()

		mylog.Call(func() {
			targetPid := uint32(0)
			walker := NewDFIRMonitor(targetPid)
			monitor := walker.(*DFIRMonitor)
			mylog.Info("DFIR Monitor initialized")

			eventChan := dbg.Packet().GetEventChan()
			dbg.HookSyscall(0)

			mylog.Info("DFIR monitoring started...")
			timeout := time.After(5 * time.Second)

			for {
				select {
				case event := <-eventChan:
					if event == nil {
						continue
					}
					monitor.Maiter(event, func() {
						t.Logf("method:syscall, pid: %x, tid: %x, syscall: %x",
							event.ProcessId, event.ThreadId, event.RAX)
					})
				case <-timeout:
					mylog.Info("DFIR monitoring finished")
					t.Log(monitor.Summary())
					return
				}
			}
		})
	})

	t.Run("NtDeviceIoControlFileMonitor", func(t *testing.T) {
		dbg := debugger.NewUserDebug()
		defer dbg.UnloadDriver()

		mylog.Call(func() {
			mylog.Info("Starting NtDeviceIoControlFile monitor")

			eventChan := dbg.Packet().GetEventChan()
			dbg.HookSyscall(SyscallNtDeviceIoControlFile)

			mylog.Info("Syscall hook registered, waiting for events...")
			timeout := time.After(5 * time.Second)
			eventCount := 0

			for {
				select {
				case event := <-eventChan:
					if event == nil {
						continue
					}
					eventCount++
					info := ParseNtDeviceIoControlFileEvent(event)
					if info != nil {
						t.Logf("[%d] %s", eventCount, info.String())
						t.Logf("  IOCTL: %s", IoControlCodeToString(info.IoControlCode))
					}
				case <-timeout:
					mylog.Info("Monitor finished, total events:", eventCount)
					return
				}
			}
		})
	})

	t.Run("HardwareHookCallback", func(t *testing.T) {
		dbg := debugger.NewUserDebug()
		defer dbg.UnloadDriver()

		mylog.Call(func() {
			mylog.Info("Starting NtDeviceIoControlFile hook with hardware spoofing")

			walker := NewHardwareHookCallback()
			hookCallback := walker.(*HardwareHookCallback)
			eventChan := dbg.Packet().GetEventChan()
			dbg.HookSyscall(SyscallNtDeviceIoControlFile)

			mylog.Info("Syscall hook registered, monitoring for hardware queries...")
			timeout := time.After(5 * time.Second)
			eventCount := 0

			for {
				select {
				case event := <-eventChan:
					if event == nil {
						continue
					}
					eventCount++
					hookCallback.Hook(event, func() {
						info := ParseNtDeviceIoControlFileEvent(event)
						if info != nil {
							t.Logf("[%d] PID: %d | IOCTL: %s",
								eventCount, info.ProcessId, IoControlCodeToString(info.IoControlCode))
						}
					})
				case <-timeout:
					mylog.Info("Hook finished")
					t.Logf("Total events: %d", eventCount)
					return
				}
			}
		})
	})

	t.Run("IopXxxControlFileMonitor", func(t *testing.T) {
		dbg := debugger.NewUserDebug()
		defer dbg.UnloadDriver()

		mylog.Call(func() {
			walker := NewIopXxxControlFileMonitor()
			monitor := walker.(*IopXxxControlFileMonitor)
			mylog.Info("IopXxxControlFile Monitor initialized")
			mylog.Info("Target IOCTL codes:")
			for code, name := range monitor.targetIoctlCodes {
				mylog.Info("  ", fmt.Sprintf("0x%08X: %s", code, name))
			}

			eventChan := dbg.Packet().GetEventChan()
			dbg.HookSyscall(SyscallNtDeviceIoControlFile)

			mylog.Info("Structured monitoring started (5s timeout)...")
			timeout := time.After(5 * time.Second)

			for {
				select {
				case event := <-eventChan:
					monitor.Maiter(event, func() {
						monitor.Hook(event, func() {
							info := ParseNtDeviceIoControlFileEvent(event)
							if info != nil {
								t.Logf("Matched IOCTL: %s", IoControlCodeToString(info.IoControlCode))
							}
						})
					})
				case <-timeout:
					mylog.Info("Monitor finished")
					t.Log(monitor.Summary())
					return
				}
			}
		})
	})

	t.Run("SSDSpoofTest", func(t *testing.T) { // todo  nvme
		dbg := debugger.NewUserDebug()
		defer dbg.UnloadDriver()

		mylog.Call(func() {
			walker := NewHardwareHookCallback()
			hookCallback := walker.(*HardwareHookCallback)
			hookInfo := hookCallback.GetHookInfo()
			mylog.Info("SSD Spoof Configuration:")
			mylog.Info("  Model: ", hookInfo.SsdNumber.Mod.String())
			mylog.Info("  Serial: ", hookInfo.SsdNumber.Serial.String())

			eventChan := dbg.Packet().GetEventChan()
			dbg.HookSyscall(SyscallNtDeviceIoControlFile)

			mylog.Info("Monitoring for SMART_RCV_DRIVE_DATA ", fmt.Sprintf("(0x%08X)", SMART_RCV_DRIVE_DATA))
			timeout := time.After(5 * time.Second)
			smartCount := 0

			for {
				select {
				case event := <-eventChan:
					if event == nil {
						continue
					}
					info := ParseNtDeviceIoControlFileEvent(event)
					if info == nil {
						continue
					}
					if info.IoControlCode == SMART_RCV_DRIVE_DATA {
						smartCount++
						t.Logf("[SMART %d] PID: %d | Handle: 0x%x",
							smartCount, info.ProcessId, info.FileHandle)
						t.Logf("  Would spoof to: Model=%s, Serial=%s",
							hookInfo.SsdNumber.Mod.String(),
							hookInfo.SsdNumber.Serial.String())
					}
				case <-timeout:
					mylog.Info("SSD spoof test finished, SMART calls:", smartCount)
					return
				}
			}
		})
	})

	t.Run("CPUSpoofTest", func(t *testing.T) {
		dbg := debugger.NewUserDebug()
		defer dbg.UnloadDriver()

		mylog.Call(func() {
			walker := NewHardwareHookCallback()
			hookCallback := walker.(*HardwareHookCallback)
			hookInfo := hookCallback.GetHookInfo()
			mylog.Info("CPU Spoof Configuration:")
			mylog.Info("  EAX: ", fmt.Sprintf("0x%08X", hookInfo.Cpu1.Eax))
			mylog.Info("  ECX: ", fmt.Sprintf("0x%08X", hookInfo.Cpu1.Ecx))
			mylog.Info("  EDX: ", fmt.Sprintf("0x%08X", hookInfo.Cpu1.Edx))

			eventChan := dbg.Packet().GetEventChan()
			dbg.HookSyscall(SyscallNtDeviceIoControlFile)

			mylog.Info("Monitoring for CPUID IOCTL ", fmt.Sprintf("(0x%08X)", IOCTL_CPUID))
			timeout := time.After(5 * time.Second)
			cpuidCount := 0

			for {
				select {
				case event := <-eventChan:
					if event == nil {
						continue
					}
					info := ParseNtDeviceIoControlFileEvent(event)
					if info == nil {
						continue
					}
					if info.IoControlCode == IOCTL_CPUID {
						cpuidCount++
						t.Logf("[CPUID %d] PID: %d | Handle: 0x%x",
							cpuidCount, info.ProcessId, info.FileHandle)
						t.Logf("  Would spoof to: EAX=0x%08X, ECX=0x%08X, EDX=0x%08X",
							hookInfo.Cpu1.Eax, hookInfo.Cpu1.Ecx, hookInfo.Cpu1.Edx)
					}
				case <-timeout:
					mylog.Info("CPU spoof test finished, CPUID calls:", cpuidCount)
					return
				}
			}
		})
	})

	t.Run("NetworkConnectionMonitor", func(t *testing.T) {
		dbg := debugger.NewUserDebug()
		defer dbg.UnloadDriver()

		mylog.Call(func() {
			mylog.Info("Starting network connection monitor")
			mylog.Info("Monitoring AFD_CONNECT ", fmt.Sprintf("(0x%08X)", AFD_CONNECT), " for network connections")

			eventChan := dbg.Packet().GetEventChan()
			dbg.HookSyscall(SyscallNtDeviceIoControlFile)

			mylog.Info("Syscall hook registered, monitoring network connections...")
			timeout := time.After(5 * time.Second)
			connectionCount := 0

			for {
				select {
				case event := <-eventChan:
					if event == nil {
						continue
					}
					info := ParseNtDeviceIoControlFileEvent(event)
					if info == nil {
						continue
					}
					if info.IoControlCode == AFD_CONNECT {
						connectionCount++
						t.Logf("[%d] PID: %d | AFD_CONNECT detected",
							connectionCount, info.ProcessId)
					}
				case <-timeout:
					mylog.Info("Network monitor finished, connections:", connectionCount)
					return
				}
			}
		})
	})
}
