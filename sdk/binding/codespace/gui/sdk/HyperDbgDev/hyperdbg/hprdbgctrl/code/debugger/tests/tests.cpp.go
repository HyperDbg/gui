package tests
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\tests\tests.cpp.back

type (
Tests interface{
SetupTestName()(ok bool)//col:29
CreateProcessAndOpenPipeConnection()(ok bool)//col:135
CloseProcessAndClosePipeConnection()(ok bool)//col:143
}
tests struct{}
)

func NewTests()Tests{ return & tests{} }

func (t *tests)SetupTestName()(ok bool){//col:29
/*SetupTestName(_Inout_updates_bytes_all_(BufferLength) PCHAR TestLocation,
              ULONG                                         BufferLength)
{
    HANDLE  fileHandle;
    DWORD   driverLocLen = 0;
    HMODULE ProcHandle   = GetModuleHandle(NULL);
    char *  Pos;
    GetModuleFileName(ProcHandle, TestLocation, BufferLength);
    Pos = strrchr(TestLocation, '\\');
    if (Pos != NULL)
    {
        *Pos = '\0';
    }
    if (FAILED(StringCbCat(TestLocation, BufferLength, "\\" TEST_PROCESS_NAME)))
    {
        return FALSE;
    }
    if ((fileHandle = CreateFile(TestLocation, GENERIC_READ, 0, NULL, OPEN_EXISTING, FILE_ATTRIBUTE_NORMAL, NULL)) ==
        INVALID_HANDLE_VALUE)
    {
        ShowMessages("%s.exe is not loaded.\n", TEST_PROCESS_NAME);
        return FALSE;
    }
    if (fileHandle)
    {
        CloseHandle(fileHandle);
    }
    return TRUE;
}*/
return true
}

func (t *tests)CreateProcessAndOpenPipeConnection()(ok bool){//col:135
/*CreateProcessAndOpenPipeConnection(PVOID   KernelInformation,
                                   UINT32  KernelInformationSize,
                                   PHANDLE ConnectionPipeHandle,
                                   PHANDLE ThreadHandle,
                                   PHANDLE ProcessHandle)
{
    HANDLE              PipeHandle;
    BOOLEAN             SentMessageResult;
    UINT32              ReadBytes;
    char *              BufferToRead;
    char *              BufferToSend;
    char                HandshakeBuffer[] = "Hello, Dear Test Process... Yes, I'm HyperDbg Debugger :)";
    PROCESS_INFORMATION ProcessInfo;
    STARTUPINFO         StartupInfo;
    char                CmdArgs[] = TEST_PROCESS_NAME " im-hyperdbg";
    PipeHandle = NamedPipeServerCreatePipe("\\\\.\\Pipe\\HyperDbgTests",
                                           TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE,
                                           TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    if (!PipeHandle)
    {
        return FALSE;
    }
    BufferToRead = (char *)malloc(TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    BufferToSend = (char *)malloc(TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    RtlZeroMemory(BufferToRead, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    RtlZeroMemory(BufferToSend, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    strcpy(BufferToSend, HandshakeBuffer);
    ZeroMemory(&StartupInfo, sizeof(StartupInfo));
    StartupInfo.cb = sizeof StartupInfo;
    if (!SetupTestName(g_TestLocation, sizeof(g_TestLocation)))
    {
        free(BufferToRead);
        free(BufferToSend);
        return FALSE;
    }
    if (CreateProcess(g_TestLocation, CmdArgs, NULL, NULL, FALSE, CREATE_NEW_CONSOLE, NULL, NULL, &StartupInfo, &ProcessInfo))
    {
        if (!NamedPipeServerWaitForClientConntection(PipeHandle))
        {
            free(BufferToRead);
            free(BufferToSend);
            return FALSE;
        }
        ReadBytes =
            NamedPipeServerReadClientMessage(PipeHandle, BufferToRead, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
        if (!ReadBytes)
        {
            free(BufferToRead);
            free(BufferToSend);
            return FALSE;
        }
        if (strcmp(BufferToRead, "Hey there, Are you HyperDbg?") == 0)
        {
            SentMessageResult = NamedPipeServerSendMessageToClient(
                PipeHandle,
                BufferToSend,
                strlen(BufferToSend) + 1);
            if (!SentMessageResult)
            {
                free(BufferToRead);
                free(BufferToSend);
                return FALSE;
            }
            RtlZeroMemory(BufferToRead, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
            ReadBytes = NamedPipeServerReadClientMessage(PipeHandle, BufferToRead, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
            if (!ReadBytes)
            {
                free(BufferToRead);
                free(BufferToSend);
                return FALSE;
            }
            if (strcmp(BufferToRead, "Wow! I miss you... Would you plz send me the "
                                     "kernel information?") == 0)
            {
                {
                    SentMessageResult = NamedPipeServerSendMessageToClient(
                        PipeHandle,
                        (char *)KernelInformation,
                        KernelInformationSize);
                    if (!SentMessageResult)
                    {
                        free(BufferToRead);
                        free(BufferToSend);
                        return FALSE;
                    }
                }
            }
            *ConnectionPipeHandle = PipeHandle;
            *ThreadHandle         = ProcessInfo.hThread;
            *ProcessHandle        = ProcessInfo.hProcess;
            free(BufferToRead);
            free(BufferToSend);
            return TRUE;
        }
        else
        {
            ShowMessages("the process could not be started\n");
            free(BufferToRead);
            free(BufferToSend);
            return FALSE;
        }
    }
    free(BufferToRead);
    free(BufferToSend);
    return FALSE;
}*/
return true
}

func (t *tests)CloseProcessAndClosePipeConnection()(ok bool){//col:143
/*CloseProcessAndClosePipeConnection(HANDLE ConnectionPipeHandle,
                                   HANDLE ThreadHandle,
                                   HANDLE ProcessHandle)
{
    NamedPipeServerCloseHandle(ConnectionPipeHandle);
    CloseHandle(ThreadHandle);
    CloseHandle(ProcessHandle);
}*/
return true
}



