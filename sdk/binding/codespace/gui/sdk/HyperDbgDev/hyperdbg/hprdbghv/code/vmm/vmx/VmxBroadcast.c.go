package vmx

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\vmx\VmxBroadcast.c.back

type (
	VmxBroadcast interface {
		VmxBroadcastHandleNmiCallback() (ok bool) //col:13
	}
	vmxBroadcast struct{}
)

func NewVmxBroadcast() VmxBroadcast { return &vmxBroadcast{} }

func (v *vmxBroadcast) VmxBroadcastHandleNmiCallback() (ok bool) { //col:13
	/*
	   VmxBroadcastHandleNmiCallback(PVOID Context, BOOLEAN Handled)

	   	{
	   	    ULONG CurrentCoreIndex;
	   	    CurrentCoreIndex = KeGetCurrentProcessorNumber();
	   	    if (!VmxBroadcastNmiHandler(CurrentCoreIndex, NULL, TRUE))
	   	    {
	   	        return Handled;
	   	    }
	   	    else
	   	    {
	   	        return TRUE;
	   	    }
	   	}
	*/
	return true
}

