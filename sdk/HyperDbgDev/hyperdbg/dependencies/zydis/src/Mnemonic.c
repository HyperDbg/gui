#include <Zydis/Mnemonic.h>
#include <Generated/EnumMnemonic.inc>
const char *
ZydisMnemonicGetString(ZydisMnemonic mnemonic) {
    if ((ZyanUSize)mnemonic >= ZYAN_ARRAY_LENGTH(STR_MNEMONIC)) {
        return ZYAN_NULL;
    }
    return (const char *)STR_MNEMONIC[mnemonic].data;
}

const ZydisShortString *
ZydisMnemonicGetStringWrapped(ZydisMnemonic mnemonic) {
    if ((ZyanUSize)mnemonic >= ZYAN_ARRAY_LENGTH(STR_MNEMONIC)) {
        return ZYAN_NULL;
    }
    return &STR_MNEMONIC[mnemonic];
}
