#ifndef ZYCORE_H
#define ZYCORE_H
#include <ZycoreExportConfig.h>
#include <Zycore/Types.h>
#ifdef __cplusplus
extern "C" {
#endif
#define ZYCORE_VERSION (ZyanU64)0x0001000000000000
#define ZYCORE_VERSION_MAJOR(version) (ZyanU16)((version & 0xFFFF000000000000) >> 48)
#define ZYCORE_VERSION_MINOR(version) (ZyanU16)((version & 0x0000FFFF00000000) >> 32)
#define ZYCORE_VERSION_PATCH(version) (ZyanU16)((version & 0x00000000FFFF0000) >> 16)
#define ZYCORE_VERSION_BUILD(version) (ZyanU16)(version & 0x000000000000FFFF)
ZYCORE_EXPORT ZyanU64 ZycoreGetVersion(void);
#ifdef __cplusplus
}

#endif
#endif 
