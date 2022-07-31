package tests
//back\HyperDbgDev\hyperdbg\hyperdbg-test\code\tests\namedpipe.cpp.back

type (
Namedpipe interface{
NamedPipeServerCreatePipe()(ok bool)//col:49
NamedPipeServerWaitForClientConntection()(ok bool)//col:77
NamedPipeServerReadClientMessage()(ok bool)//col:120
NamedPipeServerSendMessageToClient()(ok bool)//col:147
NamedPipeServerCloseHandle()(ok bool)//col:159
NamedPipeClientCreatePipe()(ok bool)//col:214
NamedPipeClientSendMessage()(ok bool)//col:262
NamedPipeClientReadMessage()(ok bool)//col:293
NamedPipeClientClosePipe()(ok bool)//col:305
NamedPipeServerExample()(ok bool)//col:376
NamedPipeClientExample()(ok bool)//col:434
}
)

func NewNamedpipe() { return & namedpipe{} }

func (n *namedpipe)NamedPipeServerCreatePipe()(ok bool){//col:49
/*NamedPipeServerCreatePipe(LPCSTR PipeName, UINT32 OutputBufferSize, UINT32 InputBufferSize)
{
    HANDLE hPipe;
    if (INVALID_HANDLE_VALUE == hPipe)
    {
        printf("err, occurred while creating the pipe (%x)\n",
               GetLastError());
        return NULL;
    }
    return hPipe;
}*/
return true
}

func (n *namedpipe)NamedPipeServerWaitForClientConntection()(ok bool){//col:77
/*NamedPipeServerWaitForClientConntection(HANDLE PipeHandle)
{
    BOOL bClientConnected = ConnectNamedPipe(PipeHandle, NULL);
    if (FALSE == bClientConnected)
    {
        printf("err, occurred while connecting to the client (%x)\n",
               GetLastError());
        CloseHandle(PipeHandle);
        return FALSE;
    }
    return TRUE;
}*/
return true
}

func (n *namedpipe)NamedPipeServerReadClientMessage()(ok bool){//col:120
/*NamedPipeServerReadClientMessage(HANDLE PipeHandle, char * BufferToSave, int MaximumReadBufferLength)
{
    DWORD cbBytes;
    if ((!bResult) || (0 == cbBytes))
    {
        printf("err, occurred while reading from the client (%x)\n",
               GetLastError());
        CloseHandle(PipeHandle);
        return 0;
    }
    return cbBytes;
}*/
return true
}

func (n *namedpipe)NamedPipeServerSendMessageToClient()(ok bool){//col:147
/*NamedPipeServerSendMessageToClient(HANDLE PipeHandle,
                                   char * BufferToSend,
                                   int    BufferSize)
{
    DWORD cbBytes;
    BOOLEAN bResult =
    if ((!bResult) || (BufferSize != cbBytes))
    {
        printf("Error occurred while writing to the client (%x)\n",
               GetLastError());
        CloseHandle(PipeHandle);
        return FALSE;
    }
    return TRUE;
}*/
return true
}

func (n *namedpipe)NamedPipeServerCloseHandle()(ok bool){//col:159
/*NamedPipeServerCloseHandle(HANDLE PipeHandle)
{
    CloseHandle(PipeHandle);
}*/
return true
}

func (n *namedpipe)NamedPipeClientCreatePipe()(ok bool){//col:214
/*NamedPipeClientCreatePipe(LPCSTR PipeName)
{
    HANDLE hPipe;
                            GENERIC_WRITE,
    if (INVALID_HANDLE_VALUE == hPipe)
    {
        printf("err, occurred while connecting to the server (%x)\n",
               GetLastError());
        return NULL;
    }
    else
    {
        return hPipe;
    }
}*/
return true
}

func (n *namedpipe)NamedPipeClientSendMessage()(ok bool){//col:262
/*NamedPipeClientSendMessage(HANDLE PipeHandle, char * BufferToSend, int BufferSize)
{
    DWORD cbBytes;
    BOOL bResult =
    if ((!bResult) || (BufferSize != cbBytes))
    {
        printf("err, occurred while writing to the server (%x)\n",
               GetLastError());
        CloseHandle(PipeHandle);
        CloseHandle(PipeHandle);
        return FALSE;
    }
    else
    {
        return TRUE;
    }
}*/
return true
}

func (n *namedpipe)NamedPipeClientReadMessage()(ok bool){//col:293
/*NamedPipeClientReadMessage(HANDLE PipeHandle, char * BufferToRead, int MaximumSizeOfBuffer)
{
    DWORD cbBytes;
    if ((!bResult) || (0 == cbBytes))
    {
        printf("err, occurred while reading from the server (%x)\n",
               GetLastError());
        CloseHandle(PipeHandle);
    }
    return cbBytes;
}*/
return true
}

func (n *namedpipe)NamedPipeClientClosePipe()(ok bool){//col:305
/*NamedPipeClientClosePipe(HANDLE PipeHandle)
{
    CloseHandle(PipeHandle);
}*/
return true
}

func (n *namedpipe)NamedPipeServerExample()(ok bool){//col:376
/*NamedPipeServerExample()
{
    HANDLE    PipeHandle;
    BOOLEAN   SentMessageResult;
    UINT32    ReadBytes;
    const int BufferSize               = 1024;
    char      BufferToRead[BufferSize] = {0};
    char      BufferToSend[BufferSize] = "test message to send from server !!!";
    PipeHandle = NamedPipeServerCreatePipe("\\\\.\\Pipe\\HyperDbgTests",
                                           BufferSize,
                                           BufferSize);
    if (!PipeHandle)
    {
        return 1;
    }
    if (!NamedPipeServerWaitForClientConntection(PipeHandle))
    {
        return 1;
    }
    ReadBytes =
        NamedPipeServerReadClientMessage(PipeHandle, BufferToRead, BufferSize);
    if (!ReadBytes)
    {
        return 1;
    }
    printf("Message from client : %s\n", BufferToRead);
    SentMessageResult = NamedPipeServerSendMessageToClient(
        PipeHandle,
        BufferToSend,
        strlen(BufferToSend) + 1);
    if (!SentMessageResult)
    {
        return 1;
    }
    NamedPipeServerCloseHandle(PipeHandle);
    return 0;
}*/
return true
}

func (n *namedpipe)NamedPipeClientExample()(ok bool){//col:434
/*NamedPipeClientExample()
{
    HANDLE    PipeHandle;
    BOOLEAN   SentMessageResult;
    UINT32    ReadBytes;
    const int BufferSize         = 1024;
    char      Buffer[BufferSize] = "test message to send from client !!!";
    PipeHandle = NamedPipeClientCreatePipe("\\\\.\\Pipe\\HyperDbgTests");
    if (!PipeHandle)
    {
        return 1;
    }
    SentMessageResult =
        NamedPipeClientSendMessage(PipeHandle, Buffer, strlen(Buffer) + 1);
    if (!SentMessageResult)
    {
        return 1;
    }
    ReadBytes = NamedPipeClientReadMessage(PipeHandle, Buffer, BufferSize);
    if (!ReadBytes)
    {
        return 1;
    }
    printf("Server sent the following message: %s\n", Buffer);
    NamedPipeClientClosePipe(PipeHandle);
    return 0;
}*/
return true
}



