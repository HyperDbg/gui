#include <Zydis/MetaInfo.h>
#include <Generated/EnumInstructionCategory.inc>
#include <Generated/EnumISASet.inc>
#include <Generated/EnumISAExt.inc>
const char *
ZydisCategoryGetString(ZydisInstructionCategory category) {
    if ((ZyanUSize)category >= ZYAN_ARRAY_LENGTH(STR_INSTRUCTIONCATEGORY)) {
        return ZYAN_NULL;
    }
    return STR_INSTRUCTIONCATEGORY[category];
}

const char *
ZydisISASetGetString(ZydisISASet isa_set) {
    if ((ZyanUSize)isa_set >= ZYAN_ARRAY_LENGTH(STR_ISASET)) {
        return ZYAN_NULL;
    }
    return STR_ISASET[isa_set];
}

const char *
ZydisISAExtGetString(ZydisISAExt isa_ext) {
    if ((ZyanUSize)isa_ext >= ZYAN_ARRAY_LENGTH(STR_ISAEXT)) {
        return ZYAN_NULL;
    }
    return STR_ISAEXT[isa_ext];
}
