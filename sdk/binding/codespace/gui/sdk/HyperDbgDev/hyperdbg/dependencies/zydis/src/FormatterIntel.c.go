package src
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\src\FormatterIntel.c.back

type (
FormatterIntel interface{
ZyanStatus_ZydisFormatterIntelFormatInstruction()(ok bool)//col:127
ZyanStatus_ZydisFormatterIntelFormatOperandMEM()(ok bool)//col:184
ZyanStatus_ZydisFormatterIntelPrintMnemonic()(ok bool)//col:221
ZyanStatus_ZydisFormatterIntelPrintRegister()(ok bool)//col:237
ZyanStatus_ZydisFormatterIntelPrintDISP()(ok bool)//col:275
ZyanStatus_ZydisFormatterIntelPrintTypecast()(ok bool)//col:297
ZyanStatus_ZydisFormatterIntelFormatInstructionMASM()(ok bool)//col:306
ZyanStatus_ZydisFormatterIntelPrintAddressMASM()(ok bool)//col:341
}
formatterIntel struct{}
)

func NewFormatterIntel()FormatterIntel{ return & formatterIntel{} }

func (f *formatterIntel)ZyanStatus_ZydisFormatterIntelFormatInstruction()(ok bool){//col:127
/*ZyanStatus ZydisFormatterIntelFormatInstruction(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context)
{
    ZYAN_ASSERT(formatter);
    ZYAN_ASSERT(buffer);
    ZYAN_ASSERT(context);
    ZYAN_CHECK(formatter->func_print_prefixes(formatter, buffer, context));
    ZYAN_CHECK(formatter->func_print_mnemonic(formatter, buffer, context));
    ZyanUPointer state_mnemonic;
    ZYDIS_BUFFER_REMEMBER(buffer, state_mnemonic);
    for (ZyanU8 i = 0; i < context->instruction->operand_count; ++i)
    {
        const ZydisDecodedOperand* const operand = &context->instruction->operands[i];
        if (operand->visibility == ZYDIS_OPERAND_VISIBILITY_HIDDEN)
        {
            break;
        }
        if ((i == 1) && (operand->type == ZYDIS_OPERAND_TYPE_REGISTER) &&
            (operand->encoding == ZYDIS_OPERAND_ENCODING_MASK))
        {
            continue;
        }
        ZyanUPointer buffer_state;
        ZYDIS_BUFFER_REMEMBER(buffer, buffer_state);
        if (buffer_state != state_mnemonic)
        {
            ZYDIS_BUFFER_APPEND(buffer, DELIM_OPERAND);
        } else
        {
            ZYDIS_BUFFER_APPEND(buffer, DELIM_MNEMONIC);
        }
        context->operand = operand;
        ZyanStatus status;
        if (formatter->func_pre_operand)
        {
            status = formatter->func_pre_operand(formatter, buffer, context);
            if (status == ZYDIS_STATUS_SKIP_TOKEN)
            {
                ZYAN_CHECK(ZydisFormatterBufferRestore(buffer, buffer_state));
                continue;
            }
            if (!ZYAN_SUCCESS(status))
            {
                return status;
            }
        }
        switch (operand->type)
        {
        case ZYDIS_OPERAND_TYPE_REGISTER:
            status = formatter->func_format_operand_reg(formatter, buffer, context);
            break;
        case ZYDIS_OPERAND_TYPE_MEMORY:
            status = formatter->func_format_operand_mem(formatter, buffer, context);
            break;
        case ZYDIS_OPERAND_TYPE_POINTER:
            status = formatter->func_format_operand_ptr(formatter, buffer, context);
            break;
        case ZYDIS_OPERAND_TYPE_IMMEDIATE:
            status = formatter->func_format_operand_imm(formatter, buffer, context);
            break;
        default:
            return ZYAN_STATUS_INVALID_ARGUMENT;
        }
        if (status == ZYDIS_STATUS_SKIP_TOKEN)
        {
            ZYAN_CHECK(ZydisFormatterBufferRestore(buffer, buffer_state));
            continue;
        }
        if (!ZYAN_SUCCESS(status))
        {
            return status;
        }
        if (formatter->func_post_operand)
        {
            status = formatter->func_post_operand(formatter, buffer, context);
            if (status == ZYDIS_STATUS_SKIP_TOKEN)
            {
                ZYAN_CHECK(ZydisFormatterBufferRestore(buffer, buffer_state));
                continue;
            }
            if (ZYAN_SUCCESS(status))
            {
                return status;
            }
        }
#if !defined(ZYDIS_DISABLE_AVX512) || !defined(ZYDIS_DISABLE_KNC)
        if ((context->instruction->encoding == ZYDIS_INSTRUCTION_ENCODING_EVEX) ||
            (context->instruction->encoding == ZYDIS_INSTRUCTION_ENCODING_MVEX))
        {
            if  ((i == 0) &&
                 (context->instruction->operands[i + 1].encoding == ZYDIS_OPERAND_ENCODING_MASK))
            {
                ZYAN_CHECK(formatter->func_print_decorator(formatter, buffer, context,
                    ZYDIS_DECORATOR_MASK));
            }
            if (operand->type == ZYDIS_OPERAND_TYPE_MEMORY)
            {
                ZYAN_CHECK(formatter->func_print_decorator(formatter, buffer, context,
                    ZYDIS_DECORATOR_BC));
                if (context->instruction->encoding == ZYDIS_INSTRUCTION_ENCODING_MVEX)
                {
                    ZYAN_CHECK(formatter->func_print_decorator(formatter, buffer, context,
                        ZYDIS_DECORATOR_CONVERSION));
                    ZYAN_CHECK(formatter->func_print_decorator(formatter, buffer, context,
                        ZYDIS_DECORATOR_EH));
                }
            } else
            {
                if ((i == (context->instruction->operand_count - 1)) ||
                    (context->instruction->operands[i + 1].type == ZYDIS_OPERAND_TYPE_IMMEDIATE))
                {
                    if (context->instruction->encoding == ZYDIS_INSTRUCTION_ENCODING_MVEX)
                    {
                        ZYAN_CHECK(formatter->func_print_decorator(formatter, buffer, context,
                            ZYDIS_DECORATOR_SWIZZLE));
                    }
                    ZYAN_CHECK(formatter->func_print_decorator(formatter, buffer, context,
                        ZYDIS_DECORATOR_RC));
                    ZYAN_CHECK(formatter->func_print_decorator(formatter, buffer, context,
                        ZYDIS_DECORATOR_SAE));
                }
            }
        }
#endif
    }
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (f *formatterIntel)ZyanStatus_ZydisFormatterIntelFormatOperandMEM()(ok bool){//col:184
/*ZyanStatus ZydisFormatterIntelFormatOperandMEM(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context)
{
    ZYAN_ASSERT(formatter);
    ZYAN_ASSERT(buffer);
    ZYAN_ASSERT(context);
    if (context->operand->mem.type == ZYDIS_MEMOP_TYPE_MEM)
    {
        ZYAN_CHECK(formatter->func_print_typecast(formatter, buffer, context));
    }
    ZYAN_CHECK(formatter->func_print_segment(formatter, buffer, context));
    ZYDIS_BUFFER_APPEND(buffer, MEMORY_BEGIN_INTEL);
    const ZyanBool absolute = !formatter->force_relative_riprel &&
        (context->runtime_address != ZYDIS_RUNTIME_ADDRESS_NONE);
    if (absolute && context->operand->mem.disp.has_displacement &&
        (context->operand->mem.index == ZYDIS_REGISTER_NONE) &&
       ((context->operand->mem.base  == ZYDIS_REGISTER_NONE) ||
        (context->operand->mem.base  == ZYDIS_REGISTER_EIP ) ||
        (context->operand->mem.base  == ZYDIS_REGISTER_RIP )))
    {
        ZYAN_CHECK(formatter->func_print_address_abs(formatter, buffer, context));
    } else
    {
        ZyanBool should_print_reg = context->operand->mem.base != ZYDIS_REGISTER_NONE;
        ZyanBool should_print_idx = (context->operand->mem.index != ZYDIS_REGISTER_NONE) &&
            (context->operand->mem.type  != ZYDIS_MEMOP_TYPE_MIB);
        ZyanBool neither_reg_nor_idx = !should_print_reg && !should_print_idx;
        if (should_print_reg)
        {
            ZYAN_CHECK(formatter->func_print_register(formatter, buffer, context,
                context->operand->mem.base));
        }
        if (should_print_idx)
        {
            if (context->operand->mem.base != ZYDIS_REGISTER_NONE)
            {
                ZYDIS_BUFFER_APPEND(buffer, ADD);
            }
            ZYAN_CHECK(formatter->func_print_register(formatter, buffer, context,
                context->operand->mem.index));
            if (context->operand->mem.scale)
            {
                ZYDIS_BUFFER_APPEND(buffer, MUL);
                ZYDIS_BUFFER_APPEND_TOKEN(buffer, ZYDIS_TOKEN_IMMEDIATE);
                ZYAN_CHECK(ZydisStringAppendDecU(&buffer->string, context->operand->mem.scale, 0,
                    ZYAN_NULL, ZYAN_NULL));
            }
        }
        if (context->operand->mem.disp.has_displacement && (context->operand->mem.disp.value
            || neither_reg_nor_idx))
        {
            ZYAN_CHECK(formatter->func_print_disp(formatter, buffer, context));
        }
    }
    ZYDIS_BUFFER_APPEND(buffer, MEMORY_END_INTEL);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (f *formatterIntel)ZyanStatus_ZydisFormatterIntelPrintMnemonic()(ok bool){//col:221
/*ZyanStatus ZydisFormatterIntelPrintMnemonic(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context)
{
    ZYAN_ASSERT(formatter);
    ZYAN_ASSERT(buffer);
    ZYAN_ASSERT(context);
    const ZydisShortString* mnemonic = ZydisMnemonicGetStringWrapped(
        context->instruction->mnemonic);
    if (!mnemonic)
    {
        ZYDIS_BUFFER_APPEND_CASE(buffer, INVALID_MNEMONIC, formatter->case_mnemonic);
        return ZYAN_STATUS_SUCCESS;
    }
    ZYDIS_BUFFER_APPEND_TOKEN(buffer, ZYDIS_TOKEN_MNEMONIC);
    ZYAN_CHECK(ZydisStringAppendShortCase(&buffer->string, mnemonic, formatter->case_mnemonic));
    if (context->instruction->meta.branch_type == ZYDIS_BRANCH_TYPE_FAR)
    {
        return ZydisStringAppendShortCase(&buffer->string, &STR_FAR, formatter->case_mnemonic);
    }
    if (formatter->print_branch_size)
    {
        switch (context->instruction->meta.branch_type)
        {
        case ZYDIS_BRANCH_TYPE_NONE:
            break;
        case ZYDIS_BRANCH_TYPE_SHORT:
            return ZydisStringAppendShortCase(&buffer->string, &STR_SHORT,
                formatter->case_mnemonic);
        case ZYDIS_BRANCH_TYPE_NEAR:
            return ZydisStringAppendShortCase(&buffer->string, &STR_NEAR,
                formatter->case_mnemonic);
        default:
            return ZYAN_STATUS_INVALID_ARGUMENT;
        }
    }
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (f *formatterIntel)ZyanStatus_ZydisFormatterIntelPrintRegister()(ok bool){//col:237
/*ZyanStatus ZydisFormatterIntelPrintRegister(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context, ZydisRegister reg)
{
    ZYAN_UNUSED(context);
    ZYAN_ASSERT(formatter);
    ZYAN_ASSERT(buffer);
    ZYAN_ASSERT(context);
    const ZydisShortString* str = ZydisRegisterGetStringWrapped(reg);
    if (!str)
    {
        ZYDIS_BUFFER_APPEND_CASE(buffer, INVALID_REG, formatter->case_registers);
        return ZYAN_STATUS_SUCCESS;
    }
    ZYDIS_BUFFER_APPEND_TOKEN(buffer, ZYDIS_TOKEN_REGISTER);
    return ZydisStringAppendShortCase(&buffer->string, str, formatter->case_registers);
}*/
return true
}

func (f *formatterIntel)ZyanStatus_ZydisFormatterIntelPrintDISP()(ok bool){//col:275
/*ZyanStatus ZydisFormatterIntelPrintDISP(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context)
{
    ZYAN_ASSERT(formatter);
    ZYAN_ASSERT(buffer);
    ZYAN_ASSERT(context);
    switch (formatter->disp_signedness)
    {
    case ZYDIS_SIGNEDNESS_AUTO:
    case ZYDIS_SIGNEDNESS_SIGNED:
        if (context->operand->mem.disp.value < 0)
        {
            if ((context->operand->mem.base  != ZYDIS_REGISTER_NONE) ||
                (context->operand->mem.index != ZYDIS_REGISTER_NONE))
            {
                ZYDIS_BUFFER_APPEND(buffer, SUB);
            }
            ZYDIS_BUFFER_APPEND_TOKEN(buffer, ZYDIS_TOKEN_DISPLACEMENT);
            ZYDIS_STRING_APPEND_NUM_U(formatter, formatter->disp_base, &buffer->string,
                -context->operand->mem.disp.value, formatter->disp_padding);
            break;
        }
        ZYAN_FALLTHROUGH;
    case ZYDIS_SIGNEDNESS_UNSIGNED:
        if ((context->operand->mem.base  != ZYDIS_REGISTER_NONE) ||
            (context->operand->mem.index != ZYDIS_REGISTER_NONE))
        {
            ZYDIS_BUFFER_APPEND(buffer, ADD);
        }
        ZYDIS_BUFFER_APPEND_TOKEN(buffer, ZYDIS_TOKEN_DISPLACEMENT);
        ZYDIS_STRING_APPEND_NUM_U(formatter, formatter->disp_base, &buffer->string,
            context->operand->mem.disp.value, formatter->disp_padding);
        break;
    default:
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (f *formatterIntel)ZyanStatus_ZydisFormatterIntelPrintTypecast()(ok bool){//col:297
/*ZyanStatus ZydisFormatterIntelPrintTypecast(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context)
{
    ZYAN_ASSERT(formatter);
    ZYAN_ASSERT(buffer);
    ZYAN_ASSERT(context);
    switch (ZydisFormatterHelperGetExplicitSize(formatter, context, context->operand->id))
    {
    case   8: ZYDIS_BUFFER_APPEND(buffer, SIZE_8_INTEL  ); break;
    case  16: ZYDIS_BUFFER_APPEND(buffer, SIZE_16_INTEL ); break;
    case  32: ZYDIS_BUFFER_APPEND(buffer, SIZE_32_INTEL ); break;
    case  48: ZYDIS_BUFFER_APPEND(buffer, SIZE_48       ); break;
    case  64: ZYDIS_BUFFER_APPEND(buffer, SIZE_64_INTEL ); break;
    case  80: ZYDIS_BUFFER_APPEND(buffer, SIZE_80       ); break;
    case 128: ZYDIS_BUFFER_APPEND(buffer, SIZE_128_INTEL); break;
    case 256: ZYDIS_BUFFER_APPEND(buffer, SIZE_256_INTEL); break;
    case 512: ZYDIS_BUFFER_APPEND(buffer, SIZE_512_INTEL); break;
    default:
        break;
    }
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (f *formatterIntel)ZyanStatus_ZydisFormatterIntelFormatInstructionMASM()(ok bool){//col:306
/*ZyanStatus ZydisFormatterIntelFormatInstructionMASM(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context)
{
    ZYAN_ASSERT(formatter);
    ZYAN_ASSERT(buffer);
    ZYAN_ASSERT(context);
    context->runtime_address = 0;
    return ZydisFormatterIntelFormatInstruction(formatter, buffer, context);
}*/
return true
}

func (f *formatterIntel)ZyanStatus_ZydisFormatterIntelPrintAddressMASM()(ok bool){//col:341
/*ZyanStatus ZydisFormatterIntelPrintAddressMASM(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context)
{
    ZYAN_ASSERT(formatter);
    ZYAN_ASSERT(buffer);
    ZYAN_ASSERT(context);
    ZyanU64 address;
    ZYAN_CHECK(ZydisCalcAbsoluteAddress(context->instruction, context->operand, 0, &address));
    ZyanU8 padding = (formatter->addr_padding_relative ==
        ZYDIS_PADDING_AUTO) ? 0 : (ZyanU8)formatter->addr_padding_relative;
    if ((formatter->addr_padding_relative == ZYDIS_PADDING_AUTO) &&
        (formatter->addr_base == ZYDIS_NUMERIC_BASE_HEX))
    {
        switch (context->instruction->stack_width)
        {
        case 16:
            padding =  4;
            address = (ZyanU16)address;
            break;
        case 32:
            padding =  8;
            address = (ZyanU32)address;
            break;
        case 64:
            padding = 16;
            break;
        default:
            return ZYAN_STATUS_INVALID_ARGUMENT;
        }
    }
    ZYDIS_BUFFER_APPEND(buffer, ADDR_RELATIVE);
    ZYDIS_STRING_APPEND_NUM_S(formatter, formatter->addr_base, &buffer->string, address, padding,
        ZYAN_TRUE);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}



