#include "pch.h"
extern BOOLEAN                  g_IsSerialConnectedToRemoteDebuggee;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;
VOID
CommandPteHelp() {
    ShowMessages("!pte : finds virtual addresses of different paging-levels.\n\n");
    ShowMessages("syntax : \t!pte [VirtualAddress (hex)] [pid ProcessId (hex)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : !pte nt!ExAllocatePoolWithTag\n");
    ShowMessages("\t\te.g : !pte nt!ExAllocatePoolWithTag+5\n");
    ShowMessages("\t\te.g : !pte fffff801deadbeef\n");
    ShowMessages("\t\te.g : !pte 0x400021000 pid 1c0\n");
}

VOID
CommandPteShowResults(UINT64 TargetVa, PDEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS PteRead) {
    ShowMessages("VA %llx\n", TargetVa);
    ShowMessages("PML4E (PXE) at %016llx\tcontains %016llx\nPDPT (PPE) at "
                 "%016llx\tcontains "
                 "%016llx\nPDE at %016llx\tcontains %016llx\n",
                 PteRead->Pml4eVirtualAddress,
                 PteRead->Pml4eValue,
                 PteRead->PdpteVirtualAddress,
                 PteRead->PdpteValue,
                 PteRead->PdeVirtualAddress,
                 PteRead->PdeValue);
    if (PteRead->PdeVirtualAddress == PteRead->PteVirtualAddress) {
        ShowMessages("PDE is a large page, so it doesn't have a PTE\n");
    } else {
        ShowMessages("PTE at %016llx\tcontains %016llx\n",
                     PteRead->PteVirtualAddress,
                     PteRead->PteValue);
    }
}

VOID
CommandPte(vector<string> SplittedCommand, string Command) {
    BOOL                                     Status;
    ULONG                                    ReturnedLength;
    UINT64                                   TargetVa;
    UINT32                                   Pid            = 0;
    DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS AddressDetails = {0};
    vector<string>                           SplittedCommandCaseSensitive {Split(Command, ' ')};
    if (SplittedCommand.size() == 1 || SplittedCommand.size() >= 5 ||
        SplittedCommand.size() == 3) {
        ShowMessages("incorrect use of '!pte'\n\n");
        CommandPteHelp();
        return;
    }
    if (g_ActiveProcessDebuggingState.IsActive) {
        Pid = g_ActiveProcessDebuggingState.ProcessId;
    }
    if (SplittedCommand.size() == 2) {
        if (!SymbolConvertNameOrExprToAddress(SplittedCommandCaseSensitive.at(1), &TargetVa)) {
            ShowMessages("err, couldn't resolve error at '%s'\n",
                         SplittedCommandCaseSensitive.at(1).c_str());
            return;
        }
    } else {
        if (!SplittedCommand.at(1).compare("pid")) {
            if (!ConvertStringToUInt32(SplittedCommand.at(2), &Pid)) {
                ShowMessages("incorrect address, please enter a valid process id\n");
                return;
            }
            if (!SymbolConvertNameOrExprToAddress(SplittedCommandCaseSensitive.at(3), &TargetVa)) {
                ShowMessages("err, couldn't resolve error at '%s'\n",
                             SplittedCommandCaseSensitive.at(3).c_str());
                return;
            }
        } else if (!SplittedCommand.at(2).compare("pid")) {
            if (!SymbolConvertNameOrExprToAddress(SplittedCommandCaseSensitive.at(1), &TargetVa)) {
                ShowMessages("err, couldn't resolve error at '%s'\n\n",
                             SplittedCommandCaseSensitive.at(1).c_str());
                return;
            }
            if (!ConvertStringToUInt32(SplittedCommand.at(3), &Pid)) {
                ShowMessages("incorrect address, please enter a valid process id\n");
                return;
            }
        } else {
            ShowMessages("incorrect use of '!pte'\n\n");
            CommandPteHelp();
            return;
        }
    }
    AddressDetails.VirtualAddress = TargetVa;
    AddressDetails.ProcessId      = Pid; // null in debugger mode
    if (g_IsSerialConnectedToRemoteDebuggee) {
        if (Pid != 0) {
            ShowMessages("err, you cannot specify 'pid' in the debugger mode\n");
            return;
        }
        KdSendPtePacketToDebuggee(&AddressDetails);
    } else {
        AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
        if (Pid == 0) {
            Pid                      = GetCurrentProcessId();
            AddressDetails.ProcessId = Pid;
        }
        Status = DeviceIoControl(
            g_DeviceHandle,                                  // Handle to device
            IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS,  // IO Control code
            &AddressDetails,                                 // Input Buffer to driver.
            SIZEOF_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS, // Input buffer length
            &AddressDetails,                                 // Output Buffer from driver.
            SIZEOF_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS, // Length of output
            &ReturnedLength,                                 // Bytes placed in buffer.
            NULL                                             // synchronous call
        );
        if (!Status) {
            ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
            return;
        }
        if (AddressDetails.KernelStatus != DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
            ShowErrorMessage(AddressDetails.KernelStatus);
            return;
        }
        CommandPteShowResults(TargetVa, &AddressDetails);
    }
}
