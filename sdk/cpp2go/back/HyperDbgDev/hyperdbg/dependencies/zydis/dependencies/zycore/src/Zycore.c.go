package src
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\src\Zycore.c.back

type (
Zycore interface{
ZyanU64_ZycoreGetVersion()(ok bool)//col:4
}
)

func NewZycore() { return & zycore{} }

func (z *zycore)ZyanU64_ZycoreGetVersion()(ok bool){//col:4
/*ZyanU64 ZycoreGetVersion(void)
{
    return ZYCORE_VERSION;
}*/
return true
}



