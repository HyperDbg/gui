package hookApiMgr

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"github.com/ddkwork/golibrary/stream/clang"
)

type (
	Interface interface {
		DecodeStack(api, stack string, args ...string) (argList []ArgList, ok bool)

		IopXxxControlFile() (ok bool)
		IopQueryXxxInformation() (ok bool)
		IopCreateFile() (ok bool) // NtCreateFile
		NtOpenFile() (ok bool)
		IopCloseFile() (ok bool)

		IopParseDevice() (ok bool)
		IopParseFile() (ok bool)
		IopQueryName() (ok bool)
		IopQueryNameInternal() (ok bool)
		IopReadyDeviceObjects() (ok bool)
		IopfCallDriver() (ok bool) // device filter
		IopGetDriverPathInformation() (ok bool)
		CpuID() (ok bool)

		// NsiEnumerateObjectsAllParameters
		// NsiEnumerateObjectsAllParametersEx
		// NsiEnumerateObjectsAllPersistentParametersWithMask
		// NTSTATUS NsiEnumerateObjectsAllParametersEx(PNSI_ENUMERATE_OBJECTS_ALL_PRAMETERS_EX param) (netio.sys)
		// or IOCTL 0x12001B to \device\nsi (nsiproxy.sys) (in, out PNSI_ENUMERATE_OBJECTS_ALL_PRAMETERS_EX, note for wow64 processes that struct will have different layout)

		// bp nt!NsiEnumerateObjectsAllParametersEx
		// netio.sys!NsiEnumerateObjectsAllParametersEx

		// 2. 在 NT6 中，我希望您可以从内核模式获取此信息
		// IPHelper，因为您可以从用户模式 ​​IPHelper 获取它
		// MIB_TCPROW_OWNER_MODULE / TCPIP_OWNER_MODULE_BASIC_INFO。
		// 事实证明，GetTcpTable 和变体最终调用了 NsiAllocateAndGetTable。 在用户空间中，此函数在 nsi.dll 中实现，该实现通过调用 \\Device\Nsi 来满足 NsiEnumerateObjectsAllParameters 来工作。 NSI 代理驱动程序验证参数并代理对 netio.sys 的调用，该调用实际上导出了一个名为 NsiEnumerateObjectsAllParameters 的函数。 查看 netio.sys 的导出列表，我发现了一个名为 NsiAllocateAndGetTable 的导出，这正是我在用户空间中跟踪的函数。
		// git clone --recursive git://source.winehq.org/git/wine.git
		MacAddress() (ok bool) // nsi.dll

		AfdSend() (ok bool)               // nsi
		NtCreateNamedPipeFile() (ok bool) // todo steam dataShare hook and transport to mitmproxy

		HookMidForDataRecovery() (ok bool) // nsi
		Themida() (ok bool)                //
		Vmprotect() (ok bool)              //
		Vmware() (ok bool)                 //

		// net filter vxk see mitmproxy
		// eb output buffer l 512 --->hexdump
		// pid
	}
	ArgList struct {
		Offset string
		VA     string
		Arg    string
	}
	object struct{}
)

func New() *object { return &object{} }

func (o *object) DecodeStack(api, stack string, argsInput ...string) (argList []ArgList, ok bool) {
	lines := stream.NewBuffer(stack).ToLines()

	argList = make([]ArgList, 0)

	fnCut := func(orig string) (list []string) {
		list = make([]string, 0)
		orig = strings.TrimSpace(orig)
		index1 := strings.Index(orig, " ") + 1
		list = append(list, strings.TrimSpace(orig[:index1]))

		orig = orig[index1:]
		orig = strings.TrimSpace(orig)
		index2 := strings.Index(orig, " ") + 1
		list = append(list, strings.TrimSpace(orig[:index2]))
		return
	}

	for _, line := range lines {
		if line == "" {
			continue
		}
		cut := fnCut(line)
		if cut[0] == "[$+000]" {
			from := "from "
			index := strings.Index(line, from)
			line = line[index+len(from):]

			index = strings.Index(line, " ")
			line = line[:index]
			mylog.Info(api, "called\tby\t"+line)
			continue
		}

		index := strings.Index(cut[0], "+")
		Offset := cut[0][index+1:]
		Offset = Offset[:len(Offset)-1]

		list := ArgList{
			Offset: Offset,
			VA:     cut[1],
			Arg:    "",
		}
		argList = append(argList, list)
	}
	for i, s := range argsInput {
		argList[i].Arg = s
	}
	return argList, true
}

func (o *object) IopXxxControlFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *object) TestIopXxxControlFile() (ok bool) {
	api := "IopXxxControlFile"
	stack := `
[$+000]   fffff8032f4c8811    (from nt!NtDeviceIoControlFile+0x51 <fffff8032f4c8811>)
[$+008]      0000000000000000 FileHandle
[$+010]      ffff9a03d06c1b60 EventCallBack
[$+018]      0000000000000001 ApcRoutine
[$+020]      0000000000000000 ApcContext
[$+028]      000000e4de1fe390 IoStatusBlock
[$+030]      00000000001b0044 IoControlCode
[$+038]      000000e4de1fe3b8 InputBuffer
[$+040]      ffffb08e00000004 InputBufferLength
[$+048]      0000000000000000 OutputBuffer
[$+050]      0000000000000000 OutputBufferLength
[$+058]      ffffb08e00000001 DeviceIoControl

`
	argList, ok := o.DecodeStack(api, stack,
		"FileHandle", // todo add type
		"EventCallBack",
		"ApcRoutine",
		"ApcContext",
		"IoStatusBlock",
		"IoControlCode",
		"InputBuffer",
		"InputBufferLength",
		"OutputBuffer",
		"OutputBufferLength",
		"DeviceIoControl",
	)
	if !ok {
		return
	}
	s := stream.NewBuffer("")
	s.WriteStringLn("void format() {")
	s.WriteStringLn("!epthook nt!" + api + " script {")
	pname := `printf("pname\t%s\n",$pname);`
	s.WriteStringLn(pname)
	for _, list := range argList {
		s.WriteStringLn(`printf("` + list.Arg + `\t%llx\n",` + `dq(rsp +` + list.Offset + `));`)
	}
	s.WriteStringLn(`printf("\n");`)
	s.WriteStringLn("}")
	s.WriteStringLn("}")
	s.WriteStringLn(`
/*
.connect local
load vmm
.sympath SRV*c:\Symbols*https://msdl.microsoft.com/download/symbol
.sym load
.sym reload
.sym download
unload vmm

first in get stack on debug module
.debug remote namedpipe \\.\pipe\HyperDbgDebug
.debug prepare serial 115200 com1
u nt!IopXxxControlFile l 74f
bp nt!IopXxxControlFile
g
kq l 60
*/`)
	dsName := api + ".ds"
	stream.WriteTruncate(dsName, s.String())
	abs, err := filepath.Abs(dsName)
	mylog.Check(err)
	c := clang.New()
	if !c.WriteClangFormatBody(filepath.Dir(abs)) {
		return
	}
	c.Format(abs)
	b, err := os.ReadFile(abs)
	mylog.Check(err)
	all := strings.ReplaceAll(string(b), " nt !", " nt!")
	return stream.WriteTruncate(abs, all)
}

/*
pname   ssd.exe
pname   ssd.exe
FileHandle      0
FileHandle      1
EventCallBack   fffffae09df8b4f3
ApcRoutine      ffff8789aaaa6080
ApcContext      262f41796b0
ApcContext      1
IoStatusBlock   8806fff9b0
IoStatusBlock   8806fff9d0
IoControlCode   500016
IoControlCode   12001b
InputBuffer     8806fff920
InputBufferLength       70
OutputBuffer    8806fff920
OutputBufferLength      70
DeviceIoControl ffff878900000001
DeviceIoControl ffffe000d27c0101

pname   ssd.exe
FileHandle      1
EventCallBack   ffff8400a2c92b60
ApcRoutine      1
ApcRoutine      1
ApcContext      1
IoStatusBlock   8806bff3f0
IoStatusBlock   bd044ff670
IoControlCode   50000f
IoControlCode   500006
InputBuffer     bd044ffa70
InputBufferLength       28
OutputBuffer    bd044ffbf8
OutputBufferLength      90
OutputBufferLength      68
OutputBufferLength      0
DeviceIoControl 1

pname   ssd.exe
pname   ssd.exe
FileHandle      1
EventCallBack   ffff8400a2c92b60
EventCallBack   1
ApcRoutine      0
ApcRoutine      ffff8400a2c92ae0
ApcContext      ffffa7806000db98
IoStatusBlock   8806bff410
IoControlCode   ffffa7d300500016
IoControlCode   12000f
InputBuffer     8806bff400
InputBuffer     8806bff440
InputBufferLength       30
OutputBuffer    0
OutputBuffer    8806bff400
OutputBufferLength      68
DeviceIoControl ffff878900000001
DeviceIoControl 1
DeviceIoControl 1
*/

func (o *object) IopQueryXxxInformation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *object) IopCreateFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *object) NtOpenFile() (ok bool) {
	/*
		!epthook nt!NtOpenFile script {
			printf("%ws\n", dq(poi(r8 + 10) + 0x8));
		}
	*/

	// TODO implement me
	panic("implement me")
}

func (o *object) IopCloseFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *object) IopParseDevice() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *object) IopParseFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *object) IopQueryName() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *object) IopQueryNameInternal() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *object) IopReadyDeviceObjects() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *object) IopfCallDriver() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *object) IopGetDriverPathInformation() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *object) CpuID() (ok bool) {
	/*
		void format() {
		    !cpuid script {
		        printf("$pname\t%s\n", $pname);
		        if ($context == 1) {
		            @rax = 0x000306A9;
		            @rcx = 0x3DBAE3BF;
		            @rdx = 0xBFEBFBFF;
		        }
		        printf("rax,rbx,rcx,rdx: %llx\t%llx\t%llx\t%llx\t\n", @rax, @rbx, @rcx, @rdx);
		    }
		}

		void format() {
		    !cpuid script {
		        if (db($pname) == 73 && db($pname + 1) == 73 && db($pname + 2) == 64 && db($pname + 3) == 2e && db($pname + 3) == 65 && db($pname + 3) == 78 && db($pname + 3) == 65) {
		            printf("$pname\t%s\n", $pname);
		            if ($context == 1) {
		                @rax = 0x000306A9;
		                @rcx = 0x3DBAE3BF;
		                @rdx = 0xBFEBFBFF;
		            }
		            printf("rax,rbx,rcx,rdx: %llx\t%llx\t%llx\t%llx\t\n", @rax, @rbx, @rcx, @rdx);
		        }
		    }
		}
	*/
	//TODO implement me
	panic("implement me")
}

func (o *object) MacAddress() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *object) AfdSend() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *object) NtCreateNamedPipeFile() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *object) HookMidForDataRecovery() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *object) Themida() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *object) Vmprotect() (ok bool) {
	// TODO implement me
	panic("implement me")
}

func (o *object) Vmware() (ok bool) {
	// TODO implement me
	panic("implement me")
}
