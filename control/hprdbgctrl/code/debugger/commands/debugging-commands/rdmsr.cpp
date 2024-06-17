#include "pch.h"
VOID CommandRdmsrHelp(){
    ShowMessages("rdmsr : reads a model-specific register (MSR).\n\n");
    ShowMessages("syntax : \trdmsr [Msr (hex)] [core CoreNumber (hex)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : rdmsr c0000082\n");
    ShowMessages("\t\te.g : rdmsr c0000082 core 2\n");
}
typedef BOOL(WINAPI * glpie_t)(
    LOGICAL_PROCESSOR_RELATIONSHIP,
    PSYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX,
    PDWORD);
static SIZE_T
GetWindowsCompatibleNumberOfCores(){
    SYSTEM_INFO SysInfo;
    GetSystemInfo(&SysInfo);
    return (SIZE_T)SysInfo.dwNumberOfProcessors;
}
static SIZE_T
GetWindowsNumaNumberOfCores(){
    uint8_t * Buffer   = NULL;
    SIZE_T    NumCores = 0;
    DWORD     Length   = 0;
    HMODULE   Kernel32 = GetModuleHandleW(L"kernel32.dll");
    glpie_t GetLogicalProcessorInformationEx = (glpie_t)GetProcAddress(Kernel32, "GetLogicalProcessorInformationEx");
    if (!GetLogicalProcessorInformationEx){
        return 0;
    }
    GetLogicalProcessorInformationEx(RelationAll, NULL, &Length);
    if (Length < 1 || GetLastError() != ERROR_INSUFFICIENT_BUFFER){
        return 0;
    }
    Buffer = (uint8_t *)malloc(Length);
    if (!Buffer || !GetLogicalProcessorInformationEx(RelationAll, (PSYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX)Buffer, &Length)){
        free(Buffer);
        return 0;
    }
    for (DWORD Offset = 0; Offset < Length;){
        PSYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX Info = (PSYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX)(Buffer + Offset);
        Offset += Info->Size;
        if (Info->Relationship != RelationProcessorCore){
            continue;
        }
        for (WORD group = 0; group < Info->Processor.GroupCount; ++group){
            for (KAFFINITY Mask = Info->Processor.GroupMask[group].Mask; Mask != 0; Mask >>= 1){
                NumCores += Mask & 1;
            }
        }
    }
    free(Buffer);
    return NumCores;
}
VOID CommandRdmsr(vector<string> SplitCommand, string Command){
    BOOL                           Status;
    SIZE_T                         NumCPU;
    DEBUGGER_READ_AND_WRITE_ON_MSR MsrReadRequest;
    ULONG                          ReturnedLength;
    UINT64                         Msr;
    BOOL                           IsNextCoreId   = FALSE;
    BOOL                           SetMsr         = FALSE;
    UINT32                         CoreNumer      = DEBUGGER_READ_AND_WRITE_ON_MSR_APPLY_ALL_CORES;
    BOOLEAN                        IsFirstCommand = TRUE;
    if (SplitCommand.size() >= 5){
        ShowMessages("incorrect use of the 'rdmsr'\n\n");
        CommandRdmsrHelp();
        return;
    }
    for (auto Section : SplitCommand){
        if (IsFirstCommand == TRUE){
            IsFirstCommand = FALSE;
            continue;
        }
        if (IsNextCoreId){
            if (!ConvertStringToUInt32(Section, &CoreNumer)){
                ShowMessages("please specify a correct hex value for core id\n\n");
                CommandRdmsrHelp();
                return;
            }
            IsNextCoreId = FALSE;
            continue;
        }
        if (!Section.compare("core")){
            IsNextCoreId = TRUE;
            continue;
        }
        if (SetMsr || !ConvertStringToUInt64(Section, &Msr)){
            ShowMessages("please specify a correct hex value to be read\n\n");
            CommandRdmsrHelp();
            return;
        }
        SetMsr = TRUE;
    }
    if (!SetMsr){
        ShowMessages("please specify a correct hex value to be read\n\n");
        CommandRdmsrHelp();
        return;
    }
    if (IsNextCoreId){
        ShowMessages("please specify a correct hex value for core\n\n");
        CommandRdmsrHelp();
        return;
    }
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
    MsrReadRequest.ActionType = DEBUGGER_MSR_READ;
    MsrReadRequest.Msr        = Msr;
    MsrReadRequest.CoreNumber = CoreNumer;
    SIZE_T NumCores = GetWindowsNumaNumberOfCores();
    NumCPU          = NumCores > 0 ? NumCores : GetWindowsCompatibleNumberOfCores();
    UINT64 * OutputBuffer = (UINT64 *)malloc(sizeof(UINT64) * NumCPU);
    ZeroMemory(OutputBuffer, sizeof(UINT64) * NumCPU);
    Status = DeviceIoControl(
        g_DeviceHandle,                        
        IOCTL_DEBUGGER_READ_OR_WRITE_MSR,      
        &MsrReadRequest,                       
        SIZEOF_DEBUGGER_READ_AND_WRITE_ON_MSR, 
        OutputBuffer,                          
        (DWORD)(sizeof(UINT64) * NumCPU),      
        &ReturnedLength,                       
        NULL                                   
    );
    if (!Status){
        ShowMessages("ioctl failed with code (%x), either msr index or core id is invalid\n",
                     GetLastError());
        free(OutputBuffer);
        return;
    }
    if (CoreNumer == DEBUGGER_READ_AND_WRITE_ON_MSR_APPLY_ALL_CORES){
        for (SIZE_T i = 0; i < NumCPU; i++){
            ShowMessages("core : 0x%x - msr[%llx] = %s\n", i, Msr, SeparateTo64BitValue((OutputBuffer[i])).c_str());
        }
    }
    else{
        ShowMessages("core : 0x%x - msr[%llx] = %s\n", CoreNumer, Msr, SeparateTo64BitValue((OutputBuffer[0])).c_str());
    }
    free(OutputBuffer);
}
