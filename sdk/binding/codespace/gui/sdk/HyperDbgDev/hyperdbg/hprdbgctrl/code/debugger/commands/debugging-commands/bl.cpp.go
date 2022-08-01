package debugging-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\bl.cpp.back

type (
	Bl interface {
		CommandBlHelp() (ok bool) //col:5
		CommandBl() (ok bool)     //col:26
	}
	bl struct{}
)

func NewBl() Bl { return &bl{} }

func (b *bl) CommandBlHelp() (ok bool) { //col:5
	/*CommandBlHelp()
	  {
	      ShowMessages("bl : lists all the enabled and disabled breakpoints.\n\n");
	      ShowMessages("syntax : \tbl\n");
	  }*/
	return true
}

func (b *bl) CommandBl() (ok bool) { //col:26
	/*CommandBl(vector<string> SplittedCommand, string Command)
	  {
	      DEBUGGEE_BP_LIST_OR_MODIFY_PACKET Request = {0};
	      if (SplittedCommand.size() != 1)
	      {
	          ShowMessages("incorrect use of 'bl'\n\n");
	          CommandBlHelp();
	          return;
	      }
	      if (g_IsSerialConnectedToRemoteDebuggee)
	      {
	          Request.Request = DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_LIST_BREAKPOINTS;
	          KdSendListOrModifyPacketToDebuggee(&Request);
	          ShowMessages("\n");
	      }
	      else
	      {
	          ShowMessages("err, listing breakpoints is only valid if you connected to "
	                       "a debuggee in debugger-mode\n");
	      }
	  }*/
	return true
}
