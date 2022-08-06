#ifndef ZYDIS_DECODER_H
#define ZYDIS_DECODER_H
#include <Zycore/Types.h>
#include <Zycore/Defines.h>
#include <Zydis/DecoderTypes.h>
#include <Zydis/Status.h>
#ifdef __cplusplus
extern "C" {
#endif
typedef enum ZydisDecoderMode_
{
    ZYDIS_DECODER_MODE_MINIMAL,
    ZYDIS_DECODER_MODE_AMD_BRANCHES,
    ZYDIS_DECODER_MODE_KNC,
    ZYDIS_DECODER_MODE_MPX,
    ZYDIS_DECODER_MODE_CET,
    ZYDIS_DECODER_MODE_LZCNT,
    ZYDIS_DECODER_MODE_TZCNT,
    ZYDIS_DECODER_MODE_WBNOINVD,
    ZYDIS_DECODER_MODE_CLDEMOTE,
    ZYDIS_DECODER_MODE_MAX_VALUE = ZYDIS_DECODER_MODE_CLDEMOTE,
    ZYDIS_DECODER_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_DECODER_MODE_MAX_VALUE)
} ZydisDecoderMode;
typedef struct ZydisDecoder_
{
    ZydisMachineMode machine_mode;
    ZydisAddressWidth address_width;
    ZyanBool decoder_mode[ZYDIS_DECODER_MODE_MAX_VALUE + 1];
} ZydisDecoder;
ZYDIS_EXPORT ZyanStatus ZydisDecoderInit(ZydisDecoder* decoder, ZydisMachineMode machine_mode,
    ZydisAddressWidth address_width);
ZYDIS_EXPORT ZyanStatus ZydisDecoderEnableMode(ZydisDecoder* decoder, ZydisDecoderMode mode,
    ZyanBool enabled);
ZYDIS_EXPORT ZyanStatus ZydisDecoderDecodeBuffer(const ZydisDecoder* decoder,
    const void* buffer, ZyanUSize length, ZydisDecodedInstruction* instruction);
#ifdef __cplusplus
}

#endif
#endif 
