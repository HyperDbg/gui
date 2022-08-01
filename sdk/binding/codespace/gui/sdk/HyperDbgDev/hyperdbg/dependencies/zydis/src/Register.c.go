package src
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\src\Register.c.back

type typedef struct ZydisRegisterMapItem_ struct{
class ZydisRegisterClass
lo ZydisRegister
hi ZydisRegister
width ZydisRegisterWidth
width64 ZydisRegisterWidth
}



type (
Register interface{
ZydisRegister_ZydisRegisterEncode()(ok bool)//col:17
ZyanI8_ZydisRegisterGetId()(ok bool)//col:36
ZydisRegisterClass_ZydisRegisterGetClass()(ok bool)//col:47
ZydisRegisterWidth_ZydisRegisterGetWidth()(ok bool)//col:85
ZydisRegister_ZydisRegisterGetLargestEnclosing()(ok bool)//col:158
const_charPtr_ZydisRegisterGetString()(ok bool)//col:166
const_ZydisShortStringPtr_ZydisRegisterGetStringWrapped()(ok bool)//col:174
ZydisRegisterWidth_ZydisRegisterClassGetWidth()(ok bool)//col:184
}
)

func NewRegister() { return & register{} }

func (r *register)ZydisRegister_ZydisRegisterEncode()(ok bool){//col:17
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

func (r *register)ZyanI8_ZydisRegisterGetId()(ok bool){//col:36
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

func (r *register)ZydisRegisterClass_ZydisRegisterGetClass()(ok bool){//col:47
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

func (r *register)ZydisRegisterWidth_ZydisRegisterGetWidth()(ok bool){//col:85
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

func (r *register)ZydisRegister_ZydisRegisterGetLargestEnclosing()(ok bool){//col:158
/*ZydisRegister ZydisRegisterGetLargestEnclosing(ZydisMachineMode mode,
    ZydisRegister reg)
{
    static const ZyanU8 GPR8_MAPPING[20] =
    {
         0,
         1,
         2,
         3,
         0,
         1,
         2,
         3,
         4,
         5,
         6,
         7,
         8,
         9,
        10,
        11,
        12,
        13,
        14,
        15,
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

func (r *register)const_charPtr_ZydisRegisterGetString()(ok bool){//col:166
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

func (r *register)const_ZydisShortStringPtr_ZydisRegisterGetStringWrapped()(ok bool){//col:174
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

func (r *register)ZydisRegisterWidth_ZydisRegisterClassGetWidth()(ok bool){//col:184
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



