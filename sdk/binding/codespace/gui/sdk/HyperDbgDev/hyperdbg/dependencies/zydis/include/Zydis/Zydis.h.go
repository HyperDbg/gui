package Zydis

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Zydis.h.back

const (
	ZYDIS_H = //col:1
	ZYDIS_VERSION = (
ZyanU64
)0x0003000100000000 //col:2
ZYDIS_VERSION_MAJOR(version) = (ZyanU16)(((version) & 0xFFFF000000000000) >> 48) //col:3
ZYDIS_VERSION_MINOR(version) = (ZyanU16)(((version) & 0x0000FFFF00000000) >> 32) //col:4
ZYDIS_VERSION_PATCH(version) = (ZyanU16)(((version) & 0x00000000FFFF0000) >> 16) //col:5
ZYDIS_VERSION_BUILD(version) = (ZyanU16)((version) & 0x000000000000FFFF) //col:6
)

const (
	ZYDIS_FEATURE_DECODER       = 1                                               //col:3
	ZYDIS_FEATURE_FORMATTER     = 2                                               //col:4
	ZYDIS_FEATURE_AVX512        = 3                                               //col:5
	ZYDIS_FEATURE_KNC           = 4                                               //col:6
	ZYDIS_FEATURE_MAX_VALUE     = ZYDIS_FEATURE_KNC                               //col:7
	ZYDIS_FEATURE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_FEATURE_MAX_VALUE) //col:8
)
