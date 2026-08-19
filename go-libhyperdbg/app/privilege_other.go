//go:build !windows

package app

// enableDebugPrivilege is a stub for non-Windows platforms. SeDebugPrivilege
// is a Windows-only concept; on other platforms there is nothing to enable.
// Returns nil so the App can still be constructed for tests.
func enableDebugPrivilege() error {
	return nil
}
