//go:build windows

package app

import (
	"fmt"

	"golang.org/x/sys/windows"
)

// enableDebugPrivilege enables SeDebugPrivilege on the current process
// token. This is required to open handles to privileged processes and to
// load kernel drivers. Mirrors the C++ SetDebugPrivilege helper used by
// HyperDbgLoadKdModule / HyperDbgLoadVmmModule.
func enableDebugPrivilege() error {
	var token windows.Token
	if err := windows.OpenProcessToken(windows.CurrentProcess(), windows.TOKEN_ADJUST_PRIVILEGES|windows.TOKEN_QUERY, &token); err != nil {
		return fmt.Errorf("OpenProcessToken: %w", err)
	}
	defer token.Close()

	var luid windows.LUID
	if err := windows.LookupPrivilegeValue(nil, windows.StringToUTF16Ptr("SeDebugPrivilege"), &luid); err != nil {
		return fmt.Errorf("LookupPrivilegeValue(SeDebugPrivilege): %w", err)
	}

	privs := []windows.LUIDAndAttributes{
		{Luid: luid, Attributes: windows.SE_PRIVILEGE_ENABLED},
	}
	if err := windows.AdjustTokenPrivileges(token, false, &windows.Tokenprivileges{
		PrivilegeCount: 1,
		Privileges:     [1]windows.LUIDAndAttributes(privs),
	}, 0, nil, nil); err != nil {
		return fmt.Errorf("AdjustTokenPrivileges: %w", err)
	}
	return nil
}
