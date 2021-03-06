#ifndef ZYCORE_FORMAT_H
#define ZYCORE_FORMAT_H
#include <ZycoreExportConfig.h>
#include <Zycore/Status.h>
#include <Zycore/String.h>
#include <Zycore/Types.h>
#ifdef __cplusplus
extern "C" {
#endif
ZYAN_PRINTF_ATTR(3, 4)
ZYCORE_EXPORT ZyanStatus
ZyanStringInsertFormat(ZyanString * string, ZyanUSize index, const char * format, ...);
ZYCORE_EXPORT ZyanStatus
ZyanStringInsertDecU(ZyanString * string, ZyanUSize index, ZyanU64 value, ZyanU8 padding_length);
ZYCORE_EXPORT ZyanStatus
ZyanStringInsertDecS(ZyanString * string, ZyanUSize index, ZyanI64 value, ZyanU8 padding_length, ZyanBool force_sign, const ZyanString * prefix);
ZYCORE_EXPORT ZyanStatus
ZyanStringInsertHexU(ZyanString * string, ZyanUSize index, ZyanU64 value, ZyanU8 padding_length, ZyanBool uppercase);
ZYCORE_EXPORT ZyanStatus
ZyanStringInsertHexS(ZyanString * string, ZyanUSize index, ZyanI64 value, ZyanU8 padding_length, ZyanBool uppercase, ZyanBool force_sign, const ZyanString * prefix);
#ifndef ZYAN_NO_LIBC
ZYAN_PRINTF_ATTR(2, 3)
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanStatus
ZyanStringAppendFormat(
    ZyanString * string,
    const char * format,
    ...);
#endif // ZYAN_NO_LIBC
ZYCORE_EXPORT ZyanStatus
ZyanStringAppendDecU(ZyanString * string, ZyanU64 value, ZyanU8 padding_length);
ZYCORE_EXPORT ZyanStatus
ZyanStringAppendDecS(ZyanString * string, ZyanI64 value, ZyanU8 padding_length, ZyanBool force_sign, const ZyanStringView * prefix);
ZYCORE_EXPORT ZyanStatus
ZyanStringAppendHexU(ZyanString * string, ZyanU64 value, ZyanU8 padding_length, ZyanBool uppercase);
ZYCORE_EXPORT ZyanStatus
ZyanStringAppendHexS(ZyanString * string, ZyanI64 value, ZyanU8 padding_length, ZyanBool uppercase, ZyanBool force_sign, const ZyanStringView * prefix);
#ifdef __cplusplus
}

#endif
#endif // ZYCORE_FORMAT_H
