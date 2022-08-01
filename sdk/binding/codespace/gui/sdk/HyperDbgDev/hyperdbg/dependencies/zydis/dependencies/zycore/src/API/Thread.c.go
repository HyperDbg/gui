package API

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\src\API\Thread.c.back

type (
	Thread interface {
		ZyanStatus_ZyanThreadGetCurrentThread() (ok bool)   //col:5
		ZYAN_STATIC_ASSERT() (ok bool)                      //col:12
		ZyanStatus_ZyanThreadTlsAlloc() (ok bool)           //col:31
		ZyanStatus_ZyanThreadTlsFree() (ok bool)            //col:35
		ZyanStatus_ZyanThreadTlsGetValue() (ok bool)        //col:40
		ZyanStatus_ZyanThreadTlsSetValue() (ok bool)        //col:53
		ZyanStatus_ZyanThreadGetCurrentThread() (ok bool)   //col:58
		ZyanStatus_ZyanThreadGetCurrentThreadId() (ok bool) //col:63
		ZyanStatus_ZyanThreadTlsAlloc() (ok bool)           //col:73
		ZyanStatus_ZyanThreadTlsFree() (ok bool)            //col:77
		ZyanStatus_ZyanThreadTlsGetValue() (ok bool)        //col:82
		ZyanStatus_ZyanThreadTlsSetValue() (ok bool)        //col:95
	}
	thread struct{}
)

func NewThread() Thread { return &thread{} }

func (t *thread) ZyanStatus_ZyanThreadGetCurrentThread() (ok bool) { //col:5
	/*
	   ZyanStatus ZyanThreadGetCurrentThread(ZyanThread* thread)

	   	{
	   	    *thread = pthread_self();
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (t *thread) ZYAN_STATIC_ASSERT() (ok bool) { //col:12
	/*
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

func (t *thread) ZyanStatus_ZyanThreadTlsAlloc() (ok bool) { //col:31
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

func (t *thread) ZyanStatus_ZyanThreadTlsFree() (ok bool) { //col:35
	/*
	   ZyanStatus ZyanThreadTlsFree(ZyanThreadTlsIndex index)

	   	{
	   	    return !pthread_key_delete(index) ? ZYAN_STATUS_SUCCESS : ZYAN_STATUS_BAD_SYSTEMCALL;
	   	}
	*/
	return true
}

func (t *thread) ZyanStatus_ZyanThreadTlsGetValue() (ok bool) { //col:40
	/*
	   ZyanStatus ZyanThreadTlsGetValue(ZyanThreadTlsIndex index, void** data)

	   	{
	   	    *data = pthread_getspecific(index);
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (t *thread) ZyanStatus_ZyanThreadTlsSetValue() (ok bool) { //col:53
	/*
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

func (t *thread) ZyanStatus_ZyanThreadGetCurrentThread() (ok bool) { //col:58
	/*
	   ZyanStatus ZyanThreadGetCurrentThread(ZyanThread* thread)

	   	{
	   	    *thread = GetCurrentThread();
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (t *thread) ZyanStatus_ZyanThreadGetCurrentThreadId() (ok bool) { //col:63
	/*
	   ZyanStatus ZyanThreadGetCurrentThreadId(ZyanThreadId* thread_id)

	   	{
	   	    *thread_id = GetCurrentThreadId();
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (t *thread) ZyanStatus_ZyanThreadTlsAlloc() (ok bool) { //col:73
	/*
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

func (t *thread) ZyanStatus_ZyanThreadTlsFree() (ok bool) { //col:77
	/*
	   ZyanStatus ZyanThreadTlsFree(ZyanThreadTlsIndex index)

	   	{
	   	    return FlsFree(index) ? ZYAN_STATUS_SUCCESS : ZYAN_STATUS_BAD_SYSTEMCALL;
	   	}
	*/
	return true
}

func (t *thread) ZyanStatus_ZyanThreadTlsGetValue() (ok bool) { //col:82
	/*
	   ZyanStatus ZyanThreadTlsGetValue(ZyanThreadTlsIndex index, void** data)

	   	{
	   	    *data = FlsGetValue(index);
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (t *thread) ZyanStatus_ZyanThreadTlsSetValue() (ok bool) { //col:95
	/*
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

