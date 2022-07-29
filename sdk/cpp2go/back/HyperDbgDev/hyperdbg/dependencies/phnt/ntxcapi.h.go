package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\ntxcapi.h.back

const(
_NTXCAPI_H =  //col:13
KCONTINUE_FLAG_TEST_ALERT = 0x00000001 // wbenny //col:63
KCONTINUE_FLAG_DELIVER_APC = 0x00000002 // wbenny //col:64
RTL_ASSERT(exp) = ((!(exp)) ? (RtlAssert((PVOID)#exp, (PVOID)__FILE__, __LINE__, NULL), FALSE) : TRUE) //col:105
RTL_ASSERTMSG(msg, exp) = ((!(exp)) ? (RtlAssert((PVOID)#exp, (PVOID)__FILE__, __LINE__, msg), FALSE) : TRUE) //col:107
RTL_SOFT_ASSERT(_exp) = ((!(_exp)) ? (DbgPrint("%s(%d): Soft assertion failedn   Expression: %sn", __FILE__, __LINE__, #_exp), FALSE) : TRUE) RTL_SOFT_ASSERTMSG(_msg, _exp) ((!(_exp)) ? (DbgPrint("%s(%d): Soft assertion failedn   Expression: %sn   Message: %sn", __FILE__, __LINE__, #_exp, (_msg)), FALSE) : TRUE)  //col:109
RTL_SOFT_ASSERTMSG(_msg, _exp) = ((!(_exp)) ? (DbgPrint("%s(%d): Soft assertion failedn   Expression: %sn   Message: %sn", __FILE__, __LINE__, #_exp, (_msg)), FALSE) : TRUE)  //col:111
)

type     KCONTINUE_UNWIND uint32
const(
    KCONTINUE_UNWIND KCONTINUE_TYPE = 1  //col:49
    KCONTINUE_RESUME KCONTINUE_TYPE = 2  //col:50
    KCONTINUE_LONGJUMP KCONTINUE_TYPE = 3  //col:51
    KCONTINUE_SET KCONTINUE_TYPE = 4  //col:52
    KCONTINUE_LAST KCONTINUE_TYPE = 5  //col:53
)



type (
Ntxcapi interface{
 * Attribution 4.0 International ()(ok bool)//col:54
}
)

func NewNtxcapi() { return & ntxcapi{} }

func (n *ntxcapi) * Attribution 4.0 International ()(ok bool){//col:54
/* * Attribution 4.0 International (CC BY 4.0) license. 
 * 
 * You must give appropriate credit, provide a link to the license, and 
 * indicate if changes were made. You may do so in any reasonable manner, but 
 * not in any way that suggests the licensor endorses you or your use.
#ifndef _NTXCAPI_H
#define _NTXCAPI_H
NTSYSAPI
BOOLEAN
NTAPI
RtlDispatchException(
    _In_ PEXCEPTION_RECORD ExceptionRecord,
    _In_ PCONTEXT ContextRecord
    );
NTSYSAPI
DECLSPEC_NORETURN
VOID
NTAPI
RtlRaiseStatus(
    _In_ NTSTATUS Status
    );
NTSYSAPI
VOID
NTAPI
RtlRaiseException(
    _In_ PEXCEPTION_RECORD ExceptionRecord
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtContinue(
    _In_ PCONTEXT ContextRecord,
    _In_ BOOLEAN TestAlert
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
typedef enum _KCONTINUE_TYPE
{
    KCONTINUE_UNWIND,
    KCONTINUE_RESUME,
    KCONTINUE_LONGJUMP,
    KCONTINUE_SET,
    KCONTINUE_LAST,
} KCONTINUE_TYPE;*/
return true
}



