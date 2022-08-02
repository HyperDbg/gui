package tools


type typedef struct IMAGE_DOS_HEADER_ struct{
e_magic ZyanU16 //col:3
e_cblp ZyanU16 //col:4
e_cp ZyanU16 //col:5
e_crlc ZyanU16 //col:6
e_cparhdr ZyanU16 //col:7
e_minalloc ZyanU16 //col:8
e_maxalloc ZyanU16 //col:9
e_ss ZyanU16 //col:10
e_sp ZyanU16 //col:11
e_csum ZyanU16 //col:12
e_ip ZyanU16 //col:13
e_cs ZyanU16 //col:14
e_lfarlc ZyanU16 //col:15
e_ovno ZyanU16 //col:16
e_res[4] ZyanU16 //col:17
e_oemid ZyanU16 //col:18
e_oeminfo ZyanU16 //col:19
e_res2[10] ZyanU16 //col:20
e_lfanew ZyanU32 //col:21
}



type typedef struct IMAGE_FILE_HEADER_ struct{
Machine ZyanU16 //col:25
NumberOfSections ZyanU16 //col:26
TimeDateStamp ZyanU32 //col:27
PointerToSymbolTable ZyanU32 //col:28
NumberOfSymbols ZyanU32 //col:29
SizeOfOptionalHeader ZyanU16 //col:30
Characteristics ZyanU16 //col:31
}



type typedef struct IMAGE_DATA_DIRECTORY_ struct{
VirtualAddress ZyanU32 //col:35
Size ZyanU32 //col:36
}



type typedef struct IMAGE_OPTIONAL_HEADER32_ struct{
Magic ZyanU16 //col:40
MajorLinkerVersion ZyanU8 //col:41
MinorLinkerVersion ZyanU8 //col:42
SizeOfCode ZyanU32 //col:43
SizeOfInitializedData ZyanU32 //col:44
SizeOfUninitializedData ZyanU32 //col:45
AddressOfEntryPoint ZyanU32 //col:46
BaseOfCode ZyanU32 //col:47
BaseOfData ZyanU32 //col:48
ImageBase ZyanU32 //col:49
SectionAlignment ZyanU32 //col:50
FileAlignment ZyanU32 //col:51
MajorOperatingSystemVersion ZyanU16 //col:52
MinorOperatingSystemVersion ZyanU16 //col:53
MajorImageVersion ZyanU16 //col:54
MinorImageVersion ZyanU16 //col:55
MajorSubsystemVersion ZyanU16 //col:56
MinorSubsystemVersion ZyanU16 //col:57
Win32VersionValue ZyanU32 //col:58
SizeOfImage ZyanU32 //col:59
SizeOfHeaders ZyanU32 //col:60
CheckSum ZyanU32 //col:61
Subsystem ZyanU16 //col:62
DllCharacteristics ZyanU16 //col:63
SizeOfStackReserve ZyanU32 //col:64
SizeOfStackCommit ZyanU32 //col:65
SizeOfHeapReserve ZyanU32 //col:66
SizeOfHeapCommit ZyanU32 //col:67
LoaderFlags ZyanU32 //col:68
NumberOfRvaAndSizes ZyanU32 //col:69
DataDirectory[IMAGE_NUMBEROF_DIRECTORY_ENTRIES] IMAGE_DATA_DIRECTORY //col:70
}



type typedef struct IMAGE_NT_HEADERS32_ struct{
Signature ZyanU32 //col:74
FileHeader IMAGE_FILE_HEADER //col:75
OptionalHeader IMAGE_OPTIONAL_HEADER32 //col:76
}



type typedef struct IMAGE_OPTIONAL_HEADER64_ struct{
Magic ZyanU16 //col:80
MajorLinkerVersion ZyanU8 //col:81
MinorLinkerVersion ZyanU8 //col:82
SizeOfCode ZyanU32 //col:83
SizeOfInitializedData ZyanU32 //col:84
SizeOfUninitializedData ZyanU32 //col:85
AddressOfEntryPoint ZyanU32 //col:86
BaseOfCode ZyanU32 //col:87
ImageBase ZyanU64 //col:88
SectionAlignment ZyanU32 //col:89
FileAlignment ZyanU32 //col:90
MajorOperatingSystemVersion ZyanU16 //col:91
MinorOperatingSystemVersion ZyanU16 //col:92
MajorImageVersion ZyanU16 //col:93
MinorImageVersion ZyanU16 //col:94
MajorSubsystemVersion ZyanU16 //col:95
MinorSubsystemVersion ZyanU16 //col:96
Win32VersionValue ZyanU32 //col:97
SizeOfImage ZyanU32 //col:98
SizeOfHeaders ZyanU32 //col:99
CheckSum ZyanU32 //col:100
Subsystem ZyanU16 //col:101
DllCharacteristics ZyanU16 //col:102
SizeOfStackReserve ZyanU64 //col:103
SizeOfStackCommit ZyanU64 //col:104
SizeOfHeapReserve ZyanU64 //col:105
SizeOfHeapCommit ZyanU64 //col:106
LoaderFlags ZyanU32 //col:107
NumberOfRvaAndSizes ZyanU32 //col:108
DataDirectory[IMAGE_NUMBEROF_DIRECTORY_ENTRIES] IMAGE_DATA_DIRECTORY //col:109
}



type typedef struct IMAGE_NT_HEADERS64_ struct{
Signature ZyanU32 //col:113
FileHeader IMAGE_FILE_HEADER //col:114
OptionalHeader IMAGE_OPTIONAL_HEADER64 //col:115
}



type typedef struct IMAGE_SECTION_HEADER_ { struct{
Name[IMAGE_SIZEOF_SHORT_NAME] ZyanU8 //col:118
Union union //col:119
PhysicalAddress ZyanU32 //col:121
VirtualSize ZyanU32 //col:122
}



type typedef struct IMAGE_EXPORT_DIRECTORY_ struct{
Characteristics ZyanU32 //col:135
TimeDateStamp ZyanU32 //col:136
MajorVersion ZyanU16 //col:137
MinorVersion ZyanU16 //col:138
Name ZyanU32 //col:139
Base ZyanU32 //col:140
NumberOfFunctions ZyanU32 //col:141
NumberOfNames ZyanU32 //col:142
AddressOfFunctions ZyanU32 //col:143
AddressOfNames ZyanU32 //col:144
AddressOfNameOrdinals ZyanU32 //col:145
}



type typedef struct IMAGE_IMPORT_DESCRIPTOR_ struct{
Union union //col:149
Characteristics ZyanU32 //col:151
OriginalFirstThunk ZyanU32 //col:152
}



type typedef struct IMAGE_THUNK_DATA32_ struct{
Union union //col:161
ForwarderString ZyanU32 //col:163
Function ZyanU32 //col:164
Ordinal ZyanU32 //col:165
AddressOfData ZyanU32 //col:166
}



type typedef struct IMAGE_THUNK_DATA64_ struct{
Union union //col:171
ForwarderString ZyanU64 //col:173
Function ZyanU64 //col:174
Ordinal ZyanU64 //col:175
AddressOfData ZyanU64 //col:176
}



type typedef struct IMAGE_IMPORT_BY_NAME_ struct{
Hint ZyanU16 //col:181
Name[1] int8 //col:182
}



type typedef struct ZydisPESymbol_ struct{
address ZyanU64 //col:186
module_name ZyanString //col:187
symbol_name ZyanString //col:188
}



type typedef struct ZydisPEContext_ struct{
void* bool //col:192
size ZyanUSize //col:193
symbols ZyanVector //col:194
image_base ZyanU64 //col:195
unique_strings ZyanVector //col:196
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
















return true
}


func (z *zydisPe)static_const_IMAGE_SECTION_HEADERPtr_GetSectionByRVA()(ok bool){//col:38






















return true
}




