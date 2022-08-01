package debugging-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\x.cpp.back

type (
X interface{
CommandXHelp()(ok bool)//col:9
CommandX()(ok bool)//col:22
}
)

func NewX() { return & x{} }

func (x *x)CommandXHelp()(ok bool){//col:9
/*CommandXHelp()
{
    ShowMessages("x : searches and shows symbols (wildcard) and corresponding addresses.\n\n");
    ShowMessages("syntax : \tx [Module!Name (wildcard string)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : x nt!ExAllocatePoolWithTag \n");
    ShowMessages("\t\te.g : x nt!ExAllocatePool* \n");
    ShowMessages("\t\te.g : x nt!* \n");
}*/
return true
}

func (x *x)CommandX()(ok bool){//col:22
/*CommandX(vector<string> SplittedCommand, string Command)
{
    if (SplittedCommand.size() == 1)
    {
        ShowMessages("incorrect use of 'x'\n\n");
        CommandXHelp();
        return;
    }
    Trim(Command);
    Command.erase(0, 1);
    Trim(Command);
    ScriptEngineSearchSymbolForMaskWrapper(Command.c_str());
}*/
return true
}



