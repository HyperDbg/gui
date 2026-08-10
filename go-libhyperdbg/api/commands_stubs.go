// Package api — commands_stubs.go
//
// 本文件集中定义所有命令 typed API 共享的"尚未实装"错误。
//
// 设计意图：
//
//	debugger/commands 下 4 个包（meta/debugging/extension/hwdbg）共注册了
//	约 100 个命令。其中只有少数（如 .connect/load/g/pause/.logopen/.start/
//	!epthook）在 core 层有真正实装，其余都是 stubs。
//
//	字符串命令路径（Exec("!sysret hook")）的问题是参数没有编译期类型检查。
//	为了让调用方走 typed 路径（dbg.SysretHook(ctx, src)），每个命令都必须
//	在 api 层有对应的方法签名——即便方法体只是返回 ErrCommandNotImplemented。
//	这样：
//	  1. 调用方有类型检查，参数错配在编译期暴露
//	  2. 未实装的命令调用时返回明确的错误，而不是 silently 打印 stub 信息
//	  3. 未来逐步实装时，签名已就位，只需替换方法体
//
// 各命令的 typed API 分别在：
//
//	commands_meta.go       — meta 命令（.connect/load/status/exit/attach/...）
//	commands_debugging.go  — debugging 命令（bp/k/r/dt/lm/sleep/...）
//	commands_extension.go  — extension 命令（!epthook/!sysret/!cpuid/...）
//	commands_hwdbg.go      — hwdbg 命令（hw/hw_clk）
package api

import "errors"

// ErrCommandNotImplemented 表示该 typed API 方法对应的命令尚未在 core 层实装。
// 调用方可以通过 errors.Is(err, ErrCommandNotImplemented) 判断。
//
// 区别于运行时错误（如 IOCTL 失败、驱动未加载），本错误明确表示"该功能
// 还没写"——调用方应该在编译期或测试期就发现并避开，而不是依赖运行时检查。
var ErrCommandNotImplemented = errors.New("command not yet implemented (typed API stub)")
