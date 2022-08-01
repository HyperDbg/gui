package meta-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\meta-commands\detach.cpp.back

type (
Detach interface{
CommandDetachHelp()(ok bool)//col:5
DetachFromProcess()(ok bool)//col:18
CommandDetach()(ok bool)//col:35
}
)

func NewDetach() { return & detach{} }

func (d *detach)CommandDetachHelp()(ok bool){//col:5
/*CommandDetachHelp()
{
    ShowMessages(".detach : detaches from debugging a user-mode process.\n\n");
    ShowMessages("syntax : \t.detach \n");
}*/
return true
}

func (d *detach)DetachFromProcess()(ok bool){//col:18
/*DetachFromProcess()
{
    BOOLEAN                                  Status;
    ULONG                                    ReturnedLength;
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS DetachRequest = {0};
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
    if (!g_ActiveProcessDebuggingState.IsActive)
    {
        ShowMessages("you're not attached to any thread\n");
        return;
    }
    UdDetachProcess(g_ActiveProcessDebuggingState.ProcessId, g_ActiveProcessDebuggingState.ProcessDebuggingToken);
}*/
return true
}

func (d *detach)CommandDetach()(ok bool){//col:35
/*CommandDetach(vector<string> SplittedCommand, string Command)
{
    if (SplittedCommand.size() >= 2)
    {
        ShowMessages("incorrect use of '.detach'\n\n");
        CommandDetachHelp();
        return;
    }
    if (g_IsSerialConnectedToRemoteDebuggee)
    {
        ShowMessages("err, '.attach', and '.detach' commands are only usable "
                     "in VMI Mode, you can use the '.process', or the '.thread' "
                     "in Debugger Mode\n");
        return;
    }
    DetachFromProcess();
}*/
return true
}



