// Package core — load_unload_cycle_test.go
//
// 验证驱动+VMM的重复加载/卸载循环。每次循环4个步骤：
// 1. 加载驱动  2. 加载VMM  3. 卸载VMM  4. 卸载驱动
//
// 运行（管理员 PowerShell）：
//
//	go test -v -count=1 -run TestLoadUnloadCycle -timeout 600s ./debugger/core/
package core

import (
	"testing"
)

// TestLoadUnloadCycle 重复加载卸载驱动+VMM 20次，确保每次都能成功加载和卸载。
func TestLoadUnloadCycle(t *testing.T) {
	const driverPath = `C:\Users\Administrator\AppData\Local\hyperdbg\hyperkd.sys`
	const cycles = 20

	for i := 1; i <= cycles; i++ {
		dbg := New()

		// 1. 加载驱动
		if err := dbg.LoadDriver(driverPath); err != nil {
			t.Fatalf("cycle %d/%d 加载驱动失败: %v", i, cycles, err)
		}

		// 2. 加载VMM
		if err := dbg.InitVMM(); err != nil {
			_ = dbg.UnloadDriver()
			t.Fatalf("cycle %d/%d 加载VMM失败: %v", i, cycles, err)
		}

		// 3. 卸载VMM
		if err := dbg.UnloadVMM(); err != nil {
			_ = dbg.UnloadDriver()
			t.Fatalf("cycle %d/%d 卸载VMM失败: %v", i, cycles, err)
		}

		// 4. 卸载驱动
		if err := dbg.UnloadDriver(); err != nil {
			t.Fatalf("cycle %d/%d 卸载驱动失败: %v", i, cycles, err)
		}

		t.Logf("cycle %d/%d OK", i, cycles)
	}
}
