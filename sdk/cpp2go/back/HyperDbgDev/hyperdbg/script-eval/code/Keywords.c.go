package code
//back\HyperDbgDev\hyperdbg\script-eval\code\Keywords.c.back

type (
Keywords interface{
ScriptEngineKeywordPoi()(ok bool)//col:51
ScriptEngineKeywordHi()(ok bool)//col:85
ScriptEngineKeywordLow()(ok bool)//col:119
ScriptEngineKeywordDb()(ok bool)//col:153
ScriptEngineKeywordDd()(ok bool)//col:187
ScriptEngineKeywordDw()(ok bool)//col:221
ScriptEngineKeywordDq()(ok bool)//col:255
}
)

func NewKeywords() { return & keywords{} }

func (k *keywords)ScriptEngineKeywordPoi()(ok bool){//col:51
/*ScriptEngineKeywordPoi(PUINT64 Address, BOOL * HasError)
{
    UINT64 Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(UINT64)))
    {
        *HasError = TRUE;
        return NULL;
    }
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(UINT64));
    return Result;
}*/
return true
}

func (k *keywords)ScriptEngineKeywordHi()(ok bool){//col:85
/*ScriptEngineKeywordHi(PUINT64 Address, BOOL * HasError)
{
    QWORD Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(UINT64)))
    {
        *HasError = TRUE;
        return NULL;
    }
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(UINT64));
    return HIWORD(Result);
}*/
return true
}

func (k *keywords)ScriptEngineKeywordLow()(ok bool){//col:119
/*ScriptEngineKeywordLow(PUINT64 Address, BOOL * HasError)
{
    QWORD Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(UINT64)))
    {
        *HasError = TRUE;
        return NULL;
    }
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(UINT64));
    return LOWORD(Result);
}*/
return true
}

func (k *keywords)ScriptEngineKeywordDb()(ok bool){//col:153
/*ScriptEngineKeywordDb(PUINT64 Address, BOOL * HasError)
{
    BYTE Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(BYTE)))
    {
        *HasError = TRUE;
        return NULL;
    }
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(BYTE));
    return Result;
}*/
return true
}

func (k *keywords)ScriptEngineKeywordDd()(ok bool){//col:187
/*ScriptEngineKeywordDd(PUINT64 Address, BOOL * HasError)
{
    WORD Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(WORD)))
    {
        *HasError = TRUE;
        return NULL;
    }
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(WORD));
    return Result;
}*/
return true
}

func (k *keywords)ScriptEngineKeywordDw()(ok bool){//col:221
/*ScriptEngineKeywordDw(PUINT64 Address, BOOL * HasError)
{
    DWORD Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(DWORD)))
    {
        *HasError = TRUE;
        return NULL;
    }
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(DWORD));
    return Result;
}*/
return true
}

func (k *keywords)ScriptEngineKeywordDq()(ok bool){//col:255
/*ScriptEngineKeywordDq(PUINT64 Address, BOOL * HasError)
{
    QWORD Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(DWORD)))
    {
        *HasError = TRUE;
        return NULL;
    }
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(QWORD));
    return Result;
}*/
return true
}



