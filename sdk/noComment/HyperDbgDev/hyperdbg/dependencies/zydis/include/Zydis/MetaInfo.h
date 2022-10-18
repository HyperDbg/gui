#ifndef ZYDIS_METAINFO_H
#define ZYDIS_METAINFO_H
#include <Zycore/Defines.h>
#include <Zycore/Types.h>
#include <ZydisExportConfig.h>
#ifdef __cplusplus
extern "C" {
#endif
#include <Zydis/Generated/EnumISAExt.h>
#include <Zydis/Generated/EnumISASet.h>
#include <Zydis/Generated/EnumInstructionCategory.h>
ZYDIS_EXPORT const char *
ZydisCategoryGetString(ZydisInstructionCategory category);
ZYDIS_EXPORT const char *ZydisISASetGetString(ZydisISASet isa_set);
ZYDIS_EXPORT const char *ZydisISAExtGetString(ZydisISAExt isa_ext);
#ifdef __cplusplus
}

#endif
#endif 
