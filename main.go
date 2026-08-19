package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"

	"github.com/ddkwork/HyperDbg/ui"
)

//go:embed Debug/hyperkd.sys
var driverSys []byte

//go:embed Debug/libipt.dll
var libiptDLL []byte

// extractAssets 把内嵌的驱动与 DLL 释放到固定目录 %LOCALAPPDATA%\hyperdbg\，
// 返回驱动 .sys 路径。固定路径确保多次运行不会产生指向已删除 temp 目录的
// stale 服务（driverloader.Install 对已存在服务不更新 binPath）。
func extractAssets() (driverPath string, err error) {
	base, err := os.UserCacheDir() // Windows: %LOCALAPPDATA%
	if err != nil {
		return "", fmt.Errorf("获取缓存目录失败: %w", err)
	}
	dir := filepath.Join(base, "hyperdbg")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", fmt.Errorf("创建目录失败: %w", err)
	}

	driverPath = filepath.Join(dir, "hyperkd.sys")
	if err := os.WriteFile(driverPath, driverSys, 0o644); err != nil {
		// 文件可能被上次运行加载的驱动占用（服务仍在 StopPending）。
		// 路径固定，已存在文件可继续使用——driverloader.Install 对已存在
		// 服务不更新 binPath，正好匹配。
		if _, statErr := os.Stat(driverPath); statErr != nil {
			return "", fmt.Errorf("写入驱动文件失败且无已有文件可回退: %w", err)
		}
		fmt.Printf("警告: 驱动文件被占用，使用已有文件: %s\n", driverPath)
	}

	dllPath := filepath.Join(dir, "libipt.dll")
	if err := os.WriteFile(dllPath, libiptDLL, 0o644); err != nil {
		if _, statErr := os.Stat(dllPath); statErr != nil {
			return "", fmt.Errorf("写入 libipt.dll 失败且无已有文件可回退: %w", err)
		}
		fmt.Printf("警告: libipt.dll 被占用，使用已有文件: %s\n", dllPath)
	}

	// 把目录加入 DLL 搜索路径，让 PT 模块能加载 libipt.dll
	setDllDirectory(dir)

	fmt.Printf("驱动路径: %s\n", driverPath)
	fmt.Printf("DLL 路径: %s\n", dllPath)
	return driverPath, nil
}

// setDllDirectory 调用 kernel32!SetDllDirectoryW 把 dir 加入 DLL 搜索路径。
func setDllDirectory(dir string) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetDllDirectoryW")
	ptr, _ := syscall.UTF16PtrFromString(dir)
	proc.Call(uintptr(unsafe.Pointer(ptr)))
}

func main() {
	// 启用调试所需的全部 Windows 特权（SeDebugPrivilege/SeLoadDriverPrivilege 等）
	enableDebugPrivileges()

	driverPath, err := extractAssets()
	if err != nil {
		fmt.Printf("释放内嵌资源失败: %v\n", err)
		return
	}
	ui.SetDriverPath(driverPath)
	ui.Run(nil)
}
