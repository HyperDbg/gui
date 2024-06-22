package sdk

import (
	"io/fs"
	"path/filepath"
	"strings"
	"testing"
	"unicode"

	"github.com/ddkwork/golibrary/stream/orderedmap"

	"github.com/can1357/gengo/clang"
	"github.com/can1357/gengo/gengo"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
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

// ContainsLetter 检查字符串中是否包含字母
func ContainsLetter(s string) bool {
	if strings.HasPrefix(s, "0x") {
		return false
	}
	for _, char := range s {
		if unicode.IsLetter(char) {
			return true
		}
	}
	return false
}

func TestBindMacros(t *testing.T) {
	headerFile := "merged_headers.h"
	macros := extractMacros(stream.NewBuffer(headerFile).ToLines())
	mylog.Trace("number of macros", macros.Len())

	for _, p := range macros.List() {
		mylog.Warning(p.Key, p.Value)
	}
	return

	g := stream.NewGeneratedFile()
	g.P("package HPRDBGCTRL")
	g.P()
	g.P("const (")
	g.P("MaxSerialPacketSize =10 * NORMAL_PAGE_SIZE") // todo need first define NORMAL_PAGE_SIZE
	g.P("PAGE_SIZE = 4096")

	for _, p := range macros.List() {
		line := strings.TrimPrefix(p.Key, "#define ")
		line = strings.TrimSpace(line)
		if strings.Count(line, " ") == 1 {
			split := strings.Split(line, " ")
			split[1] = strings.TrimSuffix(split[1], "ull")
			split[1] = strings.TrimSuffix(split[1], "U")
			if ContainsLetter(split[1]) {
				break
			}
			g.P(split[0] + "=" + split[1])
			macros.Delete(p.Key)
		}
	}
	g.P(")")
	stream.WriteGoFile("tmp/vars.go", g.Buffer)

	for _, p := range macros.List() {
		mylog.Todo(p.Key + " need handle")
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
