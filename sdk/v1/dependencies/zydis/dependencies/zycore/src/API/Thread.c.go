package API

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\src\API\Thread.c.back

type (
	Thread interface {
		ZyanStatus_ZyanThreadGetCurrentThread() (ok bool) //col:10
		ZyanStatus_ZyanThreadTlsAlloc() (ok bool)         //col:29
		ZyanStatus_ZyanThreadTlsFree() (ok bool)          //col:33
		ZyanStatus_ZyanThreadTlsGetValue() (ok bool)      //col:49
		ZyanStatus_ZyanThreadGetCurrentThread() (ok bool) //col:65
		ZyanStatus_ZyanThreadTlsFree() (ok bool)          //col:69
		ZyanStatus_ZyanThreadTlsGetValue() (ok bool)      //col:85
	}
	thread struct{}
)

func NewThread() Thread { return &thread{} }

func (t *thread) ZyanStatus_ZyanThreadGetCurrentThread() (ok bool) { //col:10
	/*
	   ZyanStatus ZyanThreadGetCurrentThread(ZyanThread* thread)

	   	{
	   	    *thread = pthread_self();

	   ZYAN_STATIC_ASSERT(sizeof(ZyanThreadId) <= sizeof(ZyanU64));
	   ZyanStatus ZyanThreadGetCurrentThreadId(ZyanThreadId* thread_id)

	   	{
	   	    pthread_t ptid = pthread_self();
	   	    *thread_id = *(ZyanThreadId*)ptid;
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (t *thread) ZyanStatus_ZyanThreadTlsAlloc() (ok bool) { //col:29
	/*
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
	*/
	return true
}

func (t *thread) ZyanStatus_ZyanThreadTlsFree() (ok bool) { //col:33
	/*
	   ZyanStatus ZyanThreadTlsFree(ZyanThreadTlsIndex index)

	   	{
	   	    return !pthread_key_delete(index) ? ZYAN_STATUS_SUCCESS : ZYAN_STATUS_BAD_SYSTEMCALL;
	   	}
	*/
	return true
}

func (t *thread) ZyanStatus_ZyanThreadTlsGetValue() (ok bool) { //col:49
	/*
	   ZyanStatus ZyanThreadTlsGetValue(ZyanThreadTlsIndex index, void** data)

	   	{
	   	    *data = pthread_getspecific(index);

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
	*/
	return true
}

func (t *thread) ZyanStatus_ZyanThreadGetCurrentThread() (ok bool) { //col:65
	/*
	   ZyanStatus ZyanThreadGetCurrentThread(ZyanThread* thread)

	   	{
	   	    *thread = GetCurrentThread();

	   ZyanStatus ZyanThreadGetCurrentThreadId(ZyanThreadId* thread_id)

	   	{
	   	    *thread_id = GetCurrentThreadId();

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
	*/
	return true
}

func (t *thread) ZyanStatus_ZyanThreadTlsFree() (ok bool) { //col:69
	/*
	   ZyanStatus ZyanThreadTlsFree(ZyanThreadTlsIndex index)

	   	{
	   	    return FlsFree(index) ? ZYAN_STATUS_SUCCESS : ZYAN_STATUS_BAD_SYSTEMCALL;
	   	}
	*/
	return true
}

func (t *thread) ZyanStatus_ZyanThreadTlsGetValue() (ok bool) { //col:85
	/*
	   ZyanStatus ZyanThreadTlsGetValue(ZyanThreadTlsIndex index, void** data)

	   	{
	   	    *data = FlsGetValue(index);

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
	*/
	return true
}

