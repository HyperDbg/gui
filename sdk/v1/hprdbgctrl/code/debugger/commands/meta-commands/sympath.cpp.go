package meta_commands

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\meta-commands\sympath.cpp.back

type (
	Sympath interface {
		CommandSympathHelp() (ok bool) //col:9
	}
	sympath struct{}
)

func NewSympath() Sympath { return &sympath{} }

func (s *sympath) CommandSympathHelp() (ok bool) { //col:9
	/*
	   CommandSympathHelp()

	   	{
	   	    ShowMessages(".sympath : shows and sets the symbol server and path.\n\n");
	   	    ShowMessages("syntax : \t.sympath\n");
	   	    ShowMessages("syntax : \t.sympath [SymServer (string)]\n");
	   	    ShowMessages("\n");
	   	    ShowMessages("\t\te.g : .sympath\n");
	   	    ShowMessages("\t\te.g : .sympath SRV*c:\\Symbols*https:
	   	}
	*/
	return true
}

