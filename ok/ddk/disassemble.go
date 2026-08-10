package ddk

import (
	"fmt"
	"iter"
	"strings"

	"golang.org/x/arch/x86/x86asm"
)

// Disassemble 是一个函数，用于将字节码代码反汇编为x86汇编指令
// 参数:
//
//	code - 要反汇编的字节码切片
//
// 返回值:
//
//	[]x86asm.Inst - 反汇编后的x86汇编指令切片
func Disassemble(code []byte) []x86asm.Inst {
	var insts []x86asm.Inst // 用于存储反汇编后的指令切片
	// 遍历DisassembleSeq2函数返回的指令序列
	for _, inst := range DisassembleSeq2(code, 0, nil) {
		insts = append(insts, inst)
		if inst.Op == x86asm.RET {
			break
		}
	}
	return insts
}

// DisassembleToString 将字节数码转换为可读的汇编字符串
// 参数:
//
//	code: 要反汇编的字节数组
//	address: 反汇编的起始地址
//
// 返回值:
//
//	string: 包含反汇编结果的字符串，每行一条汇编指令
func DisassembleToString(code []byte, address uint64) string {
	var sb strings.Builder // 使用strings.Builder高效构建字符串
	// 遍历DisassembleSeq2函数返回的反汇编序列
	for line := range DisassembleSeq2(code, int(address), nil) {
		sb.WriteString(line + "\n") // 将每条汇编指令添加到字符串中，并换行
	}
	return sb.String() // 返回完整的反汇编字符串
}

// DisassembleSeq2 是一个函数，用于将x86汇编代码反汇编为字符串和指令的迭代序列
// 参数:
//   - code: 要反汇编的字节切片
//   - baseAddr: 基址地址
//   - sym: 符号查找函数，用于获取符号信息
//
// 返回值:
//   - 返回一个迭代器函数，该函数生成字符串格式的反汇编结果和对应的指令对象
func DisassembleSeq2(code []byte, baseAddr int, sym x86asm.SymLookup) iter.Seq2[string, x86asm.Inst] {
	return func(yield func(string, x86asm.Inst) bool) {
		// 根据基址地址确定CPU模式（32位或64位）
		mode := 32
		if uint64(baseAddr) > 0xFFFFFFFF {
			mode = 64
		}

		off := 0 // 当前代码偏移量
		// 遍历整个代码字节切片
		for off < len(code) {
			// 计算当前指令的RIP值
			rip := uint64(baseAddr + off)
			// 根据当前模式解码指令
			inst, err := x86asm.Decode(code[off:], mode)

			var opHex, asm string // 操作码的十六进制表示和汇编字符串
			if err != nil || inst.Len <= 0 {
				// 如果解码失败或指令长度无效，则按字节处理
				b := code[off]
				opHex = fmt.Sprintf("%02X ", b)
				asm = fmt.Sprintf(".byte %02X", b)
				off++
			} else {
				// x86原生 Sym 符号格式化函数
				for _, b := range code[off : off+inst.Len] {
					opHex += fmt.Sprintf("%02X ", b)
				}
				asm = x86asm.IntelSyntax(inst, rip, sym)
				off += inst.Len
			}

			// x86最长15字节 opcode 固定对齐，Maximum instruction size is 15 bytes.
			opAlign := fmt.Sprintf("%-45s", opHex)

			var addrFmt string
			if uint64(baseAddr) <= 0xFFFFFFFF {
				addrFmt = "%08X"
			} else {
				addrFmt = "%016X"
			}

			line := fmt.Sprintf(addrFmt+": %s%s", rip, opAlign, asm)
			if !yield(line, inst) {
				return
			}
			if inst.Op == x86asm.RET {
				return
			}
		}
	}
}
