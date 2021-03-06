#ifndef ZYDIS_MNEMONIC_H
#define ZYDIS_MNEMONIC_H
#include <Zycore/Types.h>
#include <Zydis/ShortString.h>
#ifdef __cplusplus
extern "C" {
#endif
#include <Zydis/Generated/EnumMnemonic.h>
ZYDIS_EXPORT const char *
ZydisMnemonicGetString(ZydisMnemonic mnemonic);
ZYDIS_EXPORT const ZydisShortString *
ZydisMnemonicGetStringWrapped(ZydisMnemonic mnemonic);
#ifdef __cplusplus
}

#endif
#endif /* ZYDIS_MNEMONIC_H */
