package meta-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\meta-commands\help.cpp.back

type (
Help interface{
CommandHelpHelp()(ok bool)//col:7
}
help struct{}
)

func NewHelp()Help{ return & help{} }

func (h *help)CommandHelpHelp()(ok bool){//col:7
/*CommandHelpHelp()
{
    ShowMessages(".help : shows help and example(s) of a specific command.\n\n");
    ShowMessages("syntax : \t.help [Command (string)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : .help !monitor\n");
}*/
return true
}



