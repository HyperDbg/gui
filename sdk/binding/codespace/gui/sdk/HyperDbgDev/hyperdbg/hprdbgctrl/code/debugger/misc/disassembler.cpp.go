package misc
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\misc\disassembler.cpp.back

const(
NDEBUG =  //col:1
PaddingLength = 12 //col:2
)

type typedef struct ZydisSymbol_ struct{
address ZyanU64
char const
}



type (
Disassembler interface{
ZydisFormatterPrintAddressAbsolute()(ok bool)//col:23
DisassembleBuffer()(ok bool)//col:117
ZydisTest()(ok bool)//col:158
HyperDbgDisassembler64()(ok bool)//col:175
HyperDbgDisassembler32()(ok bool)//col:192
HyperDbgIsConditionalJumpTaken()(ok bool)//col:334
HyperDbgCheckWhetherTheCurrentInstructionIsCall()(ok bool)//col:381
HyperDbgLengthDisassemblerEngine()(ok bool)//col:419
}
)

func NewDisassembler() { return & disassembler{} }

func (d *disassembler)ZydisFormatterPrintAddressAbsolute()(ok bool){//col:23
/*ZydisFormatterPrintAddressAbsolute(const ZydisFormatter *  formatter,
                                   ZydisFormatterBuffer *  buffer,
                                   ZydisFormatterContext * context)
{
    ZyanU64                                                address;
    std::map<UINT64, LOCAL_FUNCTION_DESCRIPTION>::iterator Iterate;
    ZYAN_CHECK(ZydisCalcAbsoluteAddress(context->instruction, context->operand, context->runtime_address, &address));
    if (g_AddressConversion)
    {
        Iterate = g_DisassemblerSymbolMap.find(address);
        if (Iterate != g_DisassemblerSymbolMap.end())
        {
            ZYAN_CHECK(ZydisFormatterBufferAppend(buffer, ZYDIS_TOKEN_SYMBOL));
            ZyanString * string;
            ZYAN_CHECK(ZydisFormatterBufferGetString(buffer, &string));
            return ZyanStringAppendFormat(string,
                                          "<%s (%s)>",
                                          Iterate->second.ObjectName.c_str(),
                                          SeparateTo64BitValue(Iterate->first).c_str());
        }
    }
    return default_print_address_absolute(formatter, buffer, context);
}*/
return true
}

func (d *disassembler)DisassembleBuffer()(ok bool){//col:117
/*DisassembleBuffer(ZydisDecoder * decoder,
                  ZyanU64        runtime_address,
                  ZyanU8 *       data,
                  ZyanUSize      length,
                  uint32_t       maximum_instr,
                  BOOLEAN        is_x86_64,
                  BOOLEAN        show_of_branch_is_taken,
                  PRFLAGS        rflags)
{
    ZydisFormatter formatter;
    int            instr_decoded   = 0;
    UINT64         UsedBaseAddress = NULL;
    if (g_DisassemblerSyntax == 1)
    {
        ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_INTEL);
    }
    else if (g_DisassemblerSyntax == 2)
    {
        ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_ATT);
    }
    else if (g_DisassemblerSyntax == 3)
    {
        ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_INTEL_MASM);
    }
    else
    {
        ShowMessages("err, in selecting disassembler syntax\n");
        return;
    }
    ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SEGMENT, ZYAN_TRUE);
    ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SIZE, ZYAN_TRUE);
    default_print_address_absolute =
        (ZydisFormatterFunc)&ZydisFormatterPrintAddressAbsolute;
    ZydisFormatterSetHook(&formatter, ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS, (const void **)&default_print_address_absolute);
    ZydisDecodedInstruction instruction;
    char                    buffer[256];
    while (ZYAN_SUCCESS(
        ZydisDecoderDecodeBuffer(decoder, data, length, &instruction)))
    {
        if (g_AddressConversion)
        {
            if (SymbolShowFunctionNameBasedOnAddress(runtime_address, &UsedBaseAddress))
            {
                ShowMessages(":\n");
            }
        }
        ShowMessages("%s   ", SeparateTo64BitValue(runtime_address).c_str());
        ZydisFormatterFormatInstruction(&formatter, &instruction, &buffer[0], sizeof(buffer), runtime_address);
        for (size_t i = 0; i < instruction.length; i++)
        {
            ZyanU8 MemoryContent = data[i];
            ShowMessages(" %02X", MemoryContent);
        }
        if (instruction.length < PaddingLength)
        {
            for (size_t i = 0; i < PaddingLength - instruction.length; i++)
            {
                ShowMessages("   ");
            }
        }
        if (show_of_branch_is_taken)
        {
            RFLAGS TempRflags = {0};
            TempRflags.AsUInt = rflags->AsUInt;
            DEBUGGER_CONDITIONAL_JUMP_STATUS ResultOfCondJmp =
                HyperDbgIsConditionalJumpTaken(data, length, TempRflags, is_x86_64);
            if (ResultOfCondJmp == DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN)
            {
                ShowMessages(" %s [taken]\n", &buffer[0]);
            }
            else if (ResultOfCondJmp ==
                     DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN)
            {
                ShowMessages(" %s [not taken]\n", &buffer[0]);
            }
            else
            {
                ShowMessages(" %s\n", &buffer[0]);
            }
        }
        else
        {
            ShowMessages(" %s\n", &buffer[0]);
        }
        data += instruction.length;
        length -= instruction.length;
        runtime_address += instruction.length;
        instr_decoded++;
        if (instr_decoded == maximum_instr)
        {
            return;
        }
    }
}*/
return true
}

func (d *disassembler)ZydisTest()(ok bool){//col:158
/*ZydisTest()
{
    if (ZydisGetVersion() != ZYDIS_VERSION)
    {
        fputs("Invalid zydis version\n", ZYAN_STDERR);
        return EXIT_FAILURE;
    }
    ZyanU8 data[] = {
        0x48,
        0x8B,
        0x05,
        0x39,
        0x00,
        0x13,
        0x00, 
        0x50, 
        0xFF,
        0x15,
        0xF2,
        0x10,
        0x00,
        0x00, 
        0x85,
        0xC0, 
        0x0F,
        0x84,
        0x00,
        0x00,
        0x00,
        0x00, 
        0xE9,
        0xE5,
        0x0F,
        0x00,
        0x00 
    };
    ZydisDecoder decoder;
    ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_64, ZYDIS_ADDRESS_WIDTH_64);
    DisassembleBuffer(&decoder, 0x007FFFFFFF400000, &data[0], sizeof(data), 0xffffffff, TRUE, FALSE, NULL);
    return 0;
}*/
return true
}

func (d *disassembler)HyperDbgDisassembler64()(ok bool){//col:175
/*HyperDbgDisassembler64(unsigned char * BufferToDisassemble,
                       UINT64          BaseAddress,
                       UINT64          Size,
                       UINT32          MaximumInstrDecoded,
                       BOOLEAN         ShowBranchIsTakenOrNot,
                       PRFLAGS         Rflags)
{
    if (ZydisGetVersion() != ZYDIS_VERSION)
    {
        fputs("Invalid zydis version\n", ZYAN_STDERR);
        return EXIT_FAILURE;
    }
    ZydisDecoder decoder;
    ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_64, ZYDIS_ADDRESS_WIDTH_64);
    DisassembleBuffer(&decoder, BaseAddress, &BufferToDisassemble[0], Size, MaximumInstrDecoded, TRUE, ShowBranchIsTakenOrNot, Rflags);
    return 0;
}*/
return true
}

func (d *disassembler)HyperDbgDisassembler32()(ok bool){//col:192
/*HyperDbgDisassembler32(unsigned char * BufferToDisassemble,
                       UINT64          BaseAddress,
                       UINT64          Size,
                       UINT32          MaximumInstrDecoded,
                       BOOLEAN         ShowBranchIsTakenOrNot,
                       PRFLAGS         Rflags)
{
    if (ZydisGetVersion() != ZYDIS_VERSION)
    {
        fputs("Invalid zydis version\n", ZYAN_STDERR);
        return EXIT_FAILURE;
    }
    ZydisDecoder decoder;
    ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_COMPAT_32, ZYDIS_ADDRESS_WIDTH_32);
    DisassembleBuffer(&decoder, (UINT32)BaseAddress, &BufferToDisassemble[0], Size, MaximumInstrDecoded, FALSE, ShowBranchIsTakenOrNot, Rflags);
    return 0;
}*/
return true
}

func (d *disassembler)HyperDbgIsConditionalJumpTaken()(ok bool){//col:334
/*HyperDbgIsConditionalJumpTaken(unsigned char * BufferToDisassemble,
                               UINT64          BuffLength,
                               RFLAGS          Rflags,
                               BOOLEAN         Isx86_64)
{
    ZydisDecoder            decoder;
    ZydisFormatter          formatter;
    UINT64                  CurrentRip    = 0;
    int                     instr_decoded = 0;
    ZydisDecodedInstruction instruction;
    char                    buffer[256];
    UINT32                  MaximumInstrDecoded = 1;
    if (ZydisGetVersion() != ZYDIS_VERSION)
    {
        ShowMessages("invalid zydis version\n");
        return DEBUGGER_CONDITIONAL_JUMP_STATUS_ERROR;
    }
    if (Isx86_64)
    {
        ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_64, ZYDIS_ADDRESS_WIDTH_64);
    }
    else
    {
        ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_COMPAT_32, ZYDIS_ADDRESS_WIDTH_32);
    }
    ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_INTEL);
    ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SEGMENT, ZYAN_TRUE);
    ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SIZE, ZYAN_TRUE);
    default_print_address_absolute =
        (ZydisFormatterFunc)&ZydisFormatterPrintAddressAbsolute;
    ZydisFormatterSetHook(&formatter, ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS, (const void **)&default_print_address_absolute);
    while (ZYAN_SUCCESS(ZydisDecoderDecodeBuffer(&decoder, BufferToDisassemble, BuffLength, &instruction)))
    {
        ZydisFormatterFormatInstruction(&formatter, &instruction, &buffer[0], sizeof(buffer), (ZyanU64)CurrentRip);
        switch (instruction.mnemonic)
        {
        case ZydisMnemonic::ZYDIS_MNEMONIC_JO:
            if (Rflags.OverflowFlag)
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
            else
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
            break;
        case ZydisMnemonic::ZYDIS_MNEMONIC_JNO:
            if (!Rflags.OverflowFlag)
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
            else
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
            break;
        case ZydisMnemonic::ZYDIS_MNEMONIC_JS:
            if (Rflags.SignFlag)
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
            else
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
            break;
        case ZydisMnemonic::ZYDIS_MNEMONIC_JNS:
            if (!Rflags.SignFlag)
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
            else
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
            break;
        case ZydisMnemonic::ZYDIS_MNEMONIC_JZ:
            if (Rflags.ZeroFlag)
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
            else
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
            break;
        case ZydisMnemonic::ZYDIS_MNEMONIC_JNZ:
            if (!Rflags.ZeroFlag)
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
            else
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
            break;
        case ZydisMnemonic::ZYDIS_MNEMONIC_JB:
            if (Rflags.CarryFlag)
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
            else
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
            break;
        case ZydisMnemonic::ZYDIS_MNEMONIC_JNB:
            if (!Rflags.CarryFlag)
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
            else
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
            break;
        case ZydisMnemonic::ZYDIS_MNEMONIC_JBE:
            if (Rflags.CarryFlag || Rflags.ZeroFlag)
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
            else
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
            break;
        case ZydisMnemonic::ZYDIS_MNEMONIC_JNBE:
            if (!Rflags.CarryFlag && !Rflags.ZeroFlag)
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
            else
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
            break;
        case ZydisMnemonic::ZYDIS_MNEMONIC_JL:
            if (Rflags.SignFlag != Rflags.OverflowFlag)
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
            else
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
            break;
        case ZydisMnemonic::ZYDIS_MNEMONIC_JNL:
            if (Rflags.SignFlag == Rflags.OverflowFlag)
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
            else
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
            break;
        case ZydisMnemonic::ZYDIS_MNEMONIC_JLE:
            if (Rflags.ZeroFlag || Rflags.SignFlag != Rflags.OverflowFlag)
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
            else
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
            break;
        case ZydisMnemonic::ZYDIS_MNEMONIC_JNLE:
            if (!Rflags.ZeroFlag && Rflags.SignFlag == Rflags.OverflowFlag)
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
            else
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
            break;
        case ZydisMnemonic::ZYDIS_MNEMONIC_JP:
            if (Rflags.ParityFlag)
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
            else
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
            break;
        case ZydisMnemonic::ZYDIS_MNEMONIC_JNP:
            if (!Rflags.ParityFlag)
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
            else
                return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
            break;
        case ZydisMnemonic::ZYDIS_MNEMONIC_JCXZ:
        case ZydisMnemonic::ZYDIS_MNEMONIC_JECXZ:
            return DEBUGGER_CONDITIONAL_JUMP_STATUS_NOT_CONDITIONAL_JUMP;
        default:
            return DEBUGGER_CONDITIONAL_JUMP_STATUS_NOT_CONDITIONAL_JUMP;
            break;
        }
        return DEBUGGER_CONDITIONAL_JUMP_STATUS_ERROR;
    }
}*/
return true
}

func (d *disassembler)HyperDbgCheckWhetherTheCurrentInstructionIsCall()(ok bool){//col:381
/*HyperDbgCheckWhetherTheCurrentInstructionIsCall(
    unsigned char * BufferToDisassemble,
    UINT64          BuffLength,
    BOOLEAN         Isx86_64,
    PUINT32         CallLength)
{
    ZydisDecoder            decoder;
    ZydisFormatter          formatter;
    UINT64                  CurrentRip    = 0;
    int                     instr_decoded = 0;
    ZydisDecodedInstruction instruction;
    char                    buffer[256];
    UINT32                  MaximumInstrDecoded = 1;
    *CallLength = 0;
    if (ZydisGetVersion() != ZYDIS_VERSION)
    {
        ShowMessages("invalid zydis version\n");
        return DEBUGGER_CONDITIONAL_JUMP_STATUS_ERROR;
    }
    if (Isx86_64)
    {
        ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_64, ZYDIS_ADDRESS_WIDTH_64);
    }
    else
    {
        ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_COMPAT_32, ZYDIS_ADDRESS_WIDTH_32);
    }
    ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_INTEL);
    ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SEGMENT, ZYAN_TRUE);
    ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SIZE, ZYAN_TRUE);
    default_print_address_absolute =
        (ZydisFormatterFunc)&ZydisFormatterPrintAddressAbsolute;
    ZydisFormatterSetHook(&formatter, ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS, (const void **)&default_print_address_absolute);
    while (ZYAN_SUCCESS(ZydisDecoderDecodeBuffer(&decoder, BufferToDisassemble, BuffLength, &instruction)))
    {
        ZydisFormatterFormatInstruction(&formatter, &instruction, &buffer[0], sizeof(buffer), (ZyanU64)CurrentRip);
        if (instruction.mnemonic == ZydisMnemonic::ZYDIS_MNEMONIC_CALL)
        {
            *CallLength = instruction.length;
            return TRUE;
        }
        else
        {
            return FALSE;
        }
    }
}*/
return true
}

func (d *disassembler)HyperDbgLengthDisassemblerEngine()(ok bool){//col:419
/*HyperDbgLengthDisassemblerEngine(
    unsigned char * BufferToDisassemble,
    UINT64          BuffLength,
    BOOLEAN         Isx86_64)
{
    ZydisDecoder            decoder;
    ZydisFormatter          formatter;
    UINT64                  CurrentRip    = 0;
    int                     instr_decoded = 0;
    ZydisDecodedInstruction instruction;
    char                    buffer[256];
    UINT32                  MaximumInstrDecoded = 1;
    if (ZydisGetVersion() != ZYDIS_VERSION)
    {
        ShowMessages("invalid zydis version\n");
        return DEBUGGER_CONDITIONAL_JUMP_STATUS_ERROR;
    }
    if (Isx86_64)
    {
        ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_64, ZYDIS_ADDRESS_WIDTH_64);
    }
    else
    {
        ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_COMPAT_32, ZYDIS_ADDRESS_WIDTH_32);
    }
    ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_INTEL);
    ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SEGMENT, ZYAN_TRUE);
    ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SIZE, ZYAN_TRUE);
    default_print_address_absolute =
        (ZydisFormatterFunc)&ZydisFormatterPrintAddressAbsolute;
    ZydisFormatterSetHook(&formatter, ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS, (const void **)&default_print_address_absolute);
    while (ZYAN_SUCCESS(ZydisDecoderDecodeBuffer(&decoder, BufferToDisassemble, BuffLength, &instruction)))
    {
        ZydisFormatterFormatInstruction(&formatter, &instruction, &buffer[0], sizeof(buffer), (ZyanU64)CurrentRip);
        return instruction.length;
    }
    return 0;
}*/
return true
}



