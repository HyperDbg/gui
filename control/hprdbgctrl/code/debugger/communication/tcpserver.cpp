#undef UNICODE
#define WIN32_LEAN_AND_MEAN
#include "pch.h"
#pragma warning(disable : 4996)
int CommunicationServerCreateServerAndWaitForClient(PCSTR    Port,
                                                SOCKET * ClientSocketArg,
                                                SOCKET * ListenSocketArg){
    WSADATA wsaData;
    int     iResult;
    SOCKET ListenSocket = INVALID_SOCKET;
    SOCKET ClientSocket = INVALID_SOCKET;
    struct addrinfo * result = NULL;
    struct addrinfo   hints;
    iResult = WSAStartup(MAKEWORD(2, 2), &wsaData);
    if (iResult != 0){
        ShowMessages("err, WSAStartup failed (%x)\n", iResult);
        return 1;
    }
    ZeroMemory(&hints, sizeof(hints));
    hints.ai_family   = AF_INET;
    hints.ai_socktype = SOCK_STREAM;
    hints.ai_protocol = IPPROTO_TCP;
    hints.ai_flags    = AI_PASSIVE;
    iResult = getaddrinfo(NULL, Port, &hints, &result);
    if (iResult != 0){
        ShowMessages("err, getaddrinfo failed (%x)\n", iResult);
        WSACleanup();
        return 1;
    }
    ListenSocket =
        socket(result->ai_family, result->ai_socktype, result->ai_protocol);
    if (ListenSocket == INVALID_SOCKET){
        ShowMessages("socket failed with error: %ld\n", WSAGetLastError());
        freeaddrinfo(result);
        WSACleanup();
        return 1;
    }
    iResult = ::bind(ListenSocket, result->ai_addr, (int)result->ai_addrlen);
    if (iResult == SOCKET_ERROR){
        ShowMessages("err, bind failed (%x)\n", WSAGetLastError());
        freeaddrinfo(result);
        closesocket(ListenSocket);
        WSACleanup();
        return 1;
    }
    freeaddrinfo(result);
    iResult = listen(ListenSocket, SOMAXCONN);
    if (iResult == SOCKET_ERROR){
        ShowMessages("err, listen failed (%x)\n", WSAGetLastError());
        closesocket(ListenSocket);
        WSACleanup();
        return 1;
    }
    sockaddr_in name    = {0};
    int         addrlen = sizeof(name);
    ClientSocket = accept(ListenSocket, (struct sockaddr *)&name, &addrlen);
    if (ClientSocket == INVALID_SOCKET){
        ShowMessages("err, accept failed (%x)\n", WSAGetLastError());
        closesocket(ListenSocket);
        WSACleanup();
        return 1;
    }
    ShowMessages("connected to : %s:%d\n", inet_ntoa(name.sin_addr), ntohs(name.sin_port));
    *ClientSocketArg = ClientSocket;
    *ListenSocketArg = ListenSocket;
    return 0;
}
int CommunicationServerReceiveMessage(SOCKET ClientSocket, char * recvbuf, int recvbuflen){
    int iResult;
    iResult = recv(ClientSocket, recvbuf, recvbuflen, 0);
    if (iResult > 0){
    }
    else if (iResult == 0){
    }
    else{
        ShowMessages("err, recv failed (%x)\n", WSAGetLastError());
        closesocket(ClientSocket);
        WSACleanup();
        return 1;
    }
    return 0;
}
int CommunicationServerSendMessage(SOCKET ClientSocket, const char * sendbuf, int length){
    int iSendResult;
    iSendResult = send(ClientSocket, sendbuf, length, 0);
    if (iSendResult == SOCKET_ERROR){
        return 1;
    }
    return 0;
}
int CommunicationServerShutdownAndCleanupConnection(SOCKET ClientSocket,
                                                SOCKET ListenSocket){
    int iResult;
    closesocket(ListenSocket);
    iResult = shutdown(ClientSocket, SD_SEND);
    if (iResult == SOCKET_ERROR){
        closesocket(ClientSocket);
        WSACleanup();
        return 1;
    }
    closesocket(ClientSocket);
    WSACleanup();
    return 0;
}
