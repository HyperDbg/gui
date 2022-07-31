package phnt
const(
_NTWOW64_H =  //col:13
WOW64_SYSTEM_DIRECTORY = "SysWOW64" //col:15
WOW64_SYSTEM_DIRECTORY_U = L"SysWOW64" //col:16
WOW64_X86_TAG = " (x86)" //col:17
WOW64_X86_TAG_U = L" (x86)" //col:18
WOW64_POINTER(Type) = ULONG //col:39
LDR_DATA_TABLE_ENTRY_SIZE_WINXP_32 = FIELD_OFFSET(LDR_DATA_TABLE_ENTRY32, DdagNode) //col:108
LDR_DATA_TABLE_ENTRY_SIZE_WIN7_32 = FIELD_OFFSET(LDR_DATA_TABLE_ENTRY32, BaseNameHashValue) //col:109
LDR_DATA_TABLE_ENTRY_SIZE_WIN8_32 = FIELD_OFFSET(LDR_DATA_TABLE_ENTRY32, ImplicitPathOptions) //col:110
GDI_BATCH_BUFFER_SIZE = 310 //col:400
)
type     SharedNtdll32LdrInitializeThunk uint32
const(
    SharedNtdll32LdrInitializeThunk WOW64_SHARED_INFORMATION = 1  //col:23
    SharedNtdll32KiUserExceptionDispatcher WOW64_SHARED_INFORMATION = 2  //col:24
    SharedNtdll32KiUserApcDispatcher WOW64_SHARED_INFORMATION = 3  //col:25
    SharedNtdll32KiUserCallbackDispatcher WOW64_SHARED_INFORMATION = 4  //col:26
    SharedNtdll32ExpInterlockedPopEntrySListFault WOW64_SHARED_INFORMATION = 5  //col:27
    SharedNtdll32ExpInterlockedPopEntrySListResume WOW64_SHARED_INFORMATION = 6  //col:28
    SharedNtdll32ExpInterlockedPopEntrySListEnd WOW64_SHARED_INFORMATION = 7  //col:29
    SharedNtdll32RtlUserThreadStart WOW64_SHARED_INFORMATION = 8  //col:30
    SharedNtdll32pQueryProcessDebugInformationRemote WOW64_SHARED_INFORMATION = 9  //col:31
    SharedNtdll32BaseAddress WOW64_SHARED_INFORMATION = 10  //col:32
    SharedNtdll32LdrSystemDllInitBlock WOW64_SHARED_INFORMATION = 11  //col:33
    Wow64SharedPageEntriesCount WOW64_SHARED_INFORMATION = 12  //col:34
)
type (
Ntwow64 interface{
 * Attribution 4.0 International ()(ok bool)//col:35
#define WOW64_POINTER()(ok bool)//col:58
    WOW64_POINTER()(ok bool)//col:64
    WOW64_POINTER()(ok bool)//col:77
    WOW64_POINTER()(ok bool)//col:83
    WOW64_POINTER()(ok bool)//col:88
    WOW64_POINTER()(ok bool)//col:106
#define LDR_DATA_TABLE_ENTRY_SIZE_WINXP_32 FIELD_OFFSET()(ok bool)//col:183
    WOW64_POINTER()(ok bool)//col:189
    WOW64_POINTER()(ok bool)//col:245
    WOW64_POINTER()(ok bool)//col:390
C_ASSERT()(ok bool)//col:407
    WOW64_POINTER()(ok bool)//col:553
C_ASSERT()(ok bool)//col:581
FORCEINLINE VOID UStrToUStr32()(ok bool)//col:591
}

)
func NewNtwow64() { return & ntwow64{} }
func (n *ntwow64) * Attribution 4.0 International ()(ok bool){//col:35
return true
}

func (n *ntwow64)#define WOW64_POINTER()(ok bool){//col:58
return true
}

func (n *ntwow64)    WOW64_POINTER()(ok bool){//col:64
return true
}

func (n *ntwow64)    WOW64_POINTER()(ok bool){//col:77
return true
}

func (n *ntwow64)    WOW64_POINTER()(ok bool){//col:83
return true
}

func (n *ntwow64)    WOW64_POINTER()(ok bool){//col:88
return true
}

func (n *ntwow64)    WOW64_POINTER()(ok bool){//col:106
return true
}

func (n *ntwow64)#define LDR_DATA_TABLE_ENTRY_SIZE_WINXP_32 FIELD_OFFSET()(ok bool){//col:183
return true
}

func (n *ntwow64)    WOW64_POINTER()(ok bool){//col:189
return true
}

func (n *ntwow64)    WOW64_POINTER()(ok bool){//col:245
return true
}

func (n *ntwow64)    WOW64_POINTER()(ok bool){//col:390
return true
}

func (n *ntwow64)C_ASSERT()(ok bool){//col:407
return true
}

func (n *ntwow64)    WOW64_POINTER()(ok bool){//col:553
return true
}

func (n *ntwow64)C_ASSERT()(ok bool){//col:581
return true
}

func (n *ntwow64)FORCEINLINE VOID UStrToUStr32()(ok bool){//col:591
return true
}

