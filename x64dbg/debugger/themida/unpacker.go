package themida

import (
	"fmt"
	"log"
	"sync"

	"github.com/ddkwork/x64dbg/debugger"
	"github.com/ddkwork/x64dbg/debugger/windows"
)

type VMVersion int

const (
	VMUnknown VMVersion = iota
	VMOld
	VMNew
	VMRISC
)

type ThemidaUnpacker struct {
	dbg          *debugger.Debugger
	vmVersion    VMVersion
	wlSection    uint32
	wlSize       uint32
	setEventAddr uint64
	ioMarkerAddr uint64
	apiLogger    *APILogger
	peRebuilder  *PERebuilder
	mu           sync.Mutex
	eventChan    chan *windows.DebugEvent
	stopChan     chan struct{}
}

type Config struct {
	SetEventUserData  bool
	CheckHWID         bool
	BypassHWIDSimple  bool
	TryIATPatch       bool
	AllocSize         uint32
	AllocSizePEADS    uint32
	XBundlerAuto      bool
	UseMessageHWBP    bool
	ARImpRecPath      string
	SetEventEntryAddr uint32
	IOMarkerAddress   uint32
	SecLocation       uint32
}

func NewThemidaUnpacker(dbg *debugger.Debugger, config Config) *ThemidaUnpacker {
	return &ThemidaUnpacker{
		dbg:         dbg,
		vmVersion:   VMUnknown,
		eventChan:   make(chan *windows.DebugEvent, 100),
		stopChan:    make(chan struct{}),
		apiLogger:   NewAPILogger(config.ARImpRecPath),
		peRebuilder: NewPERebuilder(),
	}
}

func (t *ThemidaUnpacker) Start() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	log.Println("开始 Themida 脱壳过程...")

	if err := t.detectVMVersion(); err != nil {
		return fmt.Errorf("检测 VM 版本失败: %v", err)
	}

	if err := t.setupBreakpoints(); err != nil {
		return fmt.Errorf("设置断点失败: %v", err)
	}

	go t.eventLoop()

	return nil
}

func (t *ThemidaUnpacker) Stop() {
	close(t.stopChan)
	t.apiLogger.Close()
}

func (t *ThemidaUnpacker) detectVMVersion() error {
	log.Println("检测 VM 版本...")

	baseAddr := uint32(t.dbg.GetBaseAddress())
	signature := t.readVMSignature(baseAddr)
	t.vmVersion = t.analyzeSignature(signature)

	log.Printf("检测到 VM 版本: %v", t.vmVersion)
	return nil
}

func (t *ThemidaUnpacker) isWLSection(name string) bool {
	wlNames := []string{".themida", ".winlice", ".wlsec", ".packed", ".vmp0", ".vmp1", ".vmp2"}
	for _, wlName := range wlNames {
		if name == wlName {
			return true
		}
	}
	return false
}

func (t *ThemidaUnpacker) readVMSignature(addr uint32) []byte {
	data, err := t.dbg.GetMemory().ReadMemory(uint64(addr), 32)
	if err != nil {
		log.Printf("读取 VM 签名失败: %v", err)
		return nil
	}
	return data
}

func (t *ThemidaUnpacker) analyzeSignature(data []byte) VMVersion {
	if data == nil || len(data) < 16 {
		return VMUnknown
	}

	oldPattern := []byte{0x68, 0x00, 0x00, 0x00, 0x00, 0xE9, 0x00, 0x00, 0x00, 0x00, 0xFF, 0x68, 0x00, 0x00, 0x00, 0x00}
	newPattern := []byte{0x68, 0x00, 0x00, 0x00, 0x00, 0x68, 0x00, 0x00, 0x00, 0x00, 0xE9, 0x00, 0x00, 0x00, 0x00, 0xFF, 0x68, 0x00, 0x00, 0x00, 0x00}

	if t.matchPattern(data, oldPattern) {
		return VMOld
	}
	if t.matchPattern(data, newPattern) {
		return VMNew
	}

	return VMRISC
}

func (t *ThemidaUnpacker) matchPattern(data, pattern []byte) bool {
	if len(data) < len(pattern) {
		return false
	}

	for i := range pattern {
		if pattern[i] != 0x00 && data[i] != pattern[i] {
			return false
		}
	}
	return true
}

func (t *ThemidaUnpacker) setupBreakpoints() error {
	log.Println("设置断点...")

	setEventAddr := uint64(0x7FFE5FD87FB0)

	if _, err := t.dbg.GetBreakpoints().SetSoftwareBreakpoint(setEventAddr, true); err != nil {
		log.Printf("设置 SetEvent 断点失败（可能程序未使用此功能）: %v", err)
	} else {
		log.Printf("SetEvent 断点设置成功: 0x%X", setEventAddr)
	}

	return nil
}

func (t *ThemidaUnpacker) eventLoop() {
	for {
		select {
		case <-t.stopChan:
			return
		case event := <-t.eventChan:
			t.handleEvent(event)
		}
	}
}

func (t *ThemidaUnpacker) handleEvent(event *windows.DebugEvent) {
	if event.DebugEventCode() != windows.EXCEPTION_DEBUG_EVENT {
		return
	}

	if event.Exception().ExceptionRecord.ExceptionCode == windows.EXCEPTION_BREAKPOINT {
		t.handleBreakpoint(event)
	} else if event.Exception().ExceptionRecord.ExceptionCode == windows.EXCEPTION_SINGLE_STEP {
		t.handleSingleStep(event)
	}
}

func (t *ThemidaUnpacker) handleBreakpoint(event *windows.DebugEvent) {
	eip := event.Exception().ExceptionRecord.ExceptionAddress

	if eip == t.setEventAddr {
		t.handleSetEvent(event)
	}
}

func (t *ThemidaUnpacker) handleSetEvent(event *windows.DebugEvent) {
	log.Printf("SetEvent 被调用: 0x%X", event.Exception().ExceptionRecord.ExceptionAddress)

	retAddr := t.getReturnAddress(event)
	t.apiLogger.LogAPI("SetEvent", uint32(event.Exception().ExceptionRecord.ExceptionAddress), retAddr)

	if t.vmVersion == VMOld {
		t.findOldVMEntry(event, retAddr)
	} else if t.vmVersion == VMNew {
		t.findNewVMEntry(event, retAddr)
	} else if t.vmVersion == VMRISC {
		t.findRISCVMEntry(event)
	}
}

func (t *ThemidaUnpacker) getReturnAddress(event *windows.DebugEvent) uint32 {
	threadId := event.ThreadId()
	thread := t.dbg.GetThreads().GetThread(threadId)
	if thread == nil {
		return 0
	}

	esp, err := t.dbg.GetRegisters().GetRegister(thread.Handle, "ESP")
	if err != nil {
		return 0
	}

	retAddr, err := t.dbg.GetMemory().ReadMemory(esp, 4)
	if err != nil {
		return 0
	}

	return uint32(retAddr[0]) | (uint32(retAddr[1]) << 8) | (uint32(retAddr[2]) << 16) | (uint32(retAddr[3]) << 24)
}

func (t *ThemidaUnpacker) findOldVMEntry(event *windows.DebugEvent, retAddr uint32) {
	eip := event.Exception().ExceptionRecord.ExceptionAddress

	data, err := t.dbg.GetMemory().ReadMemory(eip, 8)
	if err != nil {
		log.Printf("读取内存失败: %v", err)
		return
	}

	if data[0] == 0x68 && data[5] == 0xE9 {
		pushValue := uint32(data[1]) | (uint32(data[2]) << 8) | (uint32(data[3]) << 16) | (uint32(data[4]) << 24)
		jumpOffset := uint32(data[6]) | (uint32(data[7]) << 8) | (uint32(data[8]) << 16) | (uint32(data[9]) << 24)
		jumpAddr := eip + 0x0A + uint64(jumpOffset)

		t.setEventAddr = eip
		t.ioMarkerAddr = uint64(retAddr)
		log.Printf("找到旧版 VM 入口: PUSH=0x%X, JUMP=0x%X", pushValue, jumpAddr)
		log.Printf("I/O Marker 地址: 0x%X (来自返回地址)", retAddr)
	}
}

func (t *ThemidaUnpacker) findNewVMEntry(event *windows.DebugEvent, retAddr uint32) {
	eip := event.Exception().ExceptionRecord.ExceptionAddress

	data, err := t.dbg.GetMemory().ReadMemory(eip, 12)
	if err != nil {
		log.Printf("读取内存失败: %v", err)
		return
	}

	if data[0] == 0x68 && data[5] == 0x68 && data[10] == 0xE9 {
		pushValue1 := uint32(data[1]) | (uint32(data[2]) << 8) | (uint32(data[3]) << 16) | (uint32(data[4]) << 24)
		pushValue2 := uint32(data[6]) | (uint32(data[7]) << 8) | (uint32(data[8]) << 16) | (uint32(data[9]) << 24)
		jumpOffset := uint32(data[11]) | (uint32(data[12]) << 8) | (uint32(data[13]) << 16) | (uint32(data[14]) << 24)
		jumpAddr := eip + 0x0F + uint64(jumpOffset)

		t.setEventAddr = eip
		t.ioMarkerAddr = uint64(retAddr)
		log.Printf("找到新版 VM 入口: PUSH1=0x%X, PUSH2=0x%X, JUMP=0x%X", pushValue1, pushValue2, jumpAddr)
		log.Printf("I/O Marker 地址: 0x%X (来自返回地址)", retAddr)
	}
}

func (t *ThemidaUnpacker) findRISCVMEntry(event *windows.DebugEvent) {
	log.Println("查找 RISC VM 入口...")

	eip := event.Exception().ExceptionRecord.ExceptionAddress
	threadId := event.ThreadId()
	thread := t.dbg.GetThreads().GetThread(threadId)
	if thread == nil {
		log.Printf("获取线程失败")
		return
	}

	var secLocation uint64 = 0

	for range 1000 {
		data, err := t.dbg.GetMemory().ReadMemory(eip, 2)
		if err != nil {
			log.Printf("读取内存失败: %v", err)
			return
		}

		if data[0] == 0x8B && data[1] == 0xB5 {
			secLocation = eip
			log.Printf("找到 SECLOCATION (8B B5): 0x%X", eip)
			break
		}

		if err := t.dbg.StepInto(); err != nil {
			log.Printf("单步执行失败: %v", err)
			return
		}

		eip, _ = t.dbg.GetRegisters().GetRegister(thread.Handle, "EIP")
	}

	if secLocation == 0 {
		log.Println("未找到 SECLOCATION (8B B5)")
		return
	}

	for range 1000 {
		data, err := t.dbg.GetMemory().ReadMemory(eip, 1)
		if err != nil {
			log.Printf("读取内存失败: %v", err)
			return
		}

		if data[0] == 0xF0 {
			t.ioMarkerAddr = eip
			log.Printf("找到 I/O Marker (F0): 0x%X", eip)
			ioMarkerRVA := uint32(eip - secLocation)
			log.Printf("I/O Marker RVA: 0x%X", ioMarkerRVA)
			return
		}

		if err := t.dbg.StepInto(); err != nil {
			log.Printf("单步执行失败: %v", err)
			return
		}

		eip, _ = t.dbg.GetRegisters().GetRegister(thread.Handle, "EIP")
	}

	log.Println("未找到 I/O Marker (F0)")
}

func (t *ThemidaUnpacker) handleSingleStep(event *windows.DebugEvent) {
	eip := event.Exception().ExceptionRecord.ExceptionAddress

	data, err := t.dbg.GetMemory().ReadMemory(eip, 1)
	if err != nil {
		return
	}

	if data[0] == 0xF0 {
		log.Printf("找到 RISC VM 指令: 0x%X", eip)
		t.ioMarkerAddr = eip
	}
}

func (t *ThemidaUnpacker) DumpProcess(outputPath string) error {
	log.Println("开始转储进程...")

	if err := t.peRebuilder.DumpProcess(t.dbg, outputPath); err != nil {
		return fmt.Errorf("转储进程失败: %v", err)
	}

	if err := t.peRebuilder.RebuildPE(outputPath); err != nil {
		return fmt.Errorf("重建 PE 文件失败: %v", err)
	}

	log.Printf("进程转储完成: %s", outputPath)
	return nil
}

func (t *ThemidaUnpacker) GetVMVersion() VMVersion {
	return t.vmVersion
}

func (t *ThemidaUnpacker) GetSetEventAddress() uint64 {
	return t.setEventAddr
}

func (t *ThemidaUnpacker) GetIOMarkerAddress() uint64 {
	return t.ioMarkerAddr
}
