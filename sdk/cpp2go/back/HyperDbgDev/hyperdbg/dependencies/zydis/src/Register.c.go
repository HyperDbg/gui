package src
//back\HyperDbgDev\hyperdbg\dependencies\zydis\src\Register.c.back

type (
Register interface{
  Zyan Disassembler Library ()(ok bool)//col:64
ZydisRegister ZydisRegisterEncode()(ok bool)//col:115
ZyanI8 ZydisRegisterGetId()(ok bool)//col:135
ZydisRegisterClass ZydisRegisterGetClass()(ok bool)//col:147
ZydisRegisterWidth ZydisRegisterGetWidth()(ok bool)//col:189
ZydisRegister ZydisRegisterGetLargestEnclosing()(ok bool)//col:266
const char* ZydisRegisterGetString()(ok bool)//col:275
const ZydisShortString* ZydisRegisterGetStringWrapped()(ok bool)//col:284
ZydisRegisterWidth ZydisRegisterClassGetWidth()(ok bool)//col:299
}
)

func NewRegister() { return & register{} }

func (r *register)  Zyan Disassembler Library ()(ok bool){//col:64
/*  Zyan Disassembler Library (Zydis)
  Original Author : Florian Bernd
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
#include <Zydis/Register.h>
#include <Generated/EnumRegister.inc>
 * Defines the `ZydisRegisterMapItem` struct.
typedef struct ZydisRegisterMapItem_
{
     * The register class.
    ZydisRegisterClass class;
     * The lowest register of the current class.
    ZydisRegister lo;
     * The highest register of the current class.
    ZydisRegister hi;
     * The width of registers of the current class in 16- and 32-bit mode.
    ZydisRegisterWidth width;
     * The width of registers of the current class in 64-bit mode.
    ZydisRegisterWidth width64;
} ZydisRegisterMapItem;*/
return true
}

func (r *register)ZydisRegister ZydisRegisterEncode()(ok bool){//col:115
/*ZydisRegister ZydisRegisterEncode(ZydisRegisterClass register_class, ZyanU8 id)
{
    switch (register_class)
    {
    case ZYDIS_REGCLASS_INVALID:
    case ZYDIS_REGCLASS_FLAGS:
    case ZYDIS_REGCLASS_IP:
        break;
    default:
        if (((ZyanUSize)register_class < ZYAN_ARRAY_LENGTH(REGISTER_MAP)) &&
            (id <= (REGISTER_MAP[register_class].hi - REGISTER_MAP[register_class].lo)))
        {
            return REGISTER_MAP[register_class].lo + id;
        }
    }
    return ZYDIS_REGISTER_NONE;
}*/
return true
}

func (r *register)ZyanI8 ZydisRegisterGetId()(ok bool){//col:135
/*ZyanI8 ZydisRegisterGetId(ZydisRegister reg)
{
    for (ZyanUSize i = 0; i < ZYAN_ARRAY_LENGTH(REGISTER_MAP); ++i)
    {
        switch (REGISTER_MAP[i].class)
        {
        case ZYDIS_REGCLASS_INVALID:
        case ZYDIS_REGCLASS_FLAGS:
        case ZYDIS_REGCLASS_IP:
            break;
        default:
            if ((reg >= REGISTER_MAP[i].lo) && (reg <= REGISTER_MAP[i].hi))
            {
                return (ZyanU8)(reg - REGISTER_MAP[i].lo);
            }
        }
    }
    return -1;
}*/
return true
}

func (r *register)ZydisRegisterClass ZydisRegisterGetClass()(ok bool){//col:147
/*ZydisRegisterClass ZydisRegisterGetClass(ZydisRegister reg)
{
    for (ZyanUSize i = 0; i < ZYAN_ARRAY_LENGTH(REGISTER_MAP); ++i)
    {
        if ((reg >= REGISTER_MAP[i].lo) && (reg <= REGISTER_MAP[i].hi))
        {
            return REGISTER_MAP[i].class;
        }
    }
    return ZYDIS_REGCLASS_INVALID;
}*/
return true
}

func (r *register)ZydisRegisterWidth ZydisRegisterGetWidth()(ok bool){//col:189
/*ZydisRegisterWidth ZydisRegisterGetWidth(ZydisMachineMode mode, ZydisRegister reg)
{
    switch (reg)
    {
    case ZYDIS_REGISTER_X87CONTROL:
    case ZYDIS_REGISTER_X87STATUS:
    case ZYDIS_REGISTER_X87TAG:
        return 16;
    case ZYDIS_REGISTER_IP:
    case ZYDIS_REGISTER_FLAGS:
        return 16;
    case ZYDIS_REGISTER_EIP:
    case ZYDIS_REGISTER_EFLAGS:
        return 32;
    case ZYDIS_REGISTER_RIP:
    case ZYDIS_REGISTER_RFLAGS:
        return (mode == ZYDIS_MACHINE_MODE_LONG_64) ? 64 : 0;
    case ZYDIS_REGISTER_BNDCFG:
    case ZYDIS_REGISTER_BNDSTATUS:
        return 64;
    case ZYDIS_REGISTER_XCR0:
        return 64;
    case ZYDIS_REGISTER_PKRU:
    case ZYDIS_REGISTER_MXCSR:
        return 32;
    default:
        break;
    }
    for (ZyanUSize i = 0; i < ZYAN_ARRAY_LENGTH(REGISTER_MAP); ++i)
    {
        if ((reg >= REGISTER_MAP[i].lo) && (reg <= REGISTER_MAP[i].hi))
        {
            return (mode == ZYDIS_MACHINE_MODE_LONG_64) ?
                REGISTER_MAP[i].width64 : REGISTER_MAP[i].width;
        }
    }
    return 0;
}*/
return true
}

func (r *register)ZydisRegister ZydisRegisterGetLargestEnclosing()(ok bool){//col:266
/*ZydisRegister ZydisRegisterGetLargestEnclosing(ZydisMachineMode mode,
    ZydisRegister reg)
{
    static const ZyanU8 GPR8_MAPPING[20] =
    {
    };
    for (ZyanUSize i = 0; i < ZYAN_ARRAY_LENGTH(REGISTER_MAP); ++i)
    {
        if ((reg >= REGISTER_MAP[i].lo) && (reg <= REGISTER_MAP[i].hi))
        {
            const ZydisRegisterClass reg_class = REGISTER_MAP[i].class;
            if ((reg_class == ZYDIS_REGCLASS_GPR64) && (mode != ZYDIS_MACHINE_MODE_LONG_64))
            {
                return ZYDIS_REGISTER_NONE;
            }
            ZyanU8 reg_id = (ZyanU8)(reg - REGISTER_MAP[reg_class].lo);
            switch (reg_class)
            {
            case ZYDIS_REGCLASS_GPR8:
                reg_id = GPR8_MAPPING[reg_id];
                ZYAN_FALLTHROUGH;
            case ZYDIS_REGCLASS_GPR16:
            case ZYDIS_REGCLASS_GPR32:
            case ZYDIS_REGCLASS_GPR64:
                switch (mode)
                {
                case ZYDIS_MACHINE_MODE_LONG_64:
                    return REGISTER_MAP[ZYDIS_REGCLASS_GPR64].lo + reg_id;
                case ZYDIS_MACHINE_MODE_LONG_COMPAT_32:
                case ZYDIS_MACHINE_MODE_LEGACY_32:
                    return REGISTER_MAP[ZYDIS_REGCLASS_GPR32].lo + reg_id;
                case ZYDIS_MACHINE_MODE_LONG_COMPAT_16:
                case ZYDIS_MACHINE_MODE_LEGACY_16:
                case ZYDIS_MACHINE_MODE_REAL_16:
                    return REGISTER_MAP[ZYDIS_REGCLASS_GPR16].lo + reg_id;
                default:
                    return ZYDIS_REGISTER_NONE;
                }
            case ZYDIS_REGCLASS_XMM:
            case ZYDIS_REGCLASS_YMM:
            case ZYDIS_REGCLASS_ZMM:
#if defined(ZYDIS_DISABLE_AVX512) && defined(ZYDIS_DISABLE_KNC)
                return REGISTER_MAP[ZYDIS_REGCLASS_YMM].lo + reg_id;
#else
                return REGISTER_MAP[ZYDIS_REGCLASS_ZMM].lo + reg_id;
#endif
            default:
                return ZYDIS_REGISTER_NONE;
            }
        }
    }
    return ZYDIS_REGISTER_NONE;
}*/
return true
}

func (r *register)const char* ZydisRegisterGetString()(ok bool){//col:275
/*const char* ZydisRegisterGetString(ZydisRegister reg)
{
    if ((ZyanUSize)reg >= ZYAN_ARRAY_LENGTH(STR_REGISTER))
    {
        return ZYAN_NULL;
    }
    return STR_REGISTER[reg].data;
}*/
return true
}

func (r *register)const ZydisShortString* ZydisRegisterGetStringWrapped()(ok bool){//col:284
/*const ZydisShortString* ZydisRegisterGetStringWrapped(ZydisRegister reg)
{
    if ((ZyanUSize)reg >= ZYAN_ARRAY_LENGTH(STR_REGISTER))
    {
        return ZYAN_NULL;
    }
    return &STR_REGISTER[reg];
}*/
return true
}

func (r *register)ZydisRegisterWidth ZydisRegisterClassGetWidth()(ok bool){//col:299
/*ZydisRegisterWidth ZydisRegisterClassGetWidth(ZydisMachineMode mode,
    ZydisRegisterClass register_class)
{
    if ((ZyanUSize)register_class < ZYAN_ARRAY_LENGTH(REGISTER_MAP))
    {
        return (mode == ZYDIS_MACHINE_MODE_LONG_64) ?
            REGISTER_MAP[register_class].width64 : REGISTER_MAP[register_class].width;
    }
    return 0;
}*/
return true
}



