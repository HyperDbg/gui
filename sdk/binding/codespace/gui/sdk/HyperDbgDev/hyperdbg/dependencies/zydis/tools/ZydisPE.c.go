package tools

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\tools\ZydisPE.c.back

const (
	ZYAN_MODULE_ZYDIS_PE                               = 0x101                                                                                                                                               //col:1
	ZYDIS_STATUS_INVALID_DOS_SIGNATURE                 = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYDIS_PE, 0x00)                                                                                                     //col:2
	ZYDIS_STATUS_INVALID_NT_SIGNATURE                  = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYDIS_PE, 0x01)                                                                                                     //col:4
	ZYDIS_STATUS_UNSUPPORTED_ARCHITECTURE              = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYDIS_PE, 0x02)                                                                                                     //col:6
	IMAGE_DOS_SIGNATURE                                = 0x5A4D                                                                                                                                              //col:8
	IMAGE_NT_SIGNATURE                                 = 0x00004550                                                                                                                                          //col:9
	IMAGE_NT_OPTIONAL_HDR32_MAGIC                      = 0x010B                                                                                                                                              //col:10
	IMAGE_NT_OPTIONAL_HDR64_MAGIC                      = 0x020B                                                                                                                                              //col:11
	IMAGE_NUMBEROF_DIRECTORY_ENTRIES                   = 16                                                                                                                                                  //col:12
	IMAGE_DIRECTORY_ENTRY_EXPORT                       = 0                                                                                                                                                   //col:13
	IMAGE_DIRECTORY_ENTRY_IMPORT                       = 1                                                                                                                                                   //col:14
	IMAGE_FILE_MACHINE_I386                            = 0x014C                                                                                                                                              //col:15
	IMAGE_FILE_MACHINE_IA64                            = 0x0200                                                                                                                                              //col:16
	IMAGE_FILE_MACHINE_AMD64                           = 0x8664                                                                                                                                              //col:17
	IMAGE_SCN_CNT_CODE                                 = 0x00000020                                                                                                                                          //col:18
	IMAGE_IMPORT_BY_ORDINAL32                          = 0x80000000                                                                                                                                          //col:19
	IMAGE_IMPORT_BY_ORDINAL64                          = 0x8000000000000000                                                                                                                                  //col:20
	IMAGE_SIZEOF_SHORT_NAME                            = 8                                                                                                                                                   //col:21
	IMAGE_FIRST_SECTION                   (nt_headers) = ((IMAGE_SECTION_HEADER *)((ZyanUPointer)(nt_headers)+offsetof(IMAGE_NT_HEADERS32, OptionalHeader)+(nt_headers)- > FileHeader.SizeOfOptionalHeader)) //col:22
)

type typedef struct IMAGE_DOS_HEADER_ struct {
	e_magic    ZyanU16     //col:3
	e_cblp     ZyanU16     //col:4
	e_cp       ZyanU16     //col:5
	e_crlc     ZyanU16     //col:6
	e_cparhdr  ZyanU16     //col:7
	e_minalloc ZyanU16     //col:8
	e_maxalloc ZyanU16     //col:9
	e_ss       ZyanU16     //col:10
	e_sp       ZyanU16     //col:11
	e_csum     ZyanU16     //col:12
	e_ip       ZyanU16     //col:13
	e_cs       ZyanU16     //col:14
	e_lfarlc   ZyanU16     //col:15
	e_ovno     ZyanU16     //col:16
	e_res      [4]ZyanU16  //col:17
	e_oemid    ZyanU16     //col:18
	e_oeminfo  ZyanU16     //col:19
	e_res2     [10]ZyanU16 //col:20
	e_lfanew   ZyanU32     //col:21
}


	type typedef struct IMAGE_FILE_HEADER_ struct{
	Machine ZyanU16          //col:25
	NumberOfSections ZyanU16 //col:26
	TimeDateStamp ZyanU32    //col:27
	PointerToSymbolTable ZyanU32 //col:28
	NumberOfSymbols ZyanU32      //col:29
	SizeOfOptionalHeader ZyanU16 //col:30
	Characteristics ZyanU16 //col:31
	}


	type typedef struct IMAGE_DATA_DIRECTORY_ struct{
	VirtualAddress ZyanU32 //col:35
	Size ZyanU32           //col:36
	}


	type typedef struct IMAGE_OPTIONAL_HEADER32_ struct{
	Magic ZyanU16             //col:40
	MajorLinkerVersion ZyanU8 //col:41
	MinorLinkerVersion ZyanU8 //col:42
	SizeOfCode ZyanU32              //col:43
	SizeOfInitializedData ZyanU32   //col:44
	SizeOfUninitializedData ZyanU32 //col:45
	AddressOfEntryPoint ZyanU32 //col:46
	BaseOfCode ZyanU32          //col:47
	BaseOfData ZyanU32          //col:48
	ImageBase ZyanU32           //col:49
	SectionAlignment ZyanU32            //col:50
	FileAlignment ZyanU32               //col:51
	MajorOperatingSystemVersion ZyanU16 //col:52
	MinorOperatingSystemVersion ZyanU16 //col:53
	MajorImageVersion ZyanU16           //col:54
	MinorImageVersion ZyanU16           //col:55
	MajorSubsystemVersion ZyanU16 //col:56
	MinorSubsystemVersion ZyanU16 //col:57
	Win32VersionValue ZyanU32     //col:58
	SizeOfImage ZyanU32           //col:59
	SizeOfHeaders ZyanU32 //col:60
	CheckSum ZyanU32      //col:61
	Subsystem ZyanU16     //col:62
	DllCharacteristics ZyanU16 //col:63
	SizeOfStackReserve ZyanU32 //col:64
	SizeOfStackCommit ZyanU32  //col:65
	SizeOfHeapReserve ZyanU32   //col:66
	SizeOfHeapCommit ZyanU32    //col:67
	LoaderFlags ZyanU32         //col:68
	NumberOfRvaAndSizes ZyanU32 //col:69
	DataDirectory[IMAGE_NUMBEROF_DIRECTORY_ENTRIES] IMAGE_DATA_DIRECTORY //col:70
	}


	type typedef struct IMAGE_NT_HEADERS32_ struct{
	Signature ZyanU32                      //col:74
	FileHeader IMAGE_FILE_HEADER           //col:75
	OptionalHeader IMAGE_OPTIONAL_HEADER32 //col:76
	}


	type typedef struct IMAGE_OPTIONAL_HEADER64_ struct{
	Magic ZyanU16             //col:80
	MajorLinkerVersion ZyanU8 //col:81
	MinorLinkerVersion ZyanU8 //col:82
	SizeOfCode ZyanU32              //col:83
	SizeOfInitializedData ZyanU32   //col:84
	SizeOfUninitializedData ZyanU32 //col:85
	AddressOfEntryPoint ZyanU32 //col:86
	BaseOfCode ZyanU32          //col:87
	ImageBase ZyanU64           //col:88
	SectionAlignment ZyanU32    //col:89
	FileAlignment ZyanU32               //col:90
	MajorOperatingSystemVersion ZyanU16 //col:91
	MinorOperatingSystemVersion ZyanU16 //col:92
	MajorImageVersion ZyanU16     //col:93
	MinorImageVersion ZyanU16     //col:94
	MajorSubsystemVersion ZyanU16 //col:95
	MinorSubsystemVersion ZyanU16 //col:96
	Win32VersionValue ZyanU32     //col:97
	SizeOfImage ZyanU32           //col:98
	SizeOfHeaders ZyanU32         //col:99
	CheckSum ZyanU32           //col:100
	Subsystem ZyanU16          //col:101
	DllCharacteristics ZyanU16 //col:102
	SizeOfStackReserve ZyanU64 //col:103
	SizeOfStackCommit ZyanU64  //col:104
	SizeOfHeapReserve ZyanU64  //col:105
	SizeOfHeapCommit ZyanU64    //col:106
	LoaderFlags ZyanU32         //col:107
	NumberOfRvaAndSizes ZyanU32 //col:108
	DataDirectory[IMAGE_NUMBEROF_DIRECTORY_ENTRIES] IMAGE_DATA_DIRECTORY //col:109
	}


	type typedef struct IMAGE_NT_HEADERS64_ struct{
	Signature ZyanU32                      //col:113
	FileHeader IMAGE_FILE_HEADER           //col:114
	OptionalHeader IMAGE_OPTIONAL_HEADER64 //col:115
	}


	type typedef struct IMAGE_SECTION_HEADER_ {
	struct{
	Name[IMAGE_SIZEOF_SHORT_NAME] ZyanU8 //col:118
	Union union                          //col:119
	PhysicalAddress ZyanU32 //col:121
	VirtualSize ZyanU32     //col:122
	}


	type typedef struct IMAGE_EXPORT_DIRECTORY_ struct{
	Characteristics ZyanU32 //col:135
	TimeDateStamp ZyanU32   //col:136
	MajorVersion ZyanU16    //col:137
	MinorVersion ZyanU16 //col:138
	Name ZyanU32         //col:139
	Base ZyanU32         //col:140
	NumberOfFunctions ZyanU32  //col:141
	NumberOfNames ZyanU32      //col:142
	AddressOfFunctions ZyanU32 //col:143
	AddressOfNames ZyanU32     //col:144
	AddressOfNameOrdinals ZyanU32 //col:145
	}


	type typedef struct IMAGE_IMPORT_DESCRIPTOR_ struct{
	Union union                //col:149
	Characteristics ZyanU32    //col:151
	OriginalFirstThunk ZyanU32 //col:152
	}


	type typedef struct IMAGE_THUNK_DATA32_ struct{
	Union union             //col:161
	ForwarderString ZyanU32 //col:163
	Function ZyanU32        //col:164
	Ordinal ZyanU32       //col:165
	AddressOfData ZyanU32 //col:166
	}


	type typedef struct IMAGE_THUNK_DATA64_ struct{
	Union union             //col:171
	ForwarderString ZyanU64 //col:173
	Function ZyanU64        //col:174
	Ordinal ZyanU64       //col:175
	AddressOfData ZyanU64 //col:176
	}


	type typedef struct IMAGE_IMPORT_BY_NAME_ struct{
	Hint ZyanU16 //col:181
	Name[1] int8 //col:182
	}


	type typedef struct ZydisPESymbol_ struct{
	address ZyanU64        //col:186
	module_name ZyanString //col:187
	symbol_name ZyanString //col:188
	}


	type typedef struct ZydisPEContext_ struct{
	void* bool         //col:192
	size ZyanUSize     //col:193
	symbols ZyanVector //col:194
	image_base ZyanU64        //col:195
	unique_strings ZyanVector //col:196
	}



	type (
	ZydisPe interface{
	static_const_ZyanStringView_STR_DOT_ = _ZYAN_DEFINE_STRING_VIEW()(ok bool) //col:18
	static_const_IMAGE_SECTION_HEADERPtr_GetSectionByRVA()(ok bool) //col:42
	const_voidPtr_RVAToFileOffset()(ok bool) //col:52
	static_ZyanStatus_ZydisPEContextFinalize()(ok bool) //col:65
	static_ZyanStatus_ZydisPEContextInit()(ok bool)         //col:367
	static_ZyanStatus_ZydisFormatterPrintAddress()(ok bool) //col:397
	static_ZyanStatus_ZydisFormatterPrintAddressABS()(ok bool) //col:405
	static_ZyanStatus_ZydisFormatterPrintAddressREL()(ok bool) //col:412
	static_ZyanStatus_DisassembleMappedPEFile()(ok bool) //col:530
	int_main()(ok bool) //col:610
	}
	zydisPe struct{}
	)

	func NewZydisPe()ZydisPe{ return & zydisPe{} }

	func (z *zydisPe)static_const_ZyanStringView_STR_DOT_ = _ZYAN_DEFINE_STRING_VIEW()(ok bool){//col:18
	/*static const ZyanStringView STR_DOT = ZYAN_DEFINE_STRING_VIEW(".");
	  static const ZyanStringView STR_ENTRY_POINT = ZYAN_DEFINE_STRING_VIEW("EntryPoint");
	  #pragma pack(push, 8)
	  #pragma pack(pop)
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

	func (z *zydisPe)static_const_IMAGE_SECTION_HEADERPtr_GetSectionByRVA()(ok bool){//col:42
	/*static const IMAGE_SECTION_HEADER* GetSectionByRVA(const void* base, ZyanU64 rva)
	  {
	      ZYAN_ASSERT(base);
	      const IMAGE_DOS_HEADER* dos_header = (const IMAGE_DOS_HEADER*)base;
	      ZYAN_ASSERT(dos_header->e_magic == IMAGE_DOS_SIGNATURE);
	      const IMAGE_NT_HEADERS32* nt_headers =
	          (const IMAGE_NT_HEADERS32*)((ZyanU8*)dos_header + dos_header->e_lfanew);
	      ZYAN_ASSERT(nt_headers->Signature == IMAGE_NT_SIGNATURE);
	      const IMAGE_SECTION_HEADER* section = IMAGE_FIRST_SECTION(nt_headers);
	      for (ZyanU16 i = 0; i < nt_headers->FileHeader.NumberOfSections; ++i, ++section)
	      {
	          ZyanU32 size = section->SizeOfRawData;
	          if (section->Misc.VirtualSize > 0)
	          {
	              size = ZYAN_MIN(section->Misc.VirtualSize, size);
	          }
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

	func (z *zydisPe)const_voidPtr_RVAToFileOffset()(ok bool){//col:52
	/*const void* RVAToFileOffset(const void* base, ZyanU64 rva)
	  {
	      ZYAN_ASSERT(base);
	      const IMAGE_SECTION_HEADER* section = GetSectionByRVA(base, rva);
	      if (!section)
	      {
	          return ZYAN_NULL;
	      }
	      return (void*)((ZyanU8*)base + section->PointerToRawData + (rva - section->VirtualAddress));
	  }*/
	return true
	}

	func (z *zydisPe)static_ZyanStatus_ZydisPEContextFinalize()(ok bool){//col:65
	/*static ZyanStatus ZydisPEContextFinalize(ZydisPEContext* context)
	  {
	      ZyanUSize size;
	      ZYAN_CHECK(ZyanVectorGetSize(&context->unique_strings, &size));
	      for (ZyanUSize i = 0; i < size; ++i)
	      {
	          ZyanString* string;
	          ZYAN_CHECK(ZyanVectorGetPointerMutable(&context->unique_strings, i, (void**)&string));
	          ZYAN_CHECK(ZyanStringDestroy(string));
	      }
	      ZYAN_CHECK(ZyanVectorDestroy(&context->symbols));
	      return ZyanVectorDestroy(&context->unique_strings);
	  }*/
	return true
	}

	func (z *zydisPe)static_ZyanStatus_ZydisPEContextInit()(ok bool){//col:367
	/*static ZyanStatus ZydisPEContextInit(ZydisPEContext* context, const void* base, ZyanUSize size)
	  {
	      ZYAN_ASSERT(context);
	      ZYAN_ASSERT(base);
	      ZYAN_ASSERT(size);
	      context->base = base;
	      context->size = size;
	      ZyanStatus status;
	      ZYAN_CHECK(status = ZyanVectorInit(&context->symbols, sizeof(ZydisPESymbol), 256, ZYAN_NULL));
	      if (!ZYAN_SUCCESS(status = ZyanVectorInit(&context->unique_strings, sizeof(ZyanString), 256, ZYAN_NULL)))
	      {
	          ZyanVectorDestroy(&context->symbols);
	          return status;
	      }
	      const IMAGE_DOS_HEADER* dos_header = (const IMAGE_DOS_HEADER*)base;
	      ZYAN_ASSERT(dos_header->e_magic == IMAGE_DOS_SIGNATURE);
	      const IMAGE_NT_HEADERS32* nt_headers_temp =
	          (const IMAGE_NT_HEADERS32*)((ZyanU8*)dos_header + dos_header->e_lfanew);
	      ZYAN_ASSERT(nt_headers_temp->Signature == IMAGE_NT_SIGNATURE);
	      switch (nt_headers_temp->OptionalHeader.Magic)
	      {
	      case IMAGE_NT_OPTIONAL_HDR32_MAGIC:
	      {
	          const IMAGE_NT_HEADERS32* nt_headers =
	              (const IMAGE_NT_HEADERS32*)((ZyanU8*)dos_header + dos_header->e_lfanew);
	          context->image_base = nt_headers->OptionalHeader.ImageBase;
	          ZydisPESymbol element;
	          const ZyanUPointer entry_point = nt_headers->OptionalHeader.AddressOfEntryPoint;
	          element.address = entry_point;
	          if (!ZYAN_SUCCESS((status =
	                  ZyanStringDuplicate(&element.symbol_name, &STR_ENTRY_POINT, 0))) ||
	              !ZYAN_SUCCESS((status =
	                  ZyanVectorPushBack(&context->symbols, &element))))
	          {
	              goto FatalError;
	          }
	          if (nt_headers->OptionalHeader.DataDirectory[IMAGE_DIRECTORY_ENTRY_EXPORT].VirtualAddress)
	          {
	              const IMAGE_EXPORT_DIRECTORY* export_directory =
	                  (const IMAGE_EXPORT_DIRECTORY*)RVAToFileOffset(base,
	                      nt_headers->OptionalHeader.DataDirectory[
	                          IMAGE_DIRECTORY_ENTRY_EXPORT].VirtualAddress);
	              ZyanStringView module_name;
	              if (!ZYAN_SUCCESS(status = ZyanStringViewInsideBuffer(&module_name,
	                      (char*)RVAToFileOffset(base, export_directory->Name))) ||
	                  !ZYAN_SUCCESS(status =
	                      ZyanStringDuplicate(&element.module_name, &module_name, 0)))
	              {
	                  goto FatalError;
	              }
	              ZyanISize index;
	              if (!ZYAN_SUCCESS(status = ZyanStringRPos(&module_name, &STR_DOT, &index)))
	              {
	                  goto FatalError;
	              }
	              if (index >= 0)
	              {
	                  if (!ZYAN_SUCCESS(status = ZyanStringTruncate(&element.module_name, index)) ||
	                      !ZYAN_SUCCESS(status =
	                          ZyanVectorPushBack(&context->unique_strings, &element.module_name)))
	                  {
	                      goto FatalError;
	                  }
	              } else
	              if (!ZYAN_SUCCESS(status =
	                      ZyanVectorPushBack(&context->unique_strings, &element.module_name)))
	              {
	                  goto FatalError;
	              }
	              const ZyanU32* export_addresses =
	                  (const ZyanU32*)RVAToFileOffset(base,
	                      export_directory->AddressOfFunctions);
	              const ZyanU32* export_names =
	                  (const ZyanU32*)RVAToFileOffset(base,
	                      export_directory->AddressOfNames);
	              for (ZyanU32 i = 0; i < export_directory->NumberOfFunctions; ++i)
	              {
	                  element.address = export_addresses[i];
	                  ZyanStringView symbol_name;
	                  if (!ZYAN_SUCCESS((status =
	                          ZyanStringViewInsideBuffer(&symbol_name,
	                              (const char*)RVAToFileOffset(base, export_names[i])))) ||
	                      !ZYAN_SUCCESS((status =
	                          ZyanStringDuplicate(&element.symbol_name, &symbol_name, 0))))
	                  {
	                      goto FatalError;
	                  }
	                  ZyanUSize found_index;
	                  if (!ZYAN_SUCCESS((status =
	                          ZyanVectorBinarySearch(&context->symbols, &element, &found_index,
	                              (ZyanComparison)&CompareSymbol))) ||
	                      !ZYAN_SUCCESS((status =
	                          ZyanVectorInsert(&context->symbols, found_index, &element))))
	                  {
	                      goto FatalError;
	                  }
	              }
	          }
	          if (nt_headers->OptionalHeader.DataDirectory[IMAGE_DIRECTORY_ENTRY_IMPORT].VirtualAddress)
	          {
	              const IMAGE_IMPORT_DESCRIPTOR* descriptor =
	                  (const IMAGE_IMPORT_DESCRIPTOR*)RVAToFileOffset(base,
	                      nt_headers->OptionalHeader.DataDirectory[
	                          IMAGE_DIRECTORY_ENTRY_IMPORT].VirtualAddress);
	              while (descriptor->u1.OriginalFirstThunk)
	              {
	                  ZyanStringView module_name;
	                  if (!ZYAN_SUCCESS(status = ZyanStringViewInsideBuffer(&module_name,
	                          (const char*)RVAToFileOffset(base, descriptor->Name))) ||
	                      !ZYAN_SUCCESS(status = ZyanStringDuplicate(
	                          &element.module_name, &module_name, 0)))
	                  {
	                      goto FatalError;
	                  }
	                  ZyanISize index;
	                  if (!ZYAN_SUCCESS(status = ZyanStringRPos(&module_name, &STR_DOT, &index)))
	                  {
	                      goto FatalError;
	                  }
	                  if (index >= 0)
	                  {
	                      if (!ZYAN_SUCCESS(status = ZyanStringTruncate(&element.module_name, index)) ||
	                          !ZYAN_SUCCESS(status =
	                              ZyanVectorPushBack(&context->unique_strings, &element.module_name)))
	                      {
	                          goto FatalError;
	                      }
	                  } else
	                  if (!ZYAN_SUCCESS(status =
	                          ZyanVectorPushBack(&context->unique_strings, &element.module_name)))
	                  {
	                      goto FatalError;
	                  }
	                  const IMAGE_THUNK_DATA32* original_thunk =
	                      (const IMAGE_THUNK_DATA32*)RVAToFileOffset(base,
	                          descriptor->u1.OriginalFirstThunk);
	                  element.address  = descriptor->FirstThunk;
	                  ZyanUSize found_index;
	                  if (!ZYAN_SUCCESS((status =
	                          ZyanVectorBinarySearch(&context->symbols, &element, &found_index,
	                              (ZyanComparison)&CompareSymbol))))
	                  {
	                      goto FatalError;
	                  }
	                  ZYAN_ASSERT(status == ZYAN_STATUS_FALSE);
	                  while (original_thunk->u1.ForwarderString)
	                  {
	                      if (!(original_thunk->u1.Ordinal & IMAGE_IMPORT_BY_ORDINAL32))
	                      {
	                          const IMAGE_IMPORT_BY_NAME* by_name =
	                              (const IMAGE_IMPORT_BY_NAME*)RVAToFileOffset(base,
	                                  original_thunk->u1.AddressOfData);
	                          ZyanStringView symbol_name;
	                          if (!ZYAN_SUCCESS((status =
	                                  ZyanStringViewInsideBuffer(
	                                      &symbol_name, (const char*)by_name->Name))) ||
	                              !ZYAN_SUCCESS((status =
	                                  ZyanStringDuplicate(&element.symbol_name, &symbol_name, 0))))
	                          {
	                              goto FatalError;
	                          }
	                      }
	                      if (!ZYAN_SUCCESS((status =
	                              ZyanVectorInsert(&context->symbols, found_index, &element))))
	                      {
	                          goto FatalError;
	                      }
	                      element.address += sizeof(IMAGE_THUNK_DATA32);
	                      ++original_thunk;
	                      ++found_index;
	                  }
	                  ++descriptor;
	              }
	          }
	          break;
	      }
	      case IMAGE_NT_OPTIONAL_HDR64_MAGIC:
	      {
	          const IMAGE_NT_HEADERS64* nt_headers =
	              (const IMAGE_NT_HEADERS64*)((ZyanU8*)dos_header + dos_header->e_lfanew);
	          context->image_base = nt_headers->OptionalHeader.ImageBase;
	          ZydisPESymbol element;
	          const ZyanUPointer entry_point = nt_headers->OptionalHeader.AddressOfEntryPoint;
	          element.address = entry_point;
	          if (!ZYAN_SUCCESS((status =
	                  ZyanStringDuplicate(&element.symbol_name, &STR_ENTRY_POINT, 0))) ||
	              !ZYAN_SUCCESS((status =
	                  ZyanVectorPushBack(&context->symbols, &element))))
	          {
	              goto FatalError;
	          }
	          if (nt_headers->OptionalHeader.DataDirectory[IMAGE_DIRECTORY_ENTRY_EXPORT].VirtualAddress)
	          {
	              const IMAGE_EXPORT_DIRECTORY* export_directory =
	                  (const IMAGE_EXPORT_DIRECTORY*)RVAToFileOffset(base,
	                      nt_headers->OptionalHeader.DataDirectory[
	                          IMAGE_DIRECTORY_ENTRY_EXPORT].VirtualAddress);
	              ZyanStringView module_name;
	              if (!ZYAN_SUCCESS((status =
	                      ZyanStringViewInsideBuffer(&module_name,
	                          (const char*)RVAToFileOffset(base, export_directory->Name)))) ||
	                  !ZYAN_SUCCESS((status =
	                      ZyanStringDuplicate(&element.module_name, &module_name, 0))))
	              {
	                  goto FatalError;
	              }
	              const ZyanU32* export_addresses =
	                  (const ZyanU32*)RVAToFileOffset(base,
	                      export_directory->AddressOfFunctions);
	              const ZyanU32* export_names =
	                  (const ZyanU32*)RVAToFileOffset(base,
	                      export_directory->AddressOfNames);
	              for (ZyanU32 i = 0; i < export_directory->NumberOfFunctions; ++i)
	              {
	                  element.address = export_addresses[i];
	                  ZyanStringView symbol_name;
	                  if (!ZYAN_SUCCESS((status =
	                      ZyanStringViewInsideBuffer(&symbol_name,
	                          (const char*)RVAToFileOffset(base, export_names[i])))) ||
	                      !ZYAN_SUCCESS((status =
	                          ZyanStringDuplicate(&element.symbol_name, &symbol_name, 0))))
	                  {
	                      goto FatalError;
	                  }
	                  ZyanUSize found_index;
	                  if (!ZYAN_SUCCESS((status =
	                          ZyanVectorBinarySearch(&context->symbols, &element, &found_index,
	                              (ZyanComparison)&CompareSymbol))) ||
	                      !ZYAN_SUCCESS((status =
	                          ZyanVectorInsert(&context->symbols, found_index, &element))))
	                  {
	                      goto FatalError;
	                  }
	              }
	          }
	          if (nt_headers->OptionalHeader.DataDirectory[IMAGE_DIRECTORY_ENTRY_IMPORT].VirtualAddress)
	          {
	              const IMAGE_IMPORT_DESCRIPTOR* descriptor =
	                  (const IMAGE_IMPORT_DESCRIPTOR*)RVAToFileOffset(base,
	                      nt_headers->OptionalHeader.DataDirectory[
	                          IMAGE_DIRECTORY_ENTRY_IMPORT].VirtualAddress);
	              while (descriptor->u1.OriginalFirstThunk)
	              {
	                  ZyanStringView module_name;
	                  if (!ZYAN_SUCCESS((status =
	                      ZyanStringViewInsideBuffer(&module_name,
	                          (const char*)RVAToFileOffset(base, descriptor->Name)))) ||
	                      !ZYAN_SUCCESS((status =
	                          ZyanStringDuplicate(&element.module_name, &module_name, 0))))
	                  {
	                      goto FatalError;
	                  }
	                  const IMAGE_THUNK_DATA64* original_thunk =
	                      (const IMAGE_THUNK_DATA64*)RVAToFileOffset(base,
	                          descriptor->u1.OriginalFirstThunk);
	                  element.address  = descriptor->FirstThunk;
	                  ZyanUSize found_index;
	                  if (!ZYAN_SUCCESS((status =
	                          ZyanVectorBinarySearch(&context->symbols, &element, &found_index,
	                              (ZyanComparison)&CompareSymbol))))
	                  {
	                      goto FatalError;
	                  }
	                  ZYAN_ASSERT(status == ZYAN_STATUS_FALSE);
	                  while (original_thunk->u1.ForwarderString)
	                  {
	                      if (!(original_thunk->u1.Ordinal & IMAGE_IMPORT_BY_ORDINAL64))
	                      {
	                          const IMAGE_IMPORT_BY_NAME* by_name =
	                              (const IMAGE_IMPORT_BY_NAME*)RVAToFileOffset(base,
	                                  original_thunk->u1.AddressOfData);
	                          ZyanStringView symbol_name;
	                          if (!ZYAN_SUCCESS((status = ZyanStringViewInsideBuffer(
	                                  &symbol_name, (const char*)by_name->Name))) ||
	                              !ZYAN_SUCCESS((status =
	                                  ZyanStringDuplicate(&element.symbol_name, &symbol_name, 0))))
	                          {
	                              goto FatalError;
	                          }
	                      }
	                      if (!ZYAN_SUCCESS((status =
	                              ZyanVectorInsert(&context->symbols, found_index, &element))))
	                      {
	                          goto FatalError;
	                      }
	                      element.address += sizeof(IMAGE_THUNK_DATA64);
	                      ++original_thunk;
	                      ++found_index;
	                  }
	                  ++descriptor;
	              }
	          }
	          break;
	      }
	      default:
	          ZYAN_UNREACHABLE;
	      }
	      return ZYAN_STATUS_SUCCESS;
	  FatalError:
	      ZydisPEContextFinalize(context);
	      return status;
	  }*/
	return true
	}

	func (z *zydisPe)static_ZyanStatus_ZydisFormatterPrintAddress()(ok bool){//col:397
	/*static ZyanStatus ZydisFormatterPrintAddress(const ZydisFormatter* formatter,
	      ZydisFormatterBuffer* buffer, ZydisFormatterContext* context, ZyanU64 address, ZyanBool is_abs)
	  {
	      const ZydisPEContext* data = (const ZydisPEContext*)context->user_data;
	      ZYAN_ASSERT(data);
	      ZydisPESymbol symbol;
	      symbol.address = address - data->image_base;
	      ZyanStatus status;
	      ZyanUSize found_index;
	      ZYAN_CHECK((status =
	          ZyanVectorBinarySearch(&data->symbols, &symbol, &found_index,
	              (ZyanComparison)&CompareSymbol)));
	      ZyanString* string;
	      ZYAN_CHECK(ZydisFormatterBufferGetString(buffer, &string));
	      if (status == ZYAN_STATUS_TRUE)
	      {
	          const ZydisPESymbol* element;
	          ZYAN_CHECK(ZyanVectorGetPointer(&data->symbols, found_index, (const void**)&element));
	          ZyanUSize index;
	          ZyanUSize count;
	          ZYAN_CHECK(ZyanStringGetSize(string, &index));
	          ZYAN_CHECK(ZyanStringGetSize(&element->module_name, &count));
	          ZYAN_CHECK(ZyanStringAppend(string, ZYAN_STRING_TO_VIEW(&element->module_name)));
	          ZYAN_CHECK(ZyanStringToLowerCaseEx(string, index, count));
	          ZYAN_CHECK(ZyanStringAppend(string, &STR_DOT));
	          return ZyanStringAppend(string, ZYAN_STRING_TO_VIEW(&element->symbol_name));
	      }
	      ZydisFormatterFunc fn = is_abs ? default_print_address_abs : default_print_address_rel;
	      return fn(formatter, buffer, context);
	  }*/
	return true
	}

	func (z *zydisPe)static_ZyanStatus_ZydisFormatterPrintAddressABS()(ok bool){//col:405
	/*static ZyanStatus ZydisFormatterPrintAddressABS(const ZydisFormatter* formatter,
	      ZydisFormatterBuffer* buffer, ZydisFormatterContext* context)
	  {
	      ZyanU64 address;
	      ZYAN_CHECK(ZydisCalcAbsoluteAddress(context->instruction, context->operand,
	          context->runtime_address, &address));
	      return ZydisFormatterPrintAddress(formatter, buffer, context, address, ZYAN_TRUE);
	  }*/
	return true
	}

	func (z *zydisPe)static_ZyanStatus_ZydisFormatterPrintAddressREL()(ok bool){//col:412
	/*static ZyanStatus ZydisFormatterPrintAddressREL(const ZydisFormatter* formatter,
	      ZydisFormatterBuffer* buffer, ZydisFormatterContext* context)
	  {
	      ZyanU64 address;
	      ZYAN_CHECK(ZydisCalcAbsoluteAddress(context->instruction, context->operand, 0, &address));
	      return ZydisFormatterPrintAddress(formatter, buffer, context, address, ZYAN_FALSE);
	  }*/
	return true
	}

	func (z *zydisPe)static_ZyanStatus_DisassembleMappedPEFile()(ok bool){//col:530
	/*static ZyanStatus DisassembleMappedPEFile(const ZydisPEContext* context)
	  {
	      ZYAN_ASSERT(context);
	      const IMAGE_DOS_HEADER* dos_header = (const IMAGE_DOS_HEADER*)context->base;
	      ZYAN_ASSERT(dos_header->e_magic == IMAGE_DOS_SIGNATURE);
	      const IMAGE_NT_HEADERS32* nt_headers =
	          (const IMAGE_NT_HEADERS32*)((ZyanU8*)dos_header + dos_header->e_lfanew);
	      ZYAN_ASSERT(nt_headers->Signature == IMAGE_NT_SIGNATURE);
	      ZyanStatus status;
	      ZydisMachineMode machine_mode;
	      ZydisAddressWidth address_width;
	      switch (nt_headers->FileHeader.Machine)
	      {
	      case IMAGE_FILE_MACHINE_I386:
	          machine_mode  = ZYDIS_MACHINE_MODE_LONG_COMPAT_32;
	          address_width = ZYDIS_ADDRESS_WIDTH_32;
	          break;
	      case IMAGE_FILE_MACHINE_IA64:
	      case IMAGE_FILE_MACHINE_AMD64:
	          machine_mode  = ZYDIS_MACHINE_MODE_LONG_64;
	          address_width = ZYDIS_ADDRESS_WIDTH_64;
	          break;
	      default:
	          ZYAN_UNREACHABLE;
	      }
	      ZydisDecoder decoder;
	      if (!ZYAN_SUCCESS((status = ZydisDecoderInit(&decoder, machine_mode, address_width))))
	      {
	          fputs("Failed to initialize instruction-decoder\n", stderr);
	          return status;
	      }
	      default_print_address_abs = (ZydisFormatterFunc)&ZydisFormatterPrintAddressABS;
	      default_print_address_rel = (ZydisFormatterFunc)&ZydisFormatterPrintAddressREL;
	      ZydisFormatter formatter;
	      if (!ZYAN_SUCCESS((status = ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_INTEL))) ||
	          !ZYAN_SUCCESS((status = ZydisFormatterSetProperty(&formatter,
	              ZYDIS_FORMATTER_PROP_FORCE_SEGMENT, ZYAN_TRUE))) ||
	          !ZYAN_SUCCESS((status = ZydisFormatterSetProperty(&formatter,
	              ZYDIS_FORMATTER_PROP_FORCE_SIZE, ZYAN_TRUE))) ||
	          !ZYAN_SUCCESS((status = ZydisFormatterSetHook(&formatter,
	              ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS, (const void**)&default_print_address_abs))) ||
	          !ZYAN_SUCCESS((status = ZydisFormatterSetHook(&formatter,
	              ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_REL, (const void**)&default_print_address_rel))))
	      {
	          fputs("Failed to initialize instruction-formatter\n", stderr);
	          return status;
	      }
	      ZydisDecodedInstruction instruction;
	      const IMAGE_SECTION_HEADER* section_header = IMAGE_FIRST_SECTION(nt_headers);
	      for (ZyanU16 i = 0; i < nt_headers->FileHeader.NumberOfSections; ++i)
	      {
	          if (!(section_header->Characteristics & IMAGE_SCN_CNT_CODE))
	          {
	              continue;
	          }
	          const ZyanU8* buffer = (ZyanU8*)context->base + section_header->PointerToRawData;
	          const ZyanUSize buffer_size = section_header->SizeOfRawData;
	          const ZyanU64 read_offset_base = context->image_base + section_header->VirtualAddress;
	          ZyanUSize read_offset = 0;
	          while ((status = ZydisDecoderDecodeBuffer(&decoder, buffer + read_offset,
	              buffer_size - read_offset, &instruction)) != ZYDIS_STATUS_NO_MORE_DATA)
	          {
	              const ZyanU64 runtime_address = read_offset_base + read_offset;
	              ZydisPESymbol symbol;
	              symbol.address = runtime_address - context->image_base;
	              ZyanUSize found_index;
	              ZYAN_CHECK((status =
	                  ZyanVectorBinarySearch(&context->symbols, &symbol, &found_index,
	                      (ZyanComparison)&CompareSymbol)));
	              if (status == ZYAN_STATUS_TRUE)
	              {
	                  const ZydisPESymbol* element;
	                  ZYAN_CHECK(ZyanVectorGetPointer(&context->symbols, found_index,
	                      (const void**)&element));
	                  const char* string;
	                  ZYAN_CHECK(ZyanStringGetData(&element->symbol_name, &string));
	                  printf("\n%s:\n", string);
	              }
	              switch (instruction.machine_mode)
	              {
	              case ZYDIS_MACHINE_MODE_LONG_COMPAT_32:
	                  printf("%08" PRIX32 "  ", (ZyanU32)runtime_address);
	                  break;
	              case ZYDIS_MACHINE_MODE_LONG_64:
	                  printf("%016" PRIX64 "  ", (ZyanU64)runtime_address);
	                  break;
	              default:
	                  ZYAN_UNREACHABLE;
	              }
	              for (int j = 0; j < instruction.length; ++j)
	              {
	                  printf("%02X ", buffer[read_offset + j]);
	              }
	              for (int j = instruction.length; j < 15; ++j)
	              {
	                  printf("   ");
	              }
	              if (ZYAN_SUCCESS(status))
	              {
	                  read_offset += instruction.length;
	                  char format_buffer[256];
	                  if (!ZYAN_SUCCESS((status =
	                          ZydisFormatterFormatInstructionEx(&formatter, &instruction, format_buffer,
	                              sizeof(format_buffer), runtime_address, (void*)context))))
	                  {
	                      fputs("Failed to format instruction\n", stderr);
	                      return status;
	                  }
	                  printf(" %s\n", &format_buffer[0]);
	              } else
	              {
	                  printf(" db %02x\n", buffer[read_offset++]);
	              }
	          }
	          ++section_header;
	      }
	      return ZYAN_STATUS_SUCCESS;
	  }*/
	return true
	}

	func (z *zydisPe)int_main()(ok bool){//col:610
	/*int main(int argc, char** argv)
	  {
	      if (argc != 2)
	      {
	          fprintf(stderr, "Usage: %s <input file>\n", (argc > 0 ? argv[0] : "ZydisPE"));
	          return ZYAN_STATUS_INVALID_ARGUMENT;
	      }
	      FILE* file = fopen(argv[1], "rb");
	      if (!file)
	      {
	          fprintf(stderr, "Could not open file \"%s\": %s\n", argv[1], strerror(errno));
	          return ZYAN_STATUS_BAD_SYSTEMCALL;
	      }
	      fseek(file, 0L, SEEK_END);
	      const long size = ftell(file);
	      void* buffer = ZYAN_MALLOC(size);
	      if (!buffer)
	      {
	          fprintf(stderr, "Failed to allocate %" PRIu64 " bytes on the heap\n", (ZyanU64)size);
	          fclose(file);
	          return ZYAN_STATUS_NOT_ENOUGH_MEMORY;
	      }
	      rewind(file);
	      if (fread(buffer, 1, size, file) != (ZyanUSize)size)
	      {
	          fprintf(stderr,
	              "Could not read %" PRIu64 " bytes from file \"%s\"\n", (ZyanU64)size, argv[1]);
	          ZYAN_FREE(buffer);
	          fclose(file);
	          return ZYAN_STATUS_BAD_SYSTEMCALL;
	      }
	      const IMAGE_DOS_HEADER* dos_header = (const IMAGE_DOS_HEADER*)buffer;
	      if (dos_header->e_magic != IMAGE_DOS_SIGNATURE)
	      {
	          fputs("Invalid file signature (DOS header)\n", stderr);
	          return ZYDIS_STATUS_INVALID_DOS_SIGNATURE;
	      }
	      const IMAGE_NT_HEADERS32* nt_headers =
	          (const IMAGE_NT_HEADERS32*)((ZyanU8*)dos_header + dos_header->e_lfanew);
	      if (nt_headers->Signature != IMAGE_NT_SIGNATURE)
	      {
	          fputs("Invalid file signature (NT headers)\n", stderr);
	          return ZYDIS_STATUS_INVALID_NT_SIGNATURE;
	      }
	      switch (nt_headers->FileHeader.Machine)
	      {
	      case IMAGE_FILE_MACHINE_I386:
	      case IMAGE_FILE_MACHINE_IA64:
	      case IMAGE_FILE_MACHINE_AMD64:
	          break;
	      default:
	          fputs("Unsupported architecture\n", stderr);
	          return ZYDIS_STATUS_UNSUPPORTED_ARCHITECTURE;
	      }
	      switch (nt_headers->OptionalHeader.Magic)
	      {
	      case IMAGE_NT_OPTIONAL_HDR32_MAGIC:
	      case IMAGE_NT_OPTIONAL_HDR64_MAGIC:
	          break;
	      default:
	          fputs("Unsupported architecture\n", stderr);
	          return ZYDIS_STATUS_UNSUPPORTED_ARCHITECTURE;
	      }
	      ZyanStatus status;
	      ZydisPEContext context;
	      if (!ZYAN_SUCCESS((status = ZydisPEContextInit(&context, buffer, size))))
	      {
	          goto Exit;
	      }
	      if (!ZYAN_SUCCESS((status = DisassembleMappedPEFile(&context))))
	      {
	          ZydisPEContextFinalize(&context);
	          goto Exit;
	      }
	      status = ZydisPEContextFinalize(&context);
	  Exit:
	      ZYAN_FREE(buffer);
	      fclose(file);
	      return status;
	  }*/
	return true
	}
