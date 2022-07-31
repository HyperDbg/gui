package Zydis
const(
ZYDIS_H =  //col:33
ZYDIS_VERSION = (ZyanU64)0x0003000100000000 //col:69
ZYDIS_VERSION_MAJOR(version) = (ZyanU16)(((version) & 0xFFFF000000000000) >> 48) //col:80
ZYDIS_VERSION_MINOR(version) = (ZyanU16)(((version) & 0x0000FFFF00000000) >> 32) //col:87
ZYDIS_VERSION_PATCH(version) = (ZyanU16)(((version) & 0x00000000FFFF0000) >> 16) //col:94
ZYDIS_VERSION_BUILD(version) = (ZyanU16)((version) & 0x000000000000FFFF) //col:101
)
type     ZYDIS_FEATURE_DECODER uint32
const(
    ZYDIS_FEATURE_DECODER typedef enum ZydisFeature_ = 1  //col:114
    ZYDIS_FEATURE_FORMATTER typedef enum ZydisFeature_ = 2  //col:115
    ZYDIS_FEATURE_AVX512 typedef enum ZydisFeature_ = 3  //col:116
    ZYDIS_FEATURE_KNC typedef enum ZydisFeature_ = 4  //col:117
return true
}

func (z *zydis)ZYDIS_EXPORT ZyanU64 ZydisGetVersion()(ok bool){//col:166
return true
}

