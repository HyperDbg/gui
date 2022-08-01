package meta_commands

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\meta-commands\script.cpp.back

type (
	Script interface {
		CommandScriptHelp() (ok bool)                       //col:12
		CommandScriptRunCommand() (ok bool)                 //col:38
		HyperDbgScriptReadFileAndExecuteCommand() (ok bool) //col:86
		CommandScript() (ok bool)                           //col:101
	}
	script struct{}
)

func NewScript() Script { return &script{} }

func (s *script) CommandScriptHelp() (ok bool) { //col:12
	/*
	   CommandScriptHelp()

	   	{
	   	    ShowMessages(".script : runs a HyperDbg script.\n\n");
	   	    ShowMessages("syntax : \.script [FilePath (string)] [Args (string)]\n");
	   	    ShowMessages("\n");
	   	    ShowMessages("\t\te.g : .script C:\\scripts\\script.ds\n");
	   	    ShowMessages("\t\te.g : .script \"C:\\scripts\\hello world.ds\"\n");
	   	}
	*/
	return true
}

func (s *script) CommandScriptRunCommand() (ok bool) { //col:38
	/*
	   CommandScriptRunCommand(std::string Input, vector<string> PathAndArgs)

	   	{
	   	    int    CommandExecutionResult = 0;
	   	    char * LineContent            = NULL;
	   	    int    i                      = 0;
	   	    for (auto item : PathAndArgs)
	   	    {
	   	        string ToReplace = "$arg" + std::to_string(i);
	   	        i++;
	   	        ReplaceAll(Input, ToReplace, item);
	   	    }
	   	    LineContent = (char *)Input.c_str();
	   	    InterpreterRemoveComments(LineContent);
	   	    if (IsEmptyString(LineContent))
	   	    {
	   	        return;
	   	    }
	   	    HyperDbgShowSignature();
	   	    ShowMessages("%s\n", LineContent);
	   	    CommandExecutionResult = HyperDbgInterpreter(LineContent);
	   	    ShowMessages("\n");
	   	    if (CommandExecutionResult == 1)
	   	    {
	   	        exit(0);
	   	    }
	   	}
	*/
	return true
}

func (s *script) HyperDbgScriptReadFileAndExecuteCommand() (ok bool) { //col:86
	/*
	   HyperDbgScriptReadFileAndExecuteCommand(std::vector<std::string> & PathAndArgs)

	   	{
	   	    std::string Line;
	   	    BOOLEAN     IsOpened         = FALSE;
	   	    bool        Reset            = false;
	   	    string      CommandToExecute = "";
	   	    string      PathOfScriptFile = "";
	   	    PathOfScriptFile = PathAndArgs.at(0);
	   	    ReplaceAll(PathOfScriptFile, "\"", "");
	   	    ifstream File(PathOfScriptFile);
	   	    if (File.is_open())
	   	    {
	   	        IsOpened = TRUE;
	   	        g_ExecutingScript = TRUE;
	   	        Reset = true;
	   	        while (std::getline(File, Line))
	   	        {
	   	            if (HyperDbgCheckMultilineCommand(Line, Reset))
	   	            {
	   	                if (Reset)
	   	                {
	   	                    CommandToExecute.clear();
	   	                }
	   	                Reset = false;
	   	                CommandToExecute += Line + "\n";
	   	                continue;
	   	            }
	   	            else
	   	            {
	   	                Reset = true;
	   	                CommandToExecute += Line;
	   	            }
	   	            CommandScriptRunCommand(CommandToExecute, PathAndArgs);
	   	            CommandToExecute.clear();
	   	        }
	   	        if (!CommandToExecute.empty())
	   	        {
	   	            CommandScriptRunCommand(CommandToExecute, PathAndArgs);
	   	            CommandToExecute.clear();
	   	        }
	   	        g_ExecutingScript = FALSE;
	   	        File.close();
	   	    }
	   	    if (!IsOpened)
	   	    {
	   	        ShowMessages("err, invalid file specified for the script\n");
	   	    }
	   	}
	*/
	return true
}

func (s *script) CommandScript() (ok bool) { //col:101
	/*
	   CommandScript(vector<string> SplittedCommand, string Command)

	   	{
	   	    vector<string> PathAndArgs;
	   	    if (SplittedCommand.size() == 1)
	   	    {
	   	        ShowMessages("please specify a file\n");
	   	        CommandScriptHelp();
	   	        return;
	   	    }
	   	    Trim(Command);
	   	    Command.erase(0, 7);
	   	    Trim(Command);
	   	    SplitPathAndArgs(PathAndArgs, Command);
	   	    HyperDbgScriptReadFileAndExecuteCommand(PathAndArgs);
	   	}
	*/
	return true
}

