package meta-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\meta-commands\logclose.cpp.back

type (
	Logclose interface {
		CommandLogcloseHelp() (ok bool) //col:5
		CommandLogclose() (ok bool)     //col:31
	}
	logclose struct{}
)

func NewLogclose() Logclose { return &logclose{} }

func (l *logclose) CommandLogcloseHelp() (ok bool) { //col:5
	/*CommandLogcloseHelp()
	  {
	      ShowMessages(".logclose : closes the previously opened log.\n\n");
	      ShowMessages("syntax : \t.logclose\n");
	  }*/
	return true
}

func (l *logclose) CommandLogclose() (ok bool) { //col:31
	/*CommandLogclose(vector<string> SplittedCommand, string Command)
	  {
	      if (SplittedCommand.size() != 1)
	      {
	          ShowMessages("incorrect use of '.logclose'\n\n");
	          CommandLogcloseHelp();
	          return;
	      }
	      if (!g_LogOpened)
	      {
	          ShowMessages("there is no opened log, did you use '.logopen'? \n");
	          return;
	      }
	      time_t    t  = time(NULL);
	      struct tm tm = *localtime(&t);
	      ShowMessages("log file closed (%d-%02d-%02d "
	                   "%02d:%02d:%02d)\n",
	                   tm.tm_year + 1900,
	                   tm.tm_mon + 1,
	                   tm.tm_mday,
	                   tm.tm_hour,
	                   tm.tm_min,
	                   tm.tm_sec);
	      g_LogOpenFile.close();
	      g_LogOpened = FALSE;
	  }*/
	return true
}
