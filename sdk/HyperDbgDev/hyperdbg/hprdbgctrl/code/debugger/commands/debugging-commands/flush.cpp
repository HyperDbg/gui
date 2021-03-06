#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
VOID
CommandFlushHelp() {
    ShowMessages("flush : removes all the buffer and messages from kernel-mode "
                 "buffers.\n\n");
    ShowMessages("syntax : \tflush \n");
}

VOID
CommandFlushRequestFlush() {
    BOOL                           Status;
    ULONG                          ReturnedLength;
    DEBUGGER_FLUSH_LOGGING_BUFFERS FlushRequest = {0};
    if (g_IsSerialConnectedToRemoteDebuggee) {
        KdSendFlushPacketToDebuggee();
    } else {
        AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
        Status = DeviceIoControl(
            g_DeviceHandle,                        // Handle to device
            IOCTL_DEBUGGER_FLUSH_LOGGING_BUFFERS,  // IO Control code
            &FlushRequest,                         // Input Buffer to driver.
            SIZEOF_DEBUGGER_FLUSH_LOGGING_BUFFERS, // Input buffer length
            &FlushRequest,                         // Output Buffer from driver.
            SIZEOF_DEBUGGER_FLUSH_LOGGING_BUFFERS, // Length of output buffer in
            &ReturnedLength,                       // Bytes placed in buffer.
            NULL                                   // synchronous call
        );
        if (!Status) {
            ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
            return;
        }
        if (FlushRequest.KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
            ShowMessages(
                "flushing buffers was successful, total %d messages were cleared.\n",
                FlushRequest.CountOfMessagesThatSetAsReadFromVmxNonRoot +
                    FlushRequest.CountOfMessagesThatSetAsReadFromVmxRoot);
        } else {
            ShowMessages("flushing buffers was not successful :(\n");
        }
    }
}

VOID
CommandFlush(vector<string> SplittedCommand, string Command) {
    if (SplittedCommand.size() != 1) {
        ShowMessages("incorrect use of 'flush'\n\n");
        CommandFlushHelp();
        return;
    }
    CommandFlushRequestFlush();
}
