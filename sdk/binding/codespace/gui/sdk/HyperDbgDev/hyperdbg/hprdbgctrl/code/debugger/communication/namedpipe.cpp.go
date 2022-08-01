package communication
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\communication\namedpipe.cpp.back

type (
Namedpipe interface{
NamedPipeServerCreatePipe()(ok bool)//col:21
NamedPipeServerWaitForClientConntection()(ok bool)//col:33
NamedPipeServerReadClientMessage()(ok bool)//col:50
NamedPipeServerSendMessageToClient()(ok bool)//col:70
NamedPipeServerCloseHandle()(ok bool)//col:74
NamedPipeClientCreatePipe()(ok bool)//col:96
NamedPipeClientCreatePipeOverlappedIo()(ok bool)//col:122
NamedPipeClientSendMessage()(ok bool)//col:144
NamedPipeClientReadMessage()(ok bool)//col:161
NamedPipeClientClosePipe()(ok bool)//col:165
NamedPipeServerExample()(ok bool)//col:202
NamedPipeClientExample()(ok bool)//col:229
}
namedpipe struct{}
)

func NewNamedpipe()Namedpipe{ return & namedpipe{} }

func (n *namedpipe)NamedPipeServerCreatePipe()(ok bool){//col:21
/*NamedPipeServerCreatePipe(LPCSTR PipeName, UINT32 OutputBufferSize, UINT32 InputBufferSize)
{
    HANDLE hPipe;
    hPipe = CreateNamedPipeA(PipeName,                   
                             PIPE_ACCESS_DUPLEX,         
                             PIPE_TYPE_MESSAGE |         
                                 PIPE_READMODE_MESSAGE | 
                                 PIPE_WAIT,              
                             PIPE_UNLIMITED_INSTANCES,   
                             OutputBufferSize,           
                             InputBufferSize,            
                             NMPWAIT_USE_DEFAULT_WAIT,   
                             NULL);                      
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

func (n *namedpipe)NamedPipeServerWaitForClientConntection()(ok bool){//col:33
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

func (n *namedpipe)NamedPipeServerReadClientMessage()(ok bool){//col:50
/*NamedPipeServerReadClientMessage(HANDLE PipeHandle, char * BufferToSave, int MaximumReadBufferLength)
{
    DWORD cbBytes;
    BOOL bResult = ReadFile(PipeHandle,              
                            BufferToSave,            
                            MaximumReadBufferLength, 
                            &cbBytes,                
                            NULL);                   
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

func (n *namedpipe)NamedPipeServerSendMessageToClient()(ok bool){//col:70
/*NamedPipeServerSendMessageToClient(HANDLE PipeHandle,
                                   char * BufferToSend,
                                   int    BufferSize)
{
    DWORD cbBytes;
    BOOLEAN bResult =
        WriteFile(PipeHandle,   
                  BufferToSend, 
                  BufferSize,   
                  &cbBytes,     
                  NULL);        
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

func (n *namedpipe)NamedPipeServerCloseHandle()(ok bool){//col:74
/*NamedPipeServerCloseHandle(HANDLE PipeHandle)
{
    CloseHandle(PipeHandle);
}*/
return true
}

func (n *namedpipe)NamedPipeClientCreatePipe()(ok bool){//col:96
/*NamedPipeClientCreatePipe(LPCSTR PipeName)
{
    HANDLE hPipe;
    hPipe = CreateFileA(PipeName,      
                        GENERIC_READ | 
                            GENERIC_WRITE,
                        0,             
                        NULL,          
                        OPEN_EXISTING, 
                        0,             
                        NULL);         
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

func (n *namedpipe)NamedPipeClientCreatePipeOverlappedIo()(ok bool){//col:122
/*NamedPipeClientCreatePipeOverlappedIo(LPCSTR PipeName)
{
    HANDLE hPipe;
    hPipe = CreateFileA(PipeName,      
                        GENERIC_READ | 
                            GENERIC_WRITE,
                        0,                    
                        NULL,                 
                        OPEN_EXISTING,        
                        FILE_FLAG_OVERLAPPED, 
                        NULL);                
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

func (n *namedpipe)NamedPipeClientSendMessage()(ok bool){//col:144
/*NamedPipeClientSendMessage(HANDLE PipeHandle, char * BufferToSend, int BufferSize)
{
    DWORD cbBytes;
    BOOL bResult =
        WriteFile(PipeHandle,   
                  BufferToSend, 
                  BufferSize,   
                  &cbBytes,     
                  NULL);        
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

func (n *namedpipe)NamedPipeClientReadMessage()(ok bool){//col:161
/*NamedPipeClientReadMessage(HANDLE PipeHandle, char * BufferToRead, int MaximumSizeOfBuffer)
{
    DWORD cbBytes;
    BOOL bResult = ReadFile(PipeHandle,          
                            BufferToRead,        
                            MaximumSizeOfBuffer, 
                            &cbBytes,            
                            NULL);               
    if ((!bResult) || (0 == cbBytes))
    {
        ShowMessages("err, occurred while reading from the server (%x)\n",
                     GetLastError());
        CloseHandle(PipeHandle);
        return NULL; 
    }
    return cbBytes;
}*/
return true
}

func (n *namedpipe)NamedPipeClientClosePipe()(ok bool){//col:165
/*NamedPipeClientClosePipe(HANDLE PipeHandle)
{
    CloseHandle(PipeHandle);
}*/
return true
}

func (n *namedpipe)NamedPipeServerExample()(ok bool){//col:202
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

func (n *namedpipe)NamedPipeClientExample()(ok bool){//col:229
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



