package code
//back\HyperDbgDev\hyperdbg\script-eval\code\Keywords.c.back

type (
Keywords interface{
ScriptEngineKeywordPoi()(ok bool)//col:18
ScriptEngineKeywordHi()(ok bool)//col:36
ScriptEngineKeywordLow()(ok bool)//col:54
ScriptEngineKeywordDb()(ok bool)//col:72
ScriptEngineKeywordDd()(ok bool)//col:90
ScriptEngineKeywordDw()(ok bool)//col:108
ScriptEngineKeywordDq()(ok bool)//col:126
}
)

func NewKeywords() { return & keywords{} }

func (k *keywords)ScriptEngineKeywordPoi()(ok bool){//col:18
/*ScriptEngineKeywordPoi(PUINT64 Address, BOOL * HasError)
{
    UINT64 Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(UINT64)))
    {
        *HasError = TRUE;
        return NULL;
    }
#endif 
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(UINT64));
#endif 
    return Result;
}*/
return true
}

func (k *keywords)ScriptEngineKeywordHi()(ok bool){//col:36
/*ScriptEngineKeywordHi(PUINT64 Address, BOOL * HasError)
{
    QWORD Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(UINT64)))
    {
        *HasError = TRUE;
        return NULL;
    }
#endif 
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(UINT64));
#endif 
    return HIWORD(Result);
}*/
return true
}

func (k *keywords)ScriptEngineKeywordLow()(ok bool){//col:54
/*ScriptEngineKeywordLow(PUINT64 Address, BOOL * HasError)
{
    QWORD Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(UINT64)))
    {
        *HasError = TRUE;
        return NULL;
    }
#endif 
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(UINT64));
#endif 
    return LOWORD(Result);
}*/
return true
}

func (k *keywords)ScriptEngineKeywordDb()(ok bool){//col:72
/*ScriptEngineKeywordDb(PUINT64 Address, BOOL * HasError)
{
    BYTE Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(BYTE)))
    {
        *HasError = TRUE;
        return NULL;
    }
#endif 
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(BYTE));
#endif 
    return Result;
}*/
return true
}

func (k *keywords)ScriptEngineKeywordDd()(ok bool){//col:90
/*ScriptEngineKeywordDd(PUINT64 Address, BOOL * HasError)
{
    WORD Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(WORD)))
    {
        *HasError = TRUE;
        return NULL;
    }
#endif 
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(WORD));
#endif 
    return Result;
}*/
return true
}

func (k *keywords)ScriptEngineKeywordDw()(ok bool){//col:108
/*ScriptEngineKeywordDw(PUINT64 Address, BOOL * HasError)
{
    DWORD Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(DWORD)))
    {
        *HasError = TRUE;
        return NULL;
    }
#endif 
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(DWORD));
#endif 
    return Result;
}*/
return true
}

func (k *keywords)ScriptEngineKeywordDq()(ok bool){//col:126
/*ScriptEngineKeywordDq(PUINT64 Address, BOOL * HasError)
{
    QWORD Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(DWORD)))
    {
        *HasError = TRUE;
        return NULL;
    }
#endif 
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(QWORD));
#endif 
    return Result;
}*/
return true
}



