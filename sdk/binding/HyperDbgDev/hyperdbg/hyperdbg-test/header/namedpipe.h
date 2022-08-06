










#pragma once





HANDLE
NamedPipeServerCreatePipe(LPCSTR PipeName, UINT32 OutputBufferSize, UINT32 InputBufferSize);

BOOLEAN
NamedPipeServerWaitForClientConntection(HANDLE PipeHandle);

UINT32
NamedPipeServerReadClientMessage(HANDLE PipeHandle, char * BufferToSave, int MaximumReadBufferLength);

BOOLEAN
NamedPipeServerSendMessageToClient(HANDLE PipeHandle,
                                   char * BufferToSend,
                                   int    BufferSize);

VOID
NamedPipeServerCloseHandle(HANDLE PipeHandle);





HANDLE
NamedPipeClientCreatePipe(LPCSTR PipeName);

BOOLEAN
NamedPipeClientSendMessage(HANDLE PipeHandle, char * BufferToSend, int BufferSize);

UINT32
NamedPipeClientReadMessage(HANDLE PipeHandle, char * BufferToRead, int MaximumSizeOfBuffer);

VOID
NamedPipeClientClosePipe(HANDLE PipeHandle);
