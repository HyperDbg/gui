package misc

import (
	"fmt"
	"strings"

	"golang.org/x/arch/x86/x86asm"
)

type DisasmMode int

const (
	Mode64 DisasmMode = iota
	Mode32
)

type Instruction struct {
	Length    uint32
	Opcode    string
	Operands  string
	Address   uint64
	Bytes     []byte
	IsBranch  bool
	IsCall    bool
	IsRet     bool
	IsCondJmp bool
}

func DisassembleOne(code []byte, baseAddr uint64, mode DisasmMode) (*Instruction, error) {
	var archMode int
	if mode == Mode64 {
		archMode = 64
	} else {
		archMode = 32
	}

	inst, err := x86asm.Decode(code, archMode)
	if err != nil {
		return nil, fmt.Errorf("disassemble failed at address 0x%x: %w", baseAddr, err)
	}

	result := &Instruction{
		Length:   uint32(inst.Len),
		Opcode:   inst.Op.String(),
		Operands: x86asm.IntelSyntax(inst, baseAddr, nil),
		Address:  baseAddr,
		Bytes:    make([]byte, inst.Len),
	}
	copy(result.Bytes, code[:inst.Len])

	switch inst.Op {
	case x86asm.CALL:
		result.IsCall = true
	case x86asm.RET:
		result.IsRet = true
	case x86asm.JO, x86asm.JNO, x86asm.JB, x86asm.JAE, x86asm.JE, x86asm.JNE,
		x86asm.JBE, x86asm.JA, x86asm.JS, x86asm.JNS, x86asm.JP, x86asm.JNP,
		x86asm.JL, x86asm.JGE, x86asm.JLE, x86asm.JG:
		result.IsCondJmp = true
		result.IsBranch = true
	case x86asm.JMP:
		result.IsBranch = true
	}

	return result, nil
}

func DisassembleBuffer(code []byte, baseAddr uint64, mode DisasmMode, maxCount int) ([]*Instruction, error) {
	var instructions []*Instruction
	offset := 0

	for i := 0; i < maxCount && offset < len(code); i++ {
		inst, err := DisassembleOne(code[offset:], baseAddr+uint64(offset), mode)
		if err != nil {
			break
		}
		instructions = append(instructions, inst)
		offset += int(inst.Length)
	}

	return instructions, nil
}

func FormatDisassembly(inst *Instruction) string {
	bytesStr := formatBytes(inst.Bytes)
	return fmt.Sprintf("%s: %-30s %s", bytesStr, inst.Opcode, inst.Operands)
}

func formatBytes(bytes []byte) string {
	var parts []string
	for _, b := range bytes {
		parts = append(parts, fmt.Sprintf("%02x", b))
	}
	return strings.Join(parts, " ")
}

func GetInstructionLength(code []byte, mode DisasmMode) (uint32, error) {
	inst, err := DisassembleOne(code, 0, mode)
	if err != nil {
		return 0, err
	}
	return inst.Length, nil
}

func IsCallOrBranch(inst *Instruction) bool {
	return inst.IsCall || inst.IsBranch
}

func IsConditionalJump(inst *Instruction) bool {
	return inst.IsCondJmp
}

func GetBranchTarget(inst *Instruction) (uint64, bool) {
	if !inst.IsBranch {
		return 0, false
	}
	return inst.Address + uint64(inst.Length), true
}
