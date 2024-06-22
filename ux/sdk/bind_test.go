package sdk

import (
	"github.com/ddkwork/app/bindgen/clang"
	"github.com/ddkwork/app/bindgen/gengo"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"github.com/ddkwork/golibrary/stream/orderedmap"
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestMergeHeader(t *testing.T) {
	Headers := orderedmap.New[string, bool]()
	Modules := orderedmap.New[string, bool]()
	Imports := orderedmap.New[string, bool]()

	g := stream.NewGeneratedFile()
	filepath.Walk("../../../bin", func(path string, info fs.FileInfo, err error) error {
		if strings.Contains(path, "Examples") {
			return err
		}
		switch stream.BaseName(path) {
		case "Assertions", "HyperDbgSdk":
			return err
		}
		if filepath.Ext(path) == ".h" {
			// println(path)
			switch {
			case strings.Contains(path, "Headers"):
				Headers.Set(path, true)
			case strings.Contains(path, "Imports"):
				Imports.Set(path, true)
			case strings.Contains(path, "Modules"):
				Modules.Set(path, true)
			}
		}
		return err
	})

	// todo merge step
	// bugfix
	// Headers
	// Modules
	// Imports

	bugfix = strings.TrimPrefix(bugfix, "\n")
	g.P("//bugfix.h")
	g.P(bugfix)
	g.P()
	mylog.Trace("merge", "bugfix.h")

	fnDo := func(path string) {
		g.P("//" + path)
		g.P(stream.NewBuffer(path)) // todo remove #pragma once ?
		g.P()
		mylog.Trace("merge", path)
	}

	for _, s := range Headers.Keys() {
		fnDo(s)
	}
	for _, s := range Modules.Keys() {
		fnDo(s)
	}
	for _, s := range Imports.Keys() {
		fnDo(s)
	}
	stream.WriteBinaryFile("merged_headers.h", g.Buffer)
}

func TestBindMacros(t *testing.T) {
	headerFile := "merged_headers.h"
	macros := extractMacros(stream.NewBuffer(headerFile).ToLines())
	mylog.Trace("number of macros", macros.Len())

	j := stream.NewGeneratedFile()
	j.P("var skips = []string{")

	for _, p := range macros.List() {
		j.P(strconv.Quote(p.Key), ",")
		ok := true
		for _, skip := range skips {
			if p.Key == skip {
				ok = false
				macros.Delete(p.Key)
				continue
			}
		}
		if ok {
			// mylog.Warning(p.Key, p.Value)
		}
	}
	j.P("}")
	// println(j.String())
	// return

	g := stream.NewGeneratedFile()
	g.P("package HPRDBGCTRL")
	g.P()
	g.P("var (")
	g.P("NORMAL_PAGE_SIZE=4096")
	// g.P("MaxSerialPacketSize =10 * NORMAL_PAGE_SIZE") // todo need first define NORMAL_PAGE_SIZE
	g.P("PAGE_SIZE = 4096")

	for _, p := range macros.List() {
		p.Value = strings.TrimSpace(p.Value)
		if strings.HasPrefix(p.Value, "sizeof") {
			continue
		}
		if strings.Contains(p.Value, "sizeof") {
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
		p.Value = strings.TrimSuffix(p.Value, "U")
		p.Value = strings.TrimSuffix(p.Value, "ull")

		if strings.HasPrefix(p.Value, "CTL_CODE") {
			p.Value = "IoctlsKind(" + p.Value + ")"
		}
		g.P(p.Key + "=" + p.Value)
		macros.Delete(p.Key)
	}
	g.P(")")
	stream.WriteGoFile("tmp/vars.go", g.Buffer)

	for _, p := range macros.List() {
		mylog.Todo(p.Key + " = " + p.Value)
	}
}

func TestBindSdk(t *testing.T) {
	// mylog.SetDebug(false)
	mylog.Call(func() {
		pkg := gengo.NewPackage("HPRDBGCTRL",
			gengo.WithRemovePrefix(
			//"Zydis_", "Zyan_", "Zycore_",
			//"Zydis", "Zyan", "Zycore",
			),
			gengo.WithInferredMethods([]gengo.MethodInferenceRule{
				//{Name: "ZydisDecoder", Receiver: "Decoder"},
			}),
			gengo.WithForcedSynthetic(
			//"ZydisShortString_",
			//"struct ZydisShortString_",
			),
		)
		mylog.Check(pkg.Transform("HPRDBGCTRL", &clang.Options{
			// Sources:          []string{"combined_headers.h"},
			Sources:          []string{"merged_headers.h"},
			AdditionalParams: []string{
				//"-DZYAN_NO_LIBC",
				//"-DZYAN_STATIC_ASSERT",
				//"-DZYDIS_STATIC_BUILD",
				//"-DHYPERDBG_HPRDBGCTRL",

				//"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\shared",
				//"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\ucrt",
				//"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\um",
				//"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\km",
				//"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\km\\crt",

				//"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\winrt",
				//"-IC:\\Program Files\\Microsoft Visual Studio\\2022\\Enterprise\\VC\\Tools\\MSVC\\14.40.33807\\include",

				//"-ID:\\fork\\HyperDbg\\hyperdbg\\hprdbgctrl",
				//"-ID:\\fork\\HyperDbg\\hyperdbg\\hprdbghv",
				//"-ID:\\fork\\HyperDbg\\hyperdbg\\hprdbgctrl\\header",
				//"-ID:\\fork\\HyperDbg\\hyperdbg\\include",
				//"-ID:\\fork\\HyperDbg\\hyperdbg\\dependencies",
				//"-ID:\\fork\\HyperDbg\\hyperdbg\\dependencies\\phnt",
			},
		}))
		mylog.Check(pkg.WriteToDir("./tmp"))
	})
}

var bugfix = `
typedef unsigned short wchar_t;
typedef int bool ;
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

/*
typedef struct _IRP {
  CSHORT                    Type;
  USHORT                    Size;
  PMDL                      MdlAddress;
  ULONG                     Flags;
  union {
    struct _IRP     *MasterIrp;
    __volatile LONG IrpCount;
    PVOID           SystemBuffer;
  } AssociatedIrp;
  LIST_ENTRY                ThreadListEntry;
  IO_STATUS_BLOCK           IoStatus;
  KPROCESSOR_MODE           RequestorMode;
  BOOLEAN                   PendingReturned;
  CHAR                      StackCount;
  CHAR                      CurrentLocation;
  BOOLEAN                   Cancel;
  KIRQL                     CancelIrql;
  CCHAR                     ApcEnvironment;
  UCHAR                     AllocationFlags;
  union {
    PIO_STATUS_BLOCK UserIosb;
    PVOID            IoRingContext;
  };
  PKEVENT                   UserEvent;
  union {
    struct {
      union {
        PIO_APC_ROUTINE UserApcRoutine;
        PVOID           IssuingProcess;
      };
      union {
        PVOID                 UserApcContext;
#if ...
        _IORING_OBJECT        *IoRing;
#else
        struct _IORING_OBJECT *IoRing;
#endif
      };
    } AsynchronousParameters;
    LARGE_INTEGER AllocationSize;
  } Overlay;
  __volatile PDRIVER_CANCEL CancelRoutine;
  PVOID                     UserBuffer;
  union {
    struct {
      union {
        KDEVICE_QUEUE_ENTRY DeviceQueueEntry;
        struct {
          PVOID DriverContext[4];
        };
      };
      PETHREAD     Thread;
      PCHAR        AuxiliaryBuffer;
      struct {
        LIST_ENTRY ListEntry;
        union {
          struct _IO_STACK_LOCATION *CurrentStackLocation;
          ULONG                     PacketType;
        };
      };
      PFILE_OBJECT OriginalFileObject;
    } Overlay;
    KAPC  Apc;
    PVOID CompletionKey;
  } Tail;
} IRP;
*/
`

var skips = []string{
	"PVOID",
	"HANDLE",
	"PIRP",
	"PDEVICE_OBJECT",
	"PSYMBOL_BUFFER",
	"PSYMBOL",
	//"MAX_PATH",
	"_In_",
	"_Out_",
	"_Inout_",
	"_Out_writes_bytes_(x)",
	"_In_reads_(x)",
	"_In_reads_bytes_(x)",
	"VOID",
	"NULL_ZERO",
	"NULL64_ZERO",
	"FALSE",
	"TRUE",
	//"UPPER_56_BITS",
	//"UPPER_48_BITS",
	//"UPPER_32_BITS",
	//"LOWER_32_BITS",
	//"LOWER_16_BITS",
	//"LOWER_8_BITS",
	//"SECOND_LOWER_8_BITS",
	//"UPPER_48_BITS_AND_LOWER_8_BITS",
	"VERSION_MAJOR",
	"VERSION_MINOR",
	"VERSION_PATCH",
	"BUILD_YEAR_CH0",
	"BUILD_YEAR_CH1",
	"BUILD_YEAR_CH2",
	"BUILD_YEAR_CH3",
	"BUILD_MONTH_IS_JAN",
	"BUILD_MONTH_IS_FEB",
	"BUILD_MONTH_IS_MAR",
	"BUILD_MONTH_IS_APR",
	"BUILD_MONTH_IS_MAY",
	"BUILD_MONTH_IS_JUN",
	"BUILD_MONTH_IS_JUL",
	"BUILD_MONTH_IS_AUG",
	"BUILD_MONTH_IS_SEP",
	"BUILD_MONTH_IS_OCT",
	"BUILD_MONTH_IS_NOV",
	"BUILD_MONTH_IS_DEC",
	"BUILD_MONTH_CH0",
	"BUILD_MONTH_CH1",
	"BUILD_DAY_CH0",
	"BUILD_DAY_CH1",
	"BUILD_HOUR_CH0",
	"BUILD_HOUR_CH1",
	"BUILD_MIN_CH0",
	"BUILD_MIN_CH1",
	"BUILD_SEC_CH0",
	"BUILD_SEC_CH1",
	//"MaximumPacketsCapacity",
	//"MaximumPacketsCapacityPriority",
	"NORMAL_PAGE_SIZE",
	//"PacketChunkSize",
	//"UsermodeBufferSize",
	//"MaxSerialPacketSize",
	//"LogBufferSize",
	//"LogBufferSizePriority",
	//"DbgPrintLimitation",
	//"DebuggerEventTagStartSeed",
	//"DebuggerThreadDebuggingTagStartSeed",
	//"DebuggerOutputSourceTagStartSeed",
	//"DebuggerOutputSourceMaximumRemoteSourceForSingleEvent",
	//"DebuggerScriptEngineMemcpyMovingBufferSize",
	//"MAXIMUM_NUMBER_OF_INITIAL_PREALLOCATED_EPT_HOOKS",
	//"MAXIMUM_REGULAR_INSTANT_EVENTS",
	//"MAXIMUM_BIG_INSTANT_EVENTS",
	//"REGULAR_INSTANT_EVENT_CONDITIONAL_BUFFER",
	//"BIG_INSTANT_EVENT_CONDITIONAL_BUFFER",
	//"REGULAR_INSTANT_EVENT_ACTION_BUFFER",
	//"BIG_INSTANT_EVENT_ACTION_BUFFER",
	//"REGULAR_INSTANT_EVENT_REQUESTED_SAFE_BUFFER",
	//"BIG_INSTANT_EVENT_REQUESTED_SAFE_BUFFER",
	//"DEFAULT_PORT",
	//"COMMUNICATION_BUFFER_SIZE",
	//"TOP_LEVEL_DRIVERS_VMCALL_STARTING_NUMBER",
	//"TOP_LEVEL_DRIVERS_VMCALL_ENDING_NUMBER",
	//"OPERATION_MANDATORY_DEBUGGEE_BIT",
	//"OPERATION_LOG_INFO_MESSAGE",
	//"OPERATION_LOG_WARNING_MESSAGE",
	//"OPERATION_LOG_ERROR_MESSAGE",
	//"OPERATION_LOG_NON_IMMEDIATE_MESSAGE",
	//"OPERATION_LOG_WITH_TAG",
	//"OPERATION_COMMAND_FROM_DEBUGGER_CLOSE_AND_UNLOAD_VMM",
	//"OPERATION_DEBUGGEE_USER_INPUT",
	//"OPERATION_DEBUGGEE_REGISTER_EVENT",
	//"OPERATION_DEBUGGEE_ADD_ACTION_TO_EVENT",
	//"OPERATION_DEBUGGEE_CLEAR_EVENTS",
	//"OPERATION_DEBUGGEE_CLEAR_EVENTS_WITHOUT_NOTIFYING_DEBUGGER",
	//"OPERATION_HYPERVISOR_DRIVER_IS_SUCCESSFULLY_LOADED",
	//"OPERATION_HYPERVISOR_DRIVER_END_OF_IRPS",
	//"OPERATION_COMMAND_FROM_DEBUGGER_RELOAD_SYMBOL",
	//"OPERATION_NOTIFICATION_FROM_USER_DEBUGGER_PAUSE",
	//"MAXIMUM_BREAKPOINTS_WITHOUT_CONTINUE",
	//"MAXIMUM_NUMBER_OF_THREAD_INFORMATION_FOR_TRAPS",
	//"POOLTAG",
	//"SERIAL_END_OF_BUFFER_CHARS_COUNT",
	//"SERIAL_END_OF_BUFFER_CHAR_1",
	//"SERIAL_END_OF_BUFFER_CHAR_2",
	//"SERIAL_END_OF_BUFFER_CHAR_3",
	//"SERIAL_END_OF_BUFFER_CHAR_4",
	//"TCP_END_OF_BUFFER_CHARS_COUNT",
	//"TCP_END_OF_BUFFER_CHAR_1",
	//"TCP_END_OF_BUFFER_CHAR_2",
	//"TCP_END_OF_BUFFER_CHAR_3",
	//"TCP_END_OF_BUFFER_CHAR_4",
	//"MAXIMUM_CHARACTER_FOR_OS_NAME",
	//"MAXIMUM_INSTR_SIZE",
	//"MAXIMUM_CALL_INSTR_SIZE",
	//"MAXIMUM_SUPPORTED_SYMBOLS",
	//"MAXIMUM_GUID_AND_AGE_SIZE",
	//"INDICATOR_OF_HYPERDBG_PACKET",
	//"MaximumSearchResults",
	//"X86_FLAGS_CF",
	//"X86_FLAGS_PF",
	//"X86_FLAGS_AF",
	//"X86_FLAGS_ZF",
	//"X86_FLAGS_SF",
	//"X86_FLAGS_TF",
	//"X86_FLAGS_IF",
	//"X86_FLAGS_DF",
	//"X86_FLAGS_OF",
	//"X86_FLAGS_STATUS_MASK",
	//"X86_FLAGS_IOPL_MASK",
	//"X86_FLAGS_IOPL_SHIFT",
	//"X86_FLAGS_IOPL_SHIFT_2ND_BIT",
	//"X86_FLAGS_NT",
	//"X86_FLAGS_RF",
	//"X86_FLAGS_VM",
	//"X86_FLAGS_AC",
	//"X86_FLAGS_VIF",
	//"X86_FLAGS_VIP",
	//"X86_FLAGS_ID",
	//"X86_FLAGS_RESERVED_ONES",
	//"X86_FLAGS_RESERVED",
	//"X86_FLAGS_RESERVED_BITS",
	//"X86_FLAGS_FIXED",
	//"MAX_TEMP_COUNT",
	//"MAX_STACK_BUFFER_COUNT",
	//"MAX_VAR_COUNT",
	//"MAX_FUNCTION_NAME_LENGTH",
	//"DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG",
	//"DISASSEMBLY_MAXIMUM_DISTANCE_FROM_OBJECT_NAME",
	//"DEBUGGER_READ_AND_WRITE_ON_MSR_APPLY_ALL_CORES",
	//"DEBUGGER_DEBUGGEE_IS_RUNNING_NO_CORE",
	//"DEBUGGER_EVENT_APPLY_TO_ALL_CORES",
	//"DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES",
	//"DEBUGGER_EVENT_MSR_READ_OR_WRITE_ALL_MSRS",
	//"DEBUGGER_EVENT_EXCEPTIONS_ALL_FIRST_32_ENTRIES",
	//"DEBUGGER_EVENT_SYSCALL_ALL_SYSRET_OR_SYSCALLS",
	//"DEBUGGER_EVENT_ALL_IO_PORTS",
	//"DEBUGGEE_BP_APPLY_TO_ALL_CORES",
	//"DEBUGGEE_BP_APPLY_TO_ALL_PROCESSES",
	//"DEBUGGEE_BP_APPLY_TO_ALL_THREADS",
	//"DEBUGGEE_SHOW_ALL_REGISTERS",
	//"SIZEOF_DEBUGGER_PAUSE_PACKET_RECEIVED",
	//"DEBUGGER_OPERATION_WAS_SUCCESSFUL",
	//"DEBUGGER_ERROR_TAG_NOT_EXISTS",
	//"DEBUGGER_ERROR_INVALID_ACTION_TYPE",
	//"DEBUGGER_ERROR_ACTION_BUFFER_SIZE_IS_ZERO",
	//"DEBUGGER_ERROR_EVENT_TYPE_IS_INVALID",
	//"DEBUGGER_ERROR_UNABLE_TO_CREATE_EVENT",
	//"DEBUGGER_ERROR_INVALID_ADDRESS",
	//"DEBUGGER_ERROR_INVALID_CORE_ID",
	//"DEBUGGER_ERROR_EXCEPTION_INDEX_EXCEED_FIRST_32_ENTRIES",
	//"DEBUGGER_ERROR_INTERRUPT_INDEX_IS_NOT_VALID",
	//"DEBUGGER_ERROR_UNABLE_TO_HIDE_OR_UNHIDE_DEBUGGER",
	//"DEBUGGER_ERROR_DEBUGGER_ALREADY_UHIDE",
	//"DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER",
	//"DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_CURRENT_PROCESS",
	//"DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_OTHER_PROCESS",
	//"DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TAG",
	//"DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TYPE_OF_ACTION",
	//"DEBUGGER_ERROR_STEPPING_INVALID_PARAMETER",
	//"DEBUGGER_ERROR_STEPPINGS_EITHER_THREAD_NOT_FOUND_OR_DISABLED",
	//"DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_BAUDRATE",
	//"DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_SERIAL_PORT",
	//"DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_CORE_IN_REMOTE_DEBUGGE",
	//"DEBUGGER_ERROR_PREPARING_DEBUGGEE_UNABLE_TO_SWITCH_TO_NEW_PROCESS",
	//"DEBUGGER_ERROR_PREPARING_DEBUGGEE_TO_RUN_SCRIPT",
	//"DEBUGGER_ERROR_INVALID_REGISTER_NUMBER",
	//"DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_WITHOUT_CONTINUE",
	//"DEBUGGER_ERROR_BREAKPOINT_ALREADY_EXISTS_ON_THE_ADDRESS",
	//"DEBUGGER_ERROR_BREAKPOINT_ID_NOT_FOUND",
	//"DEBUGGER_ERROR_BREAKPOINT_ALREADY_DISABLED",
	//"DEBUGGER_ERROR_BREAKPOINT_ALREADY_ENABLED",
	//"DEBUGGER_ERROR_MEMORY_TYPE_INVALID",
	//"DEBUGGER_ERROR_INVALID_PROCESS_ID",
	//"DEBUGGER_ERROR_EVENT_IS_NOT_APPLIED",
	//"DEBUGGER_ERROR_DETAILS_OR_SWITCH_PROCESS_INVALID_PARAMETER",
	//"DEBUGGER_ERROR_DETAILS_OR_SWITCH_THREAD_INVALID_PARAMETER",
	//"DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_FOR_A_SINGLE_PAGE_IS_HIT",
	//"DEBUGGER_ERROR_PRE_ALLOCATED_BUFFER_IS_EMPTY",
	//"DEBUGGER_ERROR_EPT_COULD_NOT_SPLIT_THE_LARGE_PAGE_TO_4KB_PAGES",
	//"DEBUGGER_ERROR_EPT_FAILED_TO_GET_PML1_ENTRY_OF_TARGET_ADDRESS",
	//"DEBUGGER_ERROR_EPT_MULTIPLE_HOOKS_IN_A_SINGLE_PAGE",
	//"DEBUGGER_ERROR_COULD_NOT_BUILD_THE_EPT_HOOK",
	//"DEBUGGER_ERROR_COULD_NOT_FIND_ALLOCATION_TYPE",
	//"DEBUGGER_ERROR_INVALID_TEST_QUERY_INDEX",
	//"DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_TARGET_USER_MODE_PROCESS",
	//"DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS_ENTRYPOINT_NOT_REACHED",
	//"DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS",
	//"DEBUGGER_ERROR_FUNCTIONS_FOR_INITIALIZING_PEB_ADDRESSES_ARE_NOT_INITIALIZED",
	//"DEBUGGER_ERROR_UNABLE_TO_DETECT_32_BIT_OR_64_BIT_PROCESS",
	//"DEBUGGER_ERROR_UNABLE_TO_KILL_THE_PROCESS",
	//"DEBUGGER_ERROR_INVALID_THREAD_DEBUGGING_TOKEN",
	//"DEBUGGER_ERROR_UNABLE_TO_PAUSE_THE_PROCESS_THREADS",
	//"DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_AN_ALREADY_ATTACHED_PROCESS",
	//"DEBUGGER_ERROR_THE_USER_DEBUGGER_NOT_ATTACHED_TO_THE_PROCESS",
	//"DEBUGGER_ERROR_UNABLE_TO_DETACH_AS_THERE_ARE_PAUSED_THREADS",
	//"DEBUGGER_ERROR_UNABLE_TO_SWITCH_PROCESS_ID_OR_THREAD_ID_IS_INVALID",
	//"DEBUGGER_ERROR_UNABLE_TO_SWITCH_THERE_IS_NO_THREAD_ON_THE_PROCESS",
	//"DEBUGGER_ERROR_UNABLE_TO_GET_MODULES_OF_THE_PROCESS",
	//"DEBUGGER_ERROR_UNABLE_TO_GET_CALLSTACK",
	//"DEBUGGER_ERROR_UNABLE_TO_QUERY_COUNT_OF_PROCESSES_OR_THREADS",
	//"DEBUGGER_ERROR_USING_SHORT_CIRCUITING_EVENT_WITH_POST_EVENT_MODE_IS_FORBIDDEDN",
	//"DEBUGGER_ERROR_UNKNOWN_TEST_QUERY_RECEIVED",
	//"DEBUGGER_ERROR_READING_MEMORY_INVALID_PARAMETER",
	//"DEBUGGER_ERROR_THE_TRAP_FLAG_LIST_IS_FULL",
	//"DEBUGGER_ERROR_UNABLE_TO_KILL_THE_PROCESS_DOES_NOT_EXISTS",
	//"DEBUGGER_ERROR_MODE_EXECUTION_IS_INVALID",
	//"DEBUGGER_ERROR_PROCESS_ID_CANNOT_BE_SPECIFIED_WHILE_APPLYING_EVENT_FROM_VMX_ROOT_MODE",
	//"DEBUGGER_ERROR_INSTANT_EVENT_PREALLOCATED_BUFFER_IS_NOT_ENOUGH_FOR_EVENT_AND_CONDITIONALS",
	//"DEBUGGER_ERROR_INSTANT_EVENT_REGULAR_PREALLOCATED_BUFFER_NOT_FOUND",
	//"DEBUGGER_ERROR_INSTANT_EVENT_BIG_PREALLOCATED_BUFFER_NOT_FOUND",
	//"DEBUGGER_ERROR_UNABLE_TO_CREATE_ACTION_CANNOT_ALLOCATE_BUFFER",
	//"DEBUGGER_ERROR_INSTANT_EVENT_ACTION_REGULAR_PREALLOCATED_BUFFER_NOT_FOUND",
	//"DEBUGGER_ERROR_INSTANT_EVENT_ACTION_BIG_PREALLOCATED_BUFFER_NOT_FOUND",
	//"DEBUGGER_ERROR_INSTANT_EVENT_PREALLOCATED_BUFFER_IS_NOT_ENOUGH_FOR_ACTION_BUFFER",
	//"DEBUGGER_ERROR_INSTANT_EVENT_REQUESTED_OPTIONAL_BUFFER_IS_BIGGER_THAN_DEBUGGERS_SEND_RECEIVE_STACK",
	//"DEBUGGER_ERROR_INSTANT_EVENT_REGULAR_REQUESTED_SAFE_BUFFER_NOT_FOUND",
	//"DEBUGGER_ERROR_INSTANT_EVENT_BIG_REQUESTED_SAFE_BUFFER_NOT_FOUND",
	//"DEBUGGER_ERROR_INSTANT_EVENT_PREALLOCATED_BUFFER_IS_NOT_ENOUGH_FOR_REQUESTED_SAFE_BUFFER",
	//"DEBUGGER_ERROR_UNABLE_TO_ALLOCATE_REQUESTED_SAFE_BUFFER",
	//"DEBUGGER_ERROR_COULD_NOT_FIND_PREACTIVATION_TYPE",
	//"DEBUGGER_ERROR_THE_MODE_EXEC_TRAP_IS_NOT_INITIALIZED",
	//"DEBUGGER_ERROR_THE_TARGET_EVENT_IS_DISABLED_BUT_CANNOT_BE_CLEARED_PRIRITY_BUFFER_IS_FULL",
	//"DEBUGGER_ERROR_NOT_ALL_CORES_ARE_LOCKED_FOR_APPLYING_INSTANT_EVENT",
	//"DEBUGGER_ERROR_TARGET_SWITCHING_CORE_IS_NOT_LOCKED",
	//"DEBUGGER_ERROR_INVALID_PHYSICAL_ADDRESS",
	//"SIZEOF_DEBUGGER_MODIFY_EVENTS",
	//"SIZEOF_REGISTER_EVENT",
	//"DEFAULT_INITIAL_DEBUGGEE_TO_DEBUGGER_OFFSET",
	//"DEFAULT_INITIAL_DEBUGGER_TO_DEBUGGEE_OFFSET",
	//"IOCTL_REGISTER_EVENT",
	//"IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL",
	//"IOCTL_TERMINATE_VMX",
	//"IOCTL_DEBUGGER_READ_MEMORY",
	//"IOCTL_DEBUGGER_READ_OR_WRITE_MSR",
	//"IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS",
	//"IOCTL_DEBUGGER_REGISTER_EVENT",
	//"IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT",
	//"IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER",
	//"IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS",
	//"IOCTL_DEBUGGER_EDIT_MEMORY",
	//"IOCTL_DEBUGGER_SEARCH_MEMORY",
	//"IOCTL_DEBUGGER_MODIFY_EVENTS",
	//"IOCTL_DEBUGGER_FLUSH_LOGGING_BUFFERS",
	//"IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS",
	//"IOCTL_DEBUGGER_PRINT",
	//"IOCTL_PREPARE_DEBUGGEE",
	//"IOCTL_PAUSE_PACKET_RECEIVED",
	//"IOCTL_SEND_SIGNAL_EXECUTION_IN_DEBUGGEE_FINISHED",
	//"IOCTL_SEND_USERMODE_MESSAGES_TO_DEBUGGER",
	//"IOCTL_SEND_GENERAL_BUFFER_FROM_DEBUGGEE_TO_DEBUGGER",
	//"IOCTL_PERFROM_KERNEL_SIDE_TESTS",
	//"IOCTL_RESERVE_PRE_ALLOCATED_POOLS",
	//"IOCTL_SEND_USER_DEBUGGER_COMMANDS",
	//"IOCTL_GET_DETAIL_OF_ACTIVE_THREADS_AND_PROCESSES",
	//"IOCTL_GET_USER_MODE_MODULE_DETAILS",
	//"IOCTL_QUERY_COUNT_OF_ACTIVE_PROCESSES_OR_THREADS",
	//"IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES",
	//"IOCTL_QUERY_CURRENT_PROCESS",
	//"IOCTL_QUERY_CURRENT_THREAD",
	//"IOCTL_REQUEST_REV_MACHINE_SERVICE",
	//"IOCTL_DEBUGGER_BRING_PAGES_IN",
	//"IOCTL_PREACTIVATE_FUNCTIONALITY",
	//"SIZEOF_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS",
	//"SIZEOF_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS",
	//"SIZEOF_DEBUGGER_PAGE_IN_REQUEST",
	//"SIZEOF_REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST",
	//"SIZEOF_DEBUGGER_DT_COMMAND_OPTIONS",
	//"SIZEOF_DEBUGGER_PREALLOC_COMMAND",
	//"SIZEOF_DEBUGGER_PREACTIVATE_COMMAND",
	//"SIZEOF_DEBUGGER_READ_MEMORY",
	//"SIZEOF_DEBUGGER_FLUSH_LOGGING_BUFFERS",
	//"SIZEOF_DEBUGGER_TEST_QUERY_BUFFER",
	//"SIZEOF_DEBUGGER_PERFORM_KERNEL_TESTS",
	//"SIZEOF_DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL",
	//"SIZEOF_DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER",
	//"SIZEOF_DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER",
	//"SIZEOF_DEBUGGER_READ_AND_WRITE_ON_MSR",
	//"SIZEOF_DEBUGGER_EDIT_MEMORY",
	//"SIZEOF_DEBUGGER_SEARCH_MEMORY",
	//"SIZEOF_DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE",
	//"SIZEOF_DEBUGGER_PREPARE_DEBUGGEE",
	//"SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS",
	//"SIZEOF_DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS",
	//"SIZEOF_DEBUGGER_CALLSTACK_REQUEST",
	//"SIZEOF_USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS",
	//"SIZEOF_DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET",
	//"SIZEOF_DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET",
	//"DEBUGGER_REMOTE_TRACKING_DEFAULT_COUNT_OF_STEPPING",
	"LogDebugInfo(format,",
}
