package tools
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\tools\ZydisPE.c.back

type  IMAGE_DOS_HEADER_ struct{
e_magic ZyanU16 //col:24
e_cblp ZyanU16 //col:25
e_cp ZyanU16 //col:26
e_crlc ZyanU16 //col:27
e_cparhdr ZyanU16 //col:28
e_minalloc ZyanU16 //col:29
e_maxalloc ZyanU16 //col:30
e_ss ZyanU16 //col:31
e_sp ZyanU16 //col:32
e_csum ZyanU16 //col:33
e_ip ZyanU16 //col:34
e_cs ZyanU16 //col:35
e_lfarlc ZyanU16 //col:36
e_ovno ZyanU16 //col:37
e_res[4] ZyanU16 //col:38
e_oemid ZyanU16 //col:39
e_oeminfo ZyanU16 //col:40
e_res2[10] ZyanU16 //col:41
e_lfanew ZyanU32 //col:42
}


type  IMAGE_FILE_HEADER_ struct{
Machine ZyanU16 //col:34
NumberOfSections ZyanU16 //col:35
TimeDateStamp ZyanU32 //col:36
PointerToSymbolTable ZyanU32 //col:37
NumberOfSymbols ZyanU32 //col:38
SizeOfOptionalHeader ZyanU16 //col:39
Characteristics ZyanU16 //col:40
}


type  IMAGE_DATA_DIRECTORY_ struct{
VirtualAddress ZyanU32 //col:39
Size ZyanU32 //col:40
}


type  IMAGE_OPTIONAL_HEADER32_ struct{
Magic ZyanU16 //col:73
MajorLinkerVersion ZyanU8 //col:74
MinorLinkerVersion ZyanU8 //col:75
SizeOfCode ZyanU32 //col:76
SizeOfInitializedData ZyanU32 //col:77
SizeOfUninitializedData ZyanU32 //col:78
AddressOfEntryPoint ZyanU32 //col:79
BaseOfCode ZyanU32 //col:80
BaseOfData ZyanU32 //col:81
ImageBase ZyanU32 //col:82
SectionAlignment ZyanU32 //col:83
FileAlignment ZyanU32 //col:84
MajorOperatingSystemVersion ZyanU16 //col:85
MinorOperatingSystemVersion ZyanU16 //col:86
MajorImageVersion ZyanU16 //col:87
MinorImageVersion ZyanU16 //col:88
MajorSubsystemVersion ZyanU16 //col:89
MinorSubsystemVersion ZyanU16 //col:90
Win32VersionValue ZyanU32 //col:91
SizeOfImage ZyanU32 //col:92
SizeOfHeaders ZyanU32 //col:93
CheckSum ZyanU32 //col:94
Subsystem ZyanU16 //col:95
DllCharacteristics ZyanU16 //col:96
SizeOfStackReserve ZyanU32 //col:97
SizeOfStackCommit ZyanU32 //col:98
SizeOfHeapReserve ZyanU32 //col:99
SizeOfHeapCommit ZyanU32 //col:100
LoaderFlags ZyanU32 //col:101
NumberOfRvaAndSizes ZyanU32 //col:102
DataDirectory[IMAGE_NUMBEROF_DIRECTORY_ENTRIES] IMAGE_DATA_DIRECTORY //col:103
}


type  IMAGE_NT_HEADERS32_ struct{
Signature ZyanU32 //col:79
FileHeader IMAGE_FILE_HEADER //col:80
OptionalHeader IMAGE_OPTIONAL_HEADER32 //col:81
}


type  IMAGE_OPTIONAL_HEADER64_ struct{
Magic ZyanU16 //col:112
MajorLinkerVersion ZyanU8 //col:113
MinorLinkerVersion ZyanU8 //col:114
SizeOfCode ZyanU32 //col:115
SizeOfInitializedData ZyanU32 //col:116
SizeOfUninitializedData ZyanU32 //col:117
AddressOfEntryPoint ZyanU32 //col:118
BaseOfCode ZyanU32 //col:119
ImageBase ZyanU64 //col:120
SectionAlignment ZyanU32 //col:121
FileAlignment ZyanU32 //col:122
MajorOperatingSystemVersion ZyanU16 //col:123
MinorOperatingSystemVersion ZyanU16 //col:124
MajorImageVersion ZyanU16 //col:125
MinorImageVersion ZyanU16 //col:126
MajorSubsystemVersion ZyanU16 //col:127
MinorSubsystemVersion ZyanU16 //col:128
Win32VersionValue ZyanU32 //col:129
SizeOfImage ZyanU32 //col:130
SizeOfHeaders ZyanU32 //col:131
CheckSum ZyanU32 //col:132
Subsystem ZyanU16 //col:133
DllCharacteristics ZyanU16 //col:134
SizeOfStackReserve ZyanU64 //col:135
SizeOfStackCommit ZyanU64 //col:136
SizeOfHeapReserve ZyanU64 //col:137
SizeOfHeapCommit ZyanU64 //col:138
LoaderFlags ZyanU32 //col:139
NumberOfRvaAndSizes ZyanU32 //col:140
DataDirectory[IMAGE_NUMBEROF_DIRECTORY_ENTRIES] IMAGE_DATA_DIRECTORY //col:141
}


type  IMAGE_NT_HEADERS64_ struct{
Signature ZyanU32 //col:118
FileHeader IMAGE_FILE_HEADER //col:119
OptionalHeader IMAGE_OPTIONAL_HEADER64 //col:120
}


type  IMAGE_SECTION_HEADER_  struct{
Name[IMAGE_SIZEOF_SHORT_NAME] ZyanU8 //col:124
Union union //col:125
PhysicalAddress ZyanU32 //col:127
VirtualSize ZyanU32 //col:128
}


type  IMAGE_EXPORT_DIRECTORY_ struct{
Characteristics ZyanU32 //col:148
TimeDateStamp ZyanU32 //col:149
MajorVersion ZyanU16 //col:150
MinorVersion ZyanU16 //col:151
Name ZyanU32 //col:152
Base ZyanU32 //col:153
NumberOfFunctions ZyanU32 //col:154
NumberOfNames ZyanU32 //col:155
AddressOfFunctions ZyanU32 //col:156
AddressOfNames ZyanU32 //col:157
AddressOfNameOrdinals ZyanU32 //col:158
}


type  IMAGE_IMPORT_DESCRIPTOR_ struct{
Union union //col:155
Characteristics ZyanU32 //col:157
OriginalFirstThunk ZyanU32 //col:158
}


type  IMAGE_THUNK_DATA32_ struct{
Union union //col:169
ForwarderString ZyanU32 //col:171
Function ZyanU32 //col:172
Ordinal ZyanU32 //col:173
AddressOfData ZyanU32 //col:174
}


type  IMAGE_THUNK_DATA64_ struct{
Union union //col:179
ForwarderString ZyanU64 //col:181
Function ZyanU64 //col:182
Ordinal ZyanU64 //col:183
AddressOfData ZyanU64 //col:184
}


type  IMAGE_IMPORT_BY_NAME_ struct{
Hint ZyanU16 //col:185
Name[1] int8 //col:186
}


type  ZydisPESymbol_ struct{
address ZyanU64 //col:191
module_name ZyanString //col:192
symbol_name ZyanString //col:193
}


type  ZydisPEContext_ struct{
void* bool //col:199
size ZyanUSize //col:200
symbols ZyanVector //col:201
image_base ZyanU64 //col:202
unique_strings ZyanVector //col:203
}



type (
ZydisPe interface{
static_const_ZyanStringView_STR_DOT_=_ZYAN_DEFINE_STRING_VIEW()(ok bool)//col:16
static_const_IMAGE_SECTION_HEADERPtr_GetSectionByRVA()(ok bool)//col:38
}
zydisPe struct{}
)

func NewZydisPe()ZydisPe{ return & zydisPe{} }

func (z *zydisPe)static_const_ZyanStringView_STR_DOT_=_ZYAN_DEFINE_STRING_VIEW()(ok bool){//col:16
/*static const ZyanStringView STR_DOT = ZYAN_DEFINE_STRING_VIEW(".");
static const ZyanStringView STR_ENTRY_POINT = ZYAN_DEFINE_STRING_VIEW("EntryPoint");
static ZyanI32 CompareSymbol(const ZydisPESymbol* left, const ZydisPESymbol* right)
{
    ZYAN_ASSERT(left);
    ZYAN_ASSERT(right);
    if (left->address < right->address)
    {
        return -1;
    }
    if (left->address > right->address)
    {
        return  1;
    }
    return 0;
}*/
return true
}

func (z *zydisPe)static_const_IMAGE_SECTION_HEADERPtr_GetSectionByRVA()(ok bool){//col:38
/*static const IMAGE_SECTION_HEADER* GetSectionByRVA(const void* base, ZyanU64 rva)
{
    ZYAN_ASSERT(base);
    const IMAGE_DOS_HEADER* dos_header = (const IMAGE_DOS_HEADER*)base;
    ZYAN_ASSERT(dos_header->e_magic == IMAGE_DOS_SIGNATURE);
        (const IMAGE_NT_HEADERS32*)((ZyanU8*)dos_header + dos_header->e_lfanew);
    ZYAN_ASSERT(nt_headers->Signature == IMAGE_NT_SIGNATURE);
    const IMAGE_SECTION_HEADER* section = IMAGE_FIRST_SECTION(nt_headers);
    for (ZyanU16 i = 0; i < nt_headers->FileHeader.NumberOfSections; ++i, ++section)
    {
        ZyanU32 size = section->SizeOfRawData;
        if (section->Misc.VirtualSize > 0)
        {
            size = ZYAN_MIN(section->Misc.VirtualSize, size);
        size = ZYAN_ALIGN_UP(size, nt_headers->OptionalHeader.FileAlignment);
        if ((rva >= section->VirtualAddress) && (rva < (section->VirtualAddress + size)))
        {
            return section;
        }
    }
    return ZYAN_NULL;
}*/
return true
}



