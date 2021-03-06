#ifndef ZYDIS_H
#define ZYDIS_H
#include <Zycore/Defines.h>
#include <Zycore/Types.h>
#ifndef ZYDIS_DISABLE_DECODER
#    include <Zydis/Decoder.h>
#    include <Zydis/DecoderTypes.h>
#endif
#ifndef ZYDIS_DISABLE_FORMATTER
#    include <Zydis/Formatter.h>
#endif
#include <Zydis/MetaInfo.h>
#include <Zydis/Mnemonic.h>
#include <Zydis/Register.h>
#include <Zydis/SharedTypes.h>
#include <Zydis/Status.h>
#include <Zydis/Utils.h>
#ifdef __cplusplus
extern "C" {
#endif
#define ZYDIS_VERSION                (ZyanU64)0x0003000100000000
#define ZYDIS_VERSION_MAJOR(version) (ZyanU16)(((version)&0xFFFF000000000000) >> 48)
#define ZYDIS_VERSION_MINOR(version) (ZyanU16)(((version)&0x0000FFFF00000000) >> 32)
#define ZYDIS_VERSION_PATCH(version) (ZyanU16)(((version)&0x00000000FFFF0000) >> 16)
#define ZYDIS_VERSION_BUILD(version) (ZyanU16)((version)&0x000000000000FFFF)
typedef enum ZydisFeature_ {
    ZYDIS_FEATURE_DECODER,
    ZYDIS_FEATURE_FORMATTER,
    ZYDIS_FEATURE_AVX512,
    ZYDIS_FEATURE_KNC,
    ZYDIS_FEATURE_MAX_VALUE     = ZYDIS_FEATURE_KNC,
    ZYDIS_FEATURE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_FEATURE_MAX_VALUE)
} ZydisFeature;
ZYDIS_EXPORT ZyanU64
ZydisGetVersion(void);
ZYDIS_EXPORT ZyanStatus
ZydisIsFeatureEnabled(ZydisFeature feature);
#ifdef __cplusplus
}

#endif
#endif /* ZYDIS_H */
