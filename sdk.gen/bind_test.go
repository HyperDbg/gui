package sdk_gen

import (
	"path/filepath"
	"strings"
	"testing"
	"unicode"

	"github.com/ddkwork/app/bindgen/clang"
	"github.com/ddkwork/app/bindgen/gengo"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"github.com/ddkwork/golibrary/stream/orderedmap"
)

func TestMergeHeader(t *testing.T) {
	Headers := mylog.Check2(filepath.Glob("SDK/Headers/*.h"))
	User := mylog.Check2(filepath.Glob("SDK/Imports/User/*.h"))
	g := stream.NewGeneratedFile()
	bugfix = strings.TrimPrefix(bugfix, "\n")
	g.P("//bugfix.h")
	g.P(bugfix)
	g.P()
	mylog.Trace("merge", "bugfix.h")

	fnDo := func(path string) {
		g.P("//" + path)
		g.P(stream.NewBuffer(path))
		g.P()
		mylog.Trace("merge", path)
	}
	for _, s := range Headers {
		if stream.BaseName(s) == "Assertions" {
			continue
		}
		fnDo(s)
	}
	for _, s := range User {
		fnDo(s)
	}
	stream.WriteBinaryFile("merged_headers.h", g.Buffer)
}

func TestBindMacros(t *testing.T) {
	headerFile := "merged_headers.h"
	macros := extractMacros(stream.NewBuffer(headerFile).ToLines())
	// println(macros.GoString())
	// return
	mylog.Trace("number of macros", macros.Len())

	var (
		enumDebuggers = orderedmap.New("", false)
		enumIoctls    = orderedmap.New("", false)
	)

	for _, p := range macros.List() {
		if !m.Contains(p.Key) {
			// mylog.Warning(p.Key, p.Value)
			macros.Delete(p.Key)
			continue
		}
		// mylog.Warning(p.Key, p.Value)
	}

	g := stream.NewGeneratedFile()
	g.P("package main")
	g.P()

	g.P("import \"github.com/winlabs/gowin32/wrappers\"")
	g.P()

	g.P("var (")
	for _, p := range macros.List() {
		p.Value = strings.TrimSpace(p.Value)
		if strings.HasPrefix(p.Value, "sizeof") {
			continue
		}
		if strings.HasSuffix(p.Value, "OPERATION_MANDATORY_DEBUGGEE_BIT") {
			continue
		}

		if strings.Contains(p.Value, "sizeof") {
			continue
		}
		if strings.Contains(p.Value, "TOP_LEVEL_DRIVERS_VMCALL_STARTING_NUMBER") {
			continue
		}

		p.Value = strings.ReplaceAll(p.Value, "\\", "")
		p.Value = strings.Replace(p.Value, "6U", "6", 1)
		p.Value = strings.Replace(p.Value, "7U", "7", 1)
		p.Value = strings.Replace(p.Value, "8U", "8", 1)
		p.Value = strings.Replace(p.Value, "9U", "9", 1)
		p.Value = strings.Replace(p.Value, "10U", "10", 1)
		p.Value = strings.Replace(p.Value, "11U", "11", 1)
		p.Value = strings.Replace(p.Value, "12U", "12", 1)
		p.Value = strings.Replace(p.Value, "13U", "13", 1)
		p.Value = strings.Replace(p.Value, "14U", "14", 1)
		p.Value = strings.Replace(p.Value, "15U", "15", 1)
		p.Value = strings.Replace(p.Value, "16U", "16", 1)
		p.Value = strings.Replace(p.Value, "17U", "17", 1)
		p.Value = strings.Replace(p.Value, "NORMAL_PAGE_SIZE", "NormalPageSize", 1)
		p.Value = strings.TrimSuffix(p.Value, "U")
		p.Value = strings.TrimSuffix(p.Value, "ull")

		if len(p.Value) == 0 {
			mylog.Todo(p.Key + " = " + p.Value)
			continue
		}
		if p.Value[0] == '(' && p.Value[len(p.Value)-1] == ')' {
			p.Value = p.Value[1 : len(p.Value)-1]
		}

		if strings.HasPrefix(p.Value, "0x") && !strings.Contains(p.Value, "//") && len(p.Value) > len("0xffffffff") {
			// mylog.Todo(p.Key + " = " + p.Value)
			p.Value = "uint64(" + p.Value + ")"
		}

		key := p.Key
		//if key == "DEBUGGER_OPERATION_WAS_SUCCESSFUL" || strings.HasPrefix(key, "DEBUGGER_ERROR") {
		//	key += " debuggerErrorType"
		//}
		value := p.Value
		if isAlphabetOrUnderscore(value) {
			value = stream.ToCamelUpper(value, false)
		}

		key = stream.ToCamelUpper(key, false)
		switch {
		case strings.HasPrefix(p.Key, "DEBUGGER_ERROR"):
			after, found := strings.CutPrefix(p.Key, "DEBUGGER_ERROR")
			if found {
				key = after
			}
			enumDebuggers.Set(key, true)
		case strings.HasPrefix(p.Key, "IOCTL_"):
			enumIoctls.Set(key, true)
		}

		g.P(stream.ToCamelUpper(key, false) + "=" + value)
		macros.Delete(p.Key)
	}
	g.P(")")

	g.P(`
func CTL_CODE(deviceType, function, method, access uint32) IoctlKind {
	return IoctlKind(((deviceType) << 16) | ((access) << 14) | ((function) << 2) | (method))
}

const (
	FILE_DEVICE_UNKNOWN = wrappers.FILE_DEVICE_UNKNOWN
	METHOD_BUFFERED     = wrappers.METHOD_BUFFERED
	FILE_ANY_ACCESS     = wrappers.FILE_ANY_ACCESS
)
`)

	stream.WriteGoFile("../vars.go", g.Buffer)

	stream.NewGeneratedFile().SetPackageName("main").SetFilePath("../").Enum("debuggerError", enumDebuggers.Keys(), nil)
	stream.NewGeneratedFile().SetPackageName("main").SetFilePath("../").Enum("ioctl", enumIoctls.Keys(), nil)

	for _, p := range macros.List() {
		return
		mylog.Todo(p.Key + " = " + p.Value)
	}
}

// isAlphabetOrUnderscore 检查字符串是否仅由字母或下划线组成
func isAlphabetOrUnderscore(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && r != '_' {
			return false
		}
	}
	return true
}

func TestBindSdk(t *testing.T) {
	TestMergeHeader(t)
	TestBindMacros(t)
	mylog.Call(func() {
		pkg := gengo.NewPackage("libhyperdbg", // 如果改了dll名字也改了呵呵，在不打算内嵌二进制之前但是保留这个操作
			gengo.WithRemovePrefix(
				"hyperdbg_u_",
			),
			gengo.WithInferredMethods([]gengo.MethodInferenceRule{
				//{Name: "ZydisDecoder", Receiver: "Decoder"},
			}),
			gengo.WithForcedSynthetic(
			//"ZydisShortString_",
			//"struct ZydisShortString_",
			),
		)
		mylog.Check(pkg.Transform("libhyperdbg", &clang.Options{
			Sources:          []string{"merged_headers.h"},
			AdditionalParams: []string{},
		}))
		mylog.Check(pkg.WriteToDir("../"))

		// generate bug fix
		fixs := []string{
			`	Uint64     = uint64
	Puint64    = *uint64
	GuestRegs  = GuestRegs`,
			`// @brief struct for extra registers
type GuestExtraRegisters = GuestExtraRegisters`,
		}

		b := stream.NewBuffer("../libhyperdbg.go")
		for _, fix := range fixs {
			b.ReplaceAll(fix, "")
		}
		b.Replace("package libhyperdbg", "package main", 1)
		b.Replace("\nSizeT              = uint64", "", 1)
		b.Replace("\nBool               = int32", "", 1)
		b.Replace(`__imp_hyperdbg_u_continue_debuggee = GengoLibrary.ImportNow("hyperdbg_u_continue_debuggee")`, `	return
	__imp_hyperdbg_u_continue_debuggee = GengoLibrary.ImportNow("hyperdbg_u_continue_debuggee")`, 1)
		b.Replace(`	Bool               = int32
	Long               = int64
	SizeT              = uint64`, `	Long               = int64`, 1)
		b.ReplaceAll(`func ReadVendorString(*Char) {
	bindlib.CCall1(__imp_hyperdbg_u_read_vendor_string.Addr(), bindlib.MarshallSyscall())
}`, `func ReadVendorString(b*Char) {
	bindlib.CCall1(__imp_hyperdbg_u_read_vendor_string.Addr(), bindlib.MarshallSyscall(b))
}`)

		stream.WriteGoFile("../libhyperdbg.go", b)
	})
}

var bugfix = `
#ifndef size_t
#define size_t int
#endif

#ifndef wchar_t
typedef int rune;
#define wchar_t rune
#endif

#ifndef _WINT_T_DEFINED_
 #define _WINT_T_DEFINED_
 typedef int wint_t;
#endif

#undef WCHAR_MIN
#undef WCHAR_MAX
#define WCHAR_MIN   0
#define WCHAR_MAX   65535

typedef int bool ;
typedef long LONG ;
#define PVOID void*
#define HANDLE void*
#define PIRP void*//todo
#define PDEVICE_OBJECT void*//todo
#define PSYMBOL_BUFFER void*//todo
#define PSYMBOL void*//todo
#define MAX_PATH 260
typedef unsigned __int64   SIZE_T,*PSIZE_T;
typedef unsigned __int64   time_t;
typedef unsigned __int64   NTSTATUS;
typedef char *  va_list;

typedef struct _LIST_ENTRY {
  struct _LIST_ENTRY *Flink;
  struct _LIST_ENTRY *Blink;
} LIST_ENTRY, *PLIST_ENTRY, PRLIST_ENTRY;

#ifndef _In_
#define _In_
#endif

#ifndef _Out_
#define _Out_
#endif

#ifndef _Inout_
#define _Inout_
#endif

#ifndef _Out_writes_bytes_
#define _Out_writes_bytes_(x)
#endif

#ifndef _In_reads_
#define _In_reads_(x)
#endif

#ifndef _In_reads_bytes_
#define _In_reads_bytes_(x)
#endif
`

var m = orderedmap.New("", "")

func init() {
	m.Set("PAGE_SIZE", "4096")
	m.Set("NORMAL_PAGE_SIZE", "4096")
	// m.Set("size_t", "int")
	// m.Set("wchar_t", "rune")
	m.Set("WCHAR_MIN", "0")
	m.Set("WCHAR_MAX", "65535")
	// m.Set("PVOID", "void*")
	// m.Set("HANDLE", "void*")
	// m.Set("PIRP", "void*//todo")
	// m.Set("PDEVICE_OBJECT", "void*//todo")
	// m.Set("PSYMBOL_BUFFER", "void*//todo")
	// m.Set("PSYMBOL", "void*//todo")
	m.Set("MAX_PATH", "260")
	// m.Set("_In_", "")
	// m.Set("_Out_", "")
	// m.Set("_Inout_", "")
	// m.Set("_Out_writes_bytes_(x)", "")
	// m.Set("_In_reads_(x)", "")
	// m.Set("_In_reads_bytes_(x)", "")
	// m.Set("VOID", "void")
	m.Set("NULL_ZERO", "0")
	m.Set("NULL64_ZERO", "0ull")
	m.Set("FALSE", "0")
	m.Set("TRUE", "1")
	m.Set("UPPER_56_BITS", "0xffffffffffffff00")
	m.Set("UPPER_48_BITS", "0xffffffffffff0000")
	m.Set("UPPER_32_BITS", "0xffffffff00000000")
	m.Set("LOWER_32_BITS", "0x00000000ffffffff")
	m.Set("LOWER_16_BITS", "0x000000000000ffff")
	m.Set("LOWER_8_BITS", "0x00000000000000ff")
	m.Set("SECOND_LOWER_8_BITS", "0x000000000000ff00")
	m.Set("UPPER_48_BITS_AND_LOWER_8_BITS", "0xffffffffffff00ff")
	m.Set("VERSION_MAJOR", "1")
	m.Set("VERSION_MINOR", "0")
	m.Set("VERSION_PATCH", "0")
	// m.Set("BUILD_YEAR_CH0", "(__DATE__[7])")
	// m.Set("BUILD_YEAR_CH1", "(__DATE__[8])")
	// m.Set("BUILD_YEAR_CH2", "(__DATE__[9])")
	// m.Set("BUILD_YEAR_CH3", "(__DATE__[10])")
	// m.Set("BUILD_MONTH_IS_JAN", "(__DATE__[0] == 'J' && __DATE__[1] == 'a' && __DATE__[2] == 'n')")
	// m.Set("BUILD_MONTH_IS_FEB", "(__DATE__[0] == 'F')")
	// m.Set("BUILD_MONTH_IS_MAR", "(__DATE__[0] == 'M' && __DATE__[1] == 'a' && __DATE__[2] == 'r')")
	// m.Set("BUILD_MONTH_IS_APR", "(__DATE__[0] == 'A' && __DATE__[1] == 'p')")
	// m.Set("BUILD_MONTH_IS_MAY", "(__DATE__[0] == 'M' && __DATE__[1] == 'a' && __DATE__[2] == 'y')")
	// m.Set("BUILD_MONTH_IS_JUN", "(__DATE__[0] == 'J' && __DATE__[1] == 'u' && __DATE__[2] == 'n')")
	// m.Set("BUILD_MONTH_IS_JUL", "(__DATE__[0] == 'J' && __DATE__[1] == 'u' && __DATE__[2] == 'l')")
	// m.Set("BUILD_MONTH_IS_AUG", "(__DATE__[0] == 'A' && __DATE__[1] == 'u')")
	// m.Set("BUILD_MONTH_IS_SEP", "(__DATE__[0] == 'S')")
	// m.Set("BUILD_MONTH_IS_OCT", "(__DATE__[0] == 'O')")
	// m.Set("BUILD_MONTH_IS_NOV", "(__DATE__[0] == 'N')")
	// m.Set("BUILD_MONTH_IS_DEC", "(__DATE__[0] == 'D')")
	// m.Set("BUILD_MONTH_CH0", "\\ ((BUILD_MONTH_IS_OCT || BUILD_MONTH_IS_NOV || BUILD_MONTH_IS_DEC) ? '1' : '0')")
	// m.Set("BUILD_MONTH_CH1", "\\ (                                                           \\ (BUILD_MONTH_IS_JAN) ? '1' : (BUILD_MONTH_IS_FEB) ? '2' \\ : (BUILD_MONTH_IS_MAR)   ? '3' \\ : (BUILD_MONTH_IS_APR)   ? '4' \\ : (BUILD_MONTH_IS_MAY)   ? '5' \\ : (BUILD_MONTH_IS_JUN)   ? '6' \\ : (BUILD_MONTH_IS_JUL)   ? '7' \\ : (BUILD_MONTH_IS_AUG)   ? '8' \\ : (BUILD_MONTH_IS_SEP)   ? '9' \\ : (BUILD_MONTH_IS_OCT)   ? '0' \\ : (BUILD_MONTH_IS_NOV)   ? '1' \\ : (BUILD_MONTH_IS_DEC)   ? '2' \\ : /* error default */ '?')")
	// m.Set("BUILD_DAY_CH0", "((__DATE__[4] >= '0') ? (__DATE__[4]) : '0')")
	// m.Set("BUILD_DAY_CH1", "(__DATE__[5])")
	// m.Set("BUILD_HOUR_CH0", "(__TIME__[0])")
	// m.Set("BUILD_HOUR_CH1", "(__TIME__[1])")
	// m.Set("BUILD_MIN_CH0", "(__TIME__[3])")
	// m.Set("BUILD_MIN_CH1", "(__TIME__[4])")
	// m.Set("BUILD_SEC_CH0", "(__TIME__[6])")
	// m.Set("BUILD_SEC_CH1", "(__TIME__[7])")
	m.Set("MaximumPacketsCapacity", "1000")
	m.Set("MaximumPacketsCapacityPriority", "50")
	// m.Set("NORMAL_PAGE_SIZE", "4096 // PAGE_SIZE")
	m.Set("PacketChunkSize", "NORMAL_PAGE_SIZE")
	m.Set("UsermodeBufferSize", "sizeof(UINT32) + PacketChunkSize + 1")
	m.Set("MaxSerialPacketSize", "10 * NORMAL_PAGE_SIZE")
	m.Set("LogBufferSize", "\\ MaximumPacketsCapacity *(PacketChunkSize + sizeof(BUFFER_HEADER))")
	m.Set("LogBufferSizePriority", "\\ MaximumPacketsCapacityPriority *(PacketChunkSize + sizeof(BUFFER_HEADER))")
	m.Set("DbgPrintLimitation", "512")
	m.Set("DebuggerEventTagStartSeed", "0x1000000")
	m.Set("DebuggerThreadDebuggingTagStartSeed", "0x1000000")
	m.Set("DebuggerOutputSourceTagStartSeed", "0x1")
	m.Set("DebuggerOutputSourceMaximumRemoteSourceForSingleEvent", "0x5")
	m.Set("DebuggerScriptEngineMemcpyMovingBufferSize", "64")
	m.Set("MAXIMUM_NUMBER_OF_INITIAL_PREALLOCATED_EPT_HOOKS", "5")
	m.Set("MAXIMUM_REGULAR_INSTANT_EVENTS", "20")
	m.Set("MAXIMUM_BIG_INSTANT_EVENTS", "0")
	m.Set("REGULAR_INSTANT_EVENT_CONDITIONAL_BUFFER", "sizeof(DEBUGGER_EVENT) + 100")
	m.Set("BIG_INSTANT_EVENT_CONDITIONAL_BUFFER", "sizeof(DEBUGGER_EVENT) + PAGE_SIZE")
	m.Set("REGULAR_INSTANT_EVENT_ACTION_BUFFER", "sizeof(DEBUGGER_EVENT_ACTION) + (PAGE_SIZE * 2)")
	m.Set("BIG_INSTANT_EVENT_ACTION_BUFFER", "sizeof(DEBUGGER_EVENT_ACTION) + MaxSerialPacketSize")
	m.Set("REGULAR_INSTANT_EVENT_REQUESTED_SAFE_BUFFER", "PAGE_SIZE")
	m.Set("BIG_INSTANT_EVENT_REQUESTED_SAFE_BUFFER", "MaxSerialPacketSize")
	m.Set("DEFAULT_PORT", "\"50000\"")
	m.Set("COMMUNICATION_BUFFER_SIZE", "PacketChunkSize + 0x100")
	m.Set("TOP_LEVEL_DRIVERS_VMCALL_STARTING_NUMBER", "0x00000200")
	m.Set("TOP_LEVEL_DRIVERS_VMCALL_ENDING_NUMBER", "TOP_LEVEL_DRIVERS_VMCALL_STARTING_NUMBER + 0x100")
	m.Set("OPERATION_MANDATORY_DEBUGGEE_BIT", "(1 << 31)")
	m.Set("OPERATION_LOG_INFO_MESSAGE", "1U")
	m.Set("OPERATION_LOG_WARNING_MESSAGE", "2U")
	m.Set("OPERATION_LOG_ERROR_MESSAGE", "3U")
	m.Set("OPERATION_LOG_NON_IMMEDIATE_MESSAGE", "4U")
	m.Set("OPERATION_LOG_WITH_TAG", "5U")
	m.Set("OPERATION_COMMAND_FROM_DEBUGGER_CLOSE_AND_UNLOAD_VMM", "\\ 6U | OPERATION_MANDATORY_DEBUGGEE_BIT")
	m.Set("OPERATION_DEBUGGEE_USER_INPUT", "7U | OPERATION_MANDATORY_DEBUGGEE_BIT")
	m.Set("OPERATION_DEBUGGEE_REGISTER_EVENT", "8U | OPERATION_MANDATORY_DEBUGGEE_BIT")
	m.Set("OPERATION_DEBUGGEE_ADD_ACTION_TO_EVENT", "\\ 9 | OPERATION_MANDATORY_DEBUGGEE_BIT")
	m.Set("OPERATION_DEBUGGEE_CLEAR_EVENTS", "10U | OPERATION_MANDATORY_DEBUGGEE_BIT")
	m.Set("OPERATION_DEBUGGEE_CLEAR_EVENTS_WITHOUT_NOTIFYING_DEBUGGER", "11U | OPERATION_MANDATORY_DEBUGGEE_BIT")
	m.Set("OPERATION_HYPERVISOR_DRIVER_IS_SUCCESSFULLY_LOADED", "\\ 12U | OPERATION_MANDATORY_DEBUGGEE_BIT")
	m.Set("OPERATION_HYPERVISOR_DRIVER_END_OF_IRPS", "\\ 13U | OPERATION_MANDATORY_DEBUGGEE_BIT")
	m.Set("OPERATION_COMMAND_FROM_DEBUGGER_RELOAD_SYMBOL", "\\ 14U | OPERATION_MANDATORY_DEBUGGEE_BIT")
	m.Set("OPERATION_NOTIFICATION_FROM_USER_DEBUGGER_PAUSE", "\\ 15U | OPERATION_MANDATORY_DEBUGGEE_BIT")
	m.Set("MAXIMUM_BREAKPOINTS_WITHOUT_CONTINUE", "100")
	m.Set("MAXIMUM_NUMBER_OF_THREAD_INFORMATION_FOR_TRAPS", "200")
	m.Set("POOLTAG", "0x48444247 // [H]yper[DBG] (HDBG)")
	m.Set("SERIAL_END_OF_BUFFER_CHARS_COUNT", "0x4")
	m.Set("SERIAL_END_OF_BUFFER_CHAR_1", "0x00")
	m.Set("SERIAL_END_OF_BUFFER_CHAR_2", "0x80")
	m.Set("SERIAL_END_OF_BUFFER_CHAR_3", "0xEE")
	m.Set("SERIAL_END_OF_BUFFER_CHAR_4", "0xFF")
	m.Set("TCP_END_OF_BUFFER_CHARS_COUNT", "0x4")
	m.Set("TCP_END_OF_BUFFER_CHAR_1", "0x10")
	m.Set("TCP_END_OF_BUFFER_CHAR_2", "0x20")
	m.Set("TCP_END_OF_BUFFER_CHAR_3", "0x33")
	m.Set("TCP_END_OF_BUFFER_CHAR_4", "0x44")
	m.Set("MAXIMUM_CHARACTER_FOR_OS_NAME", "256")
	m.Set("MAXIMUM_INSTR_SIZE", "16")
	m.Set("MAXIMUM_CALL_INSTR_SIZE", "7")
	m.Set("MAXIMUM_SUPPORTED_SYMBOLS", "1000")
	m.Set("MAXIMUM_GUID_AND_AGE_SIZE", "60")
	m.Set("INDICATOR_OF_HYPERDBG_PACKET", "\\ 0x4859504552444247 // HYPERDBG = 0x4859504552444247")
	m.Set("MaximumSearchResults", "0x1000")
	m.Set("X86_FLAGS_CF", "(1 << 0)")
	m.Set("X86_FLAGS_PF", "(1 << 2)")
	m.Set("X86_FLAGS_AF", "(1 << 4)")
	m.Set("X86_FLAGS_ZF", "(1 << 6)")
	m.Set("X86_FLAGS_SF", "(1 << 7)")
	m.Set("X86_FLAGS_TF", "(1 << 8)")
	m.Set("X86_FLAGS_IF", "(1 << 9)")
	m.Set("X86_FLAGS_DF", "(1 << 10)")
	m.Set("X86_FLAGS_OF", "(1 << 11)")
	m.Set("X86_FLAGS_STATUS_MASK", "(0xfff)")
	m.Set("X86_FLAGS_IOPL_MASK", "(3 << 12)")
	m.Set("X86_FLAGS_IOPL_SHIFT", "(12)")
	m.Set("X86_FLAGS_IOPL_SHIFT_2ND_BIT", "(13)")
	m.Set("X86_FLAGS_NT", "(1 << 14)")
	m.Set("X86_FLAGS_RF", "(1 << 16)")
	m.Set("X86_FLAGS_VM", "(1 << 17)")
	m.Set("X86_FLAGS_AC", "(1 << 18)")
	m.Set("X86_FLAGS_VIF", "(1 << 19)")
	m.Set("X86_FLAGS_VIP", "(1 << 20)")
	m.Set("X86_FLAGS_ID", "(1 << 21)")
	m.Set("X86_FLAGS_RESERVED_ONES", "0x2")
	m.Set("X86_FLAGS_RESERVED", "0xffc0802a")
	m.Set("X86_FLAGS_RESERVED_BITS", "0xffc38028")
	m.Set("X86_FLAGS_FIXED", "0x00000002")
	m.Set("MAX_TEMP_COUNT", "128")
	m.Set("MAX_STACK_BUFFER_COUNT", "128")
	m.Set("MAX_VAR_COUNT", "512")
	m.Set("MAX_FUNCTION_NAME_LENGTH", "32")
	m.Set("DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG", "0xffffffffffffffff")
	m.Set("DISASSEMBLY_MAXIMUM_DISTANCE_FROM_OBJECT_NAME", "0xffff")
	m.Set("DEBUGGER_READ_AND_WRITE_ON_MSR_APPLY_ALL_CORES", "0xffffffff")
	m.Set("DEBUGGER_DEBUGGEE_IS_RUNNING_NO_CORE", "0xffffffff")
	m.Set("DEBUGGER_EVENT_APPLY_TO_ALL_CORES", "0xffffffff")
	m.Set("DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES", "0xffffffff")
	m.Set("DEBUGGER_EVENT_MSR_READ_OR_WRITE_ALL_MSRS", "0xffffffff")
	m.Set("DEBUGGER_EVENT_EXCEPTIONS_ALL_FIRST_32_ENTRIES", "0xffffffff")
	m.Set("DEBUGGER_EVENT_SYSCALL_ALL_SYSRET_OR_SYSCALLS", "0xffffffff")
	m.Set("DEBUGGER_EVENT_ALL_IO_PORTS", "0xffffffff")
	m.Set("DEBUGGEE_BP_APPLY_TO_ALL_CORES", "0xffffffff")
	m.Set("DEBUGGEE_BP_APPLY_TO_ALL_PROCESSES", "0xffffffff")
	m.Set("DEBUGGEE_BP_APPLY_TO_ALL_THREADS", "0xffffffff")
	m.Set("DEBUGGEE_SHOW_ALL_REGISTERS", "0xffffffff")
	m.Set("SIZEOF_DEBUGGER_PAUSE_PACKET_RECEIVED", "\\ sizeof(DEBUGGER_PAUSE_PACKET_RECEIVED)")
	m.Set("DEBUGGER_OPERATION_WAS_SUCCESSFUL", "0xFFFFFFFF")
	m.Set("DEBUGGER_ERROR_TAG_NOT_EXISTS", "0xc0000000")
	m.Set("DEBUGGER_ERROR_INVALID_ACTION_TYPE", "0xc0000001")
	m.Set("DEBUGGER_ERROR_ACTION_BUFFER_SIZE_IS_ZERO", "0xc0000002")
	m.Set("DEBUGGER_ERROR_EVENT_TYPE_IS_INVALID", "0xc0000003")
	m.Set("DEBUGGER_ERROR_UNABLE_TO_CREATE_EVENT", "0xc0000004")
	m.Set("DEBUGGER_ERROR_INVALID_ADDRESS", "0xc0000005")
	m.Set("DEBUGGER_ERROR_INVALID_CORE_ID", "0xc0000006")
	m.Set("DEBUGGER_ERROR_EXCEPTION_INDEX_EXCEED_FIRST_32_ENTRIES", "0xc0000007")
	m.Set("DEBUGGER_ERROR_INTERRUPT_INDEX_IS_NOT_VALID", "0xc0000008")
	m.Set("DEBUGGER_ERROR_UNABLE_TO_HIDE_OR_UNHIDE_DEBUGGER", "0xc0000009")
	m.Set("DEBUGGER_ERROR_DEBUGGER_ALREADY_UHIDE", "0xc000000a")
	m.Set("DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER", "0xc000000b")
	m.Set("DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_CURRENT_PROCESS", "\\ 0xc000000c")
	m.Set("DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_OTHER_PROCESS", "\\ 0xc000000d")
	m.Set("DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TAG", "0xc000000e")
	m.Set("DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TYPE_OF_ACTION", "0xc000000f")
	m.Set("DEBUGGER_ERROR_STEPPING_INVALID_PARAMETER", "0xc0000010")
	m.Set("DEBUGGER_ERROR_STEPPINGS_EITHER_THREAD_NOT_FOUND_OR_DISABLED", "0xc0000011")
	m.Set("DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_BAUDRATE", "0xc0000012")
	m.Set("DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_SERIAL_PORT", "0xc0000013")
	m.Set("DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_CORE_IN_REMOTE_DEBUGGE", "\\ 0xc0000014")
	m.Set("DEBUGGER_ERROR_PREPARING_DEBUGGEE_UNABLE_TO_SWITCH_TO_NEW_PROCESS", "\\ 0xc0000015")
	m.Set("DEBUGGER_ERROR_PREPARING_DEBUGGEE_TO_RUN_SCRIPT", "0xc0000016")
	m.Set("DEBUGGER_ERROR_INVALID_REGISTER_NUMBER", "0xc0000017")
	m.Set("DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_WITHOUT_CONTINUE", "0xc0000018")
	m.Set("DEBUGGER_ERROR_BREAKPOINT_ALREADY_EXISTS_ON_THE_ADDRESS", "0xc0000019")
	m.Set("DEBUGGER_ERROR_BREAKPOINT_ID_NOT_FOUND", "0xc000001a")
	m.Set("DEBUGGER_ERROR_BREAKPOINT_ALREADY_DISABLED", "0xc000001b")
	m.Set("DEBUGGER_ERROR_BREAKPOINT_ALREADY_ENABLED", "0xc000001c")
	m.Set("DEBUGGER_ERROR_MEMORY_TYPE_INVALID", "0xc000001d")
	m.Set("DEBUGGER_ERROR_INVALID_PROCESS_ID", "0xc000001e")
	m.Set("DEBUGGER_ERROR_EVENT_IS_NOT_APPLIED", "0xc000001f")
	m.Set("DEBUGGER_ERROR_DETAILS_OR_SWITCH_PROCESS_INVALID_PARAMETER", "0xc0000020")
	m.Set("DEBUGGER_ERROR_DETAILS_OR_SWITCH_THREAD_INVALID_PARAMETER", "0xc0000021")
	m.Set("DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_FOR_A_SINGLE_PAGE_IS_HIT", "0xc0000022")
	m.Set("DEBUGGER_ERROR_PRE_ALLOCATED_BUFFER_IS_EMPTY", "0xc0000023")
	m.Set("DEBUGGER_ERROR_EPT_COULD_NOT_SPLIT_THE_LARGE_PAGE_TO_4KB_PAGES", "0xc0000024")
	m.Set("DEBUGGER_ERROR_EPT_FAILED_TO_GET_PML1_ENTRY_OF_TARGET_ADDRESS", "0xc0000025")
	m.Set("DEBUGGER_ERROR_EPT_MULTIPLE_HOOKS_IN_A_SINGLE_PAGE", "0xc0000026")
	m.Set("DEBUGGER_ERROR_COULD_NOT_BUILD_THE_EPT_HOOK", "0xc0000027")
	m.Set("DEBUGGER_ERROR_COULD_NOT_FIND_ALLOCATION_TYPE", "0xc0000028")
	m.Set("DEBUGGER_ERROR_INVALID_TEST_QUERY_INDEX", "0xc0000029")
	m.Set("DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_TARGET_USER_MODE_PROCESS", "0xc000002a")
	m.Set("DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS_ENTRYPOINT_NOT_REACHED", "0xc000002b")
	m.Set("DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS", "0xc000002c")
	m.Set("DEBUGGER_ERROR_FUNCTIONS_FOR_INITIALIZING_PEB_ADDRESSES_ARE_NOT_INITIALIZED", "0xc000002d")
	m.Set("DEBUGGER_ERROR_UNABLE_TO_DETECT_32_BIT_OR_64_BIT_PROCESS", "0xc000002e")
	m.Set("DEBUGGER_ERROR_UNABLE_TO_KILL_THE_PROCESS", "0xc000002f")
	m.Set("DEBUGGER_ERROR_INVALID_THREAD_DEBUGGING_TOKEN", "0xc0000030")
	m.Set("DEBUGGER_ERROR_UNABLE_TO_PAUSE_THE_PROCESS_THREADS", "0xc0000031")
	m.Set("DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_AN_ALREADY_ATTACHED_PROCESS", "0xc0000032")
	m.Set("DEBUGGER_ERROR_THE_USER_DEBUGGER_NOT_ATTACHED_TO_THE_PROCESS", "0xc0000033")
	m.Set("DEBUGGER_ERROR_UNABLE_TO_DETACH_AS_THERE_ARE_PAUSED_THREADS", "0xc0000034")
	m.Set("DEBUGGER_ERROR_UNABLE_TO_SWITCH_PROCESS_ID_OR_THREAD_ID_IS_INVALID", "0xc0000035")
	m.Set("DEBUGGER_ERROR_UNABLE_TO_SWITCH_THERE_IS_NO_THREAD_ON_THE_PROCESS", "0xc0000036")
	m.Set("DEBUGGER_ERROR_UNABLE_TO_GET_MODULES_OF_THE_PROCESS", "0xc0000037")
	m.Set("DEBUGGER_ERROR_UNABLE_TO_GET_CALLSTACK", "0xc0000038")
	m.Set("DEBUGGER_ERROR_UNABLE_TO_QUERY_COUNT_OF_PROCESSES_OR_THREADS", "0xc0000039")
	m.Set("DEBUGGER_ERROR_USING_SHORT_CIRCUITING_EVENT_WITH_POST_EVENT_MODE_IS_FORBIDDEDN", "0xc000003a")
	m.Set("DEBUGGER_ERROR_UNKNOWN_TEST_QUERY_RECEIVED", "0xc000003b")
	m.Set("DEBUGGER_ERROR_READING_MEMORY_INVALID_PARAMETER", "0xc000003c")
	m.Set("DEBUGGER_ERROR_THE_TRAP_FLAG_LIST_IS_FULL", "0xc000003d")
	m.Set("DEBUGGER_ERROR_UNABLE_TO_KILL_THE_PROCESS_DOES_NOT_EXISTS", "0xc000003e")
	m.Set("DEBUGGER_ERROR_MODE_EXECUTION_IS_INVALID", "0xc000003f")
	m.Set("DEBUGGER_ERROR_PROCESS_ID_CANNOT_BE_SPECIFIED_WHILE_APPLYING_EVENT_FROM_VMX_ROOT_MODE", "0xc0000040")
	m.Set("DEBUGGER_ERROR_INSTANT_EVENT_PREALLOCATED_BUFFER_IS_NOT_ENOUGH_FOR_EVENT_AND_CONDITIONALS", "0xc0000041")
	m.Set("DEBUGGER_ERROR_INSTANT_EVENT_REGULAR_PREALLOCATED_BUFFER_NOT_FOUND", "0xc0000042")
	m.Set("DEBUGGER_ERROR_INSTANT_EVENT_BIG_PREALLOCATED_BUFFER_NOT_FOUND", "0xc0000043")
	m.Set("DEBUGGER_ERROR_UNABLE_TO_CREATE_ACTION_CANNOT_ALLOCATE_BUFFER", "0xc0000044")
	m.Set("DEBUGGER_ERROR_INSTANT_EVENT_ACTION_REGULAR_PREALLOCATED_BUFFER_NOT_FOUND", "0xc0000045")
	m.Set("DEBUGGER_ERROR_INSTANT_EVENT_ACTION_BIG_PREALLOCATED_BUFFER_NOT_FOUND", "0xc0000046")
	m.Set("DEBUGGER_ERROR_INSTANT_EVENT_PREALLOCATED_BUFFER_IS_NOT_ENOUGH_FOR_ACTION_BUFFER", "0xc0000047")
	m.Set("DEBUGGER_ERROR_INSTANT_EVENT_REQUESTED_OPTIONAL_BUFFER_IS_BIGGER_THAN_DEBUGGERS_SEND_RECEIVE_STACK", "0xc0000048")
	m.Set("DEBUGGER_ERROR_INSTANT_EVENT_REGULAR_REQUESTED_SAFE_BUFFER_NOT_FOUND", "0xc0000049")
	m.Set("DEBUGGER_ERROR_INSTANT_EVENT_BIG_REQUESTED_SAFE_BUFFER_NOT_FOUND", "0xc000004a")
	m.Set("DEBUGGER_ERROR_INSTANT_EVENT_PREALLOCATED_BUFFER_IS_NOT_ENOUGH_FOR_REQUESTED_SAFE_BUFFER", "0xc000004b")
	m.Set("DEBUGGER_ERROR_UNABLE_TO_ALLOCATE_REQUESTED_SAFE_BUFFER", "0xc000004c")
	m.Set("DEBUGGER_ERROR_COULD_NOT_FIND_PREACTIVATION_TYPE", "0xc000004d")
	m.Set("DEBUGGER_ERROR_THE_MODE_EXEC_TRAP_IS_NOT_INITIALIZED", "0xc000004e")
	m.Set("DEBUGGER_ERROR_THE_TARGET_EVENT_IS_DISABLED_BUT_CANNOT_BE_CLEARED_PRIRITY_BUFFER_IS_FULL", "0xc000004f")
	m.Set("DEBUGGER_ERROR_NOT_ALL_CORES_ARE_LOCKED_FOR_APPLYING_INSTANT_EVENT", "0xc0000050")
	m.Set("DEBUGGER_ERROR_TARGET_SWITCHING_CORE_IS_NOT_LOCKED", "0xc0000051")
	m.Set("DEBUGGER_ERROR_INVALID_PHYSICAL_ADDRESS", "0xc0000052")
	m.Set("SIZEOF_DEBUGGER_MODIFY_EVENTS", "sizeof(DEBUGGER_MODIFY_EVENTS)")
	m.Set("SIZEOF_REGISTER_EVENT", "sizeof(REGISTER_NOTIFY_BUFFER)")
	m.Set("DEFAULT_INITIAL_DEBUGGEE_TO_DEBUGGER_OFFSET", "0x200")
	m.Set("DEFAULT_INITIAL_DEBUGGER_TO_DEBUGGEE_OFFSET", "0x0")
	m.Set("IOCTL_REGISTER_EVENT", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x800, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x801, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_TERMINATE_VMX", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x802, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_DEBUGGER_READ_MEMORY", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x803, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_DEBUGGER_READ_OR_WRITE_MSR", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x804, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x805, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_DEBUGGER_REGISTER_EVENT", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x806, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x807, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x808, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x809, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_DEBUGGER_EDIT_MEMORY", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80a, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_DEBUGGER_SEARCH_MEMORY", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80b, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_DEBUGGER_MODIFY_EVENTS", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80c, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_DEBUGGER_FLUSH_LOGGING_BUFFERS", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80d, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80e, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_DEBUGGER_PRINT", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80f, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_PREPARE_DEBUGGEE", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x810, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_PAUSE_PACKET_RECEIVED", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x811, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_SEND_SIGNAL_EXECUTION_IN_DEBUGGEE_FINISHED", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x812, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_SEND_USERMODE_MESSAGES_TO_DEBUGGER", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x813, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_SEND_GENERAL_BUFFER_FROM_DEBUGGEE_TO_DEBUGGER", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x814, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_PERFROM_KERNEL_SIDE_TESTS", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x815, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_RESERVE_PRE_ALLOCATED_POOLS", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x816, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_SEND_USER_DEBUGGER_COMMANDS", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x817, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_GET_DETAIL_OF_ACTIVE_THREADS_AND_PROCESSES", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x818, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_GET_USER_MODE_MODULE_DETAILS", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x819, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_QUERY_COUNT_OF_ACTIVE_PROCESSES_OR_THREADS", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81a, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81b, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_QUERY_CURRENT_PROCESS", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81c, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_QUERY_CURRENT_THREAD", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81d, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_REQUEST_REV_MACHINE_SERVICE", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81e, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_DEBUGGER_BRING_PAGES_IN", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81f, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("IOCTL_PREACTIVATE_FUNCTIONALITY", "\\ CTL_CODE(FILE_DEVICE_UNKNOWN, 0x820, METHOD_BUFFERED, FILE_ANY_ACCESS)")
	m.Set("SIZEOF_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS", "\\ sizeof(DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS)")
	m.Set("SIZEOF_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS", "\\ sizeof(DEBUGGER_VA2PA_AND_PA2VA_COMMANDS)")
	m.Set("SIZEOF_DEBUGGER_PAGE_IN_REQUEST", "\\ sizeof(DEBUGGER_PAGE_IN_REQUEST)")
	m.Set("SIZEOF_REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST", "\\ sizeof(REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST)")
	m.Set("SIZEOF_DEBUGGER_DT_COMMAND_OPTIONS", "\\ sizeof(DEBUGGER_DT_COMMAND_OPTIONS)")
	m.Set("SIZEOF_DEBUGGER_PREALLOC_COMMAND", "\\ sizeof(DEBUGGER_PREALLOC_COMMAND)")
	m.Set("SIZEOF_DEBUGGER_PREACTIVATE_COMMAND", "\\ sizeof(DEBUGGER_PREACTIVATE_COMMAND)")
	m.Set("SIZEOF_DEBUGGER_READ_MEMORY", "sizeof(DEBUGGER_READ_MEMORY)")
	m.Set("SIZEOF_DEBUGGER_FLUSH_LOGGING_BUFFERS", "\\ sizeof(DEBUGGER_FLUSH_LOGGING_BUFFERS)")
	m.Set("SIZEOF_DEBUGGER_TEST_QUERY_BUFFER", "\\ sizeof(DEBUGGER_TEST_QUERY_BUFFER)")
	m.Set("SIZEOF_DEBUGGER_PERFORM_KERNEL_TESTS", "\\ sizeof(DEBUGGER_PERFORM_KERNEL_TESTS)")
	m.Set("SIZEOF_DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL", "\\ sizeof(DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL)")
	m.Set("SIZEOF_DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER", "\\ sizeof(DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER)")
	m.Set("SIZEOF_DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER", "\\ sizeof(DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER)")
	m.Set("SIZEOF_DEBUGGER_READ_AND_WRITE_ON_MSR", "\\ sizeof(DEBUGGER_READ_AND_WRITE_ON_MSR)")
	m.Set("SIZEOF_DEBUGGER_EDIT_MEMORY", "sizeof(DEBUGGER_EDIT_MEMORY)")
	m.Set("SIZEOF_DEBUGGER_SEARCH_MEMORY", "sizeof(DEBUGGER_SEARCH_MEMORY)")
	m.Set("SIZEOF_DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE", "\\ sizeof(DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE)")
	m.Set("SIZEOF_DEBUGGER_PREPARE_DEBUGGEE", "sizeof(DEBUGGER_PREPARE_DEBUGGEE)")
	m.Set("SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS", "\\ sizeof(DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS)")
	m.Set("SIZEOF_DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS", "\\ sizeof(DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS)")
	m.Set("SIZEOF_DEBUGGER_CALLSTACK_REQUEST", "\\ sizeof(DEBUGGER_CALLSTACK_REQUEST)")
	m.Set("SIZEOF_USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS", "\\ sizeof(USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS)")
	m.Set("SIZEOF_DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET", "\\ sizeof(DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET)")
	m.Set("SIZEOF_DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET", "\\ sizeof(DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET)")
	m.Set("DEBUGGER_REMOTE_TRACKING_DEFAULT_COUNT_OF_STEPPING", "0xffffffff")
	// m.Set("LogDebugInfo(format,", "...) \\ if (DebugMode)                                                                    \\ LogCallbackPrepareAndSendMessageToQueue(OPERATION_LOG_INFO_MESSAGE,               \\ UseImmediateMessaging,                    \\ ShowSystemTimeOnDebugMessages,            \\ FALSE,                                    \\ \"[+] Information (%s:%d) | \" format \"\\n\", \\ __func__,                                 \\ __LINE__,                                 \\ __VA_ARGS__)")
}
