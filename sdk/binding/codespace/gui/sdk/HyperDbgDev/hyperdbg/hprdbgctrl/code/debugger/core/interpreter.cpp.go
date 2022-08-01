package core
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\core\interpreter.cpp.back

type (
Interpreter interface{
HyperDbgInterpreter()(ok bool)//col:116
InterpreterRemoveComments()(ok bool)//col:159
HyperDbgShowSignature()(ok bool)//col:181
HyperDbgCheckMultilineCommand()(ok bool)//col:237
HyperDbgContinuePreviousCommand()(ok bool)//col:250
GetCommandAttributes()(ok bool)//col:264
InitializeDebugger()(ok bool)//col:277
InitializeCommandsDictionary()(ok bool)//col:447
}
)

func NewInterpreter() { return & interpreter{} }

func (i *interpreter)HyperDbgInterpreter()(ok bool){//col:116
/*HyperDbgInterpreter(char * Command)
{
    BOOLEAN               HelpCommand       = FALSE;
    UINT64                CommandAttributes = NULL;
    CommandType::iterator Iterator;
    if (!g_IsCommandListInitialized)
    {
        InitializeDebugger();
        g_IsCommandListInitialized = TRUE;
    }
    if (g_LogOpened && !g_ExecutingScript)
    {
        LogopenSaveToFile(Command);
        LogopenSaveToFile("\n");
    }
    InterpreterRemoveComments(Command);
    string CommandString(Command);
    transform(CommandString.begin(), CommandString.end(), CommandString.begin(), [](unsigned char c) { return std::tolower(c); });
    vector<string> SplittedCommand {Split(CommandString, ' ')};
    if (SplittedCommand.empty())
    {
        ShowMessages("\n");
        return 0;
    }
    string FirstCommand = SplittedCommand.front();
    CommandAttributes = GetCommandAttributes(FirstCommand);
    if (CommandAttributes & DEBUGGER_COMMAND_ATTRIBUTE_REPEAT_ON_ENTER)
    {
        g_ShouldPreviousCommandBeContinued = TRUE;
    }
    else
    {
        g_ShouldPreviousCommandBeContinued = FALSE;
    }
    if (g_IsConnectedToRemoteDebuggee &&
        !(CommandAttributes & DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_REMOTE_CONNECTION))
    {
        if (g_BreakPrintingOutput)
        {
            g_BreakPrintingOutput = FALSE;
        }
        RemoteConnectionSendCommand(Command, strlen(Command) + 1);
        ShowMessages("\n");
        return 2;
    }
    else if (g_IsSerialConnectedToRemoteDebuggee &&
             !(CommandAttributes &
               DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE))
    {
        if (CommandAttributes & DEBUGGER_COMMAND_ATTRIBUTE_WONT_STOP_DEBUGGER_AGAIN)
        {
            KdSendUserInputPacketToDebuggee(Command, strlen(Command) + 1, TRUE);
            KdSetStatusAndWaitForPause();
        }
        else
        {
            KdSendUserInputPacketToDebuggee(Command, strlen(Command) + 1, FALSE);
        }
        return 2;
    }
    if (!FirstCommand.compare(".help") || !FirstCommand.compare("help") ||
        !FirstCommand.compare(".hh"))
    {
        if (SplittedCommand.size() == 2)
        {
            HelpCommand  = TRUE;
            FirstCommand = SplittedCommand.at(1);
        }
        else
        {
            ShowMessages("incorrect use of '%s'\n", FirstCommand.c_str());
            CommandHelpHelp();
            return 0;
        }
    }
    Iterator = g_CommandsList.find(FirstCommand);
    if (Iterator == g_CommandsList.end())
    {
        string         CaseSensitiveCommandString(Command);
        vector<string> CaseSensitiveSplittedCommand {Split(CaseSensitiveCommandString, ' ')};
        if (!HelpCommand)
        {
            ShowMessages("err, couldn't resolve command at '%s'\n", CaseSensitiveSplittedCommand.front().c_str());
        }
        else
        {
            ShowMessages("err, couldn't find the help for the command at '%s'\n",
                         CaseSensitiveSplittedCommand.at(1).c_str());
        }
    }
    else
    {
        if (HelpCommand)
        {
            Iterator->second.CommandHelpFunction();
        }
        else
        {
            if ((Iterator->second.CommandAttrib &
                 DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE))
            {
                string CaseSensitiveCommandString(Command);
                Iterator->second.CommandFunction(SplittedCommand, CaseSensitiveCommandString);
            }
            else
            {
                Iterator->second.CommandFunction(SplittedCommand, CommandString);
            }
        }
    }
    if (g_LogOpened && !g_ExecutingScript)
    {
        LogopenSaveToFile("\n");
    }
    return 0;
}*/
return true
}

func (i *interpreter)InterpreterRemoveComments()(ok bool){//col:159
/*InterpreterRemoveComments(char * CommandText)
{
    BOOLEAN IsComment       = FALSE;
    BOOLEAN IsOnString      = FALSE;
    UINT32  LengthOfCommand = strlen(CommandText);
    for (size_t i = 0; i < LengthOfCommand; i++)
    {
        if (IsComment)
        {
            if (CommandText[i] == '\n')
            {
                IsComment = FALSE;
            }
            else
            {
                if (CommandText[i] != '\0')
                {
                    memmove((void *)&CommandText[i], (const void *)&CommandText[i + 1], strlen(CommandText) - i);
                    i--;
                }
            }
        }
        else if (CommandText[i] == '#' && !IsOnString)
        {
            IsComment = TRUE;
            i--;
        }
        else if (CommandText[i] == '"')
        {
            if (i != 0 && CommandText[i - 1] == '\\')
            {
            }
            else if (IsOnString)
            {
                IsOnString = FALSE;
            }
            else
            {
                IsOnString = TRUE;
            }
        }
    }
}*/
return true
}

func (i *interpreter)HyperDbgShowSignature()(ok bool){//col:181
/*HyperDbgShowSignature()
{
    if (g_IsConnectedToRemoteDebuggee)
    {
        ShowMessages("[%s:%s] HyperDbg> ", g_ServerIp.c_str(), g_ServerPort.c_str());
    }
    else if (g_ActiveProcessDebuggingState.IsActive)
    {
        ShowMessages("%x:%x u%sHyperDbg> ",
                     g_ActiveProcessDebuggingState.ProcessId,
                     g_ActiveProcessDebuggingState.ThreadId,
                     g_ActiveProcessDebuggingState.Is32Bit ? "86" : "64");
    }
    else if (g_IsSerialConnectedToRemoteDebuggee)
    {
        ShowMessages("%x: kHyperDbg> ", g_CurrentRemoteCore);
    }
    else
    {
        ShowMessages("HyperDbg> ");
    }
}*/
return true
}

func (i *interpreter)HyperDbgCheckMultilineCommand()(ok bool){//col:237
/*HyperDbgCheckMultilineCommand(std::string & CurrentCommand, bool Reset)
{
    if (Reset)
    {
        g_IsInterpreterOnString                    = FALSE;
        g_IsInterpreterPreviousCharacterABackSlash = FALSE;
        g_InterpreterCountOfOpenCurlyBrackets      = 0;
    }
    for (size_t i = 0; i < CurrentCommand.length(); i++)
    {
        switch (CurrentCommand.at(i))
        {
        case '"':
            if (g_IsInterpreterPreviousCharacterABackSlash)
            {
                g_IsInterpreterPreviousCharacterABackSlash = FALSE;
                break; 
            }
            if (g_IsInterpreterOnString)
                g_IsInterpreterOnString = FALSE;
            else
                g_IsInterpreterOnString = TRUE;
            break;
        case '{':
            if (g_IsInterpreterPreviousCharacterABackSlash)
                g_IsInterpreterPreviousCharacterABackSlash = FALSE;
            if (!g_IsInterpreterOnString)
                g_InterpreterCountOfOpenCurlyBrackets++;
            break;
        case '}':
            if (g_IsInterpreterPreviousCharacterABackSlash)
                g_IsInterpreterPreviousCharacterABackSlash = FALSE;
            if (!g_IsInterpreterOnString && g_InterpreterCountOfOpenCurlyBrackets > 0)
                g_InterpreterCountOfOpenCurlyBrackets--;
            break;
        case '\\':
            if (g_IsInterpreterPreviousCharacterABackSlash)
                g_IsInterpreterPreviousCharacterABackSlash = FALSE; 
            else
                g_IsInterpreterPreviousCharacterABackSlash = TRUE;
            break;
        default:
            if (g_IsInterpreterPreviousCharacterABackSlash)
                g_IsInterpreterPreviousCharacterABackSlash = FALSE;
            break;
        }
    }
    if (g_IsInterpreterOnString == FALSE && g_InterpreterCountOfOpenCurlyBrackets == 0)
    {
        return false;
    }
    else
    {
        return true;
    }
}*/
return true
}

func (i *interpreter)HyperDbgContinuePreviousCommand()(ok bool){//col:250
/*HyperDbgContinuePreviousCommand()
{
    BOOLEAN Result = g_ShouldPreviousCommandBeContinued;
    g_ShouldPreviousCommandBeContinued = FALSE;
    if (Result)
    {
        return true;
    }
    else
    {
        return false;
    }
}*/
return true
}

func (i *interpreter)GetCommandAttributes()(ok bool){//col:264
/*GetCommandAttributes(const string & FirstCommand)
{
    CommandType::iterator Iterator;
    Iterator = g_CommandsList.find(FirstCommand);
    if (Iterator == g_CommandsList.end())
    {
        return DEBUGGER_COMMAND_ATTRIBUTE_ABSOLUTE_LOCAL;
    }
    else
    {
        return Iterator->second.CommandAttrib;
    }
    return NULL;
}*/
return true
}

func (i *interpreter)InitializeDebugger()(ok bool){//col:277
/*InitializeDebugger()
{
    InitializeCommandsDictionary();
    ScriptEngineSetTextMessageCallbackWrapper(ShowMessages);
    if (!SetConsoleCtrlHandler(BreakController, TRUE))
    {
        ShowMessages("err, when registering CTRL+C and CTRL+BREAK Signals "
                     "handler\n");
    }
    g_VirtualAddressWidth = Getx86VirtualAddressWidth();
    g_RtmSupport = CheckCpuSupportRtm();
    CommandSettingsLoadDefaultValuesFromConfigFile();
}*/
return true
}

func (i *interpreter)InitializeCommandsDictionary()(ok bool){//col:447
/*InitializeCommandsDictionary()
{
    g_CommandsList[".help"] = {NULL, &CommandHelpHelp, DEBUGGER_COMMAND_HELP_ATTRIBUTES};
    g_CommandsList[".hh"]   = {NULL, &CommandHelpHelp, DEBUGGER_COMMAND_HELP_ATTRIBUTES};
    g_CommandsList["help"]  = {NULL, &CommandHelpHelp, DEBUGGER_COMMAND_HELP_ATTRIBUTES};
    g_CommandsList["clear"] = {&CommandClearScreen, &CommandClearScreenHelp, DEBUGGER_COMMAND_CLEAR_ATTRIBUTES};
    g_CommandsList[".cls"]  = {&CommandClearScreen, &CommandClearScreenHelp, DEBUGGER_COMMAND_CLEAR_ATTRIBUTES};
    g_CommandsList["cls"]   = {&CommandClearScreen, &CommandClearScreenHelp, DEBUGGER_COMMAND_CLEAR_ATTRIBUTES};
    g_CommandsList[".connect"] = {&CommandConnect, &CommandConnectHelp, DEBUGGER_COMMAND_CONNECT_ATTRIBUTES};
    g_CommandsList["connect"]  = {&CommandConnect, &CommandConnectHelp, DEBUGGER_COMMAND_CONNECT_ATTRIBUTES};
    g_CommandsList[".listen"] = {&CommandListen, &CommandListenHelp, DEBUGGER_COMMAND_LISTEN_ATTRIBUTES};
    g_CommandsList["listen"]  = {&CommandListen, &CommandListenHelp, DEBUGGER_COMMAND_LISTEN_ATTRIBUTES};
    g_CommandsList["g"]  = {&CommandG, &CommandGHelp, DEBUGGER_COMMAND_G_ATTRIBUTES};
    g_CommandsList["go"] = {&CommandG, &CommandGHelp, DEBUGGER_COMMAND_G_ATTRIBUTES};
    g_CommandsList[".attach"] = {&CommandAttach, &CommandAttachHelp, DEBUGGER_COMMAND_ATTACH_ATTRIBUTES};
    g_CommandsList["attach"]  = {&CommandAttach, &CommandAttachHelp, DEBUGGER_COMMAND_ATTACH_ATTRIBUTES};
    g_CommandsList[".detach"] = {&CommandDetach, &CommandDetachHelp, DEBUGGER_COMMAND_DETACH_ATTRIBUTES};
    g_CommandsList["detach"]  = {&CommandDetach, &CommandDetachHelp, DEBUGGER_COMMAND_DETACH_ATTRIBUTES};
    g_CommandsList[".start"] = {&CommandStart, &CommandStartHelp, DEBUGGER_COMMAND_START_ATTRIBUTES};
    g_CommandsList["start"]  = {&CommandStart, &CommandStartHelp, DEBUGGER_COMMAND_START_ATTRIBUTES};
    g_CommandsList[".restart"] = {&CommandRestart, &CommandRestartHelp, DEBUGGER_COMMAND_RESTART_ATTRIBUTES};
    g_CommandsList["restart"]  = {&CommandRestart, &CommandRestartHelp, DEBUGGER_COMMAND_RESTART_ATTRIBUTES};
    g_CommandsList[".switch"] = {&CommandSwitch, &CommandSwitchHelp, DEBUGGER_COMMAND_SWITCH_ATTRIBUTES};
    g_CommandsList["switch"]  = {&CommandSwitch, &CommandSwitchHelp, DEBUGGER_COMMAND_SWITCH_ATTRIBUTES};
    g_CommandsList[".kill"] = {&CommandKill, &CommandKillHelp, DEBUGGER_COMMAND_KILL_ATTRIBUTES};
    g_CommandsList["kill"]  = {&CommandKill, &CommandKillHelp, DEBUGGER_COMMAND_KILL_ATTRIBUTES};
    g_CommandsList[".process"]  = {&CommandProcess, &CommandProcessHelp, DEBUGGER_COMMAND_PROCESS_ATTRIBUTES};
    g_CommandsList[".process2"] = {&CommandProcess, &CommandProcessHelp, DEBUGGER_COMMAND_PROCESS_ATTRIBUTES};
    g_CommandsList[".thread"]  = {&CommandThread, &CommandThreadHelp, DEBUGGER_COMMAND_THREAD_ATTRIBUTES};
    g_CommandsList[".thread2"] = {&CommandThread, &CommandThreadHelp, DEBUGGER_COMMAND_THREAD_ATTRIBUTES};
    g_CommandsList["sleep"] = {&CommandSleep, &CommandSleepHelp, DEBUGGER_COMMAND_SLEEP_ATTRIBUTES};
    g_CommandsList["event"]  = {&CommandEvents, &CommandEventsHelp, DEBUGGER_COMMAND_EVENTS_ATTRIBUTES};
    g_CommandsList["events"] = {&CommandEvents, &CommandEventsHelp, DEBUGGER_COMMAND_EVENTS_ATTRIBUTES};
    g_CommandsList["setting"]   = {&CommandSettings, &CommandSettingsHelp, DEBUGGER_COMMAND_SETTINGS_ATTRIBUTES};
    g_CommandsList["settings"]  = {&CommandSettings, &CommandSettingsHelp, DEBUGGER_COMMAND_SETTINGS_ATTRIBUTES};
    g_CommandsList[".setting"]  = {&CommandSettings, &CommandSettingsHelp, DEBUGGER_COMMAND_SETTINGS_ATTRIBUTES};
    g_CommandsList[".settings"] = {&CommandSettings, &CommandSettingsHelp, DEBUGGER_COMMAND_SETTINGS_ATTRIBUTES};
    g_CommandsList[".disconnect"] = {&CommandDisconnect, &CommandDisconnectHelp, DEBUGGER_COMMAND_DISCONNECT_ATTRIBUTES};
    g_CommandsList["disconnect"]  = {&CommandDisconnect, &CommandDisconnectHelp, DEBUGGER_COMMAND_DISCONNECT_ATTRIBUTES};
    g_CommandsList[".debug"] = {&CommandDebug, &CommandDebugHelp, DEBUGGER_COMMAND_DEBUG_ATTRIBUTES};
    g_CommandsList["debug"]  = {&CommandDebug, &CommandDebugHelp, DEBUGGER_COMMAND_DEBUG_ATTRIBUTES};
    g_CommandsList[".status"] = {&CommandStatus, &CommandStatusHelp, DEBUGGER_COMMAND_DOT_STATUS_ATTRIBUTES};
    g_CommandsList["status"]  = {&CommandStatus, &CommandStatusHelp, DEBUGGER_COMMAND_STATUS_ATTRIBUTES};
    g_CommandsList["load"] = {&CommandLoad, &CommandLoadHelp, DEBUGGER_COMMAND_LOAD_ATTRIBUTES};
    g_CommandsList["exit"]  = {&CommandExit, &CommandExitHelp, DEBUGGER_COMMAND_EXIT_ATTRIBUTES};
    g_CommandsList[".exit"] = {&CommandExit, &CommandExitHelp, DEBUGGER_COMMAND_EXIT_ATTRIBUTES};
    g_CommandsList["flush"] = {&CommandFlush, &CommandFlushHelp, DEBUGGER_COMMAND_FLUSH_ATTRIBUTES};
    g_CommandsList["pause"] = {&CommandPause, &CommandPauseHelp, DEBUGGER_COMMAND_PAUSE_ATTRIBUTES};
    g_CommandsList["unload"] = {&CommandUnload, &CommandUnloadHelp, DEBUGGER_COMMAND_UNLOAD_ATTRIBUTES};
    g_CommandsList[".script"] = {&CommandScript, &CommandScriptHelp, DEBUGGER_COMMAND_SCRIPT_ATTRIBUTES};
    g_CommandsList["script"]  = {&CommandScript, &CommandScriptHelp, DEBUGGER_COMMAND_SCRIPT_ATTRIBUTES};
    g_CommandsList["output"] = {&CommandOutput, &CommandOutputHelp, DEBUGGER_COMMAND_OUTPUT_ATTRIBUTES};
    g_CommandsList["print"] = {&CommandPrint, &CommandPrintHelp, DEBUGGER_COMMAND_PRINT_ATTRIBUTES};
    g_CommandsList["?"]        = {&CommandEval, &CommandEvalHelp, DEBUGGER_COMMAND_EVAL_ATTRIBUTES};
    g_CommandsList["eval"]     = {&CommandEval, &CommandEvalHelp, DEBUGGER_COMMAND_EVAL_ATTRIBUTES};
    g_CommandsList["evaluate"] = {&CommandEval, &CommandEvalHelp, DEBUGGER_COMMAND_EVAL_ATTRIBUTES};
    g_CommandsList[".logopen"] = {&CommandLogopen, &CommandLogopenHelp, DEBUGGER_COMMAND_LOGOPEN_ATTRIBUTES};
    g_CommandsList[".logclose"] = {&CommandLogclose, &CommandLogcloseHelp, DEBUGGER_COMMAND_LOGCLOSE_ATTRIBUTES};
    g_CommandsList["test"] = {&CommandTest, &CommandTestHelp, DEBUGGER_COMMAND_TEST_ATTRIBUTES};
    g_CommandsList["cpu"] = {&CommandCpu, &CommandCpuHelp, DEBUGGER_COMMAND_CPU_ATTRIBUTES};
    g_CommandsList["wrmsr"] = {&CommandWrmsr, &CommandWrmsrHelp, DEBUGGER_COMMAND_WRMSR_ATTRIBUTES};
    g_CommandsList["rdmsr"] = {&CommandRdmsr, &CommandRdmsrHelp, DEBUGGER_COMMAND_RDMSR_ATTRIBUTES};
    g_CommandsList["!va2pa"] = {&CommandVa2pa, &CommandVa2paHelp, DEBUGGER_COMMAND_VA2PA_ATTRIBUTES};
    g_CommandsList["!pa2va"] = {&CommandPa2va, &CommandPa2vaHelp, DEBUGGER_COMMAND_PA2VA_ATTRIBUTES};
    g_CommandsList[".formats"] = {&CommandFormats, &CommandFormatsHelp, DEBUGGER_COMMAND_FORMATS_ATTRIBUTES};
    g_CommandsList[".format"]  = {&CommandFormats, &CommandFormatsHelp, DEBUGGER_COMMAND_FORMATS_ATTRIBUTES};
    g_CommandsList["!pte"] = {&CommandPte, &CommandPteHelp, DEBUGGER_COMMAND_PTE_ATTRIBUTES};
    g_CommandsList["~"]    = {&CommandCore, &CommandCoreHelp, DEBUGGER_COMMAND_CORE_ATTRIBUTES};
    g_CommandsList["core"] = {&CommandCore, &CommandCoreHelp, DEBUGGER_COMMAND_CORE_ATTRIBUTES};
    g_CommandsList["!monitor"] = {&CommandMonitor, &CommandMonitorHelp, DEBUGGER_COMMAND_MONITOR_ATTRIBUTES};
    g_CommandsList["!vmcall"] = {&CommandVmcall, &CommandVmcallHelp, DEBUGGER_COMMAND_VMCALL_ATTRIBUTES};
    g_CommandsList["!epthook"] = {&CommandEptHook, &CommandEptHookHelp, DEBUGGER_COMMAND_EPTHOOK_ATTRIBUTES};
    g_CommandsList["bh"]       = {&CommandEptHook, &CommandEptHookHelp, DEBUGGER_COMMAND_EPTHOOK_ATTRIBUTES};
    g_CommandsList["bp"] = {&CommandBp, &CommandBpHelp, DEBUGGER_COMMAND_BP_ATTRIBUTES};
    g_CommandsList["bl"] = {&CommandBl, &CommandBlHelp, DEBUGGER_COMMAND_BD_ATTRIBUTES};
    g_CommandsList["be"] = {&CommandBe, &CommandBeHelp, DEBUGGER_COMMAND_BD_ATTRIBUTES};
    g_CommandsList["bd"] = {&CommandBd, &CommandBdHelp, DEBUGGER_COMMAND_BD_ATTRIBUTES};
    g_CommandsList["bc"] = {&CommandBc, &CommandBcHelp, DEBUGGER_COMMAND_BD_ATTRIBUTES};
    g_CommandsList["!epthook2"] = {&CommandEptHook2, &CommandEptHook2Help, DEBUGGER_COMMAND_EPTHOOK2_ATTRIBUTES};
    g_CommandsList["!cpuid"] = {&CommandCpuid, &CommandCpuidHelp, DEBUGGER_COMMAND_CPUID_ATTRIBUTES};
    g_CommandsList["!msrread"] = {&CommandMsrread, &CommandMsrreadHelp, DEBUGGER_COMMAND_MSRREAD_ATTRIBUTES};
    g_CommandsList["!msread"]  = {&CommandMsrread, &CommandMsrreadHelp, DEBUGGER_COMMAND_MSRREAD_ATTRIBUTES};
    g_CommandsList["!msrwrite"] = {&CommandMsrwrite, &CommandMsrwriteHelp, DEBUGGER_COMMAND_MSRWRITE_ATTRIBUTES};
    g_CommandsList["!tsc"] = {&CommandTsc, &CommandTscHelp, DEBUGGER_COMMAND_TSC_ATTRIBUTES};
    g_CommandsList["!pmc"] = {&CommandPmc, &CommandPmcHelp, DEBUGGER_COMMAND_PMC_ATTRIBUTES};
    g_CommandsList["!crwrite"] = {&CommandCrwrite, &CommandCrwriteHelp, DEBUGGER_COMMAND_CRWRITE_ATTRIBUTES};
    g_CommandsList["!dr"] = {&CommandDr, &CommandDrHelp, DEBUGGER_COMMAND_DR_ATTRIBUTES};
    g_CommandsList["!ioin"] = {&CommandIoin, &CommandIoinHelp, DEBUGGER_COMMAND_IOIN_ATTRIBUTES};
    g_CommandsList["!ioout"] = {&CommandIoout, &CommandIooutHelp, DEBUGGER_COMMAND_IOOUT_ATTRIBUTES};
    g_CommandsList["!exception"] = {&CommandException, &CommandExceptionHelp, DEBUGGER_COMMAND_EXCEPTION_ATTRIBUTES};
    g_CommandsList["!interrupt"] = {&CommandInterrupt, &CommandInterruptHelp, DEBUGGER_COMMAND_INTERRUPT_ATTRIBUTES};
    g_CommandsList["!syscall"]  = {&CommandSyscallAndSysret, &CommandSyscallHelp, DEBUGGER_COMMAND_SYSCALL_ATTRIBUTES};
    g_CommandsList["!syscall2"] = {&CommandSyscallAndSysret, &CommandSyscallHelp, DEBUGGER_COMMAND_SYSCALL_ATTRIBUTES};
    g_CommandsList["!sysret"]  = {&CommandSyscallAndSysret, &CommandSysretHelp, DEBUGGER_COMMAND_SYSRET_ATTRIBUTES};
    g_CommandsList["!sysret2"] = {&CommandSyscallAndSysret, &CommandSysretHelp, DEBUGGER_COMMAND_SYSRET_ATTRIBUTES};
    g_CommandsList["!hide"] = {&CommandHide, &CommandHideHelp, DEBUGGER_COMMAND_HIDE_ATTRIBUTES};
    g_CommandsList["!unhide"] = {&CommandUnhide, &CommandUnhideHelp, DEBUGGER_COMMAND_UNHIDE_ATTRIBUTES};
    g_CommandsList["!measure"] = {&CommandMeasure, &CommandMeasureHelp, DEBUGGER_COMMAND_MEASURE_ATTRIBUTES};
    g_CommandsList["lm"] = {&CommandLm, &CommandLmHelp, DEBUGGER_COMMAND_LM_ATTRIBUTES};
    g_CommandsList["p"]  = {&CommandP, &CommandPHelp, DEBUGGER_COMMAND_P_ATTRIBUTES};
    g_CommandsList["pr"] = {&CommandP, &CommandPHelp, DEBUGGER_COMMAND_P_ATTRIBUTES};
    g_CommandsList["t"]  = {&CommandT, &CommandTHelp, DEBUGGER_COMMAND_T_ATTRIBUTES};
    g_CommandsList["tr"] = {&CommandT, &CommandTHelp, DEBUGGER_COMMAND_T_ATTRIBUTES};
    g_CommandsList["i"]  = {&CommandI, &CommandIHelp, DEBUGGER_COMMAND_I_ATTRIBUTES};
    g_CommandsList["ir"] = {&CommandI, &CommandIHelp, DEBUGGER_COMMAND_I_ATTRIBUTES};
    g_CommandsList["db"]  = {&CommandReadMemoryAndDisassembler,
                             &CommandReadMemoryAndDisassemblerHelp,
                             DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
    g_CommandsList["dc"]  = {&CommandReadMemoryAndDisassembler,
                             &CommandReadMemoryAndDisassemblerHelp,
                             DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
    g_CommandsList["dd"]  = {&CommandReadMemoryAndDisassembler,
                             &CommandReadMemoryAndDisassemblerHelp,
                             DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
    g_CommandsList["dq"]  = {&CommandReadMemoryAndDisassembler,
                             &CommandReadMemoryAndDisassemblerHelp,
                             DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
    g_CommandsList["!db"] = {&CommandReadMemoryAndDisassembler,
                             &CommandReadMemoryAndDisassemblerHelp,
                             DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
    g_CommandsList["!dc"] = {&CommandReadMemoryAndDisassembler,
                             &CommandReadMemoryAndDisassemblerHelp,
                             DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
    g_CommandsList["!dd"] = {&CommandReadMemoryAndDisassembler,
                             &CommandReadMemoryAndDisassemblerHelp,
                             DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
    g_CommandsList["!dq"] = {&CommandReadMemoryAndDisassembler,
                             &CommandReadMemoryAndDisassemblerHelp,
                             DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
    g_CommandsList["!u"]  = {&CommandReadMemoryAndDisassembler,
                             &CommandReadMemoryAndDisassemblerHelp,
                             DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
    g_CommandsList["u"]   = {&CommandReadMemoryAndDisassembler,
                             &CommandReadMemoryAndDisassemblerHelp,
                             DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
    g_CommandsList["!u2"] = {&CommandReadMemoryAndDisassembler,
                             &CommandReadMemoryAndDisassemblerHelp,
                             DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
    g_CommandsList["u2"]  = {&CommandReadMemoryAndDisassembler,
                             &CommandReadMemoryAndDisassemblerHelp,
                             DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
    g_CommandsList["eb"]  = {&CommandEditMemory, &CommandEditMemoryHelp, DEBUGGER_COMMAND_E_ATTRIBUTES};
    g_CommandsList["ed"]  = {&CommandEditMemory, &CommandEditMemoryHelp, DEBUGGER_COMMAND_E_ATTRIBUTES};
    g_CommandsList["eq"]  = {&CommandEditMemory, &CommandEditMemoryHelp, DEBUGGER_COMMAND_E_ATTRIBUTES};
    g_CommandsList["!eb"] = {&CommandEditMemory, &CommandEditMemoryHelp, DEBUGGER_COMMAND_E_ATTRIBUTES};
    g_CommandsList["!ed"] = {&CommandEditMemory, &CommandEditMemoryHelp, DEBUGGER_COMMAND_E_ATTRIBUTES};
    g_CommandsList["!eq"] = {&CommandEditMemory, &CommandEditMemoryHelp, DEBUGGER_COMMAND_E_ATTRIBUTES};
    g_CommandsList["sb"]  = {&CommandSearchMemory, &CommandSearchMemoryHelp, DEBUGGER_COMMAND_S_ATTRIBUTES};
    g_CommandsList["sd"]  = {&CommandSearchMemory, &CommandSearchMemoryHelp, DEBUGGER_COMMAND_S_ATTRIBUTES};
    g_CommandsList["sq"]  = {&CommandSearchMemory, &CommandSearchMemoryHelp, DEBUGGER_COMMAND_S_ATTRIBUTES};
    g_CommandsList["!sb"] = {&CommandSearchMemory, &CommandSearchMemoryHelp, DEBUGGER_COMMAND_S_ATTRIBUTES};
    g_CommandsList["!sd"] = {&CommandSearchMemory, &CommandSearchMemoryHelp, DEBUGGER_COMMAND_S_ATTRIBUTES};
    g_CommandsList["!sq"] = {&CommandSearchMemory, &CommandSearchMemoryHelp, DEBUGGER_COMMAND_S_ATTRIBUTES};
    g_CommandsList["r"] = {&CommandR, &CommandRHelp, DEBUGGER_COMMAND_R_ATTRIBUTES};
    g_CommandsList[".sympath"] = {&CommandSympath, &CommandSympathHelp, DEBUGGER_COMMAND_SYMPATH_ATTRIBUTES};
    g_CommandsList["sympath"]  = {&CommandSympath, &CommandSympathHelp, DEBUGGER_COMMAND_SYMPATH_ATTRIBUTES};
    g_CommandsList[".sym"] = {&CommandSym, &CommandSymHelp, DEBUGGER_COMMAND_SYM_ATTRIBUTES};
    g_CommandsList["sym"]  = {&CommandSym, &CommandSymHelp, DEBUGGER_COMMAND_SYM_ATTRIBUTES};
    g_CommandsList["x"] = {&CommandX, &CommandXHelp, DEBUGGER_COMMAND_X_ATTRIBUTES};
    g_CommandsList["prealloc"]    = {&CommandPrealloc, &CommandPreallocHelp, DEBUGGER_COMMAND_PREALLOC_ATTRIBUTES};
    g_CommandsList["preallocate"] = {&CommandPrealloc, &CommandPreallocHelp, DEBUGGER_COMMAND_PREALLOC_ATTRIBUTES};
    g_CommandsList["k"]  = {&CommandK, &CommandKHelp, DEBUGGER_COMMAND_K_ATTRIBUTES};
    g_CommandsList["kd"] = {&CommandK, &CommandKHelp, DEBUGGER_COMMAND_K_ATTRIBUTES};
    g_CommandsList["kq"] = {&CommandK, &CommandKHelp, DEBUGGER_COMMAND_K_ATTRIBUTES};
    g_CommandsList["dt"]  = {&CommandDtAndStruct, &CommandDtHelp, DEBUGGER_COMMAND_DT_ATTRIBUTES};
    g_CommandsList["!dt"] = {&CommandDtAndStruct, &CommandDtHelp, DEBUGGER_COMMAND_DT_ATTRIBUTES};
    g_CommandsList["struct"]    = {&CommandDtAndStruct, &CommandStructHelp, DEBUGGER_COMMAND_STRUCT_ATTRIBUTES};
    g_CommandsList["structure"] = {&CommandDtAndStruct, &CommandStructHelp, DEBUGGER_COMMAND_STRUCT_ATTRIBUTES};
    g_CommandsList[".pe"] = {&CommandPe, &CommandPeHelp, DEBUGGER_COMMAND_PE_ATTRIBUTES};
}*/
return true
}



