package debugging-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\prealloc.cpp.back

type (
	Prealloc interface {
		CommandPreallocHelp() (ok bool) //col:8
		CommandPrealloc() (ok bool)     //col:66
	}
	prealloc struct{}
)

func NewPrealloc() Prealloc { return &prealloc{} }

func (p *prealloc) CommandPreallocHelp() (ok bool) { //col:8
	/*CommandPreallocHelp()
	  {
	      ShowMessages("prealloc : pre-allocates buffer for special purposes.\n\n");
	      ShowMessages("syntax : \tprealloc  [Type (string)] [Count (hex)]\n");
	      ShowMessages("\n");
	      ShowMessages("\t\te.g : prealloc monitor 10\n");
	      ShowMessages("\t\te.g : prealloc thread-interception 8\n");
	  }*/
	return true
}

func (p *prealloc) CommandPrealloc() (ok bool) { //col:66
	/*CommandPrealloc(vector<string> SplittedCommand, string Command)
	  {
	      BOOL                      Status;
	      ULONG                     ReturnedLength;
	      UINT64                    Count;
	      DEBUGGER_PREALLOC_COMMAND PreallocRequest = {0};
	      if (SplittedCommand.size() != 3)
	      {
	          ShowMessages("incorrect use of 'prealloc'\n\n");
	          CommandPreallocHelp();
	          return;
	      }
	      if (!SplittedCommand.at(1).compare("monitor"))
	      {
	          PreallocRequest.Type = DEBUGGER_PREALLOC_COMMAND_TYPE_MONITOR;
	      }
	      else if (!SplittedCommand.at(1).compare("thread-interception"))
	      {
	          PreallocRequest.Type = DEBUGGER_PREALLOC_COMMAND_TYPE_THREAD_INTERCEPTION;
	      }
	      else
	      {
	          ShowMessages("err, couldn't resolve error at '%s'\n",
	                       SplittedCommand.at(1).c_str());
	          return;
	      }
	      if (!SymbolConvertNameOrExprToAddress(SplittedCommand.at(2), &Count))
	      {
	          ShowMessages("err, couldn't resolve error at '%s'\n",
	                       SplittedCommand.at(2).c_str());
	          return;
	      }
	      PreallocRequest.Count = Count;
	      AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
	      Status = DeviceIoControl(
	          g_DeviceHandle,
	          IOCTL_RESERVE_PRE_ALLOCATED_POOLS,
	          &PreallocRequest,
	          SIZEOF_DEBUGGER_PREALLOC_COMMAND,
	          &PreallocRequest,
	          SIZEOF_DEBUGGER_PREALLOC_COMMAND,
	          &ReturnedLength,
	          NULL
	      );
	      if (!Status)
	      {
	          ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
	          return;
	      }
	      if (PreallocRequest.KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFULL)
	      {
	          ShowMessages("the requested pools are allocated and reserved\n");
	      }
	      else
	      {
	          ShowErrorMessage(PreallocRequest.KernelStatus);
	      }
	  }*/
	return true
}
