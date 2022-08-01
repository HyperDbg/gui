package Zycore
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\Zycore.h.back

const(
ZYCORE_H =  //col:1
ZYCORE_VERSION = (ZyanU64)0x0001000000000000 //col:2
ZYCORE_VERSION_MAJOR(version) = (ZyanU16)((version & 0xFFFF000000000000) >> 48) //col:3
ZYCORE_VERSION_MINOR(version) = (ZyanU16)((version & 0x0000FFFF00000000) >> 32) //col:4
ZYCORE_VERSION_PATCH(version) = (ZyanU16)((version & 0x00000000FFFF0000) >> 16) //col:5
ZYCORE_VERSION_BUILD(version) = (ZyanU16)(version & 0x000000000000FFFF) //col:6
)


