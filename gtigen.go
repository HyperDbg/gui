// Code generated by "core generate"; DO NOT EDIT.

package main

import (
	"cogentcore.org/core/gti"
)

var _ = gti.AddType(&gti.Type{Name: "main.Disassembly", IDName: "disassembly", Directives: []gti.Directive{{Tool: "type FastCall struct { //gti", Directive: "add"}, {Tool: "gti", Directive: "add"}}, Fields: []gti.Field{{Name: "Icon"}, {Name: "Address"}, {Name: "Opcode"}, {Name: "Instruction"}, {Name: "Comment"}}})

var _ = gti.AddType(&gti.Type{Name: "main.Stack", IDName: "stack", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Fields: []gti.Field{{Name: "Address"}, {Name: "Data"}, {Name: "Context"}}})

var _ = gti.AddType(&gti.Type{Name: "main.Register", IDName: "register", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Fields: []gti.Field{{Name: "RAX"}, {Name: "RBX"}, {Name: "RCX"}, {Name: "RDX"}, {Name: "RBP"}, {Name: "RSP"}, {Name: "RSI"}, {Name: "RDI"}, {Name: "R8"}, {Name: "R9"}, {Name: "R10"}, {Name: "R11"}, {Name: "R12"}, {Name: "R13"}, {Name: "R14"}, {Name: "R15"}, {Name: "RIP"}, {Name: "RFLAGS"}, {Name: "ZF"}, {Name: "OF"}, {Name: "CF"}, {Name: "PF"}, {Name: "SF"}, {Name: "TF"}, {Name: "AF"}, {Name: "DF"}, {Name: "IF"}, {Name: "LastError"}, {Name: "LastStatus"}, {Name: "GS"}, {Name: "ES"}, {Name: "CS"}, {Name: "FS"}, {Name: "DS"}, {Name: "SS"}, {Name: "ST0"}, {Name: "ST1"}, {Name: "ST2"}, {Name: "ST3"}, {Name: "ST4"}, {Name: "ST5"}, {Name: "ST6"}, {Name: "ST7"}, {Name: "x87TagWord"}, {Name: "x87ControlWord"}, {Name: "x87StatusWord"}, {Name: "x87TW_0"}, {Name: "x87TW_1"}, {Name: "x87TW_2"}, {Name: "x87TW_3"}, {Name: "x87TW_4"}, {Name: "x87TW_5"}, {Name: "x87TW_6"}, {Name: "x87TW_7"}, {Name: "x87SW_B"}, {Name: "x87SW_C3"}, {Name: "x87SW_TOP"}, {Name: "x87SW_C2"}, {Name: "x87SW_C1"}, {Name: "x87SW_O"}, {Name: "x87SW_ES"}, {Name: "x87SW_SF"}, {Name: "x87SW_P"}, {Name: "x87SW_U"}, {Name: "x87SW_Z"}, {Name: "x87SW_D"}, {Name: "x87SW_I"}, {Name: "x87SW_C0"}, {Name: "x87CW_IC"}, {Name: "x87CW_RC"}, {Name: "x87CW_PC"}, {Name: "x87CW_PM"}, {Name: "x87CW_UM"}, {Name: "x87CW_OM"}, {Name: "x87CW_ZM"}, {Name: "x87CW_DM"}, {Name: "x87CW_IM"}, {Name: "MxCsr"}, {Name: "MxCsr_FZ"}, {Name: "MxCsr_PM"}, {Name: "MxCsr_UM"}, {Name: "MxCsr_OM"}, {Name: "MxCsr_ZM"}, {Name: "MxCsr_IM"}, {Name: "MxCsr_DM"}, {Name: "MxCsr_D"}, {Name: "MxCsr_PE"}, {Name: "MxCsr_UE"}, {Name: "MxCsr_OE"}, {Name: "MxCsr_ZE"}, {Name: "MxCsr_DE"}, {Name: "MxCsr_IE"}, {Name: "MxCsr_RC"}, {Name: "XMM0"}, {Name: "XMM1"}, {Name: "XMM2"}, {Name: "XMM3"}, {Name: "XMM4"}, {Name: "XMM5"}, {Name: "XMM6"}, {Name: "XMM7"}, {Name: "XMM8"}, {Name: "XMM9"}, {Name: "XMM10"}, {Name: "XMM11"}, {Name: "XMM12"}, {Name: "XMM13"}, {Name: "XMM14"}, {Name: "XMM15"}, {Name: "YMM0"}, {Name: "YMM1"}, {Name: "YMM2"}, {Name: "YMM3"}, {Name: "YMM4"}, {Name: "YMM5"}, {Name: "YMM6"}, {Name: "YMM7"}, {Name: "YMM8"}, {Name: "YMM9"}, {Name: "YMM10"}, {Name: "YMM11"}, {Name: "YMM12"}, {Name: "YMM13"}, {Name: "YMM14"}, {Name: "YMM15"}, {Name: "DR0"}, {Name: "DR1"}, {Name: "DR2"}, {Name: "DR3"}, {Name: "DR6"}, {Name: "DR7"}}})
