// Package scriptengine wraps the go-bridge compiler (go-bridge/ast) and the
// HyperDbg kernel IOCTL protocol (IOCTL_DEBUGGER_REGISTER_EVENT +
// IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT) so callers can register a Go callback
// against an EPT hook (or any other VMM event) without dealing with binary
// AST encoding, struct packing or IOCTL plumbing themselves.
//
// This package is the Go counterpart to the C++ wrapper at
// HyperDbg/hyperdbg/libhyperdbg/code/debugger/script-engine/script-engine-wrapper.cpp.
// The C++ wrapper drives the old .ds bytecode VM (ScriptEngineParse +
// ScriptEngineExecute + SYMBOL_BUFFER); this package drives the new Go-subset
// interpreter (go-bridge/ast → binary AST → hyperkd/code/go-interp). Both
// paths coexist in the kernel (distinguished by HOOK_FLAG_GO_AST).
package scriptengine

import "errors"

// 包装层错误哨兵。Compile / BuildAction / BuildEvent / RegisterHook 在
// 失败时用 fmt.Errorf("%w: ...", sentinel) 包装底层原因，调用方可以用
// errors.Is(err, ErrScriptCompile) 之类的判定来区分错误类别，无需依赖
// 字符串匹配。底层错误（go-bridge 的 *protocol.ValidationError、comm 包
// 返回的 windows 错误等）通过 %v 嵌入消息，保留可读的根因。
var (
	// ErrScriptCompile 在脚本编译阶段失败时返回。通常是 go-bridge 编码器
	// 报告的"非 ValidationError"路径（例如编码器返回空负载、内部 panic
	// 透传等）。源码语法错误与子集违规走 ErrSubnetViolation。
	ErrScriptCompile = errors.New("scriptengine: script compile failed")

	// ErrSubnetViolation 在 Go 源码无法通过子集校验时返回。包装的底层
	// 错误通常是 *protocol.ValidationError，其 Pos 字段携带源码位置
	// （file.go:line:col），Reason 字段解释违规原因（例如 "interface
	// types not supported"、"goroutine not supported"、
	// "unsupported statement: *ast.GoStmt" 等）。命名沿用任务规范。
	ErrSubnetViolation = errors.New("scriptengine: Go subset violation")

	// ErrActionBuild 在 BuildAction 因参数非法（空脚本配 RunScript、
	// 未知动作类型、脚本超长等）而无法构造 DEBUGGER_GENERAL_ACTION 时
	// 返回。
	ErrActionBuild = errors.New("scriptengine: action build failed")

	// ErrEventBuild 在 BuildEvent 因参数非法而无法构造
	// DEBUGGER_GENERAL_EVENT_DETAIL 时返回。当前 BuildEvent 接受任何
	// VMM_EVENT_TYPE_ENUM 值（由内核负责最终校验），故此错误主要为
	// 未来扩展预留。
	ErrEventBuild = errors.New("scriptengine: event build failed")

	// ErrIoctlRegister 在 RegisterHook 与驱动通信失败、或驱动拒绝
	// 事件 / 动作注册时返回。底层原因通过 %v 嵌入（Windows 错误码、
	// DEBUGGER_EVENT_AND_ACTION_RESULT.Error 字段值等）。
	ErrIoctlRegister = errors.New("scriptengine: IOCTL register failed")
)
