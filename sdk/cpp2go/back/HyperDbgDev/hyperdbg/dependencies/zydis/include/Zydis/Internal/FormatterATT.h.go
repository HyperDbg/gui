package Internal
//back\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Internal\FormatterATT.h.back

const(
ZYDIS_FORMATTER_ATT_H =  //col:1
)

type (
FormatterAtt interface{
ZyanStatus_ZydisFormatterATTFormatInstruction()(ok bool)//col:84
}
)

func NewFormatterAtt() { return & formatterAtt{} }

func (f *formatterAtt)ZyanStatus_ZydisFormatterATTFormatInstruction()(ok bool){//col:84
/*ZyanStatus ZydisFormatterATTFormatInstruction(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterATTFormatOperandMEM(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterATTPrintMnemonic(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterATTPrintRegister(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context, ZydisRegister reg);
ZyanStatus ZydisFormatterATTPrintDISP(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterATTPrintIMM(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
static const ZydisFormatter FORMATTER_ATT =
{
    ZYDIS_FORMATTER_STYLE_ATT,
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
    ZYDIS_SIGNEDNESS_AUTO,
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
                &FORMATTER_ATT.number_format[
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
    &ZydisFormatterATTFormatInstruction,
    ZYAN_NULL,
    ZYAN_NULL,
    &ZydisFormatterBaseFormatOperandREG,
    &ZydisFormatterATTFormatOperandMEM,
    &ZydisFormatterBaseFormatOperandPTR,
    &ZydisFormatterBaseFormatOperandIMM,
    &ZydisFormatterATTPrintMnemonic,
    &ZydisFormatterATTPrintRegister,
    &ZydisFormatterBasePrintAddressABS,
    &ZydisFormatterBasePrintAddressREL,
    &ZydisFormatterATTPrintDISP,
    &ZydisFormatterATTPrintIMM,
    ZYAN_NULL,
    &ZydisFormatterBasePrintSegment,
    &ZydisFormatterBasePrintPrefixes,
    &ZydisFormatterBasePrintDecorator
};*/
return true
}



