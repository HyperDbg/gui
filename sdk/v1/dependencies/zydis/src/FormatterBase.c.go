package src

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\src\FormatterBase.c.back

type (
	FormatterBase interface {
		____() (ok bool)                                        //col:17
		ZyanU32_ZydisFormatterHelperGetExplicitSize() (ok bool) //col:61
	}
	formatterBase struct{}
)

func NewFormatterBase() FormatterBase { return &formatterBase{} }

func (f *formatterBase) ____() (ok bool) { //col:17
	/*
		    (const ZydisPredefinedToken* const)&TOK_DATA_PREF_REX_40,
		    (const ZydisPredefinedToken* const)&TOK_DATA_PREF_REX_41,
		    (const ZydisPredefinedToken* const)&TOK_DATA_PREF_REX_42,
		    (const ZydisPredefinedToken* const)&TOK_DATA_PREF_REX_43,
		    (const ZydisPredefinedToken* const)&TOK_DATA_PREF_REX_44,
		    (const ZydisPredefinedToken* const)&TOK_DATA_PREF_REX_45,
		    (const ZydisPredefinedToken* const)&TOK_DATA_PREF_REX_46,
		    (const ZydisPredefinedToken* const)&TOK_DATA_PREF_REX_47,
		    (const ZydisPredefinedToken* const)&TOK_DATA_PREF_REX_48,
		    (const ZydisPredefinedToken* const)&TOK_DATA_PREF_REX_49,
		    (const ZydisPredefinedToken* const)&TOK_DATA_PREF_REX_4A,
		    (const ZydisPredefinedToken* const)&TOK_DATA_PREF_REX_4B,
		    (const ZydisPredefinedToken* const)&TOK_DATA_PREF_REX_4C,
		    (const ZydisPredefinedToken* const)&TOK_DATA_PREF_REX_4D,
		    (const ZydisPredefinedToken* const)&TOK_DATA_PREF_REX_4E,
		    (const ZydisPredefinedToken* const)&TOK_DATA_PREF_REX_4F
		};
	*/
	return true
}

func (f *formatterBase) ZyanU32_ZydisFormatterHelperGetExplicitSize() (ok bool) { //col:61
	/*
	   ZyanU32 ZydisFormatterHelperGetExplicitSize(const ZydisFormatter* formatter,

	   	ZydisFormatterContext* context, ZyanU8 memop_id)

	   	{
	   	    ZYAN_ASSERT(formatter);
	   	    ZYAN_ASSERT(context);
	   	    ZYAN_ASSERT(memop_id < context->instruction->operand_count);
	   	    ZYAN_ASSERT(operand->type == ZYDIS_OPERAND_TYPE_MEMORY);
	   	    ZYAN_ASSERT(operand->mem.type == ZYDIS_MEMOP_TYPE_MEM);
	   	    if (formatter->force_memory_size)
	   	    {
	   	        return operand->size;
	   	    }
	   	    switch (operand->id)
	   	    {
	   	    case 0:
	   	        if ((context->instruction->operands[1].type == ZYDIS_OPERAND_TYPE_UNUSED) ||
	   	            (context->instruction->operands[1].type == ZYDIS_OPERAND_TYPE_IMMEDIATE))
	   	        {
	   	            return context->instruction->operands[0].size;
	   	        }
	   	        if (context->instruction->operands[0].size != context->instruction->operands[1].size)
	   	        {
	   	            return context->instruction->operands[0].size;
	   	        }
	   	        if ((context->instruction->operands[1].type == ZYDIS_OPERAND_TYPE_REGISTER) &&
	   	            (context->instruction->operands[1].visibility == ZYDIS_OPERAND_VISIBILITY_IMPLICIT) &&
	   	            (context->instruction->operands[1].reg.value == ZYDIS_REGISTER_CL))
	   	        {
	   	            return context->instruction->operands[0].size;
	   	        }
	   	        break;
	   	    case 1:
	   	    case 2:
	   	        if (context->instruction->operands[operand->id - 1].size !=
	   	            context->instruction->operands[operand->id].size)
	   	        {
	   	            return context->instruction->operands[operand->id].size;
	   	        }
	   	        break;
	   	    default:
	   	        break;
	   	    }
	   	    return 0;
	   	}
	*/
	return true
}

