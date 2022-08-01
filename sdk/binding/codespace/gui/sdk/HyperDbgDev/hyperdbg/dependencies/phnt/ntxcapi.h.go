package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntxcapi.h.back

const(
_NTXCAPI_H =  //col:1
KCONTINUE_FLAG_TEST_ALERT = 0x00000001 //col:2
KCONTINUE_FLAG_DELIVER_APC = 0x00000002 //col:3
RTL_ASSERT(exp) = ((!(exp)) ? (RtlAssert((PVOID)#exp, (PVOID)__FILE__, __LINE__, NULL), FALSE) : TRUE) //col:4
RTL_ASSERTMSG(msg, exp) = ((!(exp)) ? (RtlAssert((PVOID)#exp, (PVOID)__FILE__, __LINE__, msg), FALSE) : TRUE) //col:6
)

const(
    KCONTINUE_UNWIND = 1  //col:3
    KCONTINUE_RESUME = 2  //col:4
    KCONTINUE_LONGJUMP = 3  //col:5
    KCONTINUE_SET = 4  //col:6
    KCONTINUE_LAST = 5  //col:7
)



type KCONTINUE_ARGUMENT struct{
ContinueType KCONTINUE_TYPE
ContinueFlags ULONG
Reserved[2] ULONGLONG
}




