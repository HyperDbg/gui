package ept


type (
Vpid interface{
VpidInvvpid()(ok bool)//col:38
VpidInvvpidIndividualAddress()(ok bool)//col:54
VpidInvvpidSingleContext()(ok bool)//col:68
VpidInvvpidAllContext()(ok bool)//col:80
VpidInvvpidSingleContextRetainingGlobals()(ok bool)//col:94
}







































)

func NewVpid() { return & vpid{} }

func (v *vpid)VpidInvvpid()(ok bool){//col:38














return true
}








































func (v *vpid)VpidInvvpidIndividualAddress()(ok bool){//col:54





return true
}








































func (v *vpid)VpidInvvpidSingleContext()(ok bool){//col:68





return true
}








































func (v *vpid)VpidInvvpidAllContext()(ok bool){//col:80




return true
}








































func (v *vpid)VpidInvvpidSingleContextRetainingGlobals()(ok bool){//col:94





return true
}










































