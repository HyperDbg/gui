package communication

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\communication\remote-connection.cpp.back

type (
	RemoteConnection interface {
		RemoteConnectionListen() (ok bool) //col:98
	}
	remoteConnection struct{}
)

func NewRemoteConnection() RemoteConnection { return &remoteConnection{} }

func (r *remoteConnection) RemoteConnectionListen() (ok bool) { //col:98
	/*
	   RemoteConnectionListen(PCSTR Port)

	   	{
	   	    char recvbuf[COMMUNICATION_BUFFER_SIZE] = {0};
	   	    if (IsConnectedToAnyInstanceOfDebuggerOrDebuggee())
	   	    {
	   	        return;
	   	    }
	   	    CommunicationServerCreateServerAndWaitForClient(Port, &g_SeverSocket, &g_ServerListenSocket);
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
	   	        RtlZeroMemory(recvbuf, COMMUNICATION_BUFFER_SIZE);
	   	    ShowMessages("closing the conntection...\n");
	   	    CommunicationServerShutdownAndCleanupConnection(g_SeverSocket,
	   	                                                    g_ServerListenSocket);

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
	   	        if (g_IsEndOfMessageReceived)
	   	        {
	   	            g_IsEndOfMessageReceived = FALSE;
	   	            SetEvent(g_EndOfMessageReceivedEvent);
	   	        RtlZeroMemory(recvbuf, COMMUNICATION_BUFFER_SIZE);
	   	    HyperDbgShowSignature();

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
	   	        if (g_EndOfMessageReceivedEvent == NULL)
	   	        {
	   	            g_EndOfMessageReceivedEvent = CreateEvent(NULL, FALSE, FALSE, NULL);
	   	        g_RemoteDebuggeeListeningThread = CreateThread(
	   	            NULL,
	   	            0,
	   	            RemoteConnectionThreadListeningToDebuggee,
	   	            NULL,
	   	            0,
	   	            &ThreadId);
	   	        ShowMessages("connected to %s:%s\n", Ip, Port);

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

	   RemoteConnectionSendResultsToHost(const char * sendbuf, int len)

	   	{
	   	    if (CommunicationServerSendMessage(g_SeverSocket, sendbuf, len) != 0)
	   	    {
	   	        return 1;
	   	    }
	   	    return 0;
	   	}
	*/
	return true
}

