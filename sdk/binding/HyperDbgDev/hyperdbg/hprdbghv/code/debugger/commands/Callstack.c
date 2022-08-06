










#include "pch.h"











BOOLEAN
CallstackWalkthroughStack(PDEBUGGER_SINGLE_CALLSTACK_FRAME AddressToSaveFrames,
                          UINT64                           StackBaseAddress,
                          UINT32                           Size,
                          BOOLEAN                          Is32Bit)
{
    UINT32 FrameIndex          = 0;
    UINT16 AddressMode         = 0;
    UINT64 Value               = NULL;
    UINT64 CurrentStackAddress = NULL;

    if (Size == 0)
    {
        return FALSE;
    }

    if (Is32Bit)
    {
        
        
        
        AddressMode = sizeof(UINT32);
        FrameIndex  = Size / AddressMode;
    }
    else
    {
        
        
        
        AddressMode = sizeof(UINT64);
        FrameIndex  = Size / AddressMode;
    }

    
    
    
    for (size_t i = 0; i < FrameIndex; i++)
    {
        
        
        
        CurrentStackAddress = StackBaseAddress + (i * AddressMode);

        if (!CheckMemoryAccessSafety(CurrentStackAddress, AddressMode))
        {
            AddressToSaveFrames[i].IsStackAddressValid = FALSE;

            
            
            
            return FALSE;
        }

        
        
        
        AddressToSaveFrames[i].IsStackAddressValid = TRUE;

        
        
        
        MemoryMapperReadMemorySafeOnTargetProcess(CurrentStackAddress, &Value, AddressMode);

        
        
        
        AddressToSaveFrames[i].Value = Value;

        
        
        
        
        
        
        
        
        if (CheckMemoryAccessSafety(Value, MAXIMUM_CALL_INSTR_SIZE))
        {
            
            
            
            AddressToSaveFrames[i].IsValidAddress = TRUE;

            
            
            
            AddressToSaveFrames[i].IsExecutable = MemoryMapperCheckIfPageIsNxBitSetOnTargetProcess(Value);

            
            
            
            MemoryMapperReadMemorySafeOnTargetProcess(Value - MAXIMUM_CALL_INSTR_SIZE,
                                                      AddressToSaveFrames[i].InstructionBytesOnRip,
                                                      MAXIMUM_CALL_INSTR_SIZE);
        }
    }

    
    
    
    return TRUE;
}
