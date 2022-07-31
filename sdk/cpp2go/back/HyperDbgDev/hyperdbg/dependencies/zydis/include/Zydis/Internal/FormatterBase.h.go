package Internal


const(
ZYDIS_FORMATTER_BASE_H =  //col:33
ZYDIS_STRING_APPEND_NUM_U(formatter, base, str, value, padding_length) = switch (base) { case ZYDIS_NUMERIC_BASE_DEC: ZYAN_CHECK(ZydisStringAppendDecU(str, value, padding_length, (formatter)->number_format[base][0].string, (formatter)->number_format[base][1].string)); break; case ZYDIS_NUMERIC_BASE_HEX: ZYAN_CHECK(ZydisStringAppendHexU(str, value, padding_length, (formatter)->hex_uppercase, (formatter)->number_format[base][0].string, (formatter)->number_format[base][1].string)); break; default: return ZYAN_STATUS_INVALID_ARGUMENT; } //col:59
ZYDIS_STRING_APPEND_NUM_S(formatter, base, str, value, padding_length, force_sign) = switch (base) { case ZYDIS_NUMERIC_BASE_DEC: ZYAN_CHECK(ZydisStringAppendDecS(str, value, padding_length, force_sign, (formatter)->number_format[base][0].string, (formatter)->number_format[base][1].string)); break; case ZYDIS_NUMERIC_BASE_HEX: ZYAN_CHECK(ZydisStringAppendHexS(str, value, padding_length, (formatter)->hex_uppercase, force_sign, (formatter)->number_format[base][0].string, (formatter)->number_format[base][1].string)); break; default: return ZYAN_STATUS_INVALID_ARGUMENT; } //col:87
ZYDIS_BUFFER_APPEND_TOKEN(buffer, type) = if ((buffer)->is_token_list) { ZYAN_CHECK(ZydisFormatterBufferAppend(buffer, type)); } //col:119
ZYDIS_BUFFER_REMEMBER(buffer, state) = if ((buffer)->is_token_list) { (state) = (ZyanUPointer)(buffer)->string.vector.data; } else { (state) = (ZyanUPointer)(buffer)->string.vector.size; } //col:134
ZYDIS_BUFFER_APPEND(buffer, name) = if ((buffer)->is_token_list) { ZYAN_CHECK(ZydisFormatterBufferAppendPredefined(buffer, TOK_ ## name)); } else { ZYAN_CHECK(ZydisStringAppendShort(&buffer->string, &STR_ ## name)); } //col:149
ZYDIS_BUFFER_APPEND_CASE(buffer, name, letter_case) = if ((buffer)->is_token_list) { ZYAN_CHECK(ZydisFormatterBufferAppendPredefined(buffer, TOK_ ## name)); } else { ZYAN_CHECK(ZydisStringAppendShortCase(&buffer->string, &STR_ ## name, letter_case)); } //col:167
)

type (
FormatterBase interface{
    switch ()(ok bool)//col:199
#pragma pack()(ok bool)//col:243
ZyanU32 ZydisFormatterHelperGetExplicitSize()(ok bool)//col:316
}






)

func NewFormatterBase() { return & formatterBase{} }

func (f *formatterBase)    switch ()(ok bool){//col:199









































































return true
}







func (f *formatterBase)#pragma pack()(ok bool){//col:243

























return true
}







func (f *formatterBase)ZyanU32 ZydisFormatterHelperGetExplicitSize()(ok bool){//col:316






















return true
}









