#include "pch.h"
UINT64
ScriptEngineKeywordPoi(PUINT64 Address, BOOL * HasError) {
    UINT64 Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(UINT64))) {
        *HasError = TRUE;
        return NULL;
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(UINT64));
#endif // SCRIPT_ENGINE_KERNEL_MODE
    return Result;
}

WORD
ScriptEngineKeywordHi(PUINT64 Address, BOOL * HasError) {
    QWORD Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(UINT64))) {
        *HasError = TRUE;
        return NULL;
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(UINT64));
#endif // SCRIPT_ENGINE_KERNEL_MODE
    return HIWORD(Result);
}

WORD
ScriptEngineKeywordLow(PUINT64 Address, BOOL * HasError) {
    QWORD Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(UINT64))) {
        *HasError = TRUE;
        return NULL;
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(UINT64));
#endif // SCRIPT_ENGINE_KERNEL_MODE
    return LOWORD(Result);
}

BYTE
ScriptEngineKeywordDb(PUINT64 Address, BOOL * HasError) {
    BYTE Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(BYTE))) {
        *HasError = TRUE;
        return NULL;
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(BYTE));
#endif // SCRIPT_ENGINE_KERNEL_MODE
    return Result;
}

WORD
ScriptEngineKeywordDd(PUINT64 Address, BOOL * HasError) {
    WORD Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(WORD))) {
        *HasError = TRUE;
        return NULL;
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(WORD));
#endif // SCRIPT_ENGINE_KERNEL_MODE
    return Result;
}

DWORD
ScriptEngineKeywordDw(PUINT64 Address, BOOL * HasError) {
    DWORD Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(DWORD))) {
        *HasError = TRUE;
        return NULL;
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(DWORD));
#endif // SCRIPT_ENGINE_KERNEL_MODE
    return Result;
}

QWORD
ScriptEngineKeywordDq(PUINT64 Address, BOOL * HasError) {
    QWORD Result = NULL;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(DWORD))) {
        *HasError = TRUE;
        return NULL;
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = *Address;
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperReadMemorySafeOnTargetProcess(Address, &Result, sizeof(QWORD));
#endif // SCRIPT_ENGINE_KERNEL_MODE
    return Result;
}
