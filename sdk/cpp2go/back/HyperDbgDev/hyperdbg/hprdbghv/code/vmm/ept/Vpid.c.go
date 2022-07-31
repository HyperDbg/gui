package ept
//back\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\ept\Vpid.c.back

type (
Vpid interface{
VpidInvvpid()(ok bool)//col:38
VpidInvvpidIndividualAddress()(ok bool)//col:53
VpidInvvpidSingleContext()(ok bool)//col:66
VpidInvvpidAllContext()(ok bool)//col:77
VpidInvvpidSingleContextRetainingGlobals()(ok bool)//col:90
}
)

func NewVpid() { return & vpid{} }

func (v *vpid)VpidInvvpid()(ok bool){//col:38
/*VpidInvvpid(INVVPID_TYPE Type, INVVPID_DESCRIPTOR * Descriptor)
{
    INVVPID_DESCRIPTOR * TargetDescriptor = NULL;
    INVVPID_DESCRIPTOR   ZeroDescriptor   = {0};
    if (!Descriptor)
    {
        TargetDescriptor = &ZeroDescriptor;
    }
    else
    {
        TargetDescriptor = Descriptor;
    }
    AsmInvvpid(Type, TargetDescriptor);
}*/
return true
}

func (v *vpid)VpidInvvpidIndividualAddress()(ok bool){//col:53
/*VpidInvvpidIndividualAddress(UINT16 Vpid, UINT64 LinearAddress)
{
    INVVPID_DESCRIPTOR Descriptor = {Vpid, 0, 0, LinearAddress};
    VpidInvvpid(InvvpidIndividualAddress, &Descriptor);
}*/
return true
}

func (v *vpid)VpidInvvpidSingleContext()(ok bool){//col:66
/*VpidInvvpidSingleContext(UINT16 Vpid)
{
    INVVPID_DESCRIPTOR Descriptor = {Vpid, 0, 0, 0};
    VpidInvvpid(InvvpidSingleContext, &Descriptor);
}*/
return true
}

func (v *vpid)VpidInvvpidAllContext()(ok bool){//col:77
/*VpidInvvpidAllContext()
{
    VpidInvvpid(InvvpidAllContext, NULL);
}*/
return true
}

func (v *vpid)VpidInvvpidSingleContextRetainingGlobals()(ok bool){//col:90
/*VpidInvvpidSingleContextRetainingGlobals(UINT16 Vpid)
{
    INVVPID_DESCRIPTOR Descriptor = {Vpid, 0, 0, 0};
    VpidInvvpid(InvvpidSingleContextRetainingGlobals, &Descriptor);
}*/
return true
}



