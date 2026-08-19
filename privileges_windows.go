// Package main — privileges_windows.go
//
// 启用进程所需的 Windows 特权。即使 UAC manifest 要求管理员权限，
// Windows 的 admin token 默认只是 "可用" 状态，需要 AdjustTokenPrivileges
// 显式启用才能生效。
//
// 关键特权：
//   - SeDebugPrivilege           调试任意进程，读取内核模块地址
//   - SeLoadDriverPrivilege      加载/卸载内核驱动
//   - SeSystemProfilePrivilege   系统级性能分析
//   - SeProfileSingleProcessPrivilege  单进程性能采样
//   - SeIncreaseBasePriorityPrivilege  提高调度优先级
//   - SeIncreaseQuotaPrivilege    提高内存/CPU 配额
package main

import (
	"fmt"

	"golang.org/x/sys/windows"
)

// enablePrivilege 启用指定的 Windows 特权。
func enablePrivilege(name string) error {
	var token windows.Token
	if err := windows.OpenProcessToken(windows.CurrentProcess(),
		windows.TOKEN_ADJUST_PRIVILEGES|windows.TOKEN_QUERY, &token); err != nil {
		return fmt.Errorf("OpenProcessToken: %w", err)
	}
	defer token.Close()

	var luid windows.LUID
	if err := windows.LookupPrivilegeValue(nil, windows.StringToUTF16Ptr(name), &luid); err != nil {
		return fmt.Errorf("LookupPrivilegeValue(%s): %w", name, err)
	}

	tp := windows.Tokenprivileges{
		PrivilegeCount: 1,
		Privileges: [1]windows.LUIDAndAttributes{
			{Luid: luid, Attributes: windows.SE_PRIVILEGE_ENABLED},
		},
	}
	if err := windows.AdjustTokenPrivileges(token, false, &tp, 0, nil, nil); err != nil {
		return fmt.Errorf("AdjustTokenPrivileges(%s): %w", name, err)
	}
	return nil
}

// enableDebugPrivileges 启用 HyperDbg 所需的全部特权。
// 任一失败只打印警告不中断，因为某些特权可能不在 token 中（取决于组策略）。
func enableDebugPrivileges() {
	privileges := []struct {
		name string
		desc string
	}{
		{"SeDebugPrivilege", "调试任意进程"},
		{"SeLoadDriverPrivilege", "加载/卸载内核驱动"},
		{"SeSystemProfilePrivilege", "系统级性能分析"},
		{"SeProfileSingleProcessPrivilege", "单进程性能采样"},
		{"SeIncreaseBasePriorityPrivilege", "提高调度优先级"},
		{"SeIncreaseQuotaPrivilege", "提高内存/CPU 配额"},
		{"SeBackupPrivilege", "备份文件（绕过读取权限）"},
		{"SeRestorePrivilege", "还原文件（绕过写入权限）"},
	}

	for _, p := range privileges {
		if err := enablePrivilege(p.name); err != nil {
			fmt.Printf("⚠️  启用特权失败 %s (%s): %v\n", p.name, p.desc, err)
		} else {
			fmt.Printf("✅ 特权已启用: %s (%s)\n", p.name, p.desc)
		}
	}
}
