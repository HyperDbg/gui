package ept
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\ept\Invept.c.back

type (
Invept interface{
EptInvept()(ok bool)//col:9
EptInveptSingleContext()(ok bool)//col:16
EptInveptAllContexts()(ok bool)//col:20
}
invept struct{}
)

func NewInvept()Invept{ return & invept{} }

func (i *invept)EptInvept()(ok bool){//col:9
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

func (i *invept)EptInveptSingleContext()(ok bool){//col:16
/*EptInveptSingleContext(_In_ UINT64 EptPointer)
{
    INVEPT_DESCRIPTOR Descriptor = {0};
    Descriptor.EptPointer        = EptPointer;
    Descriptor.Reserved          = 0;
    return EptInvept(InveptSingleContext, &Descriptor);
}*/
return true
}

func (i *invept)EptInveptAllContexts()(ok bool){//col:20
/*EptInveptAllContexts()
{
    return EptInvept(InveptAllContext, NULL);
}*/
return true
}



