package themida_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ddkwork/x64dbg/debugger"
	"github.com/ddkwork/x64dbg/debugger/themida"
)

func TestSuperRecovery(t *testing.T) {
	fmt.Println("=== Themida 脱壳器 ===")
	fmt.Println()

	exePath := `SuperRecovery V4.8.1.5.exe`

	if _, err := os.Stat(exePath); os.IsNotExist(err) {
		log.Fatalf("目标文件不存在: %s", exePath)
	}

	outputDir := filepath.Join(filepath.Dir(exePath), "unpacked")
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		log.Fatalf("创建输出目录失败: %v", err)
	}

	fmt.Printf("目标文件: %s\n", exePath)
	fmt.Printf("输出目录: %s\n", outputDir)
	fmt.Println()

	config := themida.Config{
		SetEventUserData:  false,
		CheckHWID:         false,
		BypassHWIDSimple:  false,
		TryIATPatch:       true,
		AllocSize:         0x200000,
		AllocSizePEADS:    0x30000,
		XBundlerAuto:      true,
		UseMessageHWBP:    true,
		ARImpRecPath:      filepath.Join(outputDir, "api_log.txt"),
		SetEventEntryAddr: 0,
		IOMarkerAddress:   0,
		SecLocation:       0,
	}

	fmt.Println("步骤 1: 创建调试器")
	fmt.Println("----------------------------------------")
	dbg := debugger.New()
	fmt.Println("✓ 调试器创建成功")

	fmt.Println()
	fmt.Println("步骤 2: 创建脱壳器")
	fmt.Println("----------------------------------------")
	unpacker := themida.NewThemidaUnpacker(dbg, config)
	if unpacker == nil {
		log.Fatal("创建脱壳器失败")
	}
	fmt.Println("✓ 脱壳器创建成功")

	fmt.Println()
	fmt.Println("步骤 3: 启动目标进程")
	fmt.Println("----------------------------------------")
	if err := dbg.CreateProcess(exePath, ""); err != nil {
		log.Fatalf("启动进程失败: %v", err)
	}
	fmt.Println("✓ 进程启动成功")

	fmt.Println()
	fmt.Println("步骤 4: 启动脱壳器")
	fmt.Println("----------------------------------------")
	if err := unpacker.Start(); err != nil {
		log.Fatalf("启动脱壳器失败: %v", err)
	}
	fmt.Println("✓ 脱壳器启动成功")

	fmt.Println()
	fmt.Println("步骤 5: 等待 VM 版本检测")
	fmt.Println("----------------------------------------")
	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			log.Fatal("VM 版本检测超时")
		case <-ticker.C:
			vmVersion := unpacker.GetVMVersion()
			if vmVersion != themida.VMUnknown {
				fmt.Printf("✓ VM 版本检测完成: %v\n", vmVersion)
				goto VM_DETECTED
			}
			fmt.Print(".")
		}
	}

VM_DETECTED:
	fmt.Println()

	fmt.Println()
	fmt.Println("步骤 6: 等待程序运行")
	fmt.Println("----------------------------------------")
	fmt.Println("等待程序运行 10 秒...")
	time.Sleep(10 * time.Second)
	fmt.Println("✓ 程序运行完成")

	fmt.Println()
	fmt.Println("步骤 7: 转储进程内存")
	fmt.Println("----------------------------------------")
	dumpPath := filepath.Join(outputDir, "dumped.exe")
	if err := unpacker.DumpProcess(dumpPath); err != nil {
		log.Printf("转储进程失败: %v\n", err)
		fmt.Println("注意: 转储失败可能是因为进程已退出或内存保护")
	} else {
		fmt.Printf("✓ 进程转储完成: %s\n", dumpPath)
	}

	fmt.Println()
	fmt.Println("步骤 8: 停止脱壳器")
	fmt.Println("----------------------------------------")
	unpacker.Stop()
	fmt.Println("✓ 脱壳器已停止")

	fmt.Println()
	fmt.Println("步骤 9: 分离调试器")
	fmt.Println("----------------------------------------")
	if err := dbg.Detach(); err != nil {
		log.Printf("分离调试器失败: %v\n", err)
	} else {
		fmt.Println("✓ 调试器已分离")
	}

	fmt.Println()
	fmt.Println("=== 脱壳完成 ===")
	fmt.Println()
	fmt.Println("输出文件:")
	fmt.Printf("  - 转储文件: %s\n", dumpPath)
	fmt.Printf("  - API 日志: %s\n", config.ARImpRecPath)
	fmt.Println()
	fmt.Println("注意：")
	fmt.Println("1. 转储的文件需要使用 Scylla 修复 IAT")
	fmt.Println("2. 可能需要手动修复 PE 头和段表")
	fmt.Println("3. 建议使用 PE 工具验证文件完整性")
}
