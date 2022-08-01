package misc
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\misc\callstack.cpp.back

type (
Callstack interface{
CallstackReturnAddressToCallingAddress()(ok bool)//col:74
CallstackShowFrames()(ok bool)//col:162
}
)

func NewCallstack() { return & callstack{} }

func (c *callstack)CallstackReturnAddressToCallingAddress()(ok bool){//col:74
/*CallstackReturnAddressToCallingAddress(UCHAR * ReturnAddress, PUINT32 IndexOfCallFromReturnAddress)
{
    if (ReturnAddress[-7] == 0x9A)
    {
        *IndexOfCallFromReturnAddress = 7;
        return TRUE;
    }
    else if (ReturnAddress[-5] == 0xE8)
    {
        *IndexOfCallFromReturnAddress = 5;
        return TRUE;
    }
    else
    {
        const unsigned char RmMask = 0xF8;
        if (ReturnAddress[-7] == 0xFF &&
            (ReturnAddress[-6] == 0x94 || ReturnAddress[-6] == 0x9C))
        {
            *IndexOfCallFromReturnAddress = 7;
            return TRUE;
        }
        else if (ReturnAddress[-6] == 0xFF &&
                 ((ReturnAddress[-5] & RmMask) == 0x90 || (ReturnAddress[-5] & RmMask) == 0x98) &&
                 (ReturnAddress[-5] != 0x94 && ReturnAddress[-5] != 0x9C))
        {
            *IndexOfCallFromReturnAddress = 6;
            return TRUE;
        }
        else if (ReturnAddress[-6] == 0xFF &&
                 (ReturnAddress[-5] == 0x15 || ReturnAddress[-5] == 0x1D))
        {
            *IndexOfCallFromReturnAddress = 6;
            return TRUE;
        }
        else if (ReturnAddress[-4] == 0xFF &&
                 (ReturnAddress[-3] == 0x54 || ReturnAddress[-3] == 0x5C))
        {
            *IndexOfCallFromReturnAddress = 4;
            return TRUE;
        }
        else if (ReturnAddress[-3] == 0xFF &&
                 ((ReturnAddress[-2] & RmMask) == 0x50 || (ReturnAddress[-2] & RmMask) == 0x58) &&
                 (ReturnAddress[-2] != 0x54 && ReturnAddress[-2] != 0x5C))
        {
            *IndexOfCallFromReturnAddress = 3;
            return TRUE;
        }
        else if (ReturnAddress[-3] == 0xFF &&
                 (ReturnAddress[-2] == 0x14 || ReturnAddress[-2] == 0x1C))
        {
            *IndexOfCallFromReturnAddress = 3;
            return TRUE;
        }
        else if (ReturnAddress[-2] == 0xFF &&
                 ((ReturnAddress[-1] & RmMask) == 0xD0 || (ReturnAddress[-1] & RmMask) == 0xD8))
        {
            *IndexOfCallFromReturnAddress = 2;
            return TRUE;
        }
        else if (ReturnAddress[-2] == 0xFF &&
                 ((ReturnAddress[-1] & RmMask) == 0x10 || (ReturnAddress[-1] & RmMask) == 0x18) &&
                 (ReturnAddress[-1] != 0x14 && ReturnAddress[-1] != 0x15 &&
                  ReturnAddress[-1] != 0x1C && ReturnAddress[-1] != 0x1D))
        {
            *IndexOfCallFromReturnAddress = 2;
            return TRUE;
        }
        else
        {
            return FALSE;
        }
    }
    return FALSE;
}*/
return true
}

func (c *callstack)CallstackShowFrames()(ok bool){//col:162
/*CallstackShowFrames(PDEBUGGER_SINGLE_CALLSTACK_FRAME  CallstackFrames,
                    UINT32                            FrameCount,
                    DEBUGGER_CALLSTACK_DISPLAY_METHOD DisplayMethod,
                    BOOLEAN                           Is32Bit)
{
    UINT32                                                 CallLength;
    UINT64                                                 TargetAddress;
    UINT64                                                 UsedBaseAddress;
    BOOLEAN                                                IsCall = FALSE;
    std::map<UINT64, LOCAL_FUNCTION_DESCRIPTION>::iterator Iterate;
    for (size_t i = 0; i < FrameCount; i++)
    {
        IsCall = FALSE;
        if (CallstackFrames[i].IsValidAddress)
        {
            if (CallstackFrames[i].IsExecutable && CallstackReturnAddressToCallingAddress(
                                                       (unsigned char *)&CallstackFrames[i].InstructionBytesOnRip[MAXIMUM_CALL_INSTR_SIZE],
                                                       &CallLength))
            {
                TargetAddress = CallstackFrames[i].Value - CallLength;
                IsCall = TRUE;
            }
            else
            {
                if (DisplayMethod == DEBUGGER_CALLSTACK_DISPLAY_METHOD_WITHOUT_PARAMS)
                {
                    continue;
                }
                IsCall        = FALSE;
                TargetAddress = CallstackFrames[i].Value;
            }
            ShowMessages("[$+%03x] ", i * (Is32Bit ? sizeof(UINT32) : sizeof(UINT64)));
            if (IsCall)
            {
                if (Is32Bit)
                {
                    ShowMessages("  %08x    (from ", TargetAddress);
                }
                else
                {
                    ShowMessages("  %016llx    (from ", TargetAddress);
                }
            }
            else
            {
                if (Is32Bit)
                {
                    ShowMessages("     %08x (addr ", TargetAddress);
                }
                else
                {
                    ShowMessages("     %016llx (addr ", TargetAddress);
                }
            }
            if (g_AddressConversion)
            {
                if (SymbolShowFunctionNameBasedOnAddress(TargetAddress, &UsedBaseAddress))
                {
                    ShowMessages(" ");
                }
            }
            if (Is32Bit)
            {
                ShowMessages("<%08x>)\n", TargetAddress);
            }
            else
            {
                ShowMessages("<%016llx>)\n", TargetAddress);
            }
        }
        else
        {
            if (DisplayMethod == DEBUGGER_CALLSTACK_DISPLAY_METHOD_WITHOUT_PARAMS)
            {
                continue;
            }
            ShowMessages("[$+%03x] ", i * (Is32Bit ? sizeof(UINT32) : sizeof(UINT64)));
            if (Is32Bit)
            {
                ShowMessages("     %08x\n", CallstackFrames[i].Value);
            }
            else
            {
                ShowMessages("     %016llx\n", CallstackFrames[i].Value);
            }
        }
    }
}*/
return true
}



