package extension-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\extension-commands\measure.cpp.back

type (
	Measure interface {
		CommandMeasureHelp() (ok bool) //col:9
		CommandMeasure() (ok bool)     //col:79
	}
	measure struct{}
)

func NewMeasure() Measure { return &measure{} }

func (m *measure) CommandMeasureHelp() (ok bool) { //col:9
	/*CommandMeasureHelp()
	  {
	      ShowMessages(
	          "!measure : measures the arguments needs for the '!hide' command.\n\n");
	      ShowMessages("syntax : \t!measure [default]\n");
	      ShowMessages("\n");
	      ShowMessages("\t\te.g : !measure\n");
	      ShowMessages("\t\te.g : !measure default\n");
	  }*/
	return true
}

func (m *measure) CommandMeasure() (ok bool) { //col:79
	/*CommandMeasure(vector<string> SplittedCommand, string Command)
	  {
	      BOOLEAN DefaultMode = FALSE;
	      if (SplittedCommand.size() >= 3)
	      {
	          ShowMessages("incorrect use of '!measure'\n\n");
	          CommandMeasureHelp();
	          return;
	      }
	      if (SplittedCommand.size() == 2 && SplittedCommand.at(1).compare("default"))
	      {
	          ShowMessages("incorrect use of '!measure'\n\n");
	          CommandMeasureHelp();
	          return;
	      }
	      else if (SplittedCommand.size() == 2 &&
	               !SplittedCommand.at(1).compare("default"))
	      {
	          DefaultMode = TRUE;
	      }
	      if (g_DeviceHandle && !DefaultMode)
	      {
	          ShowMessages(
	              "Debugger is loaded and your machine is already in a hypervisor, you "
	              "should measure the times before 'load'-ing the debugger, please "
	              "'unload' the debugger and use '!measure' again or use '!measure "
	              "default' to use hardcoded measurements\n");
	          return;
	      }
	      if (!DefaultMode)
	      {
	          if (TransparentModeCheckHypervisorPresence(
	                  &g_CpuidAverage,
	                  &g_CpuidStandardDeviation,
	                  &g_CpuidMedian))
	          {
	              ShowMessages(
	                  "we detected that there is a hypervisor, on your system, it "
	                  "leads to wrong measurement results for our transparent-mode, please "
	                  "make sure that you're not in a hypervisor then measure the result "
	                  "again; otherwise the transparent-mode will not work but you can use "
	                  "'!measure default' to use the hardcoded measurements !\n\n");
	              return;
	          }
	          if (TransparentModeCheckRdtscpVmexit(
	                  &g_RdtscAverage,
	                  &g_RdtscStandardDeviation,
	                  &g_RdtscMedian))
	          {
	              ShowMessages(
	                  "we detected that there is a hypervisor, on your system, it "
	                  "leads to wrong measurement results for our transparent-mode, please "
	                  "make sure that you're not in a hypervisor then measure the result "
	                  "again; otherwise the transparent-mode will not work but you can use "
	                  "'!measure default' to use the hardcoded measurements !\n\n");
	              return;
	          }
	      }
	      else
	      {
	          g_CpuidAverage           = 0x5f;
	          g_CpuidStandardDeviation = 0x10;
	          g_CpuidMedian            = 0x5f;
	          g_RdtscAverage           = 0x16;
	          g_RdtscStandardDeviation = 0x5;
	          g_RdtscMedian            = 0x16;
	      }
	      ShowMessages("the measurements were successful\nyou can use the '!hide' command now\n");
	      g_TransparentResultsMeasured = TRUE;
	  }*/
	return true
}
