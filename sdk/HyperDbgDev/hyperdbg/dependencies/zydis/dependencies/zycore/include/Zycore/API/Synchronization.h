#ifndef ZYCORE_SYNCHRONIZATION_H
#define ZYCORE_SYNCHRONIZATION_H
#ifndef ZYAN_NO_LIBC
#    include <ZycoreExportConfig.h>
#    include <Zycore/Defines.h>
#    include <Zycore/Status.h>
#    ifdef __cplusplus
extern "C" {
#    endif
#    if defined(ZYAN_POSIX)
#        include <pthread.h>
typedef pthread_mutex_t ZyanCriticalSection;
#    elif defined(ZYAN_WINDOWS)
#        include <windows.h>
typedef CRITICAL_SECTION ZyanCriticalSection;
#    else
#        error "Unsupported platform detected"
#    endif
ZYCORE_EXPORT ZyanStatus
ZyanCriticalSectionInitialize(ZyanCriticalSection * critical_section);
ZYCORE_EXPORT ZyanStatus
ZyanCriticalSectionEnter(ZyanCriticalSection * critical_section);
ZYCORE_EXPORT ZyanBool
ZyanCriticalSectionTryEnter(ZyanCriticalSection * critical_section);
ZYCORE_EXPORT ZyanStatus
ZyanCriticalSectionLeave(ZyanCriticalSection * critical_section);
ZYCORE_EXPORT ZyanStatus
ZyanCriticalSectionDelete(ZyanCriticalSection * critical_section);
#    ifdef __cplusplus
}

#    endif
#endif /* ZYAN_NO_LIBC */
#endif /* ZYCORE_SYNCHRONIZATION_H */
