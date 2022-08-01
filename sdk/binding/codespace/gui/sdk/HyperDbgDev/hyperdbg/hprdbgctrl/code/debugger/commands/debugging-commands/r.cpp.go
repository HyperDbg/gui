package debugging-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\r.cpp.back

type (
	R interface {
		CommandRHelp() (ok bool)     //col:12
		ShowAllRegisters() (ok bool) //col:18
		CommandR() (ok bool)         //col:128
	}
	r struct{}
)

func NewR() R { return &r{} }

func (r *r) CommandRHelp() (ok bool) { //col:12
	/*CommandRHelp()
	  {
	      ShowMessages("r : reads or modifies registers.\n\n");
	      ShowMessages("syntax : \tr\n");
	      ShowMessages("syntax : \tr [Register (string)] [= Expr (string)]\n");
	      ShowMessages("\n");
	      ShowMessages("\t\te.g : r\n");
	      ShowMessages("\t\te.g : r rax\n");
	      ShowMessages("\t\te.g : r rax = 0x55\n");
	  }*/
	return true
}

func (r *r) ShowAllRegisters() (ok bool) { //col:18
	/*ShowAllRegisters()
	  {
	      DEBUGGEE_REGISTER_READ_DESCRIPTION RegState = {0};
	      RegState.RegisterID                         = DEBUGGEE_SHOW_ALL_REGISTERS;
	      KdSendReadRegisterPacketToDebuggee(&RegState);
	  }*/
	return true
}

func (r *r) CommandR() (ok bool) { //col:128
	/*CommandR(std::vector<std::string> SplittedCommand, std::string Command)
	  {
	      PVOID                    CodeBuffer;
	      UINT64                   BufferAddress;
	      UINT32                   BufferLength;
	      UINT32                   Pointer;
	      REGS_ENUM                RegKind;
	      std::vector<std::string> Tmp;
	      std::string SetRegValue = "SetRegValue";
	      if (SplittedCommand[0] != "r")
	      {
	          return;
	      }
	      if (SplittedCommand.size() == 1)
	      {
	          if (g_IsSerialConnectedToRemoteDebuggee)
	          {
	              ShowAllRegisters();
	          }
	          else
	          {
	              ShowMessages("err, reading registers (r) is not valid in the current "
	                           "context, you should connect to a debuggee\n");
	          }
	          return;
	      }
	      if (Command.find('=', 0) == string::npos)
	      {
	          Command.erase(0, 1);
	          ReplaceAll(Command, " ", "");
	          if (RegistersMap.find(Command) != RegistersMap.end())
	          {
	              RegKind = RegistersMap[Command];
	          }
	          else
	          {
	              RegKind = (REGS_ENUM)-1;
	          }
	          if (RegKind != -1)
	          {
	              DEBUGGEE_REGISTER_READ_DESCRIPTION RegState = {0};
	              RegState.RegisterID                         = RegKind;
	              if (g_IsSerialConnectedToRemoteDebuggee)
	              {
	                  KdSendReadRegisterPacketToDebuggee(&RegState);
	              }
	              else
	              {
	                  ShowMessages("err, reading registers (r) is not valid in the current "
	                               "context, you should connect to a debuggee\n");
	              }
	          }
	          else
	          {
	              ShowMessages("err, invalid register\n");
	          }
	      }
	      else if (Command.find('=', 0) != string::npos)
	      {
	          Command.erase(0, 1);
	          Tmp = Split(Command, '=');
	          if (Tmp.size() == 2)
	          {
	              ReplaceAll(Tmp[0], " ", "");
	              string tmp = Tmp[0];
	              if (RegistersMap.find(Tmp[0]) != RegistersMap.end())
	              {
	                  RegKind = RegistersMap[Tmp[0]];
	              }
	              else
	              {
	                  if (RegistersMap.find(tmp) != RegistersMap.end())
	                  {
	                      RegKind = RegistersMap[tmp];
	                  }
	                  else
	                  {
	                      RegKind = (REGS_ENUM)-1;
	                  }
	              }
	              if (RegKind != -1)
	              {
	                  if (g_IsSerialConnectedToRemoteDebuggee)
	                  {
	                      CodeBuffer = ScriptEngineParseWrapper((char *)SetRegValue.c_str(), TRUE);
	                      if (CodeBuffer == NULL)
	                      {
	                          return;
	                      }
	                      BufferAddress = ScriptEngineWrapperGetHead(CodeBuffer);
	                      BufferLength  = ScriptEngineWrapperGetSize(CodeBuffer);
	                      Pointer       = ScriptEngineWrapperGetPointer(CodeBuffer);
	                      KdSendScriptPacketToDebuggee(BufferAddress, BufferLength, Pointer, FALSE);
	                      ScriptEngineWrapperRemoveSymbolBuffer(CodeBuffer);
	                  }
	                  else
	                  {
	                      ShowMessages("err, you're not connected to any debuggee\n");
	                  }
	              }
	              else
	              {
	                  ShowMessages("err, invalid register\n");
	              }
	          }
	      }
	  }*/
	return true
}
