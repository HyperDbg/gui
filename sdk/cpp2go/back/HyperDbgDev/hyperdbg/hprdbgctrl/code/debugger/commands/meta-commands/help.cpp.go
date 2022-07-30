package meta-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\meta-commands\help.cpp.back

type (
Help interface{
CommandHelpHelp()(ok bool)//col:30
}
)

func NewHelp() { return & help{} }

func (h *help)CommandHelpHelp()(ok bool){//col:30
/*CommandHelpHelp()
{
    ShowMessages(".help : shows help and example(s) of a specific command.\n\n");
    ShowMessages("syntax : \t.help [Command (string)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : .help !monitor\n");
}*/
return true
}



