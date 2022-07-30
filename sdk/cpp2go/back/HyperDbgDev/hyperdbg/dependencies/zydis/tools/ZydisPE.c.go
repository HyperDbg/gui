package tools
//back\HyperDbgDev\hyperdbg\dependencies\zydis\tools\ZydisPE.c.back

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
/*  Zyan Disassembler Library (Zydis)
  Original Author : Florian Bernd
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 * Disassembles a given PE file.
#include <errno.h>
#include <inttypes.h>
#include <stdio.h>
#include <stddef.h>
#include <string.h>
#include <Zycore/LibC.h>
#include <Zycore/Vector.h>
#include <Zydis/Zydis.h>
static const ZyanStringView STR_DOT = ZYAN_DEFINE_STRING_VIEW(".");
static const ZyanStringView STR_ENTRY_POINT = ZYAN_DEFINE_STRING_VIEW("EntryPoint");
 * The zydis PE tool module id.
#define ZYAN_MODULE_ZYDIS_PE    0x101
 * The signature of the PE-files DOS header field is invalid.
#define ZYDIS_STATUS_INVALID_DOS_SIGNATURE \
    ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYDIS_PE, 0x00)
 * The signature of the PE-files NT headers field is invalid.
#define ZYDIS_STATUS_INVALID_NT_SIGNATURE \
    ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYDIS_PE, 0x01)
 * The architecture of the assembly code contained in the PE-file is not supported.
#define ZYDIS_STATUS_UNSUPPORTED_ARCHITECTURE \
    ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYDIS_PE, 0x02)
#define IMAGE_NT_OPTIONAL_HDR32_MAGIC                   0x010B
#define IMAGE_NT_OPTIONAL_HDR64_MAGIC                   0x020B
#define IMAGE_NUMBEROF_DIRECTORY_ENTRIES                    16
#define IMAGE_DIRECTORY_ENTRY_EXPORT                         0
#define IMAGE_DIRECTORY_ENTRY_IMPORT                         1
#define IMAGE_FILE_MACHINE_I386                         0x014C
#define IMAGE_FILE_MACHINE_IA64                         0x0200
#define IMAGE_FILE_MACHINE_AMD64                        0x8664
#define IMAGE_SCN_CNT_CODE                          0x00000020
#define IMAGE_IMPORT_BY_ORDINAL32                   0x80000000
#define IMAGE_IMPORT_BY_ORDINAL64           0x8000000000000000
typedef struct IMAGE_DOS_HEADER_
{
    ZyanU16 e_magic;
    ZyanU16 e_cblp;
    ZyanU16 e_cp;
    ZyanU16 e_crlc;
    ZyanU16 e_cparhdr;
    ZyanU16 e_minalloc;
    ZyanU16 e_maxalloc;
    ZyanU16 e_ss;
    ZyanU16 e_sp;
    ZyanU16 e_csum;
    ZyanU16 e_ip;
    ZyanU16 e_cs;
    ZyanU16 e_lfarlc;
    ZyanU16 e_ovno;
    ZyanU16 e_res[4];
    ZyanU16 e_oemid;
    ZyanU16 e_oeminfo;
    ZyanU16 e_res2[10];
    ZyanU32 e_lfanew;
} IMAGE_DOS_HEADER;*/
return true
}

func (z *zydisPe)#pragma pack()(ok bool){//col:310
/*#pragma pack(push, 8)
typedef struct IMAGE_THUNK_DATA64_
{
    union
    {
        ZyanU64 ForwarderString;
        ZyanU64 Function;
        ZyanU64 Ordinal;
        ZyanU64 AddressOfData;
    } u1;
} IMAGE_THUNK_DATA64;*/
return true
}

func (z *zydisPe)#pragma pack()(ok bool){//col:318
/*#pragma pack(pop)
typedef struct IMAGE_IMPORT_BY_NAME_
{
    ZyanU16 Hint;
    char    Name[1];
} IMAGE_IMPORT_BY_NAME;*/
return true
}

func (z *zydisPe)#define IMAGE_FIRST_SECTION()(ok bool){//col:356
/*#define IMAGE_FIRST_SECTION(nt_headers) \
    ((IMAGE_SECTION_HEADER*)((ZyanUPointer)(nt_headers) \
        + offsetof(IMAGE_NT_HEADERS32, OptionalHeader) \
        + ((nt_headers))->FileHeader.SizeOfOptionalHeader))
 * Defines the `ZydisPESymbol` struct.
typedef struct ZydisPESymbol_
{
     * The virtual address of the symbol.
    ZyanU64 address;
     * The module string.
    ZyanString module_name;
     * The symbol string.
    ZyanString symbol_name;
} ZydisPESymbol;*/
return true
}

func (z *zydisPe)static ZyanI32 CompareSymbol()(ok bool){//col:413
/*static ZyanI32 CompareSymbol(const ZydisPESymbol* left, const ZydisPESymbol* right)
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

func (z *zydisPe)static const IMAGE_SECTION_HEADER* GetSectionByRVA()(ok bool){//col:452
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

func (z *zydisPe)const void* RVAToFileOffset()(ok bool){//col:474
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

func (z *zydisPe)static ZyanStatus ZydisPEContextFinalize()(ok bool){//col:499
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

func (z *zydisPe)static ZyanStatus ZydisPEContextInit()(ok bool){//col:871
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
            {
                if (element.module_name.buffer[i] == '.')
                {
                    element.module_name.length = i;
                    break;
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
                {
                    if (element.module_name.buffer[i] == '.')
                    {
                        element.module_name.length = i;
                        break;
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

func (z *zydisPe)static ZyanStatus ZydisFormatterPrintAddress()(ok bool){//col:921
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

func (z *zydisPe)static ZyanStatus ZydisFormatterPrintAddressABS()(ok bool){//col:931
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

func (z *zydisPe)static ZyanStatus ZydisFormatterPrintAddressREL()(ok bool){//col:940
/*static ZyanStatus ZydisFormatterPrintAddressREL(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context)
{
    ZyanU64 address;
    ZYAN_CHECK(ZydisCalcAbsoluteAddress(context->instruction, context->operand, 0, &address));
    return ZydisFormatterPrintAddress(formatter, buffer, context, address, ZYAN_FALSE);
}*/
return true
}

func (z *zydisPe)static ZyanStatus DisassembleMappedPEFile()(ok bool){//col:1086
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

func (z *zydisPe)int main()(ok bool){//col:1186
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



