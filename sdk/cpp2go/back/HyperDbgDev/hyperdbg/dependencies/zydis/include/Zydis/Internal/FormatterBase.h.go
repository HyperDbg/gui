package Internal
//back\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Internal\FormatterBase.h.back

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
  Zyan Disassembler Library ()(ok bool)//col:199
#pragma pack()(ok bool)//col:243
 * Returns the size to be used as explicit size suffix ()(ok bool)//col:315
}
)

func NewFormatterBase() { return & formatterBase{} }

func (f *formatterBase)  Zyan Disassembler Library ()(ok bool){//col:199
/*  Zyan Disassembler Library (Zydis)
  Original Author : Florian Bernd, Joel Hoener
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 * Provides formatter functions that are shared between the different formatters.
#ifndef ZYDIS_FORMATTER_BASE_H
#define ZYDIS_FORMATTER_BASE_H
#include <Zydis/Formatter.h>
#include <Zydis/Internal/String.h>
#ifdef __cplusplus
extern "C" {
#endif
 * Appends an unsigned numeric value to the given string.
 *
#define ZYDIS_STRING_APPEND_NUM_U(formatter, base, str, value, padding_length) \
    switch (base) \
    { \
    case ZYDIS_NUMERIC_BASE_DEC: \
        ZYAN_CHECK(ZydisStringAppendDecU(str, value, padding_length, \
            (formatter)->number_format[base][0].string, \
            (formatter)->number_format[base][1].string)); \
        break; \
    case ZYDIS_NUMERIC_BASE_HEX: \
        ZYAN_CHECK(ZydisStringAppendHexU(str, value, padding_length, \
            (formatter)->hex_uppercase, \
            (formatter)->number_format[base][0].string, \
            (formatter)->number_format[base][1].string)); \
        break; \
    default: \
        return ZYAN_STATUS_INVALID_ARGUMENT; \
    }
 * Appends a signed numeric value to the given string.
 *
#define ZYDIS_STRING_APPEND_NUM_S(formatter, base, str, value, padding_length, force_sign) \
    switch (base) \
    { \
    case ZYDIS_NUMERIC_BASE_DEC: \
        ZYAN_CHECK(ZydisStringAppendDecS(str, value, padding_length, force_sign, \
            (formatter)->number_format[base][0].string, \
            (formatter)->number_format[base][1].string)); \
        break; \
    case ZYDIS_NUMERIC_BASE_HEX: \
        ZYAN_CHECK(ZydisStringAppendHexS(str, value, padding_length,  \
            (formatter)->hex_uppercase, force_sign, \
            (formatter)->number_format[base][0].string, \
            (formatter)->number_format[base][1].string)); \
        break; \
    default: \
        return ZYAN_STATUS_INVALID_ARGUMENT; \
    }
 * Invokes the `ZydisFormatterBufferAppend` routine, if tokenization is enabled for the
 * current pass.
 *
 *
 * Using this macro instead of direct calls to `ZydisFormatterBufferAppend` greatly improves the
 * performance for non-tokenizing passes.
#define ZYDIS_BUFFER_APPEND_TOKEN(buffer, type) \
    if ((buffer)->is_token_list) \
    { \
        ZYAN_CHECK(ZydisFormatterBufferAppend(buffer, type)); \
    }
 * Returns a snapshot of the buffer-state.
 *
 *
 * Using this macro instead of direct calls to `ZydisFormatterBufferRemember` improves the
 * performance for non-tokenizing passes.
#define ZYDIS_BUFFER_REMEMBER(buffer, state) \
    if ((buffer)->is_token_list) \
    { \
        (state) = (ZyanUPointer)(buffer)->string.vector.data; \
    } else \
    { \
        (state) = (ZyanUPointer)(buffer)->string.vector.size; \
    }
 * Appends a string (`STR_`-prefix) or a predefined token-list (`TOK_`-prefix).
 *
#define ZYDIS_BUFFER_APPEND(buffer, name) \
    if ((buffer)->is_token_list) \
    { \
        ZYAN_CHECK(ZydisFormatterBufferAppendPredefined(buffer, TOK_ ## name)); \
    } else \
    { \
        ZYAN_CHECK(ZydisStringAppendShort(&buffer->string, &STR_ ## name)); \
    }
 * Appends a string (`STR_`-prefix) or a predefined token-list (`TOK_`-prefix).
 *
#define ZYDIS_BUFFER_APPEND_CASE(buffer, name, letter_case) \
    if ((buffer)->is_token_list) \
    { \
        ZYAN_CHECK(ZydisFormatterBufferAppendPredefined(buffer, TOK_ ## name)); \
    } else \
    { \
        ZYAN_CHECK(ZydisStringAppendShortCase(&buffer->string, &STR_ ## name, letter_case)); \
    }
#ifdef ZYAN_MSVC
#   pragma warning(push)
#   pragma warning(disable:4200)
#endif
#pragma pack(push, 1)
typedef struct ZydisPredefinedToken_
{
    ZyanU8 size;
    ZyanU8 next;
    ZyanU8 data[];
} ZydisPredefinedToken;*/
return true
}

func (f *formatterBase)#pragma pack()(ok bool){//col:243
/*#pragma pack(pop)
#ifdef ZYAN_MSVC
#   pragma warning(pop)
#endif
 * Appends a predefined token-list to the `buffer`.
 *
 *
 *
 * This function is internally used to improve performance while adding static strings or multiple
 * tokens at once.
ZYAN_INLINE ZyanStatus ZydisFormatterBufferAppendPredefined(ZydisFormatterBuffer* buffer,
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

func (f *formatterBase) * Returns the size to be used as explicit size suffix ()(ok bool){//col:315
/* * Returns the size to be used as explicit size suffix (`AT&T`) or explicit typecast
 * (`INTEL`), if required.
 *
 *
 *
 * This function always returns a size different to `0`, if the `ZYDIS_FORMATTER_PROP_FORCE_SIZE`
 * is set to `ZYAN_TRUE`.
ZyanU32 ZydisFormatterHelperGetExplicitSize(const ZydisFormatter* formatter,
    ZydisFormatterContext* context, ZyanU8 memop_id);
ZyanStatus ZydisFormatterBaseFormatOperandREG(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterBaseFormatOperandPTR(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterBaseFormatOperandIMM(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterBasePrintAddressABS(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterBasePrintAddressREL(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterBasePrintIMM(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterBasePrintSegment(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterBasePrintPrefixes(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context);
ZyanStatus ZydisFormatterBasePrintDecorator(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context, ZydisDecorator decorator);
#ifdef __cplusplus
}*/
return true
}



