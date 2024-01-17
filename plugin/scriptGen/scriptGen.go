package scriptGen

import (
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"

	"path/filepath"
)

type (
	Interface interface {
		FormatString() (ok bool)
		FormatValue() (ok bool)
		ListenPipe() (ok bool)
		VmwareServer() (ok bool)
		SymbolDownload() (ok bool)
		SymbolLoad() (ok bool)
		SymbolReLoad() (ok bool)
		//LoadVmm() (ok bool)
		UnLoadVmm() (ok bool)
		//EptHook() (ok bool)
		//EptHook2() (ok bool)
		//StringLen = strlen(@rcx + 10);

		RunCli() (ok bool)        //hyperdbg-cli.exe
		SyscallDumper() (ok bool) //SyscallDumper.exe -d
		ConnectLocal() (ok bool)  //.connect local
		LoadVmm() (ok bool)       //load vmm
		UnloadVmm() (ok bool)     //unload vmm
		SymbolsPath(path string) (ok bool)
		/*
			这个命令比!"epthook "快得多，但它有以下限制。
			它只能在内核地址中使用，这意味着你不能在用户模式地址中使用它。
			你只能在一个内存页中使用一个钩子。例如，如果你把一个钩子放在ffff80126551006上，那么你就不能在ffff80126551000到ffff80126551fff的范围内放另一个钩子，因为它在同一个页面内（0x1000或4096字节）。
			它有经典迂回钩的限制。我们为我们的detours钩子修补了19个字节，所以当你把钩子放在汇编的任何地方时，你必须确保在钩子地址之后的19个字节内没有相对跳转或相对调用。大多数时候，函数的起始地址是与detours兼容的（不会以相对跳转或相对调用开始），特别是在x64快速调用函数中；因此，函数的起始地址是放置这些隐藏钩子的好位置。
			{% hint style="danger" %} 你不应该在同一页面（4KB）上同时使用任何一个！Monitor、！epthook和！epthook2的命令。例如，当你把一个隐藏的钩子(!epthook2)放在0x10000005时，你不应该在0x10000000到0x10000fff开始的地址上使用任何!"monitor "或!epthook命令。
			你可以在同一个页面上的两个或更多的地址上使用！epthook（只是！epthook而不是！epthook2，也不是！monitor）（意味着你可以在一个页面之间的地址上多次使用！epthook或者在一个页面上放置多个隐藏断点）。但是你不能在同一个页面上两次使用 !monitor 或 !epthook2。{% endhint %}
			这是一个事件命令，但是在当前版本的HyperDbg中（在调试器模式下），这个命令将继续调试一段时间；但是，你可以使用这个技巧来确保你不会丢失任何事件。
			!epthook2 nt!NtDeviceIoControlFile script {file 1.ds}
			HyperDbg> !epthook2 fffff800`4ed6f010 script {file:c:\users\sina\desktop\script.txt}
		*/
		EptHook2() (ok bool)               //inline hook，19字节hook，r0，不通用，快，返回地址存在相对偏移刚好在19字节之内会覆盖的危险，比如 zwUnloadDriver的原始代码，最后是指令“lea rax，KiServiceLinkage”
		EptHook() (ok bool)                //cc break,通用，慢，支持r0-3
		HookMidForDataRecovery() (ok bool) //cc break
		SaveLog() (ok bool)                //
		ProcessNameFilter(processName string, fn string) (ok bool)
	}
	object struct {
		s *stream.Stream
	}
)

func New() *object {
	return &object{
		s: stream.New(""),
	}
}

func (o *object) ProcessNameFilter(processName string, fn string) (ok bool) {
	o.s.WriteStringLn("\nif(")
	processName = processName[:len(processName)-len(filepath.Ext(processName))]
	b := stream.New(processName)
	mylog.HexDump("pname test", b.Bytes())
	for i, b2 := range b.Bytes() {
		s := `db($pname +` + fmt.Sprint(i) + `) ==` + fmt.Sprintf("%x", b2)
		if i < b.Len()-1 {
			s += `&&`
		}
		o.s.WriteStringLn(s)
	}

	o.s.WriteStringLn("){")
	o.s.WriteStringLn(" printf(\"Process name starts with: %s\\n\",$pname)")
	o.s.WriteStringLn("}")
	mylog.Info("pname gen", o.s.String())
	return true
}

func (o *object) FormatString() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (o *object) FormatValue() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (o *object) ListenPipe() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (o *object) VmwareServer() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (o *object) SymbolDownload() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (o *object) SymbolLoad() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (o *object) SymbolReLoad() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (o *object) UnLoadVmm() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (o *object) RunCli() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (o *object) SyscallDumper() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (o *object) ConnectLocal() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (o *object) LoadVmm() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (o *object) UnloadVmm() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (o *object) SymbolsPath(path string) (ok bool) {
	//.sympath SRV*c:\Symbols*https://msdl.microsoft.com/download/symbols
	s := ".sympath SRV*" + path + "*https://msdl.microsoft.com/download/symbols"
	println(s)
	return true
}

func (o *object) EptHook2() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (o *object) EptHook() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (o *object) HookMidForDataRecovery() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (o *object) SaveLog() (ok bool) {
	//TODO implement me
	panic("implement me")
}
