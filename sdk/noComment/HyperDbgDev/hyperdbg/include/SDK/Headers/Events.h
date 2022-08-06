#pragma once
typedef enum _DEBUGGER_EVENT_TYPE_ENUM
{
    HIDDEN_HOOK_READ_AND_WRITE,
    HIDDEN_HOOK_READ,
    HIDDEN_HOOK_WRITE,
    HIDDEN_HOOK_EXEC_DETOURS,
    HIDDEN_HOOK_EXEC_CC,
    SYSCALL_HOOK_EFER_SYSCALL,
    SYSCALL_HOOK_EFER_SYSRET,
    CPUID_INSTRUCTION_EXECUTION,
    RDMSR_INSTRUCTION_EXECUTION,
    WRMSR_INSTRUCTION_EXECUTION,
    IN_INSTRUCTION_EXECUTION,
    OUT_INSTRUCTION_EXECUTION,
    EXCEPTION_OCCURRED,
    EXTERNAL_INTERRUPT_OCCURRED,
    DEBUG_REGISTERS_ACCESSED,
    TSC_INSTRUCTION_EXECUTION,
    PMC_INSTRUCTION_EXECUTION,
    VMCALL_INSTRUCTION_EXECUTION,
    CONTROL_REGISTER_MODIFIED,
} DEBUGGER_EVENT_TYPE_ENUM;
typedef enum _DEBUGGER_EVENT_ACTION_TYPE_ENUM
{
    BREAK_TO_DEBUGGER,
    RUN_SCRIPT,
    RUN_CUSTOM_CODE
} DEBUGGER_EVENT_ACTION_TYPE_ENUM;
typedef enum _DEBUGGER_EVENT_SYSCALL_SYSRET_TYPE
{
    DEBUGGER_EVENT_SYSCALL_SYSRET_SAFE_ACCESS_MEMORY = 0,
    DEBUGGER_EVENT_SYSCALL_SYSRET_HANDLE_ALL_UD      = 1,
} DEBUGGER_EVENT_SYSCALL_SYSRET_TYPE;
#define SIZEOF_DEBUGGER_MODIFY_EVENTS sizeof(DEBUGGER_MODIFY_EVENTS)
typedef enum _DEBUGGER_MODIFY_EVENTS_TYPE
{
    DEBUGGER_MODIFY_EVENTS_QUERY_STATE,
    DEBUGGER_MODIFY_EVENTS_ENABLE,
    DEBUGGER_MODIFY_EVENTS_DISABLE,
    DEBUGGER_MODIFY_EVENTS_CLEAR
} DEBUGGER_MODIFY_EVENTS_TYPE;
typedef struct _DEBUGGER_MODIFY_EVENTS
{
    UINT64 Tag;          
    UINT64 KernelStatus; 
    DEBUGGER_MODIFY_EVENTS_TYPE
    TypeOfAction;      
    BOOLEAN IsEnabled; 
} DEBUGGER_MODIFY_EVENTS, *PDEBUGGER_MODIFY_EVENTS;
