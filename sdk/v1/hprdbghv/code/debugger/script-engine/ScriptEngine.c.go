package script_engine

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\script-engine\ScriptEngine.c.back

type (
	ScriptEngine interface {
		ScriptEngineWrapperGetInstructionPointer() (ok bool) //col:12
	}
	scriptEngine struct{}
)

func NewScriptEngine() ScriptEngine { return &scriptEngine{} }

func (s *scriptEngine) ScriptEngineWrapperGetInstructionPointer() (ok bool) { //col:12
	/*
	   ScriptEngineWrapperGetInstructionPointer()

	   	{
	   	    UINT64 GuestRip = NULL;
	   	    ULONG  CurrentProcessorIndex;
	   	    CurrentProcessorIndex = KeGetCurrentProcessorNumber();
	   	    if (g_GuestState[CurrentProcessorIndex].IsOnVmxRootMode)
	   	    {
	   	        __vmx_vmread(VMCS_GUEST_RIP, &GuestRip);

	   ScriptEngineWrapperGetAddressOfReservedBuffer(PDEBUGGER_EVENT_ACTION Action)

	   	{
	   	    return Action->RequestedBuffer.RequstBufferAddress;
	   	}
	*/
	return true
}
