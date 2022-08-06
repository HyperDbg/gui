






























#ifndef ZYDIS_FORMATTER_H
#define ZYDIS_FORMATTER_H

#include <Zycore/Defines.h>
#include <Zycore/String.h>
#include <Zycore/Types.h>
#include <Zydis/DecoderTypes.h>
#include <Zydis/FormatterBuffer.h>

#ifdef __cplusplus
extern "C" {
#endif









#define ZYDIS_RUNTIME_ADDRESS_NONE (ZyanU64)(-1)












typedef enum ZydisFormatterStyle_
{
    


    ZYDIS_FORMATTER_STYLE_ATT,
    


    ZYDIS_FORMATTER_STYLE_INTEL,
    





    ZYDIS_FORMATTER_STYLE_INTEL_MASM,

    


    ZYDIS_FORMATTER_STYLE_MAX_VALUE = ZYDIS_FORMATTER_STYLE_INTEL_MASM,
    


    ZYDIS_FORMATTER_STYLE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_FORMATTER_STYLE_MAX_VALUE)
} ZydisFormatterStyle;








typedef enum ZydisFormatterProperty_
{
    
    
    

    






    ZYDIS_FORMATTER_PROP_FORCE_SIZE,
    





    ZYDIS_FORMATTER_PROP_FORCE_SEGMENT,
    






    ZYDIS_FORMATTER_PROP_FORCE_RELATIVE_BRANCHES,
    






    ZYDIS_FORMATTER_PROP_FORCE_RELATIVE_RIPREL,
    







    ZYDIS_FORMATTER_PROP_PRINT_BRANCH_SIZE,

    





    ZYDIS_FORMATTER_PROP_DETAILED_PREFIXES,

    
    
    

    


    ZYDIS_FORMATTER_PROP_ADDR_BASE,
    



    ZYDIS_FORMATTER_PROP_ADDR_SIGNEDNESS,
    






    ZYDIS_FORMATTER_PROP_ADDR_PADDING_ABSOLUTE,
    






    ZYDIS_FORMATTER_PROP_ADDR_PADDING_RELATIVE,

    

    


    ZYDIS_FORMATTER_PROP_DISP_BASE,
    


    ZYDIS_FORMATTER_PROP_DISP_SIGNEDNESS,
    





    ZYDIS_FORMATTER_PROP_DISP_PADDING,

    

    


    ZYDIS_FORMATTER_PROP_IMM_BASE,
    





    ZYDIS_FORMATTER_PROP_IMM_SIGNEDNESS,
    






    ZYDIS_FORMATTER_PROP_IMM_PADDING,

    
    
    

    




    ZYDIS_FORMATTER_PROP_UPPERCASE_PREFIXES,
    




    ZYDIS_FORMATTER_PROP_UPPERCASE_MNEMONIC,
    




    ZYDIS_FORMATTER_PROP_UPPERCASE_REGISTERS,
    




    ZYDIS_FORMATTER_PROP_UPPERCASE_TYPECASTS,
    




    ZYDIS_FORMATTER_PROP_UPPERCASE_DECORATORS,

    
    
    

    







    ZYDIS_FORMATTER_PROP_DEC_PREFIX,
    







    ZYDIS_FORMATTER_PROP_DEC_SUFFIX,

    

    






    ZYDIS_FORMATTER_PROP_HEX_UPPERCASE,
    







    ZYDIS_FORMATTER_PROP_HEX_PREFIX,
    







    ZYDIS_FORMATTER_PROP_HEX_SUFFIX,

    

    


    ZYDIS_FORMATTER_PROP_MAX_VALUE = ZYDIS_FORMATTER_PROP_HEX_SUFFIX,
    


    ZYDIS_FORMATTER_PROP_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_FORMATTER_PROP_MAX_VALUE)
} ZydisFormatterProperty;






typedef enum ZydisNumericBase_
{
    


    ZYDIS_NUMERIC_BASE_DEC,
    


    ZYDIS_NUMERIC_BASE_HEX,

    


    ZYDIS_NUMERIC_BASE_MAX_VALUE = ZYDIS_NUMERIC_BASE_HEX,
    


    ZYDIS_NUMERIC_BASE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_NUMERIC_BASE_MAX_VALUE)
} ZydisNumericBase;






typedef enum ZydisSignedness_
{
    



    ZYDIS_SIGNEDNESS_AUTO,
    


    ZYDIS_SIGNEDNESS_SIGNED,
    


    ZYDIS_SIGNEDNESS_UNSIGNED,

    


    ZYDIS_SIGNEDNESS_MAX_VALUE = ZYDIS_SIGNEDNESS_UNSIGNED,
    


    ZYDIS_SIGNEDNESS_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_SIGNEDNESS_MAX_VALUE)
} ZydisSignedness;






typedef enum ZydisPadding_
{
    


    ZYDIS_PADDING_DISABLED = 0,
    



    ZYDIS_PADDING_AUTO     = (-1),

    


    ZYDIS_PADDING_MAX_VALUE = ZYDIS_PADDING_AUTO,
    


    ZYDIS_PADDING_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_PADDING_MAX_VALUE)
} ZydisPadding;











typedef enum ZydisFormatterFunction_
{
    
    
    

    


    ZYDIS_FORMATTER_FUNC_PRE_INSTRUCTION,
    


    ZYDIS_FORMATTER_FUNC_POST_INSTRUCTION,

    

    






    ZYDIS_FORMATTER_FUNC_FORMAT_INSTRUCTION,

    
    
    

    


    ZYDIS_FORMATTER_FUNC_PRE_OPERAND,
    


    ZYDIS_FORMATTER_FUNC_POST_OPERAND,

    

    


    ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_REG,
    






    ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_MEM,
    


    ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_PTR,
    






    ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_IMM,

    
    
    

    


    ZYDIS_FORMATTER_FUNC_PRINT_MNEMONIC,

    

    


    ZYDIS_FORMATTER_FUNC_PRINT_REGISTER,
    










    ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS,
    





    ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_REL,
    






    ZYDIS_FORMATTER_FUNC_PRINT_DISP,
    









    ZYDIS_FORMATTER_FUNC_PRINT_IMM,

    
    
    

    


    ZYDIS_FORMATTER_FUNC_PRINT_TYPECAST,
    


    ZYDIS_FORMATTER_FUNC_PRINT_SEGMENT,
    


    ZYDIS_FORMATTER_FUNC_PRINT_PREFIXES,
    



    ZYDIS_FORMATTER_FUNC_PRINT_DECORATOR,

    

    


    ZYDIS_FORMATTER_FUNC_MAX_VALUE = ZYDIS_FORMATTER_FUNC_PRINT_DECORATOR,
    


    ZYDIS_FORMATTER_FUNC_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_FORMATTER_FUNC_MAX_VALUE)
} ZydisFormatterFunction;








typedef enum ZydisDecorator_
{
    ZYDIS_DECORATOR_INVALID,
    


    ZYDIS_DECORATOR_MASK,
    


    ZYDIS_DECORATOR_BC,
    


    ZYDIS_DECORATOR_RC,
    


    ZYDIS_DECORATOR_SAE,
    


    ZYDIS_DECORATOR_SWIZZLE,
    


    ZYDIS_DECORATOR_CONVERSION,
    


    ZYDIS_DECORATOR_EH,

    


    ZYDIS_DECORATOR_MAX_VALUE = ZYDIS_DECORATOR_EH,
    


    ZYDIS_DECORATOR_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_DECORATOR_MAX_VALUE)
} ZydisDecorator;





typedef struct ZydisFormatter_ ZydisFormatter;




typedef struct ZydisFormatterContext_
{
    


    const ZydisDecodedInstruction* instruction;
    


    const ZydisDecodedOperand* operand;
    


    ZyanU64 runtime_address;
    


    void* user_data;
} ZydisFormatterContext;













































typedef ZyanStatus (*ZydisFormatterFunc)(const ZydisFormatter* formatter,
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
};

























ZYDIS_EXPORT ZyanStatus ZydisFormatterInit(ZydisFormatter* formatter, ZydisFormatterStyle style);

















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
}
#endif

#endif 
