package meta-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\meta-commands\status.cpp.back

type (
	Status interface {
		CommandStatusHelp() (ok bool) //col:9
		CommandStatus() (ok bool)     //col:46
	}
	status struct{}
)

func NewStatus() Status { return &status{} }

func (s *status) CommandStatusHelp() (ok bool) { //col:9
	/*CommandStatusHelp()
	  {
	      ShowMessages(".status | status : gets the status of current debugger in local "
	                   "system (if you connected to a remote system then '.status' "
	                   "shows the state of current debugger, while 'status' shows the "
	                   "state of remote debuggee).\n\n");
	      ShowMessages("syntax : \t.status\n");
	      ShowMessages("syntax : \tstatus\n");
	  }*/
	return true
}

func (s *status) CommandStatus() (ok bool) { //col:46
	/*CommandStatus(vector<string> SplittedCommand, string Command)
	  {
	      if (SplittedCommand.size() != 1)
	      {
	          ShowMessages("incorrect use of '.status'\n\n");
	          CommandStatusHelp();
	      }
	      if (g_IsSerialConnectedToRemoteDebuggee)
	      {
	          ShowMessages("remote debugging - debugger ('debugger mode')\n");
	      }
	      else if (g_IsSerialConnectedToRemoteDebugger)
	      {
	          ShowMessages("remote debugging - debuggee ('debugger mode')\n");
	      }
	      else if (g_IsConnectedToRemoteDebuggee)
	      {
	          ShowMessages("remote debugging ('vmi mode'), ip : %s:%s \n",
	                       g_ServerIp.c_str(),
	                       g_ServerPort.c_str());
	      }
	      else if (g_IsConnectedToHyperDbgLocally)
	      {
	          ShowMessages("local debugging ('vmi mode')\n");
	      }
	      else if (g_IsConnectedToRemoteDebugger)
	      {
	          ShowMessages("a remote debugger connected to this system in ('vmi "
	                       "mode'), ip : %s:%s \n",
	                       g_ServerIp.c_str(),
	                       g_ServerPort.c_str());
	      }
	      else
	      {
	          ShowMessages("err, you're not connected to any instance of HyperDbg\n");
	      }
	  }*/
	return true
}
