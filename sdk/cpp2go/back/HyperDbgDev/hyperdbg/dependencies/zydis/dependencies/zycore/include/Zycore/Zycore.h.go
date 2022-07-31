package Zycore


const(
ZYCORE_H =  //col:33
ZYCORE_VERSION = (ZyanU64)0x0001000000000000 //col:55
ZYCORE_VERSION_MAJOR(version) = (ZyanU16)((version & 0xFFFF000000000000) >> 48) //col:66
ZYCORE_VERSION_MINOR(version) = (ZyanU16)((version & 0x0000FFFF00000000) >> 32) //col:73
ZYCORE_VERSION_PATCH(version) = (ZyanU16)((version & 0x00000000FFFF0000) >> 16) //col:80
ZYCORE_VERSION_BUILD(version) = (ZyanU16)(version & 0x000000000000FFFF) //col:87
)

type (
Zycore interface{
ZYCORE_EXPORT ZyanU64 ZycoreGetVersion()(ok bool)//col:108
}






)

func NewZycore() { return & zycore{} }

func (z *zycore)ZYCORE_EXPORT ZyanU64 ZycoreGetVersion()(ok bool){//col:108



return true
}









