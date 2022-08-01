package debugging_commands

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\eval.cpp.back

type (
	Eval interface {
		CommandEvalHelp() (ok bool)          //col:8
		CommandEvalCheckTestcase() (ok bool) //col:95
		CommandEval() (ok bool)              //col:145
	}
	eval struct{}
)

func NewEval() Eval { return &eval{} }

func (e *eval) CommandEvalHelp() (ok bool) { //col:8
	/*
	   CommandEvalHelp()

	   	{
	   	    ShowMessages("? : evaluates and execute expressions in debuggee.\n\n");
	   	    ShowMessages("syntax : \t? [Expression (string)]\n");
	   	    ShowMessages("\n");
	   	}
	*/
	return true
}

func (e *eval) CommandEvalCheckTestcase() (ok bool) { //col:95
	/*
	   CommandEvalCheckTestcase()

	   	{
	   	    string                   Line;
	   	    BOOLEAN                  IsOpened      = FALSE;
	   	    UINT64                   ExpectedValue = 0;
	   	    BOOLEAN                  ExpectError   = FALSE;
	   	    string                   Expr          = "";
	   	    std::vector<std::string> TestFiles;
	   	    try
	   	    {
	   	        TestFiles = ListDirectory(SCRIPT_ENGINE_TEST_CASES_DIRECTORY, "*.txt");
	   	    }
	   	    catch (const std::exception &)
	   	    {
	   	        ShowMessages("err, test-cases not found, make sure to run test-environment.py "
	   	                     "before running test cases\n");
	   	        return FALSE;
	   	    }
	   	    for (auto item : TestFiles)
	   	    {
	   	        ifstream File(item.c_str());
	   	        if (File.is_open())
	   	        {
	   	            IsOpened = TRUE;
	   	            ShowMessages("Running from : %s\n\n", item.c_str());
	   	            while (std::getline(File, Line))
	   	            {
	   	                ShowMessages("Test-case number : %s\n", Line.c_str());
	   	                if (!std::getline(File, Line))
	   	                {
	   	                    goto ErrorMessage;
	   	                }
	   	                Expr = Line;
	   	                ShowMessages("Statement : %s\n", Line.c_str());
	   	                if (!std::getline(File, Line))
	   	                {
	   	                    goto ErrorMessage;
	   	                }
	   	                if (!Line.compare("$error$"))
	   	                {
	   	                    ShowMessages("Expected result : %s\n", Line.c_str());
	   	                    ExpectError   = TRUE;
	   	                    ExpectedValue = NULL;
	   	                }
	   	                else if (!ConvertStringToUInt64(Line, &ExpectedValue))
	   	                {
	   	                    ShowMessages("err, the expected results are in incorrect format\n");
	   	                    goto ErrorMessage;
	   	                }
	   	                else
	   	                {
	   	                    ExpectError = FALSE;
	   	                    ShowMessages("Expected result : %llx\n", ExpectedValue);
	   	                }
	   	                Expr.append(" ");
	   	                if (ScriptAutomaticStatementsTestWrapper(Expr, ExpectedValue, ExpectError))
	   	                {
	   	                    ShowMessages("Test result : Passed\n");
	   	                }
	   	                else
	   	                {
	   	                    ShowMessages("Test result : Failed\n");
	   	                }
	   	                if (!std::getline(File, Line))
	   	                {
	   	                    goto ErrorMessage;
	   	                }
	   	                if (Line.compare("$end$"))
	   	                {
	   	                    goto ErrorMessage;
	   	                }
	   	                ShowMessages("\n------------------------------------------------------------\n\n");
	   	            }
	   	            File.close();
	   	        }
	   	    }
	   	    if (!IsOpened)
	   	    {
	   	        ShowMessages("err, could not find files for script engine test-cases\n");
	   	        return FALSE;
	   	    }
	   	    return TRUE;

	   ErrorMessage:

	   	    ShowMessages("\nerr, testing failed! incorrect file format for the testing "
	   	                 "script engine\nmake sure files have a correct ending format\n");
	   	    return FALSE;
	   	}
	*/
	return true
}

func (e *eval) CommandEval() (ok bool) { //col:145
	/*
	   CommandEval(vector<string> SplittedCommand, string Command)

	   	{
	   	    PVOID  CodeBuffer;
	   	    UINT64 BufferAddress;
	   	    UINT32 BufferLength;
	   	    UINT32 Pointer;
	   	    if (SplittedCommand.size() == 1)
	   	    {
	   	        ShowMessages("incorrect use of '?'\n\n");
	   	        CommandEvalHelp();
	   	        return;
	   	    }
	   	    Trim(Command);
	   	    Command.erase(0, SplittedCommand.at(0).size());
	   	    Trim(Command);
	   	    if (!Command.compare("test"))
	   	    {
	   	        if (!CommandEvalCheckTestcase())
	   	        {
	   	            ShowMessages("err, testing script engine test-cases was not successful\n");
	   	        }
	   	        else
	   	        {
	   	            ShowMessages("testing script engine test-cases was successful\n");
	   	        }
	   	        return;
	   	    }
	   	    if (g_IsSerialConnectedToRemoteDebuggee)
	   	    {
	   	        CodeBuffer = ScriptEngineParseWrapper((char *)Command.c_str(), TRUE);
	   	        if (CodeBuffer == NULL)
	   	        {
	   	            return;
	   	        }
	   	        BufferAddress = ScriptEngineWrapperGetHead(CodeBuffer);
	   	        BufferLength  = ScriptEngineWrapperGetSize(CodeBuffer);
	   	        Pointer       = ScriptEngineWrapperGetPointer(CodeBuffer);
	   	        KdSendScriptPacketToDebuggee(BufferAddress, BufferLength, Pointer, FALSE);
	   	        ScriptEngineWrapperRemoveSymbolBuffer(CodeBuffer);
	   	    }
	   	    else
	   	    {
	   	        ShowMessages("this command should not be used while you're in VMI-Mode or not in debugger-mode, "
	   	                     "the results that you see is a simulated result for TESTING script-engine "
	   	                     "and is not based on the status of your system. You can use this command, "
	   	                     "ONLY in debugger-mode\n\n");
	   	        ShowMessages("test expression : %s \n", Command.c_str());
	   	        ScriptEngineWrapperTestParser(Command);
	   	    }
	   	}
	*/
	return true
}

