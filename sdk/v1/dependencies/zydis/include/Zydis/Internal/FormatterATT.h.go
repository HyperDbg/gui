package Internal

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Internal\FormatterATT.h.back

type (
	FormatterAtt interface {
		ZyanStatus_ZydisFormatterATTFormatInstruction() (ok bool) //col:55
	}
	formatterAtt struct{}
)

func NewFormatterAtt() FormatterAtt { return &formatterAtt{} }

func (f *formatterAtt) ZyanStatus_ZydisFormatterATTFormatInstruction() (ok bool) { //col:55
	/*
	   ZyanStatus ZydisFormatterATTFormatInstruction(const ZydisFormatter* formatter,

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
	   	};
	*/
	return true
}

