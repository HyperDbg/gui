#include "pch.h"
VOID CommandPreallocHelp(){
    ShowMessages("prealloc : pre-allocates buffer for special purposes.\n\n");
    ShowMessages("syntax : \tprealloc  [Type (string)] [Count (hex)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : prealloc thread-interception 8\n");
    ShowMessages("\t\te.g : prealloc monitor 10\n");
    ShowMessages("\t\te.g : prealloc epthook 5\n");
    ShowMessages("\t\te.g : prealloc epthook2 3\n");
    ShowMessages("\t\te.g : prealloc regular-event 12\n");
    ShowMessages("\t\te.g : prealloc big-safe-buffert 1\n");
    ShowMessages("\n");
    ShowMessages("type of allocations:\n");
    ShowMessages("\tthread-interception: used for pre-allocations of the thread holders for the thread interception mechanism\n");
    ShowMessages("\tmonitor: used for pre-allocations of the '!monitor' EPT hooks\n");
    ShowMessages("\tepthook: used for pre-allocations of the '!epthook' EPT hooks\n");
    ShowMessages("\tepthook2: used for pre-allocations of the '!epthook2' EPT hooks\n");
    ShowMessages("\tregular-event: used for pre-allocations of regular instant events\n");
    ShowMessages("\tbig-event: used for pre-allocations of big instant events\n");
    ShowMessages("\tregular-safe-buffer: used for pre-allocations of the regular event safe buffers ($buffer) for instant events\n");
    ShowMessages("\tbig-safe-buffer: used for pre-allocations of the big event safe buffers ($buffer) for instant events\n");
}
VOID CommandPrealloc(vector<string> SplitCommand, string Command){
    BOOL                      Status;
    ULONG                     ReturnedLength;
    UINT64                    Count;
    DEBUGGER_PREALLOC_COMMAND PreallocRequest = {0};
    if (SplitCommand.size() != 3){
        ShowMessages("incorrect use of the 'prealloc'\n\n");
        CommandPreallocHelp();
        return;
    }
    if (!SplitCommand.at(1).compare("thread-interception")){
        PreallocRequest.Type = DEBUGGER_PREALLOC_COMMAND_TYPE_THREAD_INTERCEPTION;
    }
    else if (!SplitCommand.at(1).compare("monitor") || !SplitCommand.at(1).compare("!monitor")){
        PreallocRequest.Type = DEBUGGER_PREALLOC_COMMAND_TYPE_MONITOR;
    }
    else if (!SplitCommand.at(1).compare("epthook") || !SplitCommand.at(1).compare("!epthook")){
        PreallocRequest.Type = DEBUGGER_PREALLOC_COMMAND_TYPE_EPTHOOK;
    }
    else if (!SplitCommand.at(1).compare("epthook2") || !SplitCommand.at(1).compare("!epthook2")){
        PreallocRequest.Type = DEBUGGER_PREALLOC_COMMAND_TYPE_EPTHOOK2;
    }
    else if (!SplitCommand.at(1).compare("regular-event")){
        PreallocRequest.Type = DEBUGGER_PREALLOC_COMMAND_TYPE_REGULAR_EVENT;
    }
    else if (!SplitCommand.at(1).compare("big-event")){
        PreallocRequest.Type = DEBUGGER_PREALLOC_COMMAND_TYPE_BIG_EVENT;
    }
    else if (!SplitCommand.at(1).compare("regular-safe-buffer")){
        PreallocRequest.Type = DEBUGGER_PREALLOC_COMMAND_TYPE_REGULAR_SAFE_BUFFER;
    }
    else if (!SplitCommand.at(1).compare("big-safe-buffer")){
        PreallocRequest.Type = DEBUGGER_PREALLOC_COMMAND_TYPE_BIG_SAFE_BUFFER;
    }
    else{
        ShowMessages("err, couldn't resolve error at '%s'\n",
                     SplitCommand.at(1).c_str());
        return;
    }
    if (!SymbolConvertNameOrExprToAddress(SplitCommand.at(2), &Count)){
        ShowMessages("err, couldn't resolve error at '%s'\n",
                     SplitCommand.at(2).c_str());
        return;
    }
    PreallocRequest.Count = (UINT32)Count;
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
    Status = DeviceIoControl(
        g_DeviceHandle,                    
        IOCTL_RESERVE_PRE_ALLOCATED_POOLS, 
        &PreallocRequest,                  
        SIZEOF_DEBUGGER_PREALLOC_COMMAND,  
        &PreallocRequest,                  
        SIZEOF_DEBUGGER_PREALLOC_COMMAND,  
        &ReturnedLength,                   
        NULL                               
    );
    if (!Status){
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return;
    }
    if (PreallocRequest.KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFUL){
        ShowMessages("the requested pools are allocated and reserved\n");
    }
    else{
        ShowErrorMessage(PreallocRequest.KernelStatus);
    }
}
