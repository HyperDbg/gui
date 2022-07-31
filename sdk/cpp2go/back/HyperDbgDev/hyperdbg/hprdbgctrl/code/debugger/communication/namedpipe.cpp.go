package communication
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\communication\namedpipe.cpp.back

type (
Namedpipe interface{
NamedPipeServerCreatePipe()(ok bool)//col:57
NamedPipeServerWaitForClientConntection()(ok bool)//col:85
NamedPipeServerReadClientMessage()(ok bool)//col:128
NamedPipeServerSendMessageToClient()(ok bool)//col:155
NamedPipeServerCloseHandle()(ok bool)//col:167
NamedPipeClientCreatePipe()(ok bool)//col:222
NamedPipeClientCreatePipeOverlappedIo()(ok bool)//col:281
NamedPipeClientSendMessage()(ok bool)//col:329
NamedPipeClientReadMessage()(ok bool)//col:360
NamedPipeClientClosePipe()(ok bool)//col:372
NamedPipeServerExample()(ok bool)//col:443
NamedPipeClientExample()(ok bool)//col:501
}
)

func NewNamedpipe() { return & namedpipe{} }

func (n *namedpipe)NamedPipeServerCreatePipe()(ok bool){//col:57
/*NamedPipeServerCreatePipe(LPCSTR PipeName, UINT32 OutputBufferSize, UINT32 InputBufferSize)
{
    HANDLE hPipe;
    if (INVALID_HANDLE_VALUE == hPipe)
    {
        ShowMessages("err, rror occurred while creating the pipe (%x)\n",
                     GetLastError());
        return NULL;
    }
    return hPipe;
}*/
return true
}

func (n *namedpipe)NamedPipeServerWaitForClientConntection()(ok bool){//col:85
/*NamedPipeServerWaitForClientConntection(HANDLE PipeHandle)
{
    BOOL bClientConnected = ConnectNamedPipe(PipeHandle, NULL);
    if (FALSE == bClientConnected)
    {
        ShowMessages("err, occurred while connecting to the client (%x)\n",
                     GetLastError());
        CloseHandle(PipeHandle);
        return FALSE;
    }
    return TRUE;
}*/
return true
}

func (n *namedpipe)NamedPipeServerReadClientMessage()(ok bool){//col:128
/*NamedPipeServerReadClientMessage(HANDLE PipeHandle, char * BufferToSave, int MaximumReadBufferLength)
{
    DWORD cbBytes;
    if ((!bResult) || (0 == cbBytes))
    {
        ShowMessages("err, occurred while reading from the client (%x)\n",
                     GetLastError());
        CloseHandle(PipeHandle);
        return 0;
    }
    return cbBytes;
}*/
return true
}

func (n *namedpipe)NamedPipeServerSendMessageToClient()(ok bool){//col:155
/*NamedPipeServerSendMessageToClient(HANDLE PipeHandle,
                                   char * BufferToSend,
                                   int    BufferSize)
{
    DWORD cbBytes;
    BOOLEAN bResult =
    if ((!bResult) || (BufferSize != cbBytes))
    {
        ShowMessages("err, occurred while writing to the client (%x)\n",
                     GetLastError());
        CloseHandle(PipeHandle);
        return FALSE;
    }
    return TRUE;
}*/
return true
}

func (n *namedpipe)NamedPipeServerCloseHandle()(ok bool){//col:167
/*NamedPipeServerCloseHandle(HANDLE PipeHandle)
{
    CloseHandle(PipeHandle);
}*/
return true
}

func (n *namedpipe)NamedPipeClientCreatePipe()(ok bool){//col:222
/*NamedPipeClientCreatePipe(LPCSTR PipeName)
{
    HANDLE hPipe;
                            GENERIC_WRITE,
    if (INVALID_HANDLE_VALUE == hPipe)
    {
        ShowMessages("err, occurred while connecting to the server (%x)\n",
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

func (n *namedpipe)NamedPipeClientCreatePipeOverlappedIo()(ok bool){//col:281
/*NamedPipeClientCreatePipeOverlappedIo(LPCSTR PipeName)
{
    HANDLE hPipe;
                            GENERIC_WRITE,
    if (INVALID_HANDLE_VALUE == hPipe)
    {
        ShowMessages("err, occurred while connecting to the server (%x)\n",
                     GetLastError());
        return NULL;
    }
    else
    {
        g_OverlappedIoStructureForReadDebugger.hEvent =
            CreateEvent(NULL, TRUE, FALSE, NULL);
        g_OverlappedIoStructureForWriteDebugger.hEvent =
            CreateEvent(NULL, TRUE, FALSE, NULL);
        return hPipe;
    }
}*/
return true
}

func (n *namedpipe)NamedPipeClientSendMessage()(ok bool){//col:329
/*NamedPipeClientSendMessage(HANDLE PipeHandle, char * BufferToSend, int BufferSize)
{
    DWORD cbBytes;
    BOOL bResult =
    if ((!bResult) || (BufferSize != cbBytes))
    {
        ShowMessages("err, occurred while writing to the server (%x)\n",
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

func (n *namedpipe)NamedPipeClientReadMessage()(ok bool){//col:360
/*NamedPipeClientReadMessage(HANDLE PipeHandle, char * BufferToRead, int MaximumSizeOfBuffer)
{
    DWORD cbBytes;
    if ((!bResult) || (0 == cbBytes))
    {
        ShowMessages("err, occurred while reading from the server (%x)\n",
                     GetLastError());
        CloseHandle(PipeHandle);
    }
    return cbBytes;
}*/
return true
}

func (n *namedpipe)NamedPipeClientClosePipe()(ok bool){//col:372
/*NamedPipeClientClosePipe(HANDLE PipeHandle)
{
    CloseHandle(PipeHandle);
}*/
return true
}

func (n *namedpipe)NamedPipeServerExample()(ok bool){//col:443
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
    ShowMessages("message from client : %s\n", BufferToRead);
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

func (n *namedpipe)NamedPipeClientExample()(ok bool){//col:501
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
    ShowMessages("server sent the following message: %s\n", Buffer);
    NamedPipeClientClosePipe(PipeHandle);
    return 0;
}*/
return true
}



