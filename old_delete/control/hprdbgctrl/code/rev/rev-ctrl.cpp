#include "pch.h"
BOOLEAN RevRequestService(REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST * RevRequest){
    BOOLEAN Status;
    ULONG   ReturnedLength;
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
    Status = DeviceIoControl(
        g_DeviceHandle,                                      
        IOCTL_REQUEST_REV_MACHINE_SERVICE,                   
        RevRequest,                                          
        SIZEOF_REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST, 
        RevRequest,                                          
        SIZEOF_REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST, 
        &ReturnedLength,                                     
        NULL                                                 
    );
    if (!Status){
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return FALSE;
    }
    if (RevRequest->KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFUL){
        ShowMessages("the reversing machine service request was successful!\n");
    }
    else{
        ShowErrorMessage(RevRequest->KernelStatus);
        return FALSE;
    }
    return FALSE;
}
