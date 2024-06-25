package libhyperdbg

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"github.com/ddkwork/golibrary/mylog"
)

// 定义SystemInfo结构
type SystemInfo struct {
	OemId                     uint32
	PageSize                  uint32
	MinimumApplicationAddress *byte
	MaximumApplicationAddress *byte
	ActiveProcessorMask       *uint64
	NumberOfProcessors        uint32
	ProcessorType             uint32
	AllocationGranularity     uint32
	ProcessorLevel            uint16
	ProcessorRevision         uint16
}

// GetSystemInfo 填充 SystemInfo 结构体
func GetSystemInfo(lpSystemInfo *SystemInfo) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("GetSystemInfo")
	proc.Call(uintptr(unsafe.Pointer(lpSystemInfo)))
}

// ReadBuffer 读取指定地址的size大小的数据
func readBuffer(hFileMapping syscall.Handle, address uintptr, size int) ([]byte, error) {
	// 获取系统页面大小
	var sysInfo SystemInfo
	GetSystemInfo(&sysInfo)
	allocationGranularity := int(sysInfo.AllocationGranularity)

	// 计算对齐的文件偏移
	alignedOffset := address & ^uintptr(allocationGranularity-1)
	offsetWithinPage := int(address - alignedOffset)
	mapViewSize := size + offsetWithinPage

	// 映射文件的一部分到内存
	view := mylog.Check2(syscall.MapViewOfFile(hFileMapping, syscall.FILE_MAP_READ, 0, uint32(alignedOffset), uintptr(mapViewSize)))

	defer syscall.UnmapViewOfFile(view)

	// 读取数据到buffer
	data := (*[1 << 30]byte)(unsafe.Pointer(view))[:mapViewSize:mapViewSize]
	buffer := make([]byte, size)
	copy(buffer, data[offsetWithinPage:offsetWithinPage+size])

	return buffer, nil
}

func main() {
	// 打开PE文件
	file := mylog.Check2(os.Open("path_to_your_pe_file.exe"))

	defer file.Close()

	// 获取文件句柄
	fileHandle := syscall.Handle(file.Fd())

	// 创建文件映射对象
	hFileMapping := mylog.Check2(syscall.CreateFileMapping(fileHandle, nil, syscall.PAGE_READONLY, 0, 0, nil))

	defer syscall.CloseHandle(hFileMapping)

	// 映射文件的头部到内存
	view := mylog.Check2(syscall.MapViewOfFile(hFileMapping, syscall.FILE_MAP_READ, 0, 0, 0))

	defer syscall.UnmapViewOfFile(view)

	// 获取PE头信息
	dosHeader := (*IMAGE_DOS_HEADER)(unsafe.Pointer(view))
	ntHeaders := (*IMAGE_NT_HEADERS)(unsafe.Pointer(uintptr(view) + uintptr(dosHeader.E_lfanew)))

	// 获取 ImageBase 和 AddressOfEntryPoint
	imageBase := ntHeaders.OptionalHeader.ImageBase
	oep := ntHeaders.OptionalHeader.AddressOfEntryPoint

	// 计算实际OEP地址
	actualOEP := uintptr(imageBase) + uintptr(oep)

	// 读取512字节数据
	buffer := mylog.Check2(readBuffer(hFileMapping, actualOEP, 512))

	// 打印读取的数据
	for i := 0; i < len(buffer); i++ {
		fmt.Printf("%02X ", buffer[i])
		if i%16 == 15 {
			fmt.Println()
		}
	}
}

// 定义PE头的结构体
type IMAGE_DOS_HEADER struct {
	E_magic    uint16
	E_cblp     uint16
	E_cp       uint16
	E_crlc     uint16
	E_cparhdr  uint16
	E_minalloc uint16
	E_maxalloc uint16
	E_ss       uint16
	E_sp       uint16
	E_csum     uint16
	E_ip       uint16
	E_cs       uint16
	E_lfarlc   uint16
	E_ovno     uint16
	E_res      [4]uint16
	E_oemid    uint16
	E_oeminfo  uint16
	E_res2     [10]uint16
	E_lfanew   int32
}

type IMAGE_FILE_HEADER struct {
	Machine              uint16
	NumberOfSections     uint16
	TimeDateStamp        uint32
	PointerToSymbolTable uint32
	NumberOfSymbols      uint32
	SizeOfOptionalHeader uint16
	Characteristics      uint16
}

type IMAGE_OPTIONAL_HEADER struct {
	// 标准字段
	Magic                   uint16
	MajorLinkerVersion      byte
	MinorLinkerVersion      byte
	SizeOfCode              uint32
	SizeOfInitializedData   uint32
	SizeOfUninitializedData uint32
	AddressOfEntryPoint     uint32
	BaseOfCode              uint32
	BaseOfData              uint32

	// NT附加字段
	ImageBase                   uint32
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
	SizeOfStackReserve          uint32
	SizeOfStackCommit           uint32
	SizeOfHeapReserve           uint32
	SizeOfHeapCommit            uint32
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               [16]IMAGE_DATA_DIRECTORY
}

type IMAGE_DATA_DIRECTORY struct {
	VirtualAddress uint32
	Size           uint32
}

type IMAGE_NT_HEADERS struct {
	Signature      uint32
	FileHeader     IMAGE_FILE_HEADER
	OptionalHeader IMAGE_OPTIONAL_HEADER
}
