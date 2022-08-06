

























#include <Zycore/API/Thread.h>

















#if   defined(ZYAN_POSIX)

#include <errno.h>





ZyanStatus ZyanThreadGetCurrentThread(ZyanThread* thread)
{
    *thread = pthread_self();

    return ZYAN_STATUS_SUCCESS;
}

ZYAN_STATIC_ASSERT(sizeof(ZyanThreadId) <= sizeof(ZyanU64));
ZyanStatus ZyanThreadGetCurrentThreadId(ZyanThreadId* thread_id)
{
    

    pthread_t ptid = pthread_self();
    *thread_id = *(ZyanThreadId*)ptid;

    return ZYAN_STATUS_SUCCESS;
}





ZyanStatus ZyanThreadTlsAlloc(ZyanThreadTlsIndex* index, ZyanThreadTlsCallback destructor)
{
    ZyanThreadTlsIndex value;
    const int error = pthread_key_create(&value, destructor);
    if (error != 0)
    {
        if (error == EAGAIN)
        {
            return ZYAN_STATUS_OUT_OF_RESOURCES;
        }
        if (error == ENOMEM)
        {
            return ZYAN_STATUS_NOT_ENOUGH_MEMORY;
        }
        return ZYAN_STATUS_BAD_SYSTEMCALL;
    }

    *index = value;
    return ZYAN_STATUS_SUCCESS;
}

ZyanStatus ZyanThreadTlsFree(ZyanThreadTlsIndex index)
{
    return !pthread_key_delete(index) ? ZYAN_STATUS_SUCCESS : ZYAN_STATUS_BAD_SYSTEMCALL;
}

ZyanStatus ZyanThreadTlsGetValue(ZyanThreadTlsIndex index, void** data)
{
    *data = pthread_getspecific(index);

    return ZYAN_STATUS_SUCCESS;
}

ZyanStatus ZyanThreadTlsSetValue(ZyanThreadTlsIndex index, void* data)
{
    const int error = pthread_setspecific(index, data);
    if (error != 0)
    {
        if (error == EINVAL)
        {
            return ZYAN_STATUS_INVALID_ARGUMENT;
        }
        return ZYAN_STATUS_BAD_SYSTEMCALL;
    }

    return ZYAN_STATUS_SUCCESS;
}



#elif defined(ZYAN_WINDOWS)





ZyanStatus ZyanThreadGetCurrentThread(ZyanThread* thread)
{
    *thread = GetCurrentThread();

    return ZYAN_STATUS_SUCCESS;
}

ZyanStatus ZyanThreadGetCurrentThreadId(ZyanThreadId* thread_id)
{
    *thread_id = GetCurrentThreadId();

    return ZYAN_STATUS_SUCCESS;
}





ZyanStatus ZyanThreadTlsAlloc(ZyanThreadTlsIndex* index, ZyanThreadTlsCallback destructor)
{
    const ZyanThreadTlsIndex value = FlsAlloc(destructor);
    if (value == FLS_OUT_OF_INDEXES)
    {
        return ZYAN_STATUS_OUT_OF_RESOURCES;
    }

    *index = value;
    return ZYAN_STATUS_SUCCESS;
}

ZyanStatus ZyanThreadTlsFree(ZyanThreadTlsIndex index)
{
    return FlsFree(index) ? ZYAN_STATUS_SUCCESS : ZYAN_STATUS_BAD_SYSTEMCALL;
}

ZyanStatus ZyanThreadTlsGetValue(ZyanThreadTlsIndex index, void** data)
{
    *data = FlsGetValue(index);

    return ZYAN_STATUS_SUCCESS;
}

ZyanStatus ZyanThreadTlsSetValue(ZyanThreadTlsIndex index, void* data)
{
    if (!FlsSetValue(index, data))
    {
        const DWORD error = GetLastError();
        if (error == ERROR_INVALID_PARAMETER)
        {
            return ZYAN_STATUS_INVALID_ARGUMENT;
        }
        return ZYAN_STATUS_BAD_SYSTEMCALL;
    }

    return ZYAN_STATUS_SUCCESS;
}



#else
#   error "Unsupported platform detected"
#endif


