package ddk

import (
	"github.com/ddkwork/golibrary/std/mylog"
	"golang.org/x/sys/windows"
)

type Privilege struct {
	adjusted map[string]bool
}

func NewPrivilege() *Privilege {
	return &Privilege{adjusted: make(map[string]bool)}
}

func (p *Privilege) Enable(name string) bool {
	if p.adjusted[name] {
		return true
	}
	var token windows.Token
	currentProcess := windows.CurrentProcess()
	mylog.Check(windows.OpenProcessToken(currentProcess, windows.TOKEN_ADJUST_PRIVILEGES|windows.TOKEN_QUERY, &token))
	defer func() { mylog.Check(token.Close()) }() // 没有func 函数体会被立即执行导致token.Close()被调用，AdjustTokenPrivileges是无效句柄， 所以defer 必须有func

	var id windows.LUID
	mylog.Check(windows.LookupPrivilegeValue(nil, windows.StringToUTF16Ptr(name), &id))
	mylog.Check(windows.AdjustTokenPrivileges(token, false, &windows.Tokenprivileges{
		PrivilegeCount: 1,
		Privileges:     [1]windows.LUIDAndAttributes{{Luid: id, Attributes: windows.SE_PRIVILEGE_ENABLED}},
	}, 0, nil, nil))
	p.adjusted[name] = true
	return true
}

func (p *Privilege) CreateToken() bool          { return p.Enable("SeCreateTokenPrivilege") }
func (p *Privilege) AssignPrimaryToken() bool   { return p.Enable("SeAssignPrimaryTokenPrivilege") }
func (p *Privilege) LockMemory() bool           { return p.Enable("SeLockMemoryPrivilege") }
func (p *Privilege) IncreaseQuota() bool        { return p.Enable("SeIncreaseQuotaPrivilege") }
func (p *Privilege) UnsolicitedInput() bool     { return p.Enable("SeUnsolicitedInputPrivilege") }
func (p *Privilege) MachineAccount() bool       { return p.Enable("SeMachineAccountPrivilege") }
func (p *Privilege) Tcb() bool                  { return p.Enable("SeTcbPrivilege") }
func (p *Privilege) Security() bool             { return p.Enable("SeSecurityPrivilege") }
func (p *Privilege) TakeOwnership() bool        { return p.Enable("SeTakeOwnershipPrivilege") }
func (p *Privilege) LoadDriver() bool           { return p.Enable("SeLoadDriverPrivilege") }
func (p *Privilege) SystemProfile() bool        { return p.Enable("SeSystemProfilePrivilege") }
func (p *Privilege) Systemtime() bool           { return p.Enable("SeSystemtimePrivilege") }
func (p *Privilege) ProfileSingleProcess() bool { return p.Enable("SeProfileSingleProcessPrivilege") }

func (p *Privilege) IncreaseBasePriority() bool { return p.Enable("SeIncreaseBasePriorityPrivilege") }
func (p *Privilege) CreatePagefile() bool       { return p.Enable("SeCreatePagefilePrivilege") }
func (p *Privilege) CreatePermanent() bool      { return p.Enable("SeCreatePermanentPrivilege") }
func (p *Privilege) Backup() bool               { return p.Enable("SeBackupPrivilege") }
func (p *Privilege) Restore() bool              { return p.Enable("SeRestorePrivilege") }
func (p *Privilege) Shutdown() bool             { return p.Enable("SeShutdownPrivilege") }
func (p *Privilege) Debug() bool                { return p.Enable("SeDebugPrivilege") }
func (p *Privilege) Audit() bool                { return p.Enable("SeAuditPrivilege") }
func (p *Privilege) SystemEnvironment() bool    { return p.Enable("SeSystemEnvironmentPrivilege") }
func (p *Privilege) ChangeNotify() bool         { return p.Enable("SeChangeNotifyPrivilege") }
func (p *Privilege) RemoteShutdown() bool       { return p.Enable("SeRemoteShutdownPrivilege") }
func (p *Privilege) Undock() bool               { return p.Enable("SeUndockPrivilege") }
func (p *Privilege) SyncAgent() bool            { return p.Enable("SeSyncAgentPrivilege") }
func (p *Privilege) EnableDelegation() bool     { return p.Enable("SeEnableDelegationPrivilege") }
func (p *Privilege) ManageVolume() bool         { return p.Enable("SeManageVolumePrivilege") }
func (p *Privilege) Impersonate() bool          { return p.Enable("SeImpersonatePrivilege") }
func (p *Privilege) CreateGlobal() bool         { return p.Enable("SeCreateGlobalPrivilege") }
