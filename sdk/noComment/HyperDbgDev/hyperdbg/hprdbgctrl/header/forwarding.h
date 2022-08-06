#pragma once
#define MAXIMUM_CHARACTERS_FOR_EVENT_FORWARDING_NAME 50
typedef enum _DEBUGGER_EVENT_FORWARDING_TYPE
{
    EVENT_FORWARDING_NAMEDPIPE,
    EVENT_FORWARDING_FILE,
    EVENT_FORWARDING_TCP
} DEBUGGER_EVENT_FORWARDING_TYPE;
typedef enum _DEBUGGER_EVENT_FORWARDING_STATE
{
    EVENT_FORWARDING_STATE_NOT_OPENED,
    EVENT_FORWARDING_STATE_OPENED,
    EVENT_FORWARDING_CLOSED
} DEBUGGER_EVENT_FORWARDING_STATE;
typedef enum _DEBUGGER_OUTPUT_SOURCE_STATUS
{
    DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_OPENED,
    DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_CLOSED,
    DEBUGGER_OUTPUT_SOURCE_STATUS_ALREADY_OPENED,
    DEBUGGER_OUTPUT_SOURCE_STATUS_ALREADY_CLOSED,
    DEBUGGER_OUTPUT_SOURCE_STATUS_UNKNOWN_ERROR,
} DEBUGGER_OUTPUT_SOURCE_STATUS;
typedef struct _DEBUGGER_EVENT_FORWARDING
{
    DEBUGGER_EVENT_FORWARDING_TYPE  Type;
    DEBUGGER_EVENT_FORWARDING_STATE State;
    HANDLE                          Handle;
    SOCKET                          Socket;
    UINT64                          OutputUniqueTag;
    LIST_ENTRY
    OutputSourcesList; 
    CHAR Name[MAXIMUM_CHARACTERS_FOR_EVENT_FORWARDING_NAME];
} DEBUGGER_EVENT_FORWARDING, *PDEBUGGER_EVENT_FORWARDING;
UINT64
ForwardingGetNewOutputSourceTag();
DEBUGGER_OUTPUT_SOURCE_STATUS
ForwardingOpenOutputSource(PDEBUGGER_EVENT_FORWARDING SourceDescriptor);
DEBUGGER_OUTPUT_SOURCE_STATUS
ForwardingCloseOutputSource(PDEBUGGER_EVENT_FORWARDING SourceDescriptor);
HANDLE
ForwardingCreateOutputSource(DEBUGGER_EVENT_FORWARDING_TYPE SourceType,
                             const string &                 Description,
                             SOCKET *                       Socket);
BOOLEAN
ForwardingPerformEventForwarding(PDEBUGGER_GENERAL_EVENT_DETAIL EventDetail,
                                 CHAR *                         Message,
                                 UINT32                         MessageLength);
BOOLEAN
ForwardingWriteToFile(HANDLE FileHandle, CHAR * Message, UINT32 MessageLength);
BOOLEAN
ForwardingSendToNamedPipe(HANDLE NamedPipeHandle, CHAR * Message, UINT32 MessageLength);
BOOLEAN
ForwardingSendToTcpSocket(SOCKET TcpSocket, CHAR * Message, UINT32 MessageLength);
