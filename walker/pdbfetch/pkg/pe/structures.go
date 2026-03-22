package pe

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"strings"
)

// DOS Header
// noinspection GoSnakeCaseUsage
type ImageDosHeader struct {
	E_magic    uint16
	E_cblp     uint16
	E_cp       uint16
	E_crlc     uint16
	E_cparhd   uint16
	E_minalloc uint16
	E_maxalloc uint16
	E_ss       uint16
	E_sp       uint16
	E_csum     uint16
	E_ip       uint16
	E_cs       uint16
	E_lfarlc   uint16
	E_ovno     uint16
	E_res      [8]uint8
	E_oemid    uint16
	E_oeminfo  uint16
	E_res2     [20]uint8
	E_lfanew   uint32
}

type DosHeader struct {
	ImageDosHeader
	fileOffset int
	size       int
	flags      map[string]bool
}

func NewDosHeader(fileOffset int) (header *DosHeader) {
	header = new(DosHeader)
	header.flags = make(map[string]bool)
	header.size = binary.Size(header.ImageDosHeader)
	header.fileOffset = fileOffset
	return header
}

func (h *DosHeader) String() string {
	return structString(h.fileOffset, "IMAGE_DOS_HEADER", h.ImageDosHeader) + flagString(h.flags)
}

type ImageNTHeader struct {
	Signature uint32
}

type NTHeader struct {
	ImageNTHeader
	fileOffset int
	size       int
	flags      map[string]bool
}

// NT Header

func NewNTHeader(fileOffset int) (header *NTHeader) {
	header = new(NTHeader)
	header.flags = make(map[string]bool)
	header.size = binary.Size(header.ImageNTHeader)
	header.fileOffset = fileOffset
	return header
}

func (h *NTHeader) String() string {
	return structString(h.fileOffset, "IMAGE_NT_HEADER", h.ImageNTHeader) + flagString(h.flags)
}

// File Header
type ImageFileHeader struct {
	Machine              uint16
	NumberOfSections     uint16
	TimeDateStamp        uint32
	PointerToSymbolTable uint32
	NumberOfSymbols      uint32
	SizeOfOptionalHeader uint16
	Characteristics      uint16
}

type FileHeader struct {
	ImageFileHeader

	fileOffset int
	size       int
	flags      map[string]bool
}

func NewFileHeader(fileOffset int) (header *FileHeader) {
	header = new(FileHeader)
	header.flags = make(map[string]bool)
	header.size = IMAGE_SIZEOF_FILE_HEADER
	header.fileOffset = fileOffset
	return header
}

func (f *FileHeader) String() string {
	return structString(f.fileOffset, "IMAGE_FILE_HEADER", f.ImageFileHeader) + flagString(f.flags)
}

// Optional Header
type ImageOptionalHeader32 struct {
	Magic                       uint16
	MajorLinkerVersion          uint8
	MinorLinkerVersion          uint8
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	BaseOfData                  uint32
	ImageBase                   uint32
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Reserved1                   uint32
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
}

type OptionalHeader32 struct {
	ImageOptionalHeader32
	DataDirs map[string]*DataDirectory

	fileOffset int
	size       int
	flags      map[string]bool
}

func NewOptionalHeader32(fileOffset int) (header *OptionalHeader32) {
	header = new(OptionalHeader32)
	header.DataDirs = make(map[string]*DataDirectory)
	header.flags = make(map[string]bool)
	header.size = binary.Size(header.ImageOptionalHeader32)
	header.fileOffset = fileOffset
	return header
}

func (o *OptionalHeader32) String() string {
	return structString(o.fileOffset, "OPTIONAL_HEADER", o.ImageOptionalHeader32) + flagString(o.flags)
}

type ImageOptionalHeader64 struct {
	Magic                       uint16
	MajorLinkerVersion          uint8
	MinorLinkerVersion          uint8
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	BaseOfData                  uint32
	ImageBase                   uint32
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Reserved1                   uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
	SizeOfStackReserve          uint64 // Different after this point, specific checks needed
	SizeOfStackCommit           uint64
	SizeOfHeapReserve           uint64
	SizeOfHeapCommit            uint64
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
}

type OptionalHeader64 struct {
	ImageOptionalHeader64
	DataDirs map[string]*DataDirectory

	fileOffset int
	size       int
	flags      map[string]bool
}

func NewOptionalHeader64(fileOffset int) (header *OptionalHeader64) {
	header = new(OptionalHeader64)
	header.DataDirs = make(map[string]*DataDirectory)
	header.flags = make(map[string]bool)
	header.size = binary.Size(header.ImageOptionalHeader64)
	header.fileOffset = fileOffset
	return header
}

func (o *OptionalHeader64) String() string {
	return structString(o.fileOffset, "OPTIONAL_HEADER64", o.ImageOptionalHeader64) + flagString(o.flags)
}

// Data directory
type ImageDataDirectory struct {
	VirtualAddress uint32
	Size           uint32
}

type DataDirectory struct {
	ImageDataDirectory
	Name string

	fileOffset int
	size       int
}

func NewDataDirectory(fileOffset int) (header *DataDirectory) {
	header = new(DataDirectory)
	header.size = binary.Size(header.ImageDataDirectory)
	header.fileOffset = fileOffset
	return header
}

func (d *DataDirectory) String() string {
	return structString(d.fileOffset, "DATA_DIRECTORY", d.ImageDataDirectory)
}

// Image Section

// noinspection GoSnakeCaseUsage
type ImageSectionHeader struct {
	Name                             [IMAGE_SIZEOF_SHORT_NAME]uint8
	Misc_VirtualSize_PhysicalAddress uint32
	VirtualAddress                   uint32
	SizeOfRawData                    uint32
	PointerToRawData                 uint32
	PointerToRelocations             uint32
	PointerToLinenumbers             uint32
	NumberOfRelocations              uint16
	NumberOfLinenumbers              uint16
	Characteristics                  uint32
}

type SectionHeader struct {
	ImageSectionHeader
	RawData     []byte
	Relocations []*Relocation

	nextHeaderRva uint32

	fileOffset int
	size       int
	flags      map[string]bool
}

func NewSectionHeader(fileOffset int) (header *SectionHeader) {
	header = new(SectionHeader)
	header.size = IMAGE_SIZEOF_SECTION_HEADER
	header.flags = make(map[string]bool)
	header.fileOffset = fileOffset
	return header
}

func (s *SectionHeader) String() string {
	var str strings.Builder
	str.WriteString(structString(s.fileOffset, "SECTION_HEADER", s.ImageSectionHeader) + flagString(s.flags))
	if len(s.Relocations) > 0 {
		str.WriteString(fmt.Sprintf("Relocations:\n%-12s%-30s%s\n", "VA", "Type", "Symbol"))
	}
	for _, reloc := range s.Relocations {
		str.WriteString(fmt.Sprintf("0x%-10X%-30s%s\n", reloc.VirtualAddress, RelocationTypesI386[reloc.Type], reloc.Symbol.Name))
	}
	return str.String()
}

// Image Symbol Table
type ImageSymbol struct {
	ShortName          [IMAGE_SIZEOF_SHORT_NAME]uint8 // if low 32bits 0 name is in string table at offset given by high 32bits.
	Value              uint32
	SectionNumber      int16
	Type               uint16
	StorageClass       int8
	NumberOfAuxSymbols int8
}

type Symbol struct {
	ImageSymbol
	Name string

	fileOffset int
	size       int
}

func NewSymbol(fileOffset int) (header *Symbol) {
	header = new(Symbol)
	header.size = IMAGE_SIZEOF_SYMBOL
	header.fileOffset = fileOffset
	return header
}

func (s *Symbol) String() string {
	return structString(s.fileOffset, "Symbol", s.ImageSymbol)
}

// Image Relocations
type ImageRelocation struct {
	VirtualAddress   uint32
	SymbolTableIndex uint32
	Type             uint16
}

type Relocation struct {
	ImageRelocation
	Symbol *Symbol

	fileOffset int
	size       int
}

func NewRelocation(fileOffset int) (header *Relocation) {
	header = new(Relocation)
	header.size = binary.Size(header.ImageRelocation)
	header.fileOffset = fileOffset
	return header
}

// BaseRelocation
type ImageBaseRelocation struct {
	VirtualAddress uint32
	SizeOfBlock    uint32
}

type BaseRelocation struct {
	ImageBaseRelocation

	fileOffset int
	size       int
}

func NewBaseRelocation(fileOffset int) (header *BaseRelocation) {
	header = new(BaseRelocation)
	header.size = binary.Size(header.ImageBaseRelocation)
	header.fileOffset = fileOffset
	return header
}

func (r *BaseRelocation) String() string {
	return structString(r.fileOffset, "BaseRelocation", r.ImageBaseRelocation)
}

// BaseRelocationEntry
type ImageBaseRelocationEntry struct {
	Data uint16
}

type BaseRelocationEntry struct {
	ImageBaseRelocationEntry
	fileOffset int
	size       int
}

func NewBaseRelocationEntry(fileOffset int) (header *BaseRelocationEntry) {
	header = new(BaseRelocationEntry)
	header.size = binary.Size(header.ImageBaseRelocationEntry)
	header.fileOffset = fileOffset
	return header
}

func (r *BaseRelocationEntry) String() string {
	return structString(r.fileOffset, "BaseRelocationEntry", r.Data)
}

// Image Import Descriptor
type ImageImportDescriptor struct {
	Characteristics uint32
	TimeDateStamp   uint32
	ForwarderChain  uint32
	Name            uint32
	FirstThunk      uint32
}

type ImportDescriptor struct {
	ImageImportDescriptor
	Module    []byte
	Imports   []*ImportData32
	Imports64 []*ImportData64

	fileOffset int
	size       int
	flags      map[string]bool
}

func NewImportDescriptor(fileOffset int) (header *ImportDescriptor) {
	header = new(ImportDescriptor)
	header.size = binary.Size(header.ImageImportDescriptor)
	header.flags = make(map[string]bool)
	header.fileOffset = fileOffset
	return header
}

func (d *ImportDescriptor) String() string {
	return structString(d.fileOffset, "IMPORT_DESCRIPTOR", d.ImageImportDescriptor) + flagString(d.flags)
}

type ImportData32 struct {
	StructTable      *ThunkData32
	StructIat        *ThunkData32
	ImportByOrdinal  bool
	Ordinal          uint32
	OrdinalOffset    int
	Hint             uint16
	Name             []byte
	NameOffset       int
	Bound            uint32
	Address          uint32
	HintNameTableRva uint32
	ThunkOffset      int
	ThunkRva         uint32
}

func (d ImportData32) String() string {
	return structString(0, "Import Data", d)
}

type ImportData64 struct {
	StructTable      *ThunkData64
	StructIat        *ThunkData64
	ImportByOrdinal  bool
	Ordinal          uint64
	OrdinalOffset    uint64
	Hint             uint16
	Name             []byte
	NameOffset       uint64
	Bound            uint64
	Address          uint64
	HintNameTableRva uint64
	ThunkOffset      uint64
	ThunkRva         uint64
}

func (d ImportData64) String() string {
	return structString(0, "Import Data 64bit", d)
}

// Delay Import Descriptor
type ImageDelayImportDescriptor struct {
	DIgrAttrs     uint32
	DIszName      uint32
	DIphmod       uint32
	DIpIAT        uint32
	DIpINT        uint32
	DIpBoundIAT   uint32
	DIpUnloadIAT  uint32
	DIdwTimeStamp uint32
}

type DelayImportDescriptor struct {
	ImageDelayImportDescriptor

	fileOffset int
	size       int
	flags      map[string]bool
}

func NewDelayImportDescriptor(fileOffset int) (header *DelayImportDescriptor) {
	header = new(DelayImportDescriptor)
	header.size = binary.Size(header.ImageDelayImportDescriptor)
	header.flags = make(map[string]bool)
	header.fileOffset = fileOffset
	return header
}

func (d *DelayImportDescriptor) String() string {
	return structString(d.fileOffset, "DELAY_IMPORT_DESCRIPTOR", d.ImageDelayImportDescriptor) + flagString(d.flags)
}

// Export Directory
type ImageExportDirectory struct {
	Characteristics       uint32
	TimeDateStamp         uint32
	MajorVersion          uint16
	MinorVersion          uint16
	Name                  uint32
	Base                  uint32
	NumberOfFunctions     uint32
	NumberOfNames         uint32
	AddressOfFunctions    uint32
	AddressOfNames        uint32
	AddressOfNameOrdinals uint32
}

type ExportDirectory struct {
	ImageExportDirectory
	Exports []*ExportData

	fileOffset int
	size       int
	flags      map[string]bool
}

func NewExportDirectory(fileOffset int) (header *ExportDirectory) {
	header = new(ExportDirectory)
	header.size = binary.Size(header.ImageExportDirectory)
	header.flags = make(map[string]bool)
	header.fileOffset = fileOffset
	return header
}

func (d *ExportDirectory) String() string {
	return structString(d.fileOffset, "EXPORT_DIRECTORY", d.ImageExportDirectory) + flagString(d.flags)
}

type ExportData struct {
	Ordinal         uint16
	OrdinalOffset   int
	Address         uint32
	AddressOffset   int
	Name            []byte //
	NameOffset      int    //
	Forwarder       []byte
	ForwarderOffset int
}

func (e ExportData) String() string {
	return structString(0, "Export Data", e)
}

// Resource Directory
type ImageResourceDirectory struct {
	Characteristics      uint32
	TimeDateStamp        uint32
	MajorVersion         uint16
	MinorVersion         uint16
	NumberOfNamedEntries uint16
	NumberOfIdEntries    uint16
}

type ResourceDirectory struct {
	ImageResourceDirectory
	fileOffset int
	size       int
	flags      map[string]bool
}

func NewResourceDirectory(fileOffset int) (header *ResourceDirectory) {
	header = new(ResourceDirectory)
	header.size = binary.Size(header.ImageResourceDirectory)
	header.flags = make(map[string]bool)
	header.fileOffset = fileOffset
	return header
}

func (r *ResourceDirectory) String() string {
	return structString(r.fileOffset, "RESOURCE_DIRECTORY", r.ImageResourceDirectory) + flagString(r.flags)
}

// Resource Directory Entry
type ImageResourceDirectoryEntry struct {
	Name         uint32
	OffsetToData uint32
}

type ResourceDirectoryEntry struct {
	ImageResourceDirectoryEntry
	fileOffset int
	size       int
}

func NewResourceDirectoryEntry(fileOffset int) (header *ResourceDirectoryEntry) {
	header = new(ResourceDirectoryEntry)
	header.size = binary.Size(header.ImageResourceDirectoryEntry)
	header.fileOffset = fileOffset
	return header
}

func (r *ResourceDirectoryEntry) String() string {
	return structString(r.fileOffset, "RESOURCE_DIRECTORY_ENTRY", r.ImageResourceDirectoryEntry)
}

// Resource Data Entry
type ImageResourceDataEntry struct {
	OffsetToData uint32 // Offset to the data of the resource.
	Size         uint32 // Size of the resource data.
	CodePage     uint32 // Code page.
	Reserved     uint32 // Reserved for use by the operating system.
}

type ResourceDataEntry struct {
	ImageResourceDataEntry
	fileOffset int
	size       int
}

func NewResourceDataEntry(fileOffset int) (header *ResourceDataEntry) {
	header = new(ResourceDataEntry)
	header.size = binary.Size(header.ImageResourceDataEntry)
	header.fileOffset = fileOffset
	return header
}

func (r *ResourceDataEntry) String() string {
	return structString(r.fileOffset, "RESOURCE_DATA_ENTRY", r.ImageResourceDataEntry)
}

// VS Version Info
type VersionInfoBlock struct {
	Length      uint16 // Length of this block (doesn't include padding)
	ValueLength uint16 // Value length (if any)
	Type        uint16 // Value type (0 = binary, 1 = text)
	// Key      uint8[1] // Value name (block key) (always NULL terminated)

	//////////
	// WORD padding1[]; // Padding, if any (ALIGNMENT)
	// xxxxx Value[]; // Value data, if any (*ALIGNED*)
	// WORD padding2[]; // Padding, if any (ALIGNMENT)
	// xxxxx Child[]; // Child block(s), if any (*ALIGNED*)
	//////////
}

type VSVersionInfo struct {
	VersionInfoBlock
	fileOffset int
	size       int
}

func NewVSVersionInfo(fileOffset int) (header *VSVersionInfo) {
	header = new(VSVersionInfo)
	header.size = binary.Size(header.VersionInfoBlock)
	header.fileOffset = fileOffset
	return header
}

func (v *VSVersionInfo) String() string {
	return structString(v.fileOffset, "RESOURCE_DATA_ENTRY", v.VersionInfoBlock)
}

// VSFixedFileInfo
type VSFixedfileinfo struct {
	Signature        uint32 // e.g. 0xfeef04bd
	StrucVersion     uint32 // e.g. 0x00000042 = "0.42"
	FileVersionMS    uint32 // e.g. 0x00030075 = "3.75"
	FileVersionLS    uint32 // e.g. 0x00000031 = "0.31"
	ProductVersionMS uint32 // e.g. 0x00030010 = "3.10"
	ProductVersionLS uint32 // e.g. 0x00000031 = "0.31"
	FileFlagsMask    uint32 // = 0x3F for version "0.42"
	FileFlags        uint32 // e.g. VFF_DEBUG | VFF_PRERELEASE
	FileOS           uint32 // e.g. VOS_DOS_WINDOWS16
	FileType         uint32 // e.g. VFT_DRIVER
	FileSubtype      uint32 // e.g. VFT2_DRV_KEYBOARD
	FileDateMS       uint32 // e.g. 0
	FileDateLS       uint32 // e.g. 0
}

type VSFixedFileInfo struct {
	VSFixedfileinfo
	fileOffset int
	size       int
}

func NewVSFixedFileInfo(fileOffset int) (header *VSFixedFileInfo) {
	header = new(VSFixedFileInfo)
	header.size = binary.Size(header.VSFixedfileinfo)
	header.fileOffset = fileOffset
	return header
}

func (v *VSFixedFileInfo) String() string {
	return structString(v.fileOffset, "VSFixedFileInfo", v.VSFixedfileinfo)
}

// StringFileInfo
type StringFileInfoD struct {
	Length      uint16
	ValueLength uint16
	Type        uint16
}

type StringFileInfo struct {
	Data       StringFileInfoD
	fileOffset int
	size       int
}

func NewStringFileInfo(fileOffset int) (header *StringFileInfo) {
	header = new(StringFileInfo)
	header.size = binary.Size(header.Data)
	header.fileOffset = fileOffset
	return header
}

func (s *StringFileInfo) String() string {
	return structString(s.fileOffset, "StringFileInfo", s.Data)
}

// VSStringTable
type VSStringTableD struct {
	Length      uint16
	ValueLength uint16
	Type        uint16
}

type VSStringTable struct {
	Data       VSStringTableD
	fileOffset int
	size       int
}

func NewStringTable(fileOffset int) (header *VSStringTable) {
	header = new(VSStringTable)
	header.size = binary.Size(header.Data)
	header.fileOffset = fileOffset
	return header
}

func (s *VSStringTable) String() string {
	return structString(s.fileOffset, "StringTable", s.Data)
}

// String
type StringD struct {
	Length      uint16
	ValueLength uint16
	Type        uint16
}

type String struct {
	Data       StringD
	fileOffset int
	size       int
}

func NewString(fileOffset int) (header *String) {
	header = new(String)
	header.size = binary.Size(header.Data)
	header.fileOffset = fileOffset
	return header
}

func (s *String) String() string {
	return structString(s.fileOffset, "String", s.Data)
}

// Var
type VarD struct {
	Length      uint16
	ValueLength uint16
	Type        uint16
}

type Var struct {
	Data       VarD
	fileOffset int
	size       int
}

func NewVar(fileOffset int) (header *Var) {
	header = new(Var)
	header.size = binary.Size(header.Data)
	header.fileOffset = fileOffset
	return header
}

func (v *Var) String() string {
	return structString(v.fileOffset, "Var", v.Data)
}

// ThunkData32
type ImageThunkData32 struct {
	AddressOfData uint32
}

type ThunkData32 struct {
	ImageThunkData32
	fileOffset int
	size       int
}

func NewThunkData32(fileOffset int) (header *ThunkData32) {
	header = new(ThunkData32)
	header.size = binary.Size(header.ImageThunkData32)
	header.fileOffset = fileOffset
	return header
}

func (d *ThunkData32) String() string {
	return structString(d.fileOffset, "ThunkData", d.ImageThunkData32)
}

// ThunkData64
type ImageThunkData64 struct {
	AddressOfData uint64
}

type ThunkData64 struct {
	ImageThunkData64
	fileOffset int
	size       int
}

func NewThunkData64(fileOffset int) (header *ThunkData64) {
	header = new(ThunkData64)
	header.size = binary.Size(header.ImageThunkData64)
	header.fileOffset = fileOffset
	return header
}

func (d *ThunkData64) String() string {
	return structString(d.fileOffset, "ThunkData64", d.ImageThunkData64)
}

// DebugDirectory
type ImageDebugDirectory struct {
	Characteristics  uint32
	TimeDateStamp    uint32
	MajorVersion     uint16
	MinorVersion     uint16
	Type             uint32
	SizeOfData       uint32
	AddressOfRawData uint32
	PointerToRawData uint32
}

// CodeView Debug OMF signature. The signature at the end of the file is
// a negative offset from the end of the file to another signature.  At
// the negative offset (base address) is another signature whose filepos
// field points to the first OMFDirHeader in a chain of directories.
// The NB05 signature is used by the link utility to indicated a completely
// unpacked file. The NB06 signature is used by ilink to indicate that the
// executable has had CodeView information from an incremental link appended
// to the executable. The NB08 signature is used by cvpack to indicate that
// the CodeView Debug OMF has been packed. CodeView will only process
// executables with the NB08 signature.
type OMFSignature struct {
	Signature uint32 // "NBxx"
	Filepos   uint32 // offset in file
}

type CvInfoPdb20 struct {
	CvSignature uint32
	Filepos     uint32
	Signature   uint32
	Age         uint32
	// PdbFileName ... Variable sized array
}

type CvInfoPdb70 struct {
	CvSignature uint32
	Signature   [16]byte
	Age         uint32
	// PdbFileName ... Variable sized array
}

type DebugDirectory struct {
	ImageDebugDirectory
	RawData    []byte
	SymbolName []byte
	InfoPdb70  *CvInfoPdb70
	InfoPdb20  *CvInfoPdb20
	fileOffset int
	size       int
	flags      map[string]bool
}

func NewDebugDirectory(fileOffset int) (header *DebugDirectory) {
	header = new(DebugDirectory)
	header.size = binary.Size(header.ImageDebugDirectory)
	header.flags = make(map[string]bool)
	header.fileOffset = fileOffset
	return header
}

func (d *DebugDirectory) String() string {
	return structString(d.fileOffset, "DebugDirectory", d.ImageDebugDirectory) + flagString(d.flags)
}

// TLSDirectory
type ImageTLSDirectory32 struct {
	StartAddressOfRawData uint32
	EndAddressOfRawData   uint32
	AddressOfIndex        uint32
	AddressOfCallBacks    uint32
	SizeOfZeroFill        uint32
	Characteristics       uint32
}

type TLSDirectory32 struct {
	ImageTLSDirectory32
	fileOffset int
	size       int
	flags      map[string]bool
}

func NewTLSDirectory(fileOffset int) (header *TLSDirectory32) {
	header = new(TLSDirectory32)
	header.size = binary.Size(header.ImageTLSDirectory32)
	header.flags = make(map[string]bool)
	header.fileOffset = fileOffset
	return header
}

func (t *TLSDirectory32) String() string {
	return structString(t.fileOffset, "TLSDirectory", t.ImageTLSDirectory32) + flagString(t.flags)
}

// TLSDirectory64
type ImageTLSDirectory64 struct {
	StartAddressOfRawData uint64
	EndAddressOfRawData   uint64
	AddressOfIndex        uint64
	AddressOfCallBacks    uint64
	SizeOfZeroFill        uint32
	Characteristics       uint32
}

type TLSDirectory64 struct {
	ImageTLSDirectory64
	fileOffset int
	size       int
	flags      map[string]bool
}

func NewTLSDirectory64(fileOffset int) (header *TLSDirectory64) {
	header = new(TLSDirectory64)
	header.size = binary.Size(header.ImageTLSDirectory64)
	header.flags = make(map[string]bool)
	header.fileOffset = fileOffset
	return header
}

func (t *TLSDirectory64) String() string {
	return structString(t.fileOffset, "TLSDirectory64", t.ImageTLSDirectory64) + flagString(t.flags)
}

// LoadConfigDirectory
type ImageLoadConfigDirectory32 struct {
	Size                          uint32
	TimeDateStamp                 uint32
	MajorVersion                  uint16
	MinorVersion                  uint16
	GlobalFlagsClear              uint32
	GlobalFlagsSet                uint32
	CriticalSectionDefaultTimeout uint32
	DeCommitFreeBlockThreshold    uint32
	DeCommitTotalFreeThreshold    uint32
	LockPrefixTable               uint32
	MaximumAllocationSize         uint32
	VirtualMemoryThreshold        uint32
	ProcessHeapFlags              uint32
	ProcessAffinityMask           uint32
	CSDVersion                    uint16
	Reserved1                     uint16
	EditList                      uint32
	SecurityCookie                uint32
	SEHandlerTable                uint32
	SEHandlerCount                uint32
	GuardCFCheckFunctionPointer   uint32
	Reserved2                     uint32
	GuardCFFunctionTable          uint32
	GuardCFFunctionCount          uint32
	GuardFlags                    uint32
}

type LoadConfigDirectory32 struct {
	ImageLoadConfigDirectory32
	fileOffset int
	size       int
	flags      map[string]bool
}

func NewLoadConfigDirectory32(fileOffset int) (header *LoadConfigDirectory32) {
	header = new(LoadConfigDirectory32)
	header.size = binary.Size(header.ImageLoadConfigDirectory32)
	header.flags = make(map[string]bool)
	header.fileOffset = fileOffset
	return header
}

func (l *LoadConfigDirectory32) String() string {
	return structString(l.fileOffset, "LoadConfigDirectory", l.ImageLoadConfigDirectory32) + flagString(l.flags)
}

// LoadConfigDirectory64
type ImageLoadConfigDirectory64 struct {
	Size                          uint32
	TimeDateStamp                 uint32
	MajorVersion                  uint16
	MinorVersion                  uint16
	GlobalFlagsClear              uint32
	GlobalFlagsSet                uint32
	CriticalSectionDefaultTimeout uint32
	DeCommitFreeBlockThreshold    uint64
	DeCommitTotalFreeThreshold    uint64
	LockPrefixTable               uint64
	MaximumAllocationSize         uint64
	VirtualMemoryThreshold        uint64
	ProcessAffinityMask           uint64
	ProcessHeapFlags              uint32
	CSDVersion                    uint16
	Reserved1                     uint16
	EditList                      uint64
	SecurityCookie                uint64
	SEHandlerTable                uint64
	SEHandlerCount                uint64
	GuardCFCheckFunctionPointer   uint64
	Reserved2                     uint64
	GuardCFFunctionTable          uint64
	GuardCFFunctionCount          uint64
	GuardFlags                    uint32
}

type LoadConfigDirectory64 struct {
	ImageLoadConfigDirectory64
	fileOffset int
	size       int
	flags      map[string]bool
}

func NewLoadConfigDirectory64(fileOffset int) (header *LoadConfigDirectory64) {
	header = new(LoadConfigDirectory64)
	header.size = binary.Size(header.ImageLoadConfigDirectory64)
	header.flags = make(map[string]bool)
	header.fileOffset = fileOffset
	return header
}

func (l *LoadConfigDirectory64) String() string {
	return structString(l.fileOffset, "LoadConfigDirectory64", l.ImageLoadConfigDirectory64) + flagString(l.flags)
}

// BoundImportDescriptor
type ImageBoundImportDescriptor struct {
	TimeDateStamp               uint32
	OffsetModuleName            uint16
	NumberOfModuleForwarderRefs uint16
}

type BoundImportDescriptor struct {
	ImageBoundImportDescriptor
	fileOffset int
	size       int
}

func NewBoundImportDescriptor(fileOffset int) (header *BoundImportDescriptor) {
	header = new(BoundImportDescriptor)
	header.size = binary.Size(header.ImageBoundImportDescriptor)
	header.fileOffset = fileOffset
	return header
}

func (d *BoundImportDescriptor) String() string {
	return structString(d.fileOffset, "BoundImportDescriptor", d.ImageBoundImportDescriptor)
}

// BoundForwarderRef
type ImageBoundForwarderRef struct {
	TimeDateStamp    uint32
	OffsetModuleName uint16
	Reserved         uint16
}

type BoundForwarderRef struct {
	ImageBoundForwarderRef
	fileOffset int
	size       int
}

func NewBoundForwarderRef(fileOffset int) (header *BoundForwarderRef) {
	header = new(BoundForwarderRef)
	header.size = binary.Size(header.ImageBoundForwarderRef)
	header.fileOffset = fileOffset
	return header
}

func (r *BoundForwarderRef) String() string {
	return structString(r.fileOffset, "BoundForwarderRef", r.ImageBoundForwarderRef)
}

// Helper functions

func structString(fileOffset int, structName string, iface any) string {
	sType := reflect.TypeOf(iface)
	sValue := reflect.ValueOf(iface)
	var values strings.Builder
	values.WriteString("[" + structName + "]\n")
	for i := 0; i < sType.NumField(); i++ {
		sField := sType.Field(i)
		vField := sValue.Field(i)
		kind := vField.Kind()

		fieldOffset := uint64(fileOffset) + uint64(sField.Offset)
		if kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 {
			values.WriteString(fmt.Sprintf("0x%-4X\t\t0x%-4X\t%-24s\t0x%X"+
				"\n", fieldOffset, sField.Offset, sField.Name, vField.Interface()))
		}

		if sValue.Kind() == reflect.Array {
			elemType := sValue.Type().Elem().Kind()
			if elemType == reflect.Struct || elemType == reflect.Pointer || elemType == reflect.Map || elemType == reflect.Func || elemType == reflect.Interface {
				continue
			}

		}
		if kind == reflect.Array || kind == reflect.Slice || kind == reflect.String {
			values.WriteString(fmt.Sprintf("0x%-4X\t\t0x%-4X\t%-24s\t%s"+
				"\n", fieldOffset, sField.Offset, sField.Name, vField.Interface()))
		}

		if kind == reflect.Bool {
			values.WriteString(fmt.Sprintf("0x%-4X\t\t0x%-4X\t%-24s\t%t"+
				"\n", fieldOffset, sField.Offset, sField.Name, vField.Interface()))
		}
	}
	return values.String()
}

func flagString(flagMap map[string]bool) string {
	if len(flagMap) == 0 {
		return "No Flags\n"
	}

	var values strings.Builder
	for key, value := range flagMap {
		if value == true {
			values.WriteString(key + " | ")
		}
	}

	return "Flags: " + strings.TrimSuffix(values.String(), " | ") + "\n"
}

func EmptyStruct(iface any) bool {
	value := reflect.ValueOf(iface)
	for _, field := range value.Fields() {
		kind := field.Kind()
		if kind == reflect.Uint8 && field.Interface().(uint8) != uint8(0) {
			return false
		}
		if kind == reflect.Uint16 && field.Interface().(uint16) != uint16(0) {
			return false
		}
		if kind == reflect.Uint32 && field.Interface().(uint32) != uint32(0) {
			return false
		}
		if kind == reflect.Uint64 && field.Interface().(uint64) != uint64(0) {
			return false
		}
		if kind == reflect.Array && len(field.Interface().([]byte)) != 0 {
			return false
		}
		if kind == reflect.String && len(field.Interface().(string)) != 0 {
			return false
		}
	}
	return true
}

// Call this function after the data has been parsed
func SetFlags(flagMap map[string]bool, charMap map[string]uint32, flags uint32) {
	for key, value := range charMap {
		if (flags & value) == value {
			flagMap[key] = true
		} else {
			flagMap[key] = false
		}
	}
}
