#ifndef ZYDIS_FORMATTER_TOKEN_H
#define ZYDIS_FORMATTER_TOKEN_H
#include <ZydisExportConfig.h>
#include <Zycore/String.h>
#include <Zycore/Types.h>
#include <Zydis/Status.h>
#ifdef __cplusplus
extern "C" {
#endif
typedef ZyanU8 ZydisTokenType;
#define ZYDIS_TOKEN_INVALID           0x00
#define ZYDIS_TOKEN_WHITESPACE        0x01
#define ZYDIS_TOKEN_DELIMITER         0x02
#define ZYDIS_TOKEN_PARENTHESIS_OPEN  0x03
#define ZYDIS_TOKEN_PARENTHESIS_CLOSE 0x04
#define ZYDIS_TOKEN_PREFIX            0x05
#define ZYDIS_TOKEN_MNEMONIC          0x06
#define ZYDIS_TOKEN_REGISTER          0x07
#define ZYDIS_TOKEN_ADDRESS_ABS       0x08
#define ZYDIS_TOKEN_ADDRESS_REL       0x09
#define ZYDIS_TOKEN_DISPLACEMENT      0x0A
#define ZYDIS_TOKEN_IMMEDIATE         0x0B
#define ZYDIS_TOKEN_TYPECAST          0x0C
#define ZYDIS_TOKEN_DECORATOR         0x0D
#define ZYDIS_TOKEN_SYMBOL            0x0E
#define ZYDIS_TOKEN_USER              0x80
#pragma pack(push, 1)
typedef struct ZydisFormatterToken_ {
    ZydisTokenType type;
    ZyanU8         next;
} ZydisFormatterToken;
#pragma pack(pop)
typedef const ZydisFormatterToken ZydisFormatterTokenConst;
typedef struct ZydisFormatterBuffer_ {
    ZyanBool   is_token_list;
    ZyanUSize  capacity;
    ZyanString string;
} ZydisFormatterBuffer;
ZYDIS_EXPORT ZyanStatus
ZydisFormatterTokenGetValue(const ZydisFormatterToken * token,
                            ZydisTokenType *            type,
                            ZyanConstCharPointer *      value);
ZYDIS_EXPORT ZyanStatus
ZydisFormatterTokenNext(ZydisFormatterTokenConst ** token);
ZYDIS_EXPORT ZyanStatus
ZydisFormatterBufferGetToken(const ZydisFormatterBuffer * buffer,
                             ZydisFormatterTokenConst **  token);
ZYDIS_EXPORT ZyanStatus
ZydisFormatterBufferGetString(ZydisFormatterBuffer * buffer,
                              ZyanString **          string);
ZYDIS_EXPORT ZyanStatus
ZydisFormatterBufferAppend(ZydisFormatterBuffer * buffer,
                           ZydisTokenType         type);
ZYDIS_EXPORT ZyanStatus
ZydisFormatterBufferRemember(const ZydisFormatterBuffer * buffer,
                             ZyanUPointer *               state);
ZYDIS_EXPORT ZyanStatus
ZydisFormatterBufferRestore(ZydisFormatterBuffer * buffer,
                            ZyanUPointer           state);
#ifdef __cplusplus
}

#endif
#endif /* ZYDIS_FORMATTER_TOKEN_H */
