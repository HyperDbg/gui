package exception

import (
	"fmt"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/golibrary/std/safemap"
)

type ExceptionType int

const (
	ExceptionAccessViolation ExceptionType = iota
	ExceptionArrayBounds
	ExceptionBreakpoint
	ExceptionDatatypeMisalignment
	ExceptionFltDenormalOperand
	ExceptionFltDivideByZero
	ExceptionFltInexactResult
	ExceptionFltInvalidOperation
	ExceptionFltOverflow
	ExceptionFltStackCheck
	ExceptionFltUnderflow
	ExceptionIllegalInstruction
	ExceptionInPageError
	ExceptionIntDivideByZero
	ExceptionIntOverflow
	ExceptionInvalidDisposition
	ExceptionNoncontinuableException
	ExceptionPrivInstruction
	ExceptionSingleStep
	ExceptionStackOverflow
	ExceptionGuardPage
	ExceptionInvalidHandle
	ExceptionDllNotFound
	ExceptionDllInitFailed
	ExceptionRpcServerUnavailable
	ExceptionRpcServerInvalid
	ExceptionOther
)

type ExceptionChance int

const (
	ExceptionFirstChance ExceptionChance = iota
	ExceptionSecondChance
	ExceptionAll
)

type ExceptionRule struct {
	ExceptionCode uint32
	ExceptionName string
	Chance        ExceptionChance
	Disabled      bool
	Module        string
	Address       uint64
	Action        string
	Notes         string
}

type ExceptionLog struct {
	Timestamp     int64
	ThreadId      uint32
	ExceptionCode uint32
	ExceptionName string
	ExceptionAddr uint64
	AccessAddr    uint64
	AccessMode    uint32
	Chance        ExceptionChance
	Handled       bool
	Notes         string
}

type Manager struct {
	rules *safemap.M[uint32, *ExceptionRule]
	logs  []*ExceptionLog
}

func New() api.Interface {
	return &Manager{
		rules: safemap.New[uint32, *ExceptionRule](),
		logs:  make([]*ExceptionLog, 0),
	}
}

func (m *Manager) AddRule(exceptionCode uint32, exceptionName string, chance ExceptionChance, disabled bool, module string, address uint64, action, notes string) error {
	if exceptionName == "" {
		return fmt.Errorf("exception name cannot be empty")
	}

	m.rules.Update(exceptionCode, &ExceptionRule{
		ExceptionCode: exceptionCode,
		ExceptionName: exceptionName,
		Chance:        chance,
		Disabled:      disabled,
		Module:        module,
		Address:       address,
		Action:        action,
		Notes:         notes,
	})

	return nil
}

func (m *Manager) GetRule(exceptionCode uint32) *ExceptionRule {
	rule, _ := m.rules.Get(exceptionCode)
	return rule
}

func (m *Manager) GetRuleByName(exceptionName string) *ExceptionRule {
	for _, rule := range m.rules.Range() {
		if rule.ExceptionName == exceptionName {
			return rule
		}
	}
	return nil
}

func (m *Manager) DeleteRule(exceptionCode uint32) {
	m.rules.Delete(exceptionCode)
}

func (m *Manager) GetAllRules() []*ExceptionRule {
	result := make([]*ExceptionRule, 0)
	for _, rule := range m.rules.Range() {
		result = append(result, rule)
	}
	return result
}

func (m *Manager) GetEnabledRules() []*ExceptionRule {
	result := make([]*ExceptionRule, 0)
	for _, rule := range m.rules.Range() {
		if !rule.Disabled {
			result = append(result, rule)
		}
	}
	return result
}

func (m *Manager) GetDisabledRules() []*ExceptionRule {
	result := make([]*ExceptionRule, 0)
	for _, rule := range m.rules.Range() {
		if rule.Disabled {
			result = append(result, rule)
		}
	}
	return result
}

func (m *Manager) ClearRules() {
	m.rules.Reset()
}

func (m *Manager) HasRule(exceptionCode uint32) bool {
	_, exists := m.rules.Get(exceptionCode)
	return exists
}

func (m *Manager) SetRuleDisabled(exceptionCode uint32, disabled bool) {
	if rule, exists := m.rules.Get(exceptionCode); exists {
		rule.Disabled = disabled
	}
}

func (m *Manager) SetRuleChance(exceptionCode uint32, chance ExceptionChance) {
	if rule, exists := m.rules.Get(exceptionCode); exists {
		rule.Chance = chance
	}
}

func (m *Manager) SetRuleAction(exceptionCode uint32, action string) {
	if rule, exists := m.rules.Get(exceptionCode); exists {
		rule.Action = action
	}
}

func (m *Manager) SetRuleModule(exceptionCode uint32, module string) {
	if rule, exists := m.rules.Get(exceptionCode); exists {
		rule.Module = module
	}
}

func (m *Manager) SetRuleAddress(exceptionCode uint32, address uint64) {
	if rule, exists := m.rules.Get(exceptionCode); exists {
		rule.Address = address
	}
}

func (m *Manager) SetRuleNotes(exceptionCode uint32, notes string) {
	if rule, exists := m.rules.Get(exceptionCode); exists {
		rule.Notes = notes
	}
}

func (m *Manager) GetRuleCount() int {
	count := 0
	for range m.rules.Range() {
		count++
	}
	return count
}

func (m *Manager) GetEnabledRuleCount() int {
	count := 0
	for _, rule := range m.rules.Range() {
		if !rule.Disabled {
			count++
		}
	}
	return count
}

func (m *Manager) GetDisabledRuleCount() int {
	count := 0
	for _, rule := range m.rules.Range() {
		if rule.Disabled {
			count++
		}
	}
	return count
}

func (m *Manager) AddLog(timestamp int64, threadId uint32, exceptionCode uint32, exceptionName string, exceptionAddr, accessAddr uint64, accessMode uint32, chance ExceptionChance, handled bool, notes string) {
	log := &ExceptionLog{
		Timestamp:     timestamp,
		ThreadId:      threadId,
		ExceptionCode: exceptionCode,
		ExceptionName: exceptionName,
		ExceptionAddr: exceptionAddr,
		AccessAddr:    accessAddr,
		AccessMode:    accessMode,
		Chance:        chance,
		Handled:       handled,
		Notes:         notes,
	}

	m.logs = append(m.logs, log)
}

func (m *Manager) GetLogs() []*ExceptionLog {
	result := make([]*ExceptionLog, len(m.logs))
	copy(result, m.logs)
	return result
}

func (m *Manager) GetLogsByThread(threadId uint32) []*ExceptionLog {
	result := make([]*ExceptionLog, 0)
	for _, log := range m.logs {
		if log.ThreadId == threadId {
			result = append(result, log)
		}
	}
	return result
}

func (m *Manager) GetLogsByExceptionCode(exceptionCode uint32) []*ExceptionLog {
	result := make([]*ExceptionLog, 0)
	for _, log := range m.logs {
		if log.ExceptionCode == exceptionCode {
			result = append(result, log)
		}
	}
	return result
}

func (m *Manager) GetLogsByChance(chance ExceptionChance) []*ExceptionLog {
	result := make([]*ExceptionLog, 0)
	for _, log := range m.logs {
		if log.Chance == chance {
			result = append(result, log)
		}
	}
	return result
}

func (m *Manager) GetUnhandledLogs() []*ExceptionLog {
	result := make([]*ExceptionLog, 0)
	for _, log := range m.logs {
		if !log.Handled {
			result = append(result, log)
		}
	}
	return result
}

func (m *Manager) ClearLogs() {
	m.logs = make([]*ExceptionLog, 0)
}

func (m *Manager) GetLogCount() int {
	return len(m.logs)
}

func (m *Manager) ShouldBreakOnException(exceptionCode uint32, chance ExceptionChance) bool {
	rule, exists := m.rules.Get(exceptionCode)
	if !exists {
		return false
	}

	if rule.Disabled {
		return false
	}

	if rule.Chance == ExceptionAll {
		return true
	}

	return rule.Chance == chance
}

func (m *Manager) GetExceptionName(exceptionCode uint32) string {
	if rule, exists := m.rules.Get(exceptionCode); exists {
		return rule.ExceptionName
	}

	switch exceptionCode {
	case 0xC0000005:
		return "EXCEPTION_ACCESS_VIOLATION"
	case 0xC000008C:
		return "EXCEPTION_ARRAY_BOUNDS_EXCEEDED"
	case 0x80000003:
		return "EXCEPTION_BREAKPOINT"
	case 0x80000002:
		return "EXCEPTION_DATATYPE_MISALIGNMENT"
	case 0xC000008D:
		return "EXCEPTION_FLT_DENORMAL_OPERAND"
	case 0xC000008E:
		return "EXCEPTION_FLT_DIVIDE_BY_ZERO"
	case 0xC000008F:
		return "EXCEPTION_FLT_INEXACT_RESULT"
	case 0xC0000090:
		return "EXCEPTION_FLT_INVALID_OPERATION"
	case 0xC0000091:
		return "EXCEPTION_FLT_OVERFLOW"
	case 0xC0000092:
		return "EXCEPTION_FLT_STACK_CHECK"
	case 0xC0000093:
		return "EXCEPTION_FLT_UNDERFLOW"
	case 0xC000001D:
		return "EXCEPTION_ILLEGAL_INSTRUCTION"
	case 0xC0000006:
		return "EXCEPTION_IN_PAGE_ERROR"
	case 0xC0000094:
		return "EXCEPTION_INT_DIVIDE_BY_ZERO"
	case 0xC0000095:
		return "EXCEPTION_INT_OVERFLOW"
	case 0xC0000026:
		return "EXCEPTION_INVALID_DISPOSITION"
	case 0xC0000025:
		return "EXCEPTION_NONCONTINUABLE_EXCEPTION"
	case 0xC0000096:
		return "EXCEPTION_PRIV_INSTRUCTION"
	case 0x80000004:
		return "EXCEPTION_SINGLE_STEP"
	case 0xC00000FD:
		return "EXCEPTION_STACK_OVERFLOW"
	case 0x80000001:
		return "EXCEPTION_GUARD_PAGE"
	case 0xC0000008:
		return "EXCEPTION_INVALID_HANDLE"
	case 0xC0000135:
		return "EXCEPTION_DLL_NOT_FOUND"
	case 0xC0000142:
		return "EXCEPTION_DLL_INIT_FAILED"
	default:
		return fmt.Sprintf("Unknown Exception (0x%X)", exceptionCode)
	}
}

func (m *Manager) GetExceptionCode(exceptionName string) uint32 {
	for code, rule := range m.rules.Range() {
		if rule.ExceptionName == exceptionName {
			return code
		}
	}

	switch exceptionName {
	case "EXCEPTION_ACCESS_VIOLATION":
		return 0xC0000005
	case "EXCEPTION_ARRAY_BOUNDS_EXCEEDED":
		return 0xC000008C
	case "EXCEPTION_BREAKPOINT":
		return 0x80000003
	case "EXCEPTION_DATATYPE_MISALIGNMENT":
		return 0x80000002
	case "EXCEPTION_FLT_DENORMAL_OPERAND":
		return 0xC000008D
	case "EXCEPTION_FLT_DIVIDE_BY_ZERO":
		return 0xC000008E
	case "EXCEPTION_FLT_INEXACT_RESULT":
		return 0xC000008F
	case "EXCEPTION_FLT_INVALID_OPERATION":
		return 0xC0000090
	case "EXCEPTION_FLT_OVERFLOW":
		return 0xC0000091
	case "EXCEPTION_FLT_STACK_CHECK":
		return 0xC0000092
	case "EXCEPTION_FLT_UNDERFLOW":
		return 0xC0000093
	case "EXCEPTION_ILLEGAL_INSTRUCTION":
		return 0xC000001D
	case "EXCEPTION_IN_PAGE_ERROR":
		return 0xC0000006
	case "EXCEPTION_INT_DIVIDE_BY_ZERO":
		return 0xC0000094
	case "EXCEPTION_INT_OVERFLOW":
		return 0xC0000095
	case "EXCEPTION_INVALID_DISPOSITION":
		return 0xC0000026
	case "EXCEPTION_NONCONTINUABLE_EXCEPTION":
		return 0xC0000025
	case "EXCEPTION_PRIV_INSTRUCTION":
		return 0xC0000096
	case "EXCEPTION_SINGLE_STEP":
		return 0x80000004
	case "EXCEPTION_STACK_OVERFLOW":
		return 0xC00000FD
	case "EXCEPTION_GUARD_PAGE":
		return 0x80000001
	case "EXCEPTION_INVALID_HANDLE":
		return 0xC0000008
	case "EXCEPTION_DLL_NOT_FOUND":
		return 0xC0000135
	case "EXCEPTION_DLL_INIT_FAILED":
		return 0xC0000142
	default:
		return 0
	}
}

func (m *Manager) GetMostRecentException() *ExceptionLog {
	if len(m.logs) == 0 {
		return nil
	}
	return m.logs[len(m.logs)-1]
}

func (m *Manager) GetExceptionCountByCode(exceptionCode uint32) int {
	count := 0
	for _, log := range m.logs {
		if log.ExceptionCode == exceptionCode {
			count++
		}
	}
	return count
}

func (m *Manager) GetExceptionCountByThread(threadId uint32) int {
	count := 0
	for _, log := range m.logs {
		if log.ThreadId == threadId {
			count++
		}
	}
	return count
}

func (m *Manager) Layout() layout.Widget {
	return nil
}

func (m *Manager) Update() error {
	return nil
}

func (m *Manager) Clear() {
	m.rules.Reset()
	m.logs = make([]*ExceptionLog, 0)
}

func (m *Manager) Self() any {
	return m
}
