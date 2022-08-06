










#include "pch.h"




extern BOOLEAN                  g_IsSerialConnectedToRemoteDebuggee;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;






VOID
CommandPa2vaHelp()
{
    ShowMessages("!pa2va : converts virtual address to physical address.\n\n");

    ShowMessages("syntax : \t!pa2va [PhysicalAddress (hex)] [pid ProcessId (hex)]\n");

    ShowMessages("\n");
    ShowMessages("\t\te.g : !pa2va nt!ExAllocatePoolWithTag\n");
    ShowMessages("\t\te.g : !pa2va nt!ExAllocatePoolWithTag+5\n");
    ShowMessages("\t\te.g : !pa2va @rax+5\n");
    ShowMessages("\t\te.g : !pa2va fffff801deadbeef\n");
    ShowMessages("\t\te.g : !pa2va fffff801deadbeef pid 0xc8\n");
}








VOID
CommandPa2va(vector<string> SplittedCommand, string Command)
{
    BOOL                              Status;
    ULONG                             ReturnedLength;
    UINT64                            TargetPa;
    UINT32                            Pid            = 0;
    DEBUGGER_VA2PA_AND_PA2VA_COMMANDS AddressDetails = {0};
    vector<string>                    SplittedCommandCaseSensitive {Split(Command, ' ')};

    if (SplittedCommand.size() == 1 || SplittedCommand.size() >= 5 ||
        SplittedCommand.size() == 3)
    {
        ShowMessages("incorrect use of '!pa2va'\n\n");
        CommandPa2vaHelp();
        return;
    }

    
    
    
    
    if (g_ActiveProcessDebuggingState.IsActive)
    {
        Pid = g_ActiveProcessDebuggingState.ProcessId;
    }

    if (SplittedCommand.size() == 2)
    {
        
        
        
        if (!SymbolConvertNameOrExprToAddress(SplittedCommandCaseSensitive.at(1), &TargetPa))
        {
            
            
            
            ShowMessages("err, couldn't resolve error at '%s'\n",
                         SplittedCommandCaseSensitive.at(1).c_str());
            return;
        }
    }
    else
    {
        
        
        
        if (!SplittedCommand.at(1).compare("pid"))
        {
            if (!ConvertStringToUInt32(SplittedCommand.at(2), &Pid))
            {
                ShowMessages("incorrect address, please enter a valid process id\n");
                return;
            }
            if (!SymbolConvertNameOrExprToAddress(SplittedCommandCaseSensitive.at(3), &TargetPa))
            {
                
                
                
                ShowMessages("err, couldn't resolve error at '%s'\n",
                             SplittedCommandCaseSensitive.at(3).c_str());
                return;
            }
        }
        else if (!SplittedCommand.at(2).compare("pid"))
        {
            if (!SymbolConvertNameOrExprToAddress(SplittedCommandCaseSensitive.at(1), &TargetPa))
            {
                
                
                
                ShowMessages("err, couldn't resolve error at '%s'\n",
                             SplittedCommandCaseSensitive.at(1).c_str());

                return;
            }
            if (!ConvertStringToUInt32(SplittedCommand.at(3), &Pid))
            {
                ShowMessages("incorrect address, please enter a valid process id\n");
                return;
            }
        }
        else
        {
            ShowMessages("incorrect use of '!pa2va'\n\n");
            CommandPa2vaHelp();
            return;
        }
    }

    
    
    
    
    AddressDetails.PhysicalAddress    = TargetPa;
    AddressDetails.ProcessId          = Pid; 
    AddressDetails.IsVirtual2Physical = FALSE;

    if (g_IsSerialConnectedToRemoteDebuggee)
    {
        
        
        
        if (Pid != 0)
        {
            ShowMessages("err, you cannot specify 'pid' in the debugger mode\n");
            return;
        }

        
        
        

        KdSendVa2paAndPa2vaPacketToDebuggee(&AddressDetails);
    }
    else
    {
        AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);

        if (Pid == 0)
        {
            Pid                      = GetCurrentProcessId();
            AddressDetails.ProcessId = Pid;
        }

        
        
        
        Status = DeviceIoControl(
            g_DeviceHandle,                           
            IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS,  
            &AddressDetails,                          
            SIZEOF_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS, 
            &AddressDetails,                          
            SIZEOF_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS, 
                                                      
            &ReturnedLength,                          
            NULL                                      
        );

        if (!Status)
        {
            ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
            return;
        }

        if (AddressDetails.KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFULL)
        {
            
            
            
            ShowMessages("%llx\n", AddressDetails.VirtualAddress);
        }
        else
        {
            
            
            
            ShowErrorMessage(AddressDetails.KernelStatus);
        }
    }
}
