#ifndef ZYDIS_REGISTER_H
#define ZYDIS_REGISTER_H
#include <Zycore/Defines.h>
#include <Zycore/Types.h>
#include <Zydis/SharedTypes.h>
#include <Zydis/ShortString.h>
#ifdef __cplusplus
extern "C" {
#endif
#include <Zydis/Generated/EnumRegister.h>
typedef enum ZydisRegisterClass_ {
  ZYDIS_REGCLASS_INVALID,
  ZYDIS_REGCLASS_GPR8,
  ZYDIS_REGCLASS_GPR16,
  ZYDIS_REGCLASS_GPR32,
  ZYDIS_REGCLASS_GPR64,
  ZYDIS_REGCLASS_X87,
  ZYDIS_REGCLASS_MMX,
  ZYDIS_REGCLASS_XMM,
  ZYDIS_REGCLASS_YMM,
  ZYDIS_REGCLASS_ZMM,
  ZYDIS_REGCLASS_FLAGS,
  ZYDIS_REGCLASS_IP,
  ZYDIS_REGCLASS_SEGMENT,
  ZYDIS_REGCLASS_TEST,
  ZYDIS_REGCLASS_CONTROL,
  ZYDIS_REGCLASS_DEBUG,
  ZYDIS_REGCLASS_MASK,
  ZYDIS_REGCLASS_BOUND,
  ZYDIS_REGCLASS_MAX_VALUE = ZYDIS_REGCLASS_BOUND,
  ZYDIS_REGCLASS_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_REGCLASS_MAX_VALUE)
} ZydisRegisterClass;
typedef ZyanU16 ZydisRegisterWidth;
typedef struct ZydisRegisterContext_ {
  ZyanU64 values[ZYDIS_REGISTER_MAX_VALUE + 1];
} ZydisRegisterContext;
ZYDIS_EXPORT ZydisRegister
ZydisRegisterEncode(ZydisRegisterClass register_class, ZyanU8 id);
ZYDIS_EXPORT ZyanI8 ZydisRegisterGetId(ZydisRegister reg);
ZYDIS_EXPORT ZydisRegisterClass ZydisRegisterGetClass(ZydisRegister reg);
ZYDIS_EXPORT ZydisRegisterWidth ZydisRegisterGetWidth(ZydisMachineMode mode,
                                                      ZydisRegister reg);
ZYDIS_EXPORT ZydisRegister
ZydisRegisterGetLargestEnclosing(ZydisMachineMode mode, ZydisRegister reg);
ZYDIS_EXPORT const char *ZydisRegisterGetString(ZydisRegister reg);
ZYDIS_EXPORT const ZydisShortString *
ZydisRegisterGetStringWrapped(ZydisRegister reg);
ZYDIS_EXPORT ZydisRegisterWidth ZydisRegisterClassGetWidth(
    ZydisMachineMode mode, ZydisRegisterClass register_class);
#ifdef __cplusplus
}

#endif
#endif 
