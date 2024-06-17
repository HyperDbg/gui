#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
VOID CommandTestHelp(){
    ShowMessages(
        "test : tests essential features of HyperDbg in current machine.\n");
    ShowMessages("syntax : \ttest [Task (string)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : test\n");
    ShowMessages("\t\te.g : test query\n");
    ShowMessages("\t\te.g : test trap-status\n");
    ShowMessages("\t\te.g : test pool\n");
    ShowMessages("\t\te.g : test query\n");
    ShowMessages("\t\te.g : test breakpoint on\n");
    ShowMessages("\t\te.g : test breakpoint off\n");
    ShowMessages("\t\te.g : test trap on\n");
    ShowMessages("\t\te.g : test trap off\n");
}
BOOLEAN CommandTestPerformKernelTestsIoctl(){
    BOOL                          Status;
    ULONG                         ReturnedLength;
    DEBUGGER_PERFORM_KERNEL_TESTS KernelTestRequest = {0};
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
    Status = DeviceIoControl(
        g_DeviceHandle,                       
        IOCTL_PERFROM_KERNEL_SIDE_TESTS,      
        &KernelTestRequest,                   
        SIZEOF_DEBUGGER_PERFORM_KERNEL_TESTS, 
        &KernelTestRequest,                   
        SIZEOF_DEBUGGER_PERFORM_KERNEL_TESTS, 
        &ReturnedLength,                      
        NULL                                  
    );
    if (!Status){
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return FALSE;
    }
    if (KernelTestRequest.KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFUL){
        return TRUE;
    }
    else{
        ShowErrorMessage(KernelTestRequest.KernelStatus);
        return FALSE;
    }
}
BOOLEAN CommandTestPerformTest(){
    BOOLEAN ResultOfTest = FALSE;
    HANDLE  PipeHandle;
    HANDLE  ThreadHandle;
    HANDLE  ProcessHandle;
    UINT32  ReadBytes;
    CHAR *  Buffer = NULL;
    Buffer = (CHAR *)malloc(TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    if (!Buffer){
        ShowMessages("err, enable allocate communication buffer\n");
        return FALSE;
    }
    RtlZeroMemory(Buffer, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    if (!CreateProcessAndOpenPipeConnection(&PipeHandle,
                                            &ThreadHandle,
                                            &ProcessHandle)){
        ShowMessages("err, enable to connect to the test process\n");
        free(Buffer);
        return FALSE;
    }
SendCommandAndWaitForResponse:
    CHAR TestCommand[] = "this is a test command";
    BOOLEAN SentMessageResult = NamedPipeServerSendMessageToClient(
        PipeHandle,
        TestCommand,
        (UINT32)strlen(TestCommand) + 1);
    if (!SentMessageResult){
        return FALSE;
    }
    RtlZeroMemory(Buffer, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    ReadBytes =
        NamedPipeServerReadClientMessage(PipeHandle, (char *)Buffer, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    if (!ReadBytes){
        free(Buffer);
        return FALSE;
    }
    goto SendCommandAndWaitForResponse;
    CloseProcessAndClosePipeConnection(PipeHandle, ThreadHandle, ProcessHandle);
    free(Buffer);
    return ResultOfTest;
}
VOID CommandTestQueryState(){
    if (!g_IsSerialConnectedToRemoteDebuggee){
        ShowMessages("err, query state of the debuggee is only possible when you connected "
                     "in debugger mode\n");
        return;
    }
    KdSendTestQueryPacketToDebuggee(TEST_QUERY_HALTING_CORE_STATUS);
}
VOID CommandTestQueryTrapState(){
    if (!g_IsSerialConnectedToRemoteDebuggee){
        ShowMessages("err, query state of the debuggee is only possible when you connected "
                     "in debugger mode\n");
        return;
    }
    KdSendTestQueryPacketToDebuggee(TEST_QUERY_TRAP_STATE);
}
VOID CommandTestQueryPreAllocPoolsState(){
    if (!g_IsSerialConnectedToRemoteDebuggee){
        ShowMessages("err, query state of the debuggee is only possible when you connected "
                     "in debugger mode\n");
        return;
    }
    KdSendTestQueryPacketToDebuggee(TEST_QUERY_PREALLOCATED_POOL_STATE);
}
VOID CommandTestSetTargetTaskToHaltedCores(BOOLEAN Synchronous){
    if (!g_IsSerialConnectedToRemoteDebuggee){
        ShowMessages("err, query state of the debuggee is only possible when you connected "
                     "in debugger mode\n");
        return;
    }
    KdSendTestQueryPacketToDebuggee(Synchronous ? TEST_SETTING_TARGET_TASKS_ON_HALTED_CORES_SYNCHRONOUS : TEST_SETTING_TARGET_TASKS_ON_HALTED_CORES_ASYNCHRONOUS);
}
VOID CommandTestSetTargetTaskToTargetCore(UINT32 CoreNumber){
    if (!g_IsSerialConnectedToRemoteDebuggee){
        ShowMessages("err, query state of the debuggee is only possible when you connected "
                     "in debugger mode\n");
        return;
    }
    KdSendTestQueryPacketWithContextToDebuggee(TEST_SETTING_TARGET_TASKS_ON_TARGET_HALTED_CORES, (UINT64)CoreNumber);
}
VOID CommandTestSetBreakpointState(BOOLEAN State){
    if (!g_IsSerialConnectedToRemoteDebuggee){
        ShowMessages("err, query state of the debuggee is only possible when you connected "
                     "in debugger mode\n");
        return;
    }
    if (State){
        KdSendTestQueryPacketToDebuggee(TEST_BREAKPOINT_TURN_ON_BPS);
    }
    else{
        KdSendTestQueryPacketToDebuggee(TEST_BREAKPOINT_TURN_OFF_BPS);
    }
}
VOID CommandTestSetDebugBreakState(BOOLEAN State){
    if (!g_IsSerialConnectedToRemoteDebuggee){
        ShowMessages("err, query state of the debuggee is only possible when you connected "
                     "in debugger mode\n");
        return;
    }
    if (State){
        KdSendTestQueryPacketToDebuggee(TEST_BREAKPOINT_TURN_ON_DBS);
    }
    else{
        KdSendTestQueryPacketToDebuggee(TEST_BREAKPOINT_TURN_OFF_DBS);
    }
}
VOID CommandTest(vector<string> SplitCommand, string Command){
    UINT64 Context = NULL;
    if (SplitCommand.size() == 1){
        CommandTestPerformTest();
    }
    else if (SplitCommand.size() == 2 && !SplitCommand.at(1).compare("query")){
        CommandTestQueryState();
    }
    else if (SplitCommand.size() == 2 && !SplitCommand.at(1).compare("trap-status")){
        CommandTestQueryTrapState();
    }
    else if (SplitCommand.size() == 2 && !SplitCommand.at(1).compare("pool")){
        CommandTestQueryPreAllocPoolsState();
    }
    else if (SplitCommand.size() == 2 && !SplitCommand.at(1).compare("sync-task")){
        CommandTestSetTargetTaskToHaltedCores(TRUE);
    }
    else if (SplitCommand.size() == 2 && !SplitCommand.at(1).compare("async-task")){
        CommandTestSetTargetTaskToHaltedCores(FALSE);
    }
    else if (SplitCommand.size() == 3 && !SplitCommand.at(1).compare("target-core-task")){
        if (!ConvertStringToUInt64(SplitCommand.at(2), &Context)){
            ShowMessages("err, you should enter a valid hex number as the core id\n\n");
            return;
        }
        CommandTestSetTargetTaskToTargetCore((UINT32)Context);
    }
    else if (SplitCommand.size() == 3 && !SplitCommand.at(1).compare("breakpoint")){
        if (!SplitCommand.at(2).compare("on")){
            CommandTestSetBreakpointState(TRUE);
        }
        else if (!SplitCommand.at(2).compare("off")){
            CommandTestSetBreakpointState(FALSE);
        }
        else{
            ShowMessages("err, couldn't resolve error at '%s'\n\n", SplitCommand.at(2).c_str());
            return;
        }
    }
    else if (SplitCommand.size() == 3 && !SplitCommand.at(1).compare("trap")){
        if (!SplitCommand.at(2).compare("on")){
            CommandTestSetDebugBreakState(TRUE);
        }
        else if (!SplitCommand.at(2).compare("off")){
            CommandTestSetDebugBreakState(FALSE);
        }
        else{
            ShowMessages("err, couldn't resolve error at '%s'\n\n", SplitCommand.at(2).c_str());
            return;
        }
    }
    else{
        ShowMessages("incorrect use of the 'test'\n\n");
        CommandTestHelp();
        return;
    }
}
