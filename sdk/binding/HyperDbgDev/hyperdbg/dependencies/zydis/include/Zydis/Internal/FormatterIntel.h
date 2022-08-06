






























#ifndef ZYDIS_FORMATTER_INTEL_H
#define ZYDIS_FORMATTER_INTEL_H

#include <Zydis/Formatter.h>
#include <Zydis/Internal/FormatterBase.h>
#include <Zydis/Internal/String.h>

#ifdef __cplusplus
extern "C" {
#endif









ZyanStatus ZydisFormatterIntelFormatInstruction(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);

ZyanStatus ZydisFormatterIntelFormatOperandMEM(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);

ZyanStatus ZydisFormatterIntelPrintMnemonic(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);

ZyanStatus ZydisFormatterIntelPrintRegister(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context, ZydisRegister reg);

ZyanStatus ZydisFormatterIntelPrintDISP(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);

ZyanStatus ZydisFormatterIntelPrintTypecast(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);





ZyanStatus ZydisFormatterIntelFormatInstructionMASM(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);

ZyanStatus ZydisFormatterIntelPrintAddressMASM(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);














static const ZydisFormatter FORMATTER_INTEL =
{
     ZYDIS_FORMATTER_STYLE_INTEL,
     ZYAN_FALSE,
     ZYAN_FALSE,
     ZYAN_FALSE,
     ZYAN_FALSE,
     ZYAN_FALSE,
     ZYAN_FALSE,
     ZYDIS_NUMERIC_BASE_HEX,
     ZYDIS_SIGNEDNESS_SIGNED,
     ZYDIS_PADDING_AUTO,
     2,
     ZYDIS_NUMERIC_BASE_HEX,
     ZYDIS_SIGNEDNESS_SIGNED,
     2,
     ZYDIS_NUMERIC_BASE_HEX,
     ZYDIS_SIGNEDNESS_UNSIGNED,
     2,
     ZYDIS_LETTER_CASE_DEFAULT,
     ZYDIS_LETTER_CASE_DEFAULT,
     ZYDIS_LETTER_CASE_DEFAULT,
     ZYDIS_LETTER_CASE_DEFAULT,
     ZYDIS_LETTER_CASE_DEFAULT,
     ZYAN_TRUE,
    
    {
        
        {
            
            {
                 ZYAN_NULL,
                 ZYAN_DEFINE_STRING_VIEW(""),
                 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
            },
            
            {
                 ZYAN_NULL,
                 ZYAN_DEFINE_STRING_VIEW(""),
                 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
            }
        },
        
        {
            
            {
                 &FORMATTER_INTEL.number_format[
                                      ZYDIS_NUMERIC_BASE_HEX][0].string_data,
                 ZYAN_DEFINE_STRING_VIEW("0x"),
                 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
            },
            
            {
                 ZYAN_NULL,
                 ZYAN_DEFINE_STRING_VIEW(""),
                 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
            }
        }
    },
     ZYAN_NULL,
     ZYAN_NULL,
     &ZydisFormatterIntelFormatInstruction,
     ZYAN_NULL,
     ZYAN_NULL,
     &ZydisFormatterBaseFormatOperandREG,
     &ZydisFormatterIntelFormatOperandMEM,
     &ZydisFormatterBaseFormatOperandPTR,
     &ZydisFormatterBaseFormatOperandIMM,
     &ZydisFormatterIntelPrintMnemonic,
     &ZydisFormatterIntelPrintRegister,
     &ZydisFormatterBasePrintAddressABS,
     &ZydisFormatterBasePrintAddressREL,
     &ZydisFormatterIntelPrintDISP,
     &ZydisFormatterBasePrintIMM,
     &ZydisFormatterIntelPrintTypecast,
     &ZydisFormatterBasePrintSegment,
     &ZydisFormatterBasePrintPrefixes,
     &ZydisFormatterBasePrintDecorator
};








static const ZydisFormatter FORMATTER_INTEL_MASM =
{
     ZYDIS_FORMATTER_STYLE_INTEL_MASM,
     ZYAN_TRUE,
     ZYAN_FALSE,
     ZYAN_FALSE,
     ZYAN_FALSE,
     ZYAN_FALSE,
     ZYAN_FALSE,
     ZYDIS_NUMERIC_BASE_HEX,
     ZYDIS_SIGNEDNESS_SIGNED,
     ZYDIS_PADDING_DISABLED,
     ZYDIS_PADDING_DISABLED,
     ZYDIS_NUMERIC_BASE_HEX,
     ZYDIS_SIGNEDNESS_SIGNED,
     ZYDIS_PADDING_DISABLED,
     ZYDIS_NUMERIC_BASE_HEX,
     ZYDIS_SIGNEDNESS_AUTO,
     ZYDIS_PADDING_DISABLED,
     ZYDIS_LETTER_CASE_DEFAULT,
     ZYDIS_LETTER_CASE_DEFAULT,
     ZYDIS_LETTER_CASE_DEFAULT,
     ZYDIS_LETTER_CASE_DEFAULT,
     ZYDIS_LETTER_CASE_DEFAULT,
     ZYAN_TRUE,
    
    {
        
        {
            
            {
                 ZYAN_NULL,
                 ZYAN_DEFINE_STRING_VIEW(""),
                 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
            },
            
            {
                 ZYAN_NULL,
                 ZYAN_DEFINE_STRING_VIEW(""),
                 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
            }
        },
        
        {
            
            {
                 ZYAN_NULL,
                 ZYAN_DEFINE_STRING_VIEW(""),
                 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
            },
            
            {
                 &FORMATTER_INTEL_MASM.number_format[
                                      ZYDIS_NUMERIC_BASE_HEX][1].string_data,
                 ZYAN_DEFINE_STRING_VIEW("h"),
                 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
            }
        }
    },
     ZYAN_NULL,
     ZYAN_NULL,
     &ZydisFormatterIntelFormatInstructionMASM,
     ZYAN_NULL,
     ZYAN_NULL,
     &ZydisFormatterBaseFormatOperandREG,
     &ZydisFormatterIntelFormatOperandMEM,
     &ZydisFormatterBaseFormatOperandPTR,
     &ZydisFormatterBaseFormatOperandIMM,
     &ZydisFormatterIntelPrintMnemonic,
     &ZydisFormatterIntelPrintRegister,
     &ZydisFormatterIntelPrintAddressMASM,
     &ZydisFormatterIntelPrintAddressMASM,
     &ZydisFormatterIntelPrintDISP,
     &ZydisFormatterBasePrintIMM,
     &ZydisFormatterIntelPrintTypecast,
     &ZydisFormatterBasePrintSegment,
     &ZydisFormatterBasePrintPrefixes,
     &ZydisFormatterBasePrintDecorator
};





#ifdef __cplusplus
}
#endif

#endif 
