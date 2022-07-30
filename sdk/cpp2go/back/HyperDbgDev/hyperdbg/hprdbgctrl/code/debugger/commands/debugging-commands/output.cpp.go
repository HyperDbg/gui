package debugging-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\output.cpp.back

type (
Output interface{
CommandOutputHelp()(ok bool)//col:42
CommandOutput()(ok bool)//col:462
}
)

func NewOutput() { return & output{} }

func (o *output)CommandOutputHelp()(ok bool){//col:42
/*CommandOutputHelp()
{
    ShowMessages("output : creates an output instance that can be used in event "
                 "forwarding.\n\n");
    ShowMessages("syntax : \toutput [create Name (string)] [file|namedpipe|tcp Address (string)]\n");
    ShowMessages("syntax : \toutput [open|close Name (string)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : output create MyOutputName1 file "
                 "c:\\users\\sina\\desktop\\output.txt\n");
    ShowMessages("\t\te.g : output create MyOutputName2 tcp 192.168.1.10:8080\n");
    ShowMessages("\t\te.g : output create MyOutputName3 namedpipe "
                 "\\\\.\\Pipe\\HyperDbgOutput\n");
    ShowMessages("\t\te.g : output open MyOutputName1\n");
    ShowMessages("\t\te.g : output close MyOutputName1\n");
}*/
return true
}

func (o *output)CommandOutput()(ok bool){//col:462
/*CommandOutput(vector<string> SplittedCommand, string Command)
{
    PDEBUGGER_EVENT_FORWARDING     EventForwardingObject;
    DEBUGGER_EVENT_FORWARDING_TYPE Type;
    DEBUGGER_OUTPUT_SOURCE_STATUS  Status;
    string                         DetailsOfSource;
    UINT32                         IndexToShowList;
    PLIST_ENTRY                    TempList          = 0;
    BOOLEAN                        OutputSourceFound = FALSE;
    HANDLE                         SourceHandle      = INVALID_HANDLE_VALUE;
    SOCKET                         Socket            = NULL;
    vector<string>                 SplittedCommandCaseSensitive {Split(Command, ' ')};
    if (SplittedCommand.size() == 1)
    {
        IndexToShowList = 0;
        if (g_OutputSourcesInitialized)
        {
            TempList = &g_OutputSources;
            while (&g_OutputSources != TempList->Blink)
            {
                TempList = TempList->Blink;
                PDEBUGGER_EVENT_FORWARDING CurrentOutputSourceDetails =
                    CONTAINING_RECORD(TempList, DEBUGGER_EVENT_FORWARDING, OutputSourcesList);
                IndexToShowList++;
                string TempStateString = "";
                string TempTypeString  = "";
                if (CurrentOutputSourceDetails->State ==
                    EVENT_FORWARDING_STATE_NOT_OPENED)
                {
                    TempStateString = "not opened";
                }
                else if (CurrentOutputSourceDetails->State ==
                         EVENT_FORWARDING_STATE_OPENED)
                {
                    TempStateString = "opened    ";
                }
                else if (CurrentOutputSourceDetails->State ==
                         EVENT_FORWARDING_CLOSED)
                {
                    TempStateString = "closed    ";
                }
                if (CurrentOutputSourceDetails->Type == EVENT_FORWARDING_NAMEDPIPE)
                {
                    TempTypeString = "namedpipe";
                }
                else if (CurrentOutputSourceDetails->Type == EVENT_FORWARDING_FILE)
                {
                    TempTypeString = "file     ";
                }
                else if (CurrentOutputSourceDetails->Type == EVENT_FORWARDING_TCP)
                {
                    TempTypeString = "tcp      ";
                }
                ShowMessages("%x  %s   %s\t%s\n", IndexToShowList, TempTypeString.c_str(), TempStateString.c_str(), CurrentOutputSourceDetails->Name);
            }
        }
        else
        {
            ShowMessages("output forwarding list is empty\n");
        }
        return;
    }
    if (SplittedCommand.size() <= 2)
    {
        ShowMessages("incorrect use of 'output'\n\n");
        CommandOutputHelp();
        return;
    }
    if (!SplittedCommand.at(1).compare("create"))
    {
        if (SplittedCommand.size() <= 4)
        {
            ShowMessages("incorrect use of 'output'\n\n");
            CommandOutputHelp();
            return;
        }
        if (!SplittedCommand.at(3).compare("file"))
        {
            Type = EVENT_FORWARDING_FILE;
        }
        else if (!SplittedCommand.at(3).compare("namedpipe"))
        {
            Type = EVENT_FORWARDING_NAMEDPIPE;
        }
        else if (!SplittedCommand.at(3).compare("tcp"))
        {
            Type = EVENT_FORWARDING_TCP;
        }
        else
        {
            ShowMessages("incorrect type near '%s'\n\n",
                         SplittedCommand.at(3).c_str());
            CommandOutputHelp();
            return;
        }
        if (SplittedCommand.at(2).size() >=
            MAXIMUM_CHARACTERS_FOR_EVENT_FORWARDING_NAME)
        {
            ShowMessages("name of the output cannot exceed form %d characters\n\n",
                         MAXIMUM_CHARACTERS_FOR_EVENT_FORWARDING_NAME);
            CommandOutputHelp();
            return;
        }
        if (g_OutputSourcesInitialized)
        {
            TempList = &g_OutputSources;
            while (&g_OutputSources != TempList->Flink)
            {
                TempList = TempList->Flink;
                PDEBUGGER_EVENT_FORWARDING CurrentOutputSourceDetails =
                    CONTAINING_RECORD(TempList, DEBUGGER_EVENT_FORWARDING, OutputSourcesList);
                if (strcmp(CurrentOutputSourceDetails->Name,
                           SplittedCommandCaseSensitive.at(2).c_str()) == 0)
                {
                    OutputSourceFound = TRUE;
                    break;
                }
            }
            if (OutputSourceFound)
            {
                ShowMessages("err, the name you entered, already exists, please choose "
                             "another name\n");
                return;
            }
        }
        DetailsOfSource = Command.substr(Command.find(SplittedCommandCaseSensitive.at(3)) +
                                             SplittedCommandCaseSensitive.at(3).size() + 1,
                                         Command.size());
        SourceHandle = ForwardingCreateOutputSource(Type, DetailsOfSource, &Socket);
        if (SourceHandle == INVALID_HANDLE_VALUE)
        {
            ShowMessages(
                "err, invalid address or cannot open or find the address\n");
            return;
        }
        EventForwardingObject =
            (PDEBUGGER_EVENT_FORWARDING)malloc(sizeof(DEBUGGER_EVENT_FORWARDING));
        if (EventForwardingObject == NULL)
        {
            ShowMessages("err, in allocating memory for event forwarding\n");
            return;
        }
        RtlZeroMemory(EventForwardingObject, sizeof(DEBUGGER_EVENT_FORWARDING));
        EventForwardingObject->State = EVENT_FORWARDING_STATE_NOT_OPENED;
        EventForwardingObject->Type = Type;
        EventForwardingObject->OutputUniqueTag = ForwardingGetNewOutputSourceTag();
        if (Type == EVENT_FORWARDING_TCP)
        {
            EventForwardingObject->Socket = Socket;
        }
        else
        {
            EventForwardingObject->Handle = SourceHandle;
        }
        strcpy_s(EventForwardingObject->Name, SplittedCommandCaseSensitive.at(2).c_str());
        if (!g_OutputSourcesInitialized)
        {
            g_OutputSourcesInitialized = TRUE;
            InitializeListHead(&g_OutputSources);
        }
        InsertHeadList(&g_OutputSources,
                       &(EventForwardingObject->OutputSourcesList));
    }
    else if (!SplittedCommand.at(1).compare("open"))
    {
        if (!g_OutputSourcesInitialized)
        {
            ShowMessages("err, the name you entered, not found\n");
            return;
        }
        TempList = &g_OutputSources;
        while (&g_OutputSources != TempList->Flink)
        {
            TempList = TempList->Flink;
            PDEBUGGER_EVENT_FORWARDING CurrentOutputSourceDetails = CONTAINING_RECORD(
                TempList,
                DEBUGGER_EVENT_FORWARDING,
                OutputSourcesList);
            if (strcmp(CurrentOutputSourceDetails->Name,
                       SplittedCommandCaseSensitive.at(2).c_str()) == 0)
            {
                OutputSourceFound = TRUE;
                Status = ForwardingOpenOutputSource(CurrentOutputSourceDetails);
                if (Status == DEBUGGER_OUTPUT_SOURCE_STATUS_ALREADY_CLOSED)
                {
                    ShowMessages("err, the name you entered was already closed\n");
                    return;
                }
                else if (Status == DEBUGGER_OUTPUT_SOURCE_STATUS_ALREADY_OPENED)
                {
                    ShowMessages("err, the name you entered was already opened\n");
                    return;
                }
                else if (Status !=
                         DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_OPENED)
                {
                    ShowMessages("err, unable to open the output source\n");
                    return;
                }
                break;
            }
        }
        if (!OutputSourceFound)
        {
            ShowMessages("err, the name you entered, not found\n");
            return;
        }
    }
    else if (!SplittedCommand.at(1).compare("close"))
    {
        if (!g_OutputSourcesInitialized)
        {
            ShowMessages("err, the name you entered, not found\n");
            return;
        }
        TempList = &g_OutputSources;
        while (&g_OutputSources != TempList->Flink)
        {
            TempList = TempList->Flink;
            PDEBUGGER_EVENT_FORWARDING CurrentOutputSourceDetails = CONTAINING_RECORD(
                TempList,
                DEBUGGER_EVENT_FORWARDING,
                OutputSourcesList);
            if (strcmp(CurrentOutputSourceDetails->Name,
                       SplittedCommandCaseSensitive.at(2).c_str()) == 0)
            {
                OutputSourceFound = TRUE;
                Status = ForwardingCloseOutputSource(CurrentOutputSourceDetails);
                if (Status == DEBUGGER_OUTPUT_SOURCE_STATUS_ALREADY_CLOSED)
                {
                    ShowMessages("err, the name you entered was already closed\n");
                    return;
                }
                else if (Status == DEBUGGER_OUTPUT_SOURCE_STATUS_UNKNOWN_ERROR)
                {
                    ShowMessages("err, unable to close the source\n");
                    return;
                }
                else if (Status !=
                         DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_CLOSED)
                {
                    ShowMessages("err, unable to close the source\n");
                    return;
                }
                break;
            }
        }
        if (!OutputSourceFound)
        {
            ShowMessages("err, the name you entered, not found\n");
            return;
        }
    }
    else
    {
        ShowMessages("incorrect option at '%s'\n\n", SplittedCommand.at(1).c_str());
        CommandOutputHelp();
        return;
    }
}*/
return true
}



