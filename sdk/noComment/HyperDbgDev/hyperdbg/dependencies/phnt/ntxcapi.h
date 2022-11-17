/*
 * This file is part of the Process Hacker project -
 * https://processhacker.sourceforge.io/
 *
 * You can redistribute this file and/or modify it under the terms of the
 * Attribution 4.0 International (CC BY 4.0) license.
 *
 * You must give appropriate credit, provide a link to the license, and
 * indicate if changes were made. You may do so in any reasonable manner, but
 * not in any way that suggests the licensor endorses you or your use.
 */
#ifndef _NTXCAPI_H
#define _NTXCAPI_H
NTSYSAPI
BOOLEAN
NTAPI
RtlDispatchException(_In_ PEXCEPTION_RECORD ExceptionRecord,
                     _In_ PCONTEXT ContextRecord);
NTSYSAPI
DECLSPEC_NORETURN
VOID NTAPI RtlRaiseStatus(_In_ NTSTATUS Status);
NTSYSAPI
VOID NTAPI RtlRaiseException(_In_ PEXCEPTION_RECORD ExceptionRecord);
NTSYSCALLAPI
NTSTATUS
NTAPI
NtContinue(_In_ PCONTEXT ContextRecord, _In_ BOOLEAN TestAlert);
#if (PHNT_VERSION >= PHNT_THRESHOLD)
typedef enum _KCONTINUE_TYPE {
  KCONTINUE_UNWIND,
  KCONTINUE_RESUME,
  KCONTINUE_LONGJUMP,
  KCONTINUE_SET,
  KCONTINUE_LAST,
} KCONTINUE_TYPE;
typedef struct _KCONTINUE_ARGUMENT {
  KCONTINUE_TYPE ContinueType;
  ULONG ContinueFlags;
  ULONGLONG Reserved[2];
} KCONTINUE_ARGUMENT, *PKCONTINUE_ARGUMENT;
#define KCONTINUE_FLAG_TEST_ALERT 0x00000001  
#define KCONTINUE_FLAG_DELIVER_APC 0x00000002 
NTSYSCALLAPI
NTSTATUS
NTAPI
NtContinueEx(
    _In_ PCONTEXT ContextRecord,
    _In_ PVOID ContinueArgument 
);
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRaiseException(_In_ PEXCEPTION_RECORD ExceptionRecord,
                 _In_ PCONTEXT ContextRecord, _In_ BOOLEAN FirstChance);
__analysis_noreturn NTSYSCALLAPI VOID NTAPI
RtlAssert(_In_ PVOID VoidFailedAssertion, _In_ PVOID VoidFileName,
          _In_ ULONG LineNumber, _In_opt_ PSTR MutableMessage);
#define RTL_ASSERT(exp)                                                        \
  ((!(exp))                                                                    \
       ? (RtlAssert((PVOID) #exp, (PVOID)__FILE__, __LINE__, NULL), FALSE)     \
       : TRUE)
#define RTL_ASSERTMSG(msg, exp)                                                \
  ((!(exp)) ? (RtlAssert((PVOID) #exp, (PVOID)__FILE__, __LINE__, msg), FALSE) \
            : TRUE)
#define RTL_SOFT_ASSERT(_exp)                                                  \
  ((!(_exp)) ? (DbgPrint("%s(%d): Soft assertion failed\n   Expression: %s\n", \
                         __FILE__, __LINE__, #_exp),                           \
                FALSE)                                                         \
             : TRUE)
#define RTL_SOFT_ASSERTMSG(_msg, _exp)                                         \
  ((!(_exp)) ? (DbgPrint("%s(%d): Soft assertion failed\n   Expression: %s\n " \
                         "  Message: %s\n",                                    \
                         __FILE__, __LINE__, #_exp, (_msg)),                   \
                FALSE)                                                         \
             : TRUE)
#endif