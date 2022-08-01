package debugging_commands

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\k.cpp.back

type (
	K interface {
		CommandKHelp() (ok bool) //col:18
		CommandK() (ok bool)     //col:121
	}
	k struct{}
)

func NewK() K { return &k{} }

func (k *k) CommandKHelp() (ok bool) { //col:18
	/*
	   CommandKHelp()

	   	{
	   	    ShowMessages(
	   	        "k, kd, kq : shows the callstack of the thread.\n\n");
	   	    ShowMessages("syntax : \tk\n");
	   	    ShowMessages("syntax : \tkd\n");
	   	    ShowMessages("syntax : \tkq\n");
	   	    ShowMessages("syntax : \tk [base StackAddress (hex)] [l Length (hex)]\n");
	   	    ShowMessages("syntax : \tkd [base StackAddress (hex)] [l Length (hex)]\n");
	   	    ShowMessages("syntax : \tkq [base StackAddress (hex)] [l Length (hex)]\n");
	   	    ShowMessages("\n");
	   	    ShowMessages("\t\te.g : k\n");
	   	    ShowMessages("\t\te.g : k l 100\n");
	   	    ShowMessages("\t\te.g : kd base 0x77356f010\n");
	   	    ShowMessages("\t\te.g : kq base fffff8077356f010\n");
	   	    ShowMessages("\t\te.g : kq base fffff8077356f010 l 100\n");
	   	}
	*/
	return true
}

func (k *k) CommandK() (ok bool) { //col:121
	/*
	   CommandK(vector<string> SplittedCommand, string Command)

	   	{
	   	    UINT64         BaseAddress = NULL;
	   	    UINT32         Length      = 0x100;
	   	    vector<string> SplittedCommandCaseSensitive {Split(Command, ' ')};
	   	    UINT32         IndexInCommandCaseSensitive = 0;
	   	    BOOLEAN        IsFirstCommand              = TRUE;
	   	    BOOLEAN        IsNextBase                  = FALSE;
	   	    BOOLEAN        IsNextLength                = FALSE;
	   	    string FirstCommand = SplittedCommand.front();
	   	    if (SplittedCommand.size() >= 6)
	   	    {
	   	        ShowMessages("incorrect use of '%s'\n\n", FirstCommand.c_str());
	   	        CommandKHelp();
	   	        return;
	   	    }
	   	    if (!g_IsSerialConnectedToRemoteDebuggee)
	   	    {
	   	        ShowMessages("err, tracing callstack is not possible when you're not "
	   	                     "connected to a debuggee\n");
	   	        return;
	   	    }
	   	    if (g_IsRunningInstruction32Bit)
	   	    {
	   	        Length = 0x100;
	   	    }
	   	    else
	   	    {
	   	        Length = 0x200;
	   	    }
	   	    for (auto Section : SplittedCommand)
	   	    {
	   	        IndexInCommandCaseSensitive++;
	   	        if (IsFirstCommand)
	   	        {
	   	            IsFirstCommand = FALSE;
	   	            continue;
	   	        }
	   	        if (IsNextBase == TRUE)
	   	        {
	   	            if (!SymbolConvertNameOrExprToAddress(SplittedCommandCaseSensitive.at(IndexInCommandCaseSensitive - 1),
	   	                                                  &BaseAddress))
	   	            {
	   	                ShowMessages("err, couldn't resolve error at '%s'\n",
	   	                             SplittedCommandCaseSensitive.at(IndexInCommandCaseSensitive - 1).c_str());
	   	                return;
	   	            }
	   	            IsNextBase = FALSE;
	   	            continue;
	   	        }
	   	        if (IsNextLength == TRUE)
	   	        {
	   	            if (!ConvertStringToUInt32(Section, &Length))
	   	            {
	   	                ShowMessages("err, you should enter a valid length\n\n");
	   	                return;
	   	            }
	   	            IsNextLength = FALSE;
	   	            continue;
	   	        }
	   	        if (!Section.compare("l"))
	   	        {
	   	            IsNextLength = TRUE;
	   	            continue;
	   	        }
	   	        if (!Section.compare("base"))
	   	        {
	   	            IsNextBase = TRUE;
	   	            continue;
	   	        }
	   	        ShowMessages("err, incorrect use of '%s' command\n\n",
	   	                     FirstCommand.c_str());
	   	        CommandKHelp();
	   	        return;
	   	    }
	   	    if (IsNextLength || IsNextBase)
	   	    {
	   	        ShowMessages("incorrect use of '%s' command\n\n", FirstCommand.c_str());
	   	        CommandKHelp();
	   	        return;
	   	    }
	   	    if (!FirstCommand.compare("k"))
	   	    {
	   	        KdSendCallStackPacketToDebuggee(BaseAddress,
	   	                                        Length,
	   	                                        DEBUGGER_CALLSTACK_DISPLAY_METHOD_WITHOUT_PARAMS,
	   	                                        g_IsRunningInstruction32Bit);
	   	    }
	   	    else if (!FirstCommand.compare("kq"))
	   	    {
	   	        KdSendCallStackPacketToDebuggee(BaseAddress,
	   	                                        Length,
	   	                                        DEBUGGER_CALLSTACK_DISPLAY_METHOD_WITH_PARAMS,
	   	                                        FALSE);
	   	    }
	   	    else if (!FirstCommand.compare("kd"))
	   	    {
	   	        KdSendCallStackPacketToDebuggee(BaseAddress,
	   	                                        Length,
	   	                                        DEBUGGER_CALLSTACK_DISPLAY_METHOD_WITH_PARAMS,
	   	                                        TRUE);
	   	    }
	   	}
	*/
	return true
}

