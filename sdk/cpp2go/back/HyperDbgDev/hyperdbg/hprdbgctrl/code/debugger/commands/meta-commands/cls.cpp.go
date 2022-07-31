package meta-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\meta-commands\cls.cpp.back

type (
Cls interface{
CommandClearScreenHelp()(ok bool)//col:25
CommandClearScreen()(ok bool)//col:38
}
)

func NewCls() { return & cls{} }

func (c *cls)CommandClearScreenHelp()(ok bool){//col:25
/*CommandClearScreenHelp()
{
    ShowMessages(".cls : clears the screen.\n\n");
    ShowMessages("syntax : \t.cls\n");
}*/
return true
}

func (c *cls)CommandClearScreen()(ok bool){//col:38
/*CommandClearScreen(vector<string> SplittedCommand, string Command)
{
    system("cls");
}*/
return true
}



