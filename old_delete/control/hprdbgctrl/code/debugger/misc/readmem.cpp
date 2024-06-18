#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
VOID HyperDbgReadMemoryAndDisassemble(DEBUGGER_SHOW_MEMORY_STYLE   Style,
                                 UINT64                       Address,
                                 DEBUGGER_READ_MEMORY_TYPE    MemoryType,
                                 DEBUGGER_READ_READING_TYPE   ReadingType,
                                 UINT32                       Pid,
                                 UINT32                       Size,
                                 PDEBUGGER_DT_COMMAND_OPTIONS DtDetails){
    BOOL                 Status;
    ULONG                ReturnedLength;
    DEBUGGER_READ_MEMORY ReadMem = {0};
    UINT32               SizeOfTargetBuffer;
    ReadMem.Address     = Address;
    ReadMem.Pid         = Pid;
    ReadMem.Size        = Size;
    ReadMem.MemoryType  = MemoryType;
    ReadMem.ReadingType = ReadingType;
    ReadMem.Style       = Style;
    ReadMem.DtDetails   = DtDetails;
    if (Style == DEBUGGER_SHOW_COMMAND_DISASSEMBLE64 ||
        Style == DEBUGGER_SHOW_COMMAND_DISASSEMBLE32){
        ReadMem.IsForDisasm = TRUE;
    }
    else{
        ReadMem.IsForDisasm = FALSE;
    }
    if (g_IsSerialConnectedToRemoteDebuggee){
        KdSendReadMemoryPacketToDebuggee(&ReadMem);
        return;
    }
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
    SizeOfTargetBuffer           = (Size * sizeof(CHAR)) + sizeof(DEBUGGER_READ_MEMORY);
    unsigned char * OutputBuffer = (unsigned char *)malloc(SizeOfTargetBuffer);
    ZeroMemory(OutputBuffer, SizeOfTargetBuffer);
    Status = DeviceIoControl(g_DeviceHandle,              
                             IOCTL_DEBUGGER_READ_MEMORY,  
                             &ReadMem,                    
                             SIZEOF_DEBUGGER_READ_MEMORY, 
                             OutputBuffer,                
                             SizeOfTargetBuffer,          
                             &ReturnedLength,             
                             NULL                         
    );
    if (!Status){
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        std::free(OutputBuffer);
        return;
    }
    if (((PDEBUGGER_READ_MEMORY)OutputBuffer)->KernelStatus != DEBUGGER_OPERATION_WAS_SUCCESSFUL){
        std::free(OutputBuffer);
        ShowErrorMessage(((PDEBUGGER_READ_MEMORY)OutputBuffer)->KernelStatus);
        return;
    }
    else{
        ReturnedLength -= SIZEOF_DEBUGGER_READ_MEMORY;
    }
    switch (Style){
    case DEBUGGER_SHOW_COMMAND_DT:
        if (Size == ReturnedLength){
            ScriptEngineShowDataBasedOnSymbolTypesWrapper(DtDetails->TypeName,
                                                          Address,
                                                          FALSE,
                                                          ((unsigned char *)OutputBuffer) + sizeof(DEBUGGER_READ_MEMORY),
                                                          DtDetails->AdditionalParameters);
        }
        else if (ReturnedLength == 0){
            ShowMessages("err, invalid address");
        }
        else{
            ShowMessages("err, invalid address or memory is smaller than the structure size");
        }
        break;
    case DEBUGGER_SHOW_COMMAND_DB:
        ShowMemoryCommandDB(
            ((unsigned char *)OutputBuffer) + sizeof(DEBUGGER_READ_MEMORY),
            Size,
            Address,
            MemoryType,
            ReturnedLength);
        break;
    case DEBUGGER_SHOW_COMMAND_DC:
        ShowMemoryCommandDC(
            ((unsigned char *)OutputBuffer) + sizeof(DEBUGGER_READ_MEMORY),
            Size,
            Address,
            MemoryType,
            ReturnedLength);
        break;
    case DEBUGGER_SHOW_COMMAND_DD:
        ShowMemoryCommandDD(
            ((unsigned char *)OutputBuffer) + sizeof(DEBUGGER_READ_MEMORY),
            Size,
            Address,
            MemoryType,
            ReturnedLength);
        break;
    case DEBUGGER_SHOW_COMMAND_DQ:
        ShowMemoryCommandDQ(
            ((unsigned char *)OutputBuffer) + sizeof(DEBUGGER_READ_MEMORY),
            Size,
            Address,
            MemoryType,
            ReturnedLength);
        break;
    case DEBUGGER_SHOW_COMMAND_DUMP:
        CommandDumpSaveIntoFile(((unsigned char *)OutputBuffer) + sizeof(DEBUGGER_READ_MEMORY), Size);
        break;
    case DEBUGGER_SHOW_COMMAND_DISASSEMBLE64:
        if (((PDEBUGGER_READ_MEMORY)OutputBuffer)->Is32BitAddress == TRUE &&
            MemoryType == DEBUGGER_READ_VIRTUAL_ADDRESS){
            ShowMessages("the target address seems to be located in a 32-bit program, if so, "
                         "please consider using the 'u32' instead to utilize the 32-bit disassembler\n");
        }
        HyperDbgDisassembler64(
            ((unsigned char *)OutputBuffer) + sizeof(DEBUGGER_READ_MEMORY),
            Address,
            ReturnedLength,
            0,
            FALSE,
            NULL);
        break;
    case DEBUGGER_SHOW_COMMAND_DISASSEMBLE32:
        if (((PDEBUGGER_READ_MEMORY)OutputBuffer)->Is32BitAddress == FALSE &&
            MemoryType == DEBUGGER_READ_VIRTUAL_ADDRESS){
            ShowMessages("the target address seems to be located in a 64-bit program, if so, "
                         "please consider using the 'u' instead to utilize the 64-bit disassembler\n");
        }
        HyperDbgDisassembler32(
            ((unsigned char *)OutputBuffer) + sizeof(DEBUGGER_READ_MEMORY),
            Address,
            ReturnedLength,
            0,
            FALSE,
            NULL);
        break;
    }
    std::free(OutputBuffer);
    ShowMessages("\n");
}
void ShowMemoryCommandDB(unsigned char * OutputBuffer, UINT32 Size, UINT64 Address, DEBUGGER_READ_MEMORY_TYPE MemoryType, UINT64 Length){
    unsigned int Character;
    for (UINT32 i = 0; i < Size; i += 16){
        if (MemoryType == DEBUGGER_READ_PHYSICAL_ADDRESS){
            ShowMessages("#\t");
        }
        ShowMessages("%s  ", SeparateTo64BitValue((UINT64)(Address + i)).c_str());
        for (size_t j = 0; j < 16; j++){
            if (i + j >= Length){
                ShowMessages("?? ");
            }
            else{
                ShowMessages("%02X ", OutputBuffer[i + j]);
            }
        }
        ShowMessages(" ");
        for (size_t j = 0; j < 16; j++){
            Character = (OutputBuffer[i + j]);
            if (isprint(Character)){
                ShowMessages("%c", Character);
            }
            else{
                ShowMessages(".");
            }
        }
        ShowMessages("\n");
    }
}
void ShowMemoryCommandDC(unsigned char * OutputBuffer, UINT32 Size, UINT64 Address, DEBUGGER_READ_MEMORY_TYPE MemoryType, UINT64 Length){
    unsigned int Character;
    for (UINT32 i = 0; i < Size; i += 16){
        if (MemoryType == DEBUGGER_READ_PHYSICAL_ADDRESS){
            ShowMessages("#\t");
        }
        ShowMessages("%s  ", SeparateTo64BitValue((UINT64)(Address + i)).c_str());
        for (size_t j = 0; j < 16; j += 4){
            if (i + j >= Length){
                ShowMessages("???????? ");
            }
            else{
                UINT32 OutputBufferVar = *((UINT32 *)&OutputBuffer[i + j]);
                ShowMessages("%08X ", OutputBufferVar);
            }
        }
        ShowMessages(" ");
        for (size_t j = 0; j < 16; j++){
            Character = (OutputBuffer[i + j]);
            if (isprint(Character)){
                ShowMessages("%c", Character);
            }
            else{
                ShowMessages(".");
            }
        }
        ShowMessages("\n");
    }
}
void ShowMemoryCommandDD(unsigned char * OutputBuffer, UINT32 Size, UINT64 Address, DEBUGGER_READ_MEMORY_TYPE MemoryType, UINT64 Length){
    for (UINT32 i = 0; i < Size; i += 16){
        if (MemoryType == DEBUGGER_READ_PHYSICAL_ADDRESS){
            ShowMessages("#\t");
        }
        ShowMessages("%s  ", SeparateTo64BitValue((UINT64)(Address + i)).c_str());
        for (size_t j = 0; j < 16; j += 4){
            if (i + j >= Length){
                ShowMessages("???????? ");
            }
            else{
                UINT32 OutputBufferVar = *((UINT32 *)&OutputBuffer[i + j]);
                ShowMessages("%08X ", OutputBufferVar);
            }
        }
        ShowMessages("\n");
    }
}
void ShowMemoryCommandDQ(unsigned char * OutputBuffer, UINT32 Size, UINT64 Address, DEBUGGER_READ_MEMORY_TYPE MemoryType, UINT64 Length){
    for (UINT32 i = 0; i < Size; i += 16){
        if (MemoryType == DEBUGGER_READ_PHYSICAL_ADDRESS){
            ShowMessages("#\t");
        }
        ShowMessages("%s  ", SeparateTo64BitValue((UINT64)(Address + i)).c_str());
        for (size_t j = 0; j < 16; j += 8){
            if (i + j >= Length){
                ShowMessages("???????? ");
            }
            else{
                UINT32 OutputBufferVar = *((UINT32 *)&OutputBuffer[i + j + 4]);
                ShowMessages("%08X`", OutputBufferVar);
                OutputBufferVar = *((UINT32 *)&OutputBuffer[i + j]);
                ShowMessages("%08X ", OutputBufferVar);
            }
        }
        ShowMessages("\n");
    }
}
