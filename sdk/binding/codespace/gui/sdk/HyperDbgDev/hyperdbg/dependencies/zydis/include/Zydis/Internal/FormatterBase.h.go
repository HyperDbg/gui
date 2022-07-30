package Internal

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Internal\FormatterBase.h.back

const (
	ZYDIS_FORMATTER_BASE_H = //col:1
	ZYDIS_STRING_APPEND_NUM_U(formatter, base, str, value, padding_length) = switch (
base
) { case ZYDIS_NUMERIC_BASE_DEC: ZYAN_CHECK(ZydisStringAppendDecU(str, value, padding_length, (formatter)->number_format[base][0].string, (formatter)->number_format[base][1].string)); break; case ZYDIS_NUMERIC_BASE_HEX: ZYAN_CHECK(ZydisStringAppendHexU(str, value, padding_length, (formatter)->hex_uppercase, (formatter)->number_format[base][0].string, (formatter)->number_format[base][1].string)); break; default: return ZYAN_STATUS_INVALID_ARGUMENT } //col:2
ZYDIS_STRING_APPEND_NUM_S(formatter, base, str, value, padding_length, force_sign) = switch (base) { case ZYDIS_NUMERIC_BASE_DEC: ZYAN_CHECK(ZydisStringAppendDecS(str, value, padding_length, force_sign, (formatter)->number_format[base][0].string, (formatter)->number_format[base][1].string)); break; case ZYDIS_NUMERIC_BASE_HEX: ZYAN_CHECK(ZydisStringAppendHexS(str, value, padding_length, (formatter)->hex_uppercase, force_sign, (formatter)->number_format[base][0].string, (formatter)->number_format[base][1].string)); break; default: return ZYAN_STATUS_INVALID_ARGUMENT; } //col:19
ZYDIS_BUFFER_APPEND_TOKEN(buffer, type
) = if ((buffer)->is_token_list) { ZYAN_CHECK(ZydisFormatterBufferAppend(buffer, type)); } //col:36
ZYDIS_BUFFER_REMEMBER(buffer, state) = if ((buffer)->is_token_list) { (state) = (ZyanUPointer)(buffer)->string.vector.data; } else { (state) = (ZyanUPointer)(buffer)->string.vector.size; } //col:41
ZYDIS_BUFFER_APPEND(buffer, name) = if ((buffer)->is_token_list) { ZYAN_CHECK(ZydisFormatterBufferAppendPredefined(buffer, TOK_ ## name)); } else { ZYAN_CHECK(ZydisStringAppendShort(&buffer->string, &STR_ ## name)); } //col:49
ZYDIS_BUFFER_APPEND_CASE(buffer, name, letter_case) = if ((buffer)->is_token_list) { ZYAN_CHECK(ZydisFormatterBufferAppendPredefined(buffer, TOK_ ## name)); } else { ZYAN_CHECK(ZydisStringAppendShortCase(&buffer->string, &STR_ ## name, letter_case)); } //col:57
)
type typedef struct ZydisPredefinedToken_ struct {
	size ZyanU8   //col:3
	next ZyanU8   //col:4
	data []ZyanU8 //col:5
}



	type (
	FormatterBase interface{
	ZYAN_INLINE_ZyanStatus_ZydisFormatterBufferAppendPredefined()(ok bool) //col:21
	}
	formatterBase struct{}
	)

	func NewFormatterBase()FormatterBase{ return & formatterBase{} }

	func (f *formatterBase)ZYAN_INLINE_ZyanStatus_ZydisFormatterBufferAppendPredefined()(ok bool){//col:21
	/*ZYAN_INLINE ZyanStatus ZydisFormatterBufferAppendPredefined(ZydisFormatterBuffer* buffer,
	      const ZydisPredefinedToken* data)
	  {
	      ZYAN_ASSERT(buffer);
	      ZYAN_ASSERT(data);
	      const ZyanUSize len = buffer->string.vector.size;
	      ZYAN_ASSERT((len > 0) && (len < 256));
	      if (buffer->capacity <= len + data->size)
	      {
	          return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
	      }
	      ZydisFormatterToken* const last = (ZydisFormatterToken*)buffer->string.vector.data - 1;
	      last->next = (ZyanU8)len;
	      ZYAN_MEMCPY((ZyanU8*)buffer->string.vector.data + len, &data->data[0], data->size);
	      const ZyanUSize delta = len + data->next;
	      buffer->capacity -= delta;
	      buffer->string.vector.data = (ZyanU8*)buffer->string.vector.data + delta;
	      buffer->string.vector.size = data->size - data->next;
	      buffer->string.vector.capacity = ZYAN_MIN(buffer->capacity, 255);
	      return ZYAN_STATUS_SUCCESS;
	  }*/
	return true
	}
