package debugging-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\print.cpp.back

type (
Print interface{
CommandPrintHelp()(ok bool)//col:7
CommandPrint()(ok bool)//col:43
}
print struct{}
)

func NewPrint()Print{ return & print{} }

func (p *print)CommandPrintHelp()(ok bool){//col:7
/*CommandPrintHelp()
{
    ShowMessages("print : evaluates expressions.\n\n");
    ShowMessages("syntax : \tprint [Expression (string)]\n");
    ShowMessages("\n");
}*/
return true
}

func (p *print)CommandPrint()(ok bool){//col:43
/*CommandPrint(vector<string> SplittedCommand, string Command)
{
    PVOID  CodeBuffer;
    UINT64 BufferAddress;
    UINT32 BufferLength;
    UINT32 Pointer;
    if (SplittedCommand.size() == 1)
    {
        ShowMessages("incorrect use of 'print'\n\n");
        CommandPrintHelp();
        return;
    }
    Trim(Command);
    Command.erase(0, 5);
    Trim(Command);
    Command.insert(0, "print(");
    Command.append(");");
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
        ShowMessages("\n");
    }
    else
    {
        ShowMessages("err, you're not connected to any debuggee\n");
    }
}*/
return true
}



