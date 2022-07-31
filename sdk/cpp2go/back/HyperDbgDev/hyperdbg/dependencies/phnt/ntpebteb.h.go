package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\ntpebteb.h.back

const(
_NTPEBTEB_H =  //col:13
GDI_BATCH_BUFFER_SIZE = 310 //col:251
)

type (
Ntpebteb interface{
 * Attribution 4.0 International ()(ok bool)//col:26
C_ASSERT()(ok bool)//col:258
}
)

func NewNtpebteb() { return & ntpebteb{} }

func (n *ntpebteb) * Attribution 4.0 International ()(ok bool){//col:26
/* * Attribution 4.0 International (CC BY 4.0) license. 
 * 
 * You must give appropriate credit, provide a link to the license, and 
 * indicate if changes were made. You may do so in any reasonable manner, but 
 * not in any way that suggests the licensor endorses you or your use.
#ifndef _NTPEBTEB_H
#define _NTPEBTEB_H
typedef struct _RTL_USER_PROCESS_PARAMETERS *PRTL_USER_PROCESS_PARAMETERS;
typedef struct _RTL_CRITICAL_SECTION *PRTL_CRITICAL_SECTION;
typedef struct _ACTIVATION_CONTEXT_STACK
{
    struct _RTL_ACTIVATION_CONTEXT_STACK_FRAME* ActiveFrame;
    LIST_ENTRY FrameListCache;
    ULONG Flags;
    ULONG NextCookieSequenceNumber;
    ULONG StackId;
} ACTIVATION_CONTEXT_STACK, *PACTIVATION_CONTEXT_STACK;*/
return true
}

func (n *ntpebteb)C_ASSERT()(ok bool){//col:258
/*C_ASSERT(FIELD_OFFSET(PEB, SessionId) == 0x2C0);
#else
C_ASSERT(FIELD_OFFSET(PEB, SessionId) == 0x1D4);
#endif
#define GDI_BATCH_BUFFER_SIZE 310
typedef struct _GDI_TEB_BATCH
{
    ULONG Offset;
    ULONG_PTR HDC;
    ULONG Buffer[GDI_BATCH_BUFFER_SIZE];
} GDI_TEB_BATCH, *PGDI_TEB_BATCH;*/
return true
}



