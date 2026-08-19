package themida

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type APILogger struct {
	filePath string
	file     *os.File
	mu       sync.Mutex
}

type APICall struct {
	Timestamp time.Time
	CallAddr  uint32
	APIAddr   uint32
	APIName   string
	Module    string
}

func NewAPILogger(filePath string) *APILogger {
	return &APILogger{
		filePath: filePath,
	}
}

func (l *APILogger) Open() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.file != nil {
		return nil
	}

	file, err := os.OpenFile(l.filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return fmt.Errorf("打开 API 日志文件失败: %v", err)
	}

	l.file = file

	l.writeHeader()
	return nil
}

func (l *APILogger) Close() {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.file != nil {
		l.file.Close()
		l.file = nil
	}
}

func (l *APILogger) LogAPI(apiName string, apiAddr, callAddr uint32) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.file == nil {
		return
	}

	call := APICall{
		Timestamp: time.Now(),
		CallAddr:  callAddr,
		APIAddr:   apiAddr,
		APIName:   apiName,
	}

	line := fmt.Sprintf("[%s] Call from: 0x%X | API: 0x%X | NAME: %s\n",
		call.Timestamp.Format("2006-01-02 15:04:05.000"),
		call.CallAddr,
		call.APIAddr,
		call.APIName)

	if _, err := l.file.WriteString(line); err != nil {
		fmt.Printf("写入 API 日志失败: %v\n", err)
	}
}

func (l *APILogger) LogSetEvent(addr uint32, pushValue, jumpAddr uint32) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.file == nil {
		return
	}

	line := fmt.Sprintf("-------------------------------------------------------\n")
	line += "--------------- SETEVENT_ENTRY_ADDRESS ----------------\n"
	line += fmt.Sprintf("Address: 0x%X | PUSH: 0x%X | JUMP: 0x%X\n", addr, pushValue, jumpAddr)
	line += "-------------------------------------------------------\n"

	if _, err := l.file.WriteString(line); err != nil {
		fmt.Printf("写入 SetEvent 日志失败: %v\n", err)
	}
}

func (l *APILogger) LogIOMarker(addr uint32, secLocation uint32, rva uint32) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.file == nil {
		return
	}

	line := fmt.Sprintf("-------------------------------------------------------\n")
	line += "--------------- I_O_MARKER_ADDRESS --------------------\n"
	line += fmt.Sprintf("Address: 0x%X | Section Location: 0x%X | I_O_MARKER_ADDRESS RVA: 0x%X\n", addr, secLocation, rva)
	line += "-------------------------------------------------------\n"

	if _, err := l.file.WriteString(line); err != nil {
		fmt.Printf("写入 I/O 标记日志失败: %v\n", err)
	}
}

func (l *APILogger) LogVMVersion(version VMVersion) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.file == nil {
		return
	}

	var versionStr string
	switch version {
	case VMOld:
		versionStr = "OLD VM"
	case VMNew:
		versionStr = "NEW VM"
	case VMRISC:
		versionStr = "RISC VM"
	default:
		versionStr = "UNKNOWN VM"
	}

	line := fmt.Sprintf("VM Version: %s\n", versionStr)
	if _, err := l.file.WriteString(line); err != nil {
		fmt.Printf("写入 VM 版本日志失败: %v\n", err)
	}
}

func (l *APILogger) LogSection(sectionAddr, sectionSize uint32) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.file == nil {
		return
	}

	line := fmt.Sprintf("WL Section: 0x%X - 0x%X (Size: 0x%X)\n", sectionAddr, sectionAddr+sectionSize, sectionSize)
	if _, err := l.file.WriteString(line); err != nil {
		fmt.Printf("写入段日志失败: %v\n", err)
	}
}

func (l *APILogger) LogError(err string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.file == nil {
		return
	}

	line := fmt.Sprintf("[ERROR] %s\n", err)
	if _, err := l.file.WriteString(line); err != nil {
		fmt.Printf("写入错误日志失败: %v\n", err)
	}
}

func (l *APILogger) writeHeader() {
	if l.file == nil {
		return
	}

	header := fmt.Sprintf("=== Themida/WinLicense Unpacker Log ===\n")
	header += fmt.Sprintf("Started: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	header += "=======================================================\n\n"

	if _, err := l.file.WriteString(header); err != nil {
		fmt.Printf("写入日志头失败: %v\n", err)
	}
}
