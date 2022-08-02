package API

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\src\API\Synchronization.c.back

type (
	Synchronization interface {
		ZyanStatus_ZyanCriticalSectionInitialize() (ok bool) //col:37
		ZyanStatus_ZyanCriticalSectionEnter() (ok bool)      //col:54
		ZyanBool_ZyanCriticalSectionTryEnter() (ok bool)     //col:58
		ZyanStatus_ZyanCriticalSectionLeave() (ok bool)      //col:75
		ZyanStatus_ZyanCriticalSectionDelete() (ok bool)     //col:88
		ZyanStatus_ZyanCriticalSectionInitialize() (ok bool) //col:98
	}
	synchronization struct{}
)

func NewSynchronization() Synchronization { return &synchronization{} }

func (s *synchronization) ZyanStatus_ZyanCriticalSectionInitialize() (ok bool) { //col:37
	/*
	   ZyanStatus ZyanCriticalSectionInitialize(ZyanCriticalSection* critical_section)

	   	{
	   	    pthread_mutexattr_t attribute;
	   	    int error = pthread_mutexattr_init(&attribute);
	   	    if (error != 0)
	   	    {
	   	        if (error == ENOMEM)
	   	        {
	   	            return ZYAN_STATUS_NOT_ENOUGH_MEMORY;
	   	        }
	   	        return ZYAN_STATUS_BAD_SYSTEMCALL;
	   	    }
	   	    pthread_mutexattr_settype(&attribute, PTHREAD_MUTEX_RECURSIVE);
	   	    error = pthread_mutex_init(critical_section, &attribute);
	   	    pthread_mutexattr_destroy(&attribute);
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
	   	        if (error == EPERM)
	   	        {
	   	            return ZYAN_STATUS_ACCESS_DENIED;
	   	        }
	   	        if ((error == EBUSY) || (error == EINVAL))
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

func (s *synchronization) ZyanStatus_ZyanCriticalSectionEnter() (ok bool) { //col:54
	/*
	   ZyanStatus ZyanCriticalSectionEnter(ZyanCriticalSection* critical_section)

	   	{
	   	    const int error = pthread_mutex_lock(critical_section);
	   	    if (error != 0)
	   	    {
	   	        if (error == EINVAL)
	   	        {
	   	            return ZYAN_STATUS_INVALID_ARGUMENT;
	   	        }
	   	        if (error == EAGAIN)
	   	        {
	   	            return ZYAN_STATUS_INVALID_OPERATION;
	   	        }
	   	        return ZYAN_STATUS_BAD_SYSTEMCALL;
	   	    }
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (s *synchronization) ZyanBool_ZyanCriticalSectionTryEnter() (ok bool) { //col:58
	/*
	   ZyanBool ZyanCriticalSectionTryEnter(ZyanCriticalSection* critical_section)

	   	{
	   	    return pthread_mutex_trylock(critical_section) ? ZYAN_FALSE : ZYAN_TRUE;
	   	}
	*/
	return true
}

func (s *synchronization) ZyanStatus_ZyanCriticalSectionLeave() (ok bool) { //col:75
	/*
	   ZyanStatus ZyanCriticalSectionLeave(ZyanCriticalSection* critical_section)

	   	{
	   	    const int error = pthread_mutex_unlock(critical_section);
	   	    if (error != 0)
	   	    {
	   	        if (error == EINVAL)
	   	        {
	   	            return ZYAN_STATUS_INVALID_ARGUMENT;
	   	        }
	   	        if (error == EPERM)
	   	        {
	   	            return ZYAN_STATUS_INVALID_OPERATION;
	   	        }
	   	        return ZYAN_STATUS_BAD_SYSTEMCALL;
	   	    }
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (s *synchronization) ZyanStatus_ZyanCriticalSectionDelete() (ok bool) { //col:88
	/*
	   ZyanStatus ZyanCriticalSectionDelete(ZyanCriticalSection* critical_section)

	   	{
	   	    const int error = pthread_mutex_destroy(critical_section);
	   	    if (error != 0)
	   	    {
	   	        if ((error == EBUSY) || (error == EINVAL))
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

func (s *synchronization) ZyanStatus_ZyanCriticalSectionInitialize() (ok bool) { //col:98
	/*
	   ZyanStatus ZyanCriticalSectionInitialize(ZyanCriticalSection* critical_section)

	   	{
	   	    InitializeCriticalSection(critical_section);

	   ZyanStatus ZyanCriticalSectionEnter(ZyanCriticalSection* critical_section)

	   	{
	   	    EnterCriticalSection(critical_section);

	   ZyanBool ZyanCriticalSectionTryEnter(ZyanCriticalSection* critical_section)

	   	{
	   	    return TryEnterCriticalSection(critical_section) ? ZYAN_TRUE : ZYAN_FALSE;
	   	}
	*/
	return true
}

