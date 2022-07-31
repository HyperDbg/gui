package tools
const(
ZYAN_MODULE_ZYDIS_PE =    0x101 //col:62
ZYDIS_STATUS_INVALID_DOS_SIGNATURE = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYDIS_PE, 0x00) //col:71
ZYDIS_STATUS_INVALID_NT_SIGNATURE = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYDIS_PE, 0x01) //col:77
ZYDIS_STATUS_UNSUPPORTED_ARCHITECTURE = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYDIS_PE, 0x02) //col:83
IMAGE_DOS_SIGNATURE =                             0x5A4D  // MZ //col:96
IMAGE_NT_SIGNATURE =                          0x00004550  // PE00 //col:97
IMAGE_NT_OPTIONAL_HDR32_MAGIC =                   0x010B //col:98
IMAGE_NT_OPTIONAL_HDR64_MAGIC =                   0x020B //col:99
IMAGE_NUMBEROF_DIRECTORY_ENTRIES =                    16 //col:101
IMAGE_DIRECTORY_ENTRY_EXPORT =                         0 //col:102
IMAGE_DIRECTORY_ENTRY_IMPORT =                         1 //col:103
IMAGE_FILE_MACHINE_I386 =                         0x014C //col:105
IMAGE_FILE_MACHINE_IA64 =                         0x0200 //col:106
IMAGE_FILE_MACHINE_AMD64 =                        0x8664 //col:107
IMAGE_SCN_CNT_CODE =                          0x00000020 //col:109
IMAGE_IMPORT_BY_ORDINAL32 =                   0x80000000 //col:111
IMAGE_IMPORT_BY_ORDINAL64 =           0x8000000000000000 //col:112
IMAGE_SIZEOF_SHORT_NAME = 8 //col:241
IMAGE_FIRST_SECTION(nt_headers) = ((IMAGE_SECTION_HEADER*)((ZyanUPointer)(nt_headers) + offsetof(IMAGE_NT_HEADERS32, OptionalHeader) + ((nt_headers))->FileHeader.SizeOfOptionalHeader)) //col:324
)
type (
ZydisPe interface{
  Zyan Disassembler Library ()(ok bool)//col:139
#pragma pack()(ok bool)//col:310
#pragma pack()(ok bool)//col:318
#define IMAGE_FIRST_SECTION()(ok bool)//col:356
static ZyanI32 CompareSymbol()(ok bool)//col:413
static const IMAGE_SECTION_HEADER* GetSectionByRVA()(ok bool)//col:452
const void* RVAToFileOffset()(ok bool)//col:474
static ZyanStatus ZydisPEContextFinalize()(ok bool)//col:499
static ZyanStatus ZydisPEContextInit()(ok bool)//col:871
static ZyanStatus ZydisFormatterPrintAddress()(ok bool)//col:921
static ZyanStatus ZydisFormatterPrintAddressABS()(ok bool)//col:931
static ZyanStatus ZydisFormatterPrintAddressREL()(ok bool)//col:940
static ZyanStatus DisassembleMappedPEFile()(ok bool)//col:1086
int main()(ok bool)//col:1186
}

)
func NewZydisPe() { return & zydisPe{} }
func (z *zydisPe)  Zyan Disassembler Library ()(ok bool){//col:139
return true
}

func (z *zydisPe)#pragma pack()(ok bool){//col:310
return true
}

func (z *zydisPe)#pragma pack()(ok bool){//col:318
return true
}

func (z *zydisPe)#define IMAGE_FIRST_SECTION()(ok bool){//col:356
return true
}

func (z *zydisPe)static ZyanI32 CompareSymbol()(ok bool){//col:413
return true
}

func (z *zydisPe)static const IMAGE_SECTION_HEADER* GetSectionByRVA()(ok bool){//col:452
return true
}

func (z *zydisPe)const void* RVAToFileOffset()(ok bool){//col:474
return true
}

func (z *zydisPe)static ZyanStatus ZydisPEContextFinalize()(ok bool){//col:499
return true
}

func (z *zydisPe)static ZyanStatus ZydisPEContextInit()(ok bool){//col:871
return true
}

func (z *zydisPe)static ZyanStatus ZydisFormatterPrintAddress()(ok bool){//col:921
return true
}

func (z *zydisPe)static ZyanStatus ZydisFormatterPrintAddressABS()(ok bool){//col:931
return true
}

func (z *zydisPe)static ZyanStatus ZydisFormatterPrintAddressREL()(ok bool){//col:940
return true
}

func (z *zydisPe)static ZyanStatus DisassembleMappedPEFile()(ok bool){//col:1086
return true
}

func (z *zydisPe)int main()(ok bool){//col:1186
return true
}

