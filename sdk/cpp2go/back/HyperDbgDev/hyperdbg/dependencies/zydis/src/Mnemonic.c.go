package src
//back\HyperDbgDev\hyperdbg\dependencies\zydis\src\Mnemonic.c.back

type (
Mnemonic interface{
const_charPtr_ZydisMnemonicGetString()(ok bool)//col:8
const_ZydisShortStringPtr_ZydisMnemonicGetStringWrapped()(ok bool)//col:16
}
)

func NewMnemonic() { return & mnemonic{} }

func (m *mnemonic)const_charPtr_ZydisMnemonicGetString()(ok bool){//col:8
/*const char* ZydisMnemonicGetString(ZydisMnemonic mnemonic)
{
    if ((ZyanUSize)mnemonic >= ZYAN_ARRAY_LENGTH(STR_MNEMONIC))
    {
        return ZYAN_NULL;
    }
    return (const char*)STR_MNEMONIC[mnemonic].data;
}*/
return true
}

func (m *mnemonic)const_ZydisShortStringPtr_ZydisMnemonicGetStringWrapped()(ok bool){//col:16
/*const ZydisShortString* ZydisMnemonicGetStringWrapped(ZydisMnemonic mnemonic)
{
    if ((ZyanUSize)mnemonic >= ZYAN_ARRAY_LENGTH(STR_MNEMONIC))
    {
        return ZYAN_NULL;
    }
    return &STR_MNEMONIC[mnemonic];
}*/
return true
}



