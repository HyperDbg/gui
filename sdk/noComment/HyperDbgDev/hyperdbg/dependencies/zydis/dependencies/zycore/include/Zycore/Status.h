#ifndef ZYCORE_STATUS_H
#define ZYCORE_STATUS_H
#ifdef __cplusplus
extern "C" {
#endif
#include <Zycore/Types.h>
typedef ZyanU32 ZyanStatus;
#define ZYAN_MAKE_STATUS(error, module, code)                                  \
  (ZyanStatus)((((error)&0x01) << 31) | (((module)&0x7FF) << 20) |             \
               ((code)&0xFFFFF))
#define ZYAN_SUCCESS(status) (!((status)&0x80000000))
#define ZYAN_FAILED(status) ((status)&0x80000000)
#define ZYAN_CHECK(status)                                                     \
  do {                                                                         \
    const ZyanStatus status_047620348 = (status);                              \
    if (!ZYAN_SUCCESS(status_047620348)) {                                     \
      return status_047620348;                                                 \
    }                                                                          \
  } while (0)
#define ZYAN_STATUS_MODULE(status) (((status) >> 20) & 0x7FF)
#define ZYAN_STATUS_CODE(status) ((status)&0xFFFFF)
#define ZYAN_MODULE_ZYCORE 0x001
#define ZYAN_MODULE_ARGPARSE 0x003
#define ZYAN_MODULE_USER 0x3FF
#define ZYAN_STATUS_SUCCESS ZYAN_MAKE_STATUS(0, ZYAN_MODULE_ZYCORE, 0x00)
#define ZYAN_STATUS_FAILED ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x01)
#define ZYAN_STATUS_TRUE ZYAN_MAKE_STATUS(0, ZYAN_MODULE_ZYCORE, 0x02)
#define ZYAN_STATUS_FALSE ZYAN_MAKE_STATUS(0, ZYAN_MODULE_ZYCORE, 0x03)
#define ZYAN_STATUS_INVALID_ARGUMENT                                           \
  ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x04)
#define ZYAN_STATUS_INVALID_OPERATION                                          \
  ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x05)
#define ZYAN_STATUS_ACCESS_DENIED ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x06)
#define ZYAN_STATUS_NOT_FOUND ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x07)
#define ZYAN_STATUS_OUT_OF_RANGE ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x08)
#define ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE                                   \
  ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x09)
#define ZYAN_STATUS_NOT_ENOUGH_MEMORY                                          \
  ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x0A)
#define ZYAN_STATUS_BAD_SYSTEMCALL ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x0B)
#define ZYAN_STATUS_OUT_OF_RESOURCES                                           \
  ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x0C)
#define ZYAN_STATUS_MISSING_DEPENDENCY                                         \
  ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x0D)
#define ZYAN_STATUS_ARG_NOT_UNDERSTOOD                                         \
  ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ARGPARSE, 0x00)
#define ZYAN_STATUS_TOO_FEW_ARGS ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ARGPARSE, 0x01)
#define ZYAN_STATUS_TOO_MANY_ARGS                                              \
  ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ARGPARSE, 0x02)
#define ZYAN_STATUS_ARG_MISSES_VALUE                                           \
  ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ARGPARSE, 0x03)
#define ZYAN_STATUS_REQUIRED_ARG_MISSING                                       \
  ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ARGPARSE, 0x04)
#ifdef __cplusplus
}

#endif
#endif 
