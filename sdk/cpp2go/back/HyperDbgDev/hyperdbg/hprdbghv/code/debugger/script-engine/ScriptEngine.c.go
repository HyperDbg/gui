package script-engine
//back\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\script-engine\ScriptEngine.c.back

type (
ScriptEngine interface{
ScriptEngineWrapperGetInstructionPointer()(ok bool)//col:15
ScriptEngineWrapperGetAddressOfReservedBuffer()(ok bool)//col:19
}
)

func NewScriptEngine() { return & scriptEngine{} }

func (s *scriptEngine)ScriptEngineWrapperGetInstructionPointer()(ok bool){//col:15
/*ScriptEngineWrapperGetInstructionPointer()
{
    UINT64 GuestRip = NULL;
    ULONG  CurrentProcessorIndex;
    CurrentProcessorIndex = KeGetCurrentProcessorNumber();
    if (g_GuestState[CurrentProcessorIndex].IsOnVmxRootMode)
    {
        __vmx_vmread(VMCS_GUEST_RIP, &GuestRip);
        return GuestRip;
    }
    else
    {
        return NULL;
    }
}*/
return true
}

func (s *scriptEngine)ScriptEngineWrapperGetAddressOfReservedBuffer()(ok bool){//col:19
/*ScriptEngineWrapperGetAddressOfReservedBuffer(PDEBUGGER_EVENT_ACTION Action)
{
    return Action->RequestedBuffer.RequstBufferAddress;
}*/
return true
}



