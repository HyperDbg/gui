#ifndef ZYCORE_ARGPARSE_H
#define ZYCORE_ARGPARSE_H
#include <Zycore/Types.h>
#include <Zycore/Status.h>
#include <Zycore/Vector.h>
#include <Zycore/String.h>
#ifdef __cplusplus
extern "C" {
#endif
typedef struct ZyanArgParseDefinition_ {
    const char * name;
    ZyanBool     boolean;
    ZyanBool     required;
} ZyanArgParseDefinition;
typedef struct ZyanArgParseConfig_ {
    const char **            argv;
    ZyanUSize                argc;
    ZyanUSize                min_unnamed_args;
    ZyanUSize                max_unnamed_args;
    ZyanArgParseDefinition * args;
} ZyanArgParseConfig;
typedef struct ZyanArgParseArg_ {
    const ZyanArgParseDefinition * def;
    ZyanBool                       has_value;
    ZyanStringView                 value;
} ZyanArgParseArg;
#ifndef ZYAN_NO_LIBC
ZYCORE_EXPORT ZyanStatus
ZyanArgParse(const ZyanArgParseConfig * cfg, ZyanVector * parsed, const char ** error_token);
#endif
ZYCORE_EXPORT ZyanStatus
ZyanArgParseEx(const ZyanArgParseConfig * cfg, ZyanVector * parsed, const char ** error_token, ZyanAllocator * allocator);
#ifdef __cplusplus
}

#endif
#endif /* ZYCORE_ARGPARSE_H */
