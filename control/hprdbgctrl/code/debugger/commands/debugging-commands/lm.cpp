#include "pch.h"
using namespace std;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;
VOID CommandLmHelp(){
    ShowMessages("lm : lists user/kernel modules' base address, size, name and path.\n\n");
    ShowMessages("syntax : \tlm [m Name (string)] [pid ProcessId (hex)] [Filter (string)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : lm\n");
    ShowMessages("\t\te.g : lm km\n");
    ShowMessages("\t\te.g : lm um\n");
    ShowMessages("\t\te.g : lm m nt\n");
    ShowMessages("\t\te.g : lm km m ntos\n");
    ShowMessages("\t\te.g : lm um m kernel32\n");
    ShowMessages("\t\te.g : lm um m kernel32 pid 1240\n");
}
std::wstring CommandLmConvertWow64CompatibilityPaths(const wchar_t * LocalFilePath){
    std::wstring filePath(LocalFilePath);
    std::transform(filePath.begin(), filePath.end(), filePath.begin(), ::tolower);
    size_t pos = filePath.find(L":\\windows\\system32");
    if (pos != std::string::npos){
        filePath.replace(pos, 18, L":\\Windows\\SysWOW64");
    }
    pos = filePath.find(L":\\program files");
    if (pos != std::string::npos){
        filePath.replace(pos, 15, L":\\Program Files (x86)");
    }
    return filePath;
}
BOOLEAN CommandLmShowUserModeModule(UINT32 ProcessId, const char * SearchModule){
    BOOLEAN                         Status;
    ULONG                           ReturnedLength;
    UINT32                          ModuleDetailsSize    = 0;
    UINT32                          ModulesCount         = 0;
    PUSERMODE_LOADED_MODULE_DETAILS ModuleDetailsRequest = NULL;
    PUSERMODE_LOADED_MODULE_SYMBOLS Modules              = NULL;
    USERMODE_LOADED_MODULE_DETAILS  ModuleCountRequest   = {0};
    size_t                          CharSize             = 0;
    wchar_t *                       WcharBuff            = NULL;
    wstring                         SearchModuleString;
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
    ModuleCountRequest.ProcessId        = ProcessId;
    ModuleCountRequest.OnlyCountModules = TRUE;
    Status = DeviceIoControl(
        g_DeviceHandle,                         
        IOCTL_GET_USER_MODE_MODULE_DETAILS,     
        &ModuleCountRequest,                    
        sizeof(USERMODE_LOADED_MODULE_DETAILS), 
        &ModuleCountRequest,                    
        sizeof(USERMODE_LOADED_MODULE_DETAILS), 
        &ReturnedLength,                        
        NULL                                    
    );
    if (!Status){
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return FALSE;
    }
    if (ModuleCountRequest.Result == DEBUGGER_OPERATION_WAS_SUCCESSFUL){
        ModulesCount = ModuleCountRequest.ModulesCount;
        ModuleDetailsSize = sizeof(USERMODE_LOADED_MODULE_DETAILS) +
                            (ModuleCountRequest.ModulesCount * sizeof(USERMODE_LOADED_MODULE_SYMBOLS));
        ModuleDetailsRequest = (PUSERMODE_LOADED_MODULE_DETAILS)malloc(ModuleDetailsSize);
        if (ModuleDetailsRequest == NULL){
            return FALSE;
        }
        RtlZeroMemory(ModuleDetailsRequest, ModuleDetailsSize);
        ModuleDetailsRequest->ProcessId        = ProcessId;
        ModuleDetailsRequest->OnlyCountModules = FALSE;
        Status = DeviceIoControl(
            g_DeviceHandle,                         
            IOCTL_GET_USER_MODE_MODULE_DETAILS,     
            ModuleDetailsRequest,                   
            sizeof(USERMODE_LOADED_MODULE_DETAILS), 
            ModuleDetailsRequest,                   
            ModuleDetailsSize,                      
            &ReturnedLength,                        
            NULL                                    
        );
        if (!Status){
            free(ModuleDetailsRequest);
            ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
            return FALSE;
        }
        if (ModuleCountRequest.Result == DEBUGGER_OPERATION_WAS_SUCCESSFUL){
            Modules = (PUSERMODE_LOADED_MODULE_SYMBOLS)((UINT64)ModuleDetailsRequest +
                                                        sizeof(USERMODE_LOADED_MODULE_DETAILS));
            ShowMessages("user mode\n");
            ShowMessages("start\t\t\tentrypoint\t\tpath\n\n");
            if (SearchModule != NULL){
                CharSize  = strlen(SearchModule) + 1;
                WcharBuff = (wchar_t *)malloc(CharSize * 2);
                if (WcharBuff == NULL){
                    return FALSE;
                }
                RtlZeroMemory(WcharBuff, CharSize);
                mbstowcs(WcharBuff, SearchModule, CharSize);
                SearchModuleString.assign(WcharBuff, wcslen(WcharBuff));
            }
            for (size_t i = 0; i < ModulesCount; i++){
                if (SearchModule != NULL){
                    std::wstring FullPathName((wchar_t *)Modules[i].FilePath);
                    if (FindCaseInsensitiveW(FullPathName, SearchModuleString, 0) == std::wstring::npos){
                        continue;
                    }
                }
                if (ModuleDetailsRequest->Is32Bit){
                    ShowMessages("%016llx\t%016llx\t%ws\n",
                                 Modules[i].BaseAddress,
                                 Modules[i].Entrypoint,
                                 CommandLmConvertWow64CompatibilityPaths(Modules[i].FilePath).c_str());
                }
                else{
                    ShowMessages("%016llx\t%016llx\t%ws\n",
                                 Modules[i].BaseAddress,
                                 Modules[i].Entrypoint,
                                 Modules[i].FilePath);
                }
            }
            if (SearchModule != NULL){
                free(WcharBuff);
            }
        }
        else{
            ShowErrorMessage(ModuleCountRequest.Result);
            return FALSE;
        }
        free(ModuleDetailsRequest);
        return TRUE;
    }
    else{
        ShowErrorMessage(ModuleCountRequest.Result);
        return FALSE;
    }
}
BOOLEAN CommandLmShowKernelModeModule(const char * SearchModule){
    NTSTATUS             Status = STATUS_UNSUCCESSFUL;
    PRTL_PROCESS_MODULES ModulesInfo;
    ULONG                SysModuleInfoBufferSize = 0;
    string               SearchModuleString;
    Status = NtQuerySystemInformation(SystemModuleInformation, NULL, NULL, &SysModuleInfoBufferSize);
    ModulesInfo = (PRTL_PROCESS_MODULES)VirtualAlloc(
        NULL,
        SysModuleInfoBufferSize,
        MEM_COMMIT | MEM_RESERVE,
        PAGE_READWRITE);
    if (!ModulesInfo){
        ShowMessages("err, unable to allocate memory for module list (%x)\n",
                     GetLastError());
        return FALSE;
    }
    Status = NtQuerySystemInformation(SystemModuleInformation,
                                      ModulesInfo,
                                      SysModuleInfoBufferSize,
                                      NULL);
    if (!NT_SUCCESS(Status)){
        ShowMessages("err, unable to query module list (%x)\n", Status);
        VirtualFree(ModulesInfo, 0, MEM_RELEASE);
        return FALSE;
    }
    if (SearchModule != NULL){
        SearchModuleString.assign(SearchModule, strlen(SearchModule));
    }
    ShowMessages("kernel mode\n");
    ShowMessages("start\t\t\tsize\tname\t\t\t\tpath\n\n");
    for (ULONG i = 0; i < ModulesInfo->NumberOfModules; i++){
        RTL_PROCESS_MODULE_INFORMATION * CurrentModule = &ModulesInfo->Modules[i];
        if (SearchModule != NULL){
            std::string FullPathName((char *)CurrentModule->FullPathName);
            if (FindCaseInsensitive(FullPathName, SearchModuleString, 0) == std::string::npos){
                continue;
            }
        }
        ShowMessages("%s\t", SeparateTo64BitValue((UINT64)CurrentModule->ImageBase).c_str());
        ShowMessages("%x\t", CurrentModule->ImageSize);
        auto   PathName    = CurrentModule->FullPathName + CurrentModule->OffsetToFileName;
        UINT32 PathNameLen = (UINT32)strlen((const char *)PathName);
        ShowMessages("%s\t", PathName);
        if (PathNameLen >= 24){
        }
        else if (PathNameLen >= 16){
            ShowMessages("\t");
        }
        else if (PathNameLen >= 8){
            ShowMessages("\t\t");
        }
        else{
            ShowMessages("\t\t\t");
        }
        ShowMessages("%s\n", CurrentModule->FullPathName);
    }
    VirtualFree(ModulesInfo, 0, MEM_RELEASE);
    return TRUE;
}
VOID CommandLm(vector<string> SplitCommand, string Command){
    BOOLEAN SetPid                = FALSE;
    BOOLEAN SetSearchFilter       = FALSE;
    BOOLEAN SearchStringEntered   = FALSE;
    BOOLEAN OnlyShowKernelModules = FALSE;
    BOOLEAN OnlyShowUserModules   = FALSE;
    UINT32  TargetPid             = NULL;
    char    Search[MAX_PATH]      = {0};
    char *  SearchString          = NULL;
    for (auto Section : SplitCommand){
        if (!Section.compare("lm")){
            continue;
        }
        else if (!Section.compare("pid") && !SetPid){
            SetPid = TRUE;
        }
        else if (!Section.compare("m") && !SetSearchFilter){
            SetSearchFilter = TRUE;
        }
        else if (SetPid){
            if (!ConvertStringToUInt32(Section, &TargetPid)){
                ShowMessages("err, couldn't resolve error at '%s'\n\n",
                             Section.c_str());
                CommandLmHelp();
                return;
            }
            SetPid = FALSE;
        }
        else if (SetSearchFilter){
            if (Section.length() >= MAX_PATH){
                ShowMessages("err, string is too large for search, please enter "
                             "smaller string\n");
                return;
            }
            SearchStringEntered = TRUE;
            strcpy(Search, Section.c_str());
            SetSearchFilter = FALSE;
        }
        else if (!Section.compare("km")){
            if (OnlyShowUserModules){
                ShowMessages("err, you cannot use both 'um', and 'km', by default "
                             "HyperDbg shows both user-mode and kernel-mode modules\n");
                return;
            }
            OnlyShowKernelModules = TRUE;
        }
        else if (!Section.compare("um")){
            if (OnlyShowKernelModules){
                ShowMessages("err, you cannot use both 'um', and 'km', by default "
                             "HyperDbg shows both user-mode and kernel-mode modules\n");
                return;
            }
            OnlyShowUserModules = TRUE;
        }
        else{
            ShowMessages("err, couldn't resolve error at '%s'\n\n",
                         Section.c_str());
            CommandLmHelp();
            return;
        }
    }
    if (SetPid){
        ShowMessages("err, please enter a valid process id in hex format, "
                     "or if you want to use it in decimal format, add '0n' "
                     "prefix to the number\n");
        return;
    }
    if (SetSearchFilter){
        ShowMessages("err, please enter a valid string to search in modules\n");
        return;
    }
    if (SearchStringEntered){
        SearchString = Search;
    }
    if (!OnlyShowKernelModules){
        if (TargetPid != NULL){
            CommandLmShowUserModeModule(TargetPid, SearchString);
        }
        else if (g_ActiveProcessDebuggingState.IsActive){
            CommandLmShowUserModeModule(g_ActiveProcessDebuggingState.ProcessId, SearchString);
        }
        else{
            CommandLmShowUserModeModule(GetCurrentProcessId(), SearchString);
        }
    }
    if (!OnlyShowUserModules){
        if (!OnlyShowKernelModules){
            ShowMessages("\n==============================================================================\n\n");
        }
        CommandLmShowKernelModeModule(SearchString);
    }
}
