package themida

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ddkwork/x64dbg/debugger"
)

func TestThemidaUnpacker_Basic(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过集成测试")
	}

	config := Config{
		SetEventUserData:  false,
		CheckHWID:         false,
		BypassHWIDSimple:  false,
		TryIATPatch:       true,
		AllocSize:         0x200000,
		AllocSizePEADS:    0x30000,
		XBundlerAuto:      true,
		UseMessageHWBP:    true,
		ARImpRecPath:      filepath.Join(os.TempDir(), "api_log.txt"),
		SetEventEntryAddr: 0,
		IOMarkerAddress:   0,
		SecLocation:       0,
	}

	testCases := []struct {
		name     string
		exePath  string
		expected VMVersion
	}{
		{
			name:     "测试 VM 版本检测",
			exePath:  "",
			expected: VMUnknown,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.exePath == "" {
				t.Skip("需要提供测试可执行文件路径")
			}

			dbg := debugger.New()
			if err := dbg.CreateProcess(tc.exePath, ""); err != nil {
				t.Fatalf("创建进程失败: %v", err)
			}
			defer dbg.Detach()

			unpacker := NewThemidaUnpacker(dbg, config)
			if err := unpacker.Start(); err != nil {
				t.Fatalf("启动脱壳器失败: %v", err)
			}
			defer unpacker.Stop()

			vmVersion := unpacker.GetVMVersion()
			if vmVersion != tc.expected {
				t.Errorf("VM 版本不匹配: 期望 %v, 实际 %v", tc.expected, vmVersion)
			}

			time.Sleep(2 * time.Second)
		})
	}
}

func TestThemidaUnpacker_APILogger(t *testing.T) {
	logPath := filepath.Join(os.TempDir(), "test_api_log.txt")
	defer os.Remove(logPath)

	logger := NewAPILogger(logPath)
	if err := logger.Open(); err != nil {
		t.Fatalf("打开 API 日志器失败: %v", err)
	}
	defer logger.Close()

	logger.LogAPI("SetEvent", 0x00401000, 0x00402000)
	logger.LogSetEvent(0x00401000, 0x12345678, 0x00402000)
	logger.LogIOMarker(0x00403000, 0x00400000, 0x00003000)
	logger.LogVMVersion(VMNew)
	logger.LogSection(0x00400000, 0x00100000)
	logger.LogError("测试错误信息")

	data, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("读取日志文件失败: %v", err)
	}

	if len(data) == 0 {
		t.Error("日志文件为空")
	}

	content := string(data)
	if !contains(content, "SetEvent") {
		t.Error("日志中未找到 SetEvent 记录")
	}
	if !contains(content, "I_O_MARKER_ADDRESS") {
		t.Error("日志中未找到 I_O_MARKER_ADDRESS 记录")
	}
}

func TestThemidaUnpacker_VMSignatureDetection(t *testing.T) {
	testCases := []struct {
		name     string
		data     []byte
		expected VMVersion
	}{
		{
			name:     "旧版 VM 签名",
			data:     []byte{0x68, 0x01, 0x00, 0x00, 0x00, 0xE9, 0x10, 0x00, 0x00, 0x00, 0xFF, 0x68, 0x20, 0x00, 0x00, 0x00},
			expected: VMOld,
		},
		{
			name:     "新版 VM 签名",
			data:     []byte{0x68, 0x01, 0x00, 0x00, 0x00, 0x68, 0x02, 0x00, 0x00, 0x00, 0xE9, 0x10, 0x00, 0x00, 0x00, 0xFF, 0x68, 0x03, 0x00, 0x00, 0x00},
			expected: VMNew,
		},
		{
			name:     "未知 VM 签名",
			data:     []byte{0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90},
			expected: VMRISC,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			unpacker := NewThemidaUnpacker(nil, Config{})
			version := unpacker.analyzeSignature(tc.data)
			if version != tc.expected {
				t.Errorf("VM 版本检测失败: 期望 %v, 实际 %v", tc.expected, version)
			}
		})
	}
}

func TestThemidaUnpacker_WLSectionDetection(t *testing.T) {
	testCases := []struct {
		name        string
		sectionName string
		expected    bool
	}{
		{"Themida 段", ".themida", true},
		{"WinLicense 段", ".winlice", true},
		{"WL 段", ".wlsec", true},
		{"Packed 段", ".packed", true},
		{"VMP 段", ".vmp0", true},
		{"普通代码段", ".text", false},
		{"普通数据段", ".data", false},
		{"普通资源段", ".rsrc", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			unpacker := NewThemidaUnpacker(nil, Config{})
			result := unpacker.isWLSection(tc.sectionName)
			if result != tc.expected {
				t.Errorf("段检测失败: %s 期望 %v, 实际 %v", tc.sectionName, tc.expected, result)
			}
		})
	}
}

func TestThemidaUnpacker_PatternMatching(t *testing.T) {
	testCases := []struct {
		name     string
		data     []byte
		pattern  []byte
		expected bool
	}{
		{
			name:     "完全匹配",
			data:     []byte{0x68, 0x01, 0x00, 0x00, 0x00, 0xE9, 0x10, 0x00, 0x00, 0x00},
			pattern:  []byte{0x68, 0x00, 0x00, 0x00, 0x00, 0xE9, 0x00, 0x00, 0x00, 0x00},
			expected: true,
		},
		{
			name:     "不匹配",
			data:     []byte{0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90},
			pattern:  []byte{0x68, 0x00, 0x00, 0x00, 0x00, 0xE9, 0x00, 0x00, 0x00, 0x00},
			expected: false,
		},
		{
			name:     "数据长度不足",
			data:     []byte{0x68, 0x01, 0x00, 0x00},
			pattern:  []byte{0x68, 0x00, 0x00, 0x00, 0x00, 0xE9, 0x00, 0x00, 0x00, 0x00},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			unpacker := NewThemidaUnpacker(nil, Config{})
			result := unpacker.matchPattern(tc.data, tc.pattern)
			if result != tc.expected {
				t.Errorf("模式匹配失败: 期望 %v, 实际 %v", tc.expected, result)
			}
		})
	}
}

func TestThemidaUnpacker_Config(t *testing.T) {
	config := Config{
		SetEventUserData:  true,
		CheckHWID:         true,
		BypassHWIDSimple:  true,
		TryIATPatch:       true,
		AllocSize:         0x200000,
		AllocSizePEADS:    0x30000,
		XBundlerAuto:      true,
		UseMessageHWBP:    true,
		ARImpRecPath:      "C:\\ARImpRec.dll",
		SetEventEntryAddr: 0x00401000,
		IOMarkerAddress:   0x00402000,
		SecLocation:       0x00400000,
	}

	unpacker := NewThemidaUnpacker(nil, config)
	if unpacker == nil {
		t.Fatal("创建脱壳器失败")
	}

	if unpacker.GetVMVersion() != VMUnknown {
		t.Error("初始 VM 版本应该是 VMUnknown")
	}

	if unpacker.GetSetEventAddress() != 0 {
		t.Error("初始 SetEvent 地址应该是 0")
	}

	if unpacker.GetIOMarkerAddress() != 0 {
		t.Error("初始 I/O 标记地址应该是 0")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || findSubstring(s, substr)))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
