package meta-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\meta-commands\cls.cpp.back

type (
	Cls interface {
		CommandClearScreenHelp() (ok bool) //col:5
		CommandClearScreen() (ok bool)     //col:9
	}
	cls struct{}
)

func NewCls() Cls { return &cls{} }

func (c *cls) CommandClearScreenHelp() (ok bool) { //col:5
	/*CommandClearScreenHelp()
	  {
	      ShowMessages(".cls : clears the screen.\n\n");
	      ShowMessages("syntax : \t.cls\n");
	  }*/
	return true
}

func (c *cls) CommandClearScreen() (ok bool) { //col:9
	/*CommandClearScreen(vector<string> SplittedCommand, string Command)
	  {
	      system("cls");
	  }*/
	return true
}
