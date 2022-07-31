package communication
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\communication\forwarding.cpp.back

type (
Forwarding interface{
ForwardingGetNewOutputSourceTag()(ok bool)//col:30
ForwardingOpenOutputSource()(ok bool)//col:92
ForwardingCloseOutputSource()(ok bool)//col:174
 * valid value for handle ()(ok bool)//col:278
ForwardingPerformEventForwarding()(ok bool)//col:370
ForwardingWriteToFile()(ok bool)//col:429
ForwardingSendToNamedPipe()(ok bool)//col:462
ForwardingSendToTcpSocket()(ok bool)//col:490
}
)

func NewForwarding() { return & forwarding{} }

func (f *forwarding)ForwardingGetNewOutputSourceTag()(ok bool){//col:30
/*ForwardingGetNewOutputSourceTag()
{
    return g_OutputSourceTag++;
}*/
return true
}

func (f *forwarding)ForwardingOpenOutputSource()(ok bool){//col:92
/*ForwardingOpenOutputSource(PDEBUGGER_EVENT_FORWARDING SourceDescriptor)
{
    if (SourceDescriptor->State == EVENT_FORWARDING_CLOSED)
    {
        return DEBUGGER_OUTPUT_SOURCE_STATUS_ALREADY_CLOSED;
    }
    if (SourceDescriptor->State == EVENT_FORWARDING_STATE_OPENED)
    {
        return DEBUGGER_OUTPUT_SOURCE_STATUS_ALREADY_OPENED;
    }
    SourceDescriptor->State = EVENT_FORWARDING_STATE_OPENED;
    if (SourceDescriptor->Type == EVENT_FORWARDING_FILE)
    {
        return DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_OPENED;
    }
    else if (SourceDescriptor->Type == EVENT_FORWARDING_NAMEDPIPE)
    {
        return DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_OPENED;
    }
    else if (SourceDescriptor->Type == EVENT_FORWARDING_TCP)
    {
        return DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_OPENED;
    }
    return DEBUGGER_OUTPUT_SOURCE_STATUS_UNKNOWN_ERROR;
}*/
return true
}

func (f *forwarding)ForwardingCloseOutputSource()(ok bool){//col:174
/*ForwardingCloseOutputSource(PDEBUGGER_EVENT_FORWARDING SourceDescriptor)
{
    if (SourceDescriptor->State == EVENT_FORWARDING_CLOSED)
    {
        return DEBUGGER_OUTPUT_SOURCE_STATUS_ALREADY_CLOSED;
    }
    if (SourceDescriptor->State == EVENT_FORWARDING_STATE_NOT_OPENED ||
        SourceDescriptor->State != EVENT_FORWARDING_STATE_OPENED)
    {
        return DEBUGGER_OUTPUT_SOURCE_STATUS_UNKNOWN_ERROR;
    }
    SourceDescriptor->State = EVENT_FORWARDING_CLOSED;
    if (SourceDescriptor->Type == EVENT_FORWARDING_FILE)
    {
        CloseHandle(SourceDescriptor->Handle);
        return DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_CLOSED;
    }
    else if (SourceDescriptor->Type == EVENT_FORWARDING_TCP)
    {
        CommunicationClientShutdownConnection(SourceDescriptor->Socket);
        CommunicationClientCleanup(SourceDescriptor->Socket);
        return DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_CLOSED;
    }
    else if (SourceDescriptor->Type == EVENT_FORWARDING_NAMEDPIPE)
    {
        NamedPipeClientClosePipe(SourceDescriptor->Handle);
        return DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_CLOSED;
    }
    return DEBUGGER_OUTPUT_SOURCE_STATUS_UNKNOWN_ERROR;
}*/
return true
}

func (f *forwarding) * valid value for handle ()(ok bool){//col:278
/* * valid value for handle (TRUE) which is not a valid handle but it
 * indicates that the operation was successful and the caller can use
 * the pointer that it passed as the socket.
 * On anything other than tcp sockets, the socket pointer in not
 * modified; thus, it's not value
 *
HANDLE
ForwardingCreateOutputSource(DEBUGGER_EVENT_FORWARDING_TYPE SourceType,
                             const string &                 Description,
                             SOCKET *                       Socket)
{
    string IpPortDelimiter;
    string Ip;
    string Port;
    if (SourceType == EVENT_FORWARDING_FILE)
    {
        HANDLE FileHandle = CreateFileA(Description.c_str(), GENERIC_WRITE, 0, NULL, OPEN_ALWAYS, FILE_ATTRIBUTE_NORMAL, NULL);
        return FileHandle;
    }
    else if (SourceType == EVENT_FORWARDING_NAMEDPIPE)
    {
        HANDLE PipeHandle = NamedPipeClientCreatePipe(Description.c_str());
        if (!PipeHandle)
        {
            return INVALID_HANDLE_VALUE;
        }
        return PipeHandle;
    }
    else if (SourceType == EVENT_FORWARDING_TCP)
    {
        if (Description.find(':') != std::string::npos)
        {
            IpPortDelimiter = ':';
            size_t find     = Description.find(IpPortDelimiter);
            Ip              = Description.substr(0, find);
            Port            = Description.substr(find + 1, find + Description.size());
            if (CommunicationClientConnectToServer(Ip.c_str(), Port.c_str(), Socket) == 0)
            {
                return (HANDLE)TRUE;
            }
            else
            {
                return INVALID_HANDLE_VALUE;
            }
        }
        else
        {
            return INVALID_HANDLE_VALUE;
        }
    }
    return INVALID_HANDLE_VALUE;
}*/
return true
}

func (f *forwarding)ForwardingPerformEventForwarding()(ok bool){//col:370
/*ForwardingPerformEventForwarding(PDEBUGGER_GENERAL_EVENT_DETAIL EventDetail,
                                 CHAR *                         Message,
                                 UINT32                         MessageLength)
{
    BOOLEAN     Result   = FALSE;
    PLIST_ENTRY TempList = 0;
    for (size_t i = 0; i < DebuggerOutputSourceMaximumRemoteSourceForSingleEvent;
         i++)
    {
        if (EventDetail->OutputSourceTags[i] == NULL)
        {
            return Result;
        }
        TempList = &g_OutputSources;
        while (&g_OutputSources != TempList->Flink)
        {
            TempList = TempList->Flink;
            PDEBUGGER_EVENT_FORWARDING CurrentOutputSourceDetails = CONTAINING_RECORD(
                TempList,
                DEBUGGER_EVENT_FORWARDING,
                OutputSourcesList);
            if (EventDetail->OutputSourceTags[i] ==
                CurrentOutputSourceDetails->OutputUniqueTag)
            {
                if (CurrentOutputSourceDetails->State ==
                    EVENT_FORWARDING_STATE_OPENED)
                {
                    switch (CurrentOutputSourceDetails->Type)
                    {
                    case EVENT_FORWARDING_NAMEDPIPE:
                        Result = ForwardingSendToNamedPipe(
                            CurrentOutputSourceDetails->Handle,
                            Message,
                            MessageLength);
                        break;
                    case EVENT_FORWARDING_FILE:
                        Result = ForwardingWriteToFile(CurrentOutputSourceDetails->Handle,
                                                       Message,
                                                       MessageLength);
                        break;
                    case EVENT_FORWARDING_TCP:
                        Result = ForwardingSendToTcpSocket(
                            CurrentOutputSourceDetails->Socket,
                            Message,
                            MessageLength);
                        break;
                    default:
                        break;
                    }
                }
                break;
            }
        }
    }
    return FALSE;
}*/
return true
}

func (f *forwarding)ForwardingWriteToFile()(ok bool){//col:429
/*ForwardingWriteToFile(HANDLE FileHandle, CHAR * Message, UINT32 MessageLength)
{
    DWORD BytesWritten = 0;
    BOOL  ErrorFlag    = FALSE;
    if (ErrorFlag == FALSE)
    {
        return FALSE;
    }
    else
    {
        if (BytesWritten != MessageLength)
        {
            return FALSE;
        }
        else
        {
            return TRUE;
        }
    }
    return FALSE;
}*/
return true
}

func (f *forwarding)ForwardingSendToNamedPipe()(ok bool){//col:462
/*ForwardingSendToNamedPipe(HANDLE NamedPipeHandle, CHAR * Message, UINT32 MessageLength)
{
    BOOLEAN SentMessageResult;
    SentMessageResult =
        NamedPipeClientSendMessage(NamedPipeHandle, Message, MessageLength);
    if (!SentMessageResult)
    {
        return FALSE;
    }
    return TRUE;
}*/
return true
}

func (f *forwarding)ForwardingSendToTcpSocket()(ok bool){//col:490
/*ForwardingSendToTcpSocket(SOCKET TcpSocket, CHAR * Message, UINT32 MessageLength)
{
    if (CommunicationClientSendMessage(TcpSocket, Message, MessageLength) != 0)
    {
        return FALSE;
    }
    return TRUE;
}*/
return true
}


