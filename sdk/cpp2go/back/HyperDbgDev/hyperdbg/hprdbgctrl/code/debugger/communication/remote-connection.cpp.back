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
VOID
RemoteConnectionListen(PCSTR Port)
{
    char recvbuf[COMMUNICATION_BUFFER_SIZE] = {0};
    if (IsConnectedToAnyInstanceOfDebuggerOrDebuggee())
    {
        return;
    }
    CommunicationServerCreateServerAndWaitForClient(Port, &g_SeverSocket, &g_ServerListenSocket);
    g_IsConnectedToRemoteDebugger = TRUE;
    g_IsConnectedToHyperDbgLocally = TRUE;
    while (true)
    {
        if (CommunicationServerReceiveMessage(g_SeverSocket, recvbuf, COMMUNICATION_BUFFER_SIZE) != 0)
        {
            break;
        }
        int CommandExecutionResult = HyperDbgInterpreter(recvbuf);
        RemoteConnectionSendResultsToHost((const char *)g_EndOfBufferCheckTcp, sizeof(g_EndOfBufferCheckTcp));
        if (CommandExecutionResult == 1)
        {
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
RemoteConnectionThreadListeningToDebuggee(LPVOID lpParam)
{
    char   recvbuf[COMMUNICATION_BUFFER_SIZE + TCP_END_OF_BUFFER_CHARS_COUNT] = {0};
    UINT32 BuffLenReceived                                                    = 0;
    while (g_IsConnectedToRemoteDebuggee)
    {
        if (CommunicationClientReceiveMessage(g_ClientConnectSocket, recvbuf, COMMUNICATION_BUFFER_SIZE, &BuffLenReceived) != 0)
        {
            break;
        }
        for (size_t i = 0; i < BuffLenReceived; i++)
        {
            if (recvbuf[i] == g_EndOfBufferCheckTcp[0] &&
                recvbuf[i + 1] == g_EndOfBufferCheckTcp[1] &&
                recvbuf[i + 2] == g_EndOfBufferCheckTcp[2] &&
                recvbuf[i + 3] == g_EndOfBufferCheckTcp[3])
            {
                g_IsEndOfMessageReceived = TRUE;
                recvbuf[i]     = '\x00';
                recvbuf[i + 1] = '\x00';
                break;
            }
        }
        if (!g_BreakPrintingOutput)
        {
            ShowMessages("%s", recvbuf);
        }
        if (g_IsEndOfMessageReceived)
        {
            g_IsEndOfMessageReceived = FALSE;
            SetEvent(g_EndOfMessageReceivedEvent);
        }
        RtlZeroMemory(recvbuf, COMMUNICATION_BUFFER_SIZE);
    }
    g_IsConnectedToHyperDbgLocally = FALSE;
    g_IsConnectedToRemoteDebuggee = FALSE;
    HyperDbgShowSignature();
    return 0;
}

VOID
RemoteConnectionConnect(PCSTR Ip, PCSTR Port)
{
    DWORD ThreadId;
    if (IsConnectedToAnyInstanceOfDebuggerOrDebuggee())
    {
        return;
    }
    if (CommunicationClientConnectToServer(Ip, Port, &g_ClientConnectSocket) ==
        1)
    {
        g_IsConnectedToHyperDbgLocally = FALSE;
        g_IsConnectedToRemoteDebuggee = FALSE;
        CommunicationClientShutdownConnection(g_ClientConnectSocket);
        CommunicationClientCleanup(g_ClientConnectSocket);
    }
    else
    {
        g_IsConnectedToHyperDbgLocally = FALSE;
        g_IsConnectedToRemoteDebuggee = TRUE;
        if (g_EndOfMessageReceivedEvent == NULL)
        {
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

int
RemoteConnectionSendCommand(const char * sendbuf, int len)
{
    if (CommunicationClientSendMessage(g_ClientConnectSocket, sendbuf, len) !=
        0)
    {
        return 1;
    }
    WaitForSingleObject(
        g_EndOfMessageReceivedEvent,
        INFINITE);
    return 0;
}

int
RemoteConnectionSendResultsToHost(const char * sendbuf, int len)
{
    if (CommunicationServerSendMessage(g_SeverSocket, sendbuf, len) != 0)
    {
        return 1;
    }
    return 0;
}

int
RemoteConnectionCloseTheConnectionWithDebuggee()
{
    CommunicationClientShutdownConnection(g_ClientConnectSocket);
    CommunicationClientCleanup(g_ClientConnectSocket);
    return 0;
}

