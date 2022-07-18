#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
VOID
CommandTestHelp() {
    ShowMessages(
        "test : tests essential features of HyperDbg in current machine.\n");
    ShowMessages("syntax : \ttest [Task (string)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : test\n");
    ShowMessages("\t\te.g : test query\n");
}

BOOLEAN
CommandTestPerformKernelTestsIoctl() {
    BOOL                          Status;
    ULONG                         ReturnedLength;
    DEBUGGER_PERFORM_KERNEL_TESTS KernelTestRequest = {0};
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
    Status = DeviceIoControl(
        g_DeviceHandle,                       // Handle to device
        IOCTL_PERFROM_KERNEL_SIDE_TESTS,      // IO Control code
        &KernelTestRequest,                   // Input Buffer to driver.
        SIZEOF_DEBUGGER_PERFORM_KERNEL_TESTS, // Input buffer length
        &KernelTestRequest,                   // Output Buffer from driver.
        SIZEOF_DEBUGGER_PERFORM_KERNEL_TESTS, // Length of output buffer in
        &ReturnedLength,                      // Bytes placed in buffer.
        NULL                                  // synchronous call
    );
    if (!Status) {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return FALSE;
    }
    if (KernelTestRequest.KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
        return TRUE;
    } else {
        ShowErrorMessage(KernelTestRequest.KernelStatus);
        return FALSE;
    }
}

BOOLEAN
CommandTestPerformTest(PDEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION KernelSideInformation, UINT32 KernelSideInformationSize) {
    BOOLEAN ResultOfTest = FALSE;
    HANDLE  PipeHandle;
    HANDLE  ThreadHandle;
    HANDLE  ProcessHandle;
    UINT32  ReadBytes;
    BOOLEAN SentMessageResult;
    CHAR *  Buffer = {0};
    Buffer         = (CHAR *)malloc(TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    RtlZeroMemory(Buffer, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    if (!CreateProcessAndOpenPipeConnection(
            KernelSideInformation,
            KernelSideInformationSize,
            &PipeHandle,
            &ThreadHandle,
            &ProcessHandle)) {
        ShowMessages("err, enable to connect to the test process\n");
        free(Buffer);
        return FALSE;
    }
WaitForResponse:
    RtlZeroMemory(Buffer, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    ReadBytes =
        NamedPipeServerReadClientMessage(PipeHandle, (char *)Buffer, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    if (!ReadBytes) {
        free(Buffer);
        return FALSE;
    }
    if (strcmp(Buffer, "perform-kernel-test") == 0) {
        CommandTestPerformKernelTestsIoctl();
        goto WaitForResponse;
    } else if (strcmp(Buffer, "success") == 0) {
        ResultOfTest = TRUE;
    } else if (Buffer[0] == 'c' &&
               Buffer[1] == 'm' &&
               Buffer[2] == 'd' &&
               Buffer[3] == ':') {
        ShowMessages("command is : %s\n", &Buffer[4]);
        HyperDbgInterpreter(&Buffer[4]);
        goto WaitForResponse;
    } else {
        ResultOfTest = FALSE;
    }
    CloseProcessAndClosePipeConnection(PipeHandle, ThreadHandle, ProcessHandle);
    free(Buffer);
    return ResultOfTest;
}

VOID
CommandTestInVmiMode() {
    BOOL  Status;
    ULONG ReturnedLength;
    PDEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION
    KernelSideTestInformationRequestArray;
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
    KernelSideTestInformationRequestArray = (DEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION *)malloc(TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    RtlZeroMemory(KernelSideTestInformationRequestArray, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    Status = DeviceIoControl(
        g_DeviceHandle,                                   // Handle to device
        IOCTL_SEND_GET_KERNEL_SIDE_TEST_INFORMATION,      // IO Control code
        KernelSideTestInformationRequestArray,            // Input Buffer to driver.
        SIZEOF_DEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION, // Input buffer
        KernelSideTestInformationRequestArray,            // Output Buffer from driver.
        TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE,         // Length
        &ReturnedLength,                                  // Bytes placed in buffer.
        NULL                                              // synchronous call
    );
    if (!Status) {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return;
    }
    if (CommandTestPerformTest(KernelSideTestInformationRequestArray, ReturnedLength)) {
        ShowMessages("all the tests were successful :)\n");
    } else {
        ShowMessages("all or at least one of the tests failed :(\n");
    }
    free(KernelSideTestInformationRequestArray);
}

VOID
CommandTestQueryState() {
    if (!g_IsSerialConnectedToRemoteDebuggee) {
        ShowMessages("err, query state of the debuggee is only possible when you connected "
                     "in debugger mode\n");
        return;
    }
    KdSendTestQueryPacketToDebuggee(TEST_QUERY_HALTING_CORE_STATUS);
}

VOID
CommandTest(vector<string> SplittedCommand, string Command) {
    if (SplittedCommand.size() == 1) {
        CommandTestInVmiMode();
    } else if (SplittedCommand.size() == 2 && !SplittedCommand.at(1).compare("query")) {
        CommandTestQueryState();
    } else {
        ShowMessages("incorrect use of 'test'\n\n");
        CommandTestHelp();
        return;
    }
}
