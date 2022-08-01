package examples
//back\HyperDbgDev\hyperdbg\dependencies\zydis\examples\ZydisWinKernel.c.back

type (
ZydisWinKernel interface{
RtlPcToFileHeader()(ok bool)//col:32
DriverEntry()(ok bool)//col:103
}
)

func NewZydisWinKernel() { return & zydisWinKernel{} }

func (z *zydisWinKernel)RtlPcToFileHeader()(ok bool){//col:32
/*RtlPcToFileHeader(
    _In_ PVOID PcValue,
    _Out_ PVOID *BaseOfImage
    );
NTKERNELAPI
PIMAGE_NT_HEADERS
NTAPI
RtlImageNtHeader(
    _In_ PVOID ImageBase
    );
#if defined(ZYAN_CLANG) || defined(ZYAN_GNUC)
__attribute__((section("INIT")))
#endif
DRIVER_INITIALIZE
DriverEntry;
#if defined(ALLOC_PRAGMA) && !(defined(ZYAN_CLANG) || defined(ZYAN_GNUC))
#pragma alloc_text(INIT, DriverEntry)
#endif
VOID
Print(
    _In_ PCCH Format,
    _In_ ...
    )
{
    CHAR message[512];
    va_list argList;
    va_start(argList, Format);
    const int n = _vsnprintf_s(message, sizeof(message), sizeof(message) - 1, Format, argList);
    message[n] = '\0';
    vDbgPrintExWithPrefix("[ZYDIS] ", DPFLTR_IHVDRIVER_ID, DPFLTR_ERROR_LEVEL, message, argList);
    va_end(argList);
}*/
return true
}

func (z *zydisWinKernel)DriverEntry()(ok bool){//col:103
/*DriverEntry(
    _In_ PDRIVER_OBJECT DriverObject,
    _In_ PUNICODE_STRING RegistryPath
    )
{
    PAGED_CODE();
    UNREFERENCED_PARAMETER(RegistryPath);
    if (ZydisGetVersion() != ZYDIS_VERSION)
    {
        Print("Invalid zydis version\n");
        return STATUS_UNKNOWN_REVISION;
    }
    ULONG_PTR imageBase;
    RtlPcToFileHeader((PVOID)DriverObject->DriverInit, (PVOID*)&imageBase);
    if (imageBase == 0)
        return STATUS_DRIVER_ENTRYPOINT_NOT_FOUND;
    const PIMAGE_NT_HEADERS ntHeaders = RtlImageNtHeader((PVOID)imageBase);
    if (ntHeaders == NULL)
        return STATUS_INVALID_IMAGE_FORMAT;
    PIMAGE_SECTION_HEADER section = IMAGE_FIRST_SECTION(ntHeaders);
    PIMAGE_SECTION_HEADER initSection = NULL;
    for (USHORT i = 0; i < ntHeaders->FileHeader.NumberOfSections; ++i)
    {
        if (memcmp(section->Name, "INIT", sizeof("INIT") - 1) == 0)
        {
            initSection = section;
            break;
        }
        section++;
    }
    if (initSection == NULL)
        return STATUS_NOT_FOUND;
    const ULONG entryPointRva = (ULONG)((ULONG_PTR)DriverObject->DriverInit - imageBase);
    const ULONG importDirRva = ntHeaders->OptionalHeader.DataDirectory[IMAGE_DIRECTORY_ENTRY_IMPORT].VirtualAddress;
    SIZE_T length = initSection->VirtualAddress + initSection->SizeOfRawData - entryPointRva;
    if (importDirRva > entryPointRva && importDirRva > initSection->VirtualAddress &&
        importDirRva < initSection->VirtualAddress + initSection->SizeOfRawData)
        length = importDirRva - entryPointRva;
    Print("Driver image base: 0x%p, size: 0x%X\n", (PVOID)imageBase, ntHeaders->OptionalHeader.SizeOfImage);
    Print("Entry point RVA: 0x%X (0x%p)\n", entryPointRva, DriverObject->DriverInit);
    ZydisDecoder decoder;
#ifdef _M_AMD64
    if (!ZYAN_SUCCESS(ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_64, ZYDIS_ADDRESS_WIDTH_64)))
#else
    if (!ZYAN_SUCCESS(ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_COMPAT_32, ZYDIS_ADDRESS_WIDTH_32)))
#endif
        return STATUS_DRIVER_INTERNAL_ERROR;
    ZydisFormatter formatter;
    if (!ZYAN_SUCCESS(ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_INTEL)))
        return STATUS_DRIVER_INTERNAL_ERROR;
    SIZE_T readOffset = 0;
    ZydisDecodedInstruction instruction;
    ZyanStatus status;
    CHAR printBuffer[128];
    while ((status = ZydisDecoderDecodeBuffer(&decoder, (PVOID)(imageBase + entryPointRva + readOffset),
        length - readOffset, &instruction)) != ZYDIS_STATUS_NO_MORE_DATA)
    {
        NT_ASSERT(ZYAN_SUCCESS(status));
        if (!ZYAN_SUCCESS(status))
        {
            readOffset++;
            continue;
        }
        const ZyanU64 instrAddress = (ZyanU64)(imageBase + entryPointRva + readOffset);
        ZydisFormatterFormatInstruction(
            &formatter, &instruction, printBuffer, sizeof(printBuffer), instrAddress);
        Print("+%-4X 0x%-16llX\t\t%hs\n", (ULONG)readOffset, instrAddress, printBuffer);
        readOffset += instruction.length;
    }
    return STATUS_UNSUCCESSFUL;
}*/
return true
}



