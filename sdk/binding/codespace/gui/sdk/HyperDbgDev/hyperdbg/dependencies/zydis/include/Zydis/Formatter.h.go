package Zydis
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Formatter.h.back

const(
ZYDIS_FORMATTER_H =  //col:1
ZYDIS_RUNTIME_ADDRESS_NONE = (ZyanU64)(-1) //col:2
)

const(
    ZYDIS_FORMATTER_STYLE_ATT = 1  //col:3
    ZYDIS_FORMATTER_STYLE_INTEL = 2  //col:4
    ZYDIS_FORMATTER_STYLE_INTEL_MASM = 3  //col:5
    ZYDIS_FORMATTER_STYLE_MAX_VALUE  =  ZYDIS_FORMATTER_STYLE_INTEL_MASM  //col:6
    ZYDIS_FORMATTER_STYLE_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_FORMATTER_STYLE_MAX_VALUE)  //col:7
)


const(
    ZYDIS_FORMATTER_PROP_FORCE_SIZE = 1  //col:11
    ZYDIS_FORMATTER_PROP_FORCE_SEGMENT = 2  //col:12
    ZYDIS_FORMATTER_PROP_FORCE_RELATIVE_BRANCHES = 3  //col:13
    ZYDIS_FORMATTER_PROP_FORCE_RELATIVE_RIPREL = 4  //col:14
    ZYDIS_FORMATTER_PROP_PRINT_BRANCH_SIZE = 5  //col:15
    ZYDIS_FORMATTER_PROP_DETAILED_PREFIXES = 6  //col:16
    ZYDIS_FORMATTER_PROP_ADDR_BASE = 7  //col:17
    ZYDIS_FORMATTER_PROP_ADDR_SIGNEDNESS = 8  //col:18
    ZYDIS_FORMATTER_PROP_ADDR_PADDING_ABSOLUTE = 9  //col:19
    ZYDIS_FORMATTER_PROP_ADDR_PADDING_RELATIVE = 10  //col:20
    ZYDIS_FORMATTER_PROP_DISP_BASE = 11  //col:21
    ZYDIS_FORMATTER_PROP_DISP_SIGNEDNESS = 12  //col:22
    ZYDIS_FORMATTER_PROP_DISP_PADDING = 13  //col:23
    ZYDIS_FORMATTER_PROP_IMM_BASE = 14  //col:24
    ZYDIS_FORMATTER_PROP_IMM_SIGNEDNESS = 15  //col:25
    ZYDIS_FORMATTER_PROP_IMM_PADDING = 16  //col:26
    ZYDIS_FORMATTER_PROP_UPPERCASE_PREFIXES = 17  //col:27
    ZYDIS_FORMATTER_PROP_UPPERCASE_MNEMONIC = 18  //col:28
    ZYDIS_FORMATTER_PROP_UPPERCASE_REGISTERS = 19  //col:29
    ZYDIS_FORMATTER_PROP_UPPERCASE_TYPECASTS = 20  //col:30
    ZYDIS_FORMATTER_PROP_UPPERCASE_DECORATORS = 21  //col:31
    ZYDIS_FORMATTER_PROP_DEC_PREFIX = 22  //col:32
    ZYDIS_FORMATTER_PROP_DEC_SUFFIX = 23  //col:33
    ZYDIS_FORMATTER_PROP_HEX_UPPERCASE = 24  //col:34
    ZYDIS_FORMATTER_PROP_HEX_PREFIX = 25  //col:35
    ZYDIS_FORMATTER_PROP_HEX_SUFFIX = 26  //col:36
    ZYDIS_FORMATTER_PROP_MAX_VALUE  =  ZYDIS_FORMATTER_PROP_HEX_SUFFIX  //col:37
    ZYDIS_FORMATTER_PROP_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_FORMATTER_PROP_MAX_VALUE)  //col:38
)


const(
    ZYDIS_NUMERIC_BASE_DEC = 1  //col:42
    ZYDIS_NUMERIC_BASE_HEX = 2  //col:43
    ZYDIS_NUMERIC_BASE_MAX_VALUE  =  ZYDIS_NUMERIC_BASE_HEX  //col:44
    ZYDIS_NUMERIC_BASE_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_NUMERIC_BASE_MAX_VALUE)  //col:45
)


const(
    ZYDIS_SIGNEDNESS_AUTO = 1  //col:49
    ZYDIS_SIGNEDNESS_SIGNED = 2  //col:50
    ZYDIS_SIGNEDNESS_UNSIGNED = 3  //col:51
    ZYDIS_SIGNEDNESS_MAX_VALUE  =  ZYDIS_SIGNEDNESS_UNSIGNED  //col:52
    ZYDIS_SIGNEDNESS_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_SIGNEDNESS_MAX_VALUE)  //col:53
)


const(
    ZYDIS_PADDING_DISABLED  =  0  //col:57
    ZYDIS_PADDING_AUTO      =  (-1)  //col:58
    ZYDIS_PADDING_MAX_VALUE  =  ZYDIS_PADDING_AUTO  //col:59
    ZYDIS_PADDING_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_PADDING_MAX_VALUE)  //col:60
)


const(
    ZYDIS_FORMATTER_FUNC_PRE_INSTRUCTION = 1  //col:64
    ZYDIS_FORMATTER_FUNC_POST_INSTRUCTION = 2  //col:65
    ZYDIS_FORMATTER_FUNC_FORMAT_INSTRUCTION = 3  //col:66
    ZYDIS_FORMATTER_FUNC_PRE_OPERAND = 4  //col:67
    ZYDIS_FORMATTER_FUNC_POST_OPERAND = 5  //col:68
    ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_REG = 6  //col:69
    ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_MEM = 7  //col:70
    ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_PTR = 8  //col:71
    ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_IMM = 9  //col:72
    ZYDIS_FORMATTER_FUNC_PRINT_MNEMONIC = 10  //col:73
    ZYDIS_FORMATTER_FUNC_PRINT_REGISTER = 11  //col:74
    ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS = 12  //col:75
    ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_REL = 13  //col:76
    ZYDIS_FORMATTER_FUNC_PRINT_DISP = 14  //col:77
    ZYDIS_FORMATTER_FUNC_PRINT_IMM = 15  //col:78
    ZYDIS_FORMATTER_FUNC_PRINT_TYPECAST = 16  //col:79
    ZYDIS_FORMATTER_FUNC_PRINT_SEGMENT = 17  //col:80
    ZYDIS_FORMATTER_FUNC_PRINT_PREFIXES = 18  //col:81
    ZYDIS_FORMATTER_FUNC_PRINT_DECORATOR = 19  //col:82
    ZYDIS_FORMATTER_FUNC_MAX_VALUE  =  ZYDIS_FORMATTER_FUNC_PRINT_DECORATOR  //col:83
    ZYDIS_FORMATTER_FUNC_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_FORMATTER_FUNC_MAX_VALUE)  //col:84
)


const(
    ZYDIS_DECORATOR_INVALID = 1  //col:88
    ZYDIS_DECORATOR_MASK = 2  //col:89
    ZYDIS_DECORATOR_BC = 3  //col:90
    ZYDIS_DECORATOR_RC = 4  //col:91
    ZYDIS_DECORATOR_SAE = 5  //col:92
    ZYDIS_DECORATOR_SWIZZLE = 6  //col:93
    ZYDIS_DECORATOR_CONVERSION = 7  //col:94
    ZYDIS_DECORATOR_EH = 8  //col:95
    ZYDIS_DECORATOR_MAX_VALUE  =  ZYDIS_DECORATOR_EH  //col:96
    ZYDIS_DECORATOR_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_DECORATOR_MAX_VALUE)  //col:97
)



type typedef struct ZydisFormatterContext_ struct{
ZydisDecodedInstruction* const
ZydisDecodedOperand* const
runtime_address ZyanU64
user_data void*
}


type typedef struct ZydisFormatterContext_ struct{
ZydisDecodedInstruction* const
ZydisDecodedOperand* const
runtime_address ZyanU64
user_data void*
}



type (
Formatter interface{
typedef_ZyanStatus_()(ok bool)//col:57
ZYDIS_EXPORT_ZyanStatus_ZydisFormatterInit()(ok bool)//col:88
}
)

func NewFormatter() { return & formatter{} }

func (f *formatter)typedef_ZyanStatus_()(ok bool){//col:57
/*typedef ZyanStatus (*ZydisFormatterFunc)(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
typedef ZyanStatus (*ZydisFormatterRegisterFunc)(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context, ZydisRegister reg);
typedef ZyanStatus (*ZydisFormatterDecoratorFunc)(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context, ZydisDecorator decorator);
struct ZydisFormatter_
{
    ZydisFormatterStyle style;
    ZyanBool force_memory_size;
    ZyanBool force_memory_segment;
    ZyanBool force_relative_branches;
    ZyanBool force_relative_riprel;
    ZyanBool print_branch_size;
    ZyanBool detailed_prefixes;
    ZydisNumericBase addr_base;
    ZydisSignedness addr_signedness;
    ZydisPadding addr_padding_absolute;
    ZydisPadding addr_padding_relative;
    ZydisNumericBase disp_base;
    ZydisSignedness disp_signedness;
    ZydisPadding disp_padding;
    ZydisNumericBase imm_base;
    ZydisSignedness imm_signedness;
    ZydisPadding imm_padding;
    ZyanI32 case_prefixes;
    ZyanI32 case_mnemonic;
    ZyanI32 case_registers;
    ZyanI32 case_typecasts;
    ZyanI32 case_decorators;
    ZyanBool hex_uppercase;
    struct
    {
        const ZyanStringView* string;
        ZyanStringView string_data;
        char buffer[11];
    } number_format[ZYDIS_NUMERIC_BASE_MAX_VALUE + 1][2];
    ZydisFormatterFunc func_pre_instruction;
    ZydisFormatterFunc func_post_instruction;
    ZydisFormatterFunc func_format_instruction;
    ZydisFormatterFunc func_pre_operand;
    ZydisFormatterFunc func_post_operand;
    ZydisFormatterFunc func_format_operand_reg;
    ZydisFormatterFunc func_format_operand_mem;
    ZydisFormatterFunc func_format_operand_ptr;
    ZydisFormatterFunc func_format_operand_imm;
    ZydisFormatterFunc func_print_mnemonic;
    ZydisFormatterRegisterFunc func_print_register;
    ZydisFormatterFunc func_print_address_abs;
    ZydisFormatterFunc func_print_address_rel;
    ZydisFormatterFunc func_print_disp;
    ZydisFormatterFunc func_print_imm;
    ZydisFormatterFunc func_print_typecast;
    ZydisFormatterFunc func_print_segment;
    ZydisFormatterFunc func_print_prefixes;
    ZydisFormatterDecoratorFunc func_print_decorator;
};*/
return true
}

func (f *formatter)ZYDIS_EXPORT_ZyanStatus_ZydisFormatterInit()(ok bool){//col:88
/*ZYDIS_EXPORT ZyanStatus ZydisFormatterInit(ZydisFormatter* formatter, ZydisFormatterStyle style);
ZYDIS_EXPORT ZyanStatus ZydisFormatterSetProperty(ZydisFormatter* formatter,
    ZydisFormatterProperty property, ZyanUPointer value);
ZYDIS_EXPORT ZyanStatus ZydisFormatterSetHook(ZydisFormatter* formatter,
    ZydisFormatterFunction type, const void** callback);
ZYDIS_EXPORT ZyanStatus ZydisFormatterFormatInstruction(const ZydisFormatter* formatter,
    const ZydisDecodedInstruction* instruction, char* buffer, ZyanUSize length,
    ZyanU64 runtime_address);
ZYDIS_EXPORT ZyanStatus ZydisFormatterFormatInstructionEx(const ZydisFormatter* formatter,
    const ZydisDecodedInstruction* instruction, char* buffer, ZyanUSize length,
    ZyanU64 runtime_address, void* user_data);
ZYDIS_EXPORT ZyanStatus ZydisFormatterFormatOperand(const ZydisFormatter* formatter,
    const ZydisDecodedInstruction* instruction, ZyanU8 index, char* buffer, ZyanUSize length,
    ZyanU64 runtime_address);
ZYDIS_EXPORT ZyanStatus ZydisFormatterFormatOperandEx(const ZydisFormatter* formatter,
    const ZydisDecodedInstruction* instruction, ZyanU8 index, char* buffer, ZyanUSize length,
    ZyanU64 runtime_address, void* user_data);
ZYDIS_EXPORT ZyanStatus ZydisFormatterTokenizeInstruction(const ZydisFormatter* formatter,
    const ZydisDecodedInstruction* instruction, void* buffer, ZyanUSize length,
    ZyanU64 runtime_address, ZydisFormatterTokenConst** token);
ZYDIS_EXPORT ZyanStatus ZydisFormatterTokenizeInstructionEx(const ZydisFormatter* formatter,
    const ZydisDecodedInstruction* instruction, void* buffer, ZyanUSize length,
    ZyanU64 runtime_address, ZydisFormatterTokenConst** token, void* user_data);
ZYDIS_EXPORT ZyanStatus ZydisFormatterTokenizeOperand(const ZydisFormatter* formatter,
    const ZydisDecodedInstruction* instruction, ZyanU8 index, void* buffer, ZyanUSize length,
    ZyanU64 runtime_address, ZydisFormatterTokenConst** token);
ZYDIS_EXPORT ZyanStatus ZydisFormatterTokenizeOperandEx(const ZydisFormatter* formatter,
    const ZydisDecodedInstruction* instruction, ZyanU8 index, void* buffer, ZyanUSize length,
    ZyanU64 runtime_address, ZydisFormatterTokenConst** token, void* user_data);
#ifdef __cplusplus
}*/
return true
}



