package debugging-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\flush.cpp.back

type (
Flush interface{
CommandFlushHelp()(ok bool)//col:6
CommandFlushRequestFlush()(ok bool)//col:46
CommandFlush()(ok bool)//col:56
}
)

func NewFlush() { return & flush{} }

func (f *flush)CommandFlushHelp()(ok bool){//col:6
/*CommandFlushHelp()
{
    ShowMessages("flush : removes all the buffer and messages from kernel-mode "
                 "buffers.\n\n");
    ShowMessages("syntax : \tflush \n");
}*/
return true
}

func (f *flush)CommandFlushRequestFlush()(ok bool){//col:46
/*CommandFlushRequestFlush()
{
    BOOL                           Status;
    ULONG                          ReturnedLength;
    DEBUGGER_FLUSH_LOGGING_BUFFERS FlushRequest = {0};
    if (g_IsSerialConnectedToRemoteDebuggee)
    {
        KdSendFlushPacketToDebuggee();
    }
    else
    {
        AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
        Status = DeviceIoControl(
            g_DeviceHandle,                        
            IOCTL_DEBUGGER_FLUSH_LOGGING_BUFFERS,  
            &FlushRequest,                         
            SIZEOF_DEBUGGER_FLUSH_LOGGING_BUFFERS, 
            &FlushRequest,                         
            SIZEOF_DEBUGGER_FLUSH_LOGGING_BUFFERS, 
            &ReturnedLength,                       
            NULL                                   
        );
        if (!Status)
        {
            ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
            return;
        }
        if (FlushRequest.KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFULL)
        {
            ShowMessages(
                "flushing buffers was successful, total %d messages were cleared.\n",
                FlushRequest.CountOfMessagesThatSetAsReadFromVmxNonRoot +
                    FlushRequest.CountOfMessagesThatSetAsReadFromVmxRoot);
        }
        else
        {
            ShowMessages("flushing buffers was not successful :(\n");
        }
    }
}*/
return true
}

func (f *flush)CommandFlush()(ok bool){//col:56
/*CommandFlush(vector<string> SplittedCommand, string Command)
{
    if (SplittedCommand.size() != 1)
    {
        ShowMessages("incorrect use of 'flush'\n\n");
        CommandFlushHelp();
        return;
    }
    CommandFlushRequestFlush();
}*/
return true
}



