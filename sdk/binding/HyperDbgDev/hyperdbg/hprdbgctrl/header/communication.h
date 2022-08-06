










#pragma once





#define COM1_PORT 0x03F8
#define COM2_PORT 0x02F8
#define COM3_PORT 0x03E8
#define COM4_PORT 0x02E8





int
CommunicationServerCreateServerAndWaitForClient(PCSTR    Port,
                                                SOCKET * ClientSocketArg,
                                                SOCKET * ListenSocketArg);

int
CommunicationServerReceiveMessage(SOCKET ClientSocket, char * recvbuf, int recvbuflen);

int
CommunicationServerSendMessage(SOCKET ClientSocket, const char * sendbuf, int length);

int
CommunicationServerShutdownAndCleanupConnection(SOCKET ClientSocket,
                                                SOCKET ListenSocket);





int
CommunicationClientConnectToServer(PCSTR Ip, PCSTR Port, SOCKET * ConnectSocketArg);

int
CommunicationClientSendMessage(SOCKET ConnectSocket, const char * sendbuf, int buflen);

int
CommunicationClientShutdownConnection(SOCKET ConnectSocket);

int
CommunicationClientReceiveMessage(SOCKET ConnectSocket, CHAR * RecvBuf, UINT32 MaxBuffLen, PUINT32 BuffLenRecvd);

int
CommunicationClientCleanup(SOCKET ConnectSocket);





VOID
RemoteConnectionListen(PCSTR Port);

VOID
RemoteConnectionConnect(PCSTR Ip, PCSTR Port);

int
RemoteConnectionSendCommand(const char * sendbuf, int len);

int
RemoteConnectionSendResultsToHost(const char * sendbuf, int len);

int
RemoteConnectionCloseTheConnectionWithDebuggee();
