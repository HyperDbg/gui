package debugging-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\exit.cpp.back

type (
	Exit interface {
		CommandExitHelp() (ok bool) //col:31
		CommandExit() (ok bool)     //col:59
	}
)

func NewExit() { return &exit{} }

func (e *exit) CommandExitHelp() (ok bool) { //col:31
	/*CommandExitHelp()
	  {
	      ShowMessages(
	          "exit : unloads and uninstalls the drivers and closes the debugger.\n\n");
	      ShowMessages("syntax : \texit\n");
	  }*/
	return true
}

func (e *exit) CommandExit() (ok bool) { //col:59
	/*CommandExit(vector<string> SplittedCommand, string Command)
	  {
	      if (SplittedCommand.size() != 1)
	      {
	          ShowMessages("incorrect use of 'exit'\n\n");
	          CommandExitHelp();
	          return;
	      }
	      if (g_DeviceHandle)
	      {
	          HyperDbgUnloadVmm();
	      }
	      exit(0);
	  }*/
	return true
}
