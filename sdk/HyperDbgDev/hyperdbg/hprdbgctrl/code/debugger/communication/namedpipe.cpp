#include "pch.h"
extern OVERLAPPED g_OverlappedIoStructureForReadDebugger;
extern OVERLAPPED g_OverlappedIoStructureForWriteDebugger;
HANDLE
NamedPipeServerCreatePipe(LPCSTR PipeName, UINT32 OutputBufferSize, UINT32 InputBufferSize) {
    HANDLE hPipe;
    hPipe = CreateNamedPipeA(PipeName,                   // pipe name
                             PIPE_ACCESS_DUPLEX,         // read/write access
                             PIPE_TYPE_MESSAGE |         // message type pipe
                                 PIPE_READMODE_MESSAGE | // message-read mode
                                 PIPE_WAIT,              // blocking mode
                             PIPE_UNLIMITED_INSTANCES,   // max. instances
                             OutputBufferSize,           // output buffer size
                             InputBufferSize,            // input buffer size
                             NMPWAIT_USE_DEFAULT_WAIT,   // client time-out
                             NULL);                      // default security attribute
    if (INVALID_HANDLE_VALUE == hPipe) {
        ShowMessages("err, rror occurred while creating the pipe (%x)\n",
                     GetLastError());
        return NULL;
    }
    return hPipe;
}

BOOLEAN
NamedPipeServerWaitForClientConntection(HANDLE PipeHandle) {
    BOOL bClientConnected = ConnectNamedPipe(PipeHandle, NULL);
    if (FALSE == bClientConnected) {
        ShowMessages("err, occurred while connecting to the client (%x)\n",
                     GetLastError());
        CloseHandle(PipeHandle);
        return FALSE;
    }
    return TRUE;
}

UINT32
NamedPipeServerReadClientMessage(HANDLE PipeHandle, char * BufferToSave, int MaximumReadBufferLength) {
    DWORD cbBytes;
    BOOL  bResult = ReadFile(PipeHandle,              // handle to pipe
                            BufferToSave,            // buffer to receive data
                            MaximumReadBufferLength, // size of buffer
                            &cbBytes,                // number of bytes read
                            NULL);                   // not overlapped I/O
    if ((!bResult) || (0 == cbBytes)) {
        ShowMessages("err, occurred while reading from the client (%x)\n",
                     GetLastError());
        CloseHandle(PipeHandle);
        return 0;
    }
    return cbBytes;
}

BOOLEAN
NamedPipeServerSendMessageToClient(HANDLE PipeHandle,
                                   char * BufferToSend,
                                   int    BufferSize) {
    DWORD   cbBytes;
    BOOLEAN bResult =
        WriteFile(PipeHandle,   // handle to pipe
                  BufferToSend, // buffer to write from
                  BufferSize,   // number of bytes to write, include the NULL
                  &cbBytes,     // number of bytes written
                  NULL);        // not overlapped I/O
    if ((!bResult) || (BufferSize != cbBytes)) {
        ShowMessages("err, occurred while writing to the client (%x)\n",
                     GetLastError());
        CloseHandle(PipeHandle);
        return FALSE;
    }
    return TRUE;
}

VOID
NamedPipeServerCloseHandle(HANDLE PipeHandle) {
    CloseHandle(PipeHandle);
}

HANDLE
NamedPipeClientCreatePipe(LPCSTR PipeName) {
    HANDLE hPipe;
    hPipe = CreateFileA(PipeName,      // pipe name
                        GENERIC_READ | // read and write access
                            GENERIC_WRITE,
                        0,             // no sharing
                        NULL,          // default security attributes
                        OPEN_EXISTING, // opens existing pipe
                        0,             // default attributes
                        NULL);         // no template file
    if (INVALID_HANDLE_VALUE == hPipe) {
        ShowMessages("err, occurred while connecting to the server (%x)\n",
                     GetLastError());
        return NULL;
    } else {
        return hPipe;
    }
}

HANDLE
NamedPipeClientCreatePipeOverlappedIo(LPCSTR PipeName) {
    HANDLE hPipe;
    hPipe = CreateFileA(PipeName,      // pipe name
                        GENERIC_READ | // read and write access
                            GENERIC_WRITE,
                        0,                    // no sharing
                        NULL,                 // default security attributes
                        OPEN_EXISTING,        // opens existing pipe
                        FILE_FLAG_OVERLAPPED, // Overlapped I/O
                        NULL);                // no template file
    if (INVALID_HANDLE_VALUE == hPipe) {
        ShowMessages("err, occurred while connecting to the server (%x)\n",
                     GetLastError());
        return NULL;
    } else {
        g_OverlappedIoStructureForReadDebugger.hEvent =
            CreateEvent(NULL, TRUE, FALSE, NULL);
        g_OverlappedIoStructureForWriteDebugger.hEvent =
            CreateEvent(NULL, TRUE, FALSE, NULL);
        return hPipe;
    }
}

BOOLEAN
NamedPipeClientSendMessage(HANDLE PipeHandle, char * BufferToSend, int BufferSize) {
    DWORD cbBytes;
    BOOL  bResult =
        WriteFile(PipeHandle,   // handle to pipe
                  BufferToSend, // buffer to write from
                  BufferSize,   // number of bytes to write, include the NULL
                  &cbBytes,     // number of bytes written
                  NULL);        // not overlapped I/O
    if ((!bResult) || (BufferSize != cbBytes)) {
        ShowMessages("err, occurred while writing to the server (%x)\n",
                     GetLastError());
        CloseHandle(PipeHandle);
        CloseHandle(PipeHandle);
        return FALSE;
    } else {
        return TRUE;
    }
}

UINT32
NamedPipeClientReadMessage(HANDLE PipeHandle, char * BufferToRead, int MaximumSizeOfBuffer) {
    DWORD cbBytes;
    BOOL  bResult = ReadFile(PipeHandle,          // handle to pipe
                            BufferToRead,        // buffer to receive data
                            MaximumSizeOfBuffer, // size of buffer
                            &cbBytes,            // number of bytes read
                            NULL);               // not overlapped I/O
    if ((!bResult) || (0 == cbBytes)) {
        ShowMessages("err, occurred while reading from the server (%x)\n",
                     GetLastError());
        CloseHandle(PipeHandle);
        return NULL; // Error
    }
    return cbBytes;
}

VOID
NamedPipeClientClosePipe(HANDLE PipeHandle) {
    CloseHandle(PipeHandle);
}

int
NamedPipeServerExample() {
    HANDLE    PipeHandle;
    BOOLEAN   SentMessageResult;
    UINT32    ReadBytes;
    const int BufferSize               = 1024;
    char      BufferToRead[BufferSize] = {0};
    char      BufferToSend[BufferSize] = "test message to send from server !!!";
    PipeHandle                         = NamedPipeServerCreatePipe("\\\\.\\Pipe\\HyperDbgTests",
                                           BufferSize,
                                           BufferSize);
    if (!PipeHandle) {
        return 1;
    }
    if (!NamedPipeServerWaitForClientConntection(PipeHandle)) {
        return 1;
    }
    ReadBytes =
        NamedPipeServerReadClientMessage(PipeHandle, BufferToRead, BufferSize);
    if (!ReadBytes) {
        return 1;
    }
    ShowMessages("message from client : %s\n", BufferToRead);
    SentMessageResult = NamedPipeServerSendMessageToClient(
        PipeHandle,
        BufferToSend,
        strlen(BufferToSend) + 1);
    if (!SentMessageResult) {
        return 1;
    }
    NamedPipeServerCloseHandle(PipeHandle);
    return 0;
}

int
NamedPipeClientExample() {
    HANDLE    PipeHandle;
    BOOLEAN   SentMessageResult;
    UINT32    ReadBytes;
    const int BufferSize         = 1024;
    char      Buffer[BufferSize] = "test message to send from client !!!";
    PipeHandle                   = NamedPipeClientCreatePipe("\\\\.\\Pipe\\HyperDbgTests");
    if (!PipeHandle) {
        return 1;
    }
    SentMessageResult =
        NamedPipeClientSendMessage(PipeHandle, Buffer, strlen(Buffer) + 1);
    if (!SentMessageResult) {
        return 1;
    }
    ReadBytes = NamedPipeClientReadMessage(PipeHandle, Buffer, BufferSize);
    if (!ReadBytes) {
        return 1;
    }
    ShowMessages("server sent the following message: %s\n", Buffer);
    NamedPipeClientClosePipe(PipeHandle);
    return 0;
}
