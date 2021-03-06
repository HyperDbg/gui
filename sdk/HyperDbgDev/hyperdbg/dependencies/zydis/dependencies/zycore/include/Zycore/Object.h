#ifndef ZYCORE_OBJECT_H
#define ZYCORE_OBJECT_H
#include <Zycore/Status.h>
#include <Zycore/Types.h>
#ifdef __cplusplus
extern "C" {
#endif
typedef void (*ZyanMemberProcedure)(void * object);
typedef void (*ZyanConstMemberProcedure)(const void * object);
typedef ZyanStatus (*ZyanMemberFunction)(void * object);
typedef ZyanStatus (*ZyanConstMemberFunction)(const void * object);
#ifdef __cplusplus
}

#endif
#endif /* ZYCORE_OBJECT_H */
