package communication

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\communication\tcpserver.cpp.back

type (
	Tcpserver interface {
		CommunicationServerCreateServerAndWaitForClient() (ok bool) //col:73
	}
	tcpserver struct{}
)

func NewTcpserver() Tcpserver { return &tcpserver{} }

func (t *tcpserver) CommunicationServerCreateServerAndWaitForClient() (ok bool) { //col:73
	/*
	   CommunicationServerCreateServerAndWaitForClient(PCSTR    Port,

	   	SOCKET * ClientSocketArg,
	   	SOCKET * ListenSocketArg)

	   	{
	   	    WSADATA wsaData;
	   	    int     iResult;
	   	    SOCKET ListenSocket = INVALID_SOCKET;
	   	    SOCKET ClientSocket = INVALID_SOCKET;
	   	    struct addrinfo * result = NULL;
	   	    struct addrinfo   hints;
	   	    iResult = WSAStartup(MAKEWORD(2, 2), &wsaData);
	   	    if (iResult != 0)
	   	    {
	   	        ShowMessages("err, WSAStartup failed (%x)\n", iResult);
	   	    ZeroMemory(&hints, sizeof(hints));
	   	    iResult = getaddrinfo(NULL, Port, &hints, &result);
	   	    if (iResult != 0)
	   	    {
	   	        ShowMessages("err, getaddrinfo failed (%x)\n", iResult);
	   	        WSACleanup();
	   	        socket(result->ai_family, result->ai_socktype, result->ai_protocol);
	   	    if (ListenSocket == INVALID_SOCKET)
	   	    {
	   	        ShowMessages("socket failed with error: %ld\n", WSAGetLastError());
	   	        freeaddrinfo(result);
	   	        WSACleanup();
	   	    iResult = ::bind(ListenSocket, result->ai_addr, (int)result->ai_addrlen);
	   	    if (iResult == SOCKET_ERROR)
	   	    {
	   	        ShowMessages("err, bind failed (%x)\n", WSAGetLastError());
	   	        freeaddrinfo(result);
	   	        closesocket(ListenSocket);
	   	        WSACleanup();
	   	    freeaddrinfo(result);
	   	    iResult = listen(ListenSocket, SOMAXCONN);
	   	    if (iResult == SOCKET_ERROR)
	   	    {
	   	        ShowMessages("err, listen failed (%x)\n", WSAGetLastError());
	   	        closesocket(ListenSocket);
	   	        WSACleanup();
	   	    int         addrlen = sizeof(name);
	   	    ClientSocket = accept(ListenSocket, (struct sockaddr *)&name, &addrlen);
	   	    if (ClientSocket == INVALID_SOCKET)
	   	    {
	   	        ShowMessages("err, accept failed (%x)\n", WSAGetLastError());
	   	        closesocket(ListenSocket);
	   	        WSACleanup();
	   	    ShowMessages("connected to : %s:%d\n", inet_ntoa(name.sin_addr), ntohs(name.sin_port));

	   CommunicationServerReceiveMessage(SOCKET ClientSocket, char * recvbuf, int recvbuflen)

	   	{
	   	    int iResult;
	   	    iResult = recv(ClientSocket, recvbuf, recvbuflen, 0);
	   	    if (iResult > 0)
	   	    {
	   	    }
	   	    else if (iResult == 0)
	   	    {
	   	    }
	   	    else
	   	    {
	   	        ShowMessages("err, recv failed (%x)\n", WSAGetLastError());
	   	        closesocket(ClientSocket);
	   	        WSACleanup();

	   CommunicationServerSendMessage(SOCKET ClientSocket, const char * sendbuf, int length)

	   	{
	   	    int iSendResult;
	   	    iSendResult = send(ClientSocket, sendbuf, length, 0);
	   	    if (iSendResult == SOCKET_ERROR)
	   	    {
	   	        return 1;
	   	    }
	   	    return 0;
	   	}
	*/
	return true
}

