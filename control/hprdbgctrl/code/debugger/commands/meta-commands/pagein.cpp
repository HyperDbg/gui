#include "pch.h"
extern BOOLEAN                  g_IsSerialConnectedToRemoteDebuggee;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;
VOID CommandPageinHelp(){
    ShowMessages(".pagein : brings the page in, making it available in the RAM.\n\n");
    ShowMessages("syntax : \t.pagein [Mode (string)] [l Length (hex)]\n");
    ShowMessages("syntax : \t.pagein [Mode (string)] [VirtualAddress (hex)] [l Length (hex)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : .pagein fffff801deadbeef\n");
    ShowMessages("\t\te.g : .pagein 00007ff8349f2224 l 1a000\n");
    ShowMessages("\t\te.g : .pagein u 00007ff8349f2224\n");
    ShowMessages("\t\te.g : .pagein w 00007ff8349f2224\n");
    ShowMessages("\t\te.g : .pagein f 00007ff8349f2224\n");
    ShowMessages("\t\te.g : .pagein pw 00007ff8349f2224\n");
    ShowMessages("\t\te.g : .pagein wu 00007ff8349f2224\n");
    ShowMessages("\t\te.g : .pagein wu 00007ff8349f2224 l 6000\n");
    ShowMessages("\t\te.g : .pagein pf @rax\n");
    ShowMessages("\t\te.g : .pagein uf @rip+@rcx\n");
    ShowMessages("\t\te.g : .pagein pwu @rax+5\n");
    ShowMessages("\t\te.g : .pagein pwu @rax l 2000\n");
    ShowMessages("\n");
    ShowMessages("valid mode formats: \n");
    ShowMessages("\tp : present\n");
    ShowMessages("\tw : write\n");
    ShowMessages("\tu : user\n");
    ShowMessages("\tf : fetch\n");
    ShowMessages("\tk : protection key\n");
    ShowMessages("\ts : shadow stack\n");
    ShowMessages("\th : hlat\n");
    ShowMessages("\tg : sgx\n");
    ShowMessages("\n");
    ShowMessages("common page-fault codes: \n");
    ShowMessages("\t0x0:  (default)\n");
    ShowMessages("\t0x2:  w (write access fault)\n");
    ShowMessages("\t0x3:  pw (present, write access fault)\n");
    ShowMessages("\t0x4:  u (user access fault)\n");
    ShowMessages("\t0x6:  wu (write, user access fault)\n");
    ShowMessages("\t0x7:  pwu (present, write, user access fault)\n");
    ShowMessages("\t0x10: f (fetch instruction fault)\n");
    ShowMessages("\t0x11: pf (present, fetch instruction fault)\n");
    ShowMessages("\t0x14: uf (user, fetch instruction fault)\n");
}
BOOLEAN CommandPageinCheckAndInterpretModeString(const std::string &    ModeString,
                                         PAGE_FAULT_EXCEPTION * PageFaultErrorCode){
    std::unordered_set<char> AllowedChars = {'p', 'w', 'u', 'f', 'k', 's', 'h', 'g'};
    std::unordered_set<char> FoundChars;
    for (char c : ModeString){
        if (AllowedChars.count(c) == 0){
            return FALSE;
        }
        if (FoundChars.count(c) > 0){
            return false;
        }
        FoundChars.insert(c);
    }
    for (char c : ModeString){
        if (c == 'p'){
            PageFaultErrorCode->Present = TRUE;
        }
        else if (c == 'w'){
            PageFaultErrorCode->Write = TRUE;
        }
        else if (c == 'u'){
            PageFaultErrorCode->UserModeAccess = TRUE;
        }
        else if (c == 'f'){
            PageFaultErrorCode->Execute = TRUE;
        }
        else if (c == 'k'){
            PageFaultErrorCode->ProtectionKeyViolation = TRUE;
        }
        else if (c == 's'){
            PageFaultErrorCode->ShadowStack = TRUE;
        }
        else if (c == 'h'){
            PageFaultErrorCode->Hlat = TRUE;
        }
        else if (c == 'g'){
            PageFaultErrorCode->Sgx = TRUE;
        }
        else{
            return FALSE;
        }
    }
    return TRUE;
}
VOID CommandPageinRequest(UINT64               TargetVirtualAddrFrom,
                     UINT64               TargetVirtualAddrTo,
                     PAGE_FAULT_EXCEPTION PageFaultErrorCode,
                     UINT32               Pid){
    BOOL                     Status;
    ULONG                    ReturnedLength;
    DEBUGGER_PAGE_IN_REQUEST PageFaultRequest = {0};
    PageFaultRequest.VirtualAddressFrom = TargetVirtualAddrFrom;
    PageFaultRequest.VirtualAddressTo   = TargetVirtualAddrTo;
    PageFaultRequest.ProcessId          = Pid; 
    if (g_IsSerialConnectedToRemoteDebuggee){
        if (PageFaultRequest.ProcessId != 0){
            ShowMessages(ASSERT_MESSAGE_CANNOT_SPECIFY_PID);
            return;
        }
        KdSendPageinPacketToDebuggee(&PageFaultRequest);
    }
    else{
        AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
        ShowMessages("the '.pagein' command can be used ONLY in the debugger mode, "
                     "it is not yet supported in VMI mode\n");
        return;
        if (Pid == 0){
            Pid                        = GetCurrentProcessId();
            PageFaultRequest.ProcessId = Pid;
        }
        Status = DeviceIoControl(
            g_DeviceHandle,                  
            IOCTL_DEBUGGER_BRING_PAGES_IN,   
            &PageFaultRequest,               
            SIZEOF_DEBUGGER_PAGE_IN_REQUEST, 
            &PageFaultRequest,               
            SIZEOF_DEBUGGER_PAGE_IN_REQUEST, 
            &ReturnedLength,                 
            NULL                             
        );
        if (!Status){
            ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
            return;
        }
        if (PageFaultRequest.KernelStatus != DEBUGGER_OPERATION_WAS_SUCCESSFUL){
            ShowErrorMessage(PageFaultRequest.KernelStatus);
            return;
        }
        ShowMessages("page-fault is delivered\n");
    }
}
VOID CommandPagein(vector<string> SplitCommand, string Command){
    UINT32               Pid               = 0;
    UINT64               Length            = 0;
    UINT64               TargetAddressFrom = NULL;
    UINT64               TargetAddressTo   = NULL;
    BOOLEAN              IsNextProcessId   = FALSE;
    BOOLEAN              IsFirstCommand    = TRUE;
    BOOLEAN              IsNextLength      = FALSE;
    vector<string>       SplitCommandCaseSensitive {Split(Command, ' ')};
    UINT32               IndexInCommandCaseSensitive = 0;
    PAGE_FAULT_EXCEPTION PageFaultErrorCode          = {0};
    if (g_ActiveProcessDebuggingState.IsActive){
        Pid = g_ActiveProcessDebuggingState.ProcessId;
    }
    if (SplitCommand.size() == 1){
        ShowMessages("incorrect use of the '.pagein' command\n\n");
        CommandPageinHelp();
        return;
    }
    for (auto Section : SplitCommand){
        IndexInCommandCaseSensitive++;
        if (IsFirstCommand){
            IsFirstCommand = FALSE;
            continue;
        }
        if (IsNextProcessId == TRUE){
            if (!ConvertStringToUInt32(Section, &Pid)){
                ShowMessages("err, you should enter a valid process id\n\n");
                return;
            }
            IsNextProcessId = FALSE;
            continue;
        }
        if (IsNextLength == TRUE){
            if (!SymbolConvertNameOrExprToAddress(Section, &Length)){
                ShowMessages("err, you should enter a valid length\n\n");
                return;
            }
            IsNextLength = FALSE;
            continue;
        }
        if (!Section.compare("l")){
            IsNextLength = TRUE;
            continue;
        }
        if (CommandPageinCheckAndInterpretModeString(Section, &PageFaultErrorCode)){
            continue;
        }
        else if (TargetAddressFrom == 0){
            if (!SymbolConvertNameOrExprToAddress(SplitCommandCaseSensitive.at(IndexInCommandCaseSensitive - 1),
                                                  &TargetAddressFrom)){
                ShowMessages("err, couldn't resolve error at '%s'\n",
                             SplitCommandCaseSensitive.at(IndexInCommandCaseSensitive - 1).c_str());
                return;
            }
        }
        else{
            ShowMessages("err, incorrect use of the '.pagein' command\n\n");
            CommandPageinHelp();
            return;
        }
    }
    if (!TargetAddressFrom){
        ShowMessages("err, please enter a valid address\n\n");
        return;
    }
    if (IsNextLength || IsNextProcessId){
        ShowMessages("incorrect use of the '.pagein' command\n\n");
        CommandPageinHelp();
        return;
    }
    if (Length == 0){
        TargetAddressTo = TargetAddressFrom;
    }
    else{
        TargetAddressTo = TargetAddressFrom + Length;
    }
    CommandPageinRequest(TargetAddressFrom,
                         TargetAddressTo,
                         PageFaultErrorCode,
                         Pid);
}
