package debugging-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\test.cpp.back

type (
Test interface{
CommandTestHelp()(ok bool)//col:35
CommandTestPerformKernelTestsIoctl()(ok bool)//col:90
CommandTestPerformTest()(ok bool)//col:197
CommandTestInVmiMode()(ok bool)//col:263
CommandTestQueryState()(ok bool)//col:284
CommandTest()(ok bool)//col:316
}
)

func NewTest() { return & test{} }

func (t *test)CommandTestHelp()(ok bool){//col:35
/*CommandTestHelp()
{
    ShowMessages(
        "test : tests essential features of HyperDbg in current machine.\n");
    ShowMessages("syntax : \ttest [Task (string)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : test\n");
    ShowMessages("\t\te.g : test query\n");
}*/
return true
}

func (t *test)CommandTestPerformKernelTestsIoctl()(ok bool){//col:90
/*CommandTestPerformKernelTestsIoctl()
{
    BOOL                          Status;
    ULONG                         ReturnedLength;
    DEBUGGER_PERFORM_KERNEL_TESTS KernelTestRequest = {0};
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
    Status = DeviceIoControl(
    );
    if (!Status)
    {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return FALSE;
    }
    if (KernelTestRequest.KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFULL)
    {
        return TRUE;
    }
    else
    {
        ShowErrorMessage(KernelTestRequest.KernelStatus);
        return FALSE;
    }
}*/
return true
}

func (t *test)CommandTestPerformTest()(ok bool){//col:197
/*CommandTestPerformTest(PDEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION KernelSideInformation, UINT32 KernelSideInformationSize)
{
    BOOLEAN ResultOfTest = FALSE;
    HANDLE  PipeHandle;
    HANDLE  ThreadHandle;
    HANDLE  ProcessHandle;
    UINT32  ReadBytes;
    BOOLEAN SentMessageResult;
    CHAR *  Buffer = {0};
    Buffer = (CHAR *)malloc(TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    RtlZeroMemory(Buffer, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    if (!CreateProcessAndOpenPipeConnection(
            KernelSideInformation,
            KernelSideInformationSize,
            &PipeHandle,
            &ThreadHandle,
            &ProcessHandle))
    {
        ShowMessages("err, enable to connect to the test process\n");
        free(Buffer);
        return FALSE;
    }
WaitForResponse:
    RtlZeroMemory(Buffer, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    ReadBytes =
        NamedPipeServerReadClientMessage(PipeHandle, (char *)Buffer, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    if (!ReadBytes)
    {
        free(Buffer);
        return FALSE;
    }
    if (strcmp(Buffer, "perform-kernel-test") == 0)
    {
        CommandTestPerformKernelTestsIoctl();
        goto WaitForResponse;
    }
    else if (strcmp(Buffer, "success") == 0)
    {
        ResultOfTest = TRUE;
    }
    else if (Buffer[0] == 'c' &&
             Buffer[1] == 'm' &&
             Buffer[2] == 'd' &&
             Buffer[3] == ':')
    {
        ShowMessages("command is : %s\n", &Buffer[4]);
        HyperDbgInterpreter(&Buffer[4]);
        goto WaitForResponse;
    }
    else
    {
        ResultOfTest = FALSE;
    }
    CloseProcessAndClosePipeConnection(PipeHandle, ThreadHandle, ProcessHandle);
    free(Buffer);
    return ResultOfTest;
}*/
return true
}

func (t *test)CommandTestInVmiMode()(ok bool){//col:263
/*CommandTestInVmiMode()
{
    BOOL  Status;
    ULONG ReturnedLength;
    PDEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION
    KernelSideTestInformationRequestArray;
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
    KernelSideTestInformationRequestArray = (DEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION *)malloc(TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    RtlZeroMemory(KernelSideTestInformationRequestArray, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    Status = DeviceIoControl(
    );
    if (!Status)
    {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return;
    }
    if (CommandTestPerformTest(KernelSideTestInformationRequestArray, ReturnedLength))
    {
        ShowMessages("all the tests were successful :)\n");
    }
    else
    {
        ShowMessages("all or at least one of the tests failed :(\n");
    }
    free(KernelSideTestInformationRequestArray);
}*/
return true
}

func (t *test)CommandTestQueryState()(ok bool){//col:284
/*CommandTestQueryState()
{
    if (!g_IsSerialConnectedToRemoteDebuggee)
    {
        ShowMessages("err, query state of the debuggee is only possible when you connected "
                     "in debugger mode\n");
        return;
    }
    KdSendTestQueryPacketToDebuggee(TEST_QUERY_HALTING_CORE_STATUS);
}*/
return true
}

func (t *test)CommandTest()(ok bool){//col:316
/*CommandTest(vector<string> SplittedCommand, string Command)
{
    if (SplittedCommand.size() == 1)
    {
        CommandTestInVmiMode();
    }
    else if (SplittedCommand.size() == 2 && !SplittedCommand.at(1).compare("query"))
    {
        CommandTestQueryState();
    }
    else
    {
        ShowMessages("incorrect use of 'test'\n\n");
        CommandTestHelp();
        return;
    }
}*/
return true
}


