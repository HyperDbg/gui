#ifndef ZYDIS_SHORTSTRING_H
#define ZYDIS_SHORTSTRING_H
#include <ZydisExportConfig.h>
#include <Zycore/Defines.h>
#include <Zycore/Types.h>
#ifdef __cplusplus
extern "C" {
#endif
#pragma pack(push, 1)
typedef struct ZydisShortString_
{
    const char* data;
    ZyanU8 size;
} ZydisShortString;
#pragma pack(pop)
#define ZYDIS_MAKE_SHORTSTRING(string) \
    { string, sizeof(string) - 1 }
#ifdef __cplusplus
}

#endif
#endif 
