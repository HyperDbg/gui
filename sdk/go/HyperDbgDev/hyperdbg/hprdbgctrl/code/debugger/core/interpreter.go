package core

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\core\\interpreter.cpp"
var interpreterBuf string

type (
	Interpreter interface {
		HyperDbgInterpreter() (ok bool)
		InterpreterRemoveComments() (ok bool)
		HyperDbgShowSignature() (ok bool)
		HyperDbgCheckMultilineCommand() (ok bool)
		HyperDbgContinuePreviousCommand() (ok bool)
		GetCommandAttributes() (ok bool)
		InitializeDebugger() (ok bool)
		InitializeCommandsDictionary() (ok bool) //todo change to map? all command  in here
	}
	interpreter struct{}
)

func (i *interpreter) HyperDbgInterpreter() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (i *interpreter) InterpreterRemoveComments() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (i *interpreter) HyperDbgShowSignature() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (i *interpreter) HyperDbgCheckMultilineCommand() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (i *interpreter) HyperDbgContinuePreviousCommand() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (i *interpreter) GetCommandAttributes() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (i *interpreter) InitializeDebugger() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (i *interpreter) InitializeCommandsDictionary() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func Newinterpreter() Interpreter { return &interpreter{} }

var name = map[int]string{
	0:   ".help",
	1:   ".hh",
	2:   "help",
	3:   "clear",
	4:   ".cls",
	5:   "cls",
	6:   ".connect",
	7:   "connect",
	8:   ".listen",
	9:   "listen",
	10:  "g",
	11:  "go",
	12:  ".attach",
	13:  "attach",
	14:  ".detach",
	15:  "detach",
	16:  ".start",
	17:  "start",
	18:  ".restart",
	19:  "restart",
	20:  ".switch",
	21:  "switch",
	22:  ".kill",
	23:  "kill",
	24:  ".process",
	25:  ".process2",
	26:  ".thread",
	27:  ".thread2",
	28:  "sleep",
	29:  "event",
	30:  "events",
	31:  "setting",
	32:  "settings",
	33:  ".disconnect",
	34:  "disconnect",
	35:  ".debug",
	36:  "debug",
	37:  ".status",
	38:  "status",
	39:  "load",
	40:  "exit",
	41:  ".exit",
	42:  "flush",
	43:  "pause",
	44:  "unload",
	45:  ".script",
	46:  "script",
	47:  "output",
	48:  "print",
	49:  "?",
	50:  "eval",
	51:  "evaluate",
	52:  ".logopen",
	53:  ".logclose",
	54:  "test",
	55:  "cpu",
	56:  "wrmsr",
	57:  "rdmsr",
	58:  "!va2pa",
	59:  "!pa2va",
	60:  ".formats",
	61:  ".format",
	62:  "!pte",
	63:  "~",
	64:  "core",
	65:  "!monitor",
	66:  "!vmcall",
	67:  "!epthook",
	68:  "bh",
	69:  "bp",
	70:  "bl",
	71:  "be",
	72:  "bd",
	73:  "bc",
	74:  "!epthook2",
	75:  "!cpuid",
	76:  "!msrread",
	77:  "!msread",
	78:  "!msrwrite",
	79:  "!tsc",
	80:  "!pmc",
	81:  "!crwrite",
	82:  "!dr",
	83:  "!ioin",
	84:  "!ioout",
	85:  "!exception",
	86:  "!interrupt",
	87:  "!syscall",
	88:  "!syscall2",
	89:  "!sysret",
	90:  "!sysret2",
	91:  "!hide",
	92:  "!unhide",
	93:  "!measure",
	94:  "lm",
	95:  "p",
	96:  "pr",
	97:  "t",
	98:  "tr",
	99:  "i",
	100: "ir",
	101: "db",
	102: "dc",
	103: "dd",
	104: "dq",
	105: "!db",
	106: "!dc",
	107: "!dd",
	108: "!dq",
	109: "!u",
	110: "u",
	111: "!u2",
	112: "u2",
	113: "eb",
	114: "ed",
	115: "eq",
	116: "!eb",
	117: "!ed",
	118: "!eq",
	119: "sb",
	120: "sd",
	121: "sq",
	122: "!sb",
	123: "!sd",
	124: "!sq",
	125: "r",
	126: ".sympath",
	127: "sympath",
	128: ".sym",
	129: "sym",
	130: "x",
	131: "prealloc",
	132: "preallocate",
	133: "k",
	134: "kd",
	135: "kq",
	136: "dt",
	137: "!dt",
	138: "struct",
	139: "structure",
	140: ".pe",
}
