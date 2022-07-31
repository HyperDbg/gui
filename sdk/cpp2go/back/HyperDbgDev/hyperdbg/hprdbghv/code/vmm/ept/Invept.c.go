package ept
//back\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\ept\Invept.c.back

type (
Invept interface{
EptInvept()(ok bool)//col:31
EptInveptSingleContext()(ok bool)//col:46
EptInveptAllContexts()(ok bool)//col:57
}
)

func NewInvept() { return & invept{} }

func (i *invept)EptInvept()(ok bool){//col:31
/*EptInvept(_In_ UINT32 Type, _In_ INVEPT_DESCRIPTOR * Descriptor)
{
    if (!Descriptor)
    {
        INVEPT_DESCRIPTOR ZeroDescriptor = {0};
        Descriptor                       = &ZeroDescriptor;
    }
    return AsmInvept(Type, Descriptor);
}*/
return true
}

func (i *invept)EptInveptSingleContext()(ok bool){//col:46
/*EptInveptSingleContext(_In_ UINT64 EptPointer)
{
    INVEPT_DESCRIPTOR Descriptor = {0};
    Descriptor.EptPointer        = EptPointer;
    Descriptor.Reserved          = 0;
    return EptInvept(InveptSingleContext, &Descriptor);
}*/
return true
}

func (i *invept)EptInveptAllContexts()(ok bool){//col:57
/*EptInveptAllContexts()
{
    return EptInvept(InveptAllContext, NULL);
}*/
return true
}



