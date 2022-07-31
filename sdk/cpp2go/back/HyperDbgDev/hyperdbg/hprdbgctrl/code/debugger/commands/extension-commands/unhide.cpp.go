package extension-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\extension-commands\unhide.cpp.back

type (
Unhide interface{
CommandUnhideHelp()(ok bool)//col:28
CommandUnhide()(ok bool)//col:96
}
)

func NewUnhide() { return & unhide{} }

func (u *unhide)CommandUnhideHelp()(ok bool){//col:28
/*CommandUnhideHelp()
{
    ShowMessages("!unhide : reverts the transparency measures of the '!hide' command.\n\n");
    ShowMessages("syntax : \t!unhide\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : !unhide\n");
}*/
return true
}

func (u *unhide)CommandUnhide()(ok bool){//col:96
/*CommandUnhide(vector<string> SplittedCommand, string Command)
{
    BOOLEAN                                     Status;
    ULONG                                       ReturnedLength;
    DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE UnhideRequest = {0};
    if (SplittedCommand.size() >= 2)
    {
        ShowMessages("incorrect use of '!unhide'\n\n");
        CommandUnhideHelp();
        return;
    }
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
    UnhideRequest.IsHide = FALSE;
    Status = DeviceIoControl(
    );
    if (!Status)
    {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return;
    }
    if (UnhideRequest.KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFULL)
    {
        ShowMessages("transparent debugging successfully disabled :)\n");
    }
    else if (UnhideRequest.KernelStatus ==
             DEBUGGER_ERROR_DEBUGGER_ALREADY_UHIDE)
    {
        ShowMessages("debugger is not in transparent-mode\n");
    }
    else
    {
        ShowMessages("unknown error occurred :(\n");
    }
}*/
return true
}


