#ifndef ZYDIS_FORMATTER_BASE_H
#define ZYDIS_FORMATTER_BASE_H
#include <Zydis/Formatter.h>
#include <Zydis/Internal/String.h>
#ifdef __cplusplus
extern "C" {
#endif
#define ZYDIS_STRING_APPEND_NUM_U(formatter, base, str, value, padding_length)                                                                                                             \
    switch (base) {                                                                                                                                                                        \
    case ZYDIS_NUMERIC_BASE_DEC:                                                                                                                                                           \
        ZYAN_CHECK(ZydisStringAppendDecU(str, value, padding_length, (formatter)->number_format[base][0].string, (formatter)->number_format[base][1].string));                             \
        break;                                                                                                                                                                             \
    case ZYDIS_NUMERIC_BASE_HEX:                                                                                                                                                           \
        ZYAN_CHECK(ZydisStringAppendHexU(str, value, padding_length, (formatter)->hex_uppercase, (formatter)->number_format[base][0].string, (formatter)->number_format[base][1].string)); \
        break;                                                                                                                                                                             \
    default:                                                                                                                                                                               \
        return ZYAN_STATUS_INVALID_ARGUMENT;                                                                                                                                               \
    }
#define ZYDIS_STRING_APPEND_NUM_S(formatter, base, str, value, padding_length, force_sign)                                                                                                             \
    switch (base) {                                                                                                                                                                                    \
    case ZYDIS_NUMERIC_BASE_DEC:                                                                                                                                                                       \
        ZYAN_CHECK(ZydisStringAppendDecS(str, value, padding_length, force_sign, (formatter)->number_format[base][0].string, (formatter)->number_format[base][1].string));                             \
        break;                                                                                                                                                                                         \
    case ZYDIS_NUMERIC_BASE_HEX:                                                                                                                                                                       \
        ZYAN_CHECK(ZydisStringAppendHexS(str, value, padding_length, (formatter)->hex_uppercase, force_sign, (formatter)->number_format[base][0].string, (formatter)->number_format[base][1].string)); \
        break;                                                                                                                                                                                         \
    default:                                                                                                                                                                                           \
        return ZYAN_STATUS_INVALID_ARGUMENT;                                                                                                                                                           \
    }
#define ZYDIS_BUFFER_APPEND_TOKEN(buffer, type)               \
    if ((buffer)->is_token_list) {                            \
        ZYAN_CHECK(ZydisFormatterBufferAppend(buffer, type)); \
    }
#define ZYDIS_BUFFER_REMEMBER(buffer, state)                  \
    if ((buffer)->is_token_list) {                            \
        (state) = (ZyanUPointer)(buffer)->string.vector.data; \
    } else {                                                  \
        (state) = (ZyanUPointer)(buffer)->string.vector.size; \
    }
#define ZYDIS_BUFFER_APPEND(buffer, name)                                     \
    if ((buffer)->is_token_list) {                                            \
        ZYAN_CHECK(ZydisFormatterBufferAppendPredefined(buffer, TOK_##name)); \
    } else {                                                                  \
        ZYAN_CHECK(ZydisStringAppendShort(&buffer->string, &STR_##name));     \
    }
#define ZYDIS_BUFFER_APPEND_CASE(buffer, name, letter_case)                                \
    if ((buffer)->is_token_list) {                                                         \
        ZYAN_CHECK(ZydisFormatterBufferAppendPredefined(buffer, TOK_##name));              \
    } else {                                                                               \
        ZYAN_CHECK(ZydisStringAppendShortCase(&buffer->string, &STR_##name, letter_case)); \
    }
#ifdef ZYAN_MSVC
#    pragma warning(push)
#    pragma warning(disable : 4200)
#endif
#pragma pack(push, 1)
typedef struct ZydisPredefinedToken_ {
    ZyanU8 size;
    ZyanU8 next;
    ZyanU8 data[];
} ZydisPredefinedToken;
#pragma pack(pop)
#ifdef ZYAN_MSVC
#    pragma warning(pop)
#endif
ZYAN_INLINE ZyanStatus
ZydisFormatterBufferAppendPredefined(ZydisFormatterBuffer *       buffer,
                                     const ZydisPredefinedToken * data) {
    ZYAN_ASSERT(buffer);
    ZYAN_ASSERT(data);
    const ZyanUSize len = buffer->string.vector.size;
    ZYAN_ASSERT((len > 0) && (len < 256));
    if (buffer->capacity <= len + data->size) {
        return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
    }
    ZydisFormatterToken * const last = (ZydisFormatterToken *)buffer->string.vector.data - 1;
    last->next                       = (ZyanU8)len;
    ZYAN_MEMCPY((ZyanU8 *)buffer->string.vector.data + len, &data->data[0], data->size);
    const ZyanUSize delta = len + data->next;
    buffer->capacity -= delta;
    buffer->string.vector.data     = (ZyanU8 *)buffer->string.vector.data + delta;
    buffer->string.vector.size     = data->size - data->next;
    buffer->string.vector.capacity = ZYAN_MIN(buffer->capacity, 255);
    return ZYAN_STATUS_SUCCESS;
}

ZyanU32
ZydisFormatterHelperGetExplicitSize(const ZydisFormatter *  formatter,
                                    ZydisFormatterContext * context,
                                    ZyanU8                  memop_id);
ZyanStatus
ZydisFormatterBaseFormatOperandREG(const ZydisFormatter *  formatter,
                                   ZydisFormatterBuffer *  buffer,
                                   ZydisFormatterContext * context);
ZyanStatus
ZydisFormatterBaseFormatOperandPTR(const ZydisFormatter *  formatter,
                                   ZydisFormatterBuffer *  buffer,
                                   ZydisFormatterContext * context);
ZyanStatus
ZydisFormatterBaseFormatOperandIMM(const ZydisFormatter *  formatter,
                                   ZydisFormatterBuffer *  buffer,
                                   ZydisFormatterContext * context);
ZyanStatus
ZydisFormatterBasePrintAddressABS(const ZydisFormatter *  formatter,
                                  ZydisFormatterBuffer *  buffer,
                                  ZydisFormatterContext * context);
ZyanStatus
ZydisFormatterBasePrintAddressREL(const ZydisFormatter *  formatter,
                                  ZydisFormatterBuffer *  buffer,
                                  ZydisFormatterContext * context);
ZyanStatus
ZydisFormatterBasePrintIMM(const ZydisFormatter *  formatter,
                           ZydisFormatterBuffer *  buffer,
                           ZydisFormatterContext * context);
ZyanStatus
ZydisFormatterBasePrintSegment(const ZydisFormatter *  formatter,
                               ZydisFormatterBuffer *  buffer,
                               ZydisFormatterContext * context);
ZyanStatus
ZydisFormatterBasePrintPrefixes(const ZydisFormatter *  formatter,
                                ZydisFormatterBuffer *  buffer,
                                ZydisFormatterContext * context);
ZyanStatus
ZydisFormatterBasePrintDecorator(const ZydisFormatter *  formatter,
                                 ZydisFormatterBuffer *  buffer,
                                 ZydisFormatterContext * context,
                                 ZydisDecorator          decorator);
#ifdef __cplusplus
}

#endif
#endif // ZYDIS_FORMATTER_BASE_H
