#include "pch.h"
extern BYTE g_EndOfBufferCheckTcp[TCP_END_OF_BUFFER_CHARS_COUNT];
extern BOOLEAN g_IsConnectedToHyperDbgLocally;
extern BOOLEAN g_IsConnectedToRemoteDebuggee;
extern BOOLEAN g_IsConnectedToRemoteDebugger;
extern BOOLEAN g_BreakPrintingOutput;
extern BOOLEAN g_IsEndOfMessageReceived;
extern SOCKET g_SeverSocket;
extern SOCKET g_ServerListenSocket;
extern SOCKET g_ClientConnectSocket;
extern HANDLE g_RemoteDebuggeeListeningThread;
extern HANDLE g_EndOfMessageReceivedEvent;
VOID RemoteConnectionListen(PCSTR Port){
    char recvbuf[COMMUNICATION_BUFFER_SIZE] = {0};
    if (IsConnectedToAnyInstanceOfDebuggerOrDebuggee()){
        return;
    }
    CommunicationServerCreateServerAndWaitForClient(Port, &g_SeverSocket, &g_ServerListenSocket);
    if (CommunicationServerReceiveMessage(g_SeverSocket, recvbuf, COMMUNICATION_BUFFER_SIZE) != 0){
        ShowMessages("err, unable to handshake with the remote debugger\n");
        CommunicationServerShutdownAndCleanupConnection(g_SeverSocket, g_ServerListenSocket);
        return;
    }
    if (strcmp((const char *)BuildSignature, recvbuf) != 0){
        ShowMessages(ASSERT_MESSAGE_BUILD_SIGNATURE_DOESNT_MATCH);
        if (CommunicationServerSendMessage(g_SeverSocket, "NO", 3) != 0){
            ShowMessages("err, unable to handshake with the remote debugger\n");
        }
        CommunicationServerShutdownAndCleanupConnection(g_SeverSocket, g_ServerListenSocket);
        return;
    }
    else{
        if (CommunicationServerSendMessage(g_SeverSocket, "OK", 3) != 0){
            ShowMessages("err, unable to handshake with the remote debugger\n");
            CommunicationServerShutdownAndCleanupConnection(g_SeverSocket, g_ServerListenSocket);
            return;
        }
    }
    g_IsConnectedToRemoteDebugger = TRUE;
    g_IsConnectedToHyperDbgLocally = TRUE;
    RtlZeroMemory(recvbuf, COMMUNICATION_BUFFER_SIZE);
    while (true){
        if (CommunicationServerReceiveMessage(g_SeverSocket, recvbuf, COMMUNICATION_BUFFER_SIZE) != 0){
            break;
        }
        int CommandExecutionResult = HyperDbgInterpreter(recvbuf);
        RemoteConnectionSendResultsToHost((const char *)g_EndOfBufferCheckTcp, sizeof(g_EndOfBufferCheckTcp));
        if (CommandExecutionResult == 1){
            exit(0);
        }
        RtlZeroMemory(recvbuf, COMMUNICATION_BUFFER_SIZE);
    }
    g_IsConnectedToHyperDbgLocally = FALSE;
    g_IsConnectedToRemoteDebugger = FALSE;
    ShowMessages("closing the conntection...\n");
    CommunicationServerShutdownAndCleanupConnection(g_SeverSocket,
                                                    g_ServerListenSocket);
}
DWORD WINAPI
RemoteConnectionThreadListeningToDebuggee(LPVOID lpParam){
    char   RecvBuf[COMMUNICATION_BUFFER_SIZE + TCP_END_OF_BUFFER_CHARS_COUNT] = {0};
    UINT32 BuffLenReceived                                                    = 0;
    while (g_IsConnectedToRemoteDebuggee){
        if (CommunicationClientReceiveMessage(g_ClientConnectSocket, RecvBuf, COMMUNICATION_BUFFER_SIZE, &BuffLenReceived) != 0){
            break;
        }
        for (size_t i = 0; i < BuffLenReceived; i++){
            if (RecvBuf[i] == g_EndOfBufferCheckTcp[0] &&
                RecvBuf[i + 1] == g_EndOfBufferCheckTcp[1] &&
                RecvBuf[i + 2] == g_EndOfBufferCheckTcp[2] &&
                RecvBuf[i + 3] == g_EndOfBufferCheckTcp[3]){
                g_IsEndOfMessageReceived = TRUE;
                RecvBuf[i]     = '\x00';
                RecvBuf[i + 1] = '\x00';
                break;
            }
        }
        if (!g_BreakPrintingOutput){
            ShowMessages("%s", RecvBuf);
        }
        if (g_IsEndOfMessageReceived){
            g_IsEndOfMessageReceived = FALSE;
            SetEvent(g_EndOfMessageReceivedEvent);
        }
        RtlZeroMemory(RecvBuf, COMMUNICATION_BUFFER_SIZE);
    }
    g_IsConnectedToHyperDbgLocally = FALSE;
    g_IsConnectedToRemoteDebuggee = FALSE;
    HyperDbgShowSignature();
    return 0;
}
VOID RemoteConnectionConnect(PCSTR Ip, PCSTR Port){
    DWORD  ThreadId;
    CHAR   Recv[3]  = {0};
    UINT32 BuffRecv = 0;
    if (IsConnectedToAnyInstanceOfDebuggerOrDebuggee()){
        return;
    }
    if (CommunicationClientConnectToServer(Ip, Port, &g_ClientConnectSocket) == 1){
        g_IsConnectedToHyperDbgLocally = FALSE;
        g_IsConnectedToRemoteDebuggee = FALSE;
        CommunicationClientShutdownConnection(g_ClientConnectSocket);
        CommunicationClientCleanup(g_ClientConnectSocket);
    }
    else{
        if (CommunicationClientSendMessage(g_ClientConnectSocket, (const char *)BuildSignature, sizeof(BuildSignature)) != 0){
            ShowMessages("err, failed to communicate with debuggee\n");
            return;
        }
        if (CommunicationClientReceiveMessage(g_ClientConnectSocket, Recv, sizeof(Recv), &BuffRecv) != 0){
            ShowMessages("err, failed to receive message from debuggee\n");
            return;
        }
        if (strcmp((const char *)"OK", Recv) != 0){
            ShowMessages(ASSERT_MESSAGE_BUILD_SIGNATURE_DOESNT_MATCH);
            return;
        }
        g_IsConnectedToHyperDbgLocally = FALSE;
        g_IsConnectedToRemoteDebuggee = TRUE;
        if (g_EndOfMessageReceivedEvent == NULL){
            g_EndOfMessageReceivedEvent = CreateEvent(NULL, FALSE, FALSE, NULL);
        }
        g_RemoteDebuggeeListeningThread = CreateThread(
            NULL,
            0,
            RemoteConnectionThreadListeningToDebuggee,
            NULL,
            0,
            &ThreadId);
        ShowMessages("connected to %s:%s\n", Ip, Port);
    }
}
int RemoteConnectionSendCommand(const char * sendbuf, int len){
    if (CommunicationClientSendMessage(g_ClientConnectSocket, sendbuf, len) != 0){
        return 1;
    }
    WaitForSingleObject(
        g_EndOfMessageReceivedEvent,
        INFINITE);
    return 0;
}
int RemoteConnectionSendResultsToHost(const char * sendbuf, int len){
    if (CommunicationServerSendMessage(g_SeverSocket, sendbuf, len) != 0){
        return 1;
    }
    return 0;
}
int RemoteConnectionCloseTheConnectionWithDebuggee(){
    CommunicationClientShutdownConnection(g_ClientConnectSocket);
    CommunicationClientCleanup(g_ClientConnectSocket);
    return 0;
}
