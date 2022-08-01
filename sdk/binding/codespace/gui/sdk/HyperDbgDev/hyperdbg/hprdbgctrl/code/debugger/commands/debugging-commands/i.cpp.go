package debugging-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\i.cpp.back

type (
	I interface {
		CommandIHelp() (ok bool) //col:17
		CommandI() (ok bool)     //col:73
	}
	i struct{}
)

func NewI() I { return &i{} }

func (i *i) CommandIHelp() (ok bool) { //col:17
	/*CommandIHelp()
	  {
	      ShowMessages(
	          "i : executes a single instruction (step-in) and guarantees that no "
	          "other instruction is executed other than the displayed instruction "
	          "including user to the kernel (syscalls) and kernel to the user "
	          "(sysrets) and exceptions and page-faults and optionally displays all "
	          "registers and flags' resulting values.\n\n");
	      ShowMessages("syntax : \ti\n");
	      ShowMessages("syntax : \ti [Count (hex)]\n");
	      ShowMessages("syntax : \tir\n");
	      ShowMessages("syntax : \tir [Count (hex)]\n");
	      ShowMessages("\n");
	      ShowMessages("\t\te.g : i\n");
	      ShowMessages("\t\te.g : ir\n");
	      ShowMessages("\t\te.g : ir 1f\n");
	  }*/
	return true
}

func (i *i) CommandI() (ok bool) { //col:73
	/*CommandI(vector<string> SplittedCommand, string Command)
	  {
	      UINT32                           StepCount;
	      DEBUGGER_REMOTE_STEPPING_REQUEST RequestFormat;
	      if (SplittedCommand.size() != 1 && SplittedCommand.size() != 2)
	      {
	          ShowMessages("incorrect use of 'i'\n\n");
	          CommandIHelp();
	          return;
	      }
	      if (g_ActiveProcessDebuggingState.IsActive)
	      {
	          ShowMessages("the instrumentation step-in is only supported in Debugger Mode\n");
	          return;
	      }
	      RequestFormat = DEBUGGER_REMOTE_STEPPING_REQUEST_INSTRUMENTATION_STEP_IN;
	      if (SplittedCommand.size() == 2)
	      {
	          if (!ConvertStringToUInt32(SplittedCommand.at(1), &StepCount))
	          {
	              ShowMessages("please specify a correct hex value for [count]\n\n");
	              CommandIHelp();
	              return;
	          }
	      }
	      else
	      {
	          StepCount = 1;
	      }
	      if (g_IsSerialConnectedToRemoteDebuggee)
	      {
	          g_IsInstrumentingInstructions = TRUE;
	          for (size_t i = 0; i < StepCount; i++)
	          {
	              KdSendStepPacketToDebuggee(RequestFormat);
	              if (!SplittedCommand.at(0).compare("ir"))
	              {
	                  ShowAllRegisters();
	                  if (i != StepCount - 1)
	                  {
	                      ShowMessages("\n");
	                  }
	              }
	              if (!g_IsInstrumentingInstructions)
	              {
	                  break;
	              }
	          }
	          g_IsInstrumentingInstructions = FALSE;
	      }
	      else
	      {
	          ShowMessages("err, stepping (i) is not valid in the current context, you "
	                       "should connect to a debuggee\n");
	      }
	  }*/
	return true
}
