package meta_commands

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\meta-commands\start.cpp.back

type (
	Start interface {
		CommandStartHelp() (ok bool) //col:7
		CommandStart() (ok bool)     //col:66
	}
	start struct{}
)

func NewStart() Start { return &start{} }

func (s *start) CommandStartHelp() (ok bool) { //col:7
	/*
	   CommandStartHelp()

	   	{
	   	    ShowMessages(".start : runs a user-mode process.\n\n");
	   	    ShowMessages("syntax : \t.start [path Path (string)] [Parameters (string)]\n");
	   	    ShowMessages("\n");
	   	    ShowMessages("\t\te.g : .start path c:\\users\\sina\\reverse eng\\my_file.exe\n");
	   	}
	*/
	return true
}

func (s *start) CommandStart() (ok bool) { //col:66
	/*
	   CommandStart(vector<string> SplittedCommand, string Command)

	   	{
	   	    vector<string> PathAndArgs;
	   	    string         Arguments = "";

	   #if ActivateUserModeDebugger == FALSE

	   	if (!g_IsSerialConnectedToRemoteDebugger)
	   	{
	   	    ShowMessages("The user-mode debugger is still in the beta version and not stable. "
	   	                 "We decided to exclude it from this release and release it in future versions. "
	   	                 "If you want to test the user-mode debugger in VMI Mode, you should build "
	   	                 "HyperDbg with special instructions. \nPlease follow the steps here: "
	   	                 "https:
	   	    return;
	   	}

	   #endif

	   	    if (SplittedCommand.size() <= 2)
	   	    {
	   	        ShowMessages("incorrect use of '.start'\n\n");
	   	        CommandStartHelp();
	   	        return;
	   	    }
	   	    if (!SplittedCommand.at(1).compare("path"))
	   	    {
	   	        Trim(Command);
	   	        Command.erase(0, 6);
	   	        Command.erase(0, 4 + 1);
	   	        Trim(Command);
	   	        SplitPathAndArgs(PathAndArgs, Command);
	   	        StringToWString(g_StartCommandPath, PathAndArgs.at(0));
	   	        if (PathAndArgs.size() != 1)
	   	        {
	   	            for (auto item : PathAndArgs)
	   	            {
	   	                Arguments += item + " ";
	   	            }
	   	            Arguments.pop_back();
	   	            StringToWString(g_StartCommandPathAndArguments, Arguments);
	   	        }
	   	    }
	   	    else
	   	    {
	   	        ShowMessages("err, couldn't resolve error at '%s'\n\n",
	   	                     SplittedCommand.at(1).c_str());
	   	        CommandStartHelp();
	   	        return;
	   	    }
	   	    if (Arguments.empty())
	   	    {
	   	        UdAttachToProcess(NULL,
	   	                          g_StartCommandPath.c_str(),
	   	                          NULL);
	   	    }
	   	    else
	   	    {
	   	        UdAttachToProcess(NULL,
	   	                          g_StartCommandPath.c_str(),
	   	                          (WCHAR *)g_StartCommandPathAndArguments.c_str());
	   	    }
	   	}
	*/
	return true
}

